package github.com/zxhoper/color

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"
)

// ==============================================================================================
// 2023-05-02_172806
//
//	 var debugOnFlag = flag.Bool("d", false, "debug flag")
//
//	 func main() {
//	   flag.Parse()
//	   MyDebug = *debugOnFlag
//		}
var MyDebug bool
var TestPrefix = "    ---   | "
var NotePrefix = ""
var DeNotePrefix = " --> "

//var NoteRed = fmt.Println
//var NoteGreen = fmt.Println
//var NoteBluef = fmt.Printf
//var NoteGreenf = fmt.Printf
//var NoteGreenHr = fmt.Println
//var DeNoteYellow = fmt.Println
//var DeNoteYellowf = fmt.Printf

//
//const (
//	clrNon = "0"
//
//	clrRed    = "31"
//	clrGreen  = "32"
//	clrYellow = "33"
//	clrBlue   = "34"
//	clrPurple = "35"
//	clrCyan   = "36"
//	clrGray   = "37"
//	clrWhite  = "97"
//)

var clr = map[string]string{
	"Non":    "0",
	"Red":    "31",
	"Green":  "32",
	"Yellow": "33",
	"Blue":   "34",
	"Purple": "35",
	"Cyan":   "36",
	"Gray":   "37",
	"White":  "97",
}

func NoteSTEP(a ...interface{}) {
	var sb strings.Builder
	sb.WriteString("STEP " + fmt.Sprintf("%s", a[0]) + ": ")

	for i := 1; i <= len(a)-1; i++ {
		sb.WriteString(fmt.Sprintf("%s", a[i]))
	}
	str := sb.String()
	NoteColor(getClrCode(clr["Red"]), str)
}

func DeNoteSTEP(a ...interface{}) {
	if MyDebug {
		var sb strings.Builder
		sb.WriteString("STEP " + fmt.Sprintf("%s", a[0]) + ": ")

		for i := 1; i <= len(a)-1; i++ {
			sb.WriteString(fmt.Sprintf("%s", a[i]))
		}
		str := sb.String()
		DeNoteColor(getClrCode(clr["Red"]), str)
	}
}
func GetColorType() int {
	return len(clr)
}

var Acc = newAutoColorControl()

func getClrCode(c string) string {
	return fmt.Sprintf("\033[%sm", c)
}

func NoteColor(c string, a ...interface{}) {
	_, _ = fmt.Fprintf(os.Stdout, c)
	eewP(os.Stdout, NotePrefix, a...)
	_, _ = fmt.Fprintf(os.Stdout, "\033[0m")
}

func DeNoteColor(c string, a ...interface{}) {
	if MyDebug {
		_, _ = fmt.Fprintf(os.Stdout, c)
		eewP(os.Stdout, DeNotePrefix, a...)
		_, _ = fmt.Fprintf(os.Stdout, "\033[0m")
	}
}

func NoteGreen(a ...interface{}) {
	NoteColor(getClrCode(clr["Green"]), a...)
}
func DeNoteGreen(a ...interface{}) {
	DeNoteColor(getClrCode(clr["Green"]), a...)
}
func NoteRed(a ...interface{}) {
	NoteColor(getClrCode(clr["Red"]), a...)
}
func DeNoteRed(a ...interface{}) {
	DeNoteColor(getClrCode(clr["Red"]), a...)
}
func NoteYellow(a ...interface{}) {
	NoteColor(getClrCode(clr["Yellow"]), a...)
}
func DeNoteYellow(a ...interface{}) {
	DeNoteColor(getClrCode(clr["Yellow"]), a...)
}
func NoteBlue(a ...interface{}) {
	NoteColor(getClrCode(clr["Blue"]), a...)
}
func DeNoteBlue(a ...interface{}) {
	DeNoteColor(getClrCode(clr["Blue"]), a...)
}
func NotePurple(a ...interface{}) {
	NoteColor(getClrCode(clr["Purple"]), a...)
}
func DeNotePurple(a ...interface{}) {
	DeNoteColor(getClrCode(clr["Purple"]), a...)
}
func NoteCyan(a ...interface{}) {
	NoteColor(getClrCode(clr["Cyan"]), a...)
}
func DeNoteCyan(a ...interface{}) {
	DeNoteColor(getClrCode(clr["Cyan"]), a...)
}
func NoteGray(a ...interface{}) {
	NoteColor(getClrCode(clr["Gray"]), a...)
}
func DeNoteGray(a ...interface{}) {
	DeNoteColor(getClrCode(clr["Gray"]), a...)
}
func NoteWhite(a ...interface{}) {
	NoteColor(getClrCode(clr["White"]), a...)
}
func DeNoteWhite(a ...interface{}) {
	DeNoteColor(getClrCode(clr["White"]), a...)
}

