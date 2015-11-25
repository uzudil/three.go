package math
import (
	"fmt"
	"math"
)

type Vector3 struct {
	X, Y, Z float64
}

func NewEmptyVector3() (*Vector3) {
	return NewVector3(0.0, 0.0, 0.0)
}

func NewVector3(x, y, z float64) (*Vector3) {
	return &Vector3{x, y, z}
}

func (v3 *Vector3) Set(x, y, z float64) (*Vector3) {
	v3.X = x
	v3.Y = y
	v3.Z = z
	return v3
}

func (v3 *Vector3) SetX(x float64) (*Vector3) {
	v3.X = x
	return v3
}

func (v3 *Vector3) SetY(y float64) (*Vector3) {
	v3.Y = y
	return v3
}

func (v3 *Vector3) SetZ(z float64) (*Vector3) {
	v3.Z = z
	return v3
}

func (v3 *Vector3) SetComponent(index int, value float64) {
	switch index {
		case 0: v3.X = value
		case 1: v3.Y = value
		case 2: v3.Z = value
		default: panic(fmt.Sprint("index is out of range: %d", index ))
	}
}

func (v3 *Vector3) GetComponent(index int) (float64) {
	switch index {
		case 0: return v3.X
		case 1: return v3.Y
		case 2: return v3.Z
		default: panic(fmt.Sprint("index is out of range: %d", index))
	}
}

func (v3 *Vector3) Clone() (*Vector3) {
	return NewVector3(v3.X, v3.Y, v3.Z)
}

func (v3 *Vector3) Copy(other *Vector3) (*Vector3) {
	v3.X = other.X
	v3.Y = other.Y
	v3.Z = other.Z
	return v3
}

func (v3 *Vector3) Add(other *Vector3) (*Vector3) {
	v3.X += other.X
	v3.Y += other.Y
	v3.Z += other.Z
	return v3
}

func (v3 *Vector3) AddScalar(s float64) (*Vector3) {
	v3.X += s
	v3.Y += s
	v3.Z += s
	return v3
}

func (v3 *Vector3) AddVectors(a, b Vector3) (*Vector3) {
	v3.X = a.X + b.X
	v3.Y = a.Y + b.Y
	v3.Z = a.Z + b.Z
	return v3
}

func (v3 *Vector3) AddScaledVector(v Vector3, s float64) (*Vector3) {
	v3.X += v.X * s
	v3.Y += v.Y * s
	v3.Z += v.Z * s
	return v3
}

func (v3 *Vector3) Sub(v Vector3) (*Vector3) {
	v3.X -= v.X
	v3.Y -= v.Y
	v3.Z -= v.Z
	return v3
}

func (v3 *Vector3) SubScalar(s float64) (*Vector3) {
	v3.X -= s
	v3.Y -= s
	v3.Z -= s
	return v3
}

func (v3 *Vector3) SubVectors(a, b Vector3) (*Vector3) {
	v3.X = a.X - b.X
	v3.Y = a.Y - b.Y
	v3.Z = a.Z - b.Z
	return v3
}

func (v3 *Vector3) Multiply(v Vector3) (*Vector3) {
	v3.X *= v.X
	v3.Y *= v.Y
	v3.Z *= v.Z
	return v3
}

func (v3 *Vector3) MultiplyScalar(scalar float64) (*Vector3) {
	v3.X *= scalar
	v3.Y *= scalar
	v3.Z *= scalar
	return v3
}

func (v3 *Vector3) MultiplyVectors(a, b Vector3) (*Vector3) {
	v3.X = a.X * b.X
	v3.Y = a.Y * b.Y
	v3.Z = a.Z * b.Z
	return v3
}

/*
	applyEuler: function () {

		var quaternion;

		return function applyEuler( euler ) {

			if ( euler instanceof THREE.Euler === false ) {

				console.error( 'THREE.Vector3: .applyEuler() now expects a Euler rotation rather than a Vector3 and order.' );

			}

			if ( quaternion === undefined ) quaternion = new THREE.Quaternion();

			this.applyQuaternion( quaternion.setFromEuler( euler ) );

			return this;

		};

	}(),

	applyAxisAngle: function () {

		var quaternion;

		return function applyAxisAngle( axis, angle ) {

			if ( quaternion === undefined ) quaternion = new THREE.Quaternion();

			this.applyQuaternion( quaternion.setFromAxisAngle( axis, angle ) );

			return this;

		};

	}(),
 */

