package main

import(
    "fmt"
    "tk"
)

func main() {
    a := 16.8
    b := 7.7

    fmt.Println(a * b - 1)

    tk.Printfln("This: %#v, timeFormatCompact: %v", a * b - 1, tk.TimeFormatCompact)
}