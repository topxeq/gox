package main

import (
	"bufio"
	"strings"

	// "context"

	"io"
	"math/rand"
	"os"
	"path/filepath"
	"reflect"
	"sync"
	"time"

	"errors"
	"fmt"

	"github.com/topxeq/qlang"
	_ "github.com/topxeq/qlang/lib/builtin" // 导入 builtin 包
	_ "github.com/topxeq/qlang/lib/chan"

	execq "github.com/topxeq/qlang/exec"
	"github.com/topxeq/qlang/spec"

	// import standard packages
	qlarchivezip "github.com/topxeq/qlang/lib/archive/zip"
	qlbufio "github.com/topxeq/qlang/lib/bufio"
	qlbytes "github.com/topxeq/qlang/lib/bytes"

	qlcrypto "github.com/topxeq/qlang/lib/crypto"
	qlcryptoaes "github.com/topxeq/qlang/lib/crypto/aes"
	qlcryptocipher "github.com/topxeq/qlang/lib/crypto/cipher"
	qlcryptohmac "github.com/topxeq/qlang/lib/crypto/hmac"
	qlcryptomd5 "github.com/topxeq/qlang/lib/crypto/md5"
	qlcryptorand "github.com/topxeq/qlang/lib/crypto/rand"
	qlcryptorsa "github.com/topxeq/qlang/lib/crypto/rsa"
	qlcryptosha1 "github.com/topxeq/qlang/lib/crypto/sha1"
	qlcryptosha256 "github.com/topxeq/qlang/lib/crypto/sha256"
	qlcryptox509 "github.com/topxeq/qlang/lib/crypto/x509"

	qldatabasesql "github.com/topxeq/qlang/lib/database/sql"

	qlencodingbase64 "github.com/topxeq/qlang/lib/encoding/base64"
	qlencodingbinary "github.com/topxeq/qlang/lib/encoding/binary"
	qlencodingcsv "github.com/topxeq/qlang/lib/encoding/csv"
	qlencodinggob "github.com/topxeq/qlang/lib/encoding/gob"
	qlencodinghex "github.com/topxeq/qlang/lib/encoding/hex"
	qlencodingjson "github.com/topxeq/qlang/lib/encoding/json"
	qlencodingpem "github.com/topxeq/qlang/lib/encoding/pem"
	qlencodingxml "github.com/topxeq/qlang/lib/encoding/xml"

	qlerrors "github.com/topxeq/qlang/lib/errors"
	qlflag "github.com/topxeq/qlang/lib/flag"
	qlfmt "github.com/topxeq/qlang/lib/fmt"

	qlhashfnv "github.com/topxeq/qlang/lib/hash/fnv"

	qlhtml "github.com/topxeq/qlang/lib/html"
	qlhtmltemplate "github.com/topxeq/qlang/lib/html/template"

	qlimage "github.com/topxeq/qlang/lib/image"
	qlimage_color "github.com/topxeq/qlang/lib/image/color"
	qlimage_color_palette "github.com/topxeq/qlang/lib/image/color/palette"
	qlimage_draw "github.com/topxeq/qlang/lib/image/draw"
	qlimage_gif "github.com/topxeq/qlang/lib/image/gif"
	qlimage_jpeg "github.com/topxeq/qlang/lib/image/jpeg"
	qlimage_png "github.com/topxeq/qlang/lib/image/png"

	qlio "github.com/topxeq/qlang/lib/io"
	qlio_fs "github.com/topxeq/qlang/lib/io/fs"
	qlioioutil "github.com/topxeq/qlang/lib/io/ioutil"

	qllog "github.com/topxeq/qlang/lib/log"

	qlmath "github.com/topxeq/qlang/lib/math"
	qlmathbig "github.com/topxeq/qlang/lib/math/big"
	qlmathbits "github.com/topxeq/qlang/lib/math/bits"
	qlmathrand "github.com/topxeq/qlang/lib/math/rand"

	qlnet "github.com/topxeq/qlang/lib/net"
	qlnethttp "github.com/topxeq/qlang/lib/net/http"
	qlnet_http_cookiejar "github.com/topxeq/qlang/lib/net/http/cookiejar"
	qlnet_http_httputil "github.com/topxeq/qlang/lib/net/http/httputil"
	qlnet_mail "github.com/topxeq/qlang/lib/net/mail"
	qlnet_rpc "github.com/topxeq/qlang/lib/net/rpc"
	qlnet_rpc_jsonrpc "github.com/topxeq/qlang/lib/net/rpc/jsonrpc"
	qlnet_smtp "github.com/topxeq/qlang/lib/net/smtp"
	qlneturl "github.com/topxeq/qlang/lib/net/url"

	qlos "github.com/topxeq/qlang/lib/os"
	qlos_exec "github.com/topxeq/qlang/lib/os/exec"
	qlos_signal "github.com/topxeq/qlang/lib/os/signal"
	qlos_user "github.com/topxeq/qlang/lib/os/user"

	qlpath "github.com/topxeq/qlang/lib/path"
	qlpathfilepath "github.com/topxeq/qlang/lib/path/filepath"

	qlreflect "github.com/topxeq/qlang/lib/reflect"
	qlregexp "github.com/topxeq/qlang/lib/regexp"
	qlruntime "github.com/topxeq/qlang/lib/runtime"
	qlruntimedebug "github.com/topxeq/qlang/lib/runtime/debug"

	qlsort "github.com/topxeq/qlang/lib/sort"
	qlstrconv "github.com/topxeq/qlang/lib/strconv"
	qlstrings "github.com/topxeq/qlang/lib/strings"
	qlsync "github.com/topxeq/qlang/lib/sync"

	qltext_template "github.com/topxeq/qlang/lib/text/template"
	qltime "github.com/topxeq/qlang/lib/time"

	qlunicode "github.com/topxeq/qlang/lib/unicode"
	qlunicode_utf8 "github.com/topxeq/qlang/lib/unicode/utf8"

	// import 3rd party packages
	qlgithubbeeviketree "github.com/topxeq/qlang/lib/github.com/beevik/etree"
	qlgithubtopxeqimagetk "github.com/topxeq/qlang/lib/github.com/topxeq/imagetk"
	qlgithubtopxeqsqltk "github.com/topxeq/qlang/lib/github.com/topxeq/sqltk"
	qlgithubtopxeqtk "github.com/topxeq/qlang/lib/github.com/topxeq/tk"

	qlgithub_fogleman_gg "github.com/topxeq/qlang/lib/github.com/fogleman/gg"

	qlgithub_360EntSecGroupSkylar_excelize "github.com/topxeq/qlang/lib/github.com/360EntSecGroup-Skylar/excelize"

	qlgithub_kbinani_screenshot "github.com/topxeq/qlang/lib/github.com/kbinani/screenshot"

	qlgithub_stretchr_objx "github.com/topxeq/qlang/lib/github.com/stretchr/objx"

	qlgithub_topxeq_doc2vec_doc2vec "github.com/topxeq/qlang/lib/github.com/topxeq/doc2vec/doc2vec"

	qlgithub_aliyun_alibabacloudsdkgo_services_dysmsapi "github.com/topxeq/qlang/lib/github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"

	// qlgithub_avfs_avfs_fs_memfs "github.com/topxeq/qlang/lib/github.com/avfs/avfs/fs/memfs"
	qlgithub_topxeq_afero "github.com/topxeq/qlang/lib/github.com/topxeq/afero"

	qlgithub_topxeq_socks "github.com/topxeq/qlang/lib/github.com/topxeq/socks"

	qlgithub_topxeq_regexpx "github.com/topxeq/qlang/lib/github.com/topxeq/regexpx"

	qlgithub_topxeq_xmlx "github.com/topxeq/qlang/lib/github.com/topxeq/xmlx"

	qlgithub_topxeq_awsapi "github.com/topxeq/qlang/lib/github.com/topxeq/awsapi"

	qlgithub_cavaliercoder_grab "github.com/topxeq/qlang/lib/github.com/cavaliercoder/grab"

	qlgithub_pterm_pterm "github.com/topxeq/qlang/lib/github.com/pterm/pterm"

	qlgithub_domodwyer_mailyak "github.com/topxeq/qlang/lib/github.com/domodwyer/mailyak"

	// GUI related start

	qlgonumorg_v1_plot "github.com/topxeq/qlang/lib/gonum.org/v1/plot"
	qlgonumorg_v1_plot_plotter "github.com/topxeq/qlang/lib/gonum.org/v1/plot/plotter"
	qlgonumorg_v1_plot_plotutil "github.com/topxeq/qlang/lib/gonum.org/v1/plot/plotutil"
	qlgonumorg_v1_plot_vg "github.com/topxeq/qlang/lib/gonum.org/v1/plot/vg"

	qlgithub_scitersdk_gosciter "github.com/topxeq/qlang/lib/github.com/sciter-sdk/go-sciter"
	qlgithub_scitersdk_gosciter_window "github.com/topxeq/qlang/lib/github.com/sciter-sdk/go-sciter/window"

	"github.com/sciter-sdk/go-sciter"

	// GUI related end

	// full version related start
	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/godror/godror"

	// full version related end

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"

	// GUI related start
	// full version related start
	"github.com/sqweek/dialog"
	// full version related end
	// GUI related end

	"github.com/topxeq/tk"
)