func Note(a ...interface{}) {
	eewP(os.Stdout, NotePrefix, a...)
}
func DeNote(a ...interface{}) {
	if MyDebug {
		eewP(os.Stdout, DeNotePrefix, a...)
	}
}

// eewP  wraps fmt.Println with prefix
func eewP(w io.Writer, pfix string, a ...interface{}) {
	var prefixedAny []interface{}
	if pfix != "" {
		prefixedAny = append(prefixedAny, pfix)
	}
	prefixedAny = append(prefixedAny, a...)
	_, _ = fmt.Fprintln(w, prefixedAny...)
}

func NoteT(s string) {
	eewPT(os.Stdout, NotePrefix, s)
}
func DeNoteT(s string) {
	if MyDebug {
		eewPT(os.Stdout, DeNotePrefix, s)
	}
}
func eewPT(w io.Writer, pfix string, s string) {
	_, _ = fmt.Fprintln(w, pfix+StringRepeat("=", 51))
	_, _ = fmt.Fprintf(w, pfix+"== %-45s ==\n", s)
	_, _ = fmt.Fprintln(w, pfix+StringRepeat("=", 51))
}

func Notef(format string, a ...interface{}) {
	eewPf(os.Stdout, NotePrefix, format, a...)
}

type AutoCololrCtl struct {
	TotalColor   int
	Colors       []string
	CurrentColor int
}

func newAutoColorControl() *AutoCololrCtl {
	var colors []string
	for _, c := range clr {
		colors = append(colors, c)
	}

	return &AutoCololrCtl{
		TotalColor:   GetColorType(),
		Colors:       colors,
		CurrentColor: 0,
	}
}

func (acc *AutoCololrCtl) popCurColor() string {
	current := acc.CurrentColor
	if acc.CurrentColor == acc.TotalColor-1 {
		acc.CurrentColor = 0
	} else {
		acc.CurrentColor = current + 1
	}
	return acc.Colors[current]
}

func NoteAutof(format string, a ...interface{}) {
	NoteColorf(getClrCode(Acc.popCurColor()), format, a...)
}

func NoteColorf(colorCode string, format string, a ...interface{}) {
	formatColor := fmt.Sprintf("%s%s%s", colorCode, format, getClrCode(clr["Non"]))
	eewPf(os.Stdout, NotePrefix, formatColor, a...)
}

func DeNoteColorf(colorCode string, format string, a ...interface{}) {
	if MyDebug {
		formatColor := fmt.Sprintf(DeNotePrefix+"%s%s%s", colorCode, format, getClrCode(clr["Non"]))
		eewPf(os.Stdout, NotePrefix, formatColor, a...)
	}
}

func NoteGreenf(format string, a ...interface{}) {
	NoteColorf(getClrCode(clr["Green"]), format, a...)
}
func DeNoteGreenf(format string, a ...interface{}) {
	DeNoteColorf(getClrCode(clr["Green"]), format, a...)
}
func NoteRedf(format string, a ...interface{}) {
	NoteColorf(getClrCode(clr["Red"]), format, a...)
}
func DeNoteRedf(format string, a ...interface{}) {
	DeNoteColorf(getClrCode(clr["Red"]), format, a...)
}
func NoteYellowf(format string, a ...interface{}) {
	NoteColorf(getClrCode(clr["Yellow"]), format, a...)
}
func DeNoteYellowf(format string, a ...interface{}) {
	DeNoteColorf(getClrCode(clr["Yellow"]), format, a...)
}
func NoteBluef(format string, a ...interface{}) {
	NoteColorf(getClrCode(clr["Blue"]), format, a...)
}
func DeNoteBluef(format string, a ...interface{}) {
	DeNoteColorf(getClrCode(clr["Blue"]), format, a...)
}
func NotePurplef(format string, a ...interface{}) {
	NoteColorf(getClrCode(clr["Purple"]), format, a...)
}
func DeNotePurplef(format string, a ...interface{}) {
	DeNoteColorf(getClrCode(clr["Purple"]), format, a...)
}
func NoteCyanf(format string, a ...interface{}) {
	NoteColorf(getClrCode(clr["Cyan"]), format, a...)
}
func DeNoteCyanf(format string, a ...interface{}) {
	DeNoteColorf(getClrCode(clr["Cyan"]), format, a...)
}
func NoteGrayf(format string, a ...interface{}) {
	NoteColorf(getClrCode(clr["Gray"]), format, a...)
}
func DeNoteGrayf(format string, a ...interface{}) {
	DeNoteColorf(getClrCode(clr["Gray"]), format, a...)
}
func NoteWhitef(format string, a ...interface{}) {
	NoteColorf(getClrCode(clr["White"]), format, a...)
}
func DeNoteWhitef(format string, a ...interface{}) {
	DeNoteColorf(getClrCode(clr["White"]), format, a...)
}

