package math

type Plane struct {
	Normal *Vector3
	Constant float64

	SetFromCoplanarPoints func(*Vector3, *Vector3, *Vector3) (*Plane)
	IntersectLine func(*Line3, *Vector3) (*Vector3)
	ApplyMatrix4 func(*Matrix4, *Matrix3) (*Plane)
}

func NewDefaultPlane() (*Plane) {
	return NewPlane(NewVector3(1.0, 0.0, 0.0), 0.0)
}

func NewPlane(normal *Vector3, constant float64) (*Plane) {
	p := &Plane{
		Normal: normal,
		Constant: constant,
	}

	p.SetFromCoplanarPoints = p.buildSetFromCoplanarPoints()
	p.IntersectLine = p.buildIntersectLine()
	p.ApplyMatrix4 = p.buildApplyMatrix4()

	return p
}

func (p *Plane) Set(normal *Vector3, constant float64) (*Plane) {
	p.Normal.Copy( normal )
	p.Constant = constant

	return p
}

func (p *Plane) SetComponents(x, y, z float64, w float64) (*Plane) {
	p.Normal.Set( x, y, z )
	p.Constant = w

	return p
}

func (p *Plane) SetFromNormalAndCoplanarPoint(normal *Vector3, point *Vector3) (*Plane) {
	p.Normal.Copy(normal)
	p.Constant = - point.Dot(p.Normal)    // must be p.Normal, not normal, as p.Normal is normalized

	return p
}

func (p *Plane) buildSetFromCoplanarPoints() (func(*Vector3, *Vector3, *Vector3) (*Plane)) {
	v1 := NewEmptyVector3()
	v2 := NewEmptyVector3()

	return func(a, b, c *Vector3) {
		var normal = v1.SubVectors( c, b ).Cross( v2.SubVectors( a, b ) ).Normalize()
		// Q: should an error be thrown if normal is zero (e.g. degenerate plane)?
		p.SetFromNormalAndCoplanarPoint( normal, a );
		return p
	}
}

func (p *Plane) Clone() (*Plane) {
	return NewDefaultPlane().Copy(p)
}

func (p *Plane) Copy(plane *Plane) (*Plane) {
	p.Normal.Copy( plane.Normal )
	p.Constant = plane.Constant

	return p
}

func (p *Plane) Normalize() (*Plane) {
	// Note: will lead to a divide by zero if the plane is invalid.
	inverseNormalLength := 1.0 / p.Normal.Length()
	p.Normal.MultiplyScalar( inverseNormalLength )
	p.Constant *= inverseNormalLength
	return p
}

func (p *Plane) Negate() (*Plane) {
	p.Constant *= - 1
	p.Normal.Negate()

	return p
}

func (p *Plane) DistanceToPoint(point *Vector3) float64 {
	return p.Normal.Dot( point ) + p.Constant
}

func (p *Plane) DistanceToSphere(sphere *Sphere) float64 {
	return p.DistanceToPoint(sphere.Center) - sphere.Radius
}

func (p *Plane) ProjectPoint(point, optionalTarget *Vector3) {
	return p.OrthoPoint(point, optionalTarget).Sub(point).Negate()
}

func (p *Plane) OrthoPoint(point, optionalTarget *Vector3) (*Vector3) {
	var perpendicularMagnitude = p.DistanceToPoint( point )

	result := optionalTarget
	if result == nil {
		result = NewEmptyVector3()
	}
	return result.Copy( p.Normal ).MultiplyScalar( perpendicularMagnitude )
}

func (p *Plane) IsIntersectionLinefunction(line *Line3) {

	// Note: this tests if a line intersects the plane, not whether it (or its end-points) are coplanar with it.

	var startSign = p.DistanceToPoint( line.Start )
	var endSign = p.DistanceToPoint( line.End )

	return ( startSign < 0 && endSign > 0 ) || ( endSign < 0 && startSign > 0 );
}

func (p *Plane) buildIntersectLine() (func(*Line3, *Vector3) (*Vector3)) {
	v1 := NewEmptyVector3()

	return func(line *Line3, optionalTarget *Vector3) (*Vector3) {
		result := optionalTarget
		if result == nil {
			result = NewEmptyVector3()
		}

		direction := line.Delta( v1 )
		denominator := p.Normal.Dot( direction )
		if denominator == 0 {
			// line is coplanar, return origin
			if p.DistanceToPoint(line.Start) == 0 {
				return result.Copy( line.Start )
			}

			// Unsure if this is the correct method to handle this case.
			return nil
		}

		t := - ( line.Start.Dot( p.Normal ) + p.Constant ) / denominator
		if t < 0 || t > 1 {
			return nil
		}

		return result.Copy( direction ).MultiplyScalar( t ).Add( line.Start )
	}
}

func (p *Plane) CoplanarPoint(optionalTarget *Vector3) (*Vector3) {
	result := optionalTarget
	if result == nil {
		result = NewEmptyVector3()
	}
	return result.Copy( p.Normal ).MultiplyScalar( - p.Constant )
}

func (p *Plane) buildApplyMatrix4() (func(*Matrix4, *Matrix3) (*Plane)) {

	v1 := NewEmptyVector3()
	v2 := NewEmptyVector3()
	m1 := NewMatrix3()

	return func(matrix *Matrix4, optionalNormalMatrix *Matrix3) (*Plane) {
		// compute new normal based on theory here:
		// http://www.songho.ca/opengl/gl_normaltransform.html
		normalMatrix := optionalNormalMatrix
		if normalMatrix == nil {
			normalMatrix = m1.GetNormalMatrix(matrix)
		}
		newNormal := v1.Copy( p.Normal ).ApplyMatrix3( normalMatrix )

		newCoplanarPoint := p.CoplanarPoint( v2 )
		newCoplanarPoint.ApplyMatrix4( matrix )

		p.SetFromNormalAndCoplanarPoint( newNormal, newCoplanarPoint )

		return p
	}
}

func (p *Plane) Translate(offset *Vector3) (*Plane) {
	p.Constant = p.Constant - offset.Dot( p.Normal )
	return p
}

func (p *Plane) Equals(plane *Plane) bool {
	return plane.Normal.Equals(p.Normal) && (plane.Constant == p.Constant)
}
