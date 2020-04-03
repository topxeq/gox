package main

import (
	"math/rand"
	"os"
	"time"

	"fmt"

	"github.com/dop251/goja"
	"github.com/mattn/anko/core"
	"github.com/mattn/anko/env"
	_ "github.com/mattn/anko/packages"
	"github.com/mattn/anko/vm"
	"github.com/topxeq/tk"
)

var resultG interface{}

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

		vmT := goja.New()

		vmT.Set("goPrintf", func(call goja.FunctionCall) goja.Value {
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

		vmT.Set("goPrintln", func(call goja.FunctionCall) goja.Value {
			callArgsT := call.Arguments

			argsBufT := make([]interface{}, len(callArgsT))

			for i, v := range callArgsT {
				argsBufT[i] = v.ToString().String()
			}

			fmt.Println(argsBufT...)

			return nil
		})

		vmT.Set("goGetRandomInt", func(call goja.FunctionCall) goja.Value {
			maxA := call.Argument(0).ToInteger()

			randomNumberT := rand.Intn(int(maxA))

			rs := vmT.ToValue(randomNumberT)

			return rs
		})

		consoleStrT := `console = { log: goPrintln };
		`

		v, errT := vmT.RunString(consoleStrT + fcT)
		if errT != nil {
			tk.Pl("failed to run script: %v", errT)

			return
		}

		resultG = v.Export()

		// tk.Pl("%#v", rs)

		return
	} else if tk.EndsWith(scriptT, ".ank") {
		fcT := tk.LoadStringFromFile(scriptT)

		if tk.IsErrorString(fcT) {
			tk.Pl("failed to load file content: %v", tk.GetErrorString(fcT))

			return
		}

		e := env.NewEnv()

		err := e.Define("pl", tk.Pl)
		if err != nil {
			tk.CheckErrf("Define error: %v\n", err)
		}

		e.Define("prl", tk.Prl)
		e.Define("println", tk.Prl)
		e.Define("prf", tk.Prf)

		e.Define("inG", map[string]interface{}{"Args": os.Args})

		core.Import(e)

		script := fcT //`println("Hello World :)")`

		_, err = vm.Execute(e, nil, script)
		if err != nil {
			tk.CheckErrf("Execute error: %v\n", err)
		}

		rs, errT := e.Get("resultG")

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
