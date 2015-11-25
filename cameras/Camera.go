package cameras
import (
	"github.com/uzudil/three.go/core"
	"github.com/uzudil/three.go/math"
)

type Camera struct {
	*core.Object3D
	Type string
	MatrixWorldInverse *math.Matrix4
	ProjectionMatrix *math.Matrix4
	LookAt func(*math.Vector3)
}

func NewCamera() (*Camera) {
	c := &Camera{
		core.NewObject3D(),
		Type: "Camera",
		MatrixWorldInverse: math.NewMatrix4(),
		ProjectionMatrix: math.NewMatrix4(),
	}
	c.LookAt = c.buildLookAt()
	return c
}

func (c *Camera) GetWorldDirection(optionalTarget *math.Vector3) (*math.Vector3) {

	var quaternion = math.NewEmptyQuaternion()

	var result = optionalTarget
	if result == nil {
		result = math.NewEmptyVector3()
	}

	c.GetWorldQuaternion(quaternion)

	return result.Set( 0, 0, - 1 ).ApplyQuaternion( quaternion )

}

func (c *Camera) buildLookAt() (func(*math.Vector3)) {
	// This routine does not support cameras with rotated and/or translated parent(s)
	m1 := math.NewMatrix4()
	return func(vector *math.Vector3) {
		m1.LookAt(c.Position, vector, c.Up)
		c.Quaternion.SetFromRotationMatrix(m1)
	}
}

func (c *Camera) Clone() (*Camera) {
	return NewCamera().Copy(c)
}

func (c *Camera) Copy(source *Camera) (*Camera) {
	c.Object3D.Copy(source.Object3D)

	c.MatrixWorldInverse.Copy(source.MatrixWorldInverse)
	c.ProjectionMatrix.Copy(source.ProjectionMatrix)

	return c
}
