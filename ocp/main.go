package ocp

type Rectangle struct {
	breadth int
	length	int
}

type Square struct {
	length int
}

type Shape interface {
	Area() int
}

func (r *Rectangle) Area() int {
	return r.breadth * r.length
}

func (s *Square) Area() int {
	return s.length * s.length
}

func sumShapes(shapes []Shape) int {
	var sum int
	for _, shape := range shapes {
		sum += shape.Area()
	}
	return sum
}