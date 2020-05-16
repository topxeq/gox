package main

import (
	"bufio"
	"bytes"

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
	"runtime"

	execq "github.com/topxeq/qlang/exec"
	specq "github.com/topxeq/qlang/spec"

	// GUI related end

	qlarchivezip "github.com/topxeq/qlang/lib/archive/zip"
	qlbytes "github.com/topxeq/qlang/lib/bytes"
	qlcrypto "github.com/topxeq/qlang/lib/crypto"
	qlcryptohmac "github.com/topxeq/qlang/lib/crypto/hmac"
	qlcryptomd5 "github.com/topxeq/qlang/lib/crypto/md5"
	qlcryptorand "github.com/topxeq/qlang/lib/crypto/rand"
	qlcryptorsa "github.com/topxeq/qlang/lib/crypto/rsa"
	qlcryptosha1 "github.com/topxeq/qlang/lib/crypto/sha1"
	qlcryptosha256 "github.com/topxeq/qlang/lib/crypto/sha256"
	qlcryptox509 "github.com/topxeq/qlang/lib/crypto/x509"
	qlencodingbase64 "github.com/topxeq/qlang/lib/encoding/base64"
	qlencodingcsv "github.com/topxeq/qlang/lib/encoding/csv"
	qlencodinghex "github.com/topxeq/qlang/lib/encoding/hex"
	qlencodingjson "github.com/topxeq/qlang/lib/encoding/json"
	qlencodingpem "github.com/topxeq/qlang/lib/encoding/pem"
	qlencodingxml "github.com/topxeq/qlang/lib/encoding/xml"
	qlioioutil "github.com/topxeq/qlang/lib/io/ioutil"

	qlnethttp "github.com/topxeq/qlang/lib/net/http"
	qlneturl "github.com/topxeq/qlang/lib/net/url"

	qlbufio "github.com/topxeq/qlang/lib/bufio"
	qlsync "github.com/topxeq/qlang/lib/sync"
	qltime "github.com/topxeq/qlang/lib/time"

	qlruntime "github.com/topxeq/qlang/lib/runtime"
	qlruntimedebug "github.com/topxeq/qlang/lib/runtime/debug"

	qlos "github.com/topxeq/qlang/lib/os"
	qlpath "github.com/topxeq/qlang/lib/path"
	qlpathfilepath "github.com/topxeq/qlang/lib/path/filepath"
	qlsort "github.com/topxeq/qlang/lib/sort"
	qlstrings "github.com/topxeq/qlang/lib/strings"

	qlgithubbeeviketree "github.com/topxeq/qlang/lib/github.com/beevik/etree"
	qlgithubtopxeqsqltk "github.com/topxeq/qlang/lib/github.com/topxeq/sqltk"
	qlgithubtopxeqtk "github.com/topxeq/qlang/lib/github.com/topxeq/tk"

	// full version related start
	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/godror/godror"

	"image"
	"image/color"
	"image/draw"
	"image/png"

	// GUI related start
	"github.com/topxeq/imagetk"
	"gonum.org/v1/plot/plotutil"
	// GUI related end

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"

	// full version related end

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"

	// full version related start

	// full version related end

	// GUI related start
	// full version related start
	"github.com/sqweek/dialog"
	// full version related end
	// GUI related end

	"github.com/topxeq/tk"

	// GUI related start
	// full version related start
	"github.com/AllenDang/giu"
	g "github.com/AllenDang/giu"
	"github.com/AllenDang/giu/imgui"

	// full version related end

	"github.com/topxeq/govcl/vcl"
	"github.com/topxeq/govcl/vcl/api"
	"github.com/topxeq/govcl/vcl/rtl"
	"github.com/topxeq/govcl/vcl/types"
	// GUI related end
)

// Non GUI related

var versionG = "0.987a"

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

// func getVar(nameA string) interface{} {
// 	varMutexG.Lock()
// 	rs, ok := variableG[nameA]
// 	varMutexG.Unlock()

// 	if !ok {
// 		tk.GenerateErrorString("no key")
// 	}
// 	return rs
// }

