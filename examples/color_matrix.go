package examples

import (
	"fmt"
	"strconv"

	"github.com/nelsonghezzi/go-color-term/coloring"
)

func Print8BitColorMatrix() {
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
