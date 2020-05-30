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
	"runtime"

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

	qlgithub_topxeq_govcl_vcl "github.com/topxeq/qlang/lib/github.com/topxeq/govcl/vcl"
	qlgithub_topxeq_govcl_vcl_api "github.com/topxeq/qlang/lib/github.com/topxeq/govcl/vcl/api"
	qlgithub_topxeq_govcl_vcl_rtl "github.com/topxeq/qlang/lib/github.com/topxeq/govcl/vcl/rtl"
	qlgithub_topxeq_govcl_vcl_types "github.com/topxeq/qlang/lib/github.com/topxeq/govcl/vcl/types"

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

	"github.com/topxeq/govcl/vcl"
	"github.com/topxeq/govcl/vcl/api"
	"github.com/topxeq/govcl/vcl/rtl"
	"github.com/topxeq/govcl/vcl/types"
	// GUI related end
)

// Non GUI related

var versionG = "0.996a"

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
	qlang.Import("github_topxeq_govcl_vcl", qlgithub_topxeq_govcl_vcl.Exports)
	qlang.Import("vcl", qlgithub_topxeq_govcl_vcl.Exports)
	qlang.Import("github_topxeq_govcl_vcl_types", qlgithub_topxeq_govcl_vcl_types.Exports)
	qlang.Import("vcl_types", qlgithub_topxeq_govcl_vcl_types.Exports)
	qlang.Import("github_topxeq_govcl_vcl_api", qlgithub_topxeq_govcl_vcl_api.Exports)
	qlang.Import("vcl_api", qlgithub_topxeq_govcl_vcl_api.Exports)
	qlang.Import("github_topxeq_govcl_vcl_rtl", qlgithub_topxeq_govcl_vcl_rtl.Exports)
	qlang.Import("vcl_rtl", qlgithub_topxeq_govcl_vcl_rtl.Exports)

	qlang.Import("github_AllenDang_giu", qlgithub_AllenDang_giu.Exports)
	qlang.Import("giu", qlgithub_AllenDang_giu.Exports)
	qlang.Import("github_AllenDang_giu_imgui", qlgithub_AllenDang_giu_imgui.Exports)
	qlang.Import("giu_imgui", qlgithub_AllenDang_giu_imgui.Exports)
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

	defer func() {
		if r := recover(); r != nil {
			// fmt.Printf("initLCLLib: %v\n", r)

			result = tk.Errf("initLCLLib: %v\n", r)
		}
	}()

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

func NewTKeyPressEvent(funcA *execq.Function) *vcl.TKeyPressEvent {
	var f vcl.TKeyPressEvent = func(sender vcl.IObject, key *types.Char) {
		funcA.Call(execq.NewStack(), sender, key)
	}

	return &f
	// return nil
}

func NewTMouseEvent(funcA *execq.Function) *vcl.TMouseEvent {
	var f vcl.TMouseEvent = func(sender vcl.IObject, button types.TMouseButton, shift types.TShiftState, x, y int32) {
		funcA.Call(execq.NewStack(), sender, button, shift, x, y)
	}

	return &f
	// return nil
}

func NewTMouseMoveEvent(funcA *execq.Function) *vcl.TMouseMoveEvent {
	var f vcl.TMouseMoveEvent = func(sender vcl.IObject, shift types.TShiftState, x, y int32) {
		funcA.Call(execq.NewStack(), sender, shift, x, y)
	}

	return &f
	// return nil
}

func NewTExceptionEvent(funcA *execq.Function) *vcl.TExceptionEvent {
	var f vcl.TExceptionEvent = func(sender vcl.IObject, e *vcl.Exception) {
		funcA.Call(execq.NewStack(), sender, e)
	}

	return &f
}

func NewTCloseEvent(funcA *execq.Function) *vcl.TCloseEvent {
	var f vcl.TCloseEvent = func(sender vcl.IObject, action *types.TCloseAction) { // TCloseAction int32
		funcA.Call(execq.NewStack(), sender, action)
	}

	return &f
}

