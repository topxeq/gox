package main

import (
	"bufio"
	"io"
	"math/rand"
	"os"
	"reflect"
	"strings"
	"time"

	"fmt"

	"github.com/dop251/goja"
	"github.com/mattn/anko/core"
	"github.com/mattn/anko/env"
	_ "github.com/mattn/anko/packages"
	"github.com/mattn/anko/parser"
	"github.com/mattn/anko/vm"
	"github.com/topxeq/tk"

	"github.com/AllenDang/giu"
	g "github.com/AllenDang/giu"
	"github.com/AllenDang/giu/imgui"
	// "github.com/sqweek/dialog"
)

// var inG interface{}
// var outG interface{}

var variableG = make(map[string]interface{})

var jsVMG *goja.Runtime = nil
var ankVMG *env.Env = nil

func getVar(nameA string) interface{} {
	return variableG[nameA]
}

func setVar(nameA string, valueA interface{}) {
	variableG[nameA] = valueA
}

func loadFont() {
	fonts := giu.Context.IO().Fonts()

	rangeVarT := getVar("FontRange")

	ranges := imgui.NewGlyphRanges()

	builder := imgui.NewFontGlyphRangesBuilder()

	if rangeVarT == nil {
		builder.AddRanges(fonts.GlyphRangesDefault())
	} else {
		rangeStrT := rangeVarT.(string)
		if rangeStrT == "" || rangeStrT == "COMMON" {
			builder.AddRanges(fonts.GlyphRangesChineseSimplifiedCommon())
		} else if rangeStrT == "FULL" {
			builder.AddRanges(fonts.GlyphRangesChineseFull())
		} else {
			builder.AddText(rangeStrT)
		}
	}

	builder.BuildRanges(ranges)

	fontPath := "c:/Windows/Fonts/simhei.ttf"

	fontVarT := getVar("Font") // "c:/Windows/Fonts/simsun.ttc"

	if fontVarT != nil {
		fontPath = fontVarT.(string)
	}

	fontSizeStrT := "16"

	fontSizeVarT := getVar("FontSize")

	if fontSizeVarT != nil {
		fontSizeStrT = fontSizeVarT.(string)
	}

	fontSizeT := tk.StrToIntWithDefaultValue(fontSizeStrT, 16)

	// fonts.AddFontFromFileTTF(fontPath, 14)
	fonts.AddFontFromFileTTFV(fontPath, float32(fontSizeT), imgui.DefaultFontConfig, ranges.Data())
}

