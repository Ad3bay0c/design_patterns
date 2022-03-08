package srp

type Rectangle struct {
	breadth int
	length	int
}

type Square struct {
	length int
}

func (r *Rectangle) Area() int {
	return r.breadth * r.length
}

func (s *Square) Area() int {
	return s.length * s.length
}

func sumShapes() int {
	r := Rectangle{breadth: 10, length: 5}
	s := Square{length: 5}
	return r.Area() + s.Area()
}