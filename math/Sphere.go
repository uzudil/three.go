package math
import "math"

type Sphere struct {
	Center *Vector3
	Radius float64
	SetFromPoints func([]*Vector3, *Vector3) (*Sphere)
}

func NewDefaultSphere() (*Sphere) {
	return NewSphere(NewEmptyVector3(), 0)
}

func NewSphere(center *Vector3, radius float64) (*Sphere) {
	c := &Sphere{center, radius}
	c.SetFromPoints = c.buildSetFromPoints()
	return c
}

func (c *Sphere) Set(center *Vector3, radius float64) (*Sphere) {
	c.Center.Copy( center )
	c.Radius = radius
	return c
}

func (c *Sphere) buildSetFromPoints() (func([]*Vector3, *Vector3) (*Sphere)) {
	box := NewDefaultBox3()
	return func(points []*Vector3, optionalCenter *Vector3) (*Sphere) {
		center := c.Center
		if optionalCenter != nil {
			center.Copy(optionalCenter)
		} else {
			box.SetFromPoints(points).Center(center)
		}

		var maxRadiusSq = 0;
		for _, p := range points {
			maxRadiusSq = math.Max( maxRadiusSq, center.DistanceToSquared(p))
		}
		c.Radius = math.Sqrt( maxRadiusSq )
		return c
	}
}

func (c *Sphere) Clone() (*Sphere) {
	return NewDefaultSphere().Copy(c)
}

func (c *Sphere) Copy(sphere *Sphere) (*Sphere) {
	c.Center.Copy( sphere.Center )
	c.Radius = sphere.Radius
	return c
}

func (c *Sphere) Empty() bool {
	return c.Radius <= 0
}

func (c *Sphere) ContainsPoint(point *Vector3) bool {
	return point.DistanceToSquared(c.Center) <= (c.Radius * c.Radius)
}

func (c *Sphere) DistanceToPoint(point *Vector3) float64 {
	return point.DistanceTo(c.Center) - c.Radius
}

func (c *Sphere) IntersectsSphere(sphere *Sphere) bool {
	var radiusSum = c.Radius + sphere.Radius
	return sphere.Center.DistanceToSquared(c.Center) <= (radiusSum * radiusSum)
}

func (c *Sphere) ClampPoint(point, optionalTarget *Vector3) (*Vector3) {
	deltaLengthSq := c.Center.DistanceToSquared(point)
	result := optionalTarget
	if result == nil {
		result = NewEmptyVector3()
	}
	result.Copy(point)

	if deltaLengthSq > ( c.Radius * c.Radius ) {
		result.Sub( c.Center ).Normalize()
		result.MultiplyScalar( c.Radius ).Add( c.Center )
	}

	return result
}

func (c *Sphere) GetBoundingBox(optionalTarget *Box3) (*Box3) {
	box := optionalTarget
	if box == nil {
		box = NewDefaultBox3()
	}

	box.Set( c.Center, c.Center )
	box.ExpandByScalar( c.Radius )

	return box
}

func (c *Sphere) ApplyMatrix4(matrix *Matrix4) (*Sphere) {
	c.Center.ApplyMatrix4( matrix )
	c.Radius = c.Radius * matrix.GetMaxScaleOnAxis()

	return c
}

func (c *Sphere) Translate(offset *Vector3) (*Sphere) {
	c.Center.Add( offset )
	return c
}

func (c *Sphere) Equals(sphere *Sphere) bool {
	return sphere.Center.Equals(c.Center) && (sphere.Radius == c.Radius)
}