// func setVar(nameA string, valueA interface{}) {
// 	varMutexG.Lock()
// 	variableG[nameA] = valueA
// 	varMutexG.Unlock()
// }

func qlEval(strA string) string {
	vmT := qlang.New()

	errT := vmT.SafeEval(`TXResultG=` + strA)

	if errT != nil {
		return errT.Error()
	}

	return tk.Spr("%v", vmT.Var("TXResultG"))
}

// func getClipText() string {
// 	textT, errT := clipboard.ReadAll()
// 	if errT != nil {
// 		return tk.GenerateErrorStringF("could not get text from clipboard: %v", errT.Error())
// 	}

// 	return textT
// }

// func setClipText(textA string) {
// 	clipboard.WriteAll(textA)
// }

func panicIt(valueA interface{}) {
	panic(valueA)
}

// func checkErrorFunc(errA error, funcA func()) {
// 	if errA != nil {
// 		tk.PlErr(errA)

// 		if funcA != nil {
// 			funcA()
// 		}

// 		os.Exit(1)
// 	}

// }

// func checkError(errA error, funcsA ...(func())) {
// 	if errA != nil {
// 		tk.PlErr(errA)

// 		if funcsA != nil {
// 			for _, v := range funcsA {
// 				v()
// 			}
// 		}

// 		os.Exit(1)
// 	}

// }

// func checkErrorString(strA string, funcA func()) {
// 	if tk.IsErrorString(strA) {
// 		tk.PlErrString(strA)

// 		if funcA != nil {
// 			funcA()
// 		}

// 		os.Exit(1)
// 	}

// }

// func remove(aryA []interface{}, startA int, endA int) []interface{} {
// 	if startA < 0 || startA >= len(aryA) {
// 		tk.Pl("Runtime error: %v", "index out of range")
// 		exit()
// 	}

// 	if endA < 0 || endA >= len(aryA) {
// 		tk.Pl("Runtime error: %v", "index out of range")
// 		exit()
// 	}

// 	return append(aryA[:startA], aryA[endA+1:]...)
// 	// if idxT == 0 {
// 	// 	return ayrA[idxT + 1:]
// 	// }

// 	// if idxT == len(aryA) - 1 {
// 	// 	return ayrA[0:len(aryA) - 1]
// 	// }

// 	// return append(aryA[:idxA], aryA[idxA+1:]...)

// }

// func toStringFromRuneSlice(sliceA []rune) string {
// 	return string(sliceA)
// }

// // toInt converts all reflect.Value-s into int.
// func toInt(vA interface{}) int {
// 	v := reflect.ValueOf(&vA)
// 	i, _ := tryToInt(v)
// 	return i
// }

// // tryToInt attempts to convert a value to an int.
// // If it cannot (in the case of a non-numeric string, a struct, etc.)
// // it returns 0 and an error.
// func tryToInt(v reflect.Value) (int, error) {
// 	if v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
// 		v = v.Elem()
// 	}
// 	switch v.Kind() {
// 	case reflect.Float64, reflect.Float32:
// 		return int(v.Float()), nil
// 	case reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8, reflect.Int:
// 		return int(v.Int()), nil
// 	case reflect.Bool:
// 		if v.Bool() {
// 			return 1, nil
// 		}
// 		return 0, nil
// 	case reflect.String:
// 		s := v.String()
// 		var i int64
// 		var err error
// 		if strings.HasPrefix(s, "0x") {
// 			i, err = strconv.ParseInt(s, 16, 64)
// 		} else {
// 			i, err = strconv.ParseInt(s, 10, 64)
// 		}
// 		if err == nil {
// 			return int(i), nil
// 		}
// 	}
// 	return 0, errors.New("couldn't convert to integer")
// }

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

		// if retG != notFoundG {
		// 	fmt.Println(retG)
		// }

		// rs, _ := vmT.GetVar("outG")

		// if !ok {
		// 	return ""
		// }

		return retG
		// full version related start
	} else {
		return tk.SystemCmd("gox", append([]string{codeA}, argsA...)...)
	}

}

// full version related start
func newRGBA(r, g, b, a uint8) color.RGBA {
	return color.RGBA{r, g, b, a}
}

