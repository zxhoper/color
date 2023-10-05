package github.com/zxhoper/color

import "fmt"

func Example() {
	NoteGreen("The color print setup is: OK")
	NoteTYellowf("Color Notef is %s!\n", "OK")

	Example_PrintAutoColorNotef()
	NoteAutof("Auto %s\n", "ONE")
	NoteAutof("Auto %s\n", "TWO")
	NoteAutof("Auto %s\n", "THREE")

	Notefn("Format line with %s appended.", "NEWLINE")

	Example_PrintAllColorNotef()

	Example_PrintAllColor()

	Example_PrintAllColorTitle()

	NotePrefix = "PERFIX "
	Example_PrintAllColor()
	Example_PrintAllColorTitle()

	ln := "line"
	nl := "newline"

	Notefn("Format %s with %s appended.", ln, nl)

	MyDebug = true

	DeNotefn("Format debug %s with %s appended.", ln, nl)

	Exaple_PrintAllDeNote()

	NotePrefix = ""
}

func Example_PrintAllColor() {
	NoteBlockFirst()

	Note("Default Note")
	Notef("Default Notef\n")

	NoteGreen("Green")
	NoteRed("Red")
	NoteYellow("Yellow")
	NoteBlue("Blue")
	NotePurple("Purple")
	NoteCyan("Cyan")
	NoteGray("Gray")
	NoteWhite("White")
	NoteBlockLast()

	NoteHr(":")
}
func Example_PrintAllColorNotef() {
	NoteBlockFirst()

	Notef("Default Notef\n")

	NoteGreenf("Green")
	NoteRedf("Red")
	NoteYellowf("Yellow")
	NoteBluef("Blue")
	NotePurplef("Purple")
	NoteCyanf("Cyan")
	NoteGrayf("Gray")
	NoteWhitef("White\n")

	NoteWhitef("Print AUTO COLOR\n")
	for i := 0; i < GetColorType(); i++ {
		NoteAutof("Print auto color number: %d  ", i)
	}
	NoteColor("")

	NoteBlockLast()

	NoteHr(":")
}
func Example_PrintAllColorTitle() {
	NoteBlockFirst()
	NoteT("Default NoteT")
	NoteTf("Default Note Title \n")
	NoteTGreenf("Green\n")
	NoteTRedf("Red\n")
	NoteTYellowf("Yellow\n")
	NoteTBluef("Blue\n")
	NoteTPurplef("Purple\n")
	NoteTCyanf("Cyan\n")
	NoteTGrayf("Gray\n")
	NoteTWhitef("White\n")
	NoteBlockLast()
}

func Exaple_PrintAllDeNote() {
	DeNoteBlockFirst()

	DeNote("Note for debug")
	DeNote("with debug prefix")
	DeNoteHr(":")
	DeNote("Default DeNore")

	DeNoteTf("Default DeNotef\n")
	DeNoteTGreenf("Green\n")
	DeNoteTRedf("Red\n")
	DeNoteTYellowf("Yellow\n")
	DeNoteTBluef("Blue\n")
	DeNoteTPurplef("Purple\n")
	DeNoteTCyanf("Cyan\n")
	DeNoteTGrayf("Gray\n")
	DeNoteTWhitef("White\n")

	DeNoteBlockLast()
}

func Example_PrintAutoColorNotef() {
	NoteBlockFirst()

	NoteWhitef("AUTO COLOR\n")

	idx := 0
	for {
		idx++
		for i := 0; i < GetColorType(); i++ {
			NoteAutof("Color %d ", i)
		}
		fmt.Printf("\n")
		if idx > 10 {
			break
		}
	}

	NoteBlockLast()

	NoteHr(":")
}