func (v3 *Vector3) ApplyMatrix3(m *Matrix3) (*Vector3) {
	x := v3.X
	y := v3.Y
	z := v3.Z

	e := m.Elements

	v3.X = e[ 0 ] * x + e[ 3 ] * y + e[ 6 ] * z
	v3.Y = e[ 1 ] * x + e[ 4 ] * y + e[ 7 ] * z
	v3.Z = e[ 2 ] * x + e[ 5 ] * y + e[ 8 ] * z

	return v3

}

func (v3 *Vector3) ApplyMatrix4(m *Matrix4) (*Vector3) {
	x := v3.X
	y := v3.Y
	z := v3.Z

	e := m.Elements

	v3.X = e[ 0 ] * x + e[ 4 ] * y + e[ 8 ] * z + e[12]
	v3.Y = e[ 1 ] * x + e[ 5 ] * y + e[ 9 ] * z + e[13]
	v3.Z = e[ 2 ] * x + e[ 6 ] * y + e[ 10 ] * z + e[14]

	return v3
}

func (v3 *Vector3) ApplyProjection(m *Matrix4) (*Vector3) {
	x := v3.X
	y := v3.Y
	z := v3.Z

	e := m.Elements
	d := 1 / ( e[ 3 ] * x + e[ 7 ] * y + e[ 11 ] * z + e[ 15 ] ) // perspective divide

	v3.X = (e[ 0 ] * x + e[ 4 ] * y + e[ 8 ] * z + e[12]) * d
	v3.Y = (e[ 1 ] * x + e[ 5 ] * y + e[ 9 ] * z + e[13]) * d
	v3.Z = (e[ 2 ] * x + e[ 6 ] * y + e[ 10 ] * z + e[14]) * d

	return v3
}

func (v3 *Vector3) ApplyQuaternion(q *Quaternion) (*Vector3) {

	x := v3.X
	y := v3.Y
	z := v3.Z

	qx := q.X
	qy := q.Y
	qz := q.Z
	qw := q.W

	// calculate quat * vector

	ix :=  qw * x + qy * z - qz * y
	iy :=  qw * y + qz * x - qx * z
	iz :=  qw * z + qx * y - qy * x
	iw := - qx * x - qy * y - qz * z

	// calculate result * inverse quat

	v3.X = ix * qw + iw * - qx + iy * - qz - iz * - qy
	v3.Y = iy * qw + iw * - qy + iz * - qx - ix * - qz
	v3.Z = iz * qw + iw * - qz + ix * - qy - iy * - qx

	return v3
}

/*
func (v3 *Vector3) project() (*Vector3) {
	project: function () {

		var matrix;

		return function project( camera ) {

			if ( matrix === undefined ) matrix = new THREE.Matrix4();

			matrix.multiplyMatrices( camera.projectionMatrix, matrix.getInverse( camera.matrixWorld ) );
			return this.applyProjection( matrix );

		};

	}(),

	unproject: function () {

		var matrix;

		return function unproject( camera ) {

			if ( matrix === undefined ) matrix = new THREE.Matrix4();

			matrix.multiplyMatrices( camera.matrixWorld, matrix.getInverse( camera.projectionMatrix ) );
			return this.applyProjection( matrix );

		};

	}(),
*/

func (v3 *Vector3) TransformDirection(m *Matrix4) (*Vector3) {
	// input: THREE.Matrix4 affine matrix
	// vector interpreted as a direction
	x := v3.X
	y := v3.Y
	z := v3.Z

	e := m.Elements

	v3.X = e[ 0 ] * x + e[ 4 ] * y + e[ 8 ] * z
	v3.Y = e[ 1 ] * x + e[ 5 ] * y + e[ 9 ] * z
	v3.Z = e[ 2 ] * x + e[ 6 ] * y + e[ 10 ] * z

	v3.Normalize()

	return v3
}

func (v3 *Vector3) Divide(v *Vector3) (*Vector3) {
	v3.X /= v.X
	v3.Y /= v.Y
	v3.Z /= v.Z
	return v3
}

func (v3 *Vector3) DivideScalar(scalar float64) (*Vector3) {
	return v3.MultiplyScalar( 1 / scalar )
}