func importAnkPackages() {
	env.Packages["gui"] = map[string]reflect.Value{
		"NewMasterWindow":         reflect.ValueOf(g.NewMasterWindow),
		"SingleWindow":            reflect.ValueOf(g.SingleWindow),
		"Window":                  reflect.ValueOf(g.Window),
		"SingleWindowWithMenuBar": reflect.ValueOf(g.SingleWindowWithMenuBar),
		"WindowV":                 reflect.ValueOf(g.WindowV),

		"MasterWindowFlagsNotResizable": reflect.ValueOf(g.MasterWindowFlagsNotResizable),
		"MasterWindowFlagsMaximized":    reflect.ValueOf(g.MasterWindowFlagsMaximized),
		"MasterWindowFlagsFloating":     reflect.ValueOf(g.MasterWindowFlagsFloating),

		// "Layout":          reflect.ValueOf(g.Layout),

		"Label":              reflect.ValueOf(g.Label),
		"Line":               reflect.ValueOf(g.Line),
		"Button":             reflect.ValueOf(g.Button),
		"InvisibleButton":    reflect.ValueOf(g.InvisibleButton),
		"ImageButton":        reflect.ValueOf(g.ImageButton),
		"InputTextMultiline": reflect.ValueOf(g.InputTextMultiline),
		"Checkbox":           reflect.ValueOf(g.Checkbox),
		"RadioButton":        reflect.ValueOf(g.RadioButton),
		"Child":              reflect.ValueOf(g.Child),
		"ComboCustom":        reflect.ValueOf(g.ComboCustom),
		"Combo":              reflect.ValueOf(g.Combo),
		"ContextMenu":        reflect.ValueOf(g.ContextMenu),
		"Group":              reflect.ValueOf(g.Group),
		"Image":              reflect.ValueOf(g.Image),
		"InputText":          reflect.ValueOf(g.InputText),
		"InputInt":           reflect.ValueOf(g.InputInt),
		"InputFloat":         reflect.ValueOf(g.InputFloat),
		"MainMenuBar":        reflect.ValueOf(g.MainMenuBar),
		"MenuBar":            reflect.ValueOf(g.MenuBar),
		"MenuItem":           reflect.ValueOf(g.MenuItem),
		"PopupModal":         reflect.ValueOf(g.PopupModal),
		"OpenPopup":          reflect.ValueOf(g.OpenPopup),
		"CloseCurrentPopup":  reflect.ValueOf(g.CloseCurrentPopup),
		"ProgressBar":        reflect.ValueOf(g.ProgressBar),
		"Separator":          reflect.ValueOf(g.Separator),
		"SliderInt":          reflect.ValueOf(g.SliderInt),
		"SliderFloat":        reflect.ValueOf(g.SliderFloat),
		"HSplitter":          reflect.ValueOf(g.HSplitter),
		"VSplitter":          reflect.ValueOf(g.VSplitter),
		"TabItem":            reflect.ValueOf(g.TabItem),
		"TabBar":             reflect.ValueOf(g.TabBar),
		"Row":                reflect.ValueOf(g.Row),
		"Table":              reflect.ValueOf(g.Table),
		"FastTable":          reflect.ValueOf(g.FastTable),
		"Tooltip":            reflect.ValueOf(g.Tooltip),
		"TreeNode":           reflect.ValueOf(g.TreeNode),
		"Spacing":            reflect.ValueOf(g.Spacing),
		"Custom":             reflect.ValueOf(g.Custom),
		"Condition":          reflect.ValueOf(g.Condition),
		"ListBox":            reflect.ValueOf(g.ListBox),
		"DatePicker":         reflect.ValueOf(g.DatePicker),
		// "Widget":             reflect.ValueOf(g.Widget),
		"loadFont": reflect.ValueOf(loadFont),
	}

	var widget g.Widget

	env.PackageTypes["gui"] = map[string]reflect.Type{
		"Layout": reflect.TypeOf(g.Layout{}),
		// "Signal": reflect.TypeOf(&signal).Elem(),
		"Widget": reflect.TypeOf(&widget).Elem(),
	}

	env.Packages["tk"] = map[string]reflect.Value{
		"CreateTXCollection":                  reflect.ValueOf(tk.CreateTXCollection),
		"TXResultFromString":                  reflect.ValueOf(tk.TXResultFromString),
		"SetGlobalEnv":                        reflect.ValueOf(tk.SetGlobalEnv),
		"RemoveGlobalEnv":                     reflect.ValueOf(tk.RemoveGlobalEnv),
		"GetGlobalEnvList":                    reflect.ValueOf(tk.GetGlobalEnvList),
		"GetGlobalEnvString":                  reflect.ValueOf(tk.GetGlobalEnvString),
		"HasGlobalEnv":                        reflect.ValueOf(tk.HasGlobalEnv),
		"IsEmptyTrim":                         reflect.ValueOf(tk.IsEmptyTrim),
		"StartsWith":                          reflect.ValueOf(tk.StartsWith),
		"StartsWithIgnoreCase":                reflect.ValueOf(tk.StartsWithIgnoreCase),
		"StartsWithUpper":                     reflect.ValueOf(tk.StartsWithUpper),
		"StartsWithDigit":                     reflect.ValueOf(tk.StartsWithDigit),
		"Contains":                            reflect.ValueOf(tk.Contains),
		"ContainsIgnoreCase":                  reflect.ValueOf(tk.ContainsIgnoreCase),
		"EndsWith":                            reflect.ValueOf(tk.EndsWith),
		"EndsWithIgnoreCase":                  reflect.ValueOf(tk.EndsWithIgnoreCase),
		"Trim":                                reflect.ValueOf(tk.Trim),
		"TrimCharSet":                         reflect.ValueOf(tk.TrimCharSet),
		"InStrings":                           reflect.ValueOf(tk.InStrings),
		"GetSliceMaxLen":                      reflect.ValueOf(tk.GetSliceMaxLen),
		"FindFirstDiffIndex":                  reflect.ValueOf(tk.FindFirstDiffIndex),
		"FindSamePrefix":                      reflect.ValueOf(tk.FindSamePrefix),
		"IsErrorString":                       reflect.ValueOf(tk.IsErrorString),
		"GetErrorString":                      reflect.ValueOf(tk.GetErrorString),
		"GetErrorStringSafely":                reflect.ValueOf(tk.GetErrorStringSafely),
		"GenerateErrorString":                 reflect.ValueOf(tk.GenerateErrorString),
		"GenerateErrorStringF":                reflect.ValueOf(tk.GenerateErrorStringF),
		"ErrorStringToError":                  reflect.ValueOf(tk.ErrorStringToError),
		"Replace":                             reflect.ValueOf(tk.Replace),
		"StringReplace":                       reflect.ValueOf(tk.StringReplace),
		"SplitLines":                          reflect.ValueOf(tk.SplitLines),
		"SplitLinesRemoveEmpty":               reflect.ValueOf(tk.SplitLinesRemoveEmpty),
		"Split":                               reflect.ValueOf(tk.Split),
		"SplitN":                              reflect.ValueOf(tk.SplitN),
		"JoinLines":                           reflect.ValueOf(tk.JoinLines),
		"JoinLinesBySeparator":                reflect.ValueOf(tk.JoinLinesBySeparator),
		"EnsureValidFileNameX":                reflect.ValueOf(tk.EnsureValidFileNameX),
		"CreateString":                        reflect.ValueOf(tk.CreateString),
		"CreateStringSimple":                  reflect.ValueOf(tk.CreateStringSimple),
		"CreateStringWithObject":              reflect.ValueOf(tk.CreateStringWithObject),
		"CreateStringEmpty":                   reflect.ValueOf(tk.CreateStringEmpty),
		"CreateStringSuccess":                 reflect.ValueOf(tk.CreateStringSuccess),
		"CreateStringError":                   reflect.ValueOf(tk.CreateStringError),
		"CreateStringErrorF":                  reflect.ValueOf(tk.CreateStringErrorF),
		"CreateStringErrorFromTXError":        reflect.ValueOf(tk.CreateStringErrorFromTXError),
		"GenerateErrorStringTX":               reflect.ValueOf(tk.GenerateErrorStringTX),
		"GenerateErrorStringFTX":              reflect.ValueOf(tk.GenerateErrorStringFTX),
		"LoadStringTX":                        reflect.ValueOf(tk.LoadStringTX),
		"RegContains":                         reflect.ValueOf(tk.RegContains),
		"RegFindFirstTX":                      reflect.ValueOf(tk.RegFindFirstTX),
		"LoadDualLineListFromString":          reflect.ValueOf(tk.LoadDualLineListFromString),
		"RegReplace":                          reflect.ValueOf(tk.RegReplace),
		"RegFindAll":                          reflect.ValueOf(tk.RegFindAll),
		"RegFindFirst":                        reflect.ValueOf(tk.RegFindFirst),
		"RegMatch":                            reflect.ValueOf(tk.RegMatch),
		"Randomize":                           reflect.ValueOf(tk.Randomize),
		"GetRandomIntLessThan":                reflect.ValueOf(tk.GetRandomIntLessThan),
		"GetRandomInt64LessThan":              reflect.ValueOf(tk.GetRandomInt64LessThan),
		"GetRandomIntInRange":                 reflect.ValueOf(tk.GetRandomIntInRange),
		"GetRandomInt64InRange":               reflect.ValueOf(tk.GetRandomInt64InRange),
		"GenerateRandomString":                reflect.ValueOf(tk.GenerateRandomString),
		"NewRandomGenerator":                  reflect.ValueOf(tk.NewRandomGenerator),
		"ShuffleStringArray":                  reflect.ValueOf(tk.ShuffleStringArray),
		"GetRandomizeStringArrayCopy":         reflect.ValueOf(tk.GetRandomizeStringArrayCopy),
		"GetRandomizeIntArrayCopy":            reflect.ValueOf(tk.GetRandomizeIntArrayCopy),
		"GetRandomizeInt64ArrayCopy":          reflect.ValueOf(tk.GetRandomizeInt64ArrayCopy),
		"GetRandomSubDualList":                reflect.ValueOf(tk.GetRandomSubDualList),
		"JoinDualList":                        reflect.ValueOf(tk.JoinDualList),
		"GetNowDateString":                    reflect.ValueOf(tk.GetNowDateString),
		"GetNowTimeString":                    reflect.ValueOf(tk.GetNowTimeString),
		"GetNowTimeStringFormat":              reflect.ValueOf(tk.GetNowTimeStringFormat),
		"GetNowTimeStringFormal":              reflect.ValueOf(tk.GetNowTimeStringFormal),
		"GetNowTimeOnlyStringBeijing":         reflect.ValueOf(tk.GetNowTimeOnlyStringBeijing),
		"GetTimeFromUnixTimeStamp":            reflect.ValueOf(tk.GetTimeFromUnixTimeStamp),
		"GetTimeFromUnixTimeStampMid":         reflect.ValueOf(tk.GetTimeFromUnixTimeStampMid),
		"GetTimeStamp":                        reflect.ValueOf(tk.GetTimeStamp),
		"GetTimeStampMid":                     reflect.ValueOf(tk.GetTimeStampMid),
		"GetTimeStampNano":                    reflect.ValueOf(tk.GetTimeStampNano),
		"NowToFileName":                       reflect.ValueOf(tk.NowToFileName),
		"GetNowTimeStringHourMinute":          reflect.ValueOf(tk.GetNowTimeStringHourMinute),
		"GetNowMinutesInDay":                  reflect.ValueOf(tk.GetNowMinutesInDay),
		"NowToStrUTC":                         reflect.ValueOf(tk.NowToStrUTC),
		"GetTimeStringDiffMS":                 reflect.ValueOf(tk.GetTimeStringDiffMS),
		"StrToTime":                           reflect.ValueOf(tk.StrToTime),
		"StrToTimeByFormat":                   reflect.ValueOf(tk.StrToTimeByFormat),
		"FormatTime":                          reflect.ValueOf(tk.FormatTime),
		"IsYesterday":                         reflect.ValueOf(tk.IsYesterday),
		"DeleteItemInStringArray":             reflect.ValueOf(tk.DeleteItemInStringArray),
		"DeleteItemInIntArray":                reflect.ValueOf(tk.DeleteItemInIntArray),
		"DeleteItemInInt64Array":              reflect.ValueOf(tk.DeleteItemInInt64Array),
		"ContainsIn":                          reflect.ValueOf(tk.ContainsIn),
		"ContainsInStringList":                reflect.ValueOf(tk.ContainsInStringList),
		"IndexInStringList":                   reflect.ValueOf(tk.IndexInStringList),
		"IndexInStringListFromEnd":            reflect.ValueOf(tk.IndexInStringListFromEnd),
		"GetStringSliceFilled":                reflect.ValueOf(tk.GetStringSliceFilled),
		"Len64":                               reflect.ValueOf(tk.Len64),
		"Int64ArrayToFloat64Array":            reflect.ValueOf(tk.Int64ArrayToFloat64Array),
		"ByteSliceToStringDec":                reflect.ValueOf(tk.ByteSliceToStringDec),
		"GetValueOfMSS":                       reflect.ValueOf(tk.GetValueOfMSS),
		"Prf":                                 reflect.ValueOf(tk.Prf),
		"Prl":                                 reflect.ValueOf(tk.Prl),
		"Printf":                              reflect.ValueOf(tk.Printf),
		"Printfln":                            reflect.ValueOf(tk.Printfln),
		"Spr":                                 reflect.ValueOf(tk.Spr),
		"Pr":                                  reflect.ValueOf(tk.Pr),
		"Pl":                                  reflect.ValueOf(tk.Pl),
		"PlVerbose":                           reflect.ValueOf(tk.PlVerbose),
		"Fpl":                                 reflect.ValueOf(tk.Fpl),
		"Fpr":                                 reflect.ValueOf(tk.Fpr),
		"PlvWithError":                        reflect.ValueOf(tk.PlvWithError),
		"PlAndExit":                           reflect.ValueOf(tk.PlAndExit),
		"PlErrSimple":                         reflect.ValueOf(tk.PlErrSimple),
		"PlErrSimpleAndExit":                  reflect.ValueOf(tk.PlErrSimpleAndExit),
		"PlErrAndExit":                        reflect.ValueOf(tk.PlErrAndExit),
		"PlTXErr":                             reflect.ValueOf(tk.PlTXErr),
		"PlSimpleErrorString":                 reflect.ValueOf(tk.PlSimpleErrorString),
		"PlErr":                               reflect.ValueOf(tk.PlErr),
		"PlErrWithPrefix":                     reflect.ValueOf(tk.PlErrWithPrefix),
		"Plv":                                 reflect.ValueOf(tk.Plv),
		"Plvs":                                reflect.ValueOf(tk.Plvs),
		"Plvsr":                               reflect.ValueOf(tk.Plvsr),
		"Errf":                                reflect.ValueOf(tk.Errf),
		"FatalErr":                            reflect.ValueOf(tk.FatalErr),
		"FatalErrf":                           reflect.ValueOf(tk.FatalErrf),
		"Fatalf":                              reflect.ValueOf(tk.Fatalf),
		"CheckErr":                            reflect.ValueOf(tk.CheckErr),
		"CheckErrf":                           reflect.ValueOf(tk.CheckErrf),
		"CheckErrCompact":                     reflect.ValueOf(tk.CheckErrCompact),
		"GetUserInput":                        reflect.ValueOf(tk.GetUserInput),
		"GetInputBufferedScan":                reflect.ValueOf(tk.GetInputBufferedScan),
		"SleepSeconds":                        reflect.ValueOf(tk.SleepSeconds),
		"SleepMilliSeconds":                   reflect.ValueOf(tk.SleepMilliSeconds),
		"GetRuntimeStack":                     reflect.ValueOf(tk.GetRuntimeStack),
		"GetOSName":                           reflect.ValueOf(tk.GetOSName),
		"GetCurrentDir":                       reflect.ValueOf(tk.GetCurrentDir),
		"GetApplicationPath":                  reflect.ValueOf(tk.GetApplicationPath),
		"EnsureMakeDirs":                      reflect.ValueOf(tk.EnsureMakeDirs),
		"EnsureMakeDirsE":                     reflect.ValueOf(tk.EnsureMakeDirsE),
		"AnalyzeCommandLineParamter":          reflect.ValueOf(tk.AnalyzeCommandLineParamter),
		"GetParameterByIndexWithDefaultValue": reflect.ValueOf(tk.GetParameterByIndexWithDefaultValue),
		"ParseCommandLine":                    reflect.ValueOf(tk.ParseCommandLine),
		"GetSwitchWithDefaultValue":           reflect.ValueOf(tk.GetSwitchWithDefaultValue),
		"GetSwitchWithDefaultIntValue":        reflect.ValueOf(tk.GetSwitchWithDefaultIntValue),
		"GetSwitchWithDefaultInt64Value":      reflect.ValueOf(tk.GetSwitchWithDefaultInt64Value),
		"IfSwitchExists":                      reflect.ValueOf(tk.IfSwitchExists),
		"IfSwitchExistsWhole":                 reflect.ValueOf(tk.IfSwitchExistsWhole),
		"StrToBool":                           reflect.ValueOf(tk.StrToBool),
		"ByteToHex":                           reflect.ValueOf(tk.ByteToHex),
		"IntToStr":                            reflect.ValueOf(tk.IntToStr),
		"Int64ToStr":                          reflect.ValueOf(tk.Int64ToStr),
		"StrToIntWithDefaultValue":            reflect.ValueOf(tk.StrToIntWithDefaultValue),
		"StrToInt":                            reflect.ValueOf(tk.StrToInt),
		"StrToInt64WithDefaultValue":          reflect.ValueOf(tk.StrToInt64WithDefaultValue),
		"StrToIntPositive":                    reflect.ValueOf(tk.StrToIntPositive),
		"StrToFloat64WithDefaultValue":        reflect.ValueOf(tk.StrToFloat64WithDefaultValue),
		"StrToFloat64":                        reflect.ValueOf(tk.StrToFloat64),
		"Float64ToStr":                        reflect.ValueOf(tk.Float64ToStr),
		"StrToTimeCompact":                    reflect.ValueOf(tk.StrToTimeCompact),
		"StrToTimeCompactNoError":             reflect.ValueOf(tk.StrToTimeCompactNoError),
		"FormatStringSliceSlice":              reflect.ValueOf(tk.FormatStringSliceSlice),
		"IntToKMGT":                           reflect.ValueOf(tk.IntToKMGT),
		"IntToWYZ":                            reflect.ValueOf(tk.IntToWYZ),
		"SetLogFile":                          reflect.ValueOf(tk.SetLogFile),
		"LogWithTime":                         reflect.ValueOf(tk.LogWithTime),
		"LogWithTimeCompact":                  reflect.ValueOf(tk.LogWithTimeCompact),
		"IfFileExists":                        reflect.ValueOf(tk.IfFileExists),
		"IsFile":                              reflect.ValueOf(tk.IsFile),
		"IsDirectory":                         reflect.ValueOf(tk.IsDirectory),
		"GetFilePathSeperator":                reflect.ValueOf(tk.GetFilePathSeperator),
		"GetLastComponentOfFilePath":          reflect.ValueOf(tk.GetLastComponentOfFilePath),
		"GetDirOfFilePath":                    reflect.ValueOf(tk.GetDirOfFilePath),
		"RemoveFileExt":                       reflect.ValueOf(tk.RemoveFileExt),
		"GetFileExt":                          reflect.ValueOf(tk.GetFileExt),
		"RemoveLastSubString":                 reflect.ValueOf(tk.RemoveLastSubString),
		"AddLastSubString":                    reflect.ValueOf(tk.AddLastSubString),
		"GenerateFileListRecursively":         reflect.ValueOf(tk.GenerateFileListRecursively),
		"GetAvailableFileName":                reflect.ValueOf(tk.GetAvailableFileName),
		"LoadStringFromFile":                  reflect.ValueOf(tk.LoadStringFromFile),
		"LoadStringFromFileWithDefault":       reflect.ValueOf(tk.LoadStringFromFileWithDefault),
		"LoadStringFromFileE":                 reflect.ValueOf(tk.LoadStringFromFileE),
		"LoadStringFromFileB":                 reflect.ValueOf(tk.LoadStringFromFileB),
		"LoadBytes":                           reflect.ValueOf(tk.LoadBytes),
		"SaveStringToFile":                    reflect.ValueOf(tk.SaveStringToFile),
		"SaveStringToFileE":                   reflect.ValueOf(tk.SaveStringToFileE),
		"AppendStringToFile":                  reflect.ValueOf(tk.AppendStringToFile),
		"LoadStringList":                      reflect.ValueOf(tk.LoadStringList),
		"LoadStringListFromFile":              reflect.ValueOf(tk.LoadStringListFromFile),
		"LoadStringListBuffered":              reflect.ValueOf(tk.LoadStringListBuffered),
		"SaveStringList":                      reflect.ValueOf(tk.SaveStringList),
		"SaveStringListWin":                   reflect.ValueOf(tk.SaveStringListWin),
		"SaveStringListBufferedByRange":       reflect.ValueOf(tk.SaveStringListBufferedByRange),
		"SaveStringListBuffered":              reflect.ValueOf(tk.SaveStringListBuffered),
		"ReadLineFromBufioReader":             reflect.ValueOf(tk.ReadLineFromBufioReader),
		"RestoreLineEnds":                     reflect.ValueOf(tk.RestoreLineEnds),
		"LoadDualLineList":                    reflect.ValueOf(tk.LoadDualLineList),
		"SaveDualLineList":                    reflect.ValueOf(tk.SaveDualLineList),
		"RemoveDuplicateInDualLineList":       reflect.ValueOf(tk.RemoveDuplicateInDualLineList),
		"AppendDualLineList":                  reflect.ValueOf(tk.AppendDualLineList),
		"LoadSimpleMapFromFile":               reflect.ValueOf(tk.LoadSimpleMapFromFile),
		"LoadSimpleMapFromFileE":              reflect.ValueOf(tk.LoadSimpleMapFromFileE),
		"SimpleMapToString":                   reflect.ValueOf(tk.SimpleMapToString),
		"LoadSimpleMapFromString":             reflect.ValueOf(tk.LoadSimpleMapFromString),
		"LoadSimpleMapFromStringE":            reflect.ValueOf(tk.LoadSimpleMapFromStringE),
		"ReplaceLineEnds":                     reflect.ValueOf(tk.ReplaceLineEnds),
		"SaveSimpleMapToFile":                 reflect.ValueOf(tk.SaveSimpleMapToFile),
		"AppendSimpleMapFromFile":             reflect.ValueOf(tk.AppendSimpleMapFromFile),
		"LoadSimpleMapFromDir":                reflect.ValueOf(tk.LoadSimpleMapFromDir),
		"EncodeToXMLString":                   reflect.ValueOf(tk.EncodeToXMLString),
		"ObjectToJSON":                        reflect.ValueOf(tk.ObjectToJSON),
		"ObjectToJSONIndent":                  reflect.ValueOf(tk.ObjectToJSONIndent),
		"JSONToMapStringString":               reflect.ValueOf(tk.JSONToMapStringString),
		"JSONToObject":                        reflect.ValueOf(tk.JSONToObject),
		"SafelyGetStringForKeyWithDefault":    reflect.ValueOf(tk.SafelyGetStringForKeyWithDefault),
		"SafelyGetFloat64ForKeyWithDefault":   reflect.ValueOf(tk.SafelyGetFloat64ForKeyWithDefault),
		"SafelyGetIntForKeyWithDefault":       reflect.ValueOf(tk.SafelyGetIntForKeyWithDefault),
		"JSONToStringArray":                   reflect.ValueOf(tk.JSONToStringArray),
		"EncodeStringSimple":                  reflect.ValueOf(tk.EncodeStringSimple),
		"EncodeStringUnderline":               reflect.ValueOf(tk.EncodeStringUnderline),
		"EncodeStringCustom":                  reflect.ValueOf(tk.EncodeStringCustom),
		"DecodeStringSimple":                  reflect.ValueOf(tk.DecodeStringSimple),
		"DecodeStringUnderline":               reflect.ValueOf(tk.DecodeStringUnderline),
		"DecodeStringCustom":                  reflect.ValueOf(tk.DecodeStringCustom),
		"MD5Encrypt":                          reflect.ValueOf(tk.MD5Encrypt),
		"BytesToHex":                          reflect.ValueOf(tk.BytesToHex),
		"HexToBytes":                          reflect.ValueOf(tk.HexToBytes),
		"GetRandomByte":                       reflect.ValueOf(tk.GetRandomByte),
		"EncryptDataByTXDEE":                  reflect.ValueOf(tk.EncryptDataByTXDEE),
		"SumBytes":                            reflect.ValueOf(tk.SumBytes),
		"EncryptDataByTXDEF":                  reflect.ValueOf(tk.EncryptDataByTXDEF),
		"EncryptStreamByTXDEF":                reflect.ValueOf(tk.EncryptStreamByTXDEF),
		"DecryptStreamByTXDEF":                reflect.ValueOf(tk.DecryptStreamByTXDEF),
		"DecryptDataByTXDEE":                  reflect.ValueOf(tk.DecryptDataByTXDEE),
		"DecryptDataByTXDEF":                  reflect.ValueOf(tk.DecryptDataByTXDEF),
		"EncryptStringByTXTE":                 reflect.ValueOf(tk.EncryptStringByTXTE),
		"DecryptStringByTXTE":                 reflect.ValueOf(tk.DecryptStringByTXTE),
		"EncryptStringByTXDEE":                reflect.ValueOf(tk.EncryptStringByTXDEE),
		"DecryptStringByTXDEE":                reflect.ValueOf(tk.DecryptStringByTXDEE),
		"EncryptStringByTXDEF":                reflect.ValueOf(tk.EncryptStringByTXDEF),
		"DecryptStringByTXDEF":                reflect.ValueOf(tk.DecryptStringByTXDEF),
		"EncryptFileByTXDEF":                  reflect.ValueOf(tk.EncryptFileByTXDEF),
		"EncryptFileByTXDEFStream":            reflect.ValueOf(tk.EncryptFileByTXDEFStream),
		"DecryptFileByTXDEFStream":            reflect.ValueOf(tk.DecryptFileByTXDEFStream),
		"ErrorToString":                       reflect.ValueOf(tk.ErrorToString),
		"EncryptFileByTXDEFS":                 reflect.ValueOf(tk.EncryptFileByTXDEFS),
		"EncryptFileByTXDEFStreamS":           reflect.ValueOf(tk.EncryptFileByTXDEFStreamS),
		"DecryptFileByTXDEF":                  reflect.ValueOf(tk.DecryptFileByTXDEF),
		"DecryptFileByTXDEFS":                 reflect.ValueOf(tk.DecryptFileByTXDEFS),
		"DecryptFileByTXDEFStreamS":           reflect.ValueOf(tk.DecryptFileByTXDEFStreamS),
		"Pkcs7Padding":                        reflect.ValueOf(tk.Pkcs7Padding),
		"AESEncrypt":                          reflect.ValueOf(tk.AESEncrypt),
		"AESDecrypt":                          reflect.ValueOf(tk.AESDecrypt),
		"AnalyzeURLParams":                    reflect.ValueOf(tk.AnalyzeURLParams),
		"UrlEncode":                           reflect.ValueOf(tk.UrlEncode),
		"UrlEncode2":                          reflect.ValueOf(tk.UrlEncode2),
		"UrlDecode":                           reflect.ValueOf(tk.UrlDecode),
		"JoinURL":                             reflect.ValueOf(tk.JoinURL),
		"AddDebug":                            reflect.ValueOf(tk.AddDebug),
		"AddDebugF":                           reflect.ValueOf(tk.AddDebugF),
		"ClearDebug":                          reflect.ValueOf(tk.ClearDebug),
		"GetDebug":                            reflect.ValueOf(tk.GetDebug),
		"DownloadPageUTF8":                    reflect.ValueOf(tk.DownloadPageUTF8),
		"DownloadPage":                        reflect.ValueOf(tk.DownloadPage),
		"GetLastComponentOfUrl":               reflect.ValueOf(tk.GetLastComponentOfUrl),
		"DownloadFile":                        reflect.ValueOf(tk.DownloadFile),
		"DownloadBytes":                       reflect.ValueOf(tk.DownloadBytes),
		"PostRequest":                         reflect.ValueOf(tk.PostRequest),
		"PostRequestX":                        reflect.ValueOf(tk.PostRequestX),
		"PostRequestBytesX":                   reflect.ValueOf(tk.PostRequestBytesX),
		"PostRequestBytesWithMSSHeaderX":      reflect.ValueOf(tk.PostRequestBytesWithMSSHeaderX),
		"PostRequestBytesWithCookieX":         reflect.ValueOf(tk.PostRequestBytesWithCookieX),
		"GetFormValueWithDefaultValue":        reflect.ValueOf(tk.GetFormValueWithDefaultValue),
		"GenerateJSONPResponse":               reflect.ValueOf(tk.GenerateJSONPResponse),
		"GenerateJSONPResponseWithObject":     reflect.ValueOf(tk.GenerateJSONPResponseWithObject),
		"GenerateJSONPResponseWith2Object":    reflect.ValueOf(tk.GenerateJSONPResponseWith2Object),
		"GenerateJSONPResponseWith3Object":    reflect.ValueOf(tk.GenerateJSONPResponseWith3Object),
		"GetSuccessValue":                     reflect.ValueOf(tk.GetSuccessValue),
		"Float32ArrayToFloat64Array":          reflect.ValueOf(tk.Float32ArrayToFloat64Array),
		"CalCosineSimilarityBetweenFloatsBig": reflect.ValueOf(tk.CalCosineSimilarityBetweenFloatsBig),
		"GetDBConnection":                     reflect.ValueOf(tk.GetDBConnection),
		"GetDBRowCount":                       reflect.ValueOf(tk.GetDBRowCount),
		"GetDBRowCountCompact":                reflect.ValueOf(tk.GetDBRowCountCompact),
		"GetDBResultString":                   reflect.ValueOf(tk.GetDBResultString),
		"GetDBResultArray":                    reflect.ValueOf(tk.GetDBResultArray),
		"ConvertToGB18030":                    reflect.ValueOf(tk.ConvertToGB18030),
		"ConvertToGB18030Bytes":               reflect.ValueOf(tk.ConvertToGB18030Bytes),
		"ConvertToUTF8":                       reflect.ValueOf(tk.ConvertToUTF8),
		"ConvertStringToUTF8":                 reflect.ValueOf(tk.ConvertStringToUTF8),
		"CreateSimpleEvent":                   reflect.ValueOf(tk.CreateSimpleEvent),
		"GetAllParameters":                    reflect.ValueOf(tk.GetAllParameters),
		"GetAllSwitches":                      reflect.ValueOf(tk.GetAllSwitches),
		"ToLower":                             reflect.ValueOf(tk.ToLower),
		"ToUpper":                             reflect.ValueOf(tk.ToUpper),
	}

	// env.Packages["dialog"] = map[string]reflect.Value{
	// 	"Message": reflect.ValueOf(dialog.Message),
	// }

}

