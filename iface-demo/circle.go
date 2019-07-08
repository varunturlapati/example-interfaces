package iface_demo

import (
	"fmt"
	"math"
)

type Circle struct {
	Radius float64
}

func (c Circle) Area() (float64, error) {
	if c.Radius < 0 {
		return -1, fmt.Errorf("Invalid radius")
	}
	return math.Pi * c.Radius * c.Radius, nil
}

func (c Circle) Perimeter() (float64, error) {
	if c.Radius < 0 {
		return -1, fmt.Errorf("Invalid radius")
	}
	return 2 * math.Pi * c.Radius, nil
}

func (c Circle) Name() string {
	return "Circle"
}