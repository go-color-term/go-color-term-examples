package examples

import (
	"fmt"
	"io/fs"
	"os"
	"strings"

	"github.com/go-color-term/go-color-term/coloring"
)

const depth = 3

type DirectoryListingExample struct{}

func (example *DirectoryListingExample) GetTitle() string {
	return "Directory listing"
}

func (example *DirectoryListingExample) GetDescription() string {
	return fmt.Sprintf("Lists the contents of the current directory up to %d folders depth.", depth-1)
}

func (example *DirectoryListingExample) Run() {
	listContents(".", 0, depth)
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
