package main

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"runtime/debug"
	"strings"
	"time"

	"github.com/topxeq/qlang"

	"github.com/topxeq/charlang"
	charex "github.com/topxeq/charlang/stdlib/ex"
	"github.com/topxeq/gox"
	tk "github.com/topxeq/tkc"
	"github.com/topxeq/xie"

	// "tinygo.org/x/bluetooth"

	// full version related start
	_ "github.com/denisenkom/go-mssqldb"
	// _ "github.com/godror/godror"
	_ "github.com/sijms/go-ora/v2"

	// full version related end

	_ "github.com/go-sql-driver/mysql"
	// _ "github.com/mattn/go-sqlite3"
	_ "github.com/glebarez/go-sqlite"
	_ "github.com/jackc/pgx/v5/stdlib"
)

// for goxn
var muxG *http.ServeMux
var portG = ":80"
var sslPortG = ":443"
var basePathG = "."
var webPathG = "."
var certPathG = "."
var verboseG = false

func doWms(res http.ResponseWriter, req *http.Request) {
	if res != nil {
		res.Header().Set("Access-Control-Allow-Origin", "*")
		res.Header().Set("Access-Control-Allow-Headers", "*")
		res.Header().Set("Content-Type", "text/html; charset=utf-8")
	}

	if req != nil {
		req.ParseForm()
		req.ParseMultipartForm(100000000)
	}

	reqT := tk.GetFormValueWithDefaultValue(req, "wms", "")
	if verboseG {
		tk.Pl("RequestURI: %v", req.RequestURI)
	}

	if reqT == "" {
		if tk.StartsWith(req.RequestURI, "/wms") {
			reqT = req.RequestURI[4:]
		}
	}

	tmps := tk.Split(reqT, "?")
	if len(tmps) > 1 {
		reqT = tmps[0]
	}

	if tk.StartsWith(reqT, "/") {
		reqT = reqT[1:]
	}

	var paraMapT map[string]string
	var errT error

	vo := tk.GetFormValueWithDefaultValue(req, "vo", "")

	if vo == "" {
		paraMapT = tk.FormToMap(req.Form)
	} else {
		paraMapT, errT = tk.MSSFromJSON(vo)

		if errT != nil {
			res.Write([]byte(tk.ErrStrf("操作失败：%v", "invalid vo format")))
			return
		}
	}

	if verboseG {
		tk.Pl("[%v] REQ: %#v (%#v)", tk.GetNowTimeStringFormal(), reqT, paraMapT)
	}

	toWriteT := ""

	fileNameT := reqT

	if !tk.EndsWith(fileNameT, ".gox") {
		fileNameT += ".gox"
	}

	fcT := tk.LoadStringFromFile(filepath.Join(basePathG, fileNameT))
	if tk.IsErrStr(fcT) {
		res.Write([]byte(tk.ErrStrf("操作失败：%v", tk.GetErrStr(fcT))))
		return
	}

	// paraMapT["_reqHost"] = req.Host
	// paraMapT["_reqInfo"] = fmt.Sprintf("%#v", req)

	toWriteT, errT = gox.RunScriptOnHttp(fcT, res, req, paraMapT["input"], nil, paraMapT, "-base="+basePathG)

	if errT != nil {
		res.Header().Set("Content-Type", "text/html; charset=utf-8")

		errStrT := tk.ErrStrf("操作失败：%v", errT)
		tk.Pln(errStrT)
		res.Write([]byte(errStrT))
		return
	}

	if toWriteT == "TX_END_RESPONSE_XT" {
		return
	}

	res.Header().Set("Content-Type", "text/html; charset=utf-8")

	res.Write([]byte(toWriteT))
}

func replaceHtml(strA string, mapA map[string]string) string {
	if mapA == nil {
		return strA
	}

	for k, v := range mapA {
		strA = tk.Replace(strA, "TX_"+k+"_XT", v)
	}

	return strA
}

func genFailCompact(titleA, msgA string, optsA ...string) string {
	// mapT := map[string]string{
	// 	"msgTitle":    titleA,
	// 	"msg":         msgA,
	// 	"subMsg":      "",
	// 	"actionTitle": "返回",
	// 	"actionHref":  "javascript:history.back();",
	// }

	// var fileNameT = "fail.html"

	// if tk.IfSwitchExists(optsA, "-compact") {
	// 	fileNameT = "failcompact.html"
	// }

	// tmplT := tk.LoadStringFromFile(filepath.Join(basePathG, "tmpl", fileNameT))

	// tmplT = replaceHtml(tmplT, mapT)

	tmplT := tk.ErrStrf("%v: %v", titleA, msgA)

	return tmplT
}

// do mix server

