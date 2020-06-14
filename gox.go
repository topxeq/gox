package main

import (
	"bufio"

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

	// GUI related start

	execq "github.com/topxeq/qlang/exec"
	specq "github.com/topxeq/qlang/spec"

	// GUI related end

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

	// GUI related start

	qlgithub_AllenDang_giu "github.com/topxeq/qlang/lib/github.com/AllenDang/giu"
	qlgithub_AllenDang_giu_imgui "github.com/topxeq/qlang/lib/github.com/AllenDang/giu/imgui"

	// GUI related end

	qlgithub_dgraphio_badger "github.com/topxeq/qlang/lib/github.com/dgraph-io/badger"

	qlgithub_fogleman_gg "github.com/topxeq/qlang/lib/github.com/fogleman/gg"

	qlgithub_360EntSecGroupSkylar_excelize "github.com/topxeq/qlang/lib/github.com/360EntSecGroup-Skylar/excelize"

	qlgonumorg_v1_plot "github.com/topxeq/qlang/lib/gonum.org/v1/plot"
	qlgonumorg_v1_plot_plotter "github.com/topxeq/qlang/lib/gonum.org/v1/plot/plotter"
	qlgonumorg_v1_plot_plotutil "github.com/topxeq/qlang/lib/gonum.org/v1/plot/plotutil"
	qlgonumorg_v1_plot_vg "github.com/topxeq/qlang/lib/gonum.org/v1/plot/vg"

	qlgithub_domodwyer_mailyak "github.com/topxeq/qlang/lib/github.com/domodwyer/mailyak"

	// GUI related start

	qlgithub_scitersdk_gosciter "github.com/topxeq/qlang/lib/github.com/sciter-sdk/go-sciter"
	qlgithub_scitersdk_gosciter_window "github.com/topxeq/qlang/lib/github.com/sciter-sdk/go-sciter/window"

	"github.com/sciter-sdk/go-sciter"

	// GUI related end

	// full version related start
	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/godror/godror"

	// "image"
	// "image/color"
	// "image/draw"
	// "image/png"

	// "gonum.org/v1/plot"
	// "gonum.org/v1/plot/plotter"
	// "gonum.org/v1/plot/vg"

	// full version related end

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"

	// GUI related start
	// full version related start
	"github.com/sqweek/dialog"
	// full version related end
	// GUI related end

	"github.com/topxeq/tk"

	// GUI related start
	// full version related start
	"github.com/AllenDang/giu"
	"github.com/AllenDang/giu/imgui"
	// full version related end
	// GUI related end
)

// Non GUI related

var versionG = "0.998a"

var verboseG = false

var variableG = make(map[string]interface{})

var qlVMG *qlang.Qlang = nil

var varMutexG sync.Mutex

func exit() {
	defer func() {
		if r := recover(); r != nil {
			tk.Printfln("发生异常，错误信息：%v", r)

			return
		}
	}()

	os.Exit(1)
}

func qlEval(strA string) string {
	vmT := qlang.New()

	retG = notFoundG

	errT := vmT.SafeEval(strA)

	if errT != nil {
		return errT.Error()
	}

	if retG != notFoundG {
		return tk.Spr("%v", retG)
	}

	rs, ok := vmT.GetVar("outG")

	if !ok {
		return ""
	}

	return tk.Spr("%v", rs)
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

		if argsA != nil && len(argsA) > 0 {
			vmT.SetVar("argsG", argsA)
		}

		retG = notFoundG

		errT := vmT.SafeEval(codeA)

		if errT != nil {
			return errT.Error()
		}

		return retG
	} else {
		return tk.SystemCmd("gox", append([]string{codeA}, argsA...)...)
	}

}

// full version related start
// func newRGBA(r, g, b, a uint8) color.RGBA {
// 	return color.RGBA{r, g, b, a}
// }

// func newNRGBAFromHex(strA string) color.NRGBA {
// 	r, g, b, a := tk.ParseHexColor(strA)

// 	return color.NRGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
// }

// func newRGBAFromHex(strA string) color.RGBA {
// 	r, g, b, a := tk.ParseHexColor(strA)

// 	return color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
// }

