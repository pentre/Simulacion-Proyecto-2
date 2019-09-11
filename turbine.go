package main

import "math"

type turbine struct {
	I         int
	J         int
	WindSpeed float64
}

func (t *turbine) calculateWindSpeed() {
	t.WindSpeed = uzero * (1.0 - ((a) / math.Pow(1+(alpha*(float64(200*(t.I+1))/rr)), 2)))
}
