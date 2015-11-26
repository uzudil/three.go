package math
import "fmt"

type Vector2 struct {
	X, Y float64
}

func NewEmptyVector2() (*Vector2) {
	return NewVector2(0.0, 0.0)
}

func NewVector2(x, y float64) (*Vector2) {
	return &Vector2{x, y}
}

func (v *Vector2) Width() float64 {
	return v.X
}

func (v *Vector2) SetWidth(value float64) {
	v.X = value
}

func (v *Vector2) Height() float64 {
	return v.Y
}

func (v *Vector2) SetHeight(value float64) {
	v.Y = value
}

func (v *Vector2) Set(x, y float64) (*Vector2) {
	v.X = x
	v.Y = y

	return v
}

func (v *Vector2) SetX(x float64) (*Vector2) {
	v.X = x
	return v
}

func (v *Vector2) SetY(y float64) (*Vector2) {
	v.Y = y
	return v
}

func (v *Vector2) SetComponent(index int, value float64) {
	switch ( index ) {
		case 0: v.X = value
		case 1: v.Y = value
	}
}

func (v *Vector2) GetComponent(index int) float64 {
	switch ( index ) {
		case 0: return v.X;
		case 1: return v.Y;
		default: panic(fmt.Sprint("index is out of range: %d", index))
	}
}

func (v *Vector2) Clone() (*Vector2) {
	return NewVector2(v.X, v.Y)
}

func (v *Vector2) Copy(vector *Vector2) (*Vector2) {
	v.X = vector.X
	v.Y = vector.Y
	return v
}
/*
add: function ( v, w ) {

	if ( w !== undefined ) {

		console.warn( 'THREE.Vector2: .add() now only accepts one argument. Use .addVectors( a, b ) instead.' );
		return this.addVectors( v, w );

	}

	v.X += v.x;
	v.Y += v.y;

	return v

},

addScalar: function ( s ) {

	v.X += s;
	v.Y += s;

	return v

},

addVectors: function ( a, b ) {

	v.X = a.x + b.x;
	v.Y = a.y + b.y;

	return v

},

addScaledVector: function ( v, s ) {

	v.X += v.x * s;
	v.Y += v.y * s;

	return v

},

sub: function ( v, w ) {

	if ( w !== undefined ) {

		console.warn( 'THREE.Vector2: .sub() now only accepts one argument. Use .subVectors( a, b ) instead.' );
		return this.subVectors( v, w );

	}

	v.X -= v.x;
	v.Y -= v.y;

	return v

},

subScalar: function ( s ) {

	v.X -= s;
	v.Y -= s;

	return v

},

subVectors: function ( a, b ) {

	v.X = a.x - b.x;
	v.Y = a.y - b.y;

	return v

},

multiply: function ( v ) {

	v.X *= v.x;
	v.Y *= v.y;

	return v

},

multiplyScalar: function ( scalar ) {

	if ( isFinite( scalar ) ) {
		v.X *= scalar;
		v.Y *= scalar;
	} else {
		v.X = 0;
		v.Y = 0;
	}

	return v

},

divide: function ( v ) {

	v.X /= v.x;
	v.Y /= v.y;

	return v

},

divideScalar: function ( scalar ) {

	return this.multiplyScalar( 1 / scalar );

},

min: function ( v ) {

	v.X = Math.min( v.X, v.x );
	v.Y = Math.min( v.Y, v.y );

	return v

},

max: function ( v ) {

	v.X = Math.max( v.X, v.x );
	v.Y = Math.max( v.Y, v.y );

	return v

},

clamp: function ( min, max ) {

	// This function assumes min < max, if this assumption isn't true it will not operate correctly

	v.X = Math.max( min.x, Math.min( max.x, v.X ) );
	v.Y = Math.max( min.y, Math.min( max.y, v.Y ) );

	return v

},

clampScalar: function () {

	var min, max;

	return function clampScalar( minVal, maxVal ) {

		if ( min === undefined ) {

			min = new THREE.Vector2();
			max = new THREE.Vector2();

		}

		min.set( minVal, minVal );
		max.set( maxVal, maxVal );

		return this.clamp( min, max );

	};

}(),

clampLength: function ( min, max ) {

	var length = this.length();

	this.multiplyScalar( Math.max( min, Math.min( max, length ) ) / length );

	return v

},

floor: function () {

	v.X = Math.floor( v.X );
	v.Y = Math.floor( v.Y );

	return v

},

ceil: function () {

	v.X = Math.ceil( v.X );
	v.Y = Math.ceil( v.Y );

	return v

},

round: function () {

	v.X = Math.round( v.X );
	v.Y = Math.round( v.Y );

	return v

},

roundToZero: function () {

	v.X = ( v.X < 0 ) ? Math.ceil( v.X ) : Math.floor( v.X );
	v.Y = ( v.Y < 0 ) ? Math.ceil( v.Y ) : Math.floor( v.Y );

	return v

},

negate: function () {

	v.X = - v.X;
	v.Y = - v.Y;

	return v

},

dot: function ( v ) {

	return v.X * v.x + v.Y * v.y;

},

lengthSq: function () {

	return v.X * v.X + v.Y * v.Y;

},

length: function () {

	return Math.sqrt( v.X * v.X + v.Y * v.Y );

},

lengthManhattan: function() {

	return Math.abs( v.X ) + Math.abs( v.Y );

},

normalize: function () {

	return this.divideScalar( this.length() );

},

distanceTo: function ( v ) {

	return Math.sqrt( this.distanceToSquared( v ) );

},

distanceToSquared: function ( v ) {

	var dx = v.X - v.x, dy = v.Y - v.y;
	return dx * dx + dy * dy;

},

setLength: function ( length ) {

	return this.multiplyScalar( length / this.length() );

},

lerp: function ( v, alpha ) {

	v.X += ( v.x - v.X ) * alpha;
	v.Y += ( v.y - v.Y ) * alpha;

	return v

},

lerpVectors: function ( v1, v2, alpha ) {

	this.subVectors( v2, v1 ).multiplyScalar( alpha ).add( v1 );

	return v

},

equals: function ( v ) {

	return ( ( v.x === v.X ) && ( v.y === v.Y ) );

},

fromArray: function ( array, offset ) {

	if ( offset === undefined ) offset = 0;

	v.X = array[ offset ];
	v.Y = array[ offset + 1 ];

	return v

},

toArray: function ( array, offset ) {

	if ( array === undefined ) array = [];
	if ( offset === undefined ) offset = 0;

	array[ offset ] = v.X;
	array[ offset + 1 ] = v.Y;

	return array;

},

fromAttribute: function ( attribute, index, offset ) {

	if ( offset === undefined ) offset = 0;

	index = index * attribute.itemSize + offset;

	v.X = attribute.array[ index ];
	v.Y = attribute.array[ index + 1 ];

	return v

},

rotateAround: function ( center, angle ) {

	var c = Math.cos( angle ), s = Math.sin( angle );

	var x = v.X - center.x;
	var y = v.Y - center.y;

	v.X = x * c - y * s + center.x;
	v.Y = x * s + y * c + center.y;

	return v

}
*/