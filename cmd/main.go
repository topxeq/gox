package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/topxeq/gox"
	"github.com/topxeq/tk"
	"github.com/topxeq/xie"
	"io"
	"math/rand"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"time"
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

// Xielang
func doXms(res http.ResponseWriter, req *http.Request) {
	if res != nil {
		res.Header().Set("Access-Control-Allow-Origin", "*")
		res.Header().Set("Access-Control-Allow-Headers", "*")
		res.Header().Set("Content-Type", "text/html; charset=utf-8")
	}

	if req != nil {
		req.ParseForm()
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

	muxG.HandleFunc("/wms/", doWms)
	muxG.HandleFunc("/wms", doWms)

	muxG.HandleFunc("/xms/", doXms)
	muxG.HandleFunc("/xms", doXms)

	muxG.HandleFunc("/", serveStaticDirHandler)

	tk.PlNow("Gox Server %v -port=%v -sslPort=%v -dir=%v -webDir=%v -certDir=%v", gox.VersionG, portG, sslPortG, basePathG, webPathG, certPathG)

	if sslPortG != "" {
		tk.PlNow("try starting ssl server on %v...", sslPortG)
		go startHttpsServer(sslPortG)
	}

	tk.Pl("try starting server on %v ...", portG)
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
		applicationPathT := tk.GetApplicationPath()

		osT := tk.GetOSName()

		if tk.Contains(osT, "inux") {
			tk.Pl("Please visit the following URL to find out how to make Sciter environment ready in Linux: ")

			return nil
		} else if tk.Contains(osT, "arwin") {
			tk.Pl("Please visit the following URL to find out how to make Sciter environment ready in Linux: ")

			return nil
		} else {
			rs := tk.DownloadFile("http://scripts.frenchfriend.net/pub/sciterts.dll", applicationPathT, "sciterts.dll")

			if tk.IsErrorString(rs) {

				return tk.Errf("failed to download Sciter DLL file.")
			}

			tk.Pl("Sciter DLL downloaded to application path.")

			// rs = tk.DownloadFile("http://scripts.frenchfriend.net/pub/webview.dll", applicationPathT, "webview.dll", false)

			// if tk.IsErrorString(rs) {

			// 	return tk.Errf("failed to download webview DLL file.")
			// }

			// rs = tk.DownloadFile("http://scripts.frenchfriend.net/pub/WebView2Loader.dll", applicationPathT, "WebView2Loader.dll", false)

			// if tk.IsErrorString(rs) {

			// 	return tk.Errf("failed to download webview DLL file.")
			// }

			// tk.Pl("webview DLL downloaded to application path.")

			return nil
		}
	}

	ifXieT := tk.IfSwitchExistsWhole(argsT, "-xie")
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
		if (!tk.EndsWith(scriptT, ".gox")) && (!tk.EndsWith(scriptT, ".xie")) {
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
		if (!tk.EndsWith(scriptT, ".gox")) && (!tk.EndsWith(scriptT, ".xie")) {
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
		if (!tk.EndsWith(scriptT, ".gox")) && (!tk.EndsWith(scriptT, ".xie")) {
			scriptT += ".gox"
		}

		fcT = gox.DownloadStringFromSSH(sshT, scriptT)

		if tk.IsErrorString(fcT) {

			return tk.Errf("failed to get script from SSH: %v", tk.GetErrorString(fcT))
		}

		gox.ScriptPathG = ""
	} else if ifGoPathT {
		if (!tk.EndsWith(scriptT, ".gox")) && (!tk.EndsWith(scriptT, ".xie")) {
			scriptT += ".gox"
		}

		gox.ScriptPathG = filepath.Join(tk.GetEnv("GOPATH"), "src", "github.com", "topxeq", "gox", "scripts", scriptT)

		fcT = tk.LoadStringFromFile(gox.ScriptPathG)
	} else if ifAppPathT {
		if (!tk.EndsWith(scriptT, ".gox")) && (!tk.EndsWith(scriptT, ".xie")) {
			scriptT += ".gox"
		}

		gox.ScriptPathG = filepath.Join(tk.GetApplicationPath(), scriptT)

		fcT = tk.LoadStringFromFile(gox.ScriptPathG)
	} else if ifLocalT {
		if (!tk.EndsWith(scriptT, ".gox")) && (!tk.EndsWith(scriptT, ".xie")) {
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

	if tk.IsErrorString(fcT) {
		return tk.Errf("failed to load script from %v: %v", scriptT, tk.GetErrorString(fcT))
	}

	if tk.StartsWith(fcT, "//TXDEF#") {
		if decryptRunCodeT == "" {
			tmps := tk.DecryptStringByTXDEF(fcT, "topxeq")

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
		rs := xie.RunCode(fcT, nil, map[string]interface{}{"gox.ScriptPathG": gox.ScriptPathG}, argsT...) // "guiG": guiHandlerG,
		if !tk.IsUndefined(rs) {
			tk.Pl("%v", rs)
		}

		return nil
	}

	gox.InitQLVM()

	gox.QlVMG.SetVar("argsG", argsT)

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

func main() {
	// var errT error

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Exception: ", err)
		}
	}()

	test()

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