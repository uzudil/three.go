package math
import (
	"math"
	"github.com/uzudil/three.go/core"
)

type Box3 struct {
	Min, Max *Vector3
	SetFromCenterAndSize func(*Vector3, *Vector3) (*Box3)
	SetFromObject func(*core.Object3D) (*Box3)
	DistanceToPoint (func(point *Vector3) float64)
	ApplyMatrix4 func(*Matrix4) (*Box3)
	GetBoundingSphere func(*Sphere) (*Sphere)
}

func NewDefaultBox3() (*Box3) {
	return NewBox3(
		NewVector3(math.Inf(1), math.Inf(1), math.Inf(1)), 
		NewVector3(math.Inf(-1), math.Inf(-1), math.Inf(-1)),
	)
}

func NewBox3(min, max *Vector3) (*Box3) {
	box := &Box3(min, max)
	box.SetFromCenterAndSize = box.buildSetFromCenterAndSize()
	box.SetFromObject = box.buildSetFromObject()
	box.DistanceToPoint = box.buildDistanceToPoint()
	box.ApplyMatrix4 = box.buildApplyMatrix4()
	box.GetBoundingSphere = box.buildGetBoundingSphere()
	return box
}

func (b *Box3) Set(min, max *Vector3) (*Box3) {
	b.Min.Copy( min )
	b.Max.Copy( max )

	return b
}

func (b *Box3) SetFromPoints(points []*Vector3) (*Box3) {

	b.MakeEmpty()

	il := len(points)
	for i := 0; i < il; i ++ {
		b.ExpandByPoint( points[ i ] )
	}

	return b
}

func (b *Box3) buildSetFromCenterAndSize() (func(*Vector3, *Vector3) (*Box3)) {
	var v1 = NewEmptyVector3()

	return func(center, size *Vector3) (*Box3) {
		var halfSize = v1.Copy( size ).MultiplyScalar( 0.5 )
		b.Min.Copy( center ).Sub( halfSize )
		b.Max.Copy( center ).Add( halfSize )
		return b
	}
}

func (b *Box3) buildSetFromObject() (*Box3) {
	// Computes the world-axis-aligned bounding box of an object (including its children),
	// accounting for both the object's, and children's, world transforms
	v1 := NewEmptyVector3()

	return func(object *core.Object3D) (*Box3) {
		var scope = b
		object.UpdateMatrixWorld( true )
		b.MakeEmpty()
		object.Traverse( func(node *core.Object3D) {
			var geometry = node.Geometry
			if geometry != nil {
				vertices := geometry.Vertices
				for _, v := range vertices {
					v1.Copy(v)
					v1.ApplyMatrix4(node.MatrixWorld)
					scope.ExpandByPoint(v1)
				}
			}
		} )
		return b
	}
}

func (b *Box3) Clone() (*Box3) {
	return NewDefaultBox3().Copy(b)
}

func (b *Box3) Copy(box *Box3) (*Box3) {
	b.Min.Copy(box.Min)
	b.Max.Copy(box.Max)

	return b
}

func (b *Box3) MakeEmpty() (*Box3) {
	b.Min.Set(math.Inf(1), math.Inf(1), math.Inf(1))
	b.Max.Set(-math.Inf(1), -math.Inf(1), -math.Inf(1))

	return b
}

func (b *Box3) Empty() bool {
	// this is a more robust check for empty than ( volume <= 0 ) because volume can get positive with two negative axes
	return ( b.Max.X < b.Min.X ) || ( b.Max.Y < b.Min.Y ) || ( b.Max.Z < b.Min.Z )
}

func (b *Box3) Center(optionalTarget *Vector3) (*Vector3) {
	result := optionalTarget
	if result == nil {
		result = NewEmptyVector3()
	}
	return result.AddVectors( b.Min, b.Max ).MultiplyScalar( 0.5 )
}

func (b *Box3) Size(optionalTarget *Vector3) (*Vector3) {
	result := optionalTarget
	if result == nil {
		result = NewEmptyVector3()
	}
	return result.SubVectors(b.Max, b.Min)
}

func (b *Box3) ExpandByPoint(point *Vector3) (*Box3) {
	b.Min.Min( point )
	b.Max.Max( point )
	return b
}

func (b *Box3) ExpandByVector(vector *Vector3) (*Box3) {
	b.Min.Sub( vector )
	b.Max.Add( vector )
	return b
}

func (b *Box3) ExpandByScalar(scalar float64) (*Box3) {
	b.Min.AddScalar( - scalar )
	b.Max.AddScalar( scalar )
	return b
}

func (b *Box3) ContainsPoint(point *Vector3) bool {
	if ( point.X < b.Min.X || point.X > b.Max.X ||
		 point.Y < b.Min.Y || point.Y > b.Max.Y ||
		 point.Z < b.Min.Z || point.Z > b.Max.Z ) {
		return false
	}
	return true
}