// func newPlotXY(xA, yA float64) plotter.XY {
// 	return plotter.XY{X: xA, Y: yA}
// }

// func loadRGBAFromImage(imageA image.Image) (*image.RGBA, error) {
// 	switch imageT := imageA.(type) {
// 	case *image.RGBA:
// 		return imageT, nil
// 	default:
// 		rgba := image.NewRGBA(imageT.Bounds())
// 		draw.Draw(rgba, imageT.Bounds(), imageT, image.Pt(0, 0), draw.Src)
// 		return rgba, nil
// 	}

// }

// func LoadPlotImage(p *plot.Plot, w vg.Length, h vg.Length) (*image.RGBA, error) {

// 	var bufT bytes.Buffer

// 	writerT, errT := p.WriterTo(w, h, "png")

// 	if errT != nil {
// 		return nil, errT
// 	}

// 	_, errT = writerT.WriteTo(&bufT)

// 	if errT != nil {
// 		return nil, errT
// 	}

// 	readerT := bytes.NewReader(bufT.Bytes())

// 	// defer readerT.Close()

// 	// imgFile, err := os.Open(imgPath)
// 	// if err != nil {
// 	// 	return nil, err
// 	// }
// 	// defer imgFile.Close()

// 	img, err := png.Decode(readerT)
// 	if err != nil {
// 		return nil, err
// 	}

// 	switch trueImg := img.(type) {
// 	case *image.RGBA:
// 		return trueImg, nil
// 	default:
// 		rgba := image.NewRGBA(trueImg.Bounds())
// 		draw.Draw(rgba, trueImg.Bounds(), trueImg, image.Pt(0, 0), draw.Src)
// 		return rgba, nil
// 	}
// }

// type MyXYer plotter.XY

// func (v MyXYer) Len() int {
// 	return 1
// }

// func (v MyXYer) XY(int) (x, y float64) {
// 	return 1
// }

// full version related end

// func setValue(p interface{}, v interface{}) {
// 	// tk.Pl("%#v", reflect.TypeOf(p).Kind())
// 	// p = v

// 	srcRef := reflect.ValueOf(v)
// 	vp := reflect.ValueOf(p)
// 	vp.Elem().Set(srcRef)
// }

// func getValue(p interface{}) interface{} {
// 	vp := reflect.Indirect(reflect.ValueOf(p))
// 	return vp.Interface()
// }

// func bitXor(p interface{}, v interface{}) interface{} {
// 	switch p.(type) {
// 	case int:
// 		return p.(int) ^ v.(int)
// 	case int64:
// 		return p.(int64) ^ v.(int64)
// 	case int32:
// 		return p.(int32) ^ v.(int32)
// 	case int16:
// 		return p.(int16) ^ v.(int16)
// 	case int8:
// 		return p.(int8) ^ v.(int8)
// 	case uint64:
// 		return p.(uint64) ^ v.(uint64)
// 	case uint32:
// 		return p.(uint32) ^ v.(uint32)
// 	case uint16:
// 		return p.(uint16) ^ v.(uint16)
// 	case uint8:
// 		return p.(uint8) ^ v.(uint8)
// 	case uint:
// 		return p.(uint) ^ v.(uint)
// 	}

// 	return 0
// }

// GUI related start

// func NewFuncIntStringError(funcA *interface{}) *(func(int) (string, error)) {
// 	funcT := (*funcA).(*execq.Function)
// 	f := func(n int) (string, error) {
// 		return funcT.Call(execq.NewStack(), n).(...interface{})
// 	}

// 	return &f
// }

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

// GUI related end

