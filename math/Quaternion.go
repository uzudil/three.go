package math
import "math"

type Quaternion struct {
	X, Y, Z, W float64
	onChangeCallback func()
}

func NewEmptyQuaternion() (*Quaternion) {
	return NewQuaternion(0, 0, 0, 1)
}

func NewQuaternion(x, y, z, w float64) (*Quaternion) {
	return &Quaternion(x, y, z, w, func() {})
}

func (e *Quaternion) SetX(value float64) {
	e.X = value
	e.onChangeCallback()
}

func (e *Quaternion) SetY(value float64) {
	e.Y = value
	e.onChangeCallback()
}

func (e *Quaternion) SetZ(value float64) {
	e.Z = value
	e.onChangeCallback()
}

func (e *Quaternion) SetW(value float64) {
	e.W = value
	e.onChangeCallback()
}

func (e *Quaternion) Set(x, y, z, w float64) (*Quaternion) {
	e.X = x
	e.Y = y
	e.Z = z
	e.W = w

	e.onChangeCallback()

	return e

}

func (e *Quaternion) Clone() (*Quaternion) {
	return NewQuaternion(e.X, e.Y, e.Z, e.W)
}

func (e *Quaternion) Copy(quat *Quaternion) (*Quaternion) {
	e.X = quat.X
	e.Y = quat.Y
	e.Z = quat.Z
	e.W = quat.W

	e.onChangeCallback()

	return e
}

func (q *Quaternion) SetFromEuler(euler *Euler, update bool) (*Quaternion) {

	// http://www.mathworks.com/matlabcentral/fileexchange/
	// 	20696-function-to-convert-between-dcm-euler-angles-quaternions-and-euler-vectors/
	//	content/SpinCalc.m

	c1 := math.Cos( euler.X / 2 )
	c2 := math.Cos( euler.Y / 2 )
	c3 := math.Cos( euler.Z / 2 )
	s1 := math.Sin( euler.X / 2 )
	s2 := math.Sin( euler.Y / 2 )
	s3 := math.Sin( euler.Z / 2 )

	order := euler.order;

	if order == "XYZ" {

		q.X = s1 * c2 * c3 + c1 * s2 * s3
		q.Y = c1 * s2 * c3 - s1 * c2 * s3
		q.Z = c1 * c2 * s3 + s1 * s2 * c3
		q.W = c1 * c2 * c3 - s1 * s2 * s3

	} else if order == "YXZ" {

		q.X = s1 * c2 * c3 + c1 * s2 * s3
		q.Y = c1 * s2 * c3 - s1 * c2 * s3
		q.Z = c1 * c2 * s3 - s1 * s2 * c3
		q.W = c1 * c2 * c3 + s1 * s2 * s3

	} else if order == "ZXY" {

		q.X = s1 * c2 * c3 - c1 * s2 * s3
		q.Y = c1 * s2 * c3 + s1 * c2 * s3
		q.Z = c1 * c2 * s3 + s1 * s2 * c3
		q.W = c1 * c2 * c3 - s1 * s2 * s3

	} else if order == "ZYX" {

		q.X = s1 * c2 * c3 - c1 * s2 * s3
		q.Y = c1 * s2 * c3 + s1 * c2 * s3
		q.Z = c1 * c2 * s3 - s1 * s2 * c3
		q.W = c1 * c2 * c3 + s1 * s2 * s3

	} else if order == "YZX" {

		q.X = s1 * c2 * c3 + c1 * s2 * s3
		q.Y = c1 * s2 * c3 + s1 * c2 * s3
		q.Z = c1 * c2 * s3 - s1 * s2 * c3
		q.W = c1 * c2 * c3 - s1 * s2 * s3

	} else if order == "XZY" {

		q.X = s1 * c2 * c3 - c1 * s2 * s3
		q.Y = c1 * s2 * c3 - s1 * c2 * s3
		q.Z = c1 * c2 * s3 + s1 * s2 * c3
		q.W = c1 * c2 * c3 + s1 * s2 * s3

	}

	if update != false {
		q.onChangeCallback()
	}

	return q
}