func newNRGBAFromHex(strA string) color.NRGBA {
	r, g, b, a := tk.ParseHexColor(strA)

	return color.NRGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
}

func newRGBAFromHex(strA string) color.RGBA {
	r, g, b, a := tk.ParseHexColor(strA)

	return color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
}

func newPlotXY(xA, yA float64) plotter.XY {
	return plotter.XY{X: xA, Y: yA}
}

func loadRGBAFromImage(imageA image.Image) (*image.RGBA, error) {
	switch imageT := imageA.(type) {
	case *image.RGBA:
		return imageT, nil
	default:
		rgba := image.NewRGBA(imageT.Bounds())
		draw.Draw(rgba, imageT.Bounds(), imageT, image.Pt(0, 0), draw.Src)
		return rgba, nil
	}

}

func LoadPlotImage(p *plot.Plot, w vg.Length, h vg.Length) (*image.RGBA, error) {

	var bufT bytes.Buffer

	writerT, errT := p.WriterTo(w, h, "png")

	if errT != nil {
		return nil, errT
	}

	_, errT = writerT.WriteTo(&bufT)

	if errT != nil {
		return nil, errT
	}

	readerT := bytes.NewReader(bufT.Bytes())

	// defer readerT.Close()

	// imgFile, err := os.Open(imgPath)
	// if err != nil {
	// 	return nil, err
	// }
	// defer imgFile.Close()

	img, err := png.Decode(readerT)
	if err != nil {
		return nil, err
	}

	switch trueImg := img.(type) {
	case *image.RGBA:
		return trueImg, nil
	default:
		rgba := image.NewRGBA(trueImg.Bounds())
		draw.Draw(rgba, trueImg.Bounds(), trueImg, image.Pt(0, 0), draw.Src)
		return rgba, nil
	}
}

// type MyXYer plotter.XY

// func (v MyXYer) Len() int {
// 	return 1
// }

// func (v MyXYer) XY(int) (x, y float64) {
// 	return 1
// }

// full version related end

func setValue(p interface{}, v interface{}) {
	// tk.Pl("%#v", reflect.TypeOf(p).Kind())
	// p = v

	srcRef := reflect.ValueOf(v)
	vp := reflect.ValueOf(p)
	vp.Elem().Set(srcRef)
}

func getValue(p interface{}) interface{} {
	vp := reflect.Indirect(reflect.ValueOf(p))
	return vp.Interface()
}

func bitXor(p interface{}, v interface{}) interface{} {
	switch p.(type) {
	case int:
		return p.(int) ^ v.(int)
	case int64:
		return p.(int64) ^ v.(int64)
	case int32:
		return p.(int32) ^ v.(int32)
	case int16:
		return p.(int16) ^ v.(int16)
	case int8:
		return p.(int8) ^ v.(int8)
	case uint64:
		return p.(uint64) ^ v.(uint64)
	case uint32:
		return p.(uint32) ^ v.(uint32)
	case uint16:
		return p.(uint16) ^ v.(uint16)
	case uint8:
		return p.(uint8) ^ v.(uint8)
	case uint:
		return p.(uint) ^ v.(uint)
	}

	return 0
}

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
		"setValue":         setValue,
		"getValue":         getValue,
		"bitXor":           bitXor,
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

	qlang.Import("tk", qlgithubtopxeqtk.Exports)
	qlang.Import("github_topxeq_tk", qlgithubtopxeqtk.Exports)

	qlang.Import("os", qlos.Exports)

	qlang.Import("strings", qlstrings.Exports)
	qlang.Import("bytes", qlbytes.Exports)
	qlang.Import("io_ioutil", qlioioutil.Exports)

	qlang.Import("sort", qlsort.Exports)

	qlang.Import("time", qltime.Exports)
	qlang.Import("bufio", qlbufio.Exports)
	qlang.Import("sync", qlsync.Exports)

	qlang.Import("net_url", qlneturl.Exports)
	qlang.Import("net_http", qlnethttp.Exports)

	qlang.Import("runtime", qlruntime.Exports)
	qlang.Import("runtime_debug", qlruntimedebug.Exports)

	qlang.Import("path_filepath", qlpathfilepath.Exports)
	qlang.Import("path", qlpath.Exports)

	qlang.Import("archive_zip", qlarchivezip.Exports)

	qlang.Import("encoding_pem", qlencodingpem.Exports)
	qlang.Import("encoding_base64", qlencodingbase64.Exports)
	qlang.Import("encoding_csv", qlencodingcsv.Exports)
	qlang.Import("encoding_hex", qlencodinghex.Exports)
	qlang.Import("encoding_json", qlencodingjson.Exports)
	qlang.Import("encoding_xml", qlencodingxml.Exports)

	qlang.Import("crypto", qlcrypto.Exports)
	qlang.Import("crypto_rand", qlcryptorand.Exports)
	qlang.Import("crypto_hmac", qlcryptohmac.Exports)
	qlang.Import("crypto_rsa", qlcryptorsa.Exports)
	qlang.Import("crypto_sha256", qlcryptosha256.Exports)
	qlang.Import("crypto_sha1", qlcryptosha1.Exports)
	qlang.Import("crypto_x509", qlcryptox509.Exports)
	qlang.Import("crypto_md5", qlcryptomd5.Exports)

	qlang.Import("github_beevik_etree", qlgithubbeeviketree.Exports)
	qlang.Import("github_topxeq_sqltk", qlgithubtopxeqsqltk.Exports)

}

