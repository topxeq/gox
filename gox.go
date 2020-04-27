package main

import (
	"bufio"
	"bytes"
	"context"
	"io"
	"math/rand"
	"os"
	"os/exec"
	"reflect"
	"strings"
	"sync"
	"time"

	"errors"
	"fmt"

	"net"
	"strconv"

	"github.com/melbahja/goph"
	"golang.org/x/crypto/ssh"

	"github.com/atotto/clipboard"
	"github.com/d5/tengo/stdlib"
	"github.com/d5/tengo/v2"
	"github.com/dop251/goja"
	"github.com/mattn/anko/core"
	"github.com/mattn/anko/env"
	_ "github.com/mattn/anko/packages"
	"github.com/mattn/anko/parser"
	"github.com/mattn/anko/vm"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/godror/godror"
	_ "github.com/mattn/go-sqlite3"

	"github.com/topxeq/sqltk"

	"github.com/dgraph-io/badger"

	"github.com/beevik/etree"

	// GUI related start
	"github.com/sqweek/dialog"
	// GUI related end

	"github.com/topxeq/tk"

	// GUI related start
	"github.com/AllenDang/giu"
	g "github.com/AllenDang/giu"
	"github.com/AllenDang/giu/imgui"

	"github.com/topxeq/govcl/vcl"
	"github.com/topxeq/govcl/vcl/api"
	"github.com/topxeq/govcl/vcl/rtl"
	"github.com/topxeq/govcl/vcl/types"
	// GUI related end
)

// Non GUI related

var versionG = "0.93a"

var variableG = make(map[string]interface{})

var jsVMG *goja.Runtime = nil
var ankVMG *env.Env = nil
var tengoModulesG = stdlib.GetModuleMap(stdlib.AllModuleNames()...)

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

func getVar(nameA string) interface{} {
	varMutexG.Lock()
	rs, ok := variableG[nameA]
	varMutexG.Unlock()

	if !ok {
		tk.GenerateErrorString("no key")
	}
	return rs
}

func setVar(nameA string, valueA interface{}) {
	varMutexG.Lock()
	variableG[nameA] = valueA
	varMutexG.Unlock()
}

func getVarTengo(objsA ...tengo.Object) (tengo.Object, error) {
	if len(objsA) < 1 {
		return tengo.FromInterface(tk.GenerateErrorString("not enough parameters"))
	}

	strT, ok := tengo.ToString(objsA[0])

	if !ok {
		return tengo.FromInterface(tk.GenerateErrorString("failed to convert value"))
	}

	varMutexG.Lock()
	objT, ok := variableG[strT]
	varMutexG.Unlock()

	if !ok {
		return tengo.FromInterface(tk.GenerateErrorString("no key"))
	}

	return tengo.FromInterface(objT)
}

func setVarTengo(nameA string, valueA interface{}) {
	varMutexG.Lock()
	variableG[nameA] = valueA
	varMutexG.Unlock()
}

func getClipText() string {
	textT, errT := clipboard.ReadAll()
	if errT != nil {
		return tk.GenerateErrorStringF("could not get text from clipboard: %v", errT.Error())
	}

	return textT
}

func setClipText(textA string) {
	clipboard.WriteAll(textA)
}

func times(objsA ...tengo.Object) (tengo.Object, error) {
	lenT := len(objsA)

	intListT := make([]int, lenT)

	// 用一个循环将函数不定个数参数中的所有数值存入整数切片中
	for i, v := range objsA {
		// 调用objects.ToInt函数将objects.Object对象转换为整数
		cT, ok := tengo.ToInt(v)

		if ok {
			intListT[i] = cT
		}
	}

	// 进行累乘与那算
	r := 1

	for i := 0; i < lenT; i++ {
		r = r * intListT[i]
	}

	// 输出结果值供参考
	fmt.Printf("result: %v\n", r)

	// 也作为函数返回值返回，返回前要转换为objects.Object类型
	// objects.Int类型实现了objects.Object类型，因此可以用作返回值
	return &tengo.Int{Value: int64(r)}, nil
}

func eval(expA string) interface{} {
	v, errT := vm.Execute(ankVMG, nil, expA)
	if errT != nil {
		return errT.Error()
	}

	return v
}

func panicIt(valueA interface{}) {
	panic(valueA)
}

func checkError(errA error, funcA func()) {
	if errA != nil {
		tk.PlErr(errA)

		if funcA != nil {
			funcA()
		}

		os.Exit(1)
	}

}

func checkErrorString(strA string, funcA func()) {
	if tk.IsErrorString(strA) {
		tk.PlErrString(strA)

		if funcA != nil {
			funcA()
		}

		os.Exit(1)
	}

}

func newSSHClient(hostA string, portA int, userA string, passA string) (*goph.Client, error) {
	authT := goph.Password(passA)

	clientT := &goph.Client{
		Addr: hostA,
		Port: portA,
		User: userA,
		Auth: authT,
	}

	errT := goph.Conn(clientT, &ssh.ClientConfig{
		User:    clientT.User,
		Auth:    clientT.Auth,
		Timeout: 20 * time.Second,
		HostKeyCallback: func(host string, remote net.Addr, key ssh.PublicKey) error {
			return nil
			// hostFound, err := goph.CheckKnownHost(host, remote, key, "")

			// if hostFound && err != nil {
			// 	return err
			// }

			// if hostFound && err == nil {
			// 	return nil
			// }

			// return goph.AddKnownHost(host, remote, key, "")
		},
	})

	// clientT, errT := goph.NewConn(userA, hostA, authT, func(host string, remote net.Addr, key ssh.PublicKey) error {

	// 	hostFound, err := goph.CheckKnownHost(host, remote, key, "")

	// 	if hostFound && err != nil {
	// 		return err
	// 	}

	// 	if hostFound && err == nil {
	// 		return nil
	// 	}

	// 	return goph.AddKnownHost(host, remote, key, "")
	// })

	return clientT, errT
}

func remove(aryA []interface{}, startA int, endA int) []interface{} {
	if startA < 0 || startA >= len(aryA) {
		tk.Pl("Runtime error: %v", "index out of range")
		exit()
	}

	if endA < 0 || endA >= len(aryA) {
		tk.Pl("Runtime error: %v", "index out of range")
		exit()
	}

	return append(aryA[:startA], aryA[endA+1:]...)
	// if idxT == 0 {
	// 	return ayrA[idxT + 1:]
	// }

	// if idxT == len(aryA) - 1 {
	// 	return ayrA[0:len(aryA) - 1]
	// }

	// return append(aryA[:idxA], aryA[idxA+1:]...)

}

// func printValue(nameA string) {
// 	if ankVMG == nil {
// 		return
// 	}

// 	v, errT := ankVMG.Get(nameA)

// 	if errT != nil {
// 		tk.Pl("%v(%T): %v", nameA, errT, errT)
// 		return
// 	}

// 	tk.Pl("%v(%T): %v", nameA, v, v)

// }

func toStringFromRuneSlice(sliceA []rune) string {
	return string(sliceA)
}

// toInt converts all reflect.Value-s into int.
func toInt(vA interface{}) int {
	v := reflect.ValueOf(&vA)
	i, _ := tryToInt(v)
	return i
}

// tryToInt attempts to convert a value to an int.
// If it cannot (in the case of a non-numeric string, a struct, etc.)
// it returns 0 and an error.
func tryToInt(v reflect.Value) (int, error) {
	if v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	switch v.Kind() {
	case reflect.Float64, reflect.Float32:
		return int(v.Float()), nil
	case reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8, reflect.Int:
		return int(v.Int()), nil
	case reflect.Bool:
		if v.Bool() {
			return 1, nil
		}
		return 0, nil
	case reflect.String:
		s := v.String()
		var i int64
		var err error
		if strings.HasPrefix(s, "0x") {
			i, err = strconv.ParseInt(s, 16, 64)
		} else {
			i, err = strconv.ParseInt(s, 10, 64)
		}
		if err == nil {
			return int(i), nil
		}
	}
	return 0, errors.New("couldn't convert to integer")
}

func getUint64Value(v reflect.Value) uint16 {
	tk.Pl("%x", v.Interface())

	var p *uint16

	p = (v.Interface().(*uint16))

	return *p
}

func initAnkoVMInstance(vmA *env.Env) {
	if vmA == nil {
		return
	}

	vmA.Define("printfln", tk.Pl)
	vmA.Define("pl", tk.Pl)
	vmA.Define("plv", tk.Plv)
	vmA.Define("plerr", tk.PlErr)
	vmA.Define("checkError", checkError)
	vmA.Define("checkErrorString", checkErrorString)

	printValue := func(nameA string) {

		v, errT := vmA.Get(nameA)

		if errT != nil {
			tk.Pl("%v(%T): %v", nameA, errT, errT)
			return
		}

		tk.Pl("%v(%T): %v", nameA, v, v)

	}

	vmA.Define("pv", printValue)

	vmA.Define("getValue", getUint64Value)

	vmA.Define("getInput", tk.GetInputBufferedScan)

	vmA.Define("exit", exit)

	vmA.Define("toStringFromRuneSlice", toStringFromRuneSlice)
	vmA.Define("remove", remove)

	vmA.Define("eval", eval)
	vmA.Define("runScript", runScript)
	vmA.Define("systemCmd", systemCmd)
	vmA.Define("typeof", typeOfValue)
	vmA.Define("panic", panicIt)

	vmA.Define("setVar", setVar)
	vmA.Define("getVar", getVar)

	vmA.Define("setClipText", setClipText)
	vmA.Define("getClipText", getClipText)

	vmA.Define("deepCopy", tk.DeepCopyFromTo)
	vmA.Define("deepClone", tk.DeepClone)

	vmA.Define("toExactInt", toInt)

	vmA.Define("newSSHClient", newSSHClient)

	// GUI related start

	vmA.Define("edit", editFile)
	vmA.Define("run", runFile)

	// GUI related end

	core.Import(vmA)

}