// Non GUI related

var versionG = "1.66a"

// add tk.ToJSONX

var verboseG = false

var variableG = make(map[string]interface{})

var codeTextG = ""

var qlVMG *qlang.Qlang = nil

var varMutexG sync.Mutex

func exit(argsA ...int) {
	defer func() {
		if r := recover(); r != nil {
			tk.Printfln("exception: %v", r)

			return
		}
	}()

	if argsA == nil || len(argsA) < 1 {
		os.Exit(1)
	}

	os.Exit(argsA[0])
}

func qlEval(strA string) string {
	vmT := qlang.New()

	retG = notFoundG

	errT := vmT.SafeEval(strA)

	if errT != nil {
		return errT.Error()
	}

	rs, ok := vmT.GetVar("outG")

	if ok {
		return tk.Spr("%v", rs)
	}

	if retG != notFoundG {
		return tk.Spr("%v", retG)
	}

	return tk.ErrStrF("no result")
}

func panicIt(valueA interface{}) {
	panic(valueA)
}

func getUint64Value(v reflect.Value) uint16 {
	tk.Pl("%x", v.Interface())

	var p *uint16

	p = (v.Interface().(*uint16))

	return *p
}

func runScript(codeA string, modeA string, argsA ...string) interface{} {

	if modeA == "" || modeA == "0" || modeA == "ql" {
		vmT := qlang.New()

		// if argsA != nil && len(argsA) > 0 {
		vmT.SetVar("argsG", argsA)
		// }

		retG = notFoundG

		errT := vmT.SafeEval(codeA)

		if errT != nil {
			return errT
		}

		rs, ok := vmT.GetVar("outG")

		if ok {
			if rs != nil {
				return rs
			}
		}

		return retG
	} else {
		return tk.SystemCmd("gox", append([]string{codeA}, argsA...)...)
	}

}

func runScriptX(codeA string, argsA ...string) interface{} {

	initQLVM()

	// if argsA != nil && len(argsA) > 0 {
	qlVMG.SetVar("argsG", argsA)
	// }

	retG = notFoundG

	errT := qlVMG.SafeEval(codeA)

	if errT != nil {
		return errT
	}

	rs, ok := qlVMG.GetVar("outG")

	if ok {
		if rs != nil {
			return rs
		}
	}

	return retG

}

func runCode(codeA string, argsA ...string) interface{} {
	initQLVM()

	vmT := qlang.New()

	// if argsA != nil && len(argsA) > 0 {
	vmT.SetVar("argsG", argsA)
	// } else {
	// 	vmT.SetVar("argsG", os.Args)
	// }

	retG = notFoundG

	errT := vmT.SafeEval(codeA)

	if errT != nil {
		return errT
	}

	rs, ok := vmT.GetVar("outG")

	if ok {
		if rs != nil {
			return rs
		}
	}

	if retG != notFoundG {
		return retG
	}

	return retG
}

func getMagic(numberA int) string {
	if numberA < 0 {
		return tk.GenerateErrorString("invalid magic number")
	}

	typeT := numberA % 10

	var fcT string

	if typeT == 8 {
		fcT = tk.DownloadPageUTF8(tk.Spr("https://gitee.com/topxeq/gox/raw/master/magic/%v.gox", numberA), nil, "", 30)

	} else if typeT == 7 {
		fcT = tk.DownloadPageUTF8(tk.Spr("https: //raw.githubusercontent.com/topxeq/gox/master/magic/%v.gox", numberA), nil, "", 30)
	} else {
		return tk.GenerateErrorString("invalid magic number")
	}

	return fcT

}

func magic(numberA int, argsA ...string) interface{} {
	fcT := getMagic(numberA)

	if tk.IsErrorString(fcT) {
		return tk.ErrorStringToError(fcT)
	}

	return runCode(fcT, argsA...)

}

func NewFuncIntString(funcA *interface{}) *(func(int) string) {
	funcT := (*funcA).(*execq.Function)
	f := func(n int) string {
		return funcT.Call(execq.NewStack(), n).(string)
	}

	return &f
}

func NewFuncFloatString(funcA *interface{}) *(func(float64) string) {
	funcT := (*funcA).(*execq.Function)
	f := func(n float64) string {
		return funcT.Call(execq.NewStack(), n).(string)
	}

	return &f
}

func NewFuncStringString(funcA *interface{}) *(func(string) string) {
	funcT := (*funcA).(*execq.Function)
	f := func(s string) string {
		return funcT.Call(execq.NewStack(), s).(string)
	}

	return &f
}

func NewFuncIntError(funcA *interface{}) *(func(int) error) {
	funcT := (*funcA).(*execq.Function)
	f := func(n int) error {
		return funcT.Call(execq.NewStack(), n).(error)
	}

	return &f
}

func NewFunInterfaceError(funcA *interface{}) *(func(interface{}) error) {
	funcT := (*funcA).(*execq.Function)
	f := func(n interface{}) error {
		return funcT.Call(execq.NewStack(), n).(error)
	}

	return &f
}

func NewFuncStringError(funcA *interface{}) *(func(string) error) {
	funcT := (*funcA).(*execq.Function)
	f := func(s string) error {
		return funcT.Call(execq.NewStack(), s).(error)
	}

	return &f
}

func NewFuncStringStringError(funcA *interface{}) *(func(string) (string, error)) {
	funcT := (*funcA).(*execq.Function)
	f := func(s string) (string, error) {
		r := funcT.Call(execq.NewStack(), s).([]interface{})

		if r == nil {
			return "", tk.Errf("nil result")
		}

		if len(r) < 2 {
			return "", tk.Errf("incorrect return argument count")
		}

		if r[1] == nil {
			return r[0].(string), nil
		}

		return r[0].(string), r[1].(error)
	}

	return &f
}

func NewFuncInterfaceInterfaceError(funcA *interface{}) *(func(interface{}) (interface{}, error)) {
	funcT := (*funcA).(*execq.Function)
	f := func(s interface{}) (interface{}, error) {
		r := funcT.Call(execq.NewStack(), s).([]interface{})

		// if r == nil {
		// 	return "", tk.Errf("nil result")
		// }

		// if len(r) < 2 {
		// 	return "", tk.Errf("incorrect return argument count")
		// }

		if r[1] == nil {
			return r[0].(interface{}), nil
		}

		return r[0].(interface{}), r[1].(error)
	}

	return &f
}

func NewFuncInterfaceError(funcA *interface{}) *(func(interface{}) error) {
	funcT := (*funcA).(*execq.Function)
	f := func(s interface{}) error {
		return funcT.Call(execq.NewStack(), s).(error)
	}

	return &f
}

func NewFunc(funcA *interface{}) *(func()) {
	funcT := (*funcA).(*execq.Function)
	f := func() {
		funcT.Call(execq.NewStack())

		return
	}

	return &f
}

func NewFuncError(funcA *interface{}) *(func() error) {
	funcT := (*funcA).(*execq.Function)
	f := func() error {
		return funcT.Call(execq.NewStack()).(error)
	}

	return &f
}

func NewFuncInterface(funcA *interface{}) *(func() interface{}) {
	funcT := (*funcA).(*execq.Function)
	f := func() interface{} {
		return funcT.Call(execq.NewStack()).(interface{})
	}

	return &f
}

func NewFuncIntStringError(funcA *interface{}) *(func(int) (string, error)) {
	funcT := (*funcA).(*execq.Function)
	f := func(n int) (string, error) {
		r := funcT.Call(execq.NewStack(), n).([]interface{})

		if r == nil {
			return "", tk.Errf("nil result")
		}

		if len(r) < 2 {
			return "", tk.Errf("incorrect return argument count")
		}

		if r[1] == nil {
			return r[0].(string), nil
		}

		return r[0].(string), r[1].(error)
	}

	return &f
}