func doMs(res http.ResponseWriter, req *http.Request) {
	if res != nil {
		res.Header().Set("Access-Control-Allow-Origin", "*")
		res.Header().Set("Access-Control-Allow-Headers", "*")
		res.Header().Set("Content-Type", "text/html; charset=utf-8")
	}

	var errT error

	if req != nil {
		req.ParseForm()
		req.ParseMultipartForm(100000000)
		
//		if errT != nil {
//			tk.Pl("failed to parse multipart form: %v", errT)
//		}
	}

	reqT := tk.GetFormValueWithDefaultValue(req, "ms", "")
	if verboseG {
		tk.Pl("RequestURI: %v", req.RequestURI)
	}

	if reqT == "" {
		if tk.StartsWith(req.RequestURI, "/ms") {
			reqT = req.RequestURI[3:]
		}
	}

	tmps := tk.Split(reqT, "?")
	if len(tmps) > 1 {
		reqT = tmps[0]
	}

	if tk.StartsWith(reqT, "/") {
		reqT = reqT[1:]
	}

	var paraMapT map[string]string

	vo := tk.GetFormValueWithDefaultValue(req, "vo", "")

	if vo == "" {
		paraMapT = tk.FormToMap(req.Form)
	} else {
		paraMapT, errT = tk.MSSFromJSON(vo)

		if errT != nil {
			res.Write([]byte(tk.ErrStrf("action failed: %v", "invalid vo format")))
			return
		}
	}

	if verboseG {
		tk.Pl("[%v] REQ: %#v (%#v)", tk.GetNowTimeStringFormal(), reqT, paraMapT)
	}

	toWriteT := ""

	fileNameT := reqT

	if tk.EndsWith(fileNameT, ".char") || tk.EndsWith(fileNameT, ".xie") || tk.EndsWith(fileNameT, ".gox") {
	} else {
		tmps := fileNameT + ".char"

		if tk.IfFileExists(filepath.Join(basePathG, tmps)) {
			fileNameT = tmps
		} else {
			tmps = fileNameT + ".xie"

			if tk.IfFileExists(filepath.Join(basePathG, tmps)) {
				fileNameT = tmps
			} else {
				tmps = fileNameT + ".goxe"

				if tk.IfFileExists(filepath.Join(basePathG, tmps)) {
					fileNameT = tmps
				} else {
					fileNameT = fileNameT + ".gox"
				}
			}
		}
	}

	fullPathT := filepath.Join(basePathG, fileNameT)

	scriptTypeT := filepath.Ext(fileNameT)

	if verboseG {
		tk.Pl("[%v] file path: %#v", tk.GetNowTimeStringFormal(), fullPathT)
	}

	fcT := tk.LoadStringFromFile(fullPathT)
	if tk.IsErrStr(fcT) {
		res.Write([]byte(tk.ErrStrf("action failed: %v", tk.GetErrStr(fcT))))
		return
	}

	// paraMapT["_reqHost"] = req.Host
	// paraMapT["_reqInfo"] = fmt.Sprintf("%#v", req)

	if scriptTypeT == ".char" {
		toWriteT, errT = charlang.RunScriptOnHttp(fcT, nil, res, req, paraMapT["input"], nil, paraMapT, map[string]interface{}{"scriptPathG": fullPathT, "runModeG": "charms", "basePathG": basePathG}, "-base="+basePathG)

		if errT != nil {
			res.Header().Set("Content-Type", "text/html; charset=utf-8")

			errStrT := tk.ErrStrf("action failed: %v", errT)
			tk.Pln(errStrT)
			res.Write([]byte(errStrT))
			return
		}

		if toWriteT == "TX_END_RESPONSE_XT" {
			return
		}

		res.Header().Set("Content-Type", "text/html; charset=utf-8")

		res.Write([]byte(toWriteT))

		return
	} else if scriptTypeT == ".xie" {
		vmT := xie.NewVMQuick()

		vmT.SetVar(vmT.Running, "argsG", paraMapT)
		vmT.SetVar(vmT.Running, "requestG", req)
		vmT.SetVar(vmT.Running, "responseG", res)
		vmT.SetVar(vmT.Running, "reqNameG", reqT)
		vmT.SetVar(vmT.Running, "basePathG", basePathG)

		// vmT.SetVar("inputG", objA)

		lrs := vmT.Load(vmT.Running, fcT)

		if tk.IsError(lrs) {
			res.Write([]byte(genFailCompact("action failed", lrs.Error(), "-compact")))
			return
		}

		// var argsT []string = tk.JSONToStringArray(tk.GetSwitch(optsA, "-args=", "[]"))

		// if argsT != nil {
		// 	vmT.VarsM["argsG"] = argsT
		// } else {
		// 	vmT.VarsM["argsG"] = []string{}
		// }

		rs := vmT.Run()

		if errT != nil {
			res.Write([]byte(genFailCompact("action failed", errT.Error(), "-compact")))
			return
		}

		if tk.IsErrX(rs) {
			res.Write([]byte(genFailCompact("action failed", tk.GetErrStrX(rs), "-compact")))
			return
		}

		toWriteT = tk.ToStr(rs)

		if toWriteT == "TX_END_RESPONSE_XT" {
			return
		}

		res.Header().Set("Content-Type", "text/html; charset=utf-8")

		res.Write([]byte(toWriteT))

		return
	}

	toWriteT, errT = gox.RunScriptOnHttp(fcT, res, req, paraMapT["input"], nil, paraMapT, "-base="+basePathG)

	if errT != nil {
		res.Header().Set("Content-Type", "text/html; charset=utf-8")

		errStrT := tk.ErrStrf("action failed: %v", errT)
		tk.Pln(errStrT)
		res.Write([]byte(errStrT))
		return
	}

	if toWriteT == "TX_END_RESPONSE_XT" {
		return
	}

	res.Header().Set("Content-Type", "text/html; charset=utf-8")

	res.Write([]byte(toWriteT))

}

// charlang server
func doCharms(res http.ResponseWriter, req *http.Request) {
	if res != nil {
		res.Header().Set("Access-Control-Allow-Origin", "*")
		res.Header().Set("Access-Control-Allow-Headers", "*")
		res.Header().Set("Content-Type", "text/html; charset=utf-8")
	}

	if req != nil {
		req.ParseForm()
		req.ParseMultipartForm(100000000)
	}

	reqT := tk.GetFormValueWithDefaultValue(req, "charms", "")
	if verboseG {
		tk.Pl("RequestURI: %v", req.RequestURI)
	}

	if reqT == "" {
		if tk.StartsWith(req.RequestURI, "/charms") {
			reqT = req.RequestURI[7:]
		}
	}

	tmps := tk.Split(reqT, "?")
	if len(tmps) > 1 {
		reqT = tmps[0]
	}

	if tk.StartsWith(reqT, "/") {
		reqT = reqT[1:]
	}

	var paraMapT map[string]string
	var errT error

	vo := tk.GetFormValueWithDefaultValue(req, "vo", "")

	if vo == "" {
		paraMapT = tk.FormToMap(req.Form)
	} else {
		paraMapT, errT = tk.MSSFromJSON(vo)

		if errT != nil {
			res.Write([]byte(tk.ErrStrf("action failed: %v", "invalid vo format")))
			return
		}
	}

	if verboseG {
		tk.Pl("[%v] REQ: %#v (%#v)", tk.GetNowTimeStringFormal(), reqT, paraMapT)
	}

	toWriteT := ""

	fileNameT := reqT

	if !tk.EndsWith(fileNameT, ".char") {
		fileNameT += ".char"
	}

	fullPathT := filepath.Join(basePathG, fileNameT)

	if verboseG {
		tk.Pl("[%v] file path: %#v", tk.GetNowTimeStringFormal(), fullPathT)
	}

	fcT := tk.LoadStringFromFile(fullPathT)
	if tk.IsErrStr(fcT) {
		res.Write([]byte(tk.ErrStrf("action failed: %v", tk.GetErrStr(fcT))))
		return
	}

	// paraMapT["_reqHost"] = req.Host
	// paraMapT["_reqInfo"] = fmt.Sprintf("%#v", req)

	toWriteT, errT = charlang.RunScriptOnHttp(fcT, nil, res, req, paraMapT["input"], nil, paraMapT, map[string]interface{}{"scriptPathG": fullPathT, "runModeG": "charms", "basePathG": basePathG}, "-base="+basePathG)

	if errT != nil {
		res.Header().Set("Content-Type", "text/html; charset=utf-8")

		errStrT := tk.ErrStrf("action failed: %v", errT)
		tk.Pln(errStrT)
		res.Write([]byte(errStrT))
		return
	}

	if toWriteT == "TX_END_RESPONSE_XT" {
		return
	}

	res.Header().Set("Content-Type", "text/html; charset=utf-8")

	res.Write([]byte(toWriteT))
}

