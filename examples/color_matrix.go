package examples

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nelsonghezzi/go-color-term/coloring"
)

const innerBoxSize = 182

func Print8BitColorMatrix() {
	drawBoxHeader()

	for i := 0; i <= 255; i++ {
		colorWhite := coloring.New().Color(coloring.BRIGHTWHITE).Background().Color(i).Print()
		colorBlack := coloring.New().Black().Background().Color(i).Print()

		var color coloring.ColorizerPrint

		if i < 16 {
			if i == 0 {
				drawCustomContent(fmt.Sprintf("%sStandard colors%sHigh-intensity colors%s", strings.Repeat(" ", 37), strings.Repeat(" ", 75), strings.Repeat(" ", 34)))
				drawBoxOpen()
			}

			padding := "    "
			if i < 10 {
				padding += " "
			}

			if i < 8 {
				color = colorWhite
			} else {
				color = colorBlack
			}

			color(padding + strconv.Itoa(i) + "     ")

			if i == 7 {
				fmt.Print("    ")
			}

			if i == 15 {
				drawBoxEnd()
			}
		} else {
			if i == 16 {
				drawEmptyBoxLine()
				drawCenteredTitles("6 x 6 x 6 matrix (216 colors)")
			}

			if i < 232 {
				if (i-16)%36 == 0 {
					drawBoxOpen()
				}

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

				if (i-15)%36 == 0 {
					drawBoxEnd()
				}
			} else {
				if i == 232 {
					drawEmptyBoxLine()
					drawCenteredTitles("Grayscale colors")
					drawBoxOpenPadded(31)
				}

				if i < 244 {
					color = colorWhite
				} else {
					color = colorBlack
				}

				color(" " + strconv.Itoa(i) + " ")

				if i == 255 {
					drawBoxEndPadded(31)
				}
			}
		}
	}

	drawBoxFooter()
}

func drawBoxHeader() {
	fmt.Printf("┌%s┐\n", strings.Repeat("─", innerBoxSize))
}

func drawBoxFooter() {
	fmt.Printf("└%s┘\n", strings.Repeat("─", innerBoxSize))
	fmt.Println()
}

func drawEmptyBoxLine() {
	fmt.Printf("│%s│\n", strings.Repeat(" ", innerBoxSize))
}

func drawBoxOpen() {
	drawBoxOpenPadded(1)
}

func drawBoxEnd() {
	drawBoxEndPadded(1)
}

func drawBoxOpenPadded(padding int) {
	fmt.Printf("│%s", strings.Repeat(" ", padding))
}

func drawBoxEndPadded(padding int) {
	fmt.Printf("%s│\n", strings.Repeat(" ", padding))
}

func drawCenteredTitles(titles ...string) {
	titlesLen := 0
	for _, title := range titles {
		titlesLen += len(title)
	}

	paddingSize := (innerBoxSize - titlesLen) / (len(titles) + 1)
	padding := strings.Repeat(" ", paddingSize)

	fmt.Printf("│%s", padding)

	for _, title := range titles {
		fmt.Printf("%s%s", title, padding)
	}

	fmt.Printf("%s│\n", strings.Repeat(" ", innerBoxSize-titlesLen-paddingSize*(len(titles)+1)))
}

func drawCustomContent(s string) {
	fmt.Printf("│%s│\n", s)
}
