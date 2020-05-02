![Gox](https://gitee.com/topxeq/gox/raw/master/docs/goxlogo2xs.jpg)

_Gox语言的中文简介可以从[这篇文章](https://mbd.baidu.com/newspage/data/landingshare?pageType=1&isBdboxFrom=1&context=%7B%22nid%22%3A%22news_9677952217244209303%22%2C%22sourceFrom%22%3A%22bjh%22%7D)开始看起，或者[这里](https://www.jianshu.com/nb/44919205)看更多的开发指南。_

# gox
Gox (or Goxlang) is a free, open-source script language or a interpreter written by Golang. It's based on [Qlang](https://github.com/qiniu/qlang), [Anko](https://github.com/mattn/anko) [Goja](https://github.com/dop251/goja) and [Tengo](https://github.com/d5/tengo), with some improvement. As a script runner(or interpreter), Gox supports various script languages such as txScript, Javascript, Qlang, Anko and Tengo. The syntax of Gox language is very similar to Golang.

Golang is not required to be installed. Gox is only in one executable file, green and fast.

And thanks to [Giu](https://github.com/AllenDang/giu), which enables Gox to provide a modern GUI programming ability, and it's cross-platform, native, no dependencies and convenient. Even more, Gox has an code editor embedded, so the external text editor may not be required for small piece of code. Note: this GUI library requires OPENGL.

And also many thanks to [Govcl](https://github.com/ying32/govcl) written by ying32, which enables Gox to provide GUI programming APIs based on the free Lazarus LCL library. It's from VCL library and very useful especially for experienced Delphi/VCL programmers. Since this library doesn't require OPENGL, it's more compatible in some server-side operating systems. It only requires a single library file(dll in Windows or so in Linux and dylib in Mac OS).

Gox supports password-protect your source code file, it's also a major difference from most script/interpretive language.

## 1. Installation

Download the latest Gox executable file from the [official website](http://gox.topget.org/) or [Github](https://github.com/topxeq/gox/releases) or [Gitee](https://gitee.com/topxeq/gox/releases) (due to the speed issue, official website and Gitee are recommended). Then put it into a directory in the system path. If you download the zip file, extract it first.

A command-line only version of gox (named goxc, no GUI features) is also available. This version will be more fast, and of course smaller in file size.

Since the more features added makes the Gox executable became very large, the Tiny Gox Version is now available to download. Some features are removed from Gox tiny version, such as OpenGL GUI, drawing, NoSQL, some database drivers(Oracle, MS-SQL), etc.

Goxg version is used to run GUI only programs, it will not open a CMD console.

From version 0.97a, Qlang became the default script-engine for Gox, and is prefered to use.

## 2.usage

### Check the version.

> gox -version

### Show help.

> gox -h

### Run a script file.

> gox test.gox

or

> gox test.js

Currently Only ECMAScript 5.1(+) is supported for Javascript.

### Run several script files in consequence, use "-m" switch to enable multi-scripts mode.

> gox -m test.gox script1.js start.ank last.txs

the result should be like below,

```
> gox -m basic.gox test.js a.tg
3.4000000000000004
this is 5 + 12.5 = 17.5
5 12.5 17.5
19 33 627
The random number is：1
```

Attention: in multi-scripts mode, all the command-line parameters will be recognizd as script file names, so inside each script, it can only retrieve its command-line arguments through switches. For example:

```
gox -m basic.gox test.js -para=abcd
```

And all the switches in Gox should be like "-type=code", not "-type code". The globle variables across VMs could be shared to other scripts in this mode.

### Start the interpreter in REPL mode.

> gox

the REPL runs like this,

```
> a=1                      
1                          
> b=2.3                    
2.3                        
> a+b                      
3.3                        
> printfln("a: %v", a)     
a: 1                       
<nil>                      
> exit()                   
exit status 1              
```

use "quit()" or "exit()" function to exit.

### Select script file to run in REPL mode

> run()

Just call run() function in REPL, a file selection dialog will popup.

### Run example script from Gitee directly

> gox -example basic.gox

Gox will download the example script directly from Github and run it.

### View/Save example script from Gitee

> gox -example -view basic

the basic.gox (the suffix could be omitted) script will be printed to standard output. Use redirect to save the script, i.e.

> gox -example -view basic > d:\scripts\a.gox


### Run online script directly

> gox -remote https://raw.githubusercontent.com/topxeq/gox/master/scripts/basica.gox

Gox will try to download the script first and then run it.

### Encrypt source code file

> gox -encrypt=mycode basic.gox

An encrypted file with an additional "e" of the file name(i.e. basic.goxe) will be created in the same directory of the original source code file. "mycode" is the password to encrypt the file, and of course remember this code to decrypt it back.

### Decrypt/restore encrypted source code file

> gox -decrypt=mycode basic.goxe

An decrypted file with an additional "d" of the file name(i.e. basic.goxed) will be created in the same directory of the original encrypted source code file. Manually rename the file to "basic.gox" if necessary.

### Running an encrypted script directly

> gox -decrun=mycode basic.goxe

or 

> gox -decrun=mycode -example basic.goxe

### Open a simple editor to edit Gox code

> gox -edit

### Open a simple editor to edit specific Gox code file

> gox -edit d:\tmp\basic.gox

the screenshot of Gox editor is as below,

![Gox Editor](https://gitee.com/topxeq/gox/raw/master/docs/goxeditor2.png)

Although Gox provides a simple code editor, the editors with more powerful functions are recommended, such as Visual Studio Code. Currently, you can set the .gox files the same syntax highlight scheme as Golang files.

## 3. User/developer guide

### File encoding

All the script file is better in UTF-8 encoding, and its line-endings are better to use "\n" or "\r\n".

### Stop/Terminate the script

Using exit(), os.Exit(1) or return in the script is valid. In interactive mode, quit() is valid, too.

### command-line parameters and swithes

The global value "argsG" could be used for retrieve command-line arguments, and the first element(the Gox executable) is removed. If you need the whole command-line, use os.Args instead.

All the command-line switches/flags used by Gox itself are not recommended to be used in the scripts.

An example for command-line handling is as below([source code](https://gitee.com/topxeq/gox/tree/master/scripts/commandLine.ank)),

```
// test command-line functions
// for example: gox scripts\commandLine.ank abc -file=a.txt

os = import("os")

println("The whole command-line: ", os.Args)
println("The whole command-line without executable: ", argsG)

lenT = len(argsG)

if lenT > 0 {
    printfln("The first command-line element is: %v", argsG[0])
}

if lenT > 1 {
    printfln("The sencod command-line element is: %v", argsG[1])
}

var tk = import("tk")

para1 = tk.GetParameterByIndexWithDefaultValue(argsG, 1, "")
pfl("para1=%v", para1)

para2 = tk.GetParameterByIndexWithDefaultValue(argsG, 2, "")
pfl("para2=%v", para2)

switch1 = tk.GetSwitchWithDefaultValue(argsG, "-file=", "")

pl("switch1:", switch1)

paras = tk.GetAllParameters(argsG)

pl("All parameters:", paras)

switches = tk.GetAllSwitches(argsG)

pl("All switches:", switches)

pl(tk.IfSwitchExistsWhole(argsG, "-file"))
pl(tk.IfSwitchExists(argsG, "-file"))
```

the output for this script should be something like,

```
> gox scripts\commandLine.gox abc -file=a.txt
The whole command-line:  [gox scripts\commandLine.gox abc -file=a.txt]
The whole command-line without executable:  [scripts\commandLine.gox abc -file=a.txt]
The first command-line element is: scripts\commandLine.gox
The sencod command-line element is: abc
para1=abc
para2=
switch1: a.txt
All parameters: [scripts\commandLine.gox abc]
All switches: [-file=a.txt]
false
true
```

### Import packages

In Gox with Anko engine, the packages can be imported any time, but remember to assign it to a variable, such as,

```
var fmt = import("fmt")

fmt.Println("abc")

os = import("os")

fmt.Printf("%v\n", os.Args)

```

Along with most of the core Golang libraries, the "tk" package([here](https://github.com/topxeq/tk)) is often used. The "tk" package provides many useful functions. See example scripts for more. And refer to the docs [here](https://pkg.go.dev/github.com/topxeq/tk?tab=doc).

Refer to the documents of these Golang packages for the detailed usage.

### Data type conversion

If using Anko engine, some "to" functions could be used as below,

```
a = 1
b = 2

println("type of a is:", typeof(a))

println("a + b =", a+b)
printfln("a + b = %#v", a+b)

a1 = toString(a)
b1 = toString(b)

printfln("type of a1 is: %T", a1)
printfln("value of a1 is: %v", a1)
printfln("internal value of a1 is: %#v", a1)

println("a1 + b1 =", a1+b1)
printfln("a1 + b1 = %#v", a1+b1)

a2 = toFloat(a1)
b2 = toFloat(b1)

printfln("a2 + b2 = %#v", a2+b2)
printfln("type of a2 + b2 is: %T", a2+b2)

```

the running result is,

```
λ gox scripts\dataTypeConversion.ank
type of a is: int64
a + b = 3
a + b = 3
type of a1 is: string
value of a1 is: 1
internal value of a1 is: "1"
a1 + b1 = 12
a1 + b1 = "12"
a2 + b2 = 3
type of a2 + b2 is: float64
```

These "to" function include:

> toString/toStringSlice, toStringFromRuneSlice, toBool(and tryToBool which returns the result like (bool, error))/toBoolSlice, toFloat64/tryToFloat64/toFloat/toFloatSlice, toInt64/tryToInt64, toInt/toExactInt/tryToInt/toIntSlice, toRune/toRuneSlice, toByteSlice, toChar, toDuration

## 4. More Topics and Sample Scripts

### Sample Javascript file:

```
var a = 5 + 12.5;

goPrintf("this is %v + %v = %v\n", 5, 12.5, a);

console.log(5, 12.5, a);

goPrintln(19, 33, 19 * 33);

goPrintfln("The random number：%v", goGetRandomInt(20));

```

### Basic Gox script:

```

// do simple add operation
x = 1.2
y = x + 1

println(x+y)

```

### Basic script with GUI

A simple calculator with GUI(OPENGL is required) in Anko engine.

```
var gui = import("gui")

text1 = ""


func onButton1Click() {
	rs = eval(text1)
	text1 = toString(rs)
}

func onButton2Click() {
	exit()
}

func loop() {

	layoutT = []gui.Widget{
		gui.Label("Enter an expression."), 
		gui.InputText("", 0, &text1), 
		gui.Line(gui.Button("Calculate", onButton1Click), gui.Button("Close", onButton2Click)),
	}

	gui.SingleWindow("Calculator", layoutT)
}

mainWindow = gui.NewMasterWindow("Calculator", 400, 200, gui.MasterWindowFlagsNotResizable, nil)

mainWindow.Main(loop)
```

The screen shot while running the script is like,

![Calculator](https://gitee.com/topxeq/gox/raw/master/docs/calculatorss.png)

### Basic script with GUI by LCL library

The script below(use Anko engine) acts almost the same as the script above, but written with LCL library support, and OPENGL is not required, so it's more compatible in the server-side environment.

```
lcl = import("lcl")
os = import("os")

errT = lcl.InitLCL()

if errT != nil {
	plerr(errT)
	exit()
}

application = lcl.GetApplication()

application.Initialize()

application.SetTitle("Calculator with LCL")
application.SetMainFormOnTaskBar(true)

mainForm = application.CreateForm()

mainForm.SetWidth(400)
mainForm.SetHeight(200)
mainForm.SetCaption("Calculator with LCL")
mainForm.SetPosition(lcl.PoScreenCenter)

mainForm.Font().SetSize(11)

mainForm.SetOnDestroy(&func(sender) {
	println("Form Destroyed.")
})

label1 = lcl.NewLabel(mainForm)
label1.SetParent(mainForm)
label1.SetLeft(10)
label1.SetTop(10)
label1.Font().SetName("Arial")
label1.Font().SetSize(18)

label1.SetCaption("Enter an expression")

edit1 = lcl.NewEdit(mainForm)
edit1.SetParent(mainForm)
edit1.SetBounds(10, 48, 200, 32)
edit1.Font().SetSize(11)

onClick1 = func (objA) {
	rs = eval(edit1.Text())
	edit1.SetText(toString(rs))
}

onClick2 = func(sender) {
	application.Terminate()
}

button1 = lcl.NewButton(mainForm)
button1.SetParent(mainForm)
button1.SetLeft(20)
button1.SetTop(90)
button1.SetCaption("Go")
button1.SetOnClick(&onClick1)

button2 = lcl.NewButton(mainForm)
button2.SetParent(mainForm)
button2.SetLeft(110)
button2.SetTop(90)
button2.SetCaption("Close")
button2.SetOnClick(&onClick2)

application.Run()


```

The screen shot while running the script is like,

![Calculator](https://gitee.com/topxeq/gox/raw/master/docs/lclgui.png)


## 5. More Examples

Browse the example scripts [here](https://gitee.com/topxeq/gox/tree/master/scripts)

## 6. Language

### Difference between Gox and Golang

* Gox is a dynamic-type language, so one of the major difference is no data-type declaration necessary in Gox.

```
a = 10

a = "this"

f = fn(p1, p2) {

}

```

It's very similar for function parameters and return values.

## 7. Library Reference

First, since Gox is based on Qlang/Anko and written by Golang, most of the core libraries of Golang will be available. So try to import the modules from Golang(but Golang installation is not required), and refer to the Golang documents([here](https://golang.org/doc/), [here](https://pkg.go.dev/) or [here](http://docscn.studygolang.com/)). In addition, you can browse Qlang's, Anko's, Goja's, [Govcl](https://gitee.com/ying32/govcl/wikis/pages)'s and Giu's documents.

The default Golang packages available(could imported in Gox code while using Anko engine) include: bytes, encoding/json, errors, flag, fmt, image, image/color, image/draw, image/jpg, image/png, io, io.ioutil, log, math, math/big, math/rand, net, net/http, net/http/cookiejar, net/url, os, os/exec, os/signal, path, path/filepath, regexp, runtime, sort, strconv, strings, sync, time.

Library tk (github.com/topxeq/tk) is the most frequently used package in Gox, the documents are [here](https://godoc.org/github.com/topxeq/tk) or [here](https://pkg.go.dev/github.com/topxeq/tk).

However, Gox provides some convenient global variables and functions decribed as below.

### Variables

#### argsG

get global variable

The global value "argsG" could be used for retrieve command-line arguments, and the first element(the Gox executable) is removed. If you need the whole command-line, use os.Args instead.

---


### Functions

Note: some functions may exist or not in different script engine, and may have some slight differences.

---


#### getVar

get global variable

---

#### setVar

set global variable

---

### setValue

assign a value by a pointer, used in Qlang engine.

```
s = new(string)

// *s = "abc" is not correct in Qlang engine

setValue(s, "abc")

println(*s)  // use * for dereference a value from pointer is allowed

```

---

#### defined

check if a variable is defined

---


#### print

the same as fmt.Print

---

#### printf

the same as fmt.Printf

---


#### println/pln

the same as fmt.Println

---


#### printfln/pl

the same as fmt.Printf but add a new-line character at the end

---

#### plv

the same as pl("%#v", v)

---

#### pv

output the name, type, value of a variable, attention: the parameter passed to this function should be a string, and only the global varibles are allowed.

```
s2 = "abcabcabc"

pv("s2")

// the output ->
// s2(string): abcabcabc
```

---

#### plerr

a convenient way to print an error value

---

#### checkError

a convenient way to print an error value if not nil, and terminate the whole program running

```
outT, errT = clientT.Run(cmdT)

checkError(errT, deferFunc)

```

deferFunc is the function which will be called before terminating the application, if none, pass nil for it, i.e. checkError(errT, nil)

---

#### checkErrorString

the same as checkError, but check a TXERROR string

---


#### getInput

get user input from command-line

```
printf("A:")

a = getInput()

printf("B:")

b = getInput()

println("A + B =", a+b)
```

---

#### getInputf

the same as getInput, but use printf to print a prompt string

```
n = 3

a = getInputf("Please enter the %v value: ", n)

printf("B:")

b = getInput()

println("A + B =", a+b)
```

---


#### eval

evaluate an expression and return the result

---

#### panic

raise a panic manually

```

try {
	panic("a manual panic")
} catch e {
	printfln("error: %v", e)
} finally {
	println("final")
}


try {
	panic(12345678)
} catch e {
	printfln(e)
}


```

---

#### typeof/typeOf/kindOf

return the string representation of the type for a variable or expression

```
a = 1
println(typeof(a))

```

---

#### keys

get the keys of a map

---

#### range

range a int64 array

---

#### remove

remove one or several items from an array

```
remove(arrayA, startIndexA, endIndexA)
```

---

#### exit

the same as os.Exit(1), used to terminate\exit the whole script running

---

#### deepClone

deep copy a struct variable and generate a new one

usage: 

```
person1 = make(struct {
	Name string,
	Age int
})

person1.Name = "John"
person1.Age = 20

pl("%#v", person1)

person2 = person1

person2.Name ="Tom"

pv("person1")
pv("person2")

p3 = deepClone(&person1)
p3 = *p3
p3.Name = "abc"
pv("person1")
pv("person2")
pv("p3")
```

---

#### deepCopy

deep copy a struct variable to another one

---

#### getClipText

get clipboard text

---

#### setClipText

set clipboard text

---

#### runScript

runScript(scriptA, modeA, argsA ...)

```
modeA == "" || modeA == "1" || modeA == "new" 
```

run Anko script in a new VM

```
modeA == "2" || modeA == "current"
```

run Anko script in a current VM

```
modeA == "3" || modeA == "js"
```

run JavaScript

```
modeA == "5" || modeA == "tg"
```

run Tengo script

---

#### systemCmd

run a system command

```
systemCmd(cmdA, argsA ...)
```

#### newSSHClient

create a SSH client to run shell commands, upload/download file from a remote server. Thanks to melbahja and visit [here](https://github.com/melbahja/goph) to find more docs.

short examples:

```
clientT, errT = newSSHClient(hostName, port, userName, password)

outT, errT = clientT.Run(`ls -p; cat abc.txt`)

errT = clientT.Upload(`./abc.txt`, tk.Replace(tk.JoinPath(pathT, `abc.txt`), `\`, "/"))

errT = clientT.Download(`down.txt`, `./down.txt`)

```

find more in this example script [here](https://gitee.com/topxeq/gox/tree/master/scripts/ssh.gox)

---


#### edit

the same as gui.EditFile, or the command-line switch "-edit", used to start an embedded code editor for Gox files.

```
edit("") // will open an editor for a new Gox code file

```

```
edit("d:\\tmp\\basic.gox") // will open an editor for the specific Gox code file

```

---

#### gui.GetConfirm

show a confirm dialog, return true or false

---

#### gui.LoadFont

for loading font for GUI display, example:

```
var gui = import("gui")

...

setVar("Font", "c:/Windows/Fonts/simsun.ttc")
setVar("FontRange", "COMMON")
setVar("FontSize", "36")

mainWindow = gui.NewMasterWindow("简易计算器", 400, 200, gui.MasterWindowFlagsNotResizable, gui.LoadFont)

```

see the example script [here](https://gitee.com/topxeq/gox/raw/master/scripts/testguic.gox).

---

#### gui.SelectFile

xxxxxxx

---

#### gui.SelectSaveFile

xxxxxxx

---

#### gui.SelectDirectory

xxxxxxx

---

#### gui.EditFile

the same as edit function, or command-line switch "-edit"

---



## Development

- Feel free to modify the source code to make a custom version of Gox. You can remove some packages not necessary for you, or add some Golang or third-party packages into it.

- The script [here](https://github.com/topxeq/gox/blob/master/scripts/generateImport.gox) is used for developers to add imported libraries to Gox.

The usage is as below:

``` 
gox generateImport.gox -file=c:\goprjs\src\package1\package1.go -package=package1 > a.txt
```

and you will got something like:

```
		"StartsWith":                          reflect.ValueOf(tk.StartsWith),
		"StartsWithIgnoreCase":                reflect.ValueOf(tk.StartsWithIgnoreCase),

```

then add it to the Gox source file in the code block to import variables/functions from the package,

```
	env.Packages["package1"] = map[string]reflect.Value{
		"StartsWith":                  reflect.ValueOf(package1.StartsWith),
		"StartsWithIgnoreCase":                  reflect.ValueOf(package1.StartsWithIgnoreCase),
    }
```

- The script [here](https://github.com/topxeq/gox/blob/master/scripts/generategoxc.gox) is used to generate command-line version and the script [here](https://github.com/topxeq/gox/blob/master/scripts/generategoxt.gox) is used to generate Gox tiny version.

- Add the following flag while building on Windows will generate a version without command-line console(for the goxg verison).

> go install -ldflags="-H windowsgui"

- Add the following flags while building will generate a more small-sized version without symbol and other debug information embedded in the executable.

> go install -ldflags="-s -w"

