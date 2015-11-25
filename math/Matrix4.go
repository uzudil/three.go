package math
import "math"

type Matrix4 struct {
	Elements [16]float64
	Decompose func(*Vector3, *Quaternion, *Vector3) (*Matrix4)
	LookAt func(*Vector3, *Vector3, *Vector3) (*Matrix4)
}

func NewMatrix4() (*Matrix4) {
	m := &Matrix4{
		Elements: [16]float64{
			1, 0, 0, 0,
			0, 1, 0, 0,
			0, 0, 1, 0,
			0, 0, 0, 1,
		},
	}
	m.Decompose = m.buildDecompose()
	m.LookAt = m.buildLookAt()
	return m
}

func (m *Matrix4) MakeRotationFromEuler(euler *Euler) (*Matrix4) {
	te := m.Elements

	x := euler.X
	y := euler.Y
	z := euler.Z
	a := math.Cos( x )
	b := math.Sin( x )
	c := math.Cos( y )
	d := math.Sin( y )
	e := math.Cos( z )
	f := math.Sin( z )

	if euler.Order == "XYZ" {

		ae := a * e
		af := a * f
		be := b * e
		bf := b * f

		te[ 0 ] = c * e
		te[ 4 ] = - c * f
		te[ 8 ] = d

		te[ 1 ] = af + be * d
		te[ 5 ] = ae - bf * d
		te[ 9 ] = - b * c

		te[ 2 ] = bf - ae * d
		te[ 6 ] = be + af * d
		te[ 10 ] = a * c

	} else if euler.Order == "YXZ" {

		ce := c * e
		cf := c * f
		de := d * e
		df := d * f

		te[ 0 ] = ce + df * b
		te[ 4 ] = de * b - cf
		te[ 8 ] = a * d

		te[ 1 ] = a * f
		te[ 5 ] = a * e
		te[ 9 ] = - b

		te[ 2 ] = cf * b - de
		te[ 6 ] = df + ce * b
		te[ 10 ] = a * c

	} else if euler.Order == "ZXY" {

		ce := c * e
		cf := c * f
		de := d * e
		df := d * f

		te[ 0 ] = ce - df * b
		te[ 4 ] = - a * f
		te[ 8 ] = de + cf * b

		te[ 1 ] = cf + de * b
		te[ 5 ] = a * e
		te[ 9 ] = df - ce * b

		te[ 2 ] = - a * d
		te[ 6 ] = b
		te[ 10 ] = a * c

	} else if euler.Order == "ZYX" {

		ae := a * e
		af := a * f
		be := b * e
		bf := b * f

		te[ 0 ] = c * e
		te[ 4 ] = be * d - af
		te[ 8 ] = ae * d + bf

		te[ 1 ] = c * f
		te[ 5 ] = bf * d + ae
		te[ 9 ] = af * d - be

		te[ 2 ] = - d
		te[ 6 ] = b * c
		te[ 10 ] = a * c

	} else if euler.Order == "YZX" {

		ac := a * c
		ad := a * d
		bc := b * c
		bd := b * d

		te[ 0 ] = c * e
		te[ 4 ] = bd - ac * f
		te[ 8 ] = bc * f + ad

		te[ 1 ] = f
		te[ 5 ] = a * e
		te[ 9 ] = - b * e

		te[ 2 ] = - d * e
		te[ 6 ] = ad * f + bc
		te[ 10 ] = ac - bd * f

	} else if euler.Order == "XZY" {

		ac := a * c
		ad := a * d
		bc := b * c
		bd := b * d

		te[ 0 ] = c * e
		te[ 4 ] = - f
		te[ 8 ] = d * e

		te[ 1 ] = ac * f + bd
		te[ 5 ] = a * e
		te[ 9 ] = ad * f - bc

		te[ 2 ] = bc * f - ad
		te[ 6 ] = b * e
		te[ 10 ] = bd * f + ac

	}

	// last column
	te[ 3 ] = 0
	te[ 7 ] = 0
	te[ 11 ] = 0

	// bottom row
	te[ 12 ] = 0
	te[ 13 ] = 0
	te[ 14 ] = 0
	te[ 15 ] = 1

	return m

}

func (m *Matrix4) MakeRotationFromQuaternion(q *Quaternion) (*Matrix4) {
	te := m.Elements

	x := q.X
	y := q.Y
	z := q.Z
	w := q.W
	x2 := x + x
	y2 := y + y
	z2 := z + z
	xx := x * x2
	xy := x * y2
	xz := x * z2
	yy := y * y2
	yz := y * z2
	zz := z * z2
	wx := w * x2
	wy := w * y2
	wz := w * z2

	te[ 0 ] = 1 - ( yy + zz )
	te[ 4 ] = xy - wz
	te[ 8 ] = xz + wy

	te[ 1 ] = xy + wz
	te[ 5 ] = 1 - ( xx + zz )
	te[ 9 ] = yz - wx

	te[ 2 ] = xz - wy
	te[ 6 ] = yz + wx
	te[ 10 ] = 1 - ( xx + yy )

	// last column
	te[ 3 ] = 0
	te[ 7 ] = 0
	te[ 11 ] = 0

	// bottom row
	te[ 12 ] = 0
	te[ 13 ] = 0
	te[ 14 ] = 0
	te[ 15 ] = 1

	return m
}

