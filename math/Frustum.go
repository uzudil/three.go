package math
import "github.com/uzudil/three.go/core"

type Frustum struct {
	Planes []*Plane

	IntersectsObject func(*core.Object3D) bool
	IntersectsBox func(*Box3) bool
}

func NewDefaultFrustum() (*Frustum) {
	return NewFrustum(
		NewDefaultPlane(),
		NewDefaultPlane(),
		NewDefaultPlane(),
		NewDefaultPlane(),
		NewDefaultPlane(),
		NewDefaultPlane(),
	)
}

func NewFrustum(p0, p1, p2, p3, p4, p5 *Plane) (*Frustum) {
	f := &Frustum{
		Planes: []*Plane{ p0, p1, p2, p3, p4, p5 },
	}

	f.IntersectsObject = f.buildIntersectsObject()
	f.IntersectsBox = f.buildIntersectsBox()

	return f
};

func (f *Frustum) Set( p0, p1, p2, p3, p4, p5 *Plane) (*Frustum) {
	var planes = f.Planes;

	planes[ 0 ].Copy( p0 )
	planes[ 1 ].Copy( p1 )
	planes[ 2 ].Copy( p2 )
	planes[ 3 ].Copy( p3 )
	planes[ 4 ].Copy( p4 )
	planes[ 5 ].Copy( p5 )

	return f
}

func (f *Frustum) Clone() (*Frustum) {
	return NewDefaultFrustum().Copy( f )
}

func (f *Frustum) Copy(frustum *Frustum) (*Frustum) {
	planes := f.Planes;
	for i := 0; i < 6; i++ {
		planes[ i ].Copy( frustum.Planes[ i ] );
	}
	return f
}

func (f *Frustum) SetFromMatrix(m *Matrix4) (*Frustum) {
	planes := f.Planes
	me := m.Elements
	me0 := me[ 0 ]; me1 := me[ 1 ]; me2 := me[ 2 ]; me3 := me[ 3 ]
	me4 := me[ 4 ]; me5 := me[ 5 ]; me6 := me[ 6 ]; me7 := me[ 7 ]
	me8 := me[ 8 ]; me9 := me[ 9 ]; me10 := me[ 10 ]; me11 := me[ 11 ]
	me12 := me[ 12 ]; me13 := me[ 13 ]; me14 := me[ 14 ]; me15 := me[ 15 ]

	planes[ 0 ].SetComponents( me3 - me0, me7 - me4, me11 - me8, me15 - me12 ).Normalize()
	planes[ 1 ].SetComponents( me3 + me0, me7 + me4, me11 + me8, me15 + me12 ).Normalize()
	planes[ 2 ].SetComponents( me3 + me1, me7 + me5, me11 + me9, me15 + me13 ).Normalize()
	planes[ 3 ].SetComponents( me3 - me1, me7 - me5, me11 - me9, me15 - me13 ).Normalize()
	planes[ 4 ].SetComponents( me3 - me2, me7 - me6, me11 - me10, me15 - me14 ).Normalize()
	planes[ 5 ].SetComponents( me3 + me2, me7 + me6, me11 + me10, me15 + me14 ).Normalize()

	return f
}

func (f *Frustum) buildIntersectsObject() (func(*core.Object3D) bool) {

	var sphere = NewDefaultSphere()

	return func(object *core.Object3D) bool {
		var geometry = object.Geometry
		if geometry.BoundingSphere == nil {
			geometry.ComputeBoundingSphere()
		}

		sphere.Copy( geometry.BoundingSphere )
		sphere.ApplyMatrix4( object.MatrixWorld )

		return f.IntersectsSphere( sphere )
	}
}

func (f *Frustum) IntersectsSphere(sphere *Sphere) bool {

	planes := f.Planes
	center := sphere.Center
	negRadius := - sphere.Radius

	for i := 0; i < 6; i++ {

		distance := planes[ i ].DistanceToPoint( center )

		if distance < negRadius {
			return false;
		}
	}
	return true;
}

func (f *Frustum) buildIntersectsBox() (func(*Box3) bool) {

	p1 := NewEmptyVector3()
	p2 := NewEmptyVector3()

	return func(box *Box3) bool {
		planes := f.Planes

		for i := 0; i < 6 ; i ++ {

			var plane = planes[ i ];

			if plane.Normal.X > 0 {
				p1.X = box.Min.X
			} else {
				p1.X = box.Max.X
			}
			if plane.Normal.X > 0 {
				p2.X = box.Max.X
			} else {
				p2.X = box.Min.X
			}
			if plane.Normal.Y > 0 {
				p1.Y = box.Min.Y
			} else {
				p1.Y = box.Max.Y
			}
			if plane.Normal.X > 0 {
				p2.Y = box.Max.Y
			} else {
				p2.Y = box.Min.Y
			}
			if plane.Normal.Z > 0 {
				p1.Z = box.Min.Z
			} else {
				p1.Z = box.Max.Z
			}
			if plane.Normal.Z > 0 {
				p2.Z = box.Max.Z
			} else {
				p2.Z = box.Min.Z
			}

			d1 := plane.DistanceToPoint( p1 )
			d2 := plane.DistanceToPoint( p2 )

			// if both outside plane, no intersection

			if d1 < 0 && d2 < 0 {
				return false
			}
		}
		return true
	}
}


func (f *Frustum) ContainsPoint(point *Vector3) bool {
	planes := f.Planes

	for i := 0; i < 6; i++ {
		if planes[ i ].DistanceToPoint( point ) < 0 {
			return false;
		}
	}

	return true
}