func NewTCloseQueryEvent(funcA *execq.Function) *vcl.TCloseQueryEvent {
	var f vcl.TCloseQueryEvent = func(sender vcl.IObject, canClose *bool) {
		funcA.Call(execq.NewStack(), sender, canClose)
	}

	return &f
}

func NewTContextPopupEvent(funcA *execq.Function) *vcl.TContextPopupEvent {
	var f vcl.TContextPopupEvent = func(sender vcl.IObject, mousePos types.TPoint, handled *bool) {
		funcA.Call(execq.NewStack(), sender, mousePos, handled)
	}

	return &f
}

func NewTDragDropEvent(funcA *execq.Function) *vcl.TDragDropEvent {
	var f vcl.TDragDropEvent = func(sender vcl.IObject, source vcl.IObject, x, y int32) {
		funcA.Call(execq.NewStack(), sender, source, x, y)
	}

	return &f
}

func NewTDragOverEvent(funcA *execq.Function) *vcl.TDragOverEvent {
	var f vcl.TDragOverEvent = func(sender vcl.IObject, source vcl.IObject, x, y int32, state types.TDragState, accept *bool) {
		funcA.Call(execq.NewStack(), sender, source, x, y, state, accept)
	}

	return &f
}

func NewTStartDragEvent(funcA *execq.Function) *vcl.TStartDragEvent {
	var f vcl.TStartDragEvent = func(sender vcl.IObject, dragObject *vcl.TDragObject) {
		funcA.Call(execq.NewStack(), sender, dragObject)
	}

	return &f
}

func NewTEndDragEvent(funcA *execq.Function) *vcl.TEndDragEvent {
	var f vcl.TEndDragEvent = func(sender vcl.IObject, target vcl.IObject, x, y int32) {
		funcA.Call(execq.NewStack(), sender, target, x, y)
	}

	return &f
}

func NewTAlignPositionEvent(funcA *execq.Function) *vcl.TAlignPositionEvent {
	var f vcl.TAlignPositionEvent = func(sender *vcl.TWinControl, control *vcl.TControl, newLeft, newTop, newWidth, newHeight *int32, alignRect *types.TRect, alignInfo types.TAlignInfo) {
		funcA.Call(execq.NewStack(), sender, control, newLeft, newTop, newWidth, newHeight, alignRect, alignInfo)
	}

	return &f
}

func NewTDockDropEvent(funcA *execq.Function) *vcl.TDockDropEvent {
	var f vcl.TDockDropEvent = func(sender vcl.IObject, source *vcl.TDragDockObject, x, y int32) {
		funcA.Call(execq.NewStack(), sender, source, x, y)
	}

	return &f
}

func NewTDockOverEvent(funcA *execq.Function) *vcl.TDockOverEvent {
	var f vcl.TDockOverEvent = func(sender vcl.IObject, source *vcl.TDragDockObject, x, y int32, state types.TDragState, accept *bool) {
		funcA.Call(execq.NewStack(), sender, source, x, y, state, accept)
	}

	return &f
}

func NewTStartDockEvent(funcA *execq.Function) *vcl.TStartDockEvent {
	var f vcl.TStartDockEvent = func(sender vcl.IObject, dragObject *vcl.TDragDockObject) {
		funcA.Call(execq.NewStack(), sender, dragObject)
	}

	return &f
}

func NewTUnDockEvent(funcA *execq.Function) *vcl.TUnDockEvent {
	var f vcl.TUnDockEvent = func(sender vcl.IObject, client *vcl.TControl, newTarget *vcl.TControl, allow *bool) {
		funcA.Call(execq.NewStack(), sender, client, newTarget, allow)
	}

	return &f
}

func NewTGetSiteInfoEvent(funcA *execq.Function) *vcl.TGetSiteInfoEvent {
	var f vcl.TGetSiteInfoEvent = func(sender vcl.IObject, dockClient *vcl.TControl, influenceRect *types.TRect, mousePos types.TPoint, canDock *bool) {
		funcA.Call(execq.NewStack(), sender, dockClient, influenceRect, mousePos, canDock)
	}

	return &f
}