func (m *Matrix4) Clone() (*Matrix4) {
	return NewMatrix4().FromArray( m.Elements )
}

func (m *Matrix4) Copy(matrix *Matrix4) (*Matrix4) {
	m.Elements = append(m.Elements, matrix.Elements...)
	return m
}

func (m *Matrix4) Equals(matrix *Matrix4) bool {
	te := m.Elements
	me := matrix.Elements

	for i := 0; i < 16; i++ {
		if te[ i ] != me[ i ] {
			return false
		}
	}
	return true
}

func (m *Matrix4) FromArray(array []float64) (*Matrix4) {
	m.Elements = append(m.Elements, array[:]...)
	return m
}

func (m *Matrix4) ToArray() {
	var te = m.Elements
	return []float64{
		te[ 0 ], te[ 1 ], te[ 2 ], te[ 3 ],
		te[ 4 ], te[ 5 ], te[ 6 ], te[ 7 ],
		te[ 8 ], te[ 9 ], te[ 10 ], te[ 11 ],
		te[ 12 ], te[ 13 ], te[ 14 ], te[ 15 ],
	}
}

func (m *Matrix4) Scale(v *Vector3) (*Matrix4) {
	te := m.Elements
	x := v.X
	y := v.Y
	z := v.Z

	te[ 0 ] *= x; te[ 4 ] *= y; te[ 8 ] *= z
	te[ 1 ] *= x; te[ 5 ] *= y; te[ 9 ] *= z
	te[ 2 ] *= x; te[ 6 ] *= y; te[ 10 ] *= z
	te[ 3 ] *= x; te[ 7 ] *= y; te[ 11 ] *= z

	return m
}

func (m *Matrix4) SetPosition(v *Vector3) (*Matrix4) {
	var te = m.Elements

	te[ 12 ] = v.X
	te[ 13 ] = v.Y
	te[ 14 ] = v.Z

	return m
}

func (m *Matrix4) Compose( position *Vector3, quaternion *Quaternion, scale *Vector3) (*Matrix4) {
	m.MakeRotationFromQuaternion( quaternion )
	m.Scale( scale )
	m.SetPosition( position )
	return m
}

func (m *Matrix4) buildDecompose() (func(*Vector3, *Quaternion, *Vector3) (*Matrix4)) {

	var vector *Vector3 = nil
	var matrix *Matrix4 = nil

	return func(position *Vector3, quaternion *Quaternion, scale *Vector3) (*Matrix4) {

		if vector == nil {
			vector = NewEmptyVector3()
		}
		if matrix == nil {
			matrix = NewMatrix4()
		}

		var te = m.Elements

		var sx = vector.Set(te[ 0 ], te[ 1 ], te[ 2 ]).Length()
		var sy = vector.Set(te[ 4 ], te[ 5 ], te[ 6 ]).Length()
		var sz = vector.Set(te[ 8 ], te[ 9 ], te[ 10 ]).Length()

		// if determine is negative, we need to invert one scale
		var det = m.Determinant()
		if ( det < 0 ) {
			sx = - sx;
		}

		position.X = te[ 12 ]
		position.Y = te[ 13 ]
		position.Z = te[ 14 ]

		// scale the rotation part
		copy(matrix.Elements, m.Elements) // at this point matrix is incomplete so we can't use .copy()

		invSX := 1 / sx
		invSY := 1 / sy
		invSZ := 1 / sz

		matrix.Elements[ 0 ] *= invSX
		matrix.Elements[ 1 ] *= invSX
		matrix.Elements[ 2 ] *= invSX

		matrix.Elements[ 4 ] *= invSY
		matrix.Elements[ 5 ] *= invSY
		matrix.Elements[ 6 ] *= invSY

		matrix.Elements[ 8 ] *= invSZ
		matrix.Elements[ 9 ] *= invSZ
		matrix.Elements[ 10 ] *= invSZ

		quaternion.SetFromRotationMatrix(matrix)

		scale.X = sx
		scale.Y = sy
		scale.Z = sz

		return m
	}
}

