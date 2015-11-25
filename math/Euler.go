package math
import (
	"math"
	"fmt"
)

type Euler struct {
	X, Y, Z float64
	Order string
	onChangeCallback func()
}

var RotationOrders [6]string = [...]string{ "XYZ", "YZX", "ZXY", "XZY", "YXZ", "ZYX" }

var DefaultOrder string = "XYZ"

func NewEmptyEuler() (*Euler) {
	return NewEuler(0, 0, 0, DefaultOrder)
}

func NewEuler(x, y, z float64, order string) (*Euler) {
	return &Euler{x, y, z, order, func() {}}
}

func (e *Euler) SetX(value float64) {
	e.X = value
	e.onChangeCallback()
}

func (e *Euler) SetY(value float64) {
	e.Y = value
	e.onChangeCallback()
}

func (e *Euler) SetZ(value float64) {
	e.Z = value
	e.onChangeCallback()
}

func (e *Euler) SetOrder(order string) {
	e.Order = order
	e.onChangeCallback()
}

func (e *Euler) Set(x, y, z float64, order string) (*Euler) {
	e.X = x
	e.Y = y
	e.Z = z
	if order != nil {
		e.Order = order
	}

	e.onChangeCallback()

	return e

}

func (e *Euler) Clone() (*Euler) {
	return NewEuler(e.X, e.Y, e.Z, e.Order)
}

func (e *Euler) Copy(euler *Euler) (*Euler) {
	e.X = euler.X
	e.Y = euler.Y
	e.Z = euler.Z
	e.Order = euler.Order

	e.onChangeCallback()

	return e
}

func (e *Euler) SetFromRotationMatrix(m *Matrix4, order string, update bool) (*Euler) {

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

	if order == nil {
		order = e.Order
	}

	if order == "XYZ" {

		e.Y = math.Asin( Clamp( m13, - 1, 1 ) )

		if math.Abs( m13 ) < 0.99999 {

			e.X = math.Atan2( - m23, m33 )
			e.Z = math.Atan2( - m12, m11 )

		} else {

			e.X = math.Atan2( m32, m22 )
			e.Z = 0

		}

	} else if order == "YXZ" {

		e.X = math.Asin( - Clamp( m23, - 1, 1 ) )

		if math.Abs( m23 ) < 0.99999 {

			e.Y = math.Atan2( m13, m33 )
			e.Z = math.Atan2( m21, m22 )

		} else {

			e.Y = math.Atan2( - m31, m11 )
			e.Z = 0

		}

	} else if order == "ZXY" {

		e.X = math.Asin( Clamp( m32, - 1, 1 ) )

		if ( math.Abs( m32 ) < 0.99999 ) {

			e.Y = math.Atan2( - m31, m33 )
			e.Z = math.Atan2( - m12, m22 )

		} else {

			e.Y = 0
			e.Z = math.Atan2( m21, m11 )

		}

	} else if order == "ZYX" {

		e.Y = math.Asin( - Clamp( m31, - 1, 1 ) )

		if ( math.Abs( m31 ) < 0.99999 ) {

			e.X = math.Atan2( m32, m33 )
			e.Z = math.Atan2( m21, m11 )

		} else {

			e.X = 0
			e.Z = math.Atan2( - m12, m22 )

		}

	} else if order == "YZX" {

		e.Z = math.Asin( Clamp( m21, - 1, 1 ) )

		if ( math.Abs( m21 ) < 0.99999 ) {

			e.X = math.Atan2( - m23, m22 )
			e.Y = math.Atan2( - m31, m11 )

		} else {

			e.X = 0
			e.Y = math.Atan2( m13, m33 )

		}

	} else if order == "XZY" {

		e.Z = math.Asin( - Clamp( m12, - 1, 1 ) )

		if ( math.Abs( m12 ) < 0.99999 ) {

			e.X = math.Atan2( m32, m22 )
			e.Y = math.Atan2( m13, m11 )

		} else {

			e.X = math.Atan2( - m23, m33 )
			e.Y = 0

		}

	} else {
		panic(fmt.Sprint("THREE.Euler: .setFromRotationMatrix() given unsupported order: %s", order ))
	}

	e.Order = order

	if update != false {
		e.onChangeCallback()
	}

	return e
}

func (e *Euler) SetFromQuaternion(q Quaternion, order string, update bool) (*Euler) {

	var matrix Matrix4 = nil

	if matrix == nil {
		matrix = NewMatrix4()
	}
	matrix.MakeRotationFromQuaternion(q)
	e.SetFromRotationMatrix(matrix, order, update)

	return e
}

func (e *Euler) SetFromVector3(v Vector3, order string) (*Euler) {
	if order == nil {
		order = e.Order
	}
	return e.Set( v.X, v.Y, v.Z, order)
}

func (e *Euler) Reorder() (func(string) (*Euler)) {

	// WARNING: this discards revolution information -bhouston

	var q = NewEmptyQuaternion()

	return func(newOrder string) {
		q.SetFromEuler(e, false)
		e.SetFromQuaternion(q, newOrder, false)
	}
}

func (e *Euler) Equals(euler *Euler) bool {
	return ( euler.X == e.X ) && ( euler.Y == e.Y ) && ( euler.Z == e.Z ) && ( euler.Order == e.Order )
}

/*
func (e *Euler) FromArray(array [](...)) {

	e.X = array[ 0 ];
	e.Y = array[ 1 ];
	e.Z = array[ 2 ];
	if ( array[ 3 ] !== undefined ) this._order = array[ 3 ];

this.onChangeCallback();

return this;

},

toArray: function ( array, offset ) {

	if ( array == undefined ) array = [];
	if ( offset == undefined ) offset = 0;

	array[ offset ] = e.X;
	array[ offset + 1 ] = e.Y;
	array[ offset + 2 ] = e.Z;
	array[ offset + 3 ] = this._order;

	return array;

},
*/
func (e *Euler) ToVector3(optionalResult *Vector3) (*Vector3) {

	if optionalResult != nil {
		return optionalResult.Set( e.X, e.Y, e.Z )
	} else {
		return NewVector3( e.X, e.Y, e.Z )
	}
}

func (e *Euler) OnChange(callback func()) (*Euler) {
	e.onChangeCallback = callback
	return e
}