func doCharmsContent(res http.ResponseWriter, req *http.Request) {
	if res != nil {
		res.Header().Set("Access-Control-Allow-Origin", "*")
		res.Header().Set("Access-Control-Allow-Headers", "*")
		res.Header().Set("Content-Type", "text/html; charset=utf-8")
	}

	if req != nil {
		req.ParseForm()
		req.ParseMultipartForm(100000000)
	}

	reqT := tk.GetFormValueWithDefaultValue(req, "dc", "")

	if charlang.GlobalsG.VerboseLevel > 0 {
		tk.Pl("RequestURI: %v", req.RequestURI)
	}

	if reqT == "" {
		if tk.StartsWith(req.RequestURI, "/dc") {
			reqT = req.RequestURI[3:]
		}
	}

	tmps := tk.Split(reqT, "?")
	if len(tmps) > 1 {
		reqT = tmps[0]
	}

	if tk.StartsWith(reqT, "/") {
		reqT = reqT[1:]
	}

	var paraMapT map[string]string
	var errT error

	vo := tk.GetFormValueWithDefaultValue(req, "vo", "")

	if vo == "" {
		paraMapT = tk.FormToMap(req.Form)
	} else {
		paraMapT, errT = tk.MSSFromJSON(vo)

		if errT != nil {
			res.Write([]byte(genFailCompact("action failed", "invalid parameter format", "-compact")))
			return
		}
	}

	if charlang.GlobalsG.VerboseLevel > 0 {
		tk.Pl("[%v] REQ: %#v (%#v)", tk.GetNowTimeStringFormal(), reqT, paraMapT)
	}

	toWriteT := ""

	fileNameT := "router"

	if !tk.EndsWith(fileNameT, ".char") {
		fileNameT += ".char"
	}

	// fcT := tk.LoadStringFromFile(filepath.Join(basePathG, "xms", fileNameT))
	// absT, _ := filepath.Abs(filepath.Join(basePathG, fileNameT))
	// tk.Pln("loading", absT)

	fullPathT := filepath.Join(basePathG, fileNameT)

	fcT := tk.LoadStringFromFile(fullPathT)
	if tk.IsErrStr(fcT) {
		res.Write([]byte(genFailCompact("action failed", tk.GetErrStr(fcT), "-compact")))
		return
	}

	// vmT := xie.NewVMQuick(nil)

	// vmT.SetVar(nil, "paraMapG", paraMapT)
	// vmT.SetVar(nil, "requestG", req)
	// vmT.SetVar(nil, "responseG", res)
	// vmT.SetVar(nil, "reqNameG", reqT)
	// vmT.SetVar(nil, "basePathG", basePathG)

	// vmT.SetVar("inputG", objA)

	// lrs := vmT.Load(nil, fcT)

	// contentTypeT := res.Header().Get("Content-Type")

	// if tk.IsErrX(lrs) {
	// 	if tk.StartsWith(contentTypeT, "text/json") {
	// 		res.Write([]byte(tk.GenerateJSONPResponse("fail", tk.Spr("action failed: %v", tk.GetErrStrX(lrs)), req)))
	// 		return
	// 	}

	// 	res.Write([]byte(genFailCompact("action failed", tk.GetErrStrX(lrs), "-compact")))
	// 	return
	// }

	// rs := vmT.Run()

	toWriteT, errT = charlang.RunScriptOnHttp(fcT, nil, res, req, paraMapT["input"], nil, paraMapT, map[string]interface{}{"scriptPathG": fullPathT, "runModeG": "chardc", "basePathG": basePathG}, "-base="+basePathG)

	if errT != nil {
		res.Header().Set("Content-Type", "text/html; charset=utf-8")

		errStrT := tk.ErrStrf("action failed: %v", errT)
		tk.Pln(errStrT)
		res.Write([]byte(errStrT))
		return
	}

	if toWriteT == "TX_END_RESPONSE_XT" {
		return
	}

	res.Header().Set("Content-Type", "text/html; charset=utf-8")

	res.Write([]byte(toWriteT))
	contentTypeT := res.Header().Get("Content-Type")

	if errT != nil {
		if tk.StartsWith(contentTypeT, "text/json") {
			res.Write([]byte(tk.GenerateJSONPResponse("fail", tk.Spr("action failed: %v", errT), req)))
			return
		}

		res.Write([]byte(genFailCompact("action failed", errT.Error(), "-compact")))
		return
	}

	if toWriteT == "TX_END_RESPONSE_XT" {
		return
	}

	res.Header().Set("Content-Type", "text/html; charset=utf-8")

	res.Write([]byte(toWriteT))

}

// Xielang
func doXms(res http.ResponseWriter, req *http.Request) {
	if res != nil {
		res.Header().Set("Access-Control-Allow-Origin", "*")
		res.Header().Set("Access-Control-Allow-Headers", "*")
		res.Header().Set("Content-Type", "text/html; charset=utf-8")
	}

	if req != nil {
		req.ParseForm()
		req.ParseMultipartForm(100000000)
	}

	// tk.Pl("xms: %v", req)

	reqT := tk.GetFormValueWithDefaultValue(req, "xms", "")

	if reqT == "" {
		if tk.StartsWith(req.RequestURI, "/xms") {
			reqT = req.RequestURI[4:]
		}
	}

	tmps := tk.Split(reqT, "?")
	if len(tmps) > 1 {
		reqT = tmps[0]
	}

	if tk.StartsWith(reqT, "/") {
		reqT = reqT[1:]
	}

	// tk.Pl("charms: %v", reqT)

	var paraMapT map[string]string
	var errT error

	vo := tk.GetFormValueWithDefaultValue(req, "vo", "")

	if vo == "" {
		paraMapT = tk.FormToMap(req.Form)
	} else {
		paraMapT, errT = tk.MSSFromJSON(vo)

		if errT != nil {
			res.Write([]byte(genFailCompact("action failed", "invalid vo format", "-compact")))
			return
		}
	}

	if verboseG {
		tk.Pl("[%v] REQ: %#v (%#v)", tk.GetNowTimeStringFormal(), reqT, paraMapT)
	}

	toWriteT := ""

	fileNameT := reqT

	if !tk.EndsWith(fileNameT, ".xie") {
		fileNameT += ".xie"
	}

	fcT := tk.LoadStringFromFile(filepath.Join(basePathG, fileNameT))
	if tk.IsErrStr(fcT) {
		res.Write([]byte(genFailCompact("action failed", tk.GetErrStr(fcT), "-compact")))
		return
	}

	// envT := make(map[string]interface{})

	// envT["argsG"] = paraMapT
	// envT["requestG"] = req
	// envT["responseG"] = res
	// envT["reqNameG"] = reqT

	vmT := xie.NewVMQuick()

	vmT.SetVar(vmT.Running, "argsG", paraMapT)
	vmT.SetVar(vmT.Running, "requestG", req)
	vmT.SetVar(vmT.Running, "responseG", res)
	vmT.SetVar(vmT.Running, "reqNameG", reqT)
	vmT.SetVar(vmT.Running, "basePathG", basePathG)

	// vmT.SetVar("inputG", objA)

	lrs := vmT.Load(vmT.Running, fcT)

	if tk.IsError(lrs) {
		res.Write([]byte(genFailCompact("action failed", lrs.Error(), "-compact")))
		return
	}

	// var argsT []string = tk.JSONToStringArray(tk.GetSwitch(optsA, "-args=", "[]"))

	// if argsT != nil {
	// 	vmT.VarsM["argsG"] = argsT
	// } else {
	// 	vmT.VarsM["argsG"] = []string{}
	// }

	rs := vmT.Run()

	if errT != nil {
		res.Write([]byte(genFailCompact("action failed", errT.Error(), "-compact")))
		return
	}

	if tk.IsErrX(rs) {
		res.Write([]byte(genFailCompact("action failed", tk.GetErrStrX(rs), "-compact")))
		return
	}

	toWriteT = tk.ToStr(rs)

	if toWriteT == "TX_END_RESPONSE_XT" {
		return
	}

	res.Header().Set("Content-Type", "text/html; charset=utf-8")

	res.Write([]byte(toWriteT))

	// paraMapT["_reqHost"] = req.Host
	// paraMapT["_reqInfo"] = fmt.Sprintf("%#v", req)

}