func DeNotef(format string, a ...interface{}) {
	if MyDebug {
		fmt.Printf("%s", DeNotePrefix)
		eewPf(os.Stdout, DeNotePrefix, format, a...)
	}
}
func eewPf(w io.Writer, pfix string, format string, a ...interface{}) {
	prefixedFormat := pfix + format
	_, _ = fmt.Fprintf(w, prefixedFormat, a...)
}

func Notefn(format string, a ...interface{}) {
	eewPfn(os.Stdout, format, NotePrefix, a...)
}
func DeNotefn(format string, a ...interface{}) {
	if MyDebug {
		eewPfn(os.Stdout, DeNotePrefix, format, a...)
	}
}
func eewPfn(w io.Writer, pfix string, format string, a ...interface{}) {
	prefixedFormat := pfix + format + "\n"
	_, _ = fmt.Fprintf(w, prefixedFormat, a...)
}

func NoteTColorf(c string, format string, a ...interface{}) {
	formatColor := fmt.Sprintf("%s%s%s%s", c, "== ", format, getClrCode(clr["Non"]))
	prefixedFormatColor := NotePrefix + formatColor
	hrColor := fmt.Sprintf("%s%s%s", c, StringRepeat("=", 51), getClrCode(clr["Non"]))
	_, _ = fmt.Fprintln(os.Stdout, NotePrefix+hrColor)
	_, _ = fmt.Fprintf(os.Stdout, prefixedFormatColor, a...)
	_, _ = fmt.Fprintln(os.Stdout, NotePrefix+hrColor)
}

func NoteTRedf(format string, a ...interface{}) {
	NoteTColorf(getClrCode(clr["Red"]), format, a...)
}
func NoteTGreenf(format string, a ...interface{}) {
	NoteTColorf(getClrCode(clr["Green"]), format, a...)
}
func NoteTYellowf(format string, a ...interface{}) {
	NoteTColorf(getClrCode(clr["Yellow"]), format, a...)
}
func NoteTBluef(format string, a ...interface{}) {
	NoteTColorf(getClrCode(clr["Blue"]), format, a...)
}
func NoteTPurplef(format string, a ...interface{}) {
	NoteTColorf(getClrCode(clr["Purple"]), format, a...)
}
func NoteTCyanf(format string, a ...interface{}) {
	NoteTColorf(getClrCode(clr["Cyan"]), format, a...)
}
func NoteTGrayf(format string, a ...interface{}) {
	NoteTColorf(getClrCode(clr["Gray"]), format, a...)
}
func NoteTWhitef(format string, a ...interface{}) {
	NoteTColorf(getClrCode(clr["White"]), format, a...)
}

func NoteTf(format string, a ...interface{}) {
	eewPft(os.Stdout, NotePrefix, format, a...)
}

func DeNoteTf(format string, a ...interface{}) {
	if MyDebug {
		eewPft(os.Stdout, DeNotePrefix, format, a...)
	}
}
func eewPft(w io.Writer, pfix string, format string, a ...interface{}) {
	prefixedFormat := pfix + "== " + format
	_, _ = fmt.Fprintln(w, pfix+StringRepeat("=", 51))
	_, _ = fmt.Fprintf(w, prefixedFormat, a...)
	_, _ = fmt.Fprintln(w, pfix+StringRepeat("=", 51))
}