func (m *Matrix4) Determinant() float64 {

	te := m.Elements

	n11 := te[ 0 ]; n12 := te[ 4 ]; n13 := te[ 8 ]; n14 := te[ 12 ]
	n21 := te[ 1 ]; n22 := te[ 5 ]; n23 := te[ 9 ]; n24 := te[ 13 ]
	n31 := te[ 2 ]; n32 := te[ 6 ]; n33 := te[ 10 ]; n34 := te[ 14 ]
	n41 := te[ 3 ]; n42 := te[ 7 ]; n43 := te[ 11 ]; n44 := te[ 15 ]

	//TODO: make this more efficient
	//( based on http://www.euclideanspace.com/maths/algebra/matrix/functions/inverse/fourD/index.htm )

	p1 := n41 * (+ n14 * n23 * n32 - n13 * n24 * n32 - n14 * n22 * n33 + n12 * n24 * n33 + n13 * n22 * n34 - n12 * n23 * n34)
	p2 := n42 * (+ n11 * n23 * n34 - n11 * n24 * n33 + n14 * n21 * n33 - n13 * n21 * n34 + n13 * n24 * n31 - n14 * n23 * n31)
	p3 := n43 * (+ n11 * n24 * n32 - n11 * n22 * n34 - n14 * n21 * n32 + n12 * n21 * n34 + n14 * n22 * n31 - n12 * n24 * n31)
	p4 := n44 * (- n13 * n22 * n31 - n11 * n23 * n32 + n11 * n22 * n33 + n13 * n21 * n32 - n12 * n21 * n33 + n12 * n23 * n31)

	return p1 + p2 + p3 + p4
}

func (m *Matrix4) MultiplyMatrices( a, b *Matrix4) (*Matrix4) {

	ae := a.Elements
	be := b.Elements
	te := m.Elements

	a11 := ae[ 0 ]; a12 := ae[ 4 ]; a13 := ae[ 8 ]; a14 := ae[ 12 ]
	a21 := ae[ 1 ]; a22 := ae[ 5 ]; a23 := ae[ 9 ]; a24 := ae[ 13 ]
	a31 := ae[ 2 ]; a32 := ae[ 6 ]; a33 := ae[ 10 ]; a34 := ae[ 14 ]
	a41 := ae[ 3 ]; a42 := ae[ 7 ]; a43 := ae[ 11 ]; a44 := ae[ 15 ]

	b11 := be[ 0 ]; b12 := be[ 4 ]; b13 := be[ 8 ]; b14 := be[ 12 ]
	b21 := be[ 1 ]; b22 := be[ 5 ]; b23 := be[ 9 ]; b24 := be[ 13 ]
	b31 := be[ 2 ]; b32 := be[ 6 ]; b33 := be[ 10 ]; b34 := be[ 14 ]
	b41 := be[ 3 ]; b42 := be[ 7 ]; b43 := be[ 11 ]; b44 := be[ 15 ]

	te[ 0 ] = a11 * b11 + a12 * b21 + a13 * b31 + a14 * b41
	te[ 4 ] = a11 * b12 + a12 * b22 + a13 * b32 + a14 * b42
	te[ 8 ] = a11 * b13 + a12 * b23 + a13 * b33 + a14 * b43
	te[ 12 ] = a11 * b14 + a12 * b24 + a13 * b34 + a14 * b44

	te[ 1 ] = a21 * b11 + a22 * b21 + a23 * b31 + a24 * b41
	te[ 5 ] = a21 * b12 + a22 * b22 + a23 * b32 + a24 * b42
	te[ 9 ] = a21 * b13 + a22 * b23 + a23 * b33 + a24 * b43
	te[ 13 ] = a21 * b14 + a22 * b24 + a23 * b34 + a24 * b44

	te[ 2 ] = a31 * b11 + a32 * b21 + a33 * b31 + a34 * b41
	te[ 6 ] = a31 * b12 + a32 * b22 + a33 * b32 + a34 * b42
	te[ 10 ] = a31 * b13 + a32 * b23 + a33 * b33 + a34 * b43
	te[ 14 ] = a31 * b14 + a32 * b24 + a33 * b34 + a34 * b44

	te[ 3 ] = a41 * b11 + a42 * b21 + a43 * b31 + a44 * b41
	te[ 7 ] = a41 * b12 + a42 * b22 + a43 * b32 + a44 * b42
	te[ 11 ] = a41 * b13 + a42 * b23 + a43 * b33 + a44 * b43
	te[ 15 ] = a41 * b14 + a42 * b24 + a43 * b34 + a44 * b44

	return m
}

func (m *Matrix4) buildLookAt() (func(*Vector3, *Vector3, *Vector3) (*Matrix4)) {

	var x, y, z Vector3

	return func(eye, target, up *Vector3) (*Matrix4) {

		if x == nil {
			x = NewEmptyVector3()
		}
		if y == nil {
			y = NewEmptyVector3()
		}
		if z == nil {
			z = NewEmptyVector3()
		}

		var te = m.Elements

		z.SubVectors( eye, target ).Normalize()

		if z.LengthSq() == 0 {
			z.Z = 1
		}

		x.CrossVectors( up, z ).Normalize()

		if x.LengthSq() == 0 {
			z.X += 0.0001
			x.CrossVectors( up, z ).Normalize()
		}

		y.CrossVectors( z, x )


		te[ 0 ] = x.X; te[ 4 ] = y.X; te[ 8 ] = z.X
		te[ 1 ] = x.Y; te[ 5 ] = y.Y; te[ 9 ] = z.Y
		te[ 2 ] = x.Z; te[ 6 ] = y.Z; te[ 10 ] = z.Z

		return m
	}
}