func showHelp() {
	tk.Pl("Gox by TopXeQ V%v\n", versionG)

	tk.Pl("Usage: gox [-v|-h] test.gox next.js, ...\n")
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

func initLCLLib() (result error) {

	api.DoLibInit()

	result = nil

	return result
}

// func syncInitLCL() error {
// 	var errT error

// 	vcl.ThreadSync(func() {
// 		initLCL()
// 	})

// 	return errT
// }

func initLCL() error {

	runtime.LockOSThread()

	// startThreadID := tk.GetCurrentThreadID()

	api.CloseLib()

	errT := initLCLLib()

	if errT != nil {
		tk.Pl("failed to init lib: %v, try to download the LCL lib...", errT)

		applicationPathT := tk.GetApplicationPath()

		osT := tk.GetOSName()

		if tk.Contains(osT, "inux") {
			rs := tk.DownloadFile("http://scripts.frenchfriend.net/pub/liblcl.so", applicationPathT, "liblcl.so", false)

			if tk.IsErrorString(rs) {
				tk.Pl("failed to download LCL file.")
				return tk.Errf("failed to download LCL file.")
			}
		} else if tk.Contains(osT, "arwin") {
			rs := tk.DownloadFile("http://scripts.frenchfriend.net/pub/liblcl.dylib", applicationPathT, "liblcl.dylib", false)

			if tk.IsErrorString(rs) {
				tk.Pl("failed to download LCL file.")
				return tk.Errf("failed to download LCL file.")
			}
		} else {
			rs := tk.DownloadFile("http://scripts.frenchfriend.net/pub/liblcl.dll", applicationPathT, "liblcl.dll", false)

			if tk.IsErrorString(rs) {
				tk.Pl("failed to download LCL file.")
				return tk.Errf("failed to download LCL file.")
			}
		}

		errT = initLCLLib()

		if errT != nil {
			tk.Pl("failed to install lib: %v", errT)
			return tk.Errf("failed to install lib: %v", errT)
		}
	}

	api.DoResInit()

	api.DoImportInit()

	api.DoDefInit()

	rtl.DoRtlInit()

	vcl.DoInit()

	// if verboseG {

	// 	endThreadID := tk.GetCurrentThreadID()

	// 	tk.Pl("start tid: %v, end tid: %v", startThreadID, endThreadID)

	// 	if endThreadID != startThreadID {
	// 		return tk.Errf("failed to init lcl lib: %v", "thread not same")
	// 	}
	// }

	return nil
}

func getVclApplication() *vcl.TApplication {
	return vcl.Application
}

// func newLclAnchors() *types.TAnchors {
// 	a := types.TAnchors(rtl.Include(0, types.AkTop, types.AkBottom, types.AkLeft, types.AkRight))

// }

// func getLCLEvent(funcA *execq.Function) *() {

// }

// func getTNotifyEvent(funcA *execq.Function) *vcl.TNotifyEvent {
// 	var f vcl.TNotifyEvent = func(sender vcl.IObject) {
// 		funcA.Call(execq.NewStack(), sender)
// 	}

// 	return &f
// }

func NewTNotifyEvent(funcA *execq.Function) *vcl.TNotifyEvent {
	var f vcl.TNotifyEvent = func(sender vcl.IObject) {
		funcA.Call(execq.NewStack(), sender)
	}

	return &f
}

func NewTKeyEvent(funcA *execq.Function) *vcl.TKeyEvent {
	var f vcl.TKeyEvent = func(sender vcl.IObject, key *types.Char, shift types.TShiftState) {
		funcA.Call(execq.NewStack(), sender, key, shift)
	}

	return &f
	// return nil
}

// full version related start

var vgInch = float64(vg.Inch)

// full version related end

func importQLGUIPackages() {
	// full version related start
	var plotExports = map[string]interface{}{
		"New": plot.New,
		// "SetTitleText":  plot.SetTitleText,
		"NewXY":         newPlotXY,
		"AddLinePoints": plotutil.AddLinePoints,
		"Inch":          vgInch,

		"XYs": specq.StructOf((*plotter.XYs)(nil)),
		"XY":  specq.StructOf((*plotter.XY)(nil)),
	}

	qlang.Import("plot", plotExports)

	var imagetkExports = map[string]interface{}{
		"NewImageTK": imagetk.NewImageTK,
	}

	qlang.Import("imagetk", imagetkExports)

	var guiExports = map[string]interface{}{
		"NewMasterWindow":         g.NewMasterWindow,
		"SingleWindow":            g.SingleWindow,
		"Window":                  g.Window,
		"SingleWindowWithMenuBar": g.SingleWindowWithMenuBar,
		"WindowV":                 g.WindowV,

		"MasterWindowFlagsNotResizable": g.MasterWindowFlagsNotResizable,
		"MasterWindowFlagsMaximized":    g.MasterWindowFlagsMaximized,
		"MasterWindowFlagsFloating":     g.MasterWindowFlagsFloating,

		// "Layout":          g.Layout,

		"NewTextureFromRgba": g.NewTextureFromRgba,

		"Label":                  g.Label,
		"Line":                   g.Line,
		"Button":                 g.Button,
		"InvisibleButton":        g.InvisibleButton,
		"ImageButton":            g.ImageButton,
		"InputTextMultiline":     g.InputTextMultiline,
		"Checkbox":               g.Checkbox,
		"RadioButton":            g.RadioButton,
		"Child":                  g.Child,
		"ComboCustom":            g.ComboCustom,
		"Combo":                  g.Combo,
		"ContextMenu":            g.ContextMenu,
		"Group":                  g.Group,
		"Image":                  g.Image,
		"ImageWithFile":          g.ImageWithFile,
		"ImageWithUrl":           g.ImageWithUrl,
		"InputText":              g.InputText,
		"InputTextV":             g.InputTextV,
		"InputTextFlagsPassword": g.InputTextFlagsPassword,
		"InputInt":               g.InputInt,
		"InputFloat":             g.InputFloat,
		"MainMenuBar":            g.MainMenuBar,
		"MenuBar":                g.MenuBar,
		"MenuItem":               g.MenuItem,
		"PopupModal":             g.PopupModal,
		"OpenPopup":              g.OpenPopup,
		"CloseCurrentPopup":      g.CloseCurrentPopup,
		"ProgressBar":            g.ProgressBar,
		"Separator":              g.Separator,
		"SliderInt":              g.SliderInt,
		"SliderFloat":            g.SliderFloat,
		"HSplitter":              g.HSplitter,
		"VSplitter":              g.VSplitter,
		"TabItem":                g.TabItem,
		"TabBar":                 g.TabBar,
		"Row":                    g.Row,
		"Table":                  g.Table,
		"FastTable":              g.FastTable,
		"Tooltip":                g.Tooltip,
		"TreeNode":               g.TreeNode,
		"Spacing":                g.Spacing,
		"Custom":                 g.Custom,
		"Condition":              g.Condition,
		"ListBox":                g.ListBox,
		"DatePicker":             g.DatePicker,
		"Dummy":                  g.Dummy,
		// "Widget":             g.Widget,

		"PrepareMessageBox": g.PrepareMsgbox,
		"MessageBox":        g.Msgbox,

		"LoadFont": loadFont,

		"GetConfirm": getConfirmGUI,

		"SimpleInfo":      simpleInfo,
		"SimpleError":     simpleError,
		"SelectFile":      selectFileGUI,
		"SelectSaveFile":  selectFileToSaveGUI,
		"SelectDirectory": selectDirectoryGUI,

		"EditFile":   editFile,
		"LoopWindow": loopWindow,

		"LayoutP": g.Layout{},

		"Layout": specq.StructOf((*g.Layout)(nil)),
		"Widget": specq.StructOf((*g.Widget)(nil)),
	}

	qlang.Import("gui", guiExports)
	// full version related end

	var lclExports = map[string]interface{}{
		"NewTNotifyEvent": NewTNotifyEvent,
		"NewTKeyEvent":    NewTKeyEvent,
		"GetApplication":  getVclApplication,
		// "NewApplication":    vcl.NewApplication,
		"InitVCL":           initLCL,
		"InitLCL":           initLCL,
		"NewCheckBox":       vcl.NewCheckBox,
		"NewLabel":          vcl.NewLabel,
		"NewButton":         vcl.NewButton,
		"NewComboBox":       vcl.NewComboBox,
		"NewEdit":           vcl.NewEdit,
		"NewCanvas":         vcl.NewCanvas,
		"NewImage":          vcl.NewImage,
		"NewList":           vcl.NewList,
		"NewListBox":        vcl.NewListBox,
		"NewListView":       vcl.NewListView,
		"NewListColumns":    vcl.NewListColumns,
		"NewListItem":       vcl.NewListItem,
		"NewListItems":      vcl.NewListItems,
		"NewMainMenu":       vcl.NewMainMenu,
		"NewMemo":           vcl.NewMemo,
		"NewMenuItem":       vcl.NewMenuItem,
		"NewMiniWebview":    vcl.NewMiniWebview,
		"NewPaintBox":       vcl.NewPaintBox,
		"NewPanel":          vcl.NewPanel,
		"NewPicture":        vcl.NewPicture,
		"NewPopupMenu":      vcl.NewPopupMenu,
		"NewProgressBar":    vcl.NewProgressBar,
		"NewRadioButton":    vcl.NewRadioButton,
		"NewRadioGroup":     vcl.NewRadioGroup,
		"NewScrollBox":      vcl.NewScrollBox,
		"NewScrollBar":      vcl.NewScrollBar,
		"NewSplitter":       vcl.NewSplitter,
		"NewStatusBar":      vcl.NewStatusBar,
		"NewStatusPanel":    vcl.NewStatusPanel,
		"NewStatusPanels":   vcl.NewStatusPanels,
		"NewTimer":          vcl.NewTimer,
		"NewToolBar":        vcl.NewToolBar,
		"NewToolButton":     vcl.NewToolButton,
		"NewTrayIcon":       vcl.NewTrayIcon,
		"NewStaticText":     vcl.NewStaticText,
		"NewSpinEdit":       vcl.NewSpinEdit,
		"NewSpeedButton":    vcl.NewSpeedButton,
		"NewShape":          vcl.NewShape,
		"NewScreen":         vcl.NewScreen,
		"NewSaveDialog":     vcl.NewSaveDialog,
		"NewReplaceDialog":  vcl.NewReplaceDialog,
		"NewPngImage":       vcl.NewPngImage,
		"NewPen":            vcl.NewPen,
		"NewPageControl":    vcl.NewPageControl,
		"NewOpenDialog":     vcl.NewOpenDialog,
		"NewObject":         vcl.NewObject,
		"NewMouse":          vcl.NewMouse,
		"NewMaskEdit":       vcl.NewMaskEdit,
		"NewLinkLabel":      vcl.NewLinkLabel,
		"NewLabeledEdit":    vcl.NewLabeledEdit,
		"NewJPEGImage":      vcl.NewJPEGImage,
		"NewImageList":      vcl.NewImageList,
		"NewImageButton":    vcl.NewImageButton,
		"NewIcon":           vcl.NewIcon,
		"NewGroupBox":       vcl.NewGroupBox,
		"NewHeaderControl":  vcl.NewHeaderControl,
		"NewHeaderSection":  vcl.NewHeaderSection,
		"NewHeaderSections": vcl.NewHeaderSections,
		"NewGraphic":        vcl.NewGraphic,
		"NewGIFImage":       vcl.NewGIFImage,
		"NewGauge":          vcl.NewGauge,
		"ShowMessage":       vcl.ShowMessage,
		"ShowMessageFmt":    vcl.ShowMessageFmt,
		"MessageDlg":        vcl.MessageDlg,
		"InputBox":          vcl.InputBox,
		"InputQuery":        vcl.InputQuery,
		"ThreadSync":        vcl.ThreadSync,
		"NewFrame":          vcl.NewFrame,
		"SelectDirectory":   vcl.SelectDirectory1,
		"SelectDirectory3":  vcl.SelectDirectory3,
		"NewForm":           vcl.NewForm,
		"NewFontDialog":     vcl.NewFontDialog,
		"NewFont":           vcl.NewFont,
		"NewFlowPanel":      vcl.NewFlowPanel,
		"NewFindDialog":     vcl.NewFindDialog,
		"NewDrawGrid":       vcl.NewDrawGrid,
		"NewDateTimePicker": vcl.NewDateTimePicker,
		"NewControl":        vcl.NewControl,
		"NewComboBoxEx":     vcl.NewComboBoxEx,
		"NewColorListBox":   vcl.NewColorListBox,
		"NewColorDialog":    vcl.NewColorDialog,
		"NewColorBox":       vcl.NewColorBox,
		"NewCheckListBox":   vcl.NewCheckListBox,
		"NewBrush":          vcl.NewBrush,
		"NewBitmap":         vcl.NewBitmap,
		"NewBitBtn":         vcl.NewBitBtn,
		"NewBevel":          vcl.NewBevel,
		"NewApplication":    vcl.NewApplication,
		"NewAction":         vcl.NewAction,
		"NewActionList":     vcl.NewActionList,
		"NewMemoryStream":   vcl.NewMemoryStream,

		// "NewAnchors": types.TAnchors,
		// "RTLInclude": rtl.Include,
		"NewSet":           types.NewSet,
		"AkTop":            types.AkTop,
		"AkBottom":         types.AkBottom,
		"AkLeft":           types.AkLeft,
		"AkRight":          types.AkRight,
		"SsNone":           types.SsNone,
		"SsHorizontal":     types.SsHorizontal,
		"SsVertical":       types.SsVertical,
		"SsBoth":           types.SsBoth,
		"SsAutoHorizontal": types.SsAutoHorizontal,
		"SsAutoVertical":   types.SsAutoVertical,
		"SsAutoBoth":       types.SsAutoBoth,

		"GetLibVersion": vcl.GetLibVersion,

		// values
		"PoDesigned":        types.PoDesigned,
		"PoDefault":         types.PoDefault,
		"PoDefaultPosOnly":  types.PoDefaultPosOnly,
		"PoDefaultSizeOnly": types.PoDefaultSizeOnly,
		"PoScreenCenter":    types.PoScreenCenter,
		"PoMainFormCenter":  types.PoMainFormCenter,
		"PoOwnerFormCenter": types.PoOwnerFormCenter,
		"PoWorkAreaCenter":  types.PoWorkAreaCenter,
	}

	qlang.Import("lcl", lclExports)
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
			g.Msgbox("Info", tk.Spr("Action cancelled by user"))
			return
		}

		g.Msgbox("Error", tk.Spr("Failed to select file: %v", tk.GetErrorString(fileNameNewT)))
		return
	}

	fcT := tk.LoadStringFromFile(fileNameNewT)

	if tk.IsErrorString(fcT) {
		g.Msgbox("Error", tk.Spr("Failed to load file content: %v", tk.GetErrorString(fileNameNewT)))
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
			g.Msgbox("Info", tk.Spr("Action cancelled by user"))
			return
		}

		g.Msgbox("Error", tk.Spr("Failed to select file: %v", tk.GetErrorString(fileNameNewT)))
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
			g.Msgbox("Error", tk.Spr("Failed to save: %v", rs))
			return
		}

		g.Msgbox("Info", tk.Spr("File saved to: %v", editFileNameG))

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
			g.Msgbox("Error", tk.Spr("Failed to save: %v", rs))
			return
		}

		g.Msgbox("Info", tk.Spr("File saved to file: %v", editFileNameG))

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
	g.OpenPopup("Please enter:##EncryptInputSecureCode")
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
	g.OpenPopup("Please enter:##DecryptInputSecureCode")
}