func NewFuncFloatStringError(funcA *interface{}) *(func(float64) (string, error)) {
	funcT := (*funcA).(*execq.Function)
	f := func(n float64) (string, error) {
		r := funcT.Call(execq.NewStack(), n).([]interface{})

		if r == nil {
			return "", tk.Errf("nil result")
		}

		if len(r) < 2 {
			return "", tk.Errf("incorrect return argument count")
		}

		if r[1] == nil {
			return r[0].(string), nil
		}

		return r[0].(string), r[1].(error)
	}

	return &f
}

var scriptPathG string

func importQLNonGUIPackages() {
	printValue := func(nameA string) {

		v, ok := qlVMG.GetVar(nameA)

		if !ok {
			tk.Pl("no variable by the name found: %v", nameA)
			return
		}

		tk.Pl("%v(%T): %v", nameA, v, v)

	}

	defined := func(nameA string) bool {

		_, ok := qlVMG.GetVar(nameA)

		return ok

	}

	// getPointer := func(nameA string) {

	// 	v, ok := qlVMG.GetVar(nameA)

	// 	if !ok {
	// 		tk.Pl("no variable by the name found: %v", nameA)
	// 		return
	// 	}

	// 	tk.Pl("%v(%T): %v", nameA, v, v)

	// }

	// setString := func(p *string, strA string) {
	// 	*p = strA
	// }

	var defaultExports = map[string]interface{}{
		"pass":                 tk.Pass,
		"defined":              defined,
		"eval":                 qlEval,
		"typeOf":               tk.TypeOfValue,
		"typeOfReflect":        tk.TypeOfValueReflect,
		"remove":               tk.RemoveItemsInArray,
		"pr":                   tk.Pr,
		"pln":                  tk.Pln,
		"prf":                  tk.Printf,
		"printfln":             tk.Pl,
		"pl":                   tk.Pl,
		"sprintf":              fmt.Sprintf,
		"fprintf":              fmt.Fprintf,
		"plv":                  tk.Plv,
		"plvx":                 tk.Plvx,
		"pv":                   printValue,
		"plvsr":                tk.Plvsr,
		"plerr":                tk.PlErr,
		"plExit":               tk.PlAndExit,
		"exit":                 tk.Exit,
		"setValue":             tk.SetValue,
		"getValue":             tk.GetValue,
		"setVar":               tk.SetVar,
		"getVar":               tk.GetVar,
		"bitXor":               tk.BitXor,
		"isNil":                tk.IsNil,
		"isError":              tk.IsError,
		"strToInt":             tk.StrToIntWithDefaultValue,
		"intToStr":             tk.IntToStr,
		"floatToStr":           tk.Float64ToStr,
		"toStr":                tk.ToStr,
		"toInt":                tk.ToInt,
		"toFloat":              tk.ToFloat,
		"toLower":              strings.ToLower,
		"toUpper":              strings.ToUpper,
		"checkError":           tk.CheckError,
		"checkErrorString":     tk.CheckErrorString,
		"checkErrf":            tk.CheckErrf,
		"checkErrStrf":         tk.CheckErrStrf,
		"fatalf":               tk.Fatalf,
		"isErrStr":             tk.IsErrStr,
		"errStr":               tk.ErrStr,
		"errStrf":              tk.ErrStrF,
		"getErrStr":            tk.GetErrStr,
		"errf":                 tk.Errf,
		"getInput":             tk.GetUserInput,
		"getInputf":            tk.GetInputf,
		"deepClone":            tk.DeepClone,
		"deepCopy":             tk.DeepCopyFromTo,
		"getClipText":          tk.GetClipText,
		"setClipText":          tk.SetClipText,
		"trim":                 tk.Trim,
		"run":                  runFile,
		"runCode":              runCode,
		"runScript":            runScript,
		"magic":                magic,
		"systemCmd":            tk.SystemCmd,
		"newSSHClient":         tk.NewSSHClient,
		"getParameter":         tk.GetParameterByIndexWithDefaultValue,
		"getSwitch":            tk.GetSwitchWithDefaultValue,
		"getIntSwitch":         tk.GetSwitchWithDefaultIntValue,
		"switchExists":         tk.IfSwitchExistsWhole,
		"ifSwitchExists":       tk.IfSwitchExistsWhole,
		"xmlEncode":            tk.EncodeToXMLString,
		"htmlEncode":           tk.EncodeHTML,
		"htmlDecode":           tk.DecodeHTML,
		"base64Encode":         tk.EncodeToBase64,
		"base64Decode":         tk.DecodeFromBase64,
		"md5Encode":            tk.MD5Encrypt,
		"md5":                  tk.MD5Encrypt,
		"hexEncode":            tk.StrToHex,
		"hexDecode":            tk.HexToStr,
		"jsonEncode":           tk.ObjectToJSON,
		"jsonDecode":           tk.JSONToObject,
		"simpleEncode":         tk.EncodeStringCustomEx,
		"simpleDecode":         tk.DecodeStringCustom,
		"getFormValue":         tk.GetFormValueWithDefaultValue,
		"generateJSONResponse": tk.GenerateJSONPResponseWithMore,

		"newFunc":     NewFunc,
		"scriptPathG": scriptPathG,
		"versionG":    versionG,

		// GUI related start

		// full version related start
		"edit": editFile,
		// full version related end
		// GUI related end
	}

	qlang.Import("", defaultExports)

	var imiscExports = map[string]interface{}{
		"NewFunc":                        NewFunc,
		"NewFuncError":                   NewFuncError,
		"NewFuncInterface":               NewFuncInterface,
		"NewFuncInterfaceError":          NewFuncInterfaceError,
		"NewFuncInterfaceInterfaceError": NewFuncInterfaceInterfaceError,
		"NewFuncIntString":               NewFuncIntString,
		"NewFuncIntError":                NewFuncIntError,
		"NewFuncFloatString":             NewFuncFloatString,
		"NewFuncFloatStringError":        NewFuncFloatStringError,
		"NewFuncStringString":            NewFuncStringString,
		"NewFuncStringError":             NewFuncStringError,
		"NewFuncStringStringError":       NewFuncStringStringError,
		"NewFuncIntStringError":          NewFuncIntStringError,
	}

	qlang.Import("imisc", imiscExports)

	qlang.Import("archive_zip", qlarchivezip.Exports)
	qlang.Import("bufio", qlbufio.Exports)
	qlang.Import("bytes", qlbytes.Exports)

	qlang.Import("crypto", qlcrypto.Exports)
	qlang.Import("crypto_aes", qlcryptoaes.Exports)
	qlang.Import("crypto_cipher", qlcryptocipher.Exports)
	qlang.Import("crypto_hmac", qlcryptohmac.Exports)
	qlang.Import("crypto_md5", qlcryptomd5.Exports)
	qlang.Import("crypto_rand", qlcryptorand.Exports)
	qlang.Import("crypto_rsa", qlcryptorsa.Exports)
	qlang.Import("crypto_sha256", qlcryptosha256.Exports)
	qlang.Import("crypto_sha1", qlcryptosha1.Exports)
	qlang.Import("crypto_x509", qlcryptox509.Exports)

	qlang.Import("database_sql", qldatabasesql.Exports)

	qlang.Import("encoding_pem", qlencodingpem.Exports)
	qlang.Import("encoding_base64", qlencodingbase64.Exports)
	qlang.Import("encoding_binary", qlencodingbinary.Exports)
	qlang.Import("encoding_csv", qlencodingcsv.Exports)
	qlang.Import("encoding_gob", qlencodinggob.Exports)
	qlang.Import("encoding_hex", qlencodinghex.Exports)
	qlang.Import("encoding_json", qlencodingjson.Exports)
	qlang.Import("encoding_xml", qlencodingxml.Exports)

	qlang.Import("errors", qlerrors.Exports)

	qlang.Import("flag", qlflag.Exports)
	qlang.Import("fmt", qlfmt.Exports)

	qlang.Import("hash_fnv", qlhashfnv.Exports)

	qlang.Import("html", qlhtml.Exports)
	qlang.Import("html_template", qlhtmltemplate.Exports)

	qlang.Import("image", qlimage.Exports)
	qlang.Import("image_color", qlimage_color.Exports)
	qlang.Import("image_color_palette", qlimage_color_palette.Exports)
	qlang.Import("image_draw", qlimage_draw.Exports)
	qlang.Import("image_gif", qlimage_gif.Exports)
	qlang.Import("image_jpeg", qlimage_jpeg.Exports)
	qlang.Import("image_png", qlimage_png.Exports)

	qlang.Import("io", qlio.Exports)
	qlang.Import("io_ioutil", qlioioutil.Exports)
	qlang.Import("io_fs", qlio_fs.Exports)

	qlang.Import("log", qllog.Exports)

	qlang.Import("math", qlmath.Exports)
	qlang.Import("math_big", qlmathbig.Exports)
	qlang.Import("math_bits", qlmathbits.Exports)
	qlang.Import("math_rand", qlmathrand.Exports)

	qlang.Import("net", qlnet.Exports)
	qlang.Import("net_http", qlnethttp.Exports)
	qlang.Import("net_http_cookiejar", qlnet_http_cookiejar.Exports)
	qlang.Import("net_http_httputil", qlnet_http_httputil.Exports)
	qlang.Import("net_mail", qlnet_mail.Exports)
	qlang.Import("net_rpc", qlnet_rpc.Exports)
	qlang.Import("net_rpc_jsonrpc", qlnet_rpc_jsonrpc.Exports)
	qlang.Import("net_smtp", qlnet_smtp.Exports)
	qlang.Import("net_url", qlneturl.Exports)

	qlang.Import("os", qlos.Exports)
	qlang.Import("os_exec", qlos_exec.Exports)
	qlang.Import("os_signal", qlos_signal.Exports)
	qlang.Import("os_user", qlos_user.Exports)
	qlang.Import("path", qlpath.Exports)
	qlang.Import("path_filepath", qlpathfilepath.Exports)

	qlang.Import("reflect", qlreflect.Exports)
	qlang.Import("regexp", qlregexp.Exports)

	qlang.Import("runtime", qlruntime.Exports)
	qlang.Import("runtime_debug", qlruntimedebug.Exports)

	qlang.Import("sort", qlsort.Exports)
	qlang.Import("strconv", qlstrconv.Exports)
	qlang.Import("strings", qlstrings.Exports)
	qlang.Import("sync", qlsync.Exports)

	qlang.Import("text_template", qltext_template.Exports)
	qlang.Import("time", qltime.Exports)

	qlang.Import("unicode", qlunicode.Exports)
	qlang.Import("unicode_utf8", qlunicode_utf8.Exports)

	// 3rd party

	qlang.Import("github_topxeq_tk", qlgithubtopxeqtk.Exports)
	qlang.Import("tk", qlgithubtopxeqtk.Exports)
	qlang.Import("github_topxeq_imagetk", qlgithubtopxeqimagetk.Exports)
	qlang.Import("imagetk", qlgithubtopxeqimagetk.Exports)

	qlang.Import("github_beevik_etree", qlgithubbeeviketree.Exports)
	qlang.Import("etree", qlgithubbeeviketree.Exports)
	qlang.Import("github_topxeq_sqltk", qlgithubtopxeqsqltk.Exports)
	qlang.Import("sqltk", qlgithubtopxeqsqltk.Exports)

	qlang.Import("github_topxeq_xmlx", qlgithub_topxeq_xmlx.Exports)

	qlang.Import("github_topxeq_awsapi", qlgithub_topxeq_awsapi.Exports)

	qlang.Import("github_cavaliercoder_grab", qlgithub_cavaliercoder_grab.Exports)

	qlang.Import("github_pterm_pterm", qlgithub_pterm_pterm.Exports)

	qlang.Import("github_domodwyer_mailyak", qlgithub_domodwyer_mailyak.Exports)

	// GUI related start

	InitGiu()

	qlgithub_scitersdk_gosciter.Exports["NewValue"] = sciter.NewValue
	qlgithub_scitersdk_gosciter.Exports["NullValue"] = sciter.NullValue

	qlgithub_scitersdk_gosciter.Exports["NewScnLoadDataFunc"] = NewScnLoadDataFunc
	qlgithub_scitersdk_gosciter.Exports["NewScnDataLoaded"] = NewScnDataLoaded

	qlang.Import("github_scitersdk_gosciter", qlgithub_scitersdk_gosciter.Exports)
	qlang.Import("github_scitersdk_gosciter_window", qlgithub_scitersdk_gosciter_window.Exports)

	qlang.Import("gonumorg_v1_plot", qlgonumorg_v1_plot.Exports)
	qlang.Import("plot", qlgonumorg_v1_plot.Exports)
	qlang.Import("gonumorg_v1_plot_plotter", qlgonumorg_v1_plot_plotter.Exports)
	qlang.Import("plot_plotter", qlgonumorg_v1_plot_plotter.Exports)
	qlang.Import("gonumorg_v1_plot_plotutil", qlgonumorg_v1_plot_plotutil.Exports)
	qlang.Import("plot_plotutil", qlgonumorg_v1_plot_plotutil.Exports)
	qlang.Import("gonumorg_v1_plot_vg", qlgonumorg_v1_plot_vg.Exports)
	qlang.Import("plot_vg", qlgonumorg_v1_plot_vg.Exports)

	// InitBlink()

	InitSysspec()

	// GUI related end

	qlang.Import("github_fogleman_gg", qlgithub_fogleman_gg.Exports)
	qlang.Import("gg", qlgithub_fogleman_gg.Exports)

	qlang.Import("github_360EntSecGroupSkylar_excelize", qlgithub_360EntSecGroupSkylar_excelize.Exports)

	qlang.Import("github_kbinani_screenshot", qlgithub_kbinani_screenshot.Exports)

	qlang.Import("github_stretchr_objx", qlgithub_stretchr_objx.Exports)

	qlang.Import("github_topxeq_doc2vec_doc2vec", qlgithub_topxeq_doc2vec_doc2vec.Exports)

	qlang.Import("github_aliyun_alibabacloudsdkgo_services_dysmsapi", qlgithub_aliyun_alibabacloudsdkgo_services_dysmsapi.Exports)
	qlang.Import("aliyunsms", qlgithub_aliyun_alibabacloudsdkgo_services_dysmsapi.Exports)

	// qlang.Import("github_avfs_avfs_fs_memfs", qlgithub_avfs_avfs_fs_memfs.Exports)
	qlang.Import("github_topxeq_afero", qlgithub_topxeq_afero.Exports)
	qlang.Import("memfs", qlgithub_topxeq_afero.Exports)

	qlang.Import("github_topxeq_socks", qlgithub_topxeq_socks.Exports)

	qlang.Import("github_topxeq_regexpx", qlgithub_topxeq_regexpx.Exports)

}