func runScript(codeA string, modeA string, argsA ...string) interface{} {

	if modeA == "" || modeA == "1" || modeA == "new" {
		var vmT *env.Env

		vmT = env.NewEnv()

		initAnkoVMInstance(vmT)

		vmT.Define("argsG", argsA)

		v, errT := vm.Execute(vmT, nil, codeA)
		if errT != nil {
			posStrT := ""

			e, ok := errT.(*parser.Error)

			if ok {
				posStrT = fmt.Sprintf("line: %v, col: %v", e.Pos.Line, e.Pos.Column)
			} else {
				e, ok := errT.(*vm.Error)

				if ok {
					posStrT = fmt.Sprintf("line: %v, col: %v", e.Pos.Line, e.Pos.Column)
				} else {
					tk.Pl("%#v", errT)
				}
			}

			return tk.GenerateErrorStringF("failed to execute script(%v) error: %v", posStrT, errT)
		}

		return v
	} else if modeA == "2" || modeA == "current" {
		ankVMG.Define("argsG", argsA)

		v, errT := vm.Execute(ankVMG, nil, codeA)
		if errT != nil {
			posStrT := ""

			e, ok := errT.(*parser.Error)

			if ok {
				posStrT = fmt.Sprintf("line: %v, col: %v", e.Pos.Line, e.Pos.Column)
			} else {
				e, ok := errT.(*vm.Error)

				if ok {
					posStrT = fmt.Sprintf("line: %v, col: %v", e.Pos.Line, e.Pos.Column)
				} else {
					tk.Pl("%#v", errT)
				}
			}

			return tk.GenerateErrorStringF("failed to execute script(%v) error: %v", posStrT, errT)
		}

		return v

	} else if modeA == "3" || modeA == "js" {
		initJSVM()

		jsVMG.Set("argsG", jsVMG.ToValue(argsA))

		_, errT := jsVMG.RunString(codeA)
		if errT != nil {
			return tk.GenerateErrorStringF("failed to run script(%v): %v", codeA, errT)
		}

		result := jsVMG.Get("resultG")

		return result

	} else if modeA == "5" || modeA == "tg" {
		scriptT := tengo.NewScript([]byte(codeA))

		scriptT.SetImports(tengoModulesG)

		_ = scriptT.Add("setVar", setVarTengo)
		errT := scriptT.Add("times", times)
		if errT != nil {
			return tk.GenerateErrorStringF("failed to add times(%v) error: %v", "", errT)
		}

		errT = scriptT.Add("getVar", getVarTengo)
		if errT != nil {
			return tk.GenerateErrorStringF("failed to add getVar(%v) error: %v", "", errT)
		}

		argsG_TG := make([]interface{}, len(argsA))
		for i, v := range argsA {
			argsG_TG[i] = v
		}

		errT = scriptT.Add("argsG", argsG_TG)
		if errT != nil {
			return tk.GenerateErrorStringF("failed to add argsA(%v) error: %v", "", errT)
		}

		compiledT, errT := scriptT.RunContext(context.Background())
		if errT != nil {
			return tk.GenerateErrorStringF("failed to execute script(%v) error: %v", "", errT)
		}

		result := compiledT.Get("resultG")
		return result

	} else {
		return systemCmd("gox", append([]string{codeA}, argsA...)...)
	}

}

func systemCmd(cmdA string, argsA ...string) string {
	var out bytes.Buffer

	cmd := exec.Command(cmdA, argsA...)

	cmd.Stdout = &out
	errT := cmd.Run()
	if errT != nil {
		return tk.GenerateErrorStringF("failed: %v", errT)
	}

	rStrT := tk.Trim(out.String())

	return rStrT
}

func typeOfValue(vA interface{}) string {
	return fmt.Sprintf("%T", vA)
}

