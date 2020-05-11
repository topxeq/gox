// text1 used to hold the string value of the text input
// notice: text1 is a pointer
text1 = new(string)

onButton1Click = fn() {
	// evaluate the expression in the text input
	rs = eval(*text1)

	// set the result back into the text input
	setValue(text1, rs)//string(rs)
}

// close the window, also terminate the application
onButton2Click = fn() {
	exit()
}

// main window loop
loop = fn() {

	// set the layout of GUI
	layoutT = []gui.Widget{
		gui.Label("Enter an expression."), 
		gui.InputText("", 0, text1), 

		// widgets in line layout is aligned left to right 
		gui.Line(gui.Button("Calculate", onButton1Click), 
			gui.Button("Close", onButton2Click)),
	}

	gui.SingleWindow("Calculator", layoutT)
}

// setup the title, size (width and height, 400*200), style and font-loading function of main window, 
mainWindow = gui.NewMasterWindow("Calculator", 400, 200, gui.MasterWindowFlagsNotResizable, nil)

// show the window and start the message loop
gui.LoopWindow(mainWindow, loop)
