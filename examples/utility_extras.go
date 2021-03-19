package examples

import (
	"fmt"

	"github.com/nelsonghezzi/go-color-term/coloring"
)

func BrightColors() {
	fmt.Printf("%s    %s\n", coloring.BgWhite(coloring.Black("Normal black   ")), coloring.BgWhite(coloring.Extras.BrightBlack("bright black   ")))
	fmt.Printf("%s    %s\n", coloring.Red("Normal red     "), coloring.Extras.BrightRed("bright red     "))
	fmt.Printf("%s    %s\n", coloring.Green("Normal green   "), coloring.Extras.BrightGreen("bright green   "))
	fmt.Printf("%s    %s\n", coloring.Yellow("Normal yellow  "), coloring.Extras.BrightYellow("bright yellow  "))
	fmt.Printf("%s    %s\n", coloring.Blue("Normal blue    "), coloring.Extras.BrightBlue("bright blue    "))
	fmt.Printf("%s    %s\n", coloring.Magenta("Normal magenta "), coloring.Extras.BrightMagenta("bright magenta "))
	fmt.Printf("%s    %s\n", coloring.Cyan("Normal cyan    "), coloring.Extras.BrightCyan("bright cyan    "))
	fmt.Printf("%s    %s\n", coloring.White("Normal white   "), coloring.Extras.BrightWhite("bright white   "))

	fmt.Println()

	fmt.Printf("%s    %s\n", coloring.BgBlack("Normal black   "), coloring.Extras.BgBrightBlack("bright black   "))
	fmt.Printf("%s    %s\n", coloring.Black(coloring.BgRed("Normal red     ")), coloring.Black(coloring.Extras.BgBrightRed("bright red     ")))
	fmt.Printf("%s    %s\n", coloring.Black(coloring.BgGreen("Normal green   ")), coloring.Black(coloring.Extras.BgBrightGreen("bright green   ")))
	fmt.Printf("%s    %s\n", coloring.Black(coloring.BgYellow("Normal yellow  ")), coloring.Black(coloring.Extras.BgBrightYellow("bright yellow  ")))
	fmt.Printf("%s    %s\n", coloring.Black(coloring.BgBlue("Normal blue    ")), coloring.Black(coloring.Extras.BgBrightBlue("bright blue    ")))
	fmt.Printf("%s    %s\n", coloring.Black(coloring.BgMagenta("Normal magenta ")), coloring.Black(coloring.Extras.BgBrightMagenta("bright magenta ")))
	fmt.Printf("%s    %s\n", coloring.Black(coloring.BgCyan("Normal cyan    ")), coloring.Black(coloring.Extras.BgBrightCyan("bright cyan    ")))
	fmt.Printf("%s    %s\n", coloring.Black(coloring.BgWhite("Normal white   ")), coloring.Black(coloring.Extras.BgBrightWhite("bright white   ")))
}