func chpHandler(strA string, w http.ResponseWriter, r *http.Request) {
	var paraMapT map[string]string
	var errT error

	r.ParseForm()

	vo := tk.GetFormValueWithDefaultValue(r, "vo", "")

	if vo == "" {
		paraMapT = tk.FormToMap(r.Form)
	} else {
		paraMapT, errT = tk.MSSFromJSON(vo)

		if errT != nil {
			paraMapT = map[string]string{}
		}
	}

	evalT := charlang.NewEvalQuick(map[string]interface{}{"versionG": charlang.VersionG, "argsG": []string{}, "scriptPathG": "", "runModeG": "chp", "paraMapG": paraMapT, "requestG": r, "responseG": w, "reqUriG": r.RequestURI}, charlang.MainCompilerOptions)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	countT := 0

	replaceFuncT := func(str1A string) string {
		countT++
		// tk.Pl("found: %v", str1A)
		lastResultT, lastBytecodeT, errT := evalT.Run(ctx, []byte(str1A[5:len(str1A)-2]))

		if verboseG {
			tk.Pln("result:", lastResultT, lastBytecodeT, errT)
		}

		if errT != nil {
			return fmt.Sprintf("[%v] %v", countT, tk.ErrorToString(errT))
		}

		if lastResultT != nil && lastResultT.TypeCode() != 0 {
			return fmt.Sprintf("%v", lastResultT)
		}

		return ""
	}

	re := regexp.MustCompile(`(?sm)<\?chp.*?\?>`)

	strT := re.ReplaceAllStringFunc(strA, replaceFuncT)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	w.Write([]byte(strT))
}

func ghpHandler(strA string, w http.ResponseWriter, r *http.Request) {
	var paraMapT map[string]string
	var errT error

	r.ParseForm()

	vo := tk.GetFormValueWithDefaultValue(r, "vo", "")

	if vo == "" {
		paraMapT = tk.FormToMap(r.Form)
	} else {
		paraMapT, errT = tk.MSSFromJSON(vo)

		if errT != nil {
			paraMapT = map[string]string{}
		}
	}

	gox.InitQLVM()

	vmT := qlang.New("-noexit")

	vmT.SetVar("paraMapG", paraMapT)

	vmT.SetVar("requestG", r)

	vmT.SetVar("responseG", w)

	countT := 0

	replaceFuncT := func(str1A string) string {
		strT := str1A[5 : len(str1A)-2]

		countT++

		retT := ""

		if tk.StartsWith(strT, "//TXDEF#") {
			tmps := tk.DecryptStringByTXDEF(strT, "topxeq")

			if !tk.IsErrStr(tmps) {
				strT = tmps
			}
		}

		errT = vmT.SafeEval(strT)

		if errT != nil {
			return fmt.Sprintf("[%v] %v", countT, tk.ErrorToString(errT))
		}

		rs, ok := vmT.GetVar("outG")

		if ok {
			if rs != nil {
				strT, ok := rs.(string)
				if ok {
					return strT
				}

				return fmt.Sprintf("%v", rs)
			}

			return retT
		}

		return retT

	}

	re := regexp.MustCompile(`(?sm)<\?ghp.*?\?>`)

	strT := re.ReplaceAllStringFunc(strA, replaceFuncT)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	w.Write([]byte(strT))
}

