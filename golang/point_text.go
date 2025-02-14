package sprint

import "fmt"

type Point4 struct {
	X    float32
	Y    float32
	Text string
}

func PointText(p Point4) Point4 {
	return Point4{p.X, p.Y, fmt.Sprintf("Text at (%f, %f)", p.X, p.Y)}
}
