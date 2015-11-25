package math

type Matrix3 struct {
	Elements [9]float64
}

func NewMatrix3() (*Matrix3) {
	return &Matrix3{
		Elements: [9]float64{
			1, 0, 0,
			0, 1, 0,
			0, 0, 1,
		},
	}
}
