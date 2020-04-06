# gox
Gox is a free, open-source script language or a interpreter written by Golang. It's based on [Anko](https://github.com/mattn/anko) and [Goja](https://github.com/dop251/goja), with some improvement. As a script runner(or interpreter), Gox supports various script languages such as txScript, Javascript, Anko, ...

Golang is not required to be installed. Gox is only in one executable file, green and fast.

And thanks to [Giu](https://github.com/AllenDang/giu), which enable Gox to provide a modern GUI programming ability, and it's cross-platform, native, no dependencies and convenient.

Gox supports password-protect your source code file, it's also a major difference from most script/interpretive language.

## 1. Installation

Download the latest Gox executable file from the [official website](http://gox.topget.org/) or [Github](https://github.com/topxeq/gox/releases). Then put it into a directory in the system path. If you download the zip file, extract it first.

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
> gox -m basic.gox test.js
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

And all the switches in Gox should be like "-type=code", not "-type code".

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

### Run example script from Github directly

> gox -example basic.gox

Gox will download the example script directly from Github and run it.

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

### Run

## 3. User/developer guide

### File encoding

All the script file is better in UTF-8 encoding, and its line-endings are better to use "\n" or "\r\n".

### command-line parameters and swithes

The global value "argsG" could be used for retrieve command-line arguments, and the first element(the Gox executable) is removed. If you need the whole command-line, use os.Args instead.

An example for command-line handling is as below([source code](https://github.com/topxeq/gox/blob/master/scripts/commandLine.gox)),

```
// test command-line functions
// for example: gox scripts\commandLine.gox abc -file=a.txt

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

In Gox, the packages can be imported any time, but remember to assign it to a variable, such as,

```
var fmt = import("fmt")

fmt.Println("abc")

os = import("os")

fmt.Printf("%v\n", os.Args)

```

Along with most of the core Golang libraries, the "tk" package([here](https://github.com/topxeq/tk)) is often used. The "tk" package provides many useful functions. See example scripts for more.

Refer to the documents of these Golang packages for the detailed usage.

### Data type conversion

some "to" functions could be used as below,

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
λ gox scripts\dataTypeConversion.gox
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

> toString, toBool(and tryToBool which returns the result like (bool, error)), toFloat64/tryToFloat64, toInt64/tryToInt64, toInt/tryToInt,

## 4. More Topics and Sample Script

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

### Base script with GUI

A simple calculator with GUI

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

![Calculator](https://github.com/topxeq/gox/blob/master/docs/calculatorss.png)


## 5. More examples

Browse the example scripts [here](https://github.com/topxeq/gox/tree/master/scripts)

## 6. Reference

First, since Gox is based on Anko and written by Golang, most of the core libraries of Golang will be available. So try to import the modules from Golang(but Golang installation is not required), and refer to the Golang documents. In addition, you can browse Anko's, Goja's and Giu's documents.

However, Gox provides some convenient global variables and functions decribed as below.

### Functions

---


#### getVar

get global variable

---

#### setVar

set global variable

---


#### print

the same as fmt.Print

---

#### printf

the same as fmt.Printf

---


#### println/pl

the same as fmt.Println

---


#### printfln/pfl

the same as fmt.Printf but add a new-line character at the end

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


#### eval

evaluate an expression and return the result

---

#### typeof

return the string representation of the type for a variable or expression

```
a = 1
println(typeof(a))

```

---


#### exit

the same as os.Exit(1), used to terminate\exit the whole script running

---

#### gui.loadFont

for loading font for GUI display, example:

```
var gui = import("gui")

...

setVar("Font", "c:/Windows/Fonts/simsun.ttc")
setVar("FontRange", "COMMON")
setVar("FontSize", "36")

mainWindow = gui.NewMasterWindow("简易计算器", 400, 200, gui.MasterWindowFlagsNotResizable, gui.loadFont)

```

see the example script [here](https://github.com/topxeq/gox/blob/master/scripts/testguic.gox).

---



## Development

The script [here](https://github.com/topxeq/gox/blob/master/scripts/generateImport.gox) is used for developers to add imported libraries to Gox.

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