![Gox](https://gitee.com/topxeq/gox/raw/master/docs/goxlogo2xs.jpg)

_Gox语言的中文简介可以从[这篇文章](https://mbd.baidu.com/newspage/data/landingshare?pageType=1&isBdboxFrom=1&context=%7B%22nid%22%3A%22news_9677952217244209303%22%2C%22sourceFrom%22%3A%22bjh%22%7D)开始看起，或者[这里](https://www.jianshu.com/nb/44919205)看更多的开发指南。_

# gox

**Notice: from version 0.986a, Goxlang removed some script-engine support for smaller size. Only Qlang-engine is kept. And from version 0.998a, [Sciter](https://sciter.com) is recommended for cross-platform GUI programming in Gox. Although it requires a dynamic library file, but it's worthy.**

Gox (or Goxlang) is a free, open-source script language or a interpreter written by Golang. It's based on [Qlang](https://github.com/qiniu/qlang), with some improvement. The syntax of Gox language is very similar to Golang.

Golang is not required to be installed. Gox is only in one executable file, green and fast.

And thanks to [Sciter](https://sciter.com) and [Go-Sciter](https://github.com/sciter-sdk/go-sciter), which enables Gox to provide a modern GUI programming ability, and it's cross-platform, native, no dependencies and convenient, only an additional library will needed. Even more, Gox has an code editor embedded, so the external text editor may not be required for small piece of code. Note: this GUI library file is packed within the zip file in Windows release, for Linux/Mac, please see the instructions in Sciter website.

And also many thanks to [Govcl](https://github.com/ying32/govcl) written by ying32, which enables Gox to provide GUI programming APIs based on the free Lazarus LCL library. It's from VCL library and very useful especially for experienced Delphi/VCL programmers. Since this library doesn't require OPENGL, it's more compatible in some server-side operating systems. It only requires a single library file(dll in Windows or so in Linux and dylib in Mac OS).

If you do not want to use Giu GUI programming, try compiling from the source code and add "-tags nogiugui" command-line parameters in "go build" step.

Gox supports password-protect your source code file, it's also a major difference from most script/interpretive language.

## 1. Installation

Download the latest Gox executable file from the [official website](http://topget.org/gox.html)(recommended) or [Github](https://github.com/topxeq/gox/releases) or [Gitee](https://gitee.com/topxeq/gox/releases) (due to the speed issue, official website and Gitee are recommended). Then put it into a directory in the system path. If you download the zip file, extract it first.

A command-line only version of gox (named goxc, no GUI features) is also available. This version will be more fast, and of course smaller in file size.

Since the more features added makes the Gox executable became very large, the Tiny Gox Version is now available to download(not updated frequently). Some features are removed from Gox tiny version, such as OpenGL GUI, drawing, NoSQL, some database drivers(Oracle, MS-SQL), etc.

Goxg version is used to run GUI only programs, it will not open a CMD console.

## 2.usage

### Check the version.

> gox -version

### Show help.

> gox -h

### Run a script file.

> gox test.gox


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

### Run with a magic number(acquired from somewhere)

> gox 18

Gox will try to download the script associated with the magic number and then run it.

MNS(Magic-Number-Script) can also be executed in the shell mode, i.e.

> gox
> \>magic(18) 

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

### Init GUI programming environment

> gox -initgui

Download Sciter and LCL dll files in the same folder with Gox application path, now only in Windows platform.

### Open a simple editor to edit Gox code

> gox -edit

### Open a simple editor to edit specific Gox code file

> gox -edit d:\tmp\basic.gox

the screenshot of Gox editor is as below,

![Gox Editor](https://gitee.com/topxeq/gox/raw/master/docs/goxeditor2.png)

Although Gox provides a simple code editor, the editors with more powerful functions are recommended, such as Visual Studio Code. Currently, you can set the .gox files the same syntax highlight scheme as Golang files.

### Batch run

First create a file(better named with .gxb suffix), content as below for example:

```
basic.gox
d:\scripts\next.gox -inputvalue=1 -verbose
18
```

Each line in Gox batch file could be considered as a common Gox command line but the first paramter(Gox executable path) is omitted.

There are 3 ways to inform Gox to know it's a batch file to run:

- run with -batch switch

  > gox -batch batchFile.txt

- use the proper file extension(.gxb)

  > gox batchFile.gxb

- add // gxb to the batch file as first line

	```
	// gxb
	basic.gox
	d:\scripts\next.gox -inputvalue=1 -verbose
	18
	```

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

println("The whole command-line: ", os.Args)
println("The whole command-line without executable: ", argsG)

lenT = len(argsG)

if lenT > 0 {
    printfln("The first command-line element is: %v", argsG[0])
}

if lenT > 1 {
    printfln("The sencod command-line element is: %v", argsG[1])
}

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

In Gox with Qlang engine, most of the standard Golang packages are supported and will be imported by default, just use them directly.

```
fmt.Println("abc")

fmt.Printf("%v\n", os.Args)

```

Along with most of the core Golang libraries, the "tk" package([here](https://github.com/topxeq/tk)) is often used. The "tk" package provides many useful functions. See example scripts for more. And refer to the docs [here](https://pkg.go.dev/github.com/topxeq/tk?tab=doc).

Refer to the documents of these Golang packages for the detailed usage.


## 4. More Topics and Sample Scripts

### Basic Gox script:

```

// do simple add operation
x = 1.2
y = x + 1

println(x+y)

```

### Basic script with GUI

A simple calculator with GUI(OPENGL is required).

```
text1 = new(string)

onButton1Click = func() {
	// evaluate the expression in the text input
	rs = eval(*text1)

	println(rs)

	// set the result back into the text input
	setValue(text1, rs)
}

// close the window, also terminate the application
onButton2Click = func() {
	os.Exit(1)
}

// main window loop
loop = func() {

	// set the layout of GUI
	layoutT := []gui.Widget{
		gui.Label("Enter an expression."),
		gui.InputText("", 0, text1),

		// widgets in line layout is aligned left to right
		gui.Line(gui.Button("Calculate", onButton1Click),
			gui.Button("Close", onButton2Click)),
	}

	gui.SingleWindow("Calculator", layoutT)
}

// text1 used to hold the string value of the text input
// notice: text1 is a pointer
// setup the title, size (width and height, 400*200), style and font-loading function of main window,
mainWindow := gui.NewMasterWindow("Calculator", 400, 200, gui.MasterWindowFlagsNotResizable, nil)

// show the window and start the message loop
gui.LoopWindow(mainWindow, loop)

```

The screen shot while running the script is like,

![Calculator](https://gitee.com/topxeq/gox/raw/master/docs/calculatorss.png)

### Basic script with GUI by LCL library

The script below acts almost the same as the script above, but written with LCL library support, and OPENGL is not required, so it's more compatible in the server-side environment.

```
errT = lcl.InitLCL()

if errT != nil {
	tk.Plerr(errT)
	return
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

onFromDestory = fn(sender) {
	println("Form Destroyed.")
}

mainForm.SetOnDestroy(lcl.NewTNotifyEvent(onFromDestory))

label1 = lcl.NewLabel(mainForm)
label1.SetParent(mainForm)
label1.SetLeft(10)
label1.SetTop(10)
label1.Font().SetName("Arial")
label1.Font().SetSize(18)

label1.SetCaption("Enter an expression")

onEdit1KeyUp = fn(sender, key, shift) {
	println("onEdit1KeyUp:", sender, *key, shift)
}

edit1 = lcl.NewEdit(mainForm)
edit1.SetParent(mainForm)
edit1.SetBounds(10, 48, 200, 32)
edit1.Font().SetSize(11)
edit1.SetOnKeyUp(lcl.NewTKeyEvent(onEdit1KeyUp))

onClick1 = fn(objA) {
	rs = edit1.Text()
	edit1.SetText(eval(rs))
}

f1 = lcl.NewTNotifyEvent(onClick1)

button1 = lcl.NewButton(mainForm)
button1.SetParent(mainForm)
button1.SetLeft(20)
button1.SetTop(90)
button1.SetCaption("Go")
button1.SetOnClick(f1)

onClick2 = fn(sender) {
	application.Terminate()
}

button2 = lcl.NewButton(mainForm)
button2.SetParent(mainForm)
button2.SetLeft(110)
button2.SetTop(90)
button2.SetCaption("Close")
button2.SetOnClick(lcl.NewTNotifyEvent(onClick2))


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

First, since Gox is based on Qlang and written by Golang, most of the core libraries of Golang will be available. So try to use the modules from Golang(but Golang installation is not required), and refer to the Golang documents([here](https://golang.org/doc/), [here](https://pkg.go.dev/) or [here](http://docscn.studygolang.com/)). In addition, you can browse Qlang's, [Govcl](https://gitee.com/ying32/govcl/wikis/pages)'s and [Giu](https://github.com/AllenDang/giu)'s documents.

Most of the standard Golang packages are available include: bytes, encoding/json, errors, flag, fmt, image, image/color, image/draw, image/jpg, image/png, io, io.ioutil, log, math, math/big, math/rand, net, net/http, net/http/cookiejar, net/url, os, os/exec, os/signal, path, path/filepath, regexp, runtime, sort, strconv, strings, sync, time.

Library tk (github.com/topxeq/tk) is the most frequently used package in Gox, the documents are [here](https://godoc.org/github.com/topxeq/tk) or [here](https://pkg.go.dev/github.com/topxeq/tk).

Other libraries embedded in Gox include: 

- gonum.org/v1/plot(scientific drawing and charts, use it as gonumorg_v1_plot), and with gonumorg_v1_plot_plotter, gonumorg_v1_plot_plotutil, gonumorg_v1_plot_vg
- github.com/domodwyer/mailyak(send mail via SMTP)
- github.com/360EntSecGroup-Skylar/excelize(Excel file processing)
- github.com/fogleman/gg(basic drawing)
- github.com/dgraph-io/badger(NoSQL DB)
- github.com/topxeq/govcl/vcl(derived from ying32's Govcl, GUI programming with LCL library)
- github.com/AllenDang/giu(OpenGL GUI)
- github.com/sciter-sdk/go-sciter(GUI programming with Sciter, HTML+CSS+TiScript+Golang)
- github.com/topxeq/blink(control a chronium WebView, Windows only)
- github.com/topxeq/sqltk(general SQL operations, enable access to Oracle, MySQL, MSSQLServer, SQLite databases)
- github.com/topxeq/imagetk(various image-processing funtions)
- github.com/beevik/etree(XML processing)
- github.com/stretchr/objx(another JSON library)
- github.com/kbinani/screenshot(cross-platform, multi-display screenshot)
- github.com/topxeq/doc2vec/doc2vec(document to vector engine and training functions)
- github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi(send Aliyun SMS)
- github.com/topxeq/afero(a memory file system)

Browse their Github pages for usage and documents, and of course the examples in the example directory.

To use an embedded package, see the examples below:

**directly use**

```
tk.Pl("1+3=%v", 1+3)
```

or

```
github_topxeq_tk.Pl("1+3=%v", 1+3)
```

Attention: To refer a 3rd-party package, '.com' is always omitted for the domain name such as github.com, and the same is '.org' for gonum.org. And the slash('/') is replaced by underscore('_'), other special characters will be removed. For some packages used frequently(such as github.com/topxeq/tk), we will give it a convenient short-name(tk), but not all. You can define a short name for the packages that will be used.

**define a short name**

```
excel = github_360EntSecGroupSkylar_excelize

f, err = excel.OpenFile(`d:\tmpx\test.xlsx`)

...

```


However, Gox provides some convenient global variables and functions decribed as below.

### Variables

#### argsG

get global variable

The global value "argsG" could be used for retrieve command-line arguments, and the first element(the Gox executable) is usually removed. If you need the whole command-line, use os.Args instead.

---

#### scriptPathG

holds the full script path(include the file name), for network scenes, will be ""(empty string)

#### versionG

holds the current version of Gox

### Functions

Note: some functions may exist or not in different script engine, and may have some slight differences.

---

#### pass

A function doing nothing. Usually used at the end of a script to ensure return nothing(so that will print nothing by the VM).

---

#### defined

check if a variable is defined

---

#### eval

evaluate an expression and return the result

---

#### typeof/typeOf/kindOf

return the string representation of the type for a variable or expression

```
a = 1
println(typeof(a))

```

---

#### remove

remove one or several items from an array

```
remove(arrayA, startIndexA, endIndexA)
```

---

#### print/pr

the same as fmt.Print

---

#### println/pln

the same as fmt.Println

---

#### printf/prf

the same as fmt.Printf

---

#### printfln/pl

the same as fmt.Printf but add a new-line character at the end

---

#### sprintf

the same as fmt.Sprintf

---

#### fprintln/fprintf

the same as fmt.Fprintln/fmt.Fprintf

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

#### plvsr

the same as pl("%#v", v), but for various arguments

---

#### plerr

a convenient way to print an error value

---

#### exit

the same as os.Exit(1), used to terminate\exit the whole script running

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

### getValue

get a value referenced by a pointer, used in Qlang engine.

```
s = new(string)

// *s = "abc" is not correct in Qlang engine

setValue(s, "abc")

println(*s)  // use * for dereference a value from pointer is allowed

v = getValue(s)

println(v) // will be "abc"

```

---

#### setVar

set global variable

---

#### getVar

get global variable

---

#### bitXor

bitwise XOR operation, since in Qlang engine, ^ is used for get address/pointer of a variable(like & in other engine),
so the origin bitwise XOR operator in Golang is used and we will use bitXor function instead.

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

#### checkErrf

the same as tk.CheckErrf, check an error, if not nil, terminate the program and printf something

---

#### checkErrStrf

the same as tk.CheckErrStrf, check a TXERROR string, if it's a TXERROR string, terminate the program and printf something

---

#### fatalf

printf then terminate the program

---

#### isErrStr

check if is TXERROR string (which starts with TXERROR:)

---

#### errStr

generate a TXERROR string (which starts with TXERROR:)

---

#### errStrf

generate a TXERROR string with format like printf, the same as tk.ErrStrF

---

#### getErrStr

remove the prefix(TXERROR:) of a TXERROR string, return the result

---

#### errf

the same as tk.Errf, generate an error

---


#### getInput(deprecated)

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

#### keys

get the keys of a map

---

#### range

range an array, actually a keyword, not a function

```
a = [1, 2, "list"]

for i, v = range a {
	println(a)
}

```

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

#### trim

trim a string

---

#### run

execute a script file(in a new created VM)

> run(argsA...)

---

#### runCode

execute a script string(in a new created VM, argsG will hold the arguments)

> runCode(codeA string, argsA ...string)

---

#### runScript

runScript(scriptA, modeA, argsA ...)

```
modeA == "" || modeA == "1" || modeA == "new" 
```

run script in a new VM

any other modeA will run it as a system command.

---

#### magic

execute a script with a magic number(in a new created VM, argsG will hold the arguments)

> magic(numberA, argsA...)

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

#### getParameter

the same as tk.GetParameterByIndexWithDefaultValue

---

#### getSwitch

the same as tk.GetSwitchWithDefaultValue

---

#### switchExists

the same as tk.IfSwitchExistsWhole

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

- The command [here](https://github.com/topxeq/qlang/cmd/qexport) is used for developers to add imported libraries to Gox.

The usage is as below:

``` 
qexport github.com/topxeq/sqltk

```

and add them to the qlang/lib directory(look at the files inside for reference):


- The script [here](https://github.com/topxeq/gox/blob/master/scripts/generategoxc.gox) is used to generate command-line version and the script [here](https://github.com/topxeq/gox/blob/master/scripts/generategoxt.gox) is used to generate Gox tiny version.

- Add the following flag while building on Windows will generate a version without command-line console(for the goxg verison).

> go install -ldflags="-H windowsgui"

- Add the following flags while building will generate a more small-sized version without symbol and other debug information embedded in the executable.

> go install -ldflags="-s -w"

