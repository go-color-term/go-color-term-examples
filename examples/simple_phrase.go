package examples

import (
	"fmt"

	"github.com/go-color-term/go-color-term/coloring"
)

type SimplePhraseExample struct{}

func (example *SimplePhraseExample) GetTitle() string {
	return "Simple Phrase"
}

func (example *SimplePhraseExample) GetDescription() string {
	return "An example showing the styling of a simple phrase using StyleBuilder."
}

func (example *SimplePhraseExample) Run() {
	quick := coloring.For("quick").Green().Italic().Underline().Styled()
	brown := coloring.For("brown").Color(15).Background().Rgb(0xB3, 0x64, 0x00).Styled()
	fox := coloring.For("fox").White().Background().Rgb(0xE8, 0x1D, 0x1D).Styled()
	jumps := coloring.For("jumps").Yellow().Bold().Background().Color(8).Styled()
	dog := coloring.For("dog").Bold().Italic().Styled()

	template := "The %s %s %s %s over the lazy %s."

	fmt.Printf(template, quick, brown, fox, jumps, dog)
	fmt.Println()
	fmt.Println()
	fmt.Printf(template, quick.Unformatted(), brown.Unformatted(), fox.Unformatted(), jumps.Unformatted(), dog.Unformatted())
}