func (b *Box3) ContainsBox(box *Box3) bool {
	if ( ( b.Min.X <= box.Min.X ) && ( box.Max.X <= b.Max.X ) &&
		 ( b.Min.Y <= box.Min.Y ) && ( box.Max.Y <= b.Max.Y ) &&
		 ( b.Min.Z <= box.Min.Z ) && ( box.Max.Z <= b.Max.Z ) ) {
		return true
	}
	return false
}

func (b *Box3) GetParameter( point, optionalTarget *Vector3) (*Vector3) {

	// This can potentially have a divide by zero if the box
	// has a size dimension of 0.

	result := optionalTarget
	if result == nil {
		result = NewEmptyVector3()
	}

	return result.Set(
		( point.X - b.Min.X ) / ( b.Max.X - b.Min.X ),
		( point.Y - b.Min.Y ) / ( b.Max.Y - b.Min.Y ),
		( point.Z - b.Min.Z ) / ( b.Max.Z - b.Min.Z ),
	)
}

func (b *Box3) IsIntersectionBox(box *Box3) bool {
	// using 6 splitting planes to rule out intersections.
	if ( box.Max.X < b.Min.X || box.Min.X > b.Max.X ||
		 box.Max.Y < b.Min.Y || box.Min.Y > b.Max.Y ||
		 box.Max.Z < b.Min.Z || box.Min.Z > b.Max.Z ) {
		return false
	}
	return true
}

func (b *Box3) ClampPoint(point, optionalTarget *Vector3) (*Vector3) {
	result := optionalTarget
	if result == nil {
		result = NewEmptyVector3()
	}
	return result.Copy( point ).Clamp( b.Min, b.Max )
}

func (b *Box3) buildDistanceToPoint() (func(point *Vector3) float64) {
	var v1 = NewEmptyVector3()
	return func(point *Vector3) float64 {
		var clampedPoint = v1.Copy( point ).Clamp( b.Min, b.Max )
		return clampedPoint.Sub( point ).Length();
	};
}

func (b *Box3) buildGetBoundingSphere() (func(*Sphere) (*Sphere)) {
	var v1 = NewEmptyVector3()
	return func(optionalTarget *Sphere) (*Sphere) {
		result := optionalTarget
		if result == nil {
			result = NewDefaultSphere()
		}
		result.Center = b.Center(nil)
		result.Radius = b.Size(v1).Length() * 0.5
		return result;
	};
}

func (b *Box3) Intersect(box *Box3) (*Box3) {
	b.Min.Max( box.Min )
	b.Max.Min( box.Max )
	return b
}

func (b *Box3) Union(box *Box3) (*Box3) {
	b.Min.Min( box.Min )
	b.Max.Max( box.Max )
	return b
}

func (b *Box3) buildApplyMatrix4() (func(*Matrix4) (*Box3)) {

	var points = []*Vector3{
		NewEmptyVector3(),
		NewEmptyVector3(),
		NewEmptyVector3(),
		NewEmptyVector3(),
		NewEmptyVector3(),
		NewEmptyVector3(),
		NewEmptyVector3(),
		NewEmptyVector3(),
	}

	return func(matrix *Matrix4) (*Box3) {

		// NOTE: I am using a binary pattern to specify all 2^3 combinations below
		points[ 0 ].Set( b.Min.X, b.Min.Y, b.Min.Z ).ApplyMatrix4( matrix ) // 000
		points[ 1 ].Set( b.Min.X, b.Min.Y, b.Max.Z ).ApplyMatrix4( matrix ) // 001
		points[ 2 ].Set( b.Min.X, b.Max.Y, b.Min.Z ).ApplyMatrix4( matrix ) // 010
		points[ 3 ].Set( b.Min.X, b.Max.Y, b.Max.Z ).ApplyMatrix4( matrix ) // 011
		points[ 4 ].Set( b.Max.X, b.Min.Y, b.Min.Z ).ApplyMatrix4( matrix ) // 100
		points[ 5 ].Set( b.Max.X, b.Min.Y, b.Max.Z ).ApplyMatrix4( matrix ) // 101
		points[ 6 ].Set( b.Max.X, b.Max.Y, b.Min.Z ).ApplyMatrix4( matrix ) // 110
		points[ 7 ].Set( b.Max.X, b.Max.Y, b.Max.Z ).ApplyMatrix4( matrix )  // 111

		b.MakeEmpty()
		b.SetFromPoints( points )

		return b
	}
}

func (b *Box3) Translate( offset *Vector3) (*Vector3) {
	b.Min.Add( offset )
	b.Max.Add( offset )
	return b
}

func (b *Box3) Equals(box *Box3) bool {
	return box.Min.Equals( b.Min ) && box.Max.Equals( b.Max )
}
