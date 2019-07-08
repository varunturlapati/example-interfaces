package iface_demo

import (
	"fmt"
)

type Trapezium struct {
	Long float64
	Short float64
	Height float64
}

func (t Trapezium) Area() (float64, error) {
	if t.Long < 0 || t.Short < 0 || t.Height < 0 {
		return -1, fmt.Errorf("Invalid geometrical parameters. All attrs must be 0 or greater")
	}
	return 0.5 * t.Height * (t.Short + t.Long), nil
}
