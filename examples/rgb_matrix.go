package examples

import (
	"fmt"

	"github.com/go-color-term/go-color-term/coloring"
)

func Rgb_matrix() {
	index := 0

	for i := 0.0005; i < 1; i += 0.0025 {
		index++

		r, g, b := Hsl2Rgb(i, 0.5, 0.5)

		fmt.Print(coloring.For(" ").Background().Rgb(r, g, b))

		if index%50 == 0 {
			fmt.Println()
		}
	}
}

func Hsl2Rgb(h, sl, l float64) (int, int, int) {
	var v float64
	var r, g, b float64

	r = l
	g = l
	b = l

	if l <= 0.5 {
		v = l * (1.0 + sl)
	} else {
		v = l + sl - l*sl
	}

	if v > 0 {
		var m float64
		var sv float64
		var sextant int
		var fract, vsf, mid1, mid2 float64

		m = l + l - v
		sv = (v - m) / v
		h *= 6.0

		sextant = int(h)
		fract = h - float64(sextant)
		vsf = v * sv * fract
		mid1 = m + vsf
		mid2 = v - vsf

		switch sextant {
		case 0:
			r = v
			g = mid1
			b = m
		case 1:
			r = mid2
			g = v
			b = m
		case 2:
			r = m
			g = v
			b = mid1
		case 3:
			r = m
			g = mid2
			b = v
		case 4:
			r = mid1
			g = m
			b = v
		case 5:
			r = v
			g = m
			b = mid2
		}
	}

	return int(r * 255.0), int(g * 255.0), int(b * 255.0)
}
