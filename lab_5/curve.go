package main

import (
	"math"
)

type Curve interface {
	FindY() float64
}

type Straight struct {
	X, K, B float64
}

type Ellipse struct {
	X, A, B float64
}

type Hyperbola struct {
	X, K, B float64
}

func (s Straight) FindY() float64 {
	return s.K*s.X + s.B
}

func (e Ellipse) FindY() float64 {
	sqrt_num := (1 - e.X/e.A) * (1 + e.X/e.A)

	if sqrt_num < 0 {
		return -1
	}

	return e.B * math.Sqrt(sqrt_num)
}

func (h Hyperbola) FindY() float64 {
	return h.K/h.X + h.B
}