func DeNoteTColorf(c string, format string, a ...interface{}) {
	if MyDebug {
		formatColor := fmt.Sprintf("%s%s%s%s", c, "== ", format, getClrCode(clr["Non"]))
		prefixedFormatColor := DeNotePrefix + formatColor
		hrColor := fmt.Sprintf("%s%s%s", c, StringRepeat("=", 51), getClrCode(clr["Non"]))
		_, _ = fmt.Fprintln(os.Stdout, DeNotePrefix+hrColor)
		_, _ = fmt.Fprintf(os.Stdout, prefixedFormatColor, a...)
		_, _ = fmt.Fprintln(os.Stdout, DeNotePrefix+hrColor)
	}
}

func DeNoteTRedf(format string, a ...interface{}) {
	DeNoteTColorf(getClrCode(clr["Red"]), format, a...)
}
func DeNoteTGreenf(format string, a ...interface{}) {
	DeNoteTColorf(getClrCode(clr["Green"]), format, a...)
}
func DeNoteTYellowf(format string, a ...interface{}) {
	DeNoteTColorf(getClrCode(clr["Yellow"]), format, a...)
}
func DeNoteTBluef(format string, a ...interface{}) {
	DeNoteTColorf(getClrCode(clr["Blue"]), format, a...)
}
func DeNoteTPurplef(format string, a ...interface{}) {
	DeNoteTColorf(getClrCode(clr["Purple"]), format, a...)
}
func DeNoteTCyanf(format string, a ...interface{}) {
	DeNoteTColorf(getClrCode(clr["Cyan"]), format, a...)
}
func DeNoteTGrayf(format string, a ...interface{}) {
	DeNoteTColorf(getClrCode(clr["Gray"]), format, a...)
}
func DeNoteTWhitef(format string, a ...interface{}) {
	DeNoteTColorf(getClrCode(clr["White"]), format, a...)
}
func NoteRedHr(str string) {
	NoteColorHr(getClrCode(clr["Red"]), str)
}
func NoteGreenHr(str string) {
	NoteColorHr(getClrCode(clr["Green"]), str)
}
func NoteBlueHr(str string) {
	NoteColorHr(getClrCode(clr["Blue"]), str)
}
func NoteHr(str string) {
	chr := takeFirstChar(str)
	eewPhr(os.Stdout, NotePrefix, chr)
}
func takeFirstChar(str string) string {
	if len(str) > 0 {
		return str[0:1] // only take first character o repeat
	}
	return "=" // default using "="
}
func NoteColorHr(c string, str string) {
	chr := takeFirstChar(str)
	_, _ = fmt.Fprintf(os.Stdout, c)
	eewPhr(os.Stdout, NotePrefix, chr)
	_, _ = fmt.Fprintf(os.Stdout, "\033[0m")
}
func eewPhr(w io.Writer, pfix string, sep string) {
	_, _ = fmt.Fprintln(w, pfix+StringRepeat(sep, 72))
}
func NoteBlockFirst() {
	_, _ = fmt.Fprintln(os.Stdout, StringRepeat("-", 64)+"\\\\")
}
func DeNotePrtHr(str string) {
	if MyDebug {
		chr := takeFirstChar(str)
		eewPhr(os.Stdout, DeNotePrefix, chr)
	}
}
func DeNoteHr(str string) {
	if MyDebug {
		chr := takeFirstChar(str)
		eewPhr(os.Stdout, DeNotePrefix, chr)
	}
}

func NoteBlockLast() {
	eewPLast(os.Stdout, NotePrefix)
}
func DeNoteBlockFirst() {
	if MyDebug {
		fmt.Printf("%s", DeNotePrefix)
		eewPFirst(os.Stdout, NotePrefix)
	}
}
func DeNoteBlockLast() {
	if MyDebug {
		fmt.Printf("%s", DeNotePrefix)
		eewPLast(os.Stdout, NotePrefix)
	}
}

func eewPFirst(w io.Writer, pfix string) {
	_, _ = fmt.Fprintln(w, pfix+StringRepeat("-", 64)+"\\\\")
}
func eewPLast(w io.Writer, pfix string) {
	_, _ = fmt.Fprintln(w, pfix+StringRepeat("-", 64)+"//")
}

func StringRepeat(s string, repeatTimes int) string {
	var sb strings.Builder
	for i := 0; i < repeatTimes; i++ {
		sb.WriteString(s)
	}
	return sb.String()
}

func NoteFailExit(action string, err error, moreInfo string) {
	NoteRedf("error when [[%s]] with error =>[[%s]], more info=>[[%s]]\n", action, err, moreInfo)
	os.Exit(1)
}