func exit() {
	os.Exit(1)
}

var versionG = "0.9a"

func showHelp() {
	tk.Pl("Gox by TopXeQ V%v\n", versionG)

	tk.Pl("Usage: gox [-v|-h] test.gox next.js, ...\n")
	tk.Pl("or just gox without arguments to start REPL instead.\n")

}

func runInteractive() int {
	var following bool
	var source string
	scanner := bufio.NewScanner(os.Stdin)

	parser.EnableErrorVerbose()

	for {
		if following {
			source += "\n"
			fmt.Print("  ")
		} else {
			fmt.Print("> ")
		}

		if !scanner.Scan() {
			break
		}
		source += scanner.Text()
		if source == "" {
			continue
		}
		if source == "quit()" {
			break
		}

		stmts, err := parser.ParseSrc(source)

		if e, ok := err.(*parser.Error); ok {
			es := e.Error()
			if strings.HasPrefix(es, "syntax error: unexpected") {
				if strings.HasPrefix(es, "syntax error: unexpected $end,") {
					following = true
					continue
				}
			} else {
				if e.Pos.Column == len(source) && !e.Fatal {
					fmt.Fprintln(os.Stderr, e)
					following = true
					continue
				}
				if e.Error() == "unexpected EOF" {
					following = true
					continue
				}
			}
		}

		following = false
		source = ""
		var v interface{}

		if err == nil {
			v, err = vm.Run(ankVMG, nil, stmts)
		}
		if err != nil {
			if ankVMG, ok := err.(*vm.Error); ok {
				fmt.Fprintf(os.Stderr, "%d:%d %s\n", ankVMG.Pos.Line, ankVMG.Pos.Column, err)
			} else if ankVMG, ok := err.(*parser.Error); ok {
				fmt.Fprintf(os.Stderr, "%d:%d %s\n", ankVMG.Pos.Line, ankVMG.Pos.Column, err)
			} else {
				fmt.Fprintln(os.Stderr, err)
			}
			continue
		}

		fmt.Printf("%#v\n", v)
	}

	if err := scanner.Err(); err != nil {
		if err != io.EOF {
			fmt.Fprintln(os.Stderr, "ReadString error:", err)
			return 12
		}
	}

	return 0
}