func xhpHandler(strA string, w http.ResponseWriter, r *http.Request) {
	var paraMapT map[string]string
	var errT error

	r.ParseForm()

	vo := tk.GetFormValueWithDefaultValue(r, "vo", "")

	if vo == "" {
		paraMapT = tk.FormToMap(r.Form)
	} else {
		paraMapT, errT = tk.MSSFromJSON(vo)

		if errT != nil {
			paraMapT = map[string]string{}
		}
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	vm0T := xie.NewVM()

	if tk.IsError(vm0T) {
		w.Write([]byte(fmt.Sprintf("failed to initialize VM: %v", tk.GetErrStrX(vm0T))))
		return
	}

	vmT := vm0T.(*xie.XieVM)

	vmT.SetVar(vmT.Running, "paraMapG", paraMapT)
	vmT.SetVar(vmT.Running, "requestG", r)
	vmT.SetVar(vmT.Running, "responseG", w)

	countT := 0

	replaceFuncT := func(str1A string) string {
		strT := str1A[5 : len(str1A)-2]

		countT++

		if tk.StartsWith(strT, "//TXDEF#") {
			tmps := tk.DecryptStringByTXDEF(strT, "topxeq")

			if !tk.IsErrStr(tmps) {
				strT = tmps
			}
		}

		retT := ""

		originalCodeLenT := vmT.GetCodeLen(vmT.Running)

		lrs := vmT.Load(vmT.Running, strT)

		if tk.IsError(lrs) {
			return fmt.Sprintf("TXERROR:[%v] failed to load source code of the script: %v", countT, tk.GetErrStrX(lrs))
		}

		rs := vmT.Run(originalCodeLenT)

		noResultT := tk.IsUndefined(rs) // == "TXERROR:no result")

		if tk.IsErrX(rs) {
			return fmt.Sprintf("TXERROR:[%v] failed to run: %v", countT, tk.GetErrStrX(rs))
		}

		if !noResultT {
			return tk.ToStr(rs)
		}

		return retT
	}

	re := regexp.MustCompile(`(?sm)<\?xhp.*?\?>`)

	strT := re.ReplaceAllStringFunc(strA, replaceFuncT)

	w.Write([]byte(strT))
}

var staticFS http.Handler = nil

func serveStaticDirHandler(w http.ResponseWriter, r *http.Request) {
	if staticFS == nil {
		// tk.Pl("staticFS: %#v", staticFS)
		// staticFS = http.StripPrefix("/w/", http.FileServer(http.Dir(filepath.Join(basePathG, "w"))))
		hdl := http.FileServer(http.Dir(webPathG))
		// tk.Pl("hdl: %#v", hdl)
		staticFS = hdl
	}

	old := r.URL.Path

	// tk.Pl("urlPath: %v", r.URL.Path)

	name := filepath.Join(webPathG, path.Clean(old))

	// tk.Pl("name: %v", name)

	info, err := os.Lstat(name)
	if err == nil {
		if !info.IsDir() {

			if strings.HasSuffix(name, ".chp") {
				chpHandler(tk.LoadStringFromFile(name), w, r)

				return
			}

			if strings.HasSuffix(name, ".ghp") {
				ghpHandler(tk.LoadStringFromFile(name), w, r)

				return
			}

			if strings.HasSuffix(name, ".xhp") {
				xhpHandler(tk.LoadStringFromFile(name), w, r)

				return
			}

			staticFS.ServeHTTP(w, r)
			// http.ServeFile(w, r, name)
		} else {
			if tk.IfFileExists(filepath.Join(name, "index.html")) {
				staticFS.ServeHTTP(w, r)
			} else {
				http.NotFound(w, r)
			}
		}
	} else {
		http.NotFound(w, r)
	}

}

func startHttpsServer(portA string) {
	if !tk.StartsWith(portA, ":") {
		portA = ":" + portA
	}

	err := http.ListenAndServeTLS(portA, filepath.Join(certPathG, "server.crt"), filepath.Join(certPathG, "server.key"), muxG)
	if err != nil {
		tk.PlNow("failed to start https: %v", err)
	}

}

func doGoxn() {
	gox.ServerModeG = true

	portG = tk.GetSwitch(os.Args, "-port=", portG)
	sslPortG = tk.GetSwitch(os.Args, "-sslPort=", sslPortG)

	verboseG = tk.IfSwitchExistsWhole(os.Args, "-verbose")

	if !tk.StartsWith(portG, ":") {
		portG = ":" + portG
	}

	if !tk.StartsWith(sslPortG, ":") {
		sslPortG = ":" + sslPortG
	}

	basePathG = tk.GetSwitch(os.Args, "-dir=", basePathG)
	webPathG = tk.GetSwitch(os.Args, "-webDir=", basePathG)
	certPathG = tk.GetSwitch(os.Args, "-certDir=", certPathG)

	muxG = http.NewServeMux()

	muxG.HandleFunc("/ms/", doMs)
	muxG.HandleFunc("/ms", doMs)

	muxG.HandleFunc("/wms/", doWms)
	muxG.HandleFunc("/wms", doWms)

	muxG.HandleFunc("/xms/", doXms)
	muxG.HandleFunc("/xms", doXms)

	muxG.HandleFunc("/charms/", doCharms)
	muxG.HandleFunc("/charms", doCharms)

	// dynamic content
	muxG.HandleFunc("/dc/", doCharmsContent)
	muxG.HandleFunc("/dc", doCharmsContent)

	muxG.HandleFunc("/", serveStaticDirHandler)

	tk.PlNow("Gox Server %v -port=%v -sslPort=%v -dir=%v -webDir=%v -certDir=%v", gox.VersionG, portG, sslPortG, basePathG, webPathG, certPathG)

	if sslPortG != "" {
		tk.PlNow("try starting ssl server on %v...", sslPortG)
		go startHttpsServer(sslPortG)
	}

	tk.PlNow("try starting server on %v ...", portG)
	err := http.ListenAndServe(portG, muxG)

	if err != nil {
		tk.PlNow("failed to start: %v", err)
	}

	// resultT, errT := goxn.RunScript(tk.LoadStringFromFile(os.Args[1]), tk.GetSwitch(os.Args, "-input="), os.Args, nil)

	// tk.CheckErrf("error: %v", errT)

	// tk.Pl("%v", resultT)

}

// for gox main
func test() {
	if tk.IfSwitchExists(os.Args, "-dotest") {
		tk.Pl("%v", CodeG)
	}
}

func showHelp() {
	tk.Pl("Gox by TopXeQ V%v\n", gox.VersionG)

	tk.Pl("Usage: gox [-v|-h] test.gox, ...\n")
	tk.Pl("or just gox without arguments to start REPL instead.\n")

}

func runInteractiveQlang() int {
	var following bool
	var source string

	tk.Pl("Gox %v", gox.VersionG)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		if following {
			source += "\n"
			fmt.Print("  ")
		} else {
			fmt.Print("> ")
		}

		if !scanner.Scan() {
			break
		}
		source += scanner.Text()
		if source == "" {
			continue
		}
		if source == "quit()" {
			break
		}

		// stmts, err := parser.ParseSrc(source)

		// if e, ok := err.(*parser.Error); ok {
		// 	es := e.Error()
		// 	if strings.HasPrefix(es, "syntax error: unexpected") {
		// 		if strings.HasPrefix(es, "syntax error: unexpected $end,") {
		// 			following = true
		// 			continue
		// 		}
		// 	} else {
		// 		if e.Pos.Column == len(source) && !e.Fatal {
		// 			fmt.Fprintln(os.Stderr, e)
		// 			following = true
		// 			continue
		// 		}
		// 		if e.Error() == "unexpected EOF" {
		// 			following = true
		// 			continue
		// 		}
		// 	}
		// }

		gox.RetG = gox.NotFoundG

		err := gox.QlVMG.SafeEval(source)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			following = false
			source = ""
			continue
		}

		if gox.RetG != gox.NotFoundG {
			fmt.Println(gox.RetG)
		}

		following = false
		source = ""
	}

	if err := scanner.Err(); err != nil {
		if err != io.EOF {
			fmt.Fprintln(os.Stderr, "ReadString error:", err)
			return 12
		}
	}

	return 0
}

func runLine(strA string) interface{} {
	argsT, errT := tk.ParseCommandLine(strA)

	if errT != nil {
		return errT
	}

	return runArgs(argsT...)
}

