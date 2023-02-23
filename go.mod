module github.com/topxeq/gox

go 1.14

require (
	github.com/TheTitanrain/w32 v0.0.0-20200114052255-2654d97dbd3d // indirect
	github.com/ajstarks/svgo v0.0.0-20200725142600-7a3c8b57fecb // indirect
	github.com/beevik/etree v1.1.1-0.20200718192613-4a2f8b9d084c // indirect
	github.com/denisenkom/go-mssqldb v0.12.3
	github.com/go-sql-driver/mysql v1.7.0
	github.com/godror/godror v0.36.0
	github.com/gopherjs/gopherjs v0.0.0-20220221023154-0b2280d3ff96 // indirect
	github.com/jchv/go-webview2 v0.0.0-20221223143126-dc24628cff85
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/mattn/go-sqlite3 v2.0.3+incompatible
	github.com/sijms/go-ora/v2 v2.5.25
	github.com/topxeq/charlang v0.0.0-20220722003130-49bba2664ad6
	github.com/topxeq/dialog v0.0.0-20211124003827-315c3296b533
	github.com/topxeq/dlgs v0.0.0-20220223083937-4d3036aff547
	github.com/topxeq/go-sciter v0.0.0-20221010031453-76f65a41d04f
	github.com/topxeq/imagetk v0.0.0-20210112052041-d3bf39e7174f // indirect
	github.com/topxeq/qlang v0.0.0
	github.com/topxeq/sqltk v0.0.0-20230223005953-f9932d23950c
	github.com/topxeq/tk v1.0.1
	github.com/topxeq/xie v0.0.0
	gonum.org/v1/plot v0.8.2-0.20210109212805-a636e72ce5af // indirect
	gopkg.in/check.v1 v1.0.0-20200902074654-038fdea0a05b // indirect
)

//replace github.com/360EntSecGroup-Skylar/excelize v1.4.1 => github.com/360EntSecGroup-Skylar/excelize/v2 v2.3.2

// replace github.com/360EntSecGroup-Skylar/excelize v1.4.1 => github.com/xuri/excelize/v2 v2.4.1

// replace github.com/360EntSecGroup-Skylar/excelize/v2 v2.3.2 => github.com/360EntSecGroup-Skylar/excelize v1.4.1

replace github.com/topxeq/tk v1.0.1 => ../tk

// replace github.com/topxeq/xmlx v0.2.0 => ../xmlx

// replace github.com/topxeq/sqltk v0.0.0 => ../sqltk

replace github.com/topxeq/qlang v0.0.0 => ../qlang

// replace github.com/topxeq/xie/go/xie v0.0.0 => ../xie/go/xie
replace github.com/topxeq/xie v0.0.0 => ../xie

// replace github.com/topxeq/text v0.0.0 => ../text

// replace github.com/topxeq/charlang v0.0.0 => ../charlang

// replace github.com/topxeq/dialog v0.0.0 => ../../topxeq/dialog

// replace github.com/topxeq/goph v0.0.0 => ../goph

// replace github.com/topxeq/go-sciter v0.0.0 => ../go-sciter

// replace github.com/topxeq/gods v0.0.0 => ../gods