func importAnkNonGUIPackages() {

	env.Packages["etree"] = map[string]reflect.Value{
		"NewDocument": reflect.ValueOf(etree.NewDocument),
	}

	env.PackageTypes["badger"] = map[string]reflect.Type{
		"IteratorOptions": reflect.TypeOf(badger.IteratorOptions{}),
		// "IteratorOptions": reflect.TypeOf(&widget).Elem(),
	}

	env.Packages["sqltk"] = map[string]reflect.Value{
		"ConnectDB":          reflect.ValueOf(sqltk.ConnectDB),
		"ConnectDBNoPing":    reflect.ValueOf(sqltk.ConnectDBNoPing),
		"ExecV":              reflect.ValueOf(sqltk.ExecV),
		"QueryDBS":           reflect.ValueOf(sqltk.QueryDBS),
		"QueryDBNS":          reflect.ValueOf(sqltk.QueryDBNS),
		"QueryDBNSS":         reflect.ValueOf(sqltk.QueryDBNSS),
		"QueryDBI":           reflect.ValueOf(sqltk.QueryDBI),
		"QueryDBCount":       reflect.ValueOf(sqltk.QueryDBCount),
		"QueryDBString":      reflect.ValueOf(sqltk.QueryDBString),
		"OneLineRecordToMap": reflect.ValueOf(sqltk.OneLineRecordToMap),
	}

	env.Packages["tk"] = map[string]reflect.Value{
		"CreateTXCollection":                  reflect.ValueOf(tk.CreateTXCollection),
		"TXResultFromString":                  reflect.ValueOf(tk.TXResultFromString),
		"SetGlobalEnv":                        reflect.ValueOf(tk.SetGlobalEnv),
		"RemoveGlobalEnv":                     reflect.ValueOf(tk.RemoveGlobalEnv),
		"GetGlobalEnvList":                    reflect.ValueOf(tk.GetGlobalEnvList),
		"GetGlobalEnvString":                  reflect.ValueOf(tk.GetGlobalEnvString),
		"HasGlobalEnv":                        reflect.ValueOf(tk.HasGlobalEnv),
		"IsEmptyTrim":                         reflect.ValueOf(tk.IsEmptyTrim),
		"StartsWith":                          reflect.ValueOf(tk.StartsWith),
		"StartsWithIgnoreCase":                reflect.ValueOf(tk.StartsWithIgnoreCase),
		"StartsWithUpper":                     reflect.ValueOf(tk.StartsWithUpper),
		"StartsWithDigit":                     reflect.ValueOf(tk.StartsWithDigit),
		"Contains":                            reflect.ValueOf(tk.Contains),
		"ContainsIgnoreCase":                  reflect.ValueOf(tk.ContainsIgnoreCase),
		"EndsWith":                            reflect.ValueOf(tk.EndsWith),
		"EndsWithIgnoreCase":                  reflect.ValueOf(tk.EndsWithIgnoreCase),
		"Trim":                                reflect.ValueOf(tk.Trim),
		"TrimCharSet":                         reflect.ValueOf(tk.TrimCharSet),
		"InStrings":                           reflect.ValueOf(tk.InStrings),
		"GetSliceMaxLen":                      reflect.ValueOf(tk.GetSliceMaxLen),
		"FindFirstDiffIndex":                  reflect.ValueOf(tk.FindFirstDiffIndex),
		"FindSamePrefix":                      reflect.ValueOf(tk.FindSamePrefix),
		"IsErrorString":                       reflect.ValueOf(tk.IsErrorString),
		"GetErrorString":                      reflect.ValueOf(tk.GetErrorString),
		"GetErrorStringSafely":                reflect.ValueOf(tk.GetErrorStringSafely),
		"GenerateErrorString":                 reflect.ValueOf(tk.GenerateErrorString),
		"GenerateErrorStringF":                reflect.ValueOf(tk.GenerateErrorStringF),
		"ErrorStringToError":                  reflect.ValueOf(tk.ErrorStringToError),
		"Replace":                             reflect.ValueOf(tk.Replace),
		"StringReplace":                       reflect.ValueOf(tk.StringReplace),
		"SplitLines":                          reflect.ValueOf(tk.SplitLines),
		"SplitLinesRemoveEmpty":               reflect.ValueOf(tk.SplitLinesRemoveEmpty),
		"Split":                               reflect.ValueOf(tk.Split),
		"SplitN":                              reflect.ValueOf(tk.SplitN),
		"JoinLines":                           reflect.ValueOf(tk.JoinLines),
		"JoinLinesBySeparator":                reflect.ValueOf(tk.JoinLinesBySeparator),
		"EnsureValidFileNameX":                reflect.ValueOf(tk.EnsureValidFileNameX),
		"CreateString":                        reflect.ValueOf(tk.CreateString),
		"CreateStringSimple":                  reflect.ValueOf(tk.CreateStringSimple),
		"CreateStringWithObject":              reflect.ValueOf(tk.CreateStringWithObject),
		"CreateStringEmpty":                   reflect.ValueOf(tk.CreateStringEmpty),
		"CreateStringSuccess":                 reflect.ValueOf(tk.CreateStringSuccess),
		"CreateStringError":                   reflect.ValueOf(tk.CreateStringError),
		"CreateStringErrorF":                  reflect.ValueOf(tk.CreateStringErrorF),
		"CreateStringErrorFromTXError":        reflect.ValueOf(tk.CreateStringErrorFromTXError),
		"GenerateErrorStringTX":               reflect.ValueOf(tk.GenerateErrorStringTX),
		"GenerateErrorStringFTX":              reflect.ValueOf(tk.GenerateErrorStringFTX),
		"LoadStringTX":                        reflect.ValueOf(tk.LoadStringTX),
		"RegContains":                         reflect.ValueOf(tk.RegContains),
		"RegFindFirstTX":                      reflect.ValueOf(tk.RegFindFirstTX),
		"LoadDualLineListFromString":          reflect.ValueOf(tk.LoadDualLineListFromString),
		"RegReplace":                          reflect.ValueOf(tk.RegReplace),
		"RegFindAll":                          reflect.ValueOf(tk.RegFindAll),
		"RegFindFirst":                        reflect.ValueOf(tk.RegFindFirst),
		"RegMatch":                            reflect.ValueOf(tk.RegMatch),
		"Randomize":                           reflect.ValueOf(tk.Randomize),
		"GetRandomIntLessThan":                reflect.ValueOf(tk.GetRandomIntLessThan),
		"GetRandomInt64LessThan":              reflect.ValueOf(tk.GetRandomInt64LessThan),
		"GetRandomIntInRange":                 reflect.ValueOf(tk.GetRandomIntInRange),
		"GetRandomInt64InRange":               reflect.ValueOf(tk.GetRandomInt64InRange),
		"GenerateRandomString":                reflect.ValueOf(tk.GenerateRandomString),
		"NewRandomGenerator":                  reflect.ValueOf(tk.NewRandomGenerator),
		"ShuffleStringArray":                  reflect.ValueOf(tk.ShuffleStringArray),
		"GetRandomizeStringArrayCopy":         reflect.ValueOf(tk.GetRandomizeStringArrayCopy),
		"GetRandomizeIntArrayCopy":            reflect.ValueOf(tk.GetRandomizeIntArrayCopy),
		"GetRandomizeInt64ArrayCopy":          reflect.ValueOf(tk.GetRandomizeInt64ArrayCopy),
		"GetRandomSubDualList":                reflect.ValueOf(tk.GetRandomSubDualList),
		"JoinDualList":                        reflect.ValueOf(tk.JoinDualList),
		"GetNowDateString":                    reflect.ValueOf(tk.GetNowDateString),
		"GetNowTimeString":                    reflect.ValueOf(tk.GetNowTimeString),
		"GetNowTimeStringFormat":              reflect.ValueOf(tk.GetNowTimeStringFormat),
		"GetNowTimeStringFormal":              reflect.ValueOf(tk.GetNowTimeStringFormal),
		"GetNowTimeOnlyStringBeijing":         reflect.ValueOf(tk.GetNowTimeOnlyStringBeijing),
		"GetTimeFromUnixTimeStamp":            reflect.ValueOf(tk.GetTimeFromUnixTimeStamp),
		"GetTimeFromUnixTimeStampMid":         reflect.ValueOf(tk.GetTimeFromUnixTimeStampMid),
		"GetTimeStamp":                        reflect.ValueOf(tk.GetTimeStamp),
		"GetTimeStampMid":                     reflect.ValueOf(tk.GetTimeStampMid),
		"GetTimeStampNano":                    reflect.ValueOf(tk.GetTimeStampNano),
		"NowToFileName":                       reflect.ValueOf(tk.NowToFileName),
		"GetNowTimeStringHourMinute":          reflect.ValueOf(tk.GetNowTimeStringHourMinute),
		"GetNowMinutesInDay":                  reflect.ValueOf(tk.GetNowMinutesInDay),
		"NowToStrUTC":                         reflect.ValueOf(tk.NowToStrUTC),
		"GetTimeStringDiffMS":                 reflect.ValueOf(tk.GetTimeStringDiffMS),
		"StrToTime":                           reflect.ValueOf(tk.StrToTime),
		"StrToTimeByFormat":                   reflect.ValueOf(tk.StrToTimeByFormat),
		"FormatTime":                          reflect.ValueOf(tk.FormatTime),
		"IsYesterday":                         reflect.ValueOf(tk.IsYesterday),
		"DeleteItemInStringArray":             reflect.ValueOf(tk.DeleteItemInStringArray),
		"DeleteItemInIntArray":                reflect.ValueOf(tk.DeleteItemInIntArray),
		"DeleteItemInInt64Array":              reflect.ValueOf(tk.DeleteItemInInt64Array),
		"ContainsIn":                          reflect.ValueOf(tk.ContainsIn),
		"ContainsInStringList":                reflect.ValueOf(tk.ContainsInStringList),
		"IndexInStringList":                   reflect.ValueOf(tk.IndexInStringList),
		"IndexInStringListFromEnd":            reflect.ValueOf(tk.IndexInStringListFromEnd),
		"GetStringSliceFilled":                reflect.ValueOf(tk.GetStringSliceFilled),
		"Len64":                               reflect.ValueOf(tk.Len64),
		"Int64ArrayToFloat64Array":            reflect.ValueOf(tk.Int64ArrayToFloat64Array),
		"ByteSliceToStringDec":                reflect.ValueOf(tk.ByteSliceToStringDec),
		"GetValueOfMSS":                       reflect.ValueOf(tk.GetValueOfMSS),
		"Prf":                                 reflect.ValueOf(tk.Prf),
		"Prl":                                 reflect.ValueOf(tk.Prl),
		"Printf":                              reflect.ValueOf(tk.Printf),
		"Printfln":                            reflect.ValueOf(tk.Printfln),
		"Spr":                                 reflect.ValueOf(tk.Spr),
		"Pr":                                  reflect.ValueOf(tk.Pr),
		"Pl":                                  reflect.ValueOf(tk.Pl),
		"PlVerbose":                           reflect.ValueOf(tk.PlVerbose),
		"Fpl":                                 reflect.ValueOf(tk.Fpl),
		"Fpr":                                 reflect.ValueOf(tk.Fpr),
		"PlvWithError":                        reflect.ValueOf(tk.PlvWithError),
		"PlAndExit":                           reflect.ValueOf(tk.PlAndExit),
		"PlErrSimple":                         reflect.ValueOf(tk.PlErrSimple),
		"PlErrSimpleAndExit":                  reflect.ValueOf(tk.PlErrSimpleAndExit),
		"PlErrAndExit":                        reflect.ValueOf(tk.PlErrAndExit),
		"PlTXErr":                             reflect.ValueOf(tk.PlTXErr),
		"PlSimpleErrorString":                 reflect.ValueOf(tk.PlSimpleErrorString),
		"PlErr":                               reflect.ValueOf(tk.PlErr),
		"PlErrWithPrefix":                     reflect.ValueOf(tk.PlErrWithPrefix),
		"Plv":                                 reflect.ValueOf(tk.Plv),
		"Plvs":                                reflect.ValueOf(tk.Plvs),
		"Plvsr":                               reflect.ValueOf(tk.Plvsr),
		"Errf":                                reflect.ValueOf(tk.Errf),
		"FatalErr":                            reflect.ValueOf(tk.FatalErr),
		"FatalErrf":                           reflect.ValueOf(tk.FatalErrf),
		"Fatalf":                              reflect.ValueOf(tk.Fatalf),
		"CheckErr":                            reflect.ValueOf(tk.CheckErr),
		"CheckErrf":                           reflect.ValueOf(tk.CheckErrf),
		"CheckErrCompact":                     reflect.ValueOf(tk.CheckErrCompact),
		"GetUserInput":                        reflect.ValueOf(tk.GetUserInput),
		"GetInputBufferedScan":                reflect.ValueOf(tk.GetInputBufferedScan),
		"SleepSeconds":                        reflect.ValueOf(tk.SleepSeconds),
		"SleepMilliSeconds":                   reflect.ValueOf(tk.SleepMilliSeconds),
		"GetRuntimeStack":                     reflect.ValueOf(tk.GetRuntimeStack),
		"GetOSName":                           reflect.ValueOf(tk.GetOSName),
		"GetCurrentDir":                       reflect.ValueOf(tk.GetCurrentDir),
		"GetApplicationPath":                  reflect.ValueOf(tk.GetApplicationPath),
		"EnsureMakeDirs":                      reflect.ValueOf(tk.EnsureMakeDirs),
		"EnsureMakeDirsE":                     reflect.ValueOf(tk.EnsureMakeDirsE),
		"AnalyzeCommandLineParamter":          reflect.ValueOf(tk.AnalyzeCommandLineParamter),
		"GetParameterByIndexWithDefaultValue": reflect.ValueOf(tk.GetParameterByIndexWithDefaultValue),
		"ParseCommandLine":                    reflect.ValueOf(tk.ParseCommandLine),
		"GetSwitchWithDefaultValue":           reflect.ValueOf(tk.GetSwitchWithDefaultValue),
		"GetSwitchWithDefaultIntValue":        reflect.ValueOf(tk.GetSwitchWithDefaultIntValue),
		"GetSwitchWithDefaultInt64Value":      reflect.ValueOf(tk.GetSwitchWithDefaultInt64Value),
		"IfSwitchExists":                      reflect.ValueOf(tk.IfSwitchExists),
		"IfSwitchExistsWhole":                 reflect.ValueOf(tk.IfSwitchExistsWhole),
		"StrToBool":                           reflect.ValueOf(tk.StrToBool),
		"ByteToHex":                           reflect.ValueOf(tk.ByteToHex),
		"IntToStr":                            reflect.ValueOf(tk.IntToStr),
		"Int64ToStr":                          reflect.ValueOf(tk.Int64ToStr),
		"StrToIntWithDefaultValue":            reflect.ValueOf(tk.StrToIntWithDefaultValue),
		"StrToInt":                            reflect.ValueOf(tk.StrToInt),
		"StrToInt64WithDefaultValue":          reflect.ValueOf(tk.StrToInt64WithDefaultValue),
		"StrToIntPositive":                    reflect.ValueOf(tk.StrToIntPositive),
		"StrToFloat64WithDefaultValue":        reflect.ValueOf(tk.StrToFloat64WithDefaultValue),
		"StrToFloat64":                        reflect.ValueOf(tk.StrToFloat64),
		"Float64ToStr":                        reflect.ValueOf(tk.Float64ToStr),
		"StrToTimeCompact":                    reflect.ValueOf(tk.StrToTimeCompact),
		"StrToTimeCompactNoError":             reflect.ValueOf(tk.StrToTimeCompactNoError),
		"FormatStringSliceSlice":              reflect.ValueOf(tk.FormatStringSliceSlice),
		"IntToKMGT":                           reflect.ValueOf(tk.IntToKMGT),
		"IntToWYZ":                            reflect.ValueOf(tk.IntToWYZ),
		"SetLogFile":                          reflect.ValueOf(tk.SetLogFile),
		"LogWithTime":                         reflect.ValueOf(tk.LogWithTime),
		"LogWithTimeCompact":                  reflect.ValueOf(tk.LogWithTimeCompact),
		"IfFileExists":                        reflect.ValueOf(tk.IfFileExists),
		"IsFile":                              reflect.ValueOf(tk.IsFile),
		"IsDirectory":                         reflect.ValueOf(tk.IsDirectory),
		"GetFilePathSeperator":                reflect.ValueOf(tk.GetFilePathSeperator),
		"GetLastComponentOfFilePath":          reflect.ValueOf(tk.GetLastComponentOfFilePath),
		"GetDirOfFilePath":                    reflect.ValueOf(tk.GetDirOfFilePath),
		"RemoveFileExt":                       reflect.ValueOf(tk.RemoveFileExt),
		"GetFileExt":                          reflect.ValueOf(tk.GetFileExt),
		"RemoveLastSubString":                 reflect.ValueOf(tk.RemoveLastSubString),
		"AddLastSubString":                    reflect.ValueOf(tk.AddLastSubString),
		"GenerateFileListRecursively":         reflect.ValueOf(tk.GenerateFileListRecursively),
		"GetAvailableFileName":                reflect.ValueOf(tk.GetAvailableFileName),
		"LoadStringFromFile":                  reflect.ValueOf(tk.LoadStringFromFile),
		"LoadStringFromFileWithDefault":       reflect.ValueOf(tk.LoadStringFromFileWithDefault),
		"LoadStringFromFileE":                 reflect.ValueOf(tk.LoadStringFromFileE),
		"LoadStringFromFileB":                 reflect.ValueOf(tk.LoadStringFromFileB),
		"LoadBytes":                           reflect.ValueOf(tk.LoadBytes),
		"SaveStringToFile":                    reflect.ValueOf(tk.SaveStringToFile),
		"SaveStringToFileE":                   reflect.ValueOf(tk.SaveStringToFileE),
		"AppendStringToFile":                  reflect.ValueOf(tk.AppendStringToFile),
		"LoadStringList":                      reflect.ValueOf(tk.LoadStringList),
		"LoadStringListFromFile":              reflect.ValueOf(tk.LoadStringListFromFile),
		"LoadStringListBuffered":              reflect.ValueOf(tk.LoadStringListBuffered),
		"SaveStringList":                      reflect.ValueOf(tk.SaveStringList),
		"SaveStringListWin":                   reflect.ValueOf(tk.SaveStringListWin),
		"SaveStringListBufferedByRange":       reflect.ValueOf(tk.SaveStringListBufferedByRange),
		"SaveStringListBuffered":              reflect.ValueOf(tk.SaveStringListBuffered),
		"ReadLineFromBufioReader":             reflect.ValueOf(tk.ReadLineFromBufioReader),
		"RestoreLineEnds":                     reflect.ValueOf(tk.RestoreLineEnds),
		"LoadDualLineList":                    reflect.ValueOf(tk.LoadDualLineList),
		"SaveDualLineList":                    reflect.ValueOf(tk.SaveDualLineList),
		"RemoveDuplicateInDualLineList":       reflect.ValueOf(tk.RemoveDuplicateInDualLineList),
		"AppendDualLineList":                  reflect.ValueOf(tk.AppendDualLineList),
		"LoadSimpleMapFromFile":               reflect.ValueOf(tk.LoadSimpleMapFromFile),
		"LoadSimpleMapFromFileE":              reflect.ValueOf(tk.LoadSimpleMapFromFileE),
		"SimpleMapToString":                   reflect.ValueOf(tk.SimpleMapToString),
		"LoadSimpleMapFromString":             reflect.ValueOf(tk.LoadSimpleMapFromString),
		"LoadSimpleMapFromStringE":            reflect.ValueOf(tk.LoadSimpleMapFromStringE),
		"ReplaceLineEnds":                     reflect.ValueOf(tk.ReplaceLineEnds),
		"SaveSimpleMapToFile":                 reflect.ValueOf(tk.SaveSimpleMapToFile),
		"AppendSimpleMapFromFile":             reflect.ValueOf(tk.AppendSimpleMapFromFile),
		"LoadSimpleMapFromDir":                reflect.ValueOf(tk.LoadSimpleMapFromDir),
		"EncodeToXMLString":                   reflect.ValueOf(tk.EncodeToXMLString),
		"ObjectToJSON":                        reflect.ValueOf(tk.ObjectToJSON),
		"ObjectToJSONIndent":                  reflect.ValueOf(tk.ObjectToJSONIndent),
		"JSONToMapStringString":               reflect.ValueOf(tk.JSONToMapStringString),
		"JSONToObject":                        reflect.ValueOf(tk.JSONToObject),
		"SafelyGetStringForKeyWithDefault":    reflect.ValueOf(tk.SafelyGetStringForKeyWithDefault),
		"SafelyGetFloat64ForKeyWithDefault":   reflect.ValueOf(tk.SafelyGetFloat64ForKeyWithDefault),
		"SafelyGetIntForKeyWithDefault":       reflect.ValueOf(tk.SafelyGetIntForKeyWithDefault),
		"JSONToStringArray":                   reflect.ValueOf(tk.JSONToStringArray),
		"EncodeStringSimple":                  reflect.ValueOf(tk.EncodeStringSimple),
		"EncodeStringUnderline":               reflect.ValueOf(tk.EncodeStringUnderline),
		"EncodeStringCustom":                  reflect.ValueOf(tk.EncodeStringCustom),
		"DecodeStringSimple":                  reflect.ValueOf(tk.DecodeStringSimple),
		"DecodeStringUnderline":               reflect.ValueOf(tk.DecodeStringUnderline),
		"DecodeStringCustom":                  reflect.ValueOf(tk.DecodeStringCustom),
		"MD5Encrypt":                          reflect.ValueOf(tk.MD5Encrypt),
		"BytesToHex":                          reflect.ValueOf(tk.BytesToHex),
		"HexToBytes":                          reflect.ValueOf(tk.HexToBytes),
		"GetRandomByte":                       reflect.ValueOf(tk.GetRandomByte),
		"EncryptDataByTXDEE":                  reflect.ValueOf(tk.EncryptDataByTXDEE),
		"SumBytes":                            reflect.ValueOf(tk.SumBytes),
		"EncryptDataByTXDEF":                  reflect.ValueOf(tk.EncryptDataByTXDEF),
		"EncryptStreamByTXDEF":                reflect.ValueOf(tk.EncryptStreamByTXDEF),
		"DecryptStreamByTXDEF":                reflect.ValueOf(tk.DecryptStreamByTXDEF),
		"DecryptDataByTXDEE":                  reflect.ValueOf(tk.DecryptDataByTXDEE),
		"DecryptDataByTXDEF":                  reflect.ValueOf(tk.DecryptDataByTXDEF),
		"EncryptStringByTXTE":                 reflect.ValueOf(tk.EncryptStringByTXTE),
		"DecryptStringByTXTE":                 reflect.ValueOf(tk.DecryptStringByTXTE),
		"EncryptStringByTXDEE":                reflect.ValueOf(tk.EncryptStringByTXDEE),
		"DecryptStringByTXDEE":                reflect.ValueOf(tk.DecryptStringByTXDEE),
		"EncryptStringByTXDEF":                reflect.ValueOf(tk.EncryptStringByTXDEF),
		"DecryptStringByTXDEF":                reflect.ValueOf(tk.DecryptStringByTXDEF),
		"EncryptFileByTXDEF":                  reflect.ValueOf(tk.EncryptFileByTXDEF),
		"EncryptFileByTXDEFStream":            reflect.ValueOf(tk.EncryptFileByTXDEFStream),
		"DecryptFileByTXDEFStream":            reflect.ValueOf(tk.DecryptFileByTXDEFStream),
		"ErrorToString":                       reflect.ValueOf(tk.ErrorToString),
		"EncryptFileByTXDEFS":                 reflect.ValueOf(tk.EncryptFileByTXDEFS),
		"EncryptFileByTXDEFStreamS":           reflect.ValueOf(tk.EncryptFileByTXDEFStreamS),
		"DecryptFileByTXDEF":                  reflect.ValueOf(tk.DecryptFileByTXDEF),
		"DecryptFileByTXDEFS":                 reflect.ValueOf(tk.DecryptFileByTXDEFS),
		"DecryptFileByTXDEFStreamS":           reflect.ValueOf(tk.DecryptFileByTXDEFStreamS),
		"Pkcs7Padding":                        reflect.ValueOf(tk.Pkcs7Padding),
		"AESEncrypt":                          reflect.ValueOf(tk.AESEncrypt),
		"AESDecrypt":                          reflect.ValueOf(tk.AESDecrypt),
		"AnalyzeURLParams":                    reflect.ValueOf(tk.AnalyzeURLParams),
		"UrlEncode":                           reflect.ValueOf(tk.UrlEncode),
		"UrlEncode2":                          reflect.ValueOf(tk.UrlEncode2),
		"UrlDecode":                           reflect.ValueOf(tk.UrlDecode),
		"JoinURL":                             reflect.ValueOf(tk.JoinURL),
		"AddDebug":                            reflect.ValueOf(tk.AddDebug),
		"AddDebugF":                           reflect.ValueOf(tk.AddDebugF),
		"ClearDebug":                          reflect.ValueOf(tk.ClearDebug),
		"GetDebug":                            reflect.ValueOf(tk.GetDebug),
		"DownloadPageUTF8":                    reflect.ValueOf(tk.DownloadPageUTF8),
		"DownloadPage":                        reflect.ValueOf(tk.DownloadPage),
		"DownloadPageByMap":                   reflect.ValueOf(tk.DownloadPageByMap),
		"GetLastComponentOfUrl":               reflect.ValueOf(tk.GetLastComponentOfUrl),
		"DownloadFile":                        reflect.ValueOf(tk.DownloadFile),
		"DownloadBytes":                       reflect.ValueOf(tk.DownloadBytes),
		"PostRequest":                         reflect.ValueOf(tk.PostRequest),
		"PostRequestX":                        reflect.ValueOf(tk.PostRequestX),
		"PostRequestBytesX":                   reflect.ValueOf(tk.PostRequestBytesX),
		"PostRequestBytesWithMSSHeaderX":      reflect.ValueOf(tk.PostRequestBytesWithMSSHeaderX),
		"PostRequestBytesWithCookieX":         reflect.ValueOf(tk.PostRequestBytesWithCookieX),
		"GetFormValueWithDefaultValue":        reflect.ValueOf(tk.GetFormValueWithDefaultValue),
		"GenerateJSONPResponse":               reflect.ValueOf(tk.GenerateJSONPResponse),
		"GenerateJSONPResponseWithObject":     reflect.ValueOf(tk.GenerateJSONPResponseWithObject),
		"GenerateJSONPResponseWith2Object":    reflect.ValueOf(tk.GenerateJSONPResponseWith2Object),
		"GenerateJSONPResponseWith3Object":    reflect.ValueOf(tk.GenerateJSONPResponseWith3Object),
		"GetSuccessValue":                     reflect.ValueOf(tk.GetSuccessValue),
		"Float32ArrayToFloat64Array":          reflect.ValueOf(tk.Float32ArrayToFloat64Array),
		"CalCosineSimilarityBetweenFloatsBig": reflect.ValueOf(tk.CalCosineSimilarityBetweenFloatsBig),
		"GetDBConnection":                     reflect.ValueOf(tk.GetDBConnection),
		"GetDBRowCount":                       reflect.ValueOf(tk.GetDBRowCount),
		"GetDBRowCountCompact":                reflect.ValueOf(tk.GetDBRowCountCompact),
		"GetDBResultString":                   reflect.ValueOf(tk.GetDBResultString),
		"GetDBResultArray":                    reflect.ValueOf(tk.GetDBResultArray),
		"ConvertToGB18030":                    reflect.ValueOf(tk.ConvertToGB18030),
		"ConvertToGB18030Bytes":               reflect.ValueOf(tk.ConvertToGB18030Bytes),
		"ConvertToUTF8":                       reflect.ValueOf(tk.ConvertToUTF8),
		"ConvertStringToUTF8":                 reflect.ValueOf(tk.ConvertStringToUTF8),
		"CreateSimpleEvent":                   reflect.ValueOf(tk.CreateSimpleEvent),
		"GetAllParameters":                    reflect.ValueOf(tk.GetAllParameters),
		"GetAllSwitches":                      reflect.ValueOf(tk.GetAllSwitches),
		"ToLower":                             reflect.ValueOf(tk.ToLower),
		"ToUpper":                             reflect.ValueOf(tk.ToUpper),
		"GetEnv":                              reflect.ValueOf(tk.GetEnv),
		"JoinPath":                            reflect.ValueOf(tk.JoinPath),
		"DeepClone":                           reflect.ValueOf(tk.DeepClone),
		"DeepCopyFromTo":                      reflect.ValueOf(tk.DeepCopyFromTo),
		"JSONToObjectE":                       reflect.ValueOf(tk.JSONToObjectE),
		"ToJSON":                              reflect.ValueOf(tk.ToJSON),
		"ToJSONIndent":                        reflect.ValueOf(tk.ToJSONIndent),
		"FromJSON":                            reflect.ValueOf(tk.FromJSON),
		"GetJSONNode":                         reflect.ValueOf(tk.GetJSONNode),
		"GetJSONNodeAny":                      reflect.ValueOf(tk.GetJSONNodeAny),
		"GetJSONSubNode":                      reflect.ValueOf(tk.GetJSONSubNode),
		"GetJSONSubNodeAny":                   reflect.ValueOf(tk.GetJSONSubNodeAny),
		"StartsWithBOM":                       reflect.ValueOf(tk.StartsWithBOM),
		"RemoveBOM":                           reflect.ValueOf(tk.RemoveBOM),
		"HexToInt":                            reflect.ValueOf(tk.HexToInt),
	}

}

