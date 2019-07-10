package main

import (
	"example-interfaces/iface-demo"
	"fmt"
)

func main() {
	fmt.Printf("hello, world\n")
	mycircle := iface_demo.Circle{
		Radius: 1.0,
	}
	mytrapezium := iface_demo.Trapezium{
		Long: 10,
		Short: 8,
		Height: 20,
	}
	area, _ := mycircle.Area()
	fmt.Printf("Area of %s mycircle is %f\n", mycircle.Name(), area)
	trapeziumArea, _ := mytrapezium.Area()
	fmt.Printf("Area of Trapezium is %f\n", trapeziumArea)

}