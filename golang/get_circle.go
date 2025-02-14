package sprint

type Circle struct {
	Radius    float32
	Diameter  float32
	Area      float32
	Perimeter float32
}

func GetCircle(r float32) Circle {
	var pi float32 = 3.14
	var Radius float32 = r
	var Diameter float32 = 2 * r
	var Area float32 = r * r * pi
	var Perimeter float32 = pi * 2 * r
	return Circle{
		Radius:    Radius,
		Diameter:  Diameter,
		Area:      Area,
		Perimeter: Perimeter,
	}
}