func showHelp() {
	tk.Pl("Gox by TopXeQ V%v\n", versionG)

	tk.Pl("Usage: gox [-v|-h] test.gox, ...\n")
	tk.Pl("or just gox without arguments to start REPL instead.\n")

}

// func compileSource(srcA string) string {
// 	vmT := qlang.New()

// 	tk.Pl("vmT: %v", vmT)

// 	errT := vmT.TXCompile(srcA)

// 	if errT != nil {
// 		return errT.Error()
// 	}

// 	tk.Pl("vmT after: %v", vmT)

// 	return ""

// }

func runInteractiveQlang() int {
	var following bool
	var source string
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

		retG = notFoundG

		err := qlVMG.SafeEval(source)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			following = false
			source = ""
			continue
		}

		if retG != notFoundG {
			fmt.Println(retG)
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

// Non GUI related end

// GUI related start

func NewScnLoadDataFunc(funcA *interface{}) *(func(*sciter.ScnLoadData) int) {
	funcT := (*funcA).(*execq.Function)
	f := func(dataA *sciter.ScnLoadData) int {
		return funcT.Call(execq.NewStack(), dataA).(int)
	}

	return &f
}

func NewScnDataLoaded(funcA *interface{}) *(func(*sciter.ScnDataLoaded) int) {
	funcT := (*funcA).(*execq.Function)
	f := func(dataA *sciter.ScnDataLoaded) int {
		return funcT.Call(execq.NewStack(), dataA).(int)
	}

	return &f
}

// full version related start

// var vgInch = float64(vg.Inch)

// full version related end

func importQLGUIPackages() {
	// full version related start
	// var plotExports = map[string]interface{}{
	// 	"New": plot.New,
	// 	// "SetTitleText":  plot.SetTitleText,
	// 	"NewXY":         newPlotXY,
	// 	"AddLinePoints": plotutil.AddLinePoints,
	// 	"Inch":          vgInch,

	// 	"XYs": specq.StructOf((*plotter.XYs)(nil)),
	// 	"XY":  specq.StructOf((*plotter.XY)(nil)),
	// }

	// qlang.Import("plot", plotExports)

	// var imagetkExports = map[string]interface{}{
	// 	"NewImageTK": imagetk.NewImageTK,
	// }

	// qlang.Import("imagetk", imagetkExports)
	InitGiuExports()

	// full version related end
}

// full version related start
func getConfirmGUI(titleA string, messageA string) bool {
	return dialog.Message("%v", messageA).Title(titleA).YesNo()
}

func simpleInfo(titleA string, messageA string) {
	dialog.Message("%v", messageA).Title(titleA).Info()
}

func simpleError(titleA string, messageA string) {
	dialog.Message("%v", messageA).Title(titleA).Error()
}

// filename, err := dialog.File().Filter("XML files", "xml").Title("Export to XML").Save()
func selectFileToSaveGUI(titleA string, filterNameA string, filterTypeA string) string {
	fileNameT, errT := dialog.File().Filter(filterNameA, filterTypeA).Title(titleA).Save()

	if errT != nil {
		return tk.GenerateErrorStringF("failed: %v", errT)
	}

	return fileNameT
}

// fileNameT, errT := dialog.File().Filter("Mp3 audio file", "mp3").Load()
func selectFileGUI(titleA string, filterNameA string, filterTypeA string) string {
	fileNameT, errT := dialog.File().Filter(filterNameA, filterTypeA).Title(titleA).Load()

	if errT != nil {
		return tk.GenerateErrorStringF("failed: %v", errT)
	}

	return fileNameT
}

// directory, err := dialog.Directory().Title("Load images").Browse()
func selectDirectoryGUI(titleA string) string {
	directoryT, errT := dialog.Directory().Title(titleA).Browse()

	if errT != nil {
		return tk.GenerateErrorStringF("failed: %v", errT)
	}

	return directoryT
}

// full version related end

// GUI related end

func runFile(argsA ...string) interface{} {
	lenT := len(argsA)

	// full version related start
	// GUI related start

	if lenT < 1 {
		rs := selectFileGUI("Please select file to run...", "All files", "*")

		if tk.IsErrorString(rs) {
			return tk.Errf("Failed to load file: %v", tk.GetErrorString(rs))
		}

		fcT := tk.LoadStringFromFile(rs)

		if tk.IsErrorString(fcT) {
			return tk.Errf("Invalid file content: %v", tk.GetErrorString(fcT))
		}

		return runScript(fcT, "")

	}
	// GUI related end
	// full version related end

	if lenT < 1 {
		return nil
	}

	fcT := tk.LoadStringFromFile(argsA[0])

	if tk.IsErrorString(fcT) {
		return tk.Errf("Invalid file content: %v", tk.GetErrorString(fcT))
	}

	return runScript(fcT, "", argsA[1:]...)
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
		tk.Pl("Gox by TopXeQ V%v", versionG)
		return nil
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
		rs := runScriptX(editFileScriptG, argsT...)

		if rs != notFoundG && rs != nil {
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
			rs := tk.DownloadFile("http://scripts.frenchfriend.net/pub/sciter.dll", applicationPathT, "sciter.dll", false)

			if tk.IsErrorString(rs) {

				return tk.Errf("failed to download Sciter DLL file.")
			}

			tk.Pl("Sciter DLL downloaded to application path.")

			return nil
		}
	}

	ifClipT := tk.IfSwitchExistsWhole(argsT, "-clip")
	ifEmbedT := (codeTextG != "") && (!tk.IfSwitchExistsWhole(argsT, "-noembed"))

	if scriptT == "" && (!ifClipT) && (!ifEmbedT) {

		// autoPathT := filepath.Join(tk.GetApplicationPath(), "auto.gox")
		// autoGxbPathT := filepath.Join(tk.GetApplicationPath(), "auto.gxb")
		autoPathT := "auto.gox"
		autoGxbPathT := "auto.gxb"

		if tk.IfFileExists(autoPathT) {
			scriptT = autoPathT
		} else if tk.IfFileExists(autoGxbPathT) {
			scriptT = autoGxbPathT
		} else {
			initQLVM()

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

	ifExampleT := tk.IfSwitchExistsWhole(argsT, "-example")
	ifGoPathT := tk.IfSwitchExistsWhole(argsT, "-gopath")
	ifLocalT := tk.IfSwitchExistsWhole(argsT, "-local")
	ifAppPathT := tk.IfSwitchExistsWhole(argsT, "-apppath")
	ifRemoteT := tk.IfSwitchExistsWhole(argsT, "-remote")
	ifCloudT := tk.IfSwitchExistsWhole(argsT, "-cloud")
	sshT := tk.GetSwitchWithDefaultValue(argsT, "-ssh=", "")
	ifViewT := tk.IfSwitchExistsWhole(argsT, "-view")
	// ifCompileT := tk.IfSwitchExistsWhole(argsT, "-compile")

	verboseG = tk.IfSwitchExistsWhole(argsT, "-verbose")

	ifMagicT := false
	magicNumberT, errT := tk.StrToIntE(scriptT)

	if errT == nil {
		ifMagicT = true
	}

	var fcT string

	if ifMagicT {
		fcT = getMagic(magicNumberT)

		scriptPathG = ""
	} else if ifExampleT {
		if (!tk.EndsWith(scriptT, ".gox")) && (!tk.EndsWith(scriptT, ".ql")) {
			scriptT += ".gox"
		}
		fcT = tk.DownloadPageUTF8("https://gitee.com/topxeq/gox/raw/master/scripts/"+scriptT, nil, "", 30)

		scriptPathG = ""
	} else if ifRemoteT {
		fcT = tk.DownloadPageUTF8(scriptT, nil, "", 30)

		scriptPathG = ""
	} else if ifClipT {
		fcT = tk.GetClipText()

		scriptPathG = ""
	} else if ifEmbedT {
		fcT = codeTextG

		scriptPathG = ""
	} else if ifCloudT {
		if (!tk.EndsWith(scriptT, ".gox")) && (!tk.EndsWith(scriptT, ".ql")) {
			scriptT += ".gox"
		}

		basePathT, errT := tk.EnsureBasePath("gox")

		gotT := false

		if errT == nil {
			cfgPathT := tk.JoinPath(basePathT, "cloud.cfg")

			cfgStrT := tk.Trim(tk.LoadStringFromFile(cfgPathT))

			if !tk.IsErrorString(cfgStrT) {
				fcT = tk.DownloadPageUTF8(cfgStrT+scriptT, nil, "", 30)

				gotT = true
			}

		}

		if !gotT {
			fcT = tk.DownloadPageUTF8(scriptT, nil, "", 30)
		}

		scriptPathG = ""
	} else if sshT != "" {
		if (!tk.EndsWith(scriptT, ".gox")) && (!tk.EndsWith(scriptT, ".ql")) {
			scriptT += ".gox"
		}

		fcT = downloadStringFromSSH(sshT, scriptT)

		if tk.IsErrorString(fcT) {

			return tk.Errf("failed to get script from SSH: %v", tk.GetErrorString(fcT))
		}

		scriptPathG = ""
	} else if ifGoPathT {
		if (!tk.EndsWith(scriptT, ".gox")) && (!tk.EndsWith(scriptT, ".ql")) {
			scriptT += ".gox"
		}

		scriptPathG = filepath.Join(tk.GetEnv("GOPATH"), "src", "github.com", "topxeq", "gox", "scripts", scriptT)

		fcT = tk.LoadStringFromFile(scriptPathG)
	} else if ifAppPathT {
		if (!tk.EndsWith(scriptT, ".gox")) && (!tk.EndsWith(scriptT, ".ql")) {
			scriptT += ".gox"
		}

		scriptPathG = filepath.Join(tk.GetApplicationPath(), scriptT)

		fcT = tk.LoadStringFromFile(scriptPathG)
	} else if ifLocalT {
		if (!tk.EndsWith(scriptT, ".gox")) && (!tk.EndsWith(scriptT, ".ql")) {
			scriptT += ".gox"
		}

		localPathT := getCfgString("localScriptPath.cfg")

		if tk.IsErrorString(localPathT) {
			// tk.Pl("failed to get local path: %v", tk.GetErrorString(localPathT))

			return tk.Errf("failed to get local path: %v", tk.GetErrorString(localPathT))
		}

		// if tk.GetEnv("GOXVERBOSE") == "true" {
		// 	tk.Pl("Try to load script from %v", filepath.Join(localPathT, scriptT))
		// }

		scriptPathG = filepath.Join(localPathT, scriptT)

		fcT = tk.LoadStringFromFile(scriptPathG)
	} else {
		scriptPathG = scriptT
		fcT = tk.LoadStringFromFile(scriptT)
	}

	if tk.IsErrorString(fcT) {
		return tk.Errf("failed to load script from %v: %v", scriptT, tk.GetErrorString(fcT))
	}

	if tk.StartsWith(fcT, "//TXDEF#") {
		if decryptRunCodeT == "" {
			tk.Prf("Password: ")
			decryptRunCodeT = tk.Trim(tk.GetInputBufferedScan())

			// fcT = fcT[8:]
		}
	}

	if decryptRunCodeT != "" {
		fcT = tk.DecryptStringByTXDEF(fcT, decryptRunCodeT)
	}

	if ifViewT {
		tk.Pl("%v", fcT)

		return nil
	}

	// if ifCompileT {
	// 	initQLVM()

	// 	qlVMG.SetVar("argsG", argsT)

	// 	retG = notFoundG

	// 	endT, errT := qlVMG.SafeCl([]byte(fcT), "")
	// 	if errT != nil {

	// 		// tk.Pl()

	// 		// f, l := qlVMG.Code.Line(qlVMG.Code.Reserve().Next())
	// 		// tk.Pl("Next line: %v, %v", f, l)

	// 		return tk.Errf("failed to compile script(%v) error: %v\n", scriptT, errT)
	// 	}

	// 	tk.Pl("endT: %v", endT)

	// 	errT = qlVMG.DumpEngine()

	// 	if errT != nil {
	// 		return tk.Errf("failed to dump engine: %v\n", errT)
	// 	}

	// 	tk.Plvsr(qlVMG.Cpl.GetCode().Len(), qlVMG.Run())

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

	initQLVM()

	qlVMG.SetVar("argsG", argsT)

	retG = notFoundG

	errT = qlVMG.SafeEval(fcT)
	if errT != nil {

		// tk.Pl()

		// f, l := qlVMG.Code.Line(qlVMG.Code.Reserve().Next())
		// tk.Pl("Next line: %v, %v", f, l)

		return tk.Errf("failed to execute script(%v) error: %v\n", scriptT, errT)
	}

	rs, ok := qlVMG.GetVar("outG")

	if ok {
		if rs != nil {
			return rs
		}
	}

	return retG
}

// init the main VM

var retG interface{}
var notFoundG = interface{}(errors.New("not found"))

func initQLVM() {
	if qlVMG == nil {
		qlang.SetOnPop(func(v interface{}) {
			retG = v
		})

		// qlang.SetDumpCode("1")

		importQLNonGUIPackages()

		// GUI related start

		importQLGUIPackages()

		// GUI related end

		qlVMG = qlang.New()
	}
}

func downloadStringFromSSH(sshA string, filePathA string) string {
	aryT := tk.Split(sshA, ":")

	basePathT, errT := tk.EnsureBasePath("gox")

	if errT != nil {
		return tk.GenerateErrorStringF("failed to find base path: %v", errT)
	}

	if len(aryT) != 5 {
		aryT = tk.Split(tk.LoadStringFromFile(tk.JoinPath(basePathT, "ssh.cfg"))+filePathA, ":")

		if len(aryT) != 5 {
			return tk.ErrStrF("invalid ssh config: %v", "")
		}

	}

	clientT, errT := tk.NewSSHClient(aryT[0], tk.StrToIntWithDefaultValue(aryT[1], 22), aryT[2], aryT[3])

	if errT != nil {
		return tk.ErrToStrF("failed to create SSH client:", errT)
	}

	tmpPathT := tk.JoinPath(basePathT, "tmp")

	errT = tk.EnsureMakeDirsE(tmpPathT)

	if errT != nil {
		return tk.ErrToStrF("failed to create tmp dir:", errT)
	}

	tmpFileT, errT := tk.CreateTempFile(tmpPathT, "")

	if errT != nil {
		return tk.ErrToStrF("failed to create tmp dir:", errT)
	}

	defer os.Remove(tmpFileT)

	errT = clientT.Download(aryT[4], tmpFileT)

	if errT != nil {
		return tk.ErrToStrF("failed to download file:", errT)
	}

	fcT := tk.LoadStringFromFile(tmpFileT)

	return fcT
}

func getCfgString(fileNameA string) string {
	basePathT, errT := tk.EnsureBasePath("gox")

	if errT == nil {
		cfgPathT := tk.JoinPath(basePathT, fileNameA)

		cfgStrT := tk.Trim(tk.LoadStringFromFile(cfgPathT))

		if !tk.IsErrorString(cfgStrT) {
			return cfgStrT
		}

		return tk.ErrStrF("failed to get config string: %v", tk.GetErrorString(cfgStrT))

	}

	return tk.ErrStrF("failed to get config string")
}

var editFileScriptG = `
sciter = github_scitersdk_gosciter
window = github_scitersdk_gosciter_window

htmlT := ` + "`" + `
<!DOCTYPE html>
<html>
<head>
	<meta charset="utf-8">
	<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
	<title>Gox Editor</title>
    <style>
    
    plaintext {
      padding:0;
      flow:vertical;
      behavior:plaintext;
      background:#333; border:1px solid #333;
      color:white;
	  overflow:scroll-indicator;
      font-rendering-mode:snap-pixel;
      size:*; 
      tab-size: 4;
    }
    plaintext > text {
      font-family:monospace;
      white-space: pre-wrap;
      background:white;
      color:black;
      margin-left: 3em;
      padding-left: 4dip;
      cursor:text;
      display:list-item;
      list-style-type: index;
      list-marker-color:#aaa;
    }
    plaintext > text:last-child {
      padding-bottom:*;
    }    
    
    plaintext > text:nth-child(10n) {
      list-marker-color:#fff;
    }
    
    
    </style>


	<script type="text/tiscript">
		function colorize() 
		{
			const apply = Selection.applyMark; // shortcut
			const isEditor = this.tag == "plaintext";
			
			// forward declarations:
			var doStyle;
			var doScript;

			// markup colorizer  
			function doMarkup(tz) 
			{
					var bnTagStart = null;
					var tagScript = false;
					var tagScriptType = false;
					var tagStyle = false;
					var textElement;
				
				while(var tt = tz.token()) {
				if( isEditor && tz.element != textElement )       
				{
					textElement = tz.element;
					textElement.attributes["type"] = "markup";
				}
				//stdout.println(tt,tz.attr,tz.value);
				switch(tt) {
					case #TAG-START: {    
						bnTagStart = tz.tokenStart; 
						const tag = tz.tag;
						tagScript = tag == "script";
						tagStyle  = tag == "style";
					} break;
					case #TAG-HEAD-END: {
						apply(bnTagStart,tz.tokenEnd,"tag"); 
						if( tagScript ) { tz.push(#source,"</sc"+"ript>"); doScript(tz, tagScriptType, true); }
						else if( tagStyle ) { tz.push(#source,"</style>"); doStyle(tz, true); }
					} break;
					case #TAG-END:      apply(tz.tokenStart,tz.tokenEnd,"tag"); break;  
					case #TAG-ATTR:     if( tagScript && tz.attr == "type") tagScriptType = tz.value; 
										if( tz.attr == "id" ) apply(tz.tokenStart,tz.tokenEnd,"tag-id"); 
										break;
				}
				}
			}
			
			// script colorizer
			doScript = function(tz, typ, embedded = false) 
			{
				const KEYWORDS = 
				{
				"type"    :true, "function":true, "var"       :true,"if"       :true,
				"else"    :true, "while"   :true, "return"    :true,"for"      :true,
				"break"   :true, "continue":true, "do"        :true,"switch"   :true,
				"case"    :true, "default" :true, "null"      :true,"super"    :true,
				"new"     :true, "try"     :true, "catch"     :true,"finally"  :true,
				"throw"   :true, "typeof"  :true, "instanceof":true,"in"       :true,
				"property":true, "const"   :true, "get"       :true,"set"      :true,
				"include" :true, "like"    :true, "class"     :true,"namespace":true,
				"this"    :true, "assert"  :true, "delete"    :true,"otherwise":true,
				"with"    :true, "__FILE__":true, "__LINE__"  :true,"__TRACE__":true,
				"debug"   :true, "await"   :true 
				};
				
				const LITERALS = { "true": true, "false": true, "null": true, "undefined": true };
				
				var firstElement;
				var lastElement;
			
				while:loop(var tt = tz.token()) {
				var el = tz.element;
				if( !firstElement ) firstElement = el;
				lastElement = el;
				switch(tt) 
				{
					case #NUMBER:       apply(tz.tokenStart,tz.tokenEnd,"number"); break; 
					case #NUMBER-UNIT:  apply(tz.tokenStart,tz.tokenEnd,"number-unit"); break; 
					case #STRING:       apply(tz.tokenStart,tz.tokenEnd,"string"); break;
					case #NAME:         
					{
					var val = tz.value;
					if( val[0] == '#' )
						apply(tz.tokenStart,tz.tokenEnd, "symbol"); 
					else if(KEYWORDS[val]) 
						apply(tz.tokenStart,tz.tokenEnd, "keyword"); 
					else if(LITERALS[val]) 
						apply(tz.tokenStart,tz.tokenEnd, "literal"); 
					break;
					}
					case #COMMENT:      apply(tz.tokenStart,tz.tokenEnd,"comment"); break;
					case #END-OF-ISLAND:  
					// got </scr ipt>
					tz.pop(); //pop tokenizer layer
					break loop;
				}
				}
				if(isEditor && embedded) {
				for( var el = firstElement; el; el = el.next ) {
					el.attributes["type"] = "script";
					if( el == lastElement )
					break;
				}
				}
			};
			
			doStyle = function(tz, embedded = false) 
			{
				const KEYWORDS = 
				{
				"rgb":true, "rgba":true, "url":true, 
				"@import":true, "@media":true, "@set":true, "@const":true
				};
				
				const LITERALS = { "inherit": true };
				
				var firstElement;
				var lastElement;
				
				while:loop(var tt = tz.token()) {
				var el = tz.element;
				if( !firstElement ) firstElement = el;
				lastElement = el;
				switch(tt) 
				{
					case #NUMBER:       apply(tz.tokenStart,tz.tokenEnd,"number"); break; 
					case #NUMBER-UNIT:  apply(tz.tokenStart,tz.tokenEnd,"number-unit"); break; 
					case #STRING:       apply(tz.tokenStart,tz.tokenEnd,"string"); break;
					case #NAME:         
					{
					var val = tz.value;
					if( val[0] == '#' )
						apply(tz.tokenStart,tz.tokenEnd, "symbol"); 
					else if(KEYWORDS[val]) 
						apply(tz.tokenStart,tz.tokenEnd, "keyword"); 
					else if(LITERALS[val]) 
						apply(tz.tokenStart,tz.tokenEnd, "literal"); 
					break;
					}
					case #COMMENT:      apply(tz.tokenStart,tz.tokenEnd,"comment"); break;
					case #END-OF-ISLAND:  
					// got </sc ript>
					tz.pop(); //pop tokenizer layer
					break loop;
				}
				}
				if(isEditor && embedded) {
				for( var el = firstElement; el; el = el.next ) {
					el.attributes["type"] = "style";
					if( el == lastElement )
					break;
				}
				}
			};
			
			var me = this;
			
			function doIt() { 
			
				var typ = me.attributes["type"];

				var syntaxKind = typ like "*html" || typ like "*xml" ? #markup : #source;
				var syntax = typ like "*css"? #style : #script;
				
				var tz = new Tokenizer( me, syntaxKind );
			
				if( syntaxKind == #markup )
				doMarkup(tz);
				else if( syntax == #style )
				doStyle(tz);
				else 
				doScript(tz,typ);
			}
			
			doIt();
			
			// redefine value property
			this[#value] = property(v) {
				get { return this.state.value; }
				set { this.state.value = v; doIt(); }
			};
			
			this.load = function(text,sourceType) 
			{
				this.attributes["type"] = sourceType;
				if( !isEditor )
				text = text.replace(/\r\n/g,"\n"); 
				this.state.value = text; 
				doIt();
			};
			
			this.sourceType = property(v) {
				get { return this.attributes["type"]; }
				set { this.attributes["type"] = v; doIt(); }
			};
			if (isEditor)
					this.on("change", function() {
						this.timer(40ms,doIt);
					});
			

		}
	</script>
	<style>

		@set colorizer < std-plaintext 
		{
			:root { aspect: colorize; }
			
			text { white-space:pre;  display:list-item; list-style-type: index; list-marker-color:#aaa; }
			/*markup*/  
			text::mark(tag) { color: olive; } /*background-color: #f0f0fa;*/
			text::mark(tag-id) { color: red; } /*background-color: #f0f0fa;*/

			/*source*/  
			text::mark(number) { color: brown; }
			text::mark(number-unit) { color: brown; }
			text::mark(string) { color: teal; }
			text::mark(keyword) { color: blue; }
			text::mark(symbol) { color: brown; }
			text::mark(literal) { color: brown; }
			text::mark(comment) { color: green; }
			
			text[type=script] {  background-color: #FFFAF0; }
			text[type=markup] {  background-color: #FFF;  }
			text[type=style]  {  background-color: #FAFFF0; }
		}

		plaintext[type] {
			style-set: colorizer;
		}

		@set element-colorizer 
		{
			:root { 
				aspect: colorize; 
				background-color: #fafaff;
					padding:4dip;
					border:1dip dashed #bbb;
				}
			
			/*markup*/  
			:root::mark(tag) { color: olive; } 
			:root::mark(tag-id) { color: red; }

			/*source*/  
			:root::mark(number) { color: brown; }
			:root::mark(number-unit) { color: brown; }
			:root::mark(string) { color: teal; }
			:root::mark(keyword) { color: blue; }
			:root::mark(symbol) { color: brown; }
			:root::mark(literal) { color: brown; }
			:root::mark(comment) { color: green; }
			}

			pre[type] {
			style-set: element-colorizer;
		}

	</style>
	<script type="text/tiscript">
		// if (view.connectToInspector) {
		// 	view.connectToInspector(rootElement, inspectorIpAddress);
		// }

		//stdout.println("__FOLDER__:", __FOLDER__);
		//stdout.println("__FILE__:", __FILE__);

		function isErrStr(strA) {
			if (strA.substr(0, 6) == "TXERROR:") {
				return true;
			}

			return false;
		}

		function getErrStr(strA) {
			if (strA.substr(0, 6) == "TXERROR:") {
				return strA.substr(6);
			}

			return strA;
		}

		function getConfirm(titelA, msgA) {
			var result = view.msgbox { 
				type:#question,
				title: titelA,
				content: msgA, 
				//buttons:[#yes,#no]
				buttons: [
					{id:#yes,text:"Ok",role:"default-button"},
					{id:#cancel,text:"Cancel",role:"cancel-button"}]                               
				};

			return result;
		}

		function showInfo(titelA, msgA) {
			var result = view.msgbox { 
				type:#information,
				title: titelA,
				content: msgA, 
				//buttons:[#yes,#no]
				buttons: [
					{id:#cancel,text:"Close",role:"cancel-button"}]                               
				};

			return result;
		}

		function showError(titelA, msgA) {
			var result = view.msgbox { 
				type:#alert,
				title: titelA,
				content: msgA, 
				//buttons:[#yes,#no]
				buttons: [
					{id:#cancel,text:"Close",role:"cancel-button"}]                               
				};

			return result;
		}

		function getScreenWH() {
//			view.prints(String.printf("screenBoxO: %v, %v", 1, 2));
			var (w, h) = view.screenBox(#frame, #dimension)
//			view.prints(String.printf("screenBox: %v, %v", w, h));

			view.move((w-800)/2, (h-600)/2, 800, 600);

			return String.printf("%v|%v", w, h);
		}

		var editFileNameG = "";
		var editFileCleanFlagG = "";

		function updateFileName() {
			$(#fileNameLabelID).html = (editFileNameG + editFileCleanFlagG);
		}

		function selectFileJS() {
			//var fn = view.selectFile(#open, "Gotx Files (*.gt,*.go)|*.gt;*.go|All Files (*.*)|*.*" , "gotx" );
			var fn = view.selectFile(#open);
			view.prints(String.printf("fn: %v", fn));
			//view.prints(String.printf("screenBox: %v", view.screenBox(#frame, #dimension)));

			if (fn == undefined) {
				return;
			}

			var fileNameT = URL.toPath(fn);

			var rs = view.loadStringFromFile(fileNameT);

			if (isErrStr(rs)) {
				showError("Error", String.printf("Failed to load file content: %v", getErrStr(rs)));
				return;
			}

			$(plaintext).attributes["type"] = "text/script";

			$(plaintext).value = rs;

			editFileNameG = fileNameT;

			editFileCleanFlagG = "";

			updateFileName();

			// return fn;
		}

		function editFileLoadClick() {
			if (editFileCleanFlagG != "") {
			
				var rs = getConfirm("Please confirm", "File modified, load another file anyway?");

				if (rs != #yes) {
					return;
				}

			}

			selectFileJS();
		}

		function editFileSaveAsClick() {
			var fn = view.selectFile(#save);
			view.prints(String.printf("fn: %v", fn));

			if (fn == undefined) {
				return;
			}

			var fileNameT = URL.toPath(fn);

			var textT = $(plaintext).value;

			var rs = view.saveStringToFile(textT, fileNameT);

			if (isErrStr(rs)) {
				showError("Error", String.printf("Failed to save file content: %v", getErrStr(rs)));
				return;
			}

			editFileNameG = fileNameT;
			editFileCleanFlagG = "";
			updateFileName();

			showInfo("Info", "Saved.");

		}

		function editFileSaveClick() {
			if (editFileNameG == "") {
				editFileSaveAsClick();

				return;
			}

			var textT = $(plaintext).value;

			var rs = view.saveStringToFile(textT, editFileNameG);

			if (isErrStr(rs)) {
				showError("Error", String.printf("Failed to save file content: %v", getErrStr(rs)));
				return;
			}

			editFileCleanFlagG = "";
			updateFileName();

			showInfo("Info", "Saved.");
		}

		function editRunClick() {
			view.close();
			// view.exit();
		}

		function getInput(msgA) {
			var res = view.dialog({ 
				html: ` + "`+\"`\"+`" + `
				<html>
				<body>
				  <center>
					  <div style="margin-top: 10px; margin-bottom: 10px;">
						  <span>` + "`+\"`\"+`" + `+msgA+` + "`+\"`\"+`" + `</span>
					  </div>
					  <div style="margin-top: 10px; margin-bottom: 10px;">
						  <input id="mainTextID" type="text" />
					  </div>
					  <div style="margin-top: 10px; margin-bottom: 10px;">
						  <input id="submitButtonID" type="button" value="Ok" />
						  <input id="cancelButtonID" type="button" value="Cancel" />
					  </div>
				  </center>
				  <script type="text/tiscript">
					  $(#submitButtonID).onClick = function() {
						  view.close($(#mainTextID).value);
					  };
  
					  $(#cancelButtonID).onClick = function() {
						  view.close();
					  };
				  </scr` + "`+\"`\"+`" + `+` + "`+\"`\"+`" + `ipt>
				</body>
				</html>
				` + "`+\"`\"+`" + `
			  });
  
			  return res;
		  }

		event click $(#btnEncrypt)
		{
		  	var res = getInput("Secure Code");

			if (res == undefined) {
				return;
			}

			var sourceT = $(plaintext).value;

			var encStrT = view.encryptText(sourceT, res);
		
			if (isErrStr(encStrT)) {
				showError("Error", String.printf("failed to encrypt content: %v",getErrStr(encStrT)));
				return;
			}
		
			$(plaintext).value = "\/\/TXDEF#" + encStrT;
			editFileCleanFlagG = "*";
			updateFileName();
		}
	
		event click $(#btnDecrypt)
		{
		  	var res = getInput("Secure Code");

			if (res == undefined) {
				return;
			}

			var sourceT = $(plaintext).value;

			var encStrT = view.decryptText(sourceT, res);
		
			if (isErrStr(encStrT)) {
				showError("Error", String.printf("failed to decrypt content: %v",getErrStr(encStrT)));
				return;
			}
		
			$(plaintext).value = encStrT;
			editFileCleanFlagG = "*";
			updateFileName();
		}
	
		event click $(#btnRun)
		{
			var res = getInput("Arguments to pass to script")

			if (res == undefined) {
				return;
			}

			var rs = view.runScript($(plaintext).value, res);

			showInfo("Result", rs)
		 
		  	// view.prints(String.printf("result = %v", rs));
		}
	

		function editCloseClick() {
			view.close();
			// view.exit();
		}

		function editFile(fileNameA) {
			var fcT string;

			//view.prints("fileNameA: "+fileNameA);

			if (fileNameA == "") {
				editFileNameG = "";

				fcT = "";

				editFileCleanFlagG = "*";
			} else {
				editFileNameG = fileNameA;

				fcT = view.loadStringFromFile(fileNameA);

//		if tk.IsErrorString(fcT) {
//			tk.Pl("failed to load file %v: %v", editFileNameG, tk.GetErrorString(fcT))
//			return

//		}

				editFileCleanFlagG = "";
			}

			//view.prints(fcT);

			$(plaintext).attributes["type"] = "text/script";

			$(plaintext).value = fcT;

			updateFileName();

		}

		function self.ready() {

			//$(plaintext).value = "<html>\n<body>\n<span>abc</span>\n</body></html>";

			$(#btnLoad).onClick = editFileLoadClick;
			$(#btnSave).onClick = editFileSaveClick;
			$(#btnSaveAs).onClick = editFileSaveAsClick;
			// $(#btnEncrypt).onClick = editFEncryptClick;
			// $(#btnDecrypt).onClick = editDecryptClick;
			// $(#btnRun).onClick = editRunClick;
			$(#btnClose).onClick = editCloseClick;

			$(plaintext#source).onControlEvent = function(evt) {
				switch (evt.type) {
					case Event.EDIT_VALUE_CHANGED:      
						editFileCleanFlagG = "*";
						updateFileName();
						return true;
				}
			};

		}
	</script>

</head>
<body>
	<div style="margin-top: 10px; margin-bottom: 10px;"><span id="fileNameLabelID"></span></div>
	<div style="margin-top: 10px; margin-bottom: 10px;">
		<button id="btn1" style="display: none">Load...</button>
		<button id="btnLoad">Load</button>
		<button id="btnSave">Save</button>
		<button id="btnSaveAs">SaveAs</button>
		<button id="btnEncrypt">Encrypt</button>
		<button id="btnDecrypt">Decrypt</button>
		<button id="btnRun">Run</button>
		<button id="btnClose">Close</button>
	</div>
	<plaintext#source type="text/html" style="font-size: 1.2em;"></plaintext>

</body>
</html>
` + "`" + `

// htmlT = tk.LoadStringFromFile(tk.JoinPath(path_filepath.Dir(scriptPathG), "editFileSciter.st"))

// tk.CheckErrorString(htmlT)

runtime.LockOSThread()

w, err := window.New(sciter.DefaultWindowCreateFlag, sciter.DefaultRect)

checkError(err)

w.SetOption(sciter.SCITER_SET_SCRIPT_RUNTIME_FEATURES, sciter.ALLOW_FILE_IO | sciter.ALLOW_SOCKET_IO | sciter.ALLOW_EVAL | sciter.ALLOW_SYSINFO)

w.LoadHtml(htmlT, "")

w.SetTitle("Gox Editor")

w.DefineFunction("prints", func(args) {
	tk.Pl("%v", args[0].String())
	return sciter.NewValue("")
})

w.DefineFunction("loadStringFromFile", func(args) {
	rs := tk.LoadStringFromFile(args[0].String())
	return sciter.NewValue(rs)
})

w.DefineFunction("saveStringToFile", func(args) {
	rs := tk.SaveStringToFile(args[0].String(), args[1].String())
	return sciter.NewValue(rs)
})

w.DefineFunction("encryptText", func(args) {
	rs := tk.EncryptStringByTXDEF(args[0].String(), args[1].String())
	return sciter.NewValue(rs)
})

w.DefineFunction("decryptText", func(args) {
	rs := tk.DecryptStringByTXDEF(args[0].String(), args[1].String())
	return sciter.NewValue(rs)
})

w.DefineFunction("runScript", func(args) {
	rs := runScript(args[0].String(), "", args[1].String())
	return sciter.NewValue(tk.Spr("%v", rs))
})

w.DefineFunction("exit", func(args) {
	os.Exit(1);
})

data, _ := w.Call("getScreenWH") //, sciter.NewValue(10), sciter.NewValue(20))
// fmt.Println("data:", data.String())

fileNameT := tk.GetParameterByIndexWithDefaultValue(argsG, 0, "")

w.Call("editFile", sciter.NewValue(fileNameT))

w.Show()

// screenshot = github_kbinani_screenshot

// tk.Plvsr(screenshot.NumActiveDisplays(), screenshot.GetDisplayBounds(0))

// bounds := screenshot.GetDisplayBounds(0)

// img, err := screenshot.CaptureRect(bounds)
// if err != nil {
// 	panic(err)
// }
// fileName := fmt.Sprintf("%d_%dx%d.png", i, bounds.Dx(), bounds.Dy())
// file, _ := os.Create(fileName)

// image_png.Encode(file, img)

// file.Close()

// fmt.Printf("#%d : %v \"%s\"\n", i, bounds, fileName)

w.Run()

`

func editFile(fileNameA string, argsA ...string) {
	rs := runScriptX(editFileScriptG, argsA...)

	if rs != notFoundG {
		// tk.Pl("%v", rs)
	}

}

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
			if valueT != spec.Undefined && valueT != notFoundG {
				tk.Pl("Error: %T %v", valueT, valueT)
			}
		} else {
			tk.Pl("%v", rs)
		}
	}

}

func test() {

}
