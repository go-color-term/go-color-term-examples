package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

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

	fmt.Println("Select example:")

	for i, example := range examples {
		fmt.Printf("%d) %s\n", i+1, example.GetTitle())
	}

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	selectedExample, err := strconv.Atoi(scanner.Text())
	selectedExample -= 1

	if err != nil || selectedExample < 0 || selectedExample > len(examples)-1 {
		fmt.Println(coloring.Bold("Example not found!"))
		os.Exit(1)
	}

	fmt.Println(examples[selectedExample].GetDescription())

	examples[selectedExample].Run()

	fmt.Println()
}