func NewTMouseWheelEvent(funcA *execq.Function) *vcl.TMouseWheelEvent {
	var f vcl.TMouseWheelEvent = func(sender vcl.IObject, shift types.TShiftState, wheelDelta, x, y int32, handled *bool) {
		funcA.Call(execq.NewStack(), sender, shift, wheelDelta, x, y, handled)
	}

	return &f
}

func NewTMouseWheelUpDownEvent(funcA *execq.Function) *vcl.TMouseWheelUpDownEvent {
	var f vcl.TMouseWheelUpDownEvent = func(sender vcl.IObject, shift types.TShiftState, mousePos types.TPoint, handled *bool) {
		funcA.Call(execq.NewStack(), sender, shift, mousePos, handled)
	}

	return &f
}

func NewTMessageEvent(funcA *execq.Function) *vcl.TMessageEvent {
	var f vcl.TMessageEvent = func(msg *types.TMsg, handled *bool) {
		funcA.Call(execq.NewStack(), msg, handled)
	}

	return &f
}

func NewTHelpEvent(funcA *execq.Function) *vcl.THelpEvent {
	var f vcl.THelpEvent = func(command uint16, data types.THelpEventData, callhelp, result *bool) {
		funcA.Call(execq.NewStack(), command, data, callhelp, result)
	}

	return &f
}

func NewTWebTitleChangeEvent(funcA *execq.Function) *vcl.TWebTitleChangeEvent {
	var f vcl.TWebTitleChangeEvent = func(sender vcl.IObject, text string) {
		funcA.Call(execq.NewStack(), sender, text)
	}

	return &f
}

func NewTWebJSExternalEvent(funcA *execq.Function) *vcl.TWebJSExternalEvent {
	var f vcl.TWebJSExternalEvent = func(sender vcl.IObject, funcName, args string, retVal *string) {
		funcA.Call(execq.NewStack(), sender, funcName, args, retVal)
	}

	return &f
}

func NewTMeasureItemEvent(funcA *execq.Function) *vcl.TMeasureItemEvent {
	var f vcl.TMeasureItemEvent = func(control *vcl.TWinControl, index int32, height *int32) {
		funcA.Call(execq.NewStack(), control, index, height)
	}

	return &f
}

func NewTMovedEvent(funcA *execq.Function) *vcl.TMovedEvent {
	var f vcl.TMovedEvent = func(sender vcl.IObject, fromIndex, toIndex int32) {
		funcA.Call(execq.NewStack(), sender, fromIndex, toIndex)
	}

	return &f
}

func NewTDrawCellEvent(funcA *execq.Function) *vcl.TDrawCellEvent {
	var f vcl.TDrawCellEvent = func(sender vcl.IObject, aCol, aRow int32, aRect types.TRect, state types.TGridDrawState) {
		funcA.Call(execq.NewStack(), sender, aCol, aRow, aRect, state)
	}

	return &f
}

func NewTSelectCellEvent(funcA *execq.Function) *vcl.TSelectCellEvent {
	var f vcl.TSelectCellEvent = func(sender vcl.IObject, aCol, aRow int32, canSelect *bool) {
		funcA.Call(execq.NewStack(), sender, aCol, aRow, canSelect)
	}

	return &f
}

func NewTGetEditEvent(funcA *execq.Function) *vcl.TGetEditEvent {
	var f vcl.TGetEditEvent = func(sender vcl.IObject, aCol, aRow int32, value *string) {
		funcA.Call(execq.NewStack(), sender, aCol, aRow, value)
	}

	return &f
}

func NewTSetEditEvent(funcA *execq.Function) *vcl.TSetEditEvent {
	var f vcl.TSetEditEvent = func(sender vcl.IObject, aCol, aRow int32, value string) {
		funcA.Call(execq.NewStack(), sender, aCol, aRow, value)
	}

	return &f
}

