package float

import "math"

func round(val float64, roundOn float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
}

//Round float to two places after separator, breaking on 0.5
func Round(val float64) (roundedVal float64) {
	return round(val, 0.5, 2)
}
