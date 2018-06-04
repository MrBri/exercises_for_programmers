package tip

import "math"

func Tip(v []float64) []float64 {
	tip := v[0] * (v[1] / 100)
	tip = math.Round(tip*100) / 100
	total := v[0] + tip
	return []float64{tip, total}
}
