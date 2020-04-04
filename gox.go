package main

import (
	"math/rand"
	"os"
	"reflect"
	"time"

	"fmt"

	"github.com/dop251/goja"
	"github.com/mattn/anko/core"
	"github.com/mattn/anko/env"
	_ "github.com/mattn/anko/packages"
	"github.com/mattn/anko/vm"
	"github.com/topxeq/tk"

	"github.com/sqweek/dialog"
)

// var inG interface{}
// var outG interface{}

var variableG = make(map[string]interface{})

var jsVMG *goja.Runtime = nil
var ankVMG *env.Env = nil

func getVar(nameA string) interface{} {
	return variableG[nameA]
}

func setVar(nameA string, valueA interface{}) {
	variableG[nameA] = valueA
}

func importAnkPackages() {
	env.Packages["tk"] = map[string]reflect.Value{
		"Pl":                reflect.ValueOf(tk.Pl),
		"Prl":               reflect.ValueOf(tk.Prl),
		"Prf":               reflect.ValueOf(tk.Prf),
		"SleepMilliSeconds": reflect.ValueOf(tk.SleepMilliSeconds),
		"SleepSeconds":      reflect.ValueOf(tk.SleepSeconds),
	}

	env.Packages["dialog"] = map[string]reflect.Value{
		"Message": reflect.ValueOf(dialog.Message),
	}

}

func main() {
	rand.Seed(time.Now().Unix())

	argsT := os.Args

	scriptT := tk.GetParameterByIndexWithDefaultValue(argsT, 1, "")

	if scriptT == "" {
		tk.Pl("not enough parameters")

		return
	}

	if tk.EndsWith(scriptT, ".js") {
		fcT := tk.LoadStringFromFile(scriptT)

		if tk.IsErrorString(fcT) {
			tk.Pl("failed to load file content: %v", tk.GetErrorString(fcT))

			return
		}

		if jsVMG == nil {
			jsVMG = goja.New()

			jsVMG.Set("goPrintf", func(call goja.FunctionCall) goja.Value {
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

			jsVMG.Set("goPrintln", func(call goja.FunctionCall) goja.Value {
				callArgsT := call.Arguments

				argsBufT := make([]interface{}, len(callArgsT))

				for i, v := range callArgsT {
					argsBufT[i] = v.ToString().String()
				}

				fmt.Println(argsBufT...)

				return nil
			})

			jsVMG.Set("goGetRandomInt", func(call goja.FunctionCall) goja.Value {
				maxA := call.Argument(0).ToInteger()

				randomNumberT := rand.Intn(int(maxA))

				rs := jsVMG.ToValue(randomNumberT)

				return rs
			})

			consoleStrT := `console = { log: goPrintln };`

			_, errT := jsVMG.RunString(consoleStrT)
			if errT != nil {
				tk.Pl("failed to run script: %v", errT)

				return
			}

		}

		v, errT := jsVMG.RunString(fcT)
		if errT != nil {
			tk.Pl("failed to run script: %v", errT)

			return
		}

		variableG["Out"] = v.Export()

		// tk.Pl("%#v", rs)

		return
	} else if tk.EndsWith(scriptT, ".ank") || tk.EndsWith(scriptT, ".gox") {
		fcT := tk.LoadStringFromFile(scriptT)

		if tk.IsErrorString(fcT) {
			tk.Pl("failed to load file content: %v", tk.GetErrorString(fcT))

			return
		}

		if ankVMG == nil {
			importAnkPackages()

			ankVMG = env.NewEnv()

			// err := e.Define("pl", tk.Pl)
			// if err != nil {
			// 	tk.CheckErrf("Define error: %v\n", err)
			// }

			// e.Define("prl", tk.Prl)
			ankVMG.Define("println", tk.Prl)
			// e.Define("prf", tk.Prf)

			ankVMG.Define("setVar", setVar)
			ankVMG.Define("getVar", getVar)

			core.Import(ankVMG)

		}

		ankVMG.Define("inG", map[string]interface{}{"Args": os.Args})

		script := fcT //`println("Hello World :)")`

		_, errT := vm.Execute(ankVMG, nil, script)
		if errT != nil {
			tk.CheckErrf("Execute error: %v\n", errT)
		}

		rs, errT := ankVMG.Get("outG")

		// tk.CheckErrCompact(errT)

		if errT == nil && rs != nil {
			tk.Pl("%#v", rs)
		}

		// tk.Pl("%#v", rs)

		return
	}

	tk.Pl("Gox by TopXeQ V0.9a")

	fmt.Println("")

}