func (q *Quaternion) SetFromAxisAngle(axis *Vector3, angle float64) (*Quaternion) {

	// http://www.euclideanspace.com/maths/geometry/rotations/conversions/angleToQuaternion/index.htm

	// assumes axis is normalized

	halfAngle := angle / 2
	s := math.Sin( halfAngle )

	q.X = axis.X * s
	q.Y = axis.Y * s
	q.Z = axis.Z * s
	q.W = math.Cos( halfAngle )

	q.onChangeCallback()

	return q
}

func (q *Quaternion) SetFromRotationMatrix(m *Matrix4) (*Quaternion) {

	// http://www.euclideanspace.com/maths/geometry/rotations/conversions/matrixToQuaternion/index.htm

	// assumes the upper 3x3 of m is a pure rotation matrix (i.e, unscaled)

	te := m.Elements

	m11 := te[ 0 ]
	m12 := te[ 4 ]
	m13 := te[ 8 ]
	m21 := te[ 1 ]
	m22 := te[ 5 ]
	m23 := te[ 9 ]
	m31 := te[ 2 ]
	m32 := te[ 6 ]
	m33 := te[ 10 ]

	trace := m11 + m22 + m33
	var s float64

	if trace > 0 {

		s = 0.5 / math.Sqrt( trace + 1.0 )

		q.W = 0.25 / s
		q.X = ( m32 - m23 ) * s
		q.Y = ( m13 - m31 ) * s
		q.Z = ( m21 - m12 ) * s

	} else if m11 > m22 && m11 > m33 {

		s = 2.0 * math.Sqrt( 1.0 + m11 - m22 - m33 )

		q.W = ( m32 - m23 ) / s
		q.X = 0.25 * s
		q.Y = ( m12 + m21 ) / s
		q.Z = ( m13 + m31 ) / s

	} else if m22 > m33 {

		s = 2.0 * math.Sqrt( 1.0 + m22 - m11 - m33 )

		q.W = ( m13 - m31 ) / s
		q.X = ( m12 + m21 ) / s
		q.Y = 0.25 * s
		q.Z = ( m23 + m32 ) / s

	} else {

		s = 2.0 * math.Sqrt( 1.0 + m33 - m11 - m22 )

		q.W = ( m21 - m12 ) / s
		q.X = ( m13 + m31 ) / s
		q.Y = ( m23 + m32 ) / s
		q.Z = 0.25 * s

	}

	q.onChangeCallback();

	return q
}

/*
func (q *Quaternion) SetFromUnitVectors() (*Quaternion) {

	// http://lolengine.net/blog/2014/02/24/quaternion-from-two-vectors-final

	// assumes direction vectors vFrom and vTo are normalized

	var v1, r;

	var EPS = 0.000001;

	return function ( vFrom, vTo ) {

		if ( v1 === undefined ) v1 = new THREE.Vector3();

		r = vFrom.dot( vTo ) + 1;

		if ( r < EPS ) {

			r = 0;

			if ( Math.abs( vFrom.x ) > Math.abs( vFrom.z ) ) {

				v1.set( - vFrom.y, vFrom.x, 0 );

			} else {

				v1.set( 0, - vFrom.z, vFrom.y );

			}

		} else {

			v1.crossVectors( vFrom, vTo );

		}

		q.X = v1.x;
		q.Y = v1.y;
		q.Z = v1.z;
		q.W = r;

		this.normalize();

		return this;

	}

}(),
*/
func (q *Quaternion) Inverse() (*Quaternion) {
	q.Conjugate().Normalize()
	return q
}

func (q *Quaternion) Conjugate() (*Quaternion) {
	q.X *= - 1
	q.Y *= - 1
	q.Z *= - 1

	q.onChangeCallback()

	return q
}

func (q *Quaternion) Dot(v *Quaternion) float64 {
	return q.X * v.X + q.Y * v.Y + q.Z * v.Z + q.W * v.W
}

func (q *Quaternion) LengthSq() float64 {
	return q.X * q.X + q.Y * q.Y + q.Z * q.Z + q.W * q.W;
}

func (q *Quaternion) Length() float64 {
	return math.Sqrt( q.X * q.X + q.Y * q.Y + q.Z * q.Z + q.W * q.W )
}

