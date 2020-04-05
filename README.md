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

Run several script files in consequence.

> gox test.gox script1.js start.ank last.txs

the result should be like below,

```
> gox basic.gox test.js
3.4000000000000004
this is 5 + 12.5 = 17.5
5 12.5 17.5
19 33 627
The random number is：1
```

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