func importQLNonGUIPackages() {
	printValue := func(nameA string) {

		v, ok := qlVMG.GetVar(nameA)

		if !ok {
			tk.Pl("no variable by the name found: %v", nameA)
			return
		}

		tk.Pl("%v(%T): %v", nameA, v, v)

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
		"eval":             qlEval,
		"printfln":         tk.Pl,
		"fprintf":          fmt.Fprintf,
		"pl":               tk.Pl,
		"pln":              fmt.Println,
		"plv":              tk.Plv,
		"plvsr":            tk.Plvsr,
		"plerr":            tk.PlErr,
		"pv":               printValue,
		"exit":             exit,
		"setValue":         tk.SetValue,
		"getValue":         tk.GetValue,
		"bitXor":           tk.BitXor,
		"setVar":           tk.SetVar,
		"getVar":           tk.GetVar,
		"checkError":       tk.CheckError,
		"checkErrorString": tk.CheckErrorString,
		"getInput":         tk.GetUserInput,
		"getInputf":        tk.GetInputf,
		"run":              runFile,
		"typeOf":           tk.TypeOfValueReflect,
		"remove":           tk.RemoveItemsInArray,
		"runScript":        runScript,
		"getClipText":      tk.GetClipText,
		"setClipText":      tk.SetClipText,
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

	// GUI related start

	qlang.Import("github_AllenDang_giu", qlgithub_AllenDang_giu.Exports)
	qlang.Import("giu", qlgithub_AllenDang_giu.Exports)
	qlang.Import("github_AllenDang_giu_imgui", qlgithub_AllenDang_giu_imgui.Exports)
	qlang.Import("giu_imgui", qlgithub_AllenDang_giu_imgui.Exports)

	qlgithub_scitersdk_gosciter.Exports["NewValue"] = sciter.NewValue
	qlgithub_scitersdk_gosciter.Exports["NullValue"] = sciter.NullValue

	qlgithub_scitersdk_gosciter.Exports["NewScnLoadDataFunc"] = NewScnLoadDataFunc
	qlgithub_scitersdk_gosciter.Exports["NewScnDataLoaded"] = NewScnDataLoaded

	qlang.Import("github_scitersdk_gosciter", qlgithub_scitersdk_gosciter.Exports)
	qlang.Import("github_scitersdk_gosciter_window", qlgithub_scitersdk_gosciter_window.Exports)

	InitBlink()
	// GUI related end

	qlang.Import("github_dgraphio_badger", qlgithub_dgraphio_badger.Exports)
	qlang.Import("badger", qlgithub_dgraphio_badger.Exports)

	qlang.Import("github_fogleman_gg", qlgithub_fogleman_gg.Exports)
	qlang.Import("gg", qlgithub_fogleman_gg.Exports)

	qlang.Import("gonumorg_v1_plot", qlgonumorg_v1_plot.Exports)
	qlang.Import("plot", qlgonumorg_v1_plot.Exports)
	qlang.Import("gonumorg_v1_plot_plotter", qlgonumorg_v1_plot_plotter.Exports)
	qlang.Import("plot_plotter", qlgonumorg_v1_plot_plotter.Exports)
	qlang.Import("gonumorg_v1_plot_plotutil", qlgonumorg_v1_plot_plotutil.Exports)
	qlang.Import("plot_plotutil", qlgonumorg_v1_plot_plotutil.Exports)
	qlang.Import("gonumorg_v1_plot_vg", qlgonumorg_v1_plot_vg.Exports)
	qlang.Import("plot_vg", qlgonumorg_v1_plot_vg.Exports)

	qlang.Import("github_360EntSecGroupSkylar_excelize", qlgithub_360EntSecGroupSkylar_excelize.Exports)

	qlang.Import("github_domodwyer_mailyak", qlgithub_domodwyer_mailyak.Exports)

}

func showHelp() {
	tk.Pl("Gox by TopXeQ V%v\n", versionG)

	tk.Pl("Usage: gox [-v|-h] test.gox, ...\n")
	tk.Pl("or just gox without arguments to start REPL instead.\n")

}

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

func loadFont() {
	fonts := giu.Context.IO().Fonts()

	rangeVarT := tk.GetVar("FontRange")

	ranges := imgui.NewGlyphRanges()

	builder := imgui.NewFontGlyphRangesBuilder()

	if rangeVarT == nil {
		builder.AddRanges(fonts.GlyphRangesDefault())
	} else {
		rangeStrT := rangeVarT.(string)
		if rangeStrT == "" || tk.StartsWith(rangeStrT, "COMMON") {
			builder.AddRanges(fonts.GlyphRangesChineseSimplifiedCommon())
			builder.AddText("辑" + rangeStrT[6:])
		} else if rangeStrT == "FULL" {
			builder.AddRanges(fonts.GlyphRangesChineseFull())
		} else {
			builder.AddText(rangeStrT)
		}
	}

	builder.BuildRanges(ranges)

	fontPath := "c:/Windows/Fonts/simhei.ttf"

	if tk.Contains(tk.GetOSName(), "rwin") {
		fontPath = "/Library/Fonts/Microsoft/SimHei.ttf"
	}

	fontVarT := tk.GetVar("Font") // "c:/Windows/Fonts/simsun.ttc"

	if fontVarT != nil {
		fontPath = fontVarT.(string)
	}

	fontSizeStrT := "16"

	fontSizeVarT := tk.GetVar("FontSize")

	if fontSizeVarT != nil {
		fontSizeStrT = fontSizeVarT.(string)
	}

	fontSizeT := tk.StrToIntWithDefaultValue(fontSizeStrT, 16)

	// fonts.AddFontFromFileTTF(fontPath, 14)
	fonts.AddFontFromFileTTFV(fontPath, float32(fontSizeT), imgui.DefaultFontConfig, ranges.Data())
}

// full version related end

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

	var guiExports = map[string]interface{}{
		"NewMasterWindow":         giu.NewMasterWindow,
		"SingleWindow":            giu.SingleWindow,
		"Window":                  giu.Window,
		"SingleWindowWithMenuBar": giu.SingleWindowWithMenuBar,
		"WindowV":                 giu.WindowV,

		"MasterWindowFlagsNotResizable": giu.MasterWindowFlagsNotResizable,
		"MasterWindowFlagsMaximized":    giu.MasterWindowFlagsMaximized,
		"MasterWindowFlagsFloating":     giu.MasterWindowFlagsFloating,

		// "Layout":          giu.Layout,

		"NewTextureFromRgba": giu.NewTextureFromRgba,

		"Label":                  giu.Label,
		"Line":                   giu.Line,
		"Button":                 giu.Button,
		"InvisibleButton":        giu.InvisibleButton,
		"ImageButton":            giu.ImageButton,
		"InputTextMultiline":     giu.InputTextMultiline,
		"Checkbox":               giu.Checkbox,
		"RadioButton":            giu.RadioButton,
		"Child":                  giu.Child,
		"ComboCustom":            giu.ComboCustom,
		"Combo":                  giu.Combo,
		"ContextMenu":            giu.ContextMenu,
		"Group":                  giu.Group,
		"Image":                  giu.Image,
		"ImageWithFile":          giu.ImageWithFile,
		"ImageWithUrl":           giu.ImageWithUrl,
		"InputText":              giu.InputText,
		"InputTextV":             giu.InputTextV,
		"InputTextFlagsPassword": giu.InputTextFlagsPassword,
		"InputInt":               giu.InputInt,
		"InputFloat":             giu.InputFloat,
		"MainMenuBar":            giu.MainMenuBar,
		"MenuBar":                giu.MenuBar,
		"MenuItem":               giu.MenuItem,
		"PopupModal":             giu.PopupModal,
		"OpenPopup":              giu.OpenPopup,
		"CloseCurrentPopup":      giu.CloseCurrentPopup,
		"ProgressBar":            giu.ProgressBar,
		"Separator":              giu.Separator,
		"SliderInt":              giu.SliderInt,
		"SliderFloat":            giu.SliderFloat,
		"HSplitter":              giu.HSplitter,
		"VSplitter":              giu.VSplitter,
		"TabItem":                giu.TabItem,
		"TabBar":                 giu.TabBar,
		"Row":                    giu.Row,
		"Table":                  giu.Table,
		"FastTable":              giu.FastTable,
		"Tooltip":                giu.Tooltip,
		"TreeNode":               giu.TreeNode,
		"Spacing":                giu.Spacing,
		"Custom":                 giu.Custom,
		"Condition":              giu.Condition,
		"ListBox":                giu.ListBox,
		"DatePicker":             giu.DatePicker,
		"Dummy":                  giu.Dummy,
		// "Widget":             giu.Widget,

		"PrepareMessageBox": giu.PrepareMsgbox,
		"MessageBox":        giu.Msgbox,

		"LoadFont": loadFont,

		"GetConfirm": getConfirmGUI,

		"SimpleInfo":      simpleInfo,
		"SimpleError":     simpleError,
		"SelectFile":      selectFileGUI,
		"SelectSaveFile":  selectFileToSaveGUI,
		"SelectDirectory": selectDirectoryGUI,

		"EditFile":   editFile,
		"LoopWindow": loopWindow,

		"LayoutP": giu.Layout{},

		"Layout": specq.StructOf((*giu.Layout)(nil)),
		"Widget": specq.StructOf((*giu.Widget)(nil)),
	}

	qlang.Import("gui", guiExports)
	// full version related end

	InitLCLFirst()
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

var (
	editorG            imgui.TextEditor
	errMarkersG        imgui.ErrorMarkers
	editFileNameG      string
	editFileCleanFlagG string
	editSecureCodeG    string
	editArgsG          string
)

func editorLoad() {
	if editorG.IsTextChanged() {
		editFileCleanFlagG = "*"
	}

	if editFileCleanFlagG != "" {
		rs := getConfirmGUI("Please confirm", "File modified, load another file anyway?")

		if rs == false {
			return
		}
	}

	fileNameNewT := selectFileGUI("Select the file to open...", "All files", "*")

	if tk.IsErrorString(fileNameNewT) {
		if tk.EndsWith(fileNameNewT, "Cancelled") {
			giu.Msgbox("Info", tk.Spr("Action cancelled by user"))
			return
		}

		giu.Msgbox("Error", tk.Spr("Failed to select file: %v", tk.GetErrorString(fileNameNewT)))
		return
	}

	fcT := tk.LoadStringFromFile(fileNameNewT)

	if tk.IsErrorString(fcT) {
		giu.Msgbox("Error", tk.Spr("Failed to load file content: %v", tk.GetErrorString(fileNameNewT)))
		return
	}

	editFileNameG = fileNameNewT
	editorG.SetText(fcT)
	editFileCleanFlagG = ""

}

func editorSaveAs() {
	fileNameNewT := selectFileToSaveGUI("Select the file to save...", "All file", "*")

	if tk.IsErrorString(fileNameNewT) {
		if tk.EndsWith(fileNameNewT, "Cancelled") {
			giu.Msgbox("Info", tk.Spr("Action cancelled by user"))
			return
		}

		giu.Msgbox("Error", tk.Spr("Failed to select file: %v", tk.GetErrorString(fileNameNewT)))
		return
	}

	editFileNameG = fileNameNewT

	rs := true
	// if tk.IfFileExists(editFileNameG) {
	// 	rs = getConfirmGUI("请再次确认", "文件已存在，是否覆盖?")
	// }

	if rs == true {
		rs1 := tk.SaveStringToFile(editorG.GetText(), editFileNameG)

		if rs1 != "" {
			giu.Msgbox("Error", tk.Spr("Failed to save: %v", rs))
			return
		}

		giu.Msgbox("Info", tk.Spr("File saved to: %v", editFileNameG))

		editFileCleanFlagG = ""
	}

}

func editorSave() {
	if editFileNameG == "" {
		editorSaveAs()

		return
	}

	rs := false

	if tk.IfFileExists(editFileNameG) {
		rs = getConfirmGUI("Please confirm", "The file already exists, confirm to overwrite?")
	}

	if rs == true {
		rs1 := tk.SaveStringToFile(editorG.GetText(), editFileNameG)

		if rs1 != "" {
			giu.Msgbox("Error", tk.Spr("Failed to save: %v", rs))
			return
		}

		giu.Msgbox("Info", tk.Spr("File saved to file: %v", editFileNameG))

		editFileCleanFlagG = ""
	}

}

func editEncrypt() {
	imgui.CloseCurrentPopup()

	sourceT := editorG.GetText()

	encStrT := tk.EncryptStringByTXDEF(sourceT, editSecureCodeG)

	if tk.IsErrorString(encStrT) {
		simpleError("Error", tk.Spr("failed to encrypt content: %v", tk.GetErrorString(encStrT)))
		return
	}

	editorG.SetText("//TXDEF#" + encStrT)
	editFileCleanFlagG = "*"

	editSecureCodeG = ""
}

func editEncryptClick() {
	giu.OpenPopup("Please enter:##EncryptInputSecureCode")
}

func editDecrypt() {
	imgui.CloseCurrentPopup()

	sourceT := tk.Trim(editorG.GetText())

	encStrT := tk.DecryptStringByTXDEF(sourceT, editSecureCodeG)

	if tk.IsErrorString(encStrT) {
		simpleError("Error", tk.Spr("failed to decrypt content: %v", tk.GetErrorString(encStrT)))
		return
	}

	editorG.SetText(encStrT)
	editFileCleanFlagG = "*"
	editSecureCodeG = ""

}

func editDecryptClick() {
	giu.OpenPopup("Please enter:##DecryptInputSecureCode")
}

func editRun() {
	imgui.CloseCurrentPopup()

	runScript(editorG.GetText(), "", editArgsG)
}

func editRunClick() {
	giu.OpenPopup("Please enter:##RunInputArgs")
}

func onButtonCloseClick() {
	exit()
}

func loopWindow(windowA *giu.MasterWindow, loopA func()) {
	// wnd := giu.NewMasterWindow("Gox Editor", 800, 600, 0, loadFont)

	windowA.Main(loopA)

}

func editorLoop() {
	giu.SingleWindow("Gox Editor", giu.Layout{
		giu.Label(editFileNameG + editFileCleanFlagG),
		giu.Dummy(30, 0),
		giu.Line(
			giu.Button("Load", editorLoad),
			giu.Button("Save", editorSave),
			giu.Button("Save As...", editorSaveAs),
			giu.Button("Check", func() {

				// sourceT := editorG.GetText()

				// parser.EnableErrorVerbose()
				// _, errT := parser.ParseSrc(sourceT)
				// // tk.Plv(stmts)

				// e, ok := errT.(*parser.Error)

				// if ok {
				// 	errMarkersG.Clear()
				// 	errMarkersG.Insert(e.Pos.Line, tk.Spr("[col: %v, size: %v] %v", e.Pos.Column, errMarkersG.Size(), e.Error()))

				// 	editorG.SetErrorMarkers(errMarkersG)

				// } else if errT != nil {
				// 	giu.Msgbox("Error", tk.Spr("%#v", errT))
				// } else {
				// 	giu.Msgbox("Info", "Syntax check passed.")
				// }

			}),
			giu.Button("Encrypt", editEncryptClick),
			giu.Button("Decrypt", editDecryptClick),
			giu.Button("Run", editRunClick),
			giu.Button("Close", onButtonCloseClick),
			// giu.Button("Get Text", func() {
			// 	if editorG.HasSelection() {
			// 		fmt.Println(editorG.GetSelectedText())
			// 	} else {
			// 		fmt.Println(editorG.GetText())
			// 	}

			// 	column, line := editorG.GetCursorPos()
			// 	fmt.Println("Cursor pos:", column, line)

			// 	column, line = editorG.GetSelectionStart()
			// 	fmt.Println("Selection start:", column, line)

			// 	fmt.Println("Current line is", editorG.GetCurrentLineText())
			// }),
			// giu.Button("Set Text", func() {
			// 	editorG.SetText("Set text")
			// 	editFileNameG = "Set text"
			// }),
			// giu.Button("Set Error Marker", func() {
			// 	errMarkersG.Clear()
			// 	errMarkersG.Insert(1, "Error message")
			// 	fmt.Println("ErrMarkers Size:", errMarkersG.Size())

			// 	editorG.SetErrorMarkers(errMarkersG)
			// }),
		),
		giu.PopupModal("Please enter:##EncryptInputSecureCode", giu.Layout{
			giu.Line(
				giu.Label("Secure code"),
				giu.InputTextV("", 40, &editSecureCodeG, giu.InputTextFlagsPassword, nil, nil),
			),
			giu.Line(
				giu.Button("Ok", editEncrypt),
				giu.Button("Cancel", func() { imgui.CloseCurrentPopup() }),
			),
		}),
		giu.PopupModal("Please enter:##DecryptInputSecureCode", giu.Layout{
			giu.Line(
				giu.Label("Secure code"),
				giu.InputTextV("", 40, &editSecureCodeG, giu.InputTextFlagsPassword, nil, nil),
			),
			giu.Line(
				giu.Button("Ok", editDecrypt),
				giu.Button("Cancel", func() { imgui.CloseCurrentPopup() }),
			),
		}),
		giu.PopupModal("Please enter:##RunInputArgs", giu.Layout{
			giu.Line(
				giu.Label("Arguments to pass to VM"),
				giu.InputText("", 80, &editArgsG),
			),
			giu.Line(
				giu.Button("Ok", editRun),
				giu.Button("Cancel", func() { imgui.CloseCurrentPopup() }),
			),
		}),
		giu.Custom(func() {
			editorG.Render("Hello", imgui.Vec2{X: 0, Y: 0}, true)
			if giu.IsItemHovered() {
				if editorG.IsTextChanged() {
					editFileCleanFlagG = "*"
				}
			}
		}),
		giu.PrepareMsgbox(),
	})
}

func editFile(fileNameA string) {
	var fcT string

	if fileNameA == "" {
		editFileNameG = ""

		fcT = ""

		editFileCleanFlagG = "*"
	} else {
		editFileNameG = fileNameA

		fcT = tk.LoadStringFromFile(editFileNameG)

		if tk.IsErrorString(fcT) {
			tk.Pl("failed to load file %v: %v", editFileNameG, tk.GetErrorString(fcT))
			return

		}

		editFileCleanFlagG = ""

	}

	errMarkersG = imgui.NewErrorMarkers()

	editorG = imgui.NewTextEditor()

	editorG.SetShowWhitespaces(false)
	editorG.SetTabSize(2)
	editorG.SetText(fcT)
	editorG.SetLanguageDefinitionC()

	// setVar("Font", "c:/Windows/Fonts/simsun.ttc")
	tk.SetVar("FontRange", "COMMON")
	tk.SetVar("FontSize", "15")

	wnd := giu.NewMasterWindow("Gox Editor", 800, 600, 0, loadFont)
	// tk.Pl("%T", wnd)
	wnd.Main(editorLoop)

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

	argsT := os.Args

	if tk.IfSwitchExistsWhole(argsT, "-version") {
		tk.Pl("Gox by TopXeQ V%v", versionG)
		return
	}

	if tk.IfSwitchExistsWhole(argsT, "-h") {
		showHelp()
		return
	}

	scriptT := tk.GetParameterByIndexWithDefaultValue(argsT, 1, "")

	// GUI related start

	// full version related start
	if tk.IfSwitchExistsWhole(argsT, "-edit") {
		editFile(scriptT)

		return
	}
	// full version related end

	// GUI related end

	if scriptT == "" {

		autoPathT := filepath.Join(tk.GetApplicationPath(), "auto.gox")

		if tk.IfFileExists(autoPathT) {
			scriptT = autoPathT
		} else {
			initQLVM()

			runInteractiveQlang()

			// tk.Pl("not enough parameters")

			return
		}

	}

	encryptCodeT := tk.GetSwitchWithDefaultValue(argsT, "-encrypt=", "")

	if encryptCodeT != "" {
		fcT := tk.LoadStringFromFile(scriptT)

		if tk.IsErrorString(fcT) {
			tk.Pl("failed to load file [%v]: %v", scriptT, tk.GetErrorString(fcT))
			return
		}

		encStrT := tk.EncryptStringByTXDEF(fcT, encryptCodeT)

		if tk.IsErrorString(encStrT) {
			tk.Pl("failed to encrypt content [%v]: %v", scriptT, tk.GetErrorString(encStrT))
			return
		}

		rsT := tk.SaveStringToFile("//TXDEF#"+encStrT, scriptT+"e")

		if tk.IsErrorString(rsT) {
			tk.Pl("failed to encrypt file [%v]: %v", scriptT, tk.GetErrorString(rsT))
			return
		}

		return
	}

	decryptCodeT := tk.GetSwitchWithDefaultValue(argsT, "-decrypt=", "")

	if decryptCodeT != "" {
		fcT := tk.LoadStringFromFile(scriptT)

		if tk.IsErrorString(fcT) {
			tk.Pl("failed to load file [%v]: %v", scriptT, tk.GetErrorString(fcT))
			return
		}

		decStrT := tk.DecryptStringByTXDEF(fcT, decryptCodeT)

		if tk.IsErrorString(decStrT) {
			tk.Pl("failed to decrypt content [%v]: %v", scriptT, tk.GetErrorString(decStrT))
			return
		}

		rsT := tk.SaveStringToFile(decStrT, scriptT+"d")

		if tk.IsErrorString(rsT) {
			tk.Pl("failed to decrypt file [%v]: %v", scriptT, tk.GetErrorString(rsT))
			return
		}

		return
	}

	decryptRunCodeT := tk.GetSwitchWithDefaultValue(argsT, "-decrun=", "")

	ifExampleT := tk.IfSwitchExistsWhole(argsT, "-example")
	ifGoPathT := tk.IfSwitchExistsWhole(argsT, "-gopath")
	ifLocalT := tk.IfSwitchExistsWhole(argsT, "-local")
	ifRemoteT := tk.IfSwitchExistsWhole(argsT, "-remote")
	ifCloudT := tk.IfSwitchExistsWhole(argsT, "-cloud")
	sshT := tk.GetSwitchWithDefaultValue(argsT, "-ssh=", "")
	ifViewT := tk.IfSwitchExistsWhole(argsT, "-view")

	verboseG = tk.IfSwitchExistsWhole(argsT, "-verbose")

	var fcT string

	if ifExampleT {
		if (!tk.EndsWith(scriptT, ".gox")) && (!tk.EndsWith(scriptT, ".ql")) {
			scriptT += ".gox"
		}
		fcT = tk.DownloadPageUTF8("https://gitee.com/topxeq/gox/raw/master/scripts/"+scriptT, nil, "", 30)
	} else if ifRemoteT {
		fcT = tk.DownloadPageUTF8(scriptT, nil, "", 30)
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
	} else if sshT != "" {
		if (!tk.EndsWith(scriptT, ".gox")) && (!tk.EndsWith(scriptT, ".ql")) {
			scriptT += ".gox"
		}

		fcT = downloadStringFromSSH(sshT, scriptT)

		if tk.IsErrorString(fcT) {
			tk.Pl("failed to get script from SSH: %v", tk.GetErrorString(fcT))
			return

		}
	} else if ifGoPathT {
		if (!tk.EndsWith(scriptT, ".gox")) && (!tk.EndsWith(scriptT, ".ql")) {
			scriptT += ".gox"
		}

		fcT = tk.LoadStringFromFile(filepath.Join(tk.GetEnv("GOPATH"), "src", "github.com", "topxeq", "gox", "scripts", scriptT))
	} else if ifLocalT {
		if (!tk.EndsWith(scriptT, ".gox")) && (!tk.EndsWith(scriptT, ".ql")) {
			scriptT += ".gox"
		}

		localPathT := getCfgString("localScriptPath.cfg")

		if tk.IsErrorString(localPathT) {
			tk.Pl("failed to get local path: %v", tk.GetErrorString(localPathT))

			return
		}

		// if tk.GetEnv("GOXVERBOSE") == "true" {
		// 	tk.Pl("Try to load script from %v", filepath.Join(localPathT, scriptT))
		// }

		fcT = tk.LoadStringFromFile(filepath.Join(localPathT, scriptT))
	} else {
		fcT = tk.LoadStringFromFile(scriptT)
	}

	if tk.IsErrorString(fcT) {
		tk.Pl("failed to load script from %v: %v", scriptT, tk.GetErrorString(fcT))

		return
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

		return
	}

	initQLVM()

	errT := qlVMG.SafeEval(fcT)
	if errT != nil {

		tk.Pl("failed to execute script(%v) error: %v\n", scriptT, errT)

		// f, l := qlVMG.Code.Line(qlVMG.Code.Reserve().Next())
		// tk.Pl("Next line: %v, %v", f, l)

		return
	}

	rs, ok := qlVMG.GetVar("outG")

	if ok {
		tk.Pl("%#v", rs)
	}

}

func test() {
	// var v *vcl.TKeyEvent

	// tk.Pl("%#v, %T", v, v)

	// f := func(sender vcl.IObject, key *types.Char, shift types.TShiftState) {
	// 	// funcA.Call(execq.NewStack(), sender, key, shift)
	// }

	// v = &f

	// tk.Pl("%#v, %T", v, v)

	// return &f

}
