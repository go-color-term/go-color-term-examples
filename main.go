package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/go-color-term/go-color-term-examples/examples"
	"github.com/go-color-term/go-color-term/coloring"
)

type Example struct {
	Title string
	Func  func()
}

func main() {
	examples := [...]Example{
		{"Test drive", examples.TestDrive},
		{"List directory", examples.DirTree},
		{"8-bit color matrix", examples.Print8BitColorMatrix},
		{"Phrase", examples.SimplePhrase},
		{"Complex phrases", examples.ComplexPhrases},
		{"Bright colors (utility extras)", examples.BrightColors},
	}

	fmt.Println("Select example:")

	for i, example := range examples {
		fmt.Printf("%d) %s\n", i+1, example.Title)
	}

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	selectedExample, err := strconv.Atoi(scanner.Text())
	selectedExample -= 1

	if err != nil || selectedExample < 0 || selectedExample > len(examples)-1 {
		fmt.Println(coloring.Bold("Example not found!"))
		os.Exit(1)
	}

	examples[selectedExample].Func()

	fmt.Println()
}
