package math
import "fmt"

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

func (m *Matrix3) GetNormalMatrix(matrix *Matrix4) (*Matrix3) {
	m.GetInverse(matrix, false).Transpose()
	return m
}

func (m *Matrix3) MultiplyScalar(s float64) (*Matrix3) {
	te := m.Elements

	te[ 0 ] *= s; te[ 3 ] *= s; te[ 6 ] *= s;
	te[ 1 ] *= s; te[ 4 ] *= s; te[ 7 ] *= s;
	te[ 2 ] *= s; te[ 5 ] *= s; te[ 8 ] *= s;

	return m
}

func (m *Matrix3) Determinant() float64 {

	te := m.Elements

	a := te[ 0 ]; b := te[ 1 ]; c := te[ 2 ]
	d := te[ 3 ]; e := te[ 4 ]; f := te[ 5 ]
	g := te[ 6 ]; h := te[ 7 ]; i := te[ 8 ]

	return a * e * i - a * f * h - b * d * i + b * f * g + c * d * h - c * e * g
}

func (m *Matrix3) GetInverse(matrix *Matrix4, panicOnInvertible bool) (*Matrix3) {
	// ( based on http://code.google.com/p/webgl-mjs/ )
	me := matrix.Elements
	te := m.Elements

	te[ 0 ] =   me[ 10 ] * me[ 5 ] - me[ 6 ] * me[ 9 ]
	te[ 1 ] = - me[ 10 ] * me[ 1 ] + me[ 2 ] * me[ 9 ]
	te[ 2 ] =   me[ 6 ] * me[ 1 ] - me[ 2 ] * me[ 5 ]
	te[ 3 ] = - me[ 10 ] * me[ 4 ] + me[ 6 ] * me[ 8 ]
	te[ 4 ] =   me[ 10 ] * me[ 0 ] - me[ 2 ] * me[ 8 ]
	te[ 5 ] = - me[ 6 ] * me[ 0 ] + me[ 2 ] * me[ 4 ]
	te[ 6 ] =   me[ 9 ] * me[ 4 ] - me[ 5 ] * me[ 8 ]
	te[ 7 ] = - me[ 9 ] * me[ 0 ] + me[ 1 ] * me[ 8 ]
	te[ 8 ] =   me[ 5 ] * me[ 0 ] - me[ 1 ] * me[ 4 ]

	det := me[ 0 ] * te[ 0 ] + me[ 1 ] * te[ 3 ] + me[ 2 ] * te[ 6 ]

	// no inverse
	if det == 0 {
		msg := "Matrix3.getInverse(): can't invert matrix, determinant is 0"
		if panicOnInvertible {
			panic( msg )
		} else {
			fmt.Println( msg )
		}
		m.Identity()
		return m
	}
	m.MultiplyScalar( 1.0 / det )
	return m
}

func (mm *Matrix3) Transpose() (*Matrix3) {
	var tmp float64
	m := mm.Elements

	tmp = m[ 1 ]; m[ 1 ] = m[ 3 ]; m[ 3 ] = tmp
	tmp = m[ 2 ]; m[ 2 ] = m[ 6 ]; m[ 6 ] = tmp
	tmp = m[ 5 ]; m[ 5 ] = m[ 7 ]; m[ 7 ] = tmp

	return mm
}

func (m *Matrix3) Set(n11, n12, n13, n21, n22, n23, n31, n32, n33 float64) (*Matrix3) {
	var te = m.Elements

	te[ 0 ] = n11; te[ 3 ] = n12; te[ 6 ] = n13;
	te[ 1 ] = n21; te[ 4 ] = n22; te[ 7 ] = n23;
	te[ 2 ] = n31; te[ 5 ] = n32; te[ 8 ] = n33;

	return m
}

func (m *Matrix3) Identity() (*Matrix3) {
	m.Set(
		1, 0, 0,
		0, 1, 0,
		0, 0, 1,
	)
	return m
}

func (m *Matrix3) Clone() (*Matrix3) {
	me := m.Elements
	return NewMatrix3().Set(
			me[ 0 ], me[ 3 ], me[ 6 ],
			me[ 1 ], me[ 4 ], me[ 7 ],
			me[ 2 ], me[ 5 ], me[ 8 ],
		)
}

func (m *Matrix3) Copy(matrix *Matrix3) (*Matrix3) {
	me := matrix.Elements
	m.Set(
		me[ 0 ], me[ 3 ], me[ 6 ],
		me[ 1 ], me[ 4 ], me[ 7 ],
		me[ 2 ], me[ 5 ], me[ 8 ],
	)
	return m
}