func (q *Quaternion) Normalize() (*Quaternion) {
	l := q.Length()

	if l == 0 {

		q.X = 0
		q.Y = 0
		q.Z = 0
		q.W = 1

	} else {

		l = 1 / l

		q.X = q.X * l
		q.Y = q.Y * l
		q.Z = q.Z * l
		q.W = q.W * l

	}

	q.onChangeCallback()

	return q
}

func (q *Quaternion) Multiply(quat *Quaternion) (*Quaternion) {
	return q.MultiplyQuaternions( q, quat )
}

func (q *Quaternion) MultiplyQuaternions( a, b *Quaternion) (*Quaternion) {

	// from http://www.euclideanspace.com/maths/algebra/realNormedAlgebra/quaternions/code/index.htm

	qax := a.X
	qay := a.Y
	qaz := a.Z
	qaw := a.W
	qbx := b.X
	qby := b.Y
	qbz := b.Z
	qbw := b.W

	q.X = qax * qbw + qaw * qbx + qay * qbz - qaz * qby
	q.Y = qay * qbw + qaw * qby + qaz * qbx - qax * qbz
	q.Z = qaz * qbw + qaw * qbz + qax * qby - qay * qbx
	q.W = qaw * qbw - qax * qbx - qay * qby - qaz * qbz

	q.onChangeCallback();

	return q
}

func (q *Quaternion) Slerp(qb *Quaternion, t int) (*Quaternion) {

	if t == 0 {
		return q
	}
	if t == 1 {
		return q.Copy( qb )
	}

	x := q.X
	y := q.Y
	z := q.Z
	w := q.W

	// http://www.euclideanspace.com/maths/algebra/realNormedAlgebra/quaternions/slerp/

	cosHalfTheta := w * qb.W + x * qb.X + y * qb.Y + z * qb.Z

	if cosHalfTheta < 0 {

		q.W = - qb.W
		q.X = - qb.X
		q.Y = - qb.Y
		q.Z = - qb.Z

		cosHalfTheta = - cosHalfTheta

	} else {

		q.Copy( qb )

	}

	if cosHalfTheta >= 1.0 {

		q.W = w
		q.X = x
		q.Y = y
		q.Z = z

		return q

	}

	halfTheta := math.Acos( cosHalfTheta )
	sinHalfTheta := math.Sqrt( 1.0 - cosHalfTheta * cosHalfTheta )

	if math.Abs( sinHalfTheta ) < 0.001 {

		q.W = 0.5 * ( w + q.W )
		q.X = 0.5 * ( x + q.X )
		q.Y = 0.5 * ( y + q.Y )
		q.Z = 0.5 * ( z + q.Z )

		return q

	}

	ratioA := math.Sin( ( 1 - t ) * halfTheta ) / sinHalfTheta
	ratioB := math.Sin( t * halfTheta ) / sinHalfTheta

	q.W = ( w * ratioA + q.W * ratioB )
	q.X = ( x * ratioA + q.X * ratioB )
	q.Y = ( y * ratioA + q.Y * ratioB )
	q.Z = ( z * ratioA + q.Z * ratioB )

	q.onChangeCallback()

	return q
}

func (q *Quaternion) Equals(quaternion *Quaternion) bool {
	return ( quaternion.X == q.X ) && ( quaternion.Y == q.Y ) && ( quaternion.Z == q.Z ) && ( quaternion.W == q.W )
}

func (q *Quaternion) FromArray(array []float64, offset int) (*Quaternion) {
	q.X = array[ offset ]
	q.Y = array[ offset + 1 ]
	q.Z = array[ offset + 2 ]
	q.W = array[ offset + 3 ]

	q.onChangeCallback()

	return q
}

func (q *Quaternion) ToArray(array []float64, offset int) ([]float64) {
	array[ offset ] = q.X
	array[ offset + 1 ] = q.Y
	array[ offset + 2 ] = q.Z
	array[ offset + 3 ] = q.W

	return array
}

func (q *Quaternion) OnChange(callback func()) (*Quaternion) {
	q.onChangeCallback = callback
	return q
}

func SlerpQuaternions( qa, qb, qm *Quaternion, t int) (*Quaternion) {
	return qm.Copy( qa ).Slerp( qb, t )
}
