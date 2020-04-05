# gox
Gox is a script language or a interpreter written by Golang. It's based on [Anko](https://github.com/mattn/anko) and [Goja](https://github.com/dop251/goja), with some improvement. As a script runner(or interpreter), Gox supports various script languages such as txScript, Javascript, Anko, ...

Golang is not required to be installed. Gox is only in one executable file, green and fast.

## Installation

Download the latest Gox executable file from [here](https://github.com/topxeq/gox/releases). Then put it into a directory in the system path. If you download the zip file, extract it first.

## usage

Check the version.

> gox -version

Show help.

> gox -h

Run a script file.

> gox test.gox

Run several script files in consequence, use "-m" switch to enable multi-scripts mode.

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

Start the interpreter in REPL mode.

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

## User/developer guide

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

## Sample script

Sample Javascript file:

```
var a = 5 + 12.5;

goPrintf("this is %v + %v = %v\n", 5, 12.5, a);

console.log(5, 12.5, a);

goPrintln(19, 33, 19 * 33);

goPrintfln("The random number：%v", goGetRandomInt(20));

```

Basic Gox script:

```

// do simple add operation
x = 1.2
y = x + 1

println(x+y)

```

## More examples

Browse the example scripts [here](https://github.com/topxeq/gox/tree/master/scripts)

## Reference

First, since Gox is based on Anko and written by Golang, most of the core libraries of Golang will be available. So try to import the modules from Golang(but Golang installation is not required), and refer to the Golang documents.

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


#### exit

the same as os.Exit(1), used to terminate\exit the whole script running


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