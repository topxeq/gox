package main

import (
	"math/rand"
	"os"
	"time"

	"fmt"

	"github.com/d5/tengo/v2"
	"github.com/dop251/goja"
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

	if tk.EndsWith(scriptT, ".tg") {
		fcT := tk.LoadBytes(scriptT, -1)

		if fcT == nil {
			tk.Pl("failed to load file content")

			return
		}

		s := tengo.NewScript(fcT)
		// s.SetImports((*tengo.ModuleMap)(stdlib.GetModuleMap("fmt")))
		// s.SetImports(stdlib.GetModuleMap("fmt"))
		// s.SetImports(stdlib.GetModuleMap(stdlib.AllModuleNames()...))
		// s.SetImports(*tengo.ModuleMapstdlib.GetModuleMap(stdlib.AllModuleNames()...))

		if _, errT := s.Run(); errT != nil {
			tk.Pl("failed to run script: %v", errT)

			return
		}

		return
	} else if tk.EndsWith(scriptT, ".js") {
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
	}

	tk.Pl("Gox by TopXeQ V0.9a")

	fmt.Println("")

}
