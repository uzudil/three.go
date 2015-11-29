package math

type Line3 struct {
	Start, End *Vector3

	ClosestPointToPointParameter func(*Vector3, bool) float64
}

func NewDefaultLine3() (*Line3) {
	return NewLine3(NewEmptyVector3(), NewEmptyVector3())
}

func NewLine3( start, end *Vector3) (*Line3) {
	l := &Line3{
		Start: start,
		End: end,
	}

	l.ClosestPointToPointParameter = l.buildClosestPointToPointParameter()

	return l
}

func (l *Line3) Set( start, end *Vector3) (*Line3) {
	l.Start.Copy( start );
	l.End.Copy( end );

	return l
}

func (l *Line3) Clone() (*Line3) {
	return NewDefaultLine3().Copy(l)
}

func (l *Line3) Copy(line *Line3) (*Line3) {
	l.Start.Copy( line.Start )
	l.End.Copy( line.End )

	return l
}

func (l *Line3) Center(optionalTarget *Vector3) (*Vector3) {
	result := optionalTarget
	if result == nil {
		result = NewEmptyVector3()
	}
	return result.AddVectors( l.Start, l.End ).MultiplyScalar( 0.5 )
}

func (l *Line3) Delta(optionalTarget *Vector3) (*Vector3) {
	result := optionalTarget
	if result == nil {
		result = NewEmptyVector3()
	}
	return result.SubVectors( l.End, l.Start )
}

func (l *Line3) DistanceSq() float64 {
	return l.Start.DistanceToSquared( l.End )
}

func (l *Line3) Distance() float64 {
	return l.Start.DistanceTo( l.End )
}

func (l *Line3) At(t float64, optionalTarget *Vector3) (*Vector3) {
	result := optionalTarget
	if result == nil {
		result = NewEmptyVector3()
	}
	return l.Delta( result ).MultiplyScalar( t ).Add( l.Start )
}

func (l *Line3) buildClosestPointToPointParameter() (func(*Vector3, bool) float64) {

	startP := NewEmptyVector3()
	startEnd := NewEmptyVector3()

	return func(point *Vector3, clampToLine bool) {

		startP.SubVectors( point, l.Start )
		startEnd.SubVectors( l.End, l.Start )

		startEnd2 := startEnd.Dot( startEnd )
		startEnd_startP := startEnd.Dot( startP )

		t := startEnd_startP / startEnd2
		if clampToLine {
			t = Clamp( t, 0, 1 )
		}

		return t;
	}
}

func (l *Line3) ClosestPointToPoint(point *Vector3, clampToLine bool, optionalTarget *Vector3) {
	var t = l.ClosestPointToPointParameter(point, clampToLine)

	result := optionalTarget
	if result == nil {
		result = NewEmptyVector3()
	}

	return l.Delta( result ).MultiplyScalar( t ).Add( l.Start )
}

func (l *Line3) ApplyMatrix4(matrix *Matrix4) (*Line3) {
	l.Start.ApplyMatrix4( matrix )
	l.End.ApplyMatrix4( matrix )

	return l
}

func (l *Line3) Equals(line *Line3) bool {
	return line.Start.Equals( l.Start ) && line.End.Equals( l.End )
}