func NoteSimpleStructPrint(theStruct interface{}) {
	// 2023-04-28_222524
	values := reflect.ValueOf(theStruct)
	types := values.Type()

	maxKeyLen := 0
	maxTypeLen := 0
	for i := 0; i < values.NumField(); i++ {
		key := types.Field(i).Name
		if len(key) > maxKeyLen {
			maxKeyLen = len(key)
		}
		typeName := types.Field(i).Type.Name()
		if len(typeName) > maxTypeLen {
			maxTypeLen = len(typeName)
		}
	}

	fmtKeyFieldWidth := maxKeyLen + 1
	fmtTypeFieldWidth := maxTypeLen + 1
	for i := 0; i < values.NumField(); i++ {

		index := types.Field(i).Index[0]
		key := types.Field(i).Name
		thisType := types.Field(i).Type.Name()
		val := values.Field(i)

		NoteYellowf("[%d] ", index)
		NoteGreenf("%s%s", key, StringRepeat(" ", fmtKeyFieldWidth-len(key)))
		NotePurplef("(%s)%s=> ", thisType, StringRepeat(" ", fmtTypeFieldWidth-len(thisType)))
		NoteBluef("%v \n", val)
	}
}

type NotePrintItem struct {
	Key     string
	ValType string
	Value   string
}

func SimpleMapPrint(mymap any) {

	values := reflect.ValueOf(mymap)
	//types := values.Type()

	maxLen := maxLenInSlice([]string{"VALUE", "KIND"})
	NoteYellowf("VALUE%s%s %v\n", StringRepeat(" ", maxLen-len("VALUE")+1), "=", values)
	NoteYellowf("KIND%s%s %v\n", StringRepeat(" ", maxLen-len("KIND")+1), "=", values.Kind())

	reportSlice := []NotePrintItem{}
	if values.Kind() == reflect.Map {
		for _, e := range values.MapKeys() {
			v := values.MapIndex(e)
			switch t := v.Interface().(type) {
			case int:
				reportSlice = append(reportSlice, NotePrintItem{Key: fmt.Sprintf("%s", e.String()),
					ValType: "int",
					Value:   fmt.Sprintf("%d", t)})
				//fmt.Println(e, t)
			case string:
				reportSlice = append(reportSlice, NotePrintItem{Key: fmt.Sprintf("%s", e.String()),
					ValType: "string",
					Value:   fmt.Sprintf("%s", t)})
				//fmt.Println(e, t)
			case bool:
				reportSlice = append(reportSlice, NotePrintItem{Key: fmt.Sprintf("%s", e.String()),
					ValType: "bool",
					Value:   fmt.Sprintf("%t", t)})
				//fmt.Println(e, t)
			default:
				reportSlice = append(reportSlice, NotePrintItem{Key: fmt.Sprintf("%s", e.String()),
					ValType: "UNKNOW",
					Value:   fmt.Sprintf("%s", fmt.Sprintln(t))})
				//fmt.Println("don not know the type")

			}
		}
	}

	maxKeyLen := 0
	maxTypeLen := 0
	for i := 0; i < len(reportSlice); i++ {
		key := reportSlice[i].Key
		if len(key) > maxKeyLen {
			maxKeyLen = len(key)
		}

		typeName := reportSlice[i].ValType
		if len(typeName) > maxTypeLen {
			maxTypeLen = len(typeName)
		}
	}

	fmtKeyFieldWidth := maxKeyLen + 1
	fmtTypeFieldWidth := maxTypeLen + 1

	for i := 0; i < len(reportSlice); i++ {

		index := i
		key := reportSlice[i].Key
		thisType := reportSlice[i].ValType
		val := reportSlice[i].Value

		NoteYellowf("[%d] ", index)
		NoteGreenf("%s%s", key, StringRepeat(" ", fmtKeyFieldWidth-len(key)))
		NotePurplef("(%s)%s=> ", thisType, StringRepeat(" ", fmtTypeFieldWidth-len(thisType)))
		NoteBluef("[>%v<] \n", val)
	}
}

func maxLenInSlice(strSlice []string) int {
	maxLenth := 0
	for _, s := range strSlice {
		if len(s) > maxLenth {
			maxLenth = len(s)
		}
	}
	return maxLenth
}

// ===============================================================================================