func NewTDropFilesEvent(funcA *execq.Function) *vcl.TDropFilesEvent {
	var f vcl.TDropFilesEvent = func(sender vcl.IObject, aFileNames []string) {
		funcA.Call(execq.NewStack(), sender, aFileNames)
	}

	return &f
}

func NewTConstrainedResizeEvent(funcA *execq.Function) *vcl.TConstrainedResizeEvent {
	var f vcl.TConstrainedResizeEvent = func(sender vcl.IObject, minWidth, minHeight, maxWidth, maxHeight *int32) {
		funcA.Call(execq.NewStack(), sender, minWidth, minHeight, maxWidth, maxHeight)
	}

	return &f
}

func NewTWndProcEvent(funcA *execq.Function) *vcl.TWndProcEvent {
	var f vcl.TWndProcEvent = func(msg *types.TMessage) {
		funcA.Call(execq.NewStack(), msg)
	}

	return &f
}

func NewTSectionNotifyEvent(funcA *execq.Function) *vcl.TSectionNotifyEvent {
	var f vcl.TSectionNotifyEvent = func(headerControl *vcl.THeaderControl, section *vcl.THeaderSection) {
		funcA.Call(execq.NewStack(), headerControl, section)
	}

	return &f
}

func NewTSectionTrackEvent(funcA *execq.Function) *vcl.TSectionTrackEvent {
	var f vcl.TSectionTrackEvent = func(headerControl *vcl.THeaderControl, section *vcl.THeaderSection, width int32, state types.TSectionTrackState) {
		funcA.Call(execq.NewStack(), headerControl, section, width, state)
	}

	return &f
}

func NewTSectionDragEvent(funcA *execq.Function) *vcl.TSectionDragEvent {
	var f vcl.TSectionDragEvent = func(sender vcl.IObject, fromSection, toSection *vcl.THeaderSection, allowDrag *bool) {
		funcA.Call(execq.NewStack(), sender, fromSection, toSection, allowDrag)
	}

	return &f
}

func NewTSysLinkEvent(funcA *execq.Function) *vcl.TSysLinkEvent {
	var f vcl.TSysLinkEvent = func(sender vcl.IObject, link string, linkType types.TSysLinkType) {
		funcA.Call(execq.NewStack(), sender, link, linkType)
	}

	return &f
}

func NewTDrawItemEvent(funcA *execq.Function) *vcl.TDrawItemEvent {
	var f vcl.TDrawItemEvent = func(control vcl.IWinControl, index int32, aRect types.TRect, state types.TOwnerDrawState) {
		funcA.Call(execq.NewStack(), control, index, aRect, state)
	}

	return &f
}

func NewTLVSelectItemEvent(funcA *execq.Function) *vcl.TLVSelectItemEvent {
	var f vcl.TLVSelectItemEvent = func(sender vcl.IObject, item *vcl.TListItem, selected bool) {
		funcA.Call(execq.NewStack(), sender, item, selected)
	}

	return &f
}

func NewTLVCheckedItemEvent(funcA *execq.Function) *vcl.TLVCheckedItemEvent {
	var f vcl.TLVCheckedItemEvent = func(sender vcl.IObject, item *vcl.TListItem) {
		funcA.Call(execq.NewStack(), sender, item)
	}

	return &f
}

func NewTLVAdvancedCustomDrawEvent(funcA *execq.Function) *vcl.TLVAdvancedCustomDrawEvent {
	var f vcl.TLVAdvancedCustomDrawEvent = func(sender *vcl.TListView, aRect types.TRect, stage types.TCustomDrawStage, defaultDraw *bool) {
		funcA.Call(execq.NewStack(), sender, aRect, stage, defaultDraw)
	}

	return &f
}

func NewTLVAdvancedCustomDrawItemEvent(funcA *execq.Function) *vcl.TLVAdvancedCustomDrawItemEvent {
	var f vcl.TLVAdvancedCustomDrawItemEvent = func(sender *vcl.TListView, item *vcl.TListItem, state types.TCustomDrawState, Stage types.TCustomDrawStage, defaultDraw *bool) {
		funcA.Call(execq.NewStack(), sender, item, state, Stage, defaultDraw)
	}

	return &f
}