func showHelp() {
	tk.Pl("Gox by TopXeQ V%v\n", versionG)

	tk.Pl("Usage: gox [-v|-h] test.gox next.js, ...\n")
	tk.Pl("or just gox without arguments to start REPL instead.\n")

}

func runInteractive() int {
	var following bool
	var source string
	scanner := bufio.NewScanner(os.Stdin)

	parser.EnableErrorVerbose()

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

		stmts, err := parser.ParseSrc(source)

		if e, ok := err.(*parser.Error); ok {
			es := e.Error()
			if strings.HasPrefix(es, "syntax error: unexpected") {
				if strings.HasPrefix(es, "syntax error: unexpected $end,") {
					following = true
					continue
				}
			} else {
				if e.Pos.Column == len(source) && !e.Fatal {
					fmt.Fprintln(os.Stderr, e)
					following = true
					continue
				}
				if e.Error() == "unexpected EOF" {
					following = true
					continue
				}
			}
		}

		following = false
		source = ""
		var v interface{}

		if err == nil {
			v, err = vm.Run(ankVMG, nil, stmts)
		}
		if err != nil {
			if ankVMG, ok := err.(*vm.Error); ok {
				fmt.Fprintf(os.Stderr, "%d:%d %s\n", ankVMG.Pos.Line, ankVMG.Pos.Column, err)
			} else if ankVMG, ok := err.(*parser.Error); ok {
				fmt.Fprintf(os.Stderr, "%d:%d %s\n", ankVMG.Pos.Line, ankVMG.Pos.Column, err)
			} else {
				fmt.Fprintln(os.Stderr, err)
			}
			continue
		}

		fmt.Printf("%#v\n", v)
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

func loadFont() {
	fonts := giu.Context.IO().Fonts()

	rangeVarT := getVar("FontRange")

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

	fontVarT := getVar("Font") // "c:/Windows/Fonts/simsun.ttc"

	if fontVarT != nil {
		fontPath = fontVarT.(string)
	}

	fontSizeStrT := "16"

	fontSizeVarT := getVar("FontSize")

	if fontSizeVarT != nil {
		fontSizeStrT = fontSizeVarT.(string)
	}

	fontSizeT := tk.StrToIntWithDefaultValue(fontSizeStrT, 16)

	// fonts.AddFontFromFileTTF(fontPath, 14)
	fonts.AddFontFromFileTTFV(fontPath, float32(fontSizeT), imgui.DefaultFontConfig, ranges.Data())
}

func initLCLLib() (result error) {
	defer func() {
		if r := recover(); r != nil {
			// tk.Printfln("发生异常，错误信息：%v", r)

			result = tk.Errf("发生异常，错误信息：%v", r)

			return
		}
	}()

	api.DoLibInit()

	result = nil

	return result
}

func initLCL() {

	errT := initLCLLib()

	if errT != nil {
		tk.Pl("failed to init lib: %v, try to download the LCL lib...", errT)

		applicationPathT := tk.GetApplicationPath()

		osT := tk.GetOSName()

		if tk.Contains(osT, "inux") {
			rs := tk.DownloadFile("http://scripts.frenchfriend.net/pub/liblcl.so", applicationPathT, "liblcl.so", false)

			if tk.IsErrorString(rs) {
				tk.Pl("failed to download LCL file.")
				return
			}
		} else if tk.Contains(osT, "arwin") {
			rs := tk.DownloadFile("http://scripts.frenchfriend.net/pub/liblcl.dylib", applicationPathT, "liblcl.dylib", false)

			if tk.IsErrorString(rs) {
				tk.Pl("failed to download LCL file.")
				return
			}
		} else {
			rs := tk.DownloadFile("http://scripts.frenchfriend.net/pub/liblcl.dll", applicationPathT, "liblcl.dll", false)

			if tk.IsErrorString(rs) {
				tk.Pl("failed to download LCL file.")
				return
			}
		}

		errT = initLCLLib()

		if errT != nil {
			tk.Pl("failed to install lib: %v", errT)
			return
		}
	}

	api.DoResInit()

	api.DoImportInit()

	api.DoDefInit()

	// api.DoStyleInit()

	rtl.DoRtlInit()

	vcl.DoInit()
}

func getVclApplication() *vcl.TApplication {
	return vcl.Application
}

func importAnkGUIPackages() {
	env.Packages["gui"] = map[string]reflect.Value{
		"NewMasterWindow":         reflect.ValueOf(g.NewMasterWindow),
		"SingleWindow":            reflect.ValueOf(g.SingleWindow),
		"Window":                  reflect.ValueOf(g.Window),
		"SingleWindowWithMenuBar": reflect.ValueOf(g.SingleWindowWithMenuBar),
		"WindowV":                 reflect.ValueOf(g.WindowV),

		"MasterWindowFlagsNotResizable": reflect.ValueOf(g.MasterWindowFlagsNotResizable),
		"MasterWindowFlagsMaximized":    reflect.ValueOf(g.MasterWindowFlagsMaximized),
		"MasterWindowFlagsFloating":     reflect.ValueOf(g.MasterWindowFlagsFloating),

		// "Layout":          reflect.ValueOf(g.Layout),

		"Label":                  reflect.ValueOf(g.Label),
		"Line":                   reflect.ValueOf(g.Line),
		"Button":                 reflect.ValueOf(g.Button),
		"InvisibleButton":        reflect.ValueOf(g.InvisibleButton),
		"ImageButton":            reflect.ValueOf(g.ImageButton),
		"InputTextMultiline":     reflect.ValueOf(g.InputTextMultiline),
		"Checkbox":               reflect.ValueOf(g.Checkbox),
		"RadioButton":            reflect.ValueOf(g.RadioButton),
		"Child":                  reflect.ValueOf(g.Child),
		"ComboCustom":            reflect.ValueOf(g.ComboCustom),
		"Combo":                  reflect.ValueOf(g.Combo),
		"ContextMenu":            reflect.ValueOf(g.ContextMenu),
		"Group":                  reflect.ValueOf(g.Group),
		"Image":                  reflect.ValueOf(g.Image),
		"InputText":              reflect.ValueOf(g.InputText),
		"InputTextV":             reflect.ValueOf(g.InputTextV),
		"InputTextFlagsPassword": reflect.ValueOf(g.InputTextFlagsPassword),
		"InputInt":               reflect.ValueOf(g.InputInt),
		"InputFloat":             reflect.ValueOf(g.InputFloat),
		"MainMenuBar":            reflect.ValueOf(g.MainMenuBar),
		"MenuBar":                reflect.ValueOf(g.MenuBar),
		"MenuItem":               reflect.ValueOf(g.MenuItem),
		"PopupModal":             reflect.ValueOf(g.PopupModal),
		"OpenPopup":              reflect.ValueOf(g.OpenPopup),
		"CloseCurrentPopup":      reflect.ValueOf(g.CloseCurrentPopup),
		"ProgressBar":            reflect.ValueOf(g.ProgressBar),
		"Separator":              reflect.ValueOf(g.Separator),
		"SliderInt":              reflect.ValueOf(g.SliderInt),
		"SliderFloat":            reflect.ValueOf(g.SliderFloat),
		"HSplitter":              reflect.ValueOf(g.HSplitter),
		"VSplitter":              reflect.ValueOf(g.VSplitter),
		"TabItem":                reflect.ValueOf(g.TabItem),
		"TabBar":                 reflect.ValueOf(g.TabBar),
		"Row":                    reflect.ValueOf(g.Row),
		"Table":                  reflect.ValueOf(g.Table),
		"FastTable":              reflect.ValueOf(g.FastTable),
		"Tooltip":                reflect.ValueOf(g.Tooltip),
		"TreeNode":               reflect.ValueOf(g.TreeNode),
		"Spacing":                reflect.ValueOf(g.Spacing),
		"Custom":                 reflect.ValueOf(g.Custom),
		"Condition":              reflect.ValueOf(g.Condition),
		"ListBox":                reflect.ValueOf(g.ListBox),
		"DatePicker":             reflect.ValueOf(g.DatePicker),
		"Dummy":                  reflect.ValueOf(g.Dummy),
		// "Widget":             reflect.ValueOf(g.Widget),

		"PrepareMessageBox": reflect.ValueOf(g.PrepareMsgbox),
		"MessageBox":        reflect.ValueOf(g.Msgbox),

		"LoadFont": reflect.ValueOf(loadFont),

		"GetConfirm": reflect.ValueOf(getConfirmGUI),

		"SimpleInfo":      reflect.ValueOf(simpleInfo),
		"SimpleError":     reflect.ValueOf(simpleError),
		"SelectFile":      reflect.ValueOf(selectFileGUI),
		"SelectSaveFile":  reflect.ValueOf(selectFileToSaveGUI),
		"SelectDirectory": reflect.ValueOf(selectDirectoryGUI),

		"EditFile":   reflect.ValueOf(editFile),
		"LoopWindow": reflect.ValueOf(loopWindow),

		"LayoutP": reflect.ValueOf(g.Layout{}),
	}

	env.Packages["lcl"] = map[string]reflect.Value{
		"GetApplication":    reflect.ValueOf(getVclApplication),
		"InitVCL":           reflect.ValueOf(initLCL),
		"InitLCL":           reflect.ValueOf(initLCL),
		"NewCheckBox":       reflect.ValueOf(vcl.NewCheckBox),
		"NewLabel":          reflect.ValueOf(vcl.NewLabel),
		"NewButton":         reflect.ValueOf(vcl.NewButton),
		"NewComboBox":       reflect.ValueOf(vcl.NewComboBox),
		"NewEdit":           reflect.ValueOf(vcl.NewEdit),
		"NewCanvas":         reflect.ValueOf(vcl.NewCanvas),
		"NewImage":          reflect.ValueOf(vcl.NewImage),
		"NewList":           reflect.ValueOf(vcl.NewList),
		"NewListBox":        reflect.ValueOf(vcl.NewListBox),
		"NewListView":       reflect.ValueOf(vcl.NewListView),
		"NewListColumns":    reflect.ValueOf(vcl.NewListColumns),
		"NewListItem":       reflect.ValueOf(vcl.NewListItem),
		"NewListItems":      reflect.ValueOf(vcl.NewListItems),
		"NewMainMenu":       reflect.ValueOf(vcl.NewMainMenu),
		"NewMemo":           reflect.ValueOf(vcl.NewMemo),
		"NewMenuItem":       reflect.ValueOf(vcl.NewMenuItem),
		"NewMiniWebview":    reflect.ValueOf(vcl.NewMiniWebview),
		"NewPaintBox":       reflect.ValueOf(vcl.NewPaintBox),
		"NewPanel":          reflect.ValueOf(vcl.NewPanel),
		"NewPicture":        reflect.ValueOf(vcl.NewPicture),
		"NewPopupMenu":      reflect.ValueOf(vcl.NewPopupMenu),
		"NewProgressBar":    reflect.ValueOf(vcl.NewProgressBar),
		"NewRadioButton":    reflect.ValueOf(vcl.NewRadioButton),
		"NewRadioGroup":     reflect.ValueOf(vcl.NewRadioGroup),
		"NewScrollBox":      reflect.ValueOf(vcl.NewScrollBox),
		"NewScrollBar":      reflect.ValueOf(vcl.NewScrollBar),
		"NewSplitter":       reflect.ValueOf(vcl.NewSplitter),
		"NewStatusBar":      reflect.ValueOf(vcl.NewStatusBar),
		"NewStatusPanel":    reflect.ValueOf(vcl.NewStatusPanel),
		"NewStatusPanels":   reflect.ValueOf(vcl.NewStatusPanels),
		"NewTimer":          reflect.ValueOf(vcl.NewTimer),
		"NewToolBar":        reflect.ValueOf(vcl.NewToolBar),
		"NewToolButton":     reflect.ValueOf(vcl.NewToolButton),
		"NewTrayIcon":       reflect.ValueOf(vcl.NewTrayIcon),
		"NewStaticText":     reflect.ValueOf(vcl.NewStaticText),
		"NewSpinEdit":       reflect.ValueOf(vcl.NewSpinEdit),
		"NewSpeedButton":    reflect.ValueOf(vcl.NewSpeedButton),
		"NewShape":          reflect.ValueOf(vcl.NewShape),
		"NewScreen":         reflect.ValueOf(vcl.NewScreen),
		"NewSaveDialog":     reflect.ValueOf(vcl.NewSaveDialog),
		"NewReplaceDialog":  reflect.ValueOf(vcl.NewReplaceDialog),
		"NewPngImage":       reflect.ValueOf(vcl.NewPngImage),
		"NewPen":            reflect.ValueOf(vcl.NewPen),
		"NewPageControl":    reflect.ValueOf(vcl.NewPageControl),
		"NewOpenDialog":     reflect.ValueOf(vcl.NewOpenDialog),
		"NewObject":         reflect.ValueOf(vcl.NewObject),
		"NewMouse":          reflect.ValueOf(vcl.NewMouse),
		"NewMaskEdit":       reflect.ValueOf(vcl.NewMaskEdit),
		"NewLinkLabel":      reflect.ValueOf(vcl.NewLinkLabel),
		"NewLabeledEdit":    reflect.ValueOf(vcl.NewLabeledEdit),
		"NewJPEGImage":      reflect.ValueOf(vcl.NewJPEGImage),
		"NewImageList":      reflect.ValueOf(vcl.NewImageList),
		"NewImageButton":    reflect.ValueOf(vcl.NewImageButton),
		"NewIcon":           reflect.ValueOf(vcl.NewIcon),
		"NewGroupBox":       reflect.ValueOf(vcl.NewGroupBox),
		"NewHeaderControl":  reflect.ValueOf(vcl.NewHeaderControl),
		"NewHeaderSection":  reflect.ValueOf(vcl.NewHeaderSection),
		"NewHeaderSections": reflect.ValueOf(vcl.NewHeaderSections),
		"NewGraphic":        reflect.ValueOf(vcl.NewGraphic),
		"NewGIFImage":       reflect.ValueOf(vcl.NewGIFImage),
		"NewGauge":          reflect.ValueOf(vcl.NewGauge),
		"ShowMessage":       reflect.ValueOf(vcl.ShowMessage),
		"ShowMessageFmt":    reflect.ValueOf(vcl.ShowMessageFmt),
		"MessageDlg":        reflect.ValueOf(vcl.MessageDlg),
		"InputBox":          reflect.ValueOf(vcl.InputBox),
		"InputQuery":        reflect.ValueOf(vcl.InputQuery),
		"ThreadSync":        reflect.ValueOf(vcl.ThreadSync),
		"NewFrame":          reflect.ValueOf(vcl.NewFrame),
		"SelectDirectory":   reflect.ValueOf(vcl.SelectDirectory1),
		"SelectDirectory3":  reflect.ValueOf(vcl.SelectDirectory3),
		"NewForm":           reflect.ValueOf(vcl.NewForm),
		"NewFontDialog":     reflect.ValueOf(vcl.NewFontDialog),
		"NewFont":           reflect.ValueOf(vcl.NewFont),
		"NewFlowPanel":      reflect.ValueOf(vcl.NewFlowPanel),
		"NewFindDialog":     reflect.ValueOf(vcl.NewFindDialog),
		"NewDrawGrid":       reflect.ValueOf(vcl.NewDrawGrid),
		"NewDateTimePicker": reflect.ValueOf(vcl.NewDateTimePicker),
		"NewControl":        reflect.ValueOf(vcl.NewControl),
		"NewComboBoxEx":     reflect.ValueOf(vcl.NewComboBoxEx),
		"NewColorListBox":   reflect.ValueOf(vcl.NewColorListBox),
		"NewColorDialog":    reflect.ValueOf(vcl.NewColorDialog),
		"NewColorBox":       reflect.ValueOf(vcl.NewColorBox),
		"NewCheckListBox":   reflect.ValueOf(vcl.NewCheckListBox),
		"NewBrush":          reflect.ValueOf(vcl.NewBrush),
		"NewBitmap":         reflect.ValueOf(vcl.NewBitmap),
		"NewBitBtn":         reflect.ValueOf(vcl.NewBitBtn),
		"NewBevel":          reflect.ValueOf(vcl.NewBevel),
		"NewApplication":    reflect.ValueOf(vcl.NewApplication),
		"NewAction":         reflect.ValueOf(vcl.NewAction),
		"NewActionList":     reflect.ValueOf(vcl.NewActionList),

		"GetLibVersion": reflect.ValueOf(vcl.GetLibVersion),

		// values
		"PoDesigned":        reflect.ValueOf(types.PoDesigned),
		"PoDefault":         reflect.ValueOf(types.PoDefault),
		"PoDefaultPosOnly":  reflect.ValueOf(types.PoDefaultPosOnly),
		"PoDefaultSizeOnly": reflect.ValueOf(types.PoDefaultSizeOnly),
		"PoScreenCenter":    reflect.ValueOf(types.PoScreenCenter),
		"PoMainFormCenter":  reflect.ValueOf(types.PoMainFormCenter),
		"PoOwnerFormCenter": reflect.ValueOf(types.PoOwnerFormCenter),
		"PoWorkAreaCenter":  reflect.ValueOf(types.PoWorkAreaCenter),
	}

	// env.Packages["lcl/types"] = map[string]reflect.Value{
	// }

	var widget g.Widget

	env.PackageTypes["gui"] = map[string]reflect.Type{
		"Layout": reflect.TypeOf(g.Layout{}),
		// "Signal": reflect.TypeOf(&signal).Elem(),
		"Widget": reflect.TypeOf(&widget).Elem(),
	}

}

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

	runScript(editorG.GetText(), "new", editArgsG)
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

				sourceT := editorG.GetText()

				parser.EnableErrorVerbose()
				_, errT := parser.ParseSrc(sourceT)
				// tk.Plv(stmts)

				e, ok := errT.(*parser.Error)

				if ok {
					errMarkersG.Clear()
					errMarkersG.Insert(e.Pos.Line, tk.Spr("[col: %v, size: %v] %v", e.Pos.Column, errMarkersG.Size(), e.Error()))

					editorG.SetErrorMarkers(errMarkersG)

				} else if errT != nil {
					g.Msgbox("Error", tk.Spr("%#v", errT))
				} else {
					g.Msgbox("Info", "Syntax check passed.")
				}

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
	setVar("FontRange", "COMMON")
	setVar("FontSize", "15")

	wnd := g.NewMasterWindow("Gox Editor", 800, 600, 0, loadFont)
	// tk.Pl("%T", wnd)
	wnd.Main(editorLoop)

}

