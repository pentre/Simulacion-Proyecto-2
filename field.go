package main

import (
	"encoding/json"
	"fmt"
	"math"
)

type field struct {
	Turbines []*turbine
	Count    int
	Cost     float64
	Power    float64
	Fitness  float64
}

/* func remove(s []*turbine, i int) []turbine {
	ns := []turbine{}
	for ii, t := range s {
		if i == ii {
			continue
		}

		ns = append(ns, *t)
	}
	return ns
}

func meanU(us []turbine) float64 {
	r := 0.0
	for i := 0; i < len(us); i++ {
		r += math.Pow(1.0-(us[i].WindSpeed/uzero), 2)
	}

	return uzero * (1.0 - math.Sqrt(r))
} */

func (f *field) calculateWindSpeeds() {
	for _, t := range f.Turbines {
		t.calculateWindSpeed()
	}
}

func (f *field) calculatePower() {
	f.calculateWindSpeeds()

	p := 0.0
	for i := 0; i < f.Count; i++ {
		p += math.Pow(f.Turbines[i].WindSpeed, 3) //meanU(remove(f.Turbines, i)), 3)
	}
	f.Power = 0.3 * p
}

func (f *field) calculateCost() {
	f.Cost = float64(f.Count) * ((2.0 / 3.0) + ((1.0 / 3.0) * math.Exp(-0.00174*math.Pow(float64(f.Count), 2))))
}

func (f *field) calculateFitness() {
	f.calculateCost()
	f.calculatePower()
	f.Fitness = f.Cost / f.Power
}

func (f *field) containsTurbine(t *turbine) bool {
	for _, v := range f.Turbines {
		if v.I == t.I && v.J == t.J {
			return true
		}
	}

	return false
}

func (f *field) addTurbine(t *turbine) {
	if len(f.Turbines) == 0 {
		f.Turbines = append(f.Turbines, t)
		return
	}

	turbines := []*turbine{}
	flag := true
	for _, tur := range f.Turbines {
		if flag && t.I < tur.I || (t.I == tur.I && t.J < tur.J) {
			turbines = append(turbines, t)
			flag = false
		}

		turbines = append(turbines, tur)
	}

	f.Turbines = turbines
}

func (f *field) serialize() string {
	b, err := json.MarshalIndent(f, "", "    ")
	if err != nil {
		panic(err)
	}

	return string(b)
}

func (f *field) printTurbine(i, j int) {
	for _, t := range f.Turbines {
		if t.I == i && t.J == j {
			fmt.Print("X ")
			return
		}
	}
	fmt.Print("- ")
}

func (f *field) printField() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			f.printTurbine(i, j)
		}
		fmt.Println("")
	}
}