func runArgs(argsA ...string) interface{} {
	argsT := argsA

	if tk.IfSwitchExistsWhole(argsT, "-version") {
		tk.Pl("Gox by TopXeQ V%v", gox.VersionG)
		return "github.com/topxeq/gox"

	}

	if tk.IfSwitchExistsWhole(argsT, "-h") {
		showHelp()
		return nil
	}

	scriptT := tk.GetParameterByIndexWithDefaultValue(argsT, 0, "")

	// GUI related start

	// full version related start
	if tk.IfSwitchExistsWhole(argsT, "-edit") {
		// editFile(scriptT)
		rs := gox.RunScriptX(gox.EditFileScriptG, argsT...)

		if rs != gox.NotFoundG && rs != nil {
			tk.Pl("%v", rs)
		}

		return nil
	}
	// full version related end

	// GUI related end

	if tk.IfSwitchExistsWhole(argsT, "-initgui") {
		// applicationPathT := tk.GetApplicationPath()

		// osT := tk.GetOSName()

		// if tk.Contains(osT, "inux") {
		// 	tk.Pl("Please visit the following URL to find out how to make Sciter environment ready in Linux: ")

		// 	return nil
		// } else if tk.Contains(osT, "arwin") {
		// 	tk.Pl("Please visit the following URL to find out how to make Sciter environment ready in Linux: ")

		// 	return nil
		// } else {
		// 	// rs := tk.DownloadFile("http://scripts.frenchfriend.net/pub/sciterts.dll", applicationPathT, "sciterts.dll")

		// 	// if tk.IsErrorString(rs) {

		// 	// 	return tk.Errf("failed to download Sciter DLL file.")
		// 	// }

		// 	// tk.Pl("Sciter DLL downloaded to application path.")

		// 	// rs = tk.DownloadFile("http://scripts.frenchfriend.net/pub/webview.dll", applicationPathT, "webview.dll", false)

		// 	// if tk.IsErrorString(rs) {

		// 	// 	return tk.Errf("failed to download webview DLL file.")
		// 	// }

		// 	// rs = tk.DownloadFile("http://scripts.frenchfriend.net/pub/WebView2Loader.dll", applicationPathT, "WebView2Loader.dll", false)

		// 	// if tk.IsErrorString(rs) {

		// 	// 	return tk.Errf("failed to download webview DLL file.")
		// 	// }

		// 	// tk.Pl("webview DLL downloaded to application path.")

		// 	return nil
		// }
	}

	ifXieT := tk.IfSwitchExistsWhole(argsT, "-xie")
	ifCharT := tk.IfSwitchExistsWhole(argsT, "-char")
	ifClipT := tk.IfSwitchExistsWhole(argsT, "-clip")
	ifEmbedT := (gox.CodeTextG != "") && (!tk.IfSwitchExistsWhole(argsT, "-noembed"))

	ifInExeT := false
	inExeCodeT := ""

	binNameT, errT := os.Executable()
	if errT != nil {
		binNameT = ""
	}

	baseBinNameT := filepath.Base(binNameT)

	text1T := tk.Trim(`740404`)
	text2T := tk.Trim(`690415`)
	text3T := tk.Trim(`040626`)

	if binNameT != "" {
		if !tk.StartsWith(baseBinNameT, "gox") {
			buf1, errT := tk.LoadBytesFromFileE(binNameT)
			if errT == nil {
				re := regexp.MustCompile(text1T + text2T + text3T + `(.*?) *` + text3T + text2T + text1T)
				matchT := re.FindAllSubmatch(buf1, -1)

				if matchT != nil && len(matchT) > 0 {
					codeStrT := string(matchT[len(matchT)-1][1])

					decCodeT := tk.DecryptStringByTXDEF(codeStrT, "topxeq")
					if !tk.IsErrStr(decCodeT) {
						ifInExeT = true
						inExeCodeT = decCodeT
					}

				}
			}
		}
	}

	if tk.IfSwitchExistsWhole(argsT, "-shell") {
		gox.InitQLVM()

		var guiHandlerG tk.TXDelegate = guiHandler

		gox.QlVMG.SetVar("argsG", argsT)
		gox.QlVMG.SetVar("guiG", guiHandlerG)

		runInteractiveQlang()

		// tk.Pl("not enough parameters")

		return nil
	}

	if tk.IfSwitchExistsWhole(argsT, "-server") {

		// gox.InitQLVM()
		doGoxn()

		// tk.Pl("not enough parameters")

		return nil
	}

	cmdT := tk.GetSwitchWithDefaultValue(argsT, "-cmd=", "")

	if cmdT != "" {
		scriptT = "CMD"
	}

	if scriptT == "" && (!ifClipT) && (!ifEmbedT) && (!ifInExeT) {

		// autoPathT := filepath.Join(tk.GetApplicationPath(), "auto.gox")
		// autoGxbPathT := filepath.Join(tk.GetApplicationPath(), "auto.gxb")
		autoPathT := "auto.gox"
		autoGxbPathT := "auto.gxb"

		if tk.IfFileExists(autoPathT) {
			scriptT = autoPathT
		} else if tk.IfFileExists(autoGxbPathT) {
			scriptT = autoGxbPathT
		} else {
			gox.InitQLVM()

			var guiHandlerG tk.TXDelegate = guiHandler

			gox.QlVMG.SetVar("argsG", argsT)
			gox.QlVMG.SetVar("guiG", guiHandlerG)

			runInteractiveQlang()

			// tk.Pl("not enough parameters")

			return nil
		}

	}

	encryptCodeT := tk.GetSwitchWithDefaultValue(argsT, "-encrypt=", "")

	if encryptCodeT != "" {
		fcT := tk.LoadStringFromFile(scriptT)

		if tk.IsErrorString(fcT) {

			return tk.Errf("failed to load file [%v]: %v", scriptT, tk.GetErrorString(fcT))
		}

		encStrT := tk.EncryptStringByTXDEF(fcT, encryptCodeT)

		if tk.IsErrorString(encStrT) {

			return tk.Errf("failed to encrypt content [%v]: %v", scriptT, tk.GetErrorString(encStrT))
		}

		rsT := tk.SaveStringToFile("//TXDEF#"+encStrT, scriptT+"e")

		if tk.IsErrorString(rsT) {

			return tk.Errf("failed to encrypt file [%v]: %v", scriptT, tk.GetErrorString(rsT))
		}

		return nil
	}

	decryptCodeT := tk.GetSwitchWithDefaultValue(argsT, "-decrypt=", "")

	if decryptCodeT != "" {
		fcT := tk.LoadStringFromFile(scriptT)

		if tk.IsErrorString(fcT) {

			return tk.Errf("failed to load file [%v]: %v", scriptT, tk.GetErrorString(fcT))
		}

		decStrT := tk.DecryptStringByTXDEF(fcT, decryptCodeT)

		if tk.IsErrorString(decStrT) {

			return tk.Errf("failed to decrypt content [%v]: %v", scriptT, tk.GetErrorString(decStrT))
		}

		rsT := tk.SaveStringToFile(decStrT, scriptT+"d")

		if tk.IsErrorString(rsT) {

			return tk.Errf("failed to decrypt file [%v]: %v", scriptT, tk.GetErrorString(rsT))
		}

		return nil
	}

	decryptRunCodeT := tk.GetSwitchWithDefaultValue(argsT, "-decrun=", "")

	ifBatchT := tk.IfSwitchExistsWhole(argsT, "-batch")

	if !ifBatchT {
		if tk.EndsWithIgnoreCase(scriptT, ".gxb") {
			ifBatchT = true
		}
	}

	ifBinT := tk.IfSwitchExistsWhole(argsT, "-bin")
	if ifBinT {
	}

	ifRunT := tk.IfSwitchExistsWhole(argsT, "-run")
	ifExampleT := tk.IfSwitchExistsWhole(argsT, "-example")
	ifGoPathT := tk.IfSwitchExistsWhole(argsT, "-gopath")
	ifLocalT := tk.IfSwitchExistsWhole(argsT, "-local")
	ifAppPathT := tk.IfSwitchExistsWhole(argsT, "-apppath")
	ifRemoteT := tk.IfSwitchExistsWhole(argsT, "-remote")
	ifCloudT := tk.IfSwitchExistsWhole(argsT, "-cloud")
	sshT := tk.GetSwitchWithDefaultValue(argsT, "-ssh=", "")
	ifViewT := tk.IfSwitchExistsWhole(argsT, "-view")
	ifOpenT := tk.IfSwitchExistsWhole(argsT, "-open")
	ifCompileT := tk.IfSwitchExistsWhole(argsT, "-compile")

	gox.VerboseG = tk.IfSwitchExistsWhole(argsT, "-verbose")

	ifMagicT := false
	magicNumberT, errT := tk.StrToIntE(scriptT)

	if errT == nil {
		ifMagicT = true
	}

	var fcT string

	if ifInExeT && inExeCodeT != "" && !tk.IfSwitchExistsWhole(os.Args, "-noin") {
		fcT = inExeCodeT

		gox.ScriptPathG = ""
	} else if cmdT != "" {
		fcT = cmdT

		if tk.IfSwitchExistsWhole(os.Args, "-urlDecode") {
			fcT = tk.UrlDecode(fcT)
		}

		gox.ScriptPathG = ""
	} else if ifMagicT {
		fcT = gox.GetMagic(magicNumberT)

		gox.ScriptPathG = ""
	} else if ifRunT {
		if tk.IfSwitchExistsWhole(os.Args, "-urlDecode") {
			fcT = tk.UrlDecode(scriptT)
		} else {
			fcT = scriptT
		}
		tk.Pl("run cmd(%v)", fcT)

		gox.ScriptPathG = ""
	} else if ifExampleT {
		if (!tk.EndsWith(scriptT, ".gox")) && (!tk.EndsWith(scriptT, ".xie")) && (!tk.EndsWith(scriptT, ".char")) {
			scriptT += ".gox"
		}

		gox.ScriptPathG = "https://gitee.com/topxeq/gox/raw/master/scripts/" + scriptT

		fcT = tk.DownloadPageUTF8("https://gitee.com/topxeq/gox/raw/master/scripts/"+scriptT, nil, "", 30)

	} else if ifRemoteT {
		gox.ScriptPathG = scriptT
		fcT = tk.DownloadPageUTF8(scriptT, nil, "", 30)

	} else if ifClipT {
		fcT = tk.GetClipText()

		gox.ScriptPathG = ""
	} else if ifEmbedT {
		fcT = gox.CodeTextG

		gox.ScriptPathG = ""
	} else if ifCloudT {
		if (!tk.EndsWith(scriptT, ".gox")) && (!tk.EndsWith(scriptT, ".xie")) && (!tk.EndsWith(scriptT, ".char")) {
			scriptT += ".gox"
		}

		basePathT, errT := tk.EnsureBasePath("gox")

		gotT := false

		if errT == nil {
			cfgPathT := tk.JoinPath(basePathT, "cloud.cfg")

			cfgStrT := tk.Trim(tk.LoadStringFromFile(cfgPathT))

			if !tk.IsErrorString(cfgStrT) {
				gox.ScriptPathG = cfgStrT + scriptT

				fcT = tk.DownloadPageUTF8(cfgStrT+scriptT, nil, "", 30)

				gotT = true
			}

		}

		if !gotT {
			gox.ScriptPathG = scriptT
			fcT = tk.DownloadPageUTF8(scriptT, nil, "", 30)
		}

	} else if sshT != "" {
		if (!tk.EndsWith(scriptT, ".gox")) && (!tk.EndsWith(scriptT, ".xie")) && (!tk.EndsWith(scriptT, ".char")) {
			scriptT += ".gox"
		}

		fcT = gox.DownloadStringFromSSH(sshT, scriptT)

		if tk.IsErrorString(fcT) {

			return tk.Errf("failed to get script from SSH: %v", tk.GetErrorString(fcT))
		}

		gox.ScriptPathG = ""
	} else if ifGoPathT {
		if (!tk.EndsWith(scriptT, ".gox")) && (!tk.EndsWith(scriptT, ".xie")) && (!tk.EndsWith(scriptT, ".char")) {
			scriptT += ".gox"
		}

		gox.ScriptPathG = filepath.Join(tk.GetEnv("GOPATH"), "src", "github.com", "topxeq", "gox", "scripts", scriptT)

		fcT = tk.LoadStringFromFile(gox.ScriptPathG)
	} else if ifAppPathT {
		if (!tk.EndsWith(scriptT, ".gox")) && (!tk.EndsWith(scriptT, ".xie")) && (!tk.EndsWith(scriptT, ".char")) {
			scriptT += ".gox"
		}

		gox.ScriptPathG = filepath.Join(tk.GetApplicationPath(), scriptT)

		fcT = tk.LoadStringFromFile(gox.ScriptPathG)
	} else if ifLocalT {
		if (!tk.EndsWith(scriptT, ".gox")) && (!tk.EndsWith(scriptT, ".xie")) && (!tk.EndsWith(scriptT, ".char")) {
			scriptT += ".gox"
		}

		localPathT := gox.GetCfgString("localScriptPath.cfg")

		if tk.IsErrorString(localPathT) {
			// tk.Pl("failed to get local path: %v", tk.GetErrorString(localPathT))

			return tk.Errf("failed to get local path: %v", tk.GetErrorString(localPathT))
		}

		// if tk.GetEnv("GOXVERBOSE") == "true" {
		// 	tk.Pl("Try to load script from %v", filepath.Join(localPathT, scriptT))
		// }

		gox.ScriptPathG = filepath.Join(localPathT, scriptT)

		fcT = tk.LoadStringFromFile(gox.ScriptPathG)
	} else {
		gox.ScriptPathG = scriptT
		fcT = tk.LoadStringFromFile(scriptT)

	}

	if strings.HasSuffix(scriptT, ".xie") {
		ifXieT = true
	} else if strings.HasSuffix(scriptT, ".char") {
		ifCharT = true
	}

	if tk.IsErrorString(fcT) {
		return tk.Errf("failed to load script from %v: %v", scriptT, tk.GetErrorString(fcT))
	}

	if tk.StartsWith(fcT, "//TXDEF#") {
		if decryptRunCodeT == "" {
			tmps := tk.DecryptStringByTXDEF(fcT, "goxxie")

			if tk.IsErrStr(tmps) {
				tk.Prf("Password: ")
				decryptRunCodeT = tk.Trim(tk.GetInputBufferedScan())
			} else {
				fcT = tmps
			}

			// fcT = fcT[8:]
		}
	}

	if decryptRunCodeT != "" {
		fcT = tk.DecryptStringByTXDEF(fcT, decryptRunCodeT)
	}

	if ifViewT {
		if !ifInExeT {
			tk.Pl("%v", fcT)
		}

		return nil
	}

	if ifCompileT {
		appPathT, errT := os.Executable()

		tk.CheckError(errT)

		outputT := tk.Trim(tk.GetSwitch(os.Args, "-output=", "output.exe"))

		if fcT == "" {
			tk.Fatalf("code empty")
		}

		buf1, errT := tk.LoadBytesFromFileE(appPathT)
		if errT != nil {
			tk.Fatalf("loading bin failed: %v", errT)
		}

		encTextT := tk.EncryptStringByTXDEF(fcT, "topxeq")

		encBytesT := []byte(encTextT)

		lenEncT := len(encBytesT)

		text1T := tk.Trim("740404")
		text2T := tk.Trim("690415")
		text3T := tk.Trim("040626")

		re := regexp.MustCompile(text1T + text2T + text3T + `(.*)` + text3T + text2T + text1T)
		matchT := re.FindSubmatchIndex(buf1)
		if matchT == nil {
			tk.Fatalf("invald bin")
		}

		bufCodeLenT := matchT[3] - matchT[2]

		var buf3 bytes.Buffer

		if bufCodeLenT < lenEncT {

			buf3.Write(buf1)
			buf3.Write([]byte("74040469" + "0415840215"))
			buf3.Write(encBytesT)
			buf3.Write([]byte("840215690" + "415740404"))
		} else {
			buf3.Write(buf1[:matchT[2]])
			buf3.Write(encBytesT)
			buf3.Write(buf1[matchT[2]+lenEncT:])
		}

		errT = tk.SaveBytesToFileE(buf3.Bytes(), outputT)
		tk.CheckError(errT)

		return nil

	}

	if ifOpenT {
		tk.RunWinFileWithSystemDefault(gox.ScriptPathG)

		return nil
	}

	// if ifCompileT {
	// 	initQLVM()

	// 	gox.QlVMG.SetVar("argsG", argsT)

	// 	retG = gox.NotFoundG

	// 	endT, errT := gox.QlVMG.SafeCl([]byte(fcT), "")
	// 	if errT != nil {

	// 		// tk.Pl()

	// 		// f, l := gox.QlVMG.Code.Line(gox.QlVMG.Code.Reserve().Next())
	// 		// tk.Pl("Next line: %v, %v", f, l)

	// 		return tk.Errf("failed to compile script(%v) error: %v\n", scriptT, errT)
	// 	}

	// 	tk.Pl("endT: %v", endT)

	// 	errT = gox.QlVMG.DumpEngine()

	// 	if errT != nil {
	// 		return tk.Errf("failed to dump engine: %v\n", errT)
	// 	}

	// 	tk.Plvsr(gox.QlVMG.Cpl.GetCode().Len(), gox.QlVMG.Run())

	// 	return nil
	// }

	if !ifBatchT {
		if tk.RegStartsWith(fcT, `//\s*(GXB|gxb)`) {
			ifBatchT = true
		}
	}

	if ifBatchT {
		listT := tk.SplitLinesRemoveEmpty(fcT)

		// tk.Plv(fcT)
		// tk.Plv(listT)

		for _, v := range listT {
			// tk.Pl("Run line: %#v", v)
			v = tk.Trim(v)

			if tk.StartsWith(v, "//") {
				continue
			}

			rsT := runLine(v)

			if rsT != nil {
				valueT, ok := rsT.(error)

				if ok {
					return valueT
				} else {
					tk.Pl("%v", rsT)
				}
			}

		}

		return nil
	}

	if ifXieT {
		var guiHandlerG tk.TXDelegate = guiHandler

		rs := xie.RunCode(fcT, nil, map[string]interface{}{"guiG": guiHandlerG, "scriptPathG": gox.ScriptPathG, "basePathG": basePathG}, argsT...) // "guiG": guiHandlerG,
		if !tk.IsUndefined(rs) {
			tk.Pl("%v", rs)
		}

		return nil
	}

	if ifCharT {
		moduleMap := charlang.NewModuleMap()
		moduleMap.AddBuiltinModule("ex", charex.Module)
		// moduleMap.AddBuiltinModule("fmt", charfmt.Module)

		charlang.MainCompilerOptions = &charlang.CompilerOptions{
			// ModulePath:        "", //"(repl)",
			ModuleMap: moduleMap,
			// SymbolTable:       charlang.NewSymbolTable(),
			// OptimizerMaxCycle: charlang.TraceCompilerOptions.OptimizerMaxCycle,
			// TraceParser:       true,
			// TraceOptimizer:    true,
			// TraceCompiler:     true,
			// OptimizeConst:     !noOptimizer,
			// OptimizeExpr:      !noOptimizer,

			// Trace:             os.Stdout,
			// TraceParser:       true,
			// TraceCompiler:     true,
			// TraceOptimizer:    true,
			// OptimizerMaxCycle: 1<<8 - 1,
			// OptimizeConst:     false,
			// OptimizeExpr:      false,
		}

		bytecodeT, errT := charlang.Compile([]byte(fcT), charlang.MainCompilerOptions) // charlang.DefaultCompilerOptions)
		if errT != nil {
			return errT
		}

		envT := charlang.Map{}

		envT["argsG"] = charlang.ConvertToObject(os.Args)
		envT["versionG"] = charlang.ToStringObject(charlang.VersionG)
		envT["scriptPathG"] = charlang.ToStringObject(gox.ScriptPathG)
		envT["runModeG"] = charlang.ToStringObject("script")

		vmT := charlang.NewVM(bytecodeT)

		retT, errT := vmT.Run(envT) // inParasT,

		if errT != nil {
			return fmt.Errorf("failed to execute script(%v) error: %v\n", gox.ScriptPathG, errT)
		}

		if !charlang.IsUndefInternal(retT) {
			tk.Pl("%v", retT)
		}

		return nil
	}

	gox.InitQLVM()

	var guiHandlerG tk.TXDelegate = guiHandler

	gox.QlVMG.SetVar("argsG", argsT)
	gox.QlVMG.SetVar("guiG", guiHandlerG)

	gox.RetG = gox.NotFoundG

	errT = gox.QlVMG.SafeEval(fcT)
	if errT != nil {

		// tk.Pl()

		// f, l := QlVMG.Code.Line(QlVMG.Code.Reserve().Next())
		// tk.Pl("Next line: %v, %v", f, l)

		return tk.Errf("failed to execute script(%v) error: %v\n", scriptT, errT)
	}

	rs, ok := gox.QlVMG.GetVar("outG")

	if ok {
		if rs != nil {
			return rs
		}
	}

	return gox.RetG
}

