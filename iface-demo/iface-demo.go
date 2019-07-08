package iface_demo

type geometry interface {
	area() float64
	perimeter() float64
	name() string
}