func (v3 *Vector3) Min(v *Vector3) (*Vector3) {
	v3.X = math.Min( v3.X, v.X )
	v3.Y = math.Min( v3.Y, v.Y )
	v3.Z = math.Min( v3.Z, v.Z )

	return v3
}

func (v3 *Vector3) Max(v *Vector3) (*Vector3) {
	v3.X = math.Max( v3.X, v.X )
	v3.Y = math.Max( v3.Y, v.Y )
	v3.Z = math.Max( v3.Z, v.Z )

	return v3
}

func (v3 *Vector3) Clamp(min, max *Vector3) (*Vector3) {
	v3.X = math.Max( min.X, math.Min(max.X, v3.X ))
	v3.Y = math.Max( min.Y, math.Min(max.Y, v3.Y ))
	v3.Z = math.Max( min.Z, math.Min(max.Z, v3.Z ))

	return v3
}
/*
	clampScalar: function () {

		var min, max;

		return function clampScalar( minVal, maxVal ) {

			if ( min === undefined ) {

				min = new THREE.Vector3();
				max = new THREE.Vector3();

			}

			min.set( minVal, minVal, minVal );
			max.set( maxVal, maxVal, maxVal );

			return this.clamp( min, max );

		};

	}(),
*/

func (v3 *Vector3) ClampLength(min, max float64) (*Vector3) {
	length := v3.Length()

	v3.MultiplyScalar( math.Max( min, math.Min( max, length ) ) / length )

	return v3

}

func (v3 *Vector3) Floor() (*Vector3) {
	v3.X = math.Floor( v3.X )
	v3.Y = math.Floor( v3.Y )
	v3.Z = math.Floor( v3.Z )

	return v3
}

func (v3 *Vector3) Ceil() (*Vector3) {
	v3.X = math.Ceil( v3.X )
	v3.Y = math.Ceil( v3.Y )
	v3.Z = math.Ceil( v3.Z )

	return v3
}

func (v3 *Vector3) Round() (*Vector3) {
	v3.X = Round( v3.X )
	v3.Y = Round( v3.Y )
	v3.Z = Round( v3.Z )

	return v3
}

func (v3 *Vector3) RoundToZero() (*Vector3) {
	if v3.X < 0 {
		v3.X = math.Ceil(v3.X)
	} else {
		v3.X = math.Floor(v3.X)
	}
	if v3.Y < 0 {
		v3.Y = math.Ceil(v3.Y)
	} else {
		v3.Y = math.Floor(v3.Y)
	}
	if v3.Z < 0 {
		v3.Z = math.Ceil(v3.Z)
	} else {
		v3.Z = math.Floor(v3.Z)
	}

	return v3
}

func (v3 *Vector3) Negate() (*Vector3) {
	v3.X = - v3.X
	v3.Y = - v3.Y
	v3.Z = - v3.Z
	return v3
}

func (v3 *Vector3) Dot(v *Vector3) (float64) {
	return v3.X * v.X + v3.Y * v.Y + v3.Z * v.Z
}

func (v3 *Vector3) LengthSq() (float64) {
	return v3.X * v3.X + v3.Y * v3.Y + v3.Z * v3.Z
}

func (v3 *Vector3) Length() (float64) {
	return math.Sqrt(v3.X * v3.X + v3.Y * v3.Y + v3.Z * v3.Z)
}

func (v3 *Vector3) LengthManhattan() (float64) {
	return math.Abs( v3.X ) + math.Abs( v3.Y ) + math.Abs( v3.Z )
}

func (v3 *Vector3) Normalize() (*Vector3) {
	return v3.DivideScalar( v3.Length() )
}

func (v3 *Vector3) SetLength(length float64) (*Vector3) {
	return v3.MultiplyScalar( length / v3.Length() )
}

func (v3 *Vector3) Lerp(v *Vector3, alpha float64) (*Vector3) {
	v3.X += ( v.X - v3.X ) * alpha
	v3.Y += ( v.Y - v3.Y ) * alpha
	v3.Z += ( v.Z - v3.Z ) * alpha
	return v3
}

func (v3 *Vector3) LerpVectors(v1, v2 *Vector3, alpha float64) (*Vector3) {
		v3.SubVectors( v2, v1 ).MultiplyScalar( alpha ).Add( v1 )
		return v3
}

