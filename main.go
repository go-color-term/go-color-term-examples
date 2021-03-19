package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"strconv"
	"strings"

	"github.com/nelsonghezzi/go-color-term/coloring"
)

type Example struct {
	Title string
	Func  func()
}

func main() {
	examples := [...]Example{
		{"List directory", dirTree},
		{"8-bit color matrix", print8BitColorMatrix},
		{"Phrase", printPhrase},
		{"Complex phrases", complexPhrases},
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

func dirTree() {
	listContents(".", 0, 3)
}

func listContents(path string, level, depth int) {
	if depth == 0 {
		return
	}

	entries, err := os.ReadDir(path)

	if err != nil {
		fmt.Printf("%v", err)
	}

	for _, entry := range entries {
		line := coloredEntry(entry)

		if level > 0 {
			line = strings.Repeat("  ", level-1) + "|-" + line
		}

		fmt.Println(line)

		if entry.IsDir() {
			walk := entry.Name()
			if path != "." {
				walk = path + "/" + walk
			}

			listContents(walk, level+1, depth-1)
		}
	}
}

var colorizeRegularFile = coloring.New().Func()
var colorizeDirectory = coloring.New().Bold().Underline().Func()
var colorizeGoFile = coloring.New().Bold().Blue().Func()
var colorizeTxtFile = coloring.New().Italic().White().Background().Rgb(127, 127, 255).Func()
var colorizeExecutableFile = coloring.New().Bold().Green().Func()

var colorizeHidden = coloring.New().Color(coloring.BRIGHTBLACK).Bold().Background().White()
var colorizeHiddenFile = colorizeHidden.Func()
var colorizeHiddenFolder = colorizeHidden.Underline().Func()

func coloredEntry(entry fs.DirEntry) string {
	colorizerFunc := colorizeRegularFile
	isHidden := strings.HasPrefix(entry.Name(), ".")

	if entry.IsDir() {
		if isHidden {
			colorizerFunc = colorizeHiddenFolder
		} else {
			colorizerFunc = colorizeDirectory
		}
	} else {
		if isHidden {
			colorizerFunc = colorizeHiddenFile
		} else if strings.HasSuffix(entry.Name(), ".go") {
			colorizerFunc = colorizeGoFile
		} else if strings.HasSuffix(entry.Name(), ".txt") {
			colorizerFunc = colorizeTxtFile
		} else {
			entryInfo, _ := entry.Info()
			if entryInfo.Mode().String()[3] == 'x' {
				colorizerFunc = colorizeExecutableFile
			}
		}
	}

	return colorizerFunc(entry.Name())
}

func print8BitColorMatrix() {
	for i := 0; i <= 255; i++ {
		colorWhite := coloring.New().White().Background().Color(i).Print()
		colorBlack := coloring.New().Black().Background().Color(i).Print()

		var color coloring.ColorizerPrint

		if i < 16 {
			padding := "    "
			if i < 10 {
				padding += " "
			}

			if i < 7 {
				color = colorWhite
			} else {
				color = colorBlack
			}

			color(padding + strconv.Itoa(i) + "     ")

			if i == 7 {
				fmt.Print("    ")
			}
		} else {
			if (i-16)%36 == 0 {
				fmt.Println()
			}

			if i < 232 {
				if (i-17)-(i-16)/36*36 <= 16 {
					color = colorWhite
				} else {
					color = colorBlack
				}

				padding := " "
				if i < 100 {
					padding += " "
				}

				color(" " + strconv.Itoa(i) + padding)
			} else {
				if i < 244 {
					color = colorWhite
				} else {
					color = colorBlack
				}

				color("  " + strconv.Itoa(i) + "  ")

				if (i-231)%4 == 0 {
					fmt.Print("  ")
				}

				if i == 243 {
					fmt.Print("  ")
				}
			}
		}
	}
}

func printPhrase() {
	quick := coloring.For("quick").Green().Italic().Underline().Decoration()
	brown := coloring.For("brown").Color(15).Background().Rgb(0xB3, 0x64, 0x00).Decoration()
	fox := coloring.For("fox").White().Background().Rgb(0xE8, 0x1D, 0x1D).Decoration()
	jumps := coloring.For("jumps").Yellow().Bold().Background().Color(8).Decoration()
	dog := coloring.For("dog").Bold().Italic().Decoration()

	template := "The %s %s %s %s over the lazy %s."

	fmt.Printf(template, quick, brown, fox, jumps, dog)
	fmt.Println()
	fmt.Println()
	fmt.Printf(template, quick.Unformatted(), brown.Unformatted(), fox.Unformatted(), jumps.Unformatted(), dog.Unformatted())
}

func complexPhrases() {
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