// func qlEval(strA string) string {
// 	vmT := qlang.New()

// 	retG = gox.NotFoundG

// 	errT := vmT.SafeEval(strA)

// 	if errT != nil {
// 		return errT.Error()
// 	}

// 	rs, ok := vmT.GetVar("outG")

// 	if ok {
// 		return tk.Spr("%v", rs)
// 	}

// 	if retG != gox.NotFoundG {
// 		return tk.Spr("%v", retG)
// 	}

// 	return tk.ErrStrF("no result")
// }

// var bluetoothAdapter = bluetooth.DefaultAdapter

func main() {
	// var errT error

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Exception: ", err)
			fmt.Printf("runtime error: %v\n%v\n", err, string(debug.Stack()))
		}
	}()

	test()

	// errT := bluetoothAdapter.Enable()

	// if errT != nil {
	// 	tk.Pl("enable Bluetooth function failed: %v", errT)
	// 	// exit()
	// }

	rand.Seed(time.Now().Unix())

	rs := runArgs(os.Args[1:]...)

	if rs != nil {
		valueT, ok := rs.(error)

		if ok {
			if !gox.IsUndefined(valueT) && !gox.IsNotFound(valueT) {
				tk.Pl("Error: %T %v", valueT, valueT)
			}
		} else {
			tk.Pl("%v", rs)
		}
	}

}
