function calPi(pointCount) {
	var inCircleCount = 0

	for (var i = 0; i < pointCount; i++) {
		var x = Math.random()
		var y = Math.random()

		if (x*x+y*y < 1) {
			inCircleCount++
		}
	}

	var Pi = (4.0 * inCircleCount) / pointCount

	return Pi
}


var countStrT = getVar("Count");

println("Count:", countStrT);

var countT = 1000;

if (!countStrT.startsWith("TXERROR:")) {
	countT = parseInt(countStrT);
}

if (countT == NaN) {
	countT = 1000;
}

resultG = calPi(countT)

