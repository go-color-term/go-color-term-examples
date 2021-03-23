package examples

import (
	"fmt"

	"github.com/nelsonghezzi/go-color-term/coloring"
)

func TestDrive() {
	fmt.Println(coloring.Black("Black"))
	fmt.Println(coloring.Red("Red"))
	fmt.Println(coloring.Green("Green"))
	fmt.Println(coloring.Yellow("Yellow"))
	fmt.Println(coloring.Blue("Blue"))
	fmt.Println(coloring.Magenta("Magenta"))
	fmt.Println(coloring.Cyan("Cyan"))
	fmt.Println(coloring.White("White"))

	fmt.Println(coloring.BgBlack("Black background"))
	fmt.Println(coloring.BgRed("Red background"))
	fmt.Println(coloring.BgGreen("Green background"))
	fmt.Println(coloring.BgYellow("Yellow background"))
	fmt.Println(coloring.BgBlue("Blue background"))
	fmt.Println(coloring.BgMagenta("Magenta background"))
	fmt.Println(coloring.BgCyan("Cyan background"))
	fmt.Println(coloring.BgWhite("White background"))

	fmt.Println(coloring.Bold("Bold"))
	fmt.Println(coloring.Faint("Faint"))
	fmt.Println(coloring.Italic("Italic"))
	fmt.Println(coloring.Underline("Underline"))
	fmt.Println(coloring.Blink("Blink"))
	fmt.Println(coloring.Invert("Invert"))
	fmt.Println(coloring.Conceal("Conceal"))
	fmt.Println(coloring.Strikethrough("Strikethrough"))

	fmt.Println(coloring.Extras.BrightBlack("Bright black"))
	fmt.Println(coloring.Extras.BrightRed("Bright red"))
	fmt.Println(coloring.Extras.BrightGreen("Bright green"))
	fmt.Println(coloring.Extras.BrightYellow("Bright yellow"))
	fmt.Println(coloring.Extras.BrightBlue("Bright blue"))
	fmt.Println(coloring.Extras.BrightMagenta("Bright magenta"))
	fmt.Println(coloring.Extras.BrightCyan("Bright cyan"))
	fmt.Println(coloring.Extras.BrightWhite("Bright white"))

	fmt.Println(coloring.Extras.BgBrightBlack("Bright black background"))
	fmt.Println(coloring.Extras.BgBrightRed("Bright red background"))
	fmt.Println(coloring.Extras.BgBrightGreen("Bright green background"))
	fmt.Println(coloring.Extras.BgBrightYellow("Bright yellow background"))
	fmt.Println(coloring.Extras.BgBrightBlue("Bright blue background"))
	fmt.Println(coloring.Extras.BgBrightMagenta("Bright magenta background"))
	fmt.Println(coloring.Extras.BgBrightCyan("Bright cyan background"))
	fmt.Println(coloring.Extras.BgBrightWhite("Bright white background"))
}