func runFile(argsA ...string) interface{} {
	lenT := len(argsA)

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

	fcT := tk.LoadStringFromFile(argsA[0])

	if tk.IsErrorString(fcT) {
		return tk.Errf("Invalid file content: %v", tk.GetErrorString(fcT))
	}

	return runScript(fcT, "", argsA[1:]...)
}

// GUI related end

func initTengoVM() {

}

func initJSVM() {
	if jsVMG == nil {
		jsVMG = goja.New()

		jsVMG.Set("printf", func(call goja.FunctionCall) goja.Value {
			callArgsT := call.Arguments

			argsBufT := make([]interface{}, len(callArgsT)-1)

			formatA := callArgsT[0].ToString().String()

			for i, v := range callArgsT {
				if i == 0 {
					continue
				}

				argsBufT[i-1] = v.ToString().String()
			}

			tk.Prf(formatA, argsBufT...)

			return nil
		})

		jsVMG.Set("printfln", func(call goja.FunctionCall) goja.Value {
			callArgsT := call.Arguments

			argsBufT := make([]interface{}, len(callArgsT)-1)

			formatA := callArgsT[0].ToString().String()

			for i, v := range callArgsT {
				if i == 0 {
					continue
				}

				argsBufT[i-1] = v.ToString().String()
			}

			tk.Prf(formatA+"\n", argsBufT...)

			return nil
		})

		jsVMG.Set("println", func(call goja.FunctionCall) goja.Value {
			callArgsT := call.Arguments

			argsBufT := make([]interface{}, len(callArgsT))

			for i, v := range callArgsT {
				argsBufT[i] = v.ToString().String()
			}

			fmt.Println(argsBufT...)

			return nil
		})

		// jsVMG.Set("goGetRandomInt", func(call goja.FunctionCall) goja.Value {
		// 	maxA := call.Argument(0).ToInteger()

		// 	randomNumberT := rand.Intn(int(maxA))

		// 	rs := jsVMG.ToValue(randomNumberT)

		// 	return rs
		// })

		jsVMG.Set("getVar", func(call goja.FunctionCall) goja.Value {
			nameA := call.Argument(0).ToString().String()

			objT, ok := variableG[nameA]

			if !ok {
				return jsVMG.ToValue(tk.GenerateErrorString("no key"))
			}

			rs := jsVMG.ToValue(objT)

			return rs
		})

		jsVMG.Set("setVar", func(call goja.FunctionCall) goja.Value {
			nameA := call.Argument(0).ToString().String()

			objT := call.Argument(1).ToString().String()

			variableG[nameA] = objT

			return nil
		})

		consoleStrT := `console = { log: println };
		String.prototype.startsWith = function (s) {
			if (s == null || s == "" || this.length == 0 || s.length > this.length)
				return false;
			if (this.substr(0, s.length) == s)
				return true;
			else
				return false;
			return true;
		}
		
		String.prototype.endsWith = function (s) {
			if (s == null || s == "" || this.length == 0 || s.length > this.length)
				return false;
			if (this.substring(this.length - s.length) == s)
				return true;
			else
				return false;
			return true;
		}
		
		
		`

		_, errT := jsVMG.RunString(consoleStrT)
		if errT != nil {
			tk.Pl("failed to initialize JS VM: %v", errT)

			return
		}

	}

}

