package examples

import "github.com/go-color-term/go-color-term/coloring"

type ComplexPhrasesExample struct{}

func (example *ComplexPhrasesExample) GetTitle() string {
	return "Complex phrases"
}

func (example *ComplexPhrasesExample) GetDescription() string {
	return "An example showing the styling of sentence using SentenceBuilder."
}

func (example *ComplexPhrasesExample) Run() {
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