func eval(expA string) interface{} {
	v, errT := vm.Execute(ankVMG, nil, expA)
	if errT != nil {
		return errT.Error()
	}

	return v
}

func initAnkVM() {
	if ankVMG == nil {
		importAnkPackages()

		ankVMG = env.NewEnv()

		// err := e.Define("pl", tk.Pl)
		// if err != nil {
		// 	tk.CheckErrf("Define error: %v\n", err)
		// }

		// e.Define("prl", tk.Prl)
		ankVMG.Define("print", fmt.Print)
		ankVMG.Define("println", fmt.Println)
		ankVMG.Define("printf", fmt.Printf)
		ankVMG.Define("pl", fmt.Println)
		ankVMG.Define("printfln", tk.Pl)
		ankVMG.Define("pfl", tk.Pl)
		ankVMG.Define("exit", exit)

		ankVMG.Define("eval", eval)

		ankVMG.Define("setVar", setVar)
		ankVMG.Define("getVar", getVar)

		ankVMG.Define("argsG", os.Args[1:])

		// for GUI
		// ankVMG.Define("loadChineseFont", loadChineseFont)

		core.Import(ankVMG)

	}

}

func main() {
	rand.Seed(time.Now().Unix())

	argsT := os.Args

	if tk.IfSwitchExistsWhole(argsT, "-version") {
		tk.Pl("Gox by TopXeQ V%v", versionG)
		return
	}

	if tk.IfSwitchExistsWhole(argsT, "-h") {
		showHelp()
		return
	}

	scriptsT := tk.GetAllParameters(argsT)[1:]
	// tk.Plv(scriptsT)

	lenT := len(scriptsT)

	if lenT < 1 {
		initAnkVM()

		runInteractive()

		// tk.Pl("not enough parameters")

		return
	}

	if !tk.IfSwitchExistsWhole(argsT, "-m") {
		scriptsT = scriptsT[0:1]
	}

	for _, scriptT := range scriptsT {
		if tk.EndsWith(scriptT, ".js") {
			fcT := tk.LoadStringFromFile(scriptT)

			if tk.IsErrorString(fcT) {
				tk.Pl("failed to load file content of %v: %v", scriptT, tk.GetErrorString(fcT))

				continue
			}

			if jsVMG == nil {
				jsVMG = goja.New()

				jsVMG.Set("goPrintf", func(call goja.FunctionCall) goja.Value {
					callArgsT := call.Arguments

					argsBufT := make([]interface{}, len(callArgsT)-1)

					formatA := callArgsT[0].ToString().String()

					for i, v := range callArgsT {
						if i == 0 {
							continue
						}

						argsBufT[i-1] = v.ToString().String()
					}

					tk.Prf(formatA, argsBufT...)

					return nil
				})

				jsVMG.Set("goPrintfln", func(call goja.FunctionCall) goja.Value {
					callArgsT := call.Arguments

					argsBufT := make([]interface{}, len(callArgsT)-1)

					formatA := callArgsT[0].ToString().String()

					for i, v := range callArgsT {
						if i == 0 {
							continue
						}

						argsBufT[i-1] = v.ToString().String()
					}

					tk.Prf(formatA+"\n", argsBufT...)

					return nil
				})

				jsVMG.Set("goPrintln", func(call goja.FunctionCall) goja.Value {
					callArgsT := call.Arguments

					argsBufT := make([]interface{}, len(callArgsT))

					for i, v := range callArgsT {
						argsBufT[i] = v.ToString().String()
					}

					fmt.Println(argsBufT...)

					return nil
				})

				jsVMG.Set("goGetRandomInt", func(call goja.FunctionCall) goja.Value {
					maxA := call.Argument(0).ToInteger()

					randomNumberT := rand.Intn(int(maxA))

					rs := jsVMG.ToValue(randomNumberT)

					return rs
				})

				consoleStrT := `console = { log: goPrintln };`

				_, errT := jsVMG.RunString(consoleStrT)
				if errT != nil {
					tk.Pl("failed to run script(%v): %v", scriptT, errT)

					continue
				}

			}

			v, errT := jsVMG.RunString(fcT)
			if errT != nil {
				tk.Pl("failed to run script(%v): %v", scriptT, errT)

				continue
			}

			variableG["Out"] = v.Export()

			// tk.Pl("%#v", rs)

			return
		} else if tk.EndsWith(scriptT, ".ank") || tk.EndsWith(scriptT, ".gox") {
			fcT := tk.LoadStringFromFile(scriptT)

			if tk.IsErrorString(fcT) {
				tk.Pl("failed to load file content of %v: %v", scriptT, tk.GetErrorString(fcT))

				continue
			}

			// if ankVMG == nil {
			// 	importAnkPackages()

			// 	ankVMG = env.NewEnv()

			// 	// err := e.Define("pl", tk.Pl)
			// 	// if err != nil {
			// 	// 	tk.CheckErrf("Define error: %v\n", err)
			// 	// }

			// 	// e.Define("prl", tk.Prl)
			// 	ankVMG.Define("print", fmt.Print)
			// 	ankVMG.Define("println", fmt.Println)
			// 	ankVMG.Define("printf", fmt.Printf)
			// 	ankVMG.Define("pl", fmt.Println)
			// 	ankVMG.Define("printfln", tk.Pl)
			// 	ankVMG.Define("pfl", tk.Pl)
			// 	ankVMG.Define("exit", exit)

			// 	ankVMG.Define("setVar", setVar)
			// 	ankVMG.Define("getVar", getVar)

			// 	core.Import(ankVMG)

			// }

			// ankVMG.Define("inG", map[string]interface{}{"Args": os.Args})

			initAnkVM()

			script := fcT //`println("Hello World :)")`

			_, errT := vm.Execute(ankVMG, nil, script)
			if errT != nil {
				tk.Pl("failed to execute script(%v) error: %v", scriptT, errT)
				continue
			}

			rs, errT := ankVMG.Get("outG")

			// tk.CheckErrCompact(errT)

			if errT == nil && rs != nil {
				tk.Pl("%#v", rs)
			}

			// tk.Pl("%#v", rs)

		}
	}

	// tk.Pl("Gox by TopXeQ V%v", versionG)

	// fmt.Println("")

}
