package main

import (
	"os"
	"fmt"
	"github.com/fatih/color"
)

var preColor, postColor *color.Color

type Buffer interface {
	SetText(s string)
	GetText() string
	Insert(s string)
	Delete()
	Backspace()
	CursorNext()
	CursorPrevious()
	debugPrint()
}

func printUsage() {
	fmt.Println("Usage: ./main <gap|split>")
}

func main() {
	if len(os.Args) != 2 {
		printUsage()
		return
	}

	var buffer Buffer

	switch os.Args[1] {
	case "gap":
		buffer = &GapBuffer{}
	case "split":
		buffer = &SplitBuffer{}
	default:
		printUsage()
		return
	}

	preColor = color.New(color.FgBlack).Add(color.BgCyan)
	postColor = color.New(color.FgBlack).Add(color.BgGreen)

	buffer.SetText("Greetings!")
	buffer.debugPrint()

	buffer.CursorNext()
	buffer.CursorNext()
	buffer.CursorNext()
	buffer.debugPrint()

	buffer.CursorPrevious()
	buffer.debugPrint()

	buffer.Insert(":)")
	buffer.debugPrint()

	buffer.Delete()
	buffer.Delete()
	buffer.debugPrint()

	buffer.CursorNext()
	buffer.CursorNext()
	buffer.CursorNext()
	buffer.debugPrint()

	buffer.Backspace()
	buffer.debugPrint()

	fmt.Println(buffer.GetText())
}