func NewTLVAdvancedCustomDrawSubItemEvent(funcA *execq.Function) *vcl.TLVAdvancedCustomDrawSubItemEvent {
	var f vcl.TLVAdvancedCustomDrawSubItemEvent = func(sender *vcl.TListView, item *vcl.TListItem, subItem int32, state types.TCustomDrawState, stage types.TCustomDrawStage, defaultDraw *bool) {
		funcA.Call(execq.NewStack(), sender, item, subItem, state, stage, defaultDraw)
	}

	return &f
}

func NewTLVChangeEvent(funcA *execq.Function) *vcl.TLVChangeEvent {
	var f vcl.TLVChangeEvent = func(sender vcl.IObject, item *vcl.TListItem, change types.TItemChange) {
		funcA.Call(execq.NewStack(), sender, item, change)
	}

	return &f
}

func NewTLVColumnClickEvent(funcA *execq.Function) *vcl.TLVColumnClickEvent {
	var f vcl.TLVColumnClickEvent = func(sender vcl.IObject, column *vcl.TListColumn) {
		funcA.Call(execq.NewStack(), sender, column)
	}

	return &f
}

func NewTLVCompareEvent(funcA *execq.Function) *vcl.TLVCompareEvent {
	var f vcl.TLVCompareEvent = func(sender vcl.IObject, item1, item2 *vcl.TListItem, data int32, compare *int32) {
		funcA.Call(execq.NewStack(), sender, item1, item2, data, compare)
	}

	return &f
}

func NewTLVOwnerDataEvent(funcA *execq.Function) *vcl.TLVOwnerDataEvent {
	var f vcl.TLVOwnerDataEvent = func(sender vcl.IObject, item *vcl.TListItem) {
		funcA.Call(execq.NewStack(), sender, item)
	}

	return &f
}

func NewTLVOwnerDataFindEvent(funcA *execq.Function) *vcl.TLVOwnerDataFindEvent {
	var f vcl.TLVOwnerDataFindEvent = func(sender vcl.IObject, find types.TItemFind, findString string, findPosition types.TPoint, findData types.TCustomData, startIndex int32, direction types.TSearchDirection, warp bool, index *int32) {
		funcA.Call(execq.NewStack(), sender, find, findString, findPosition, findData, startIndex, direction, warp, index)
	}

	return &f
}

func NewTLVOwnerDataHintEvent(funcA *execq.Function) *vcl.TLVOwnerDataHintEvent {
	var f vcl.TLVOwnerDataHintEvent = func(sender vcl.IObject, startIndex, endIndex int32) {
		funcA.Call(execq.NewStack(), sender, startIndex, endIndex)
	}

	return &f
}

func NewTLVDataHintEvent(funcA *execq.Function) *vcl.TLVDataHintEvent {
	var f vcl.TLVDataHintEvent = func(sender vcl.IObject, startIndex, endIndex int32) {
		funcA.Call(execq.NewStack(), sender, startIndex, endIndex)
	}

	return &f
}

func NewTLVDeletedEvent(funcA *execq.Function) *vcl.TLVDeletedEvent {
	var f vcl.TLVDeletedEvent = func(sender vcl.IObject, item *vcl.TListItem) {
		funcA.Call(execq.NewStack(), sender, item)
	}

	return &f
}

func NewTLVEditedEvent(funcA *execq.Function) *vcl.TLVEditedEvent {
	var f vcl.TLVEditedEvent = func(sender vcl.IObject, item *vcl.TListItem, s *string) {
		funcA.Call(execq.NewStack(), sender, item, s)
	}

	return &f
}

func NewTLVEditingEvent(funcA *execq.Function) *vcl.TLVEditingEvent {
	var f vcl.TLVEditingEvent = func(sender vcl.IObject, item *vcl.TListItem, allowEdit *bool) {
		funcA.Call(execq.NewStack(), sender, item, allowEdit)
	}

	return &f
}

