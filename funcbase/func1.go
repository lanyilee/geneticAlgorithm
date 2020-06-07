package funcbase

import "math"

//y =x1^3+x2^2-5*x1*x2
func GetFuncValue(a int, b int) int {
	c := math.Pow(float64(a), 3)
	d := math.Pow(float64(b), 2)
	e := 3 * a * b
	return int(c) + int(d) - e
}
