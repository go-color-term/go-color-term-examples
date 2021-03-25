package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/go-color-term/go-color-term-examples/examples"
	"github.com/go-color-term/go-color-term/coloring"
)

type Example interface {
	GetTitle() string
	GetDescription() string
	Run()
}

func main() {
	examples := [...]Example{
		&examples.TestDriveExample{},
		&examples.DirectoryListingExample{},
		&examples.ColorMatrixExample{},
		&examples.RgbColorMatrixExample{},
		&examples.SimplePhraseExample{},
		&examples.ComplexPhrasesExample{},
		&examples.UtilityExtrasExample{},
	}

	for {
		fmt.Println("Select example (press ENTER to exit):")

		for i, example := range examples {
			fmt.Printf("%d) %s\n", i+1, example.GetTitle())
		}

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()

		text := strings.TrimSpace(scanner.Text())
		if text == "" {
			fmt.Println(coloring.Sentence().
				BoldStart().
				Text("Thanks for trying ").
				ItalicStart().
				Color("go", coloring.BRIGHTBLUE).
				Text("-").
				Color("color", coloring.BRIGHTRED).
				Text("-").
				Color("term", coloring.BRIGHTGREEN).
				ItalicEnd().
				Text(". Goodbye!"))

			os.Exit(1)
		}

		selectedExample, err := strconv.Atoi(text)
		selectedExample--

		if err != nil || selectedExample < 0 || selectedExample > len(examples)-1 {
			fmt.Println(coloring.Bold("Example not found!"))

			continue
		}

		fmt.Println(examples[selectedExample].GetDescription())

		examples[selectedExample].Run()

		fmt.Println()
	}
}