// init the main VM
func initAnkVM() {
	if ankVMG == nil {
		importAnkNonGUIPackages()

		// GUI related start

		importAnkGUIPackages()

		// GUI related end

		ankVMG = env.NewEnv()

		initAnkoVMInstance(ankVMG)

		ankVMG.Define("argsG", os.Args[1:])

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

	scriptsT := tk.GetAllParameters(argsT)[1:]

	lenT := len(scriptsT)

	// GUI related start

	if tk.IfSwitchExistsWhole(argsT, "-edit") {
		if lenT < 1 {
			editFile("")
		} else {
			editFile(scriptsT[0])
		}

		return
	}

	// GUI related end

	if lenT < 1 {
		initAnkVM()

		runInteractive()

		// tk.Pl("not enough parameters")

		return
	}

	encryptCodeT := tk.GetSwitchWithDefaultValue(argsT, "-encrypt=", "")

	if encryptCodeT != "" {
		for i, v := range scriptsT {
			fcT := tk.LoadStringFromFile(v)

			if tk.IsErrorString(fcT) {
				tk.Pl("failed to load file [%v] %v: %v", i, v, tk.GetErrorString(fcT))
				continue
			}

			encStrT := tk.EncryptStringByTXDEF(fcT, encryptCodeT)

			if tk.IsErrorString(encStrT) {
				tk.Pl("failed to encrypt content [%v] %v: %v", i, v, tk.GetErrorString(encStrT))
				continue
			}

			rsT := tk.SaveStringToFile("//TXDEF#"+encStrT, v+"e")

			if tk.IsErrorString(rsT) {
				tk.Pl("failed to encrypt file [%v] %v: %v", i, v, tk.GetErrorString(rsT))
				continue
			}
		}

		return
	}

	decryptCodeT := tk.GetSwitchWithDefaultValue(argsT, "-decrypt=", "")

	if decryptCodeT != "" {
		for i, v := range scriptsT {
			fcT := tk.LoadStringFromFile(v)

			if tk.IsErrorString(fcT) {
				tk.Pl("failed to load file [%v] %v: %v", i, v, tk.GetErrorString(fcT))
				continue
			}

			decStrT := tk.DecryptStringByTXDEF(fcT, decryptCodeT)

			if tk.IsErrorString(decStrT) {
				tk.Pl("failed to decrypt content [%v] %v: %v", i, v, tk.GetErrorString(decStrT))
				continue
			}

			rsT := tk.SaveStringToFile(decStrT, v+"d")

			if tk.IsErrorString(rsT) {
				tk.Pl("failed to decrypt file [%v] %v: %v", i, v, tk.GetErrorString(rsT))
				continue
			}
		}

		return
	}

	decryptRunCodeT := tk.GetSwitchWithDefaultValue(argsT, "-decrun=", "")

	if !tk.IfSwitchExistsWhole(argsT, "-m") {
		scriptsT = scriptsT[0:1]
	}

	ifExampleT := tk.IfSwitchExistsWhole(argsT, "-example")
	ifRemoteT := tk.IfSwitchExistsWhole(argsT, "-remote")
	ifCloudT := tk.IfSwitchExistsWhole(argsT, "-cloud")
	ifViewT := tk.IfSwitchExistsWhole(argsT, "-view")

	for _, scriptT := range scriptsT {
		if tk.EndsWith(scriptT, ".js") {
			var fcT string

			if ifExampleT {
				fcT = tk.DownloadPageUTF8("https://raw.githubusercontent.com/topxeq/gox/master/scripts/"+scriptT, nil, "", 30)
			} else if ifRemoteT {
				fcT = tk.DownloadPageUTF8(scriptT, nil, "", 30)
			} else if ifCloudT {
				fcT = tk.DownloadPageUTF8("http://scripts.frenchfriend.net/xaf/scripts/"+scriptT, nil, "", 30)
			} else {
				fcT = tk.LoadStringFromFile(scriptT)
			}

			if tk.IsErrorString(fcT) {
				tk.Pl("failed to load script from %v: %v", scriptT, tk.GetErrorString(fcT))

				continue
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

			initJSVM()

			jsVMG.Set("argsG", jsVMG.ToValue(argsT))

			_, errT := jsVMG.RunString(fcT)
			if errT != nil {
				tk.Pl("failed to run script(%v): %v", scriptT, errT)

				continue
			}

			result := jsVMG.Get("resultG")

			if result != nil {
				tk.Pl("%#v", result)
			}

			return
		} else if tk.EndsWith(scriptT, ".tg") || tk.EndsWith(scriptT, ".tengo") {
			var fcT string

			if ifExampleT {
				fcT = tk.DownloadPageUTF8("https://raw.githubusercontent.com/topxeq/gox/master/scripts/"+scriptT, nil, "", 30)
			} else if ifRemoteT {
				fcT = tk.DownloadPageUTF8(scriptT, nil, "", 30)
			} else if ifCloudT {
				fcT = tk.DownloadPageUTF8("http://scripts.frenchfriend.net/xaf/scripts/"+scriptT, nil, "", 30)
			} else {
				fcT = tk.LoadStringFromFile(scriptT)
			}

			if tk.IsErrorString(fcT) {
				tk.Pl("failed to load script from %v: %v", scriptT, tk.GetErrorString(fcT))

				continue
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

			initTengoVM()

			scriptT := tengo.NewScript([]byte(fcT))

			scriptT.SetImports(tengoModulesG)

			_ = scriptT.Add("setVar", setVarTengo)
			errT := scriptT.Add("times", times)
			if errT != nil {
				tk.Pl("failed to add times(%v) error: %v", "", errT)
				continue
			}

			errT = scriptT.Add("getVar", getVarTengo)
			if errT != nil {
				tk.Pl("failed to add getVar(%v) error: %v", "", errT)
				continue
			}

			argsG_TG := make([]interface{}, len(argsT))
			for i, v := range argsT {
				argsG_TG[i] = v
			}

			errT = scriptT.Add("argsG", argsG_TG)
			if errT != nil {
				tk.Pl("failed to add argsA(%v) error: %v", "", errT)
				continue
			}

			compiledT, errT := scriptT.RunContext(context.Background())
			if errT != nil {
				tk.Pl("failed to execute script(%v) error: %v", "", errT)
				continue
			}

			rs := compiledT.Get("resultG")

			// if errT == nil && rs != nil {
			tk.Pl("%#v", rs)
			// }

		} else { // if tk.EndsWith(scriptT, ".ank") || tk.EndsWith(scriptT, ".gox") {
			var fcT string

			if ifExampleT {
				if (!tk.EndsWith(scriptT, ".gox")) && (!tk.EndsWith(scriptT, ".js")) && (!tk.EndsWith(scriptT, ".tg")) {
					scriptT += ".gox"
				}
				fcT = tk.DownloadPageUTF8("https://gitee.com/topxeq/gox/raw/master/scripts/"+scriptT, nil, "", 30)
			} else if ifRemoteT {
				fcT = tk.DownloadPageUTF8(scriptT, nil, "", 30)
			} else if ifCloudT {
				if (!tk.EndsWith(scriptT, ".gox")) && (!tk.EndsWith(scriptT, ".js")) && (!tk.EndsWith(scriptT, ".tg")) {
					scriptT += ".gox"
				}
				fcT = tk.DownloadPageUTF8("http://scripts.frenchfriend.net/xaf/scripts/"+scriptT, nil, "", 30)
			} else {
				fcT = tk.LoadStringFromFile(scriptT)
			}

			if tk.IsErrorString(fcT) {
				tk.Pl("failed to load script from %v: %v", scriptT, tk.GetErrorString(fcT))

				continue
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

			initAnkVM()

			script := fcT //`println("Hello World :)")`

			rs1, errT := vm.Execute(ankVMG, nil, script)
			if errT != nil {

				posStrT := ""

				e, ok := errT.(*parser.Error)

				// tk.Pl("%#v", ankVMG)
				if ok {
					posStrT = fmt.Sprintf(", line: %v, col: %v", e.Pos.Line, e.Pos.Column)
				} else {
					e, ok := errT.(*vm.Error)

					if ok {
						posStrT = fmt.Sprintf(", line: %v, col: %v", e.Pos.Line, e.Pos.Column)
					} else {
						tk.Pl("%#v", errT)
					}
				}

				tk.Pl("failed to execute script(%v%v) error: %v\n%#v\n", scriptT, posStrT, errT, rs1)
				continue
			}

			rs, errT := ankVMG.Get("outG")

			if errT == nil && rs != nil {
				tk.Pl("%#v", rs)
			}

		}
	}
}