func editRun() {
	imgui.CloseCurrentPopup()

	runScript(editorG.GetText(), "", editArgsG)
}

func editRunClick() {
	g.OpenPopup("Please enter:##RunInputArgs")
}

func onButtonCloseClick() {
	exit()
}

func loopWindow(windowA *g.MasterWindow, loopA func()) {
	// wnd := g.NewMasterWindow("Gox Editor", 800, 600, 0, loadFont)

	windowA.Main(loopA)

}

func editorLoop() {
	g.SingleWindow("Gox Editor", g.Layout{
		g.Label(editFileNameG + editFileCleanFlagG),
		g.Dummy(30, 0),
		g.Line(
			g.Button("Load", editorLoad),
			g.Button("Save", editorSave),
			g.Button("Save As...", editorSaveAs),
			g.Button("Check", func() {

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
				// 	g.Msgbox("Error", tk.Spr("%#v", errT))
				// } else {
				// 	g.Msgbox("Info", "Syntax check passed.")
				// }

			}),
			g.Button("Encrypt", editEncryptClick),
			g.Button("Decrypt", editDecryptClick),
			g.Button("Run", editRunClick),
			g.Button("Close", onButtonCloseClick),
			// g.Button("Get Text", func() {
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
			// g.Button("Set Text", func() {
			// 	editorG.SetText("Set text")
			// 	editFileNameG = "Set text"
			// }),
			// g.Button("Set Error Marker", func() {
			// 	errMarkersG.Clear()
			// 	errMarkersG.Insert(1, "Error message")
			// 	fmt.Println("ErrMarkers Size:", errMarkersG.Size())

			// 	editorG.SetErrorMarkers(errMarkersG)
			// }),
		),
		g.PopupModal("Please enter:##EncryptInputSecureCode", g.Layout{
			g.Line(
				g.Label("Secure code"),
				g.InputTextV("", 40, &editSecureCodeG, g.InputTextFlagsPassword, nil, nil),
			),
			g.Line(
				g.Button("Ok", editEncrypt),
				g.Button("Cancel", func() { imgui.CloseCurrentPopup() }),
			),
		}),
		g.PopupModal("Please enter:##DecryptInputSecureCode", g.Layout{
			g.Line(
				g.Label("Secure code"),
				g.InputTextV("", 40, &editSecureCodeG, g.InputTextFlagsPassword, nil, nil),
			),
			g.Line(
				g.Button("Ok", editDecrypt),
				g.Button("Cancel", func() { imgui.CloseCurrentPopup() }),
			),
		}),
		g.PopupModal("Please enter:##RunInputArgs", g.Layout{
			g.Line(
				g.Label("Arguments to pass to VM"),
				g.InputText("", 80, &editArgsG),
			),
			g.Line(
				g.Button("Ok", editRun),
				g.Button("Cancel", func() { imgui.CloseCurrentPopup() }),
			),
		}),
		g.Custom(func() {
			editorG.Render("Hello", imgui.Vec2{X: 0, Y: 0}, true)
			if giu.IsItemHovered() {
				if editorG.IsTextChanged() {
					editFileCleanFlagG = "*"
				}
			}
		}),
		g.PrepareMsgbox(),
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

	wnd := g.NewMasterWindow("Gox Editor", 800, 600, 0, loadFont)
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

	script := fcT

	errT := qlVMG.SafeEval(script)
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
