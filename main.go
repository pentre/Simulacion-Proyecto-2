package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const (
	z     = 60.0
	zzero = 0.3
	uzero = 12.0
	rzero = 40.0
	ct    = 0.88
)

var (
	alpha = alphaFunc()
	a     = aFunc()
	rr    = rrFunc()

	existingFields = map[string]bool{}
	bestField      = &field{Fitness: 100}
)

func alphaFunc() float64 {
	return 0.5 / math.Log(z/zzero)
}

func aFunc() float64 {
	return (1 - math.Sqrt(1-ct)) / 2.0
}

func rrFunc() float64 {
	return rzero * math.Sqrt((1-a)/(1-(2*a)))
}

func generateField() *field {
	rand.Seed(time.Now().UnixNano())
	turbinesCount := rand.Intn(100) + 1
	f := &field{
		Turbines: []*turbine{},
		Count:    turbinesCount,
	}

	for i := 0; i < turbinesCount; {
		t := &turbine{
			I: rand.Intn(10),
			J: rand.Intn(10),
		}

		if f.containsTurbine(t) {
			continue
		}

		f.Turbines = append(f.Turbines, t)
		i++
	}

	return f
}

func main() {
	t := time.Now()
	for i := 0; i < 10000000; i++ {
		f := generateField()
		s := f.serialize()
		if _, ok := existingFields[s]; ok {
			i--
			continue
		}

		//existingFields[s] = true
		f.calculateFitness()
		if f.Fitness < bestField.Fitness {
			bestField = f
		}
	}

	fmt.Println(bestField.serialize())
	bestField.printField()
	fmt.Println(time.Since(t))
}