func NewTMenuChangeEvent(funcA *execq.Function) *vcl.TMenuChangeEvent {
	var f vcl.TMenuChangeEvent = func(sender vcl.IObject, source *vcl.TMenuItem, rebuild bool) {
		funcA.Call(execq.NewStack(), sender, source, rebuild)
	}

	return &f
}

func NewTMenuMeasureItemEvent(funcA *execq.Function) *vcl.TMenuMeasureItemEvent {
	var f vcl.TMenuMeasureItemEvent = func(sender vcl.IObject, aCanvas *vcl.TCanvas, width, height *int32) {
		funcA.Call(execq.NewStack(), sender, aCanvas, width, height)
	}

	return &f
}

func NewTTabChangingEvent(funcA *execq.Function) *vcl.TTabChangingEvent {
	var f vcl.TTabChangingEvent = func(sender vcl.IObject, allowChange *bool) {
		funcA.Call(execq.NewStack(), sender, allowChange)
	}

	return &f
}

func NewTUDChangingEvent(funcA *execq.Function) *vcl.TUDChangingEvent {
	var f vcl.TUDChangingEvent = func(sender vcl.IObject, allowChange *bool) {
		funcA.Call(execq.NewStack(), sender, allowChange)
	}

	return &f
}

func NewTUDClickEvent(funcA *execq.Function) *vcl.TUDClickEvent {
	var f vcl.TUDClickEvent = func(sender vcl.IObject, button types.TUDBtnType) {
		funcA.Call(execq.NewStack(), sender, button)
	}

	return &f
}

func NewTTaskDlgClickEvent(funcA *execq.Function) *vcl.TTaskDlgClickEvent {
	var f vcl.TTaskDlgClickEvent = func(sender vcl.IObject, modalResult types.TModalResult, canClose *bool) {
		funcA.Call(execq.NewStack(), sender, modalResult, canClose)
	}

	return &f
}

func NewTTVExpandedEvent(funcA *execq.Function) *vcl.TTVExpandedEvent {
	var f vcl.TTVExpandedEvent = func(sender vcl.IObject, node *vcl.TTreeNode) {
		funcA.Call(execq.NewStack(), sender, node)
	}

	return &f
}

func NewTTVAdvancedCustomDrawEvent(funcA *execq.Function) *vcl.TTVAdvancedCustomDrawEvent {
	var f vcl.TTVAdvancedCustomDrawEvent = func(sender *vcl.TTreeView, aRect types.TRect, stage types.TCustomDrawStage, defaultDraw *bool) {
		funcA.Call(execq.NewStack(), sender, aRect, stage, defaultDraw)
	}

	return &f
}

func NewTTVAdvancedCustomDrawItemEvent(funcA *execq.Function) *vcl.TTVAdvancedCustomDrawItemEvent {
	var f vcl.TTVAdvancedCustomDrawItemEvent = func(sender *vcl.TTreeView, node *vcl.TTreeNode, state types.TCustomDrawState, stage types.TCustomDrawStage, paintImages, defaultDraw *bool) {
		funcA.Call(execq.NewStack(), sender, node, state, stage, paintImages, defaultDraw)
	}

	return &f
}

func NewTTVChangedEvent(funcA *execq.Function) *vcl.TTVChangedEvent {
	var f vcl.TTVChangedEvent = func(sender vcl.IObject, node *vcl.TTreeNode) {
		funcA.Call(execq.NewStack(), sender, node)
	}

	return &f
}

func NewTTVChangingEvent(funcA *execq.Function) *vcl.TTVChangingEvent {
	var f vcl.TTVChangingEvent = func(sender vcl.IObject, node *vcl.TTreeNode, allowChange *bool) {
		funcA.Call(execq.NewStack(), sender, node, allowChange)
	}

	return &f
}

func NewTTVCollapsingEvent(funcA *execq.Function) *vcl.TTVCollapsingEvent {
	var f vcl.TTVCollapsingEvent = func(sender vcl.IObject, node *vcl.TTreeNode, allowCollapse *bool) {
		funcA.Call(execq.NewStack(), sender, node, allowCollapse)
	}

	return &f
}

