module github.com/topxeq/gox

go 1.14

require (
	github.com/AllenDang/giu v0.5.0
	github.com/TheTitanrain/w32 v0.0.0-20200114052255-2654d97dbd3d // indirect
	github.com/ajstarks/svgo v0.0.0-20200725142600-7a3c8b57fecb // indirect
	github.com/beevik/etree v1.1.1-0.20200718192613-4a2f8b9d084c // indirect
	github.com/denisenkom/go-mssqldb v0.9.0
	github.com/go-gl/glfw/v3.3/glfw v0.0.0-20201108214237-06ea97f0c265 // indirect
	github.com/go-sql-driver/mysql v1.5.0
	github.com/godror/godror v0.23.1
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/mattn/go-sqlite3 v2.0.3+incompatible
	github.com/sciter-sdk/go-sciter v0.5.1
	github.com/spf13/afero v1.5.1 // indirect
	github.com/sqweek/dialog v0.0.0
	github.com/stretchr/objx v0.3.0 // indirect
	github.com/topxeq/imagetk v0.0.0-20210112052041-d3bf39e7174f // indirect
	github.com/topxeq/qlang v0.0.0
	github.com/topxeq/sqltk v0.0.0 // indirect
	github.com/topxeq/tk v0.0.0
	github.com/webview/webview v0.0.0-20210330151455-f540d88dde4e // indirect
	golang.org/x/text v0.3.5 // indirect
	gonum.org/v1/plot v0.8.2-0.20210109212805-a636e72ce5af // indirect
	gopkg.in/ini.v1 v1.62.0 // indirect
)

replace github.com/360EntSecGroup-Skylar/excelize v1.4.1 => github.com/360EntSecGroup-Skylar/excelize/v2 v2.3.2

replace github.com/360EntSecGroup-Skylar/excelize/v2 v2.3.2 => github.com/360EntSecGroup-Skylar/excelize v1.4.1

replace github.com/topxeq/tk v0.0.0 => ../tk

replace github.com/topxeq/xmlx v0.2.0 => ../xmlx

replace github.com/topxeq/sqltk v0.0.0 => ../sqltk

replace github.com/topxeq/qlang v0.0.0 => ../qlang

replace github.com/sciter-sdk/go-sciter v0.5.1 => ../../sciter-sdk/go-sciter

replace github.com/sqweek/dialog v0.0.0 => ../../sqweek/dialog
