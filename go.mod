module github.com/topxeq/gox

go 1.14

require (
	github.com/TheTitanrain/w32 v0.0.0-20200114052255-2654d97dbd3d // indirect
	github.com/ajstarks/svgo v0.0.0-20200725142600-7a3c8b57fecb // indirect
	github.com/beevik/etree v1.1.1-0.20200718192613-4a2f8b9d084c // indirect
	github.com/denisenkom/go-mssqldb v0.10.0
	github.com/go-sql-driver/mysql v1.6.0
	github.com/godror/godror v0.25.3
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/mattn/go-sqlite3 v2.0.3+incompatible
	github.com/spf13/afero v1.5.1 // indirect
	github.com/stretchr/objx v0.3.0 // indirect
	github.com/topxeq/charlang v0.0.0
	github.com/topxeq/dialog v0.0.0
	github.com/topxeq/dlgs v0.0.0-20211122010615-d49596e82836
	github.com/topxeq/go-sciter v0.0.0
	github.com/topxeq/gods v0.0.0 // indirect
	github.com/topxeq/imagetk v0.0.0-20210112052041-d3bf39e7174f // indirect
	github.com/topxeq/qlang v0.0.0
	github.com/topxeq/sqltk v0.0.0
	github.com/topxeq/tk v0.0.0
	golang.org/x/crypto v0.0.0-20211215153901-e495a2d5b3d3 // indirect
	golang.org/x/sys v0.0.0-20211216021012-1d35b9e2eb4e // indirect
	gonum.org/v1/plot v0.8.2-0.20210109212805-a636e72ce5af // indirect
	gopkg.in/check.v1 v1.0.0-20200902074654-038fdea0a05b // indirect
	gopkg.in/ini.v1 v1.62.0 // indirect
	gopkg.in/yaml.v2 v2.3.0 // indirect
)

//replace github.com/360EntSecGroup-Skylar/excelize v1.4.1 => github.com/360EntSecGroup-Skylar/excelize/v2 v2.3.2

// replace github.com/360EntSecGroup-Skylar/excelize v1.4.1 => github.com/xuri/excelize/v2 v2.4.1

// replace github.com/360EntSecGroup-Skylar/excelize/v2 v2.3.2 => github.com/360EntSecGroup-Skylar/excelize v1.4.1

replace github.com/topxeq/tk v0.0.0 => ../tk

replace github.com/topxeq/xmlx v0.2.0 => ../xmlx

replace github.com/topxeq/sqltk v0.0.0 => ../sqltk

replace github.com/topxeq/qlang v0.0.0 => ../qlang

replace github.com/topxeq/text v0.0.0 => ../text

replace github.com/topxeq/charlang v0.0.0 => ../charlang

replace github.com/topxeq/dialog v0.0.0 => ../../topxeq/dialog

replace github.com/topxeq/goph v0.0.0 => ../goph

replace github.com/topxeq/go-sciter v0.0.0 => ../go-sciter

replace github.com/topxeq/gods v0.0.0 => ../gods
