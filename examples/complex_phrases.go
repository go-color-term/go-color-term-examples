package examples

import "github.com/nelsonghezzi/go-color-term/coloring"

func ComplexPhrases() {
	coloring.Sentence().
		Text("This is ").
		ColorSet(coloring.GREEN).
		Text("some ").
		UnderlineStart().
		Text("nice").ColorReset().
		Text(" way").
		UnderlineEnd().
		Text(" of writing ").
		BoldStart().Color("red", coloring.RED).
		Text(" text!!!").
		Println()

	coloring.Sentence().
		BoldStart().StrikethroughStart().
		Color("Red", coloring.RED).
		Text(", ").
		Color("green", coloring.GREEN).
		Text(" and ").
		Color("yellow", coloring.YELLOW).
		StrikethroughEnd().
		Text(" all crossed!").
		Println()

	coloring.Sentence().BoldStart().Invert("BLACK").Text("WHITE").Println()

	coloring.Sentence().Blink("Blinking").Println()

	coloring.Sentence().BackgroundSet(coloring.BLUE).
		Color("This is some text ", coloring.MAGENTA).
		Bold("with a bold ending").
		Println()
}