func NewTTVCompareEvent(funcA *execq.Function) *vcl.TTVCompareEvent {
	var f vcl.TTVCompareEvent = func(sender vcl.IObject, node1, node2 *vcl.TTreeNode, data int32, compare *int32) {
		funcA.Call(execq.NewStack(), sender, node1, node2, data, compare)
	}

	return &f
}

func NewTTVCustomDrawEvent(funcA *execq.Function) *vcl.TTVCustomDrawEvent {
	var f vcl.TTVCustomDrawEvent = func(sender *vcl.TTreeView, aRect types.TRect, defaultDraw *bool) {
		funcA.Call(execq.NewStack(), sender, aRect, defaultDraw)
	}

	return &f
}

func NewTTVCustomDrawItemEvent(funcA *execq.Function) *vcl.TTVCustomDrawItemEvent {
	var f vcl.TTVCustomDrawItemEvent = func(sender *vcl.TTreeView, node *vcl.TTreeNode, state types.TCustomDrawStage, defaultDraw *bool) {
		funcA.Call(execq.NewStack(), sender, node, state, defaultDraw)
	}

	return &f
}

func NewTTVEditedEvent(funcA *execq.Function) *vcl.TTVEditedEvent {
	var f vcl.TTVEditedEvent = func(sender vcl.IObject, node *vcl.TTreeNode, s *string) {
		funcA.Call(execq.NewStack(), sender, node, s)
	}

	return &f
}

func NewTTVEditingEvent(funcA *execq.Function) *vcl.TTVEditingEvent {
	var f vcl.TTVEditingEvent = func(sender vcl.IObject, node *vcl.TTreeNode, allowEdit *bool) {
		funcA.Call(execq.NewStack(), sender, node, allowEdit)
	}

	return &f
}