func (v3 *Vector3) Cross(v *Vector3) (*Vector3) {
	x := v3.X
	y := v3.Y
	z := v3.Z

	v3.X = y * v.Z - z * v.Y
	v3.Y = z * v.X - x * v.Z
	v3.Z = x * v.Y - y * v.X

	return v3
}

func (v3 *Vector3) CrossVectors(a, b *Vector3) (*Vector3) {
	ax := a.X
	ay := a.Y
	az := a.Z
	bx := b.X
	by := b.Y
	bz := b.Z

	v3.X = ay * bz - az * by;
	v3.Y = az * bx - ax * bz;
	v3.Z = ax * by - ay * bx;

	return v3
}

/*
	projectOnVector: function () {

		var v1, dot;

		return function projectOnVector( vector ) {

			if ( v1 === undefined ) v1 = new THREE.Vector3();

			v1.copy( vector ).normalize();

			dot = this.dot( v1 );

			return this.copy( v1 ).multiplyScalar( dot );

		};

	}(),

	projectOnPlane: function () {

		var v1;

		return function projectOnPlane( planeNormal ) {

			if ( v1 === undefined ) v1 = new THREE.Vector3();

			v1.copy( this ).projectOnVector( planeNormal );

			return this.sub( v1 );

		}

	}(),

	reflect: function () {

		// reflect incident vector off plane orthogonal to normal
		// normal is assumed to have unit length

		var v1;

		return function reflect( normal ) {

			if ( v1 === undefined ) v1 = new THREE.Vector3();

			return this.sub( v1.copy( normal ).multiplyScalar( 2 * this.dot( normal ) ) );

		}

	}(),
*/

func (v3 *Vector3) AngleTo(v *Vector3) (float64) {
	theta := v3.Dot( v ) / ( v3.Length() * v.Length() )

	// clamp, to handle numerical problems
	return math.Acos( Clamp( theta, - 1, 1 ) )
}

func (v3 *Vector3) DistanceTo(v *Vector3) (float64) {
	return math.Sqrt( v3.DistanceToSquared( v ) )
}

func (v3 *Vector3) DistanceToSquared(v *Vector3) (float64) {
	dx := v3.X - v.X
	dy := v3.Y - v.Y
	dz := v3.Z - v.Z

	return dx * dx + dy * dy + dz * dz
}

func (v3 *Vector3) SetFromMatrixPosition(m *Matrix4) (*Vector3) {
	v3.X = m.Elements[ 12 ]
	v3.Y = m.Elements[ 13 ]
	v3.Z = m.Elements[ 14 ]

	return v3
}

func (v3 *Vector3) SetFromMatrixScale(m *Matrix4) (*Vector3) {
	var sx = v3.Set( m.Elements[ 0 ], m.Elements[ 1 ], m.Elements[ 2 ] ).Length()
	var sy = v3.Set( m.Elements[ 4 ], m.Elements[ 5 ], m.Elements[ 6 ] ).Length()
	var sz = v3.Set( m.Elements[ 8 ], m.Elements[ 9 ], m.Elements[ 10 ] ).Length()

	v3.X = sx
	v3.Y = sy
	v3.Z = sz

	return v3
}

func (v3 *Vector3) SetFromMatrixColumn(index int, m *Matrix4) (*Vector3) {
	offset := index * 4

	var me = m.Elements

	v3.X = me[ offset ]
	v3.Y = me[ offset + 1 ]
	v3.Z = me[ offset + 2 ]

	return v3

}

func (v3 *Vector3) Equals(v *Vector3) (bool) {
	return ( ( v.X == v3.X ) && ( v.Y == v3.Y ) && ( v.Z == v3.Z ) )
}

func (v3 *Vector3) FromArray(array []float64, offset int) (*Vector3) {
	v3.X = array[ offset ]
	v3.Y = array[ offset + 1 ]
	v3.Z = array[ offset + 2 ]

	return v3
}

func (v3 *Vector3) ToArray(array []float64, offset int) ([]float64) {
	array[ offset ] = v3.X
	array[ offset + 1 ] = v3.Y
	array[ offset + 2 ] = v3.Z

	return array

}

/*
func (v3 *Vector3) FromAttribute(attribute []float64, index, offset int) (*Vector3) {

	if ( offset === undefined ) offset = 0;

	index = index * attribute.itemSize + offset;

	this.x = attribute.array[ index ];
	this.y = attribute.array[ index + 1 ];
	this.z = attribute.array[ index + 2 ];

	return this;

}
*/