func NewTTVExpandingEvent(funcA *execq.Function) *vcl.TTVExpandingEvent {
	var f vcl.TTVExpandingEvent = func(sender vcl.IObject, node *vcl.TTreeNode, allowExpansion *bool) {
		funcA.Call(execq.NewStack(), sender, node, allowExpansion)
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

	var lclExports = map[string]interface{}{
		// "NewTNotifyEvent":                      NewTNotifyEvent,
		// "NewTKeyEvent":                         NewTKeyEvent,
		"NewTNotifyEvent":                      NewTNotifyEvent,
		"NewTKeyEvent":                         NewTKeyEvent,
		"NewTKeyPressEvent":                    NewTKeyPressEvent,
		"NewTMouseEvent":                       NewTMouseEvent,
		"NewTMouseMoveEvent":                   NewTMouseMoveEvent,
		"NewTExceptionEvent":                   NewTExceptionEvent,
		"NewTCloseEvent":                       NewTCloseEvent,
		"NewTCloseQueryEvent":                  NewTCloseQueryEvent,
		"NewTContextPopupEvent":                NewTContextPopupEvent,
		"NewTDragDropEvent":                    NewTDragDropEvent,
		"NewTDragOverEvent":                    NewTDragOverEvent,
		"NewTStartDragEvent":                   NewTStartDragEvent,
		"NewTEndDragEvent":                     NewTEndDragEvent,
		"NewTAlignPositionEvent":               NewTAlignPositionEvent,
		"NewTDockDropEvent":                    NewTDockDropEvent,
		"NewTDockOverEvent":                    NewTDockOverEvent,
		"NewTStartDockEvent":                   NewTStartDockEvent,
		"NewTUnDockEvent":                      NewTUnDockEvent,
		"NewTGetSiteInfoEvent":                 NewTGetSiteInfoEvent,
		"NewTMouseWheelEvent":                  NewTMouseWheelEvent,
		"NewTMouseWheelUpDownEvent":            NewTMouseWheelUpDownEvent,
		"NewTMessageEvent":                     NewTMessageEvent,
		"NewTHelpEvent":                        NewTHelpEvent,
		"NewTWebTitleChangeEvent":              NewTWebTitleChangeEvent,
		"NewTWebJSExternalEvent":               NewTWebJSExternalEvent,
		"NewTMeasureItemEvent":                 NewTMeasureItemEvent,
		"NewTMovedEvent":                       NewTMovedEvent,
		"NewTDrawCellEvent":                    NewTDrawCellEvent,
		"NewTSelectCellEvent":                  NewTSelectCellEvent,
		"NewTGetEditEvent":                     NewTGetEditEvent,
		"NewTSetEditEvent":                     NewTSetEditEvent,
		"NewTDropFilesEvent":                   NewTDropFilesEvent,
		"NewTConstrainedResizeEvent":           NewTConstrainedResizeEvent,
		"NewTWndProcEvent":                     NewTWndProcEvent,
		"NewTSectionNotifyEvent":               NewTSectionNotifyEvent,
		"NewTSectionTrackEvent":                NewTSectionTrackEvent,
		"NewTSectionDragEvent":                 NewTSectionDragEvent,
		"NewTSysLinkEvent":                     NewTSysLinkEvent,
		"NewTDrawItemEvent":                    NewTDrawItemEvent,
		"NewTLVSelectItemEvent":                NewTLVSelectItemEvent,
		"NewTLVCheckedItemEvent":               NewTLVCheckedItemEvent,
		"NewTLVAdvancedCustomDrawEvent":        NewTLVAdvancedCustomDrawEvent,
		"NewTLVAdvancedCustomDrawItemEvent":    NewTLVAdvancedCustomDrawItemEvent,
		"NewTLVAdvancedCustomDrawSubItemEvent": NewTLVAdvancedCustomDrawSubItemEvent,
		"NewTLVChangeEvent":                    NewTLVChangeEvent,
		"NewTLVColumnClickEvent":               NewTLVColumnClickEvent,
		"NewTLVCompareEvent":                   NewTLVCompareEvent,
		"NewTLVOwnerDataEvent":                 NewTLVOwnerDataEvent,
		"NewTLVOwnerDataFindEvent":             NewTLVOwnerDataFindEvent,
		"NewTLVOwnerDataHintEvent":             NewTLVOwnerDataHintEvent,
		"NewTLVDataHintEvent":                  NewTLVDataHintEvent,
		"NewTLVDeletedEvent":                   NewTLVDeletedEvent,
		"NewTLVEditedEvent":                    NewTLVEditedEvent,
		"NewTLVEditingEvent":                   NewTLVEditingEvent,
		"NewTMenuChangeEvent":                  NewTMenuChangeEvent,
		"NewTMenuMeasureItemEvent":             NewTMenuMeasureItemEvent,
		"NewTTabChangingEvent":                 NewTTabChangingEvent,
		"NewTUDChangingEvent":                  NewTUDChangingEvent,
		"NewTUDClickEvent":                     NewTUDClickEvent,
		"NewTTaskDlgClickEvent":                NewTTaskDlgClickEvent,
		"NewTTVExpandedEvent":                  NewTTVExpandedEvent,
		"NewTTVAdvancedCustomDrawEvent":        NewTTVAdvancedCustomDrawEvent,
		"NewTTVAdvancedCustomDrawItemEvent":    NewTTVAdvancedCustomDrawItemEvent,
		"NewTTVChangedEvent":                   NewTTVChangedEvent,
		"NewTTVChangingEvent":                  NewTTVChangingEvent,
		"NewTTVCollapsingEvent":                NewTTVCollapsingEvent,
		"NewTTVCompareEvent":                   NewTTVCompareEvent,
		"NewTTVCustomDrawEvent":                NewTTVCustomDrawEvent,
		"NewTTVCustomDrawItemEvent":            NewTTVCustomDrawItemEvent,
		"NewTTVEditedEvent":                    NewTTVEditedEvent,
		"NewTTVEditingEvent":                   NewTTVEditingEvent,
		"NewTTVExpandingEvent":                 NewTTVExpandingEvent,

		"GetApplication": getVclApplication,
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
