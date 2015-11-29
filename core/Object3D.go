package core
import (
	"github.com/uzudil/three.go/math"
	"fmt"
)

type Object3D struct {
	*EventDispatcher
	Id int
	Uuid string
	Name string
	Type string
	Parent Object3D
	Children []Object3D
	Up *math.Vector3
	Position *math.Vector3
	Rotation *math.Euler
	Quaternion *math.Quaternion
	Scale *math.Vector3
	RotationAutoUpdate bool
	Matrix math.Matrix4
	MatrixWorld math.Matrix4
	MatrixAutoUpdate bool
	MatrixWorldNeedsUpdate bool
	Visible bool
	CastShadow, ReceiveShadow bool
	FrustumCulled bool
	RenderOrder int
	UserData map[string]string
	ModelViewMatrix *math.Matrix4
	NormalMatrix *math.Matrix3
	Geometry *Geometry

	GetWorldQuaternion func(*math.Vector3) (*math.Quaternion)
	GetWorldRotation func(*math.Euler) (*math.Euler)
	GetWorldScale func(*math.Vector3) (*math.Vector3)
	GetWorldDirection func(*math.Vector3) (*math.Vector3)
	LookAt func(*math.Vector3)
}

var Object3DIdCount int = 0
var DefaultUp = math.NewVector3( 0.0, 1.0, 0.0 )
var DefaultMatrixAutoUpdate bool = true

func NewObject3D() (*Object3D) {
	Object3DIdCount++
	object3d := Object3D{
		NewEventDispatcher(),
		Id: Object3DIdCount,
		Uuid: math.GenerateUUID(),
		Name: "",
		Type: "Object3D",
		Parent: nil,
		Channels: NewChannels(),
		Children: make([]Object3D, 0),
		Up: DefaultUp.Clone(),
		Position: math.NewEmptyVector3(),
		Rotation: math.NewEmptyEuler(),
		Quaternion: math.NewEmptyQuaternion(),
		Scale: math.NewVector3(1.0, 1.0, 1.0),
		RotationAutoUpdate: true,
		Matrix: math.NewMatrix4(),
		MatrixWorld: math.NewMatrix4(),
		MatrixAutoUpdate: DefaultMatrixAutoUpdate,
		MatrixWorldNeedsUpdate: false,
		Visible: true,
		CastShadow: false,
		ReceiveShadow: false,
		FrustumCulled: true,
		RenderOrder: 0,
		UserData: make(map[string]string),
		ModelViewMatrix: math.NewMatrix4(),
		NormalMatrix: *math.NewMatrix3(),
		Geometry: nil,
	}

	onRotationChange := func() {
		object3d.Quaternion.SetFromEuler( object3d.Rotation, false )
	}

	onQuaternionChange := func() {
		object3d.Rotation.SetFromQuaternion( object3d.Quaternion, nil, false )
	}

	object3d.Rotation.OnChange( onRotationChange )
	object3d.Quaternion.OnChange( onQuaternionChange )

	object3d.GetWorldQuaternion = object3d.buildGetWorldQuaternion()
	object3d.GetWorldRotation = object3d.buildGetWorldRotation()
	object3d.GetWorldScale() = object3d.buildGetWorldScale()
	object3d.GetWorldDirection = object3d.buildGetWorldDirection()
	object3d.LookAt = object3d.buildLookAt()

	return &object3d
}

func (o *Object3D) Clone(recursive bool) (*Object3D) {
	return NewObject3D().Copy(o, recursive)
}

func (o *Object3D) Copy(source *Object3D, recursive bool) (*Object3D) {
	// if ( recursive === undefined ) recursive = true;

	o.Name = source.Name

	o.Up.Copy( source.Up )

	o.Position.Copy( source.Position )
	o.Quaternion.Copy( source.Quaternion )
	o.Scale.Copy( source.Scale )

	o.RotationAutoUpdate = source.RotationAutoUpdate

	o.Matrix.Copy( source.Matrix )
	o.MatrixWorld.Copy( source.MatrixWorld )

	o.MatrixAutoUpdate = source.MatrixAutoUpdate
	o.MatrixWorldNeedsUpdate = source.MatrixWorldNeedsUpdate

	o.Visible = source.Visible

	o.CastShadow = source.CastShadow
	o.ReceiveShadow = source.ReceiveShadow

	o.FrustumCulled = source.FrustumCulled
	o.RenderOrder = source.RenderOrder

	// o.UserData = JSON.parse( JSON.stringify( source.userData ) );

	if recursive {
		for i := 0; i < len(source.Children); i ++ {
			var child = source.Children[ i ];
			o.Add( child.Clone(true) );
		}
	}

	return o
}

func (o *Object3D) Add(object *Object3D) (*Object3D) {

	// what is this?
//	if ( arguments.length > 1 ) {
//		for ( var i = 0; i < arguments.length; i ++ ) {
//			this.add( arguments[ i ] );
//		}
//		return this;
//	}

	if ( object == o ) {
		fmt.Println("THREE.Object3D.add: object can't be added as a child of itself.", object)
		return o
	}

	if object.Parent != nil {
		object.Parent.Remove(object)
	}

	object.Parent = o
	object.DispatchEvent(NewEvent("added"))

	o.Children = append(o.Children, object )

	return o
}

func (o *Object3D) indexOfChild(object *Object3D) int {
	for index, child := range o.Children {
		if child == object {
			return index
		}
	}
	return -1
}

func (o *Object3D) Remove(object *Object3D) {

	// what is this?
//	if ( arguments.length > 1 ) {
//		for ( var i = 0; i < arguments.length; i ++ ) {
//			this.remove( arguments[ i ] );
//		}
//	}

	var index = o.indexOfChild( object )
	if index != - 1 {
		object.Parent = nil
		object.DispatchEvent(NewEvent("removed"))

		// go's version of remove
		o.Children = append(o.Children[:index], o.Children[index+1:]...)
	}
}

func (o *Object3D) UpdateMatrix() {
	o.Matrix.Compose( o.Position, o.Quaternion, o.Scale )
	o.MatrixWorldNeedsUpdate = true
}

func (o *Object3D) UpdateMatrixWorld(force bool) {

	if o.MatrixAutoUpdate {
		o.UpdateMatrix()
	}

	if o.MatrixWorldNeedsUpdate || force {

		if o.Parent == nil {
			o.MatrixWorld.Copy(o.Matrix)
		} else {
			o.MatrixWorld.MultiplyMatrices(o.Parent.MatrixWorld, o.Matrix)
		}
		o.MatrixWorldNeedsUpdate = false
		force = true
	}

	// update children
	l := len(o.Children)
	for i := 0; i < l; i ++ {
		o.Children[ i ].UpdateMatrixWorld( force )
	}
}

func (o *Object3D) GetWorldPosition(optionalTarget *math.Vector3) (*math.Vector3) {

	result := optionalTarget
	if result == nil {
		result = math.NewEmptyVector3()
	}

	o.UpdateMatrixWorld( true )

	return result.SetFromMatrixPosition( o.MatrixWorld )
}

func (o *Object3D) buildGetWorldQuaternion() (func(*math.Vector3) (*math.Quaternion)) {

	position := math.NewEmptyVector3()
	scale := math.NewEmptyVector3()

	return func(optionalTarget *math.Vector3) (*math.Quaternion) {

		var result = optionalTarget
		if result == nil {
			result = math.NewEmptyQuaternion()
		}

		o.UpdateMatrixWorld(true)

		o.MatrixWorld.Decompose(position, result, scale);

		return result;
	}
}

func (o *Object3D) buildGetWorldRotation() (func(*math.Euler) (*math.Euler)) {

	quaternion := math.NewEmptyQuaternion()

	return func(optionalTarget *math.Euler) {

		result := optionalTarget
		if result == nil {
			result = math.NewEmptyEuler()
		}

		o.GetWorldQuaternion( quaternion )

		return result.SetFromQuaternion( quaternion, o.Rotation.Order, false )
	}
}

func (o *Object3D) buildGetWorldScale() (func(*math.Vector3) (*math.Vector3)) {

	var position = math.NewEmptyVector3()
	var quaternion = math.NewEmptyQuaternion()

	return func(optionalTarget *math.Vector3) (*math.Vector3) {

		result := optionalTarget
		if result == nil {
			result = math.NewEmptyVector3()
		}

		o.UpdateMatrixWorld(true)

		o.MatrixWorld.Decompose(position, quaternion, result)

		return result
	}
}

func (o *Object3D) buildGetWorldDirection() (func(*math.Vector3) (*math.Vector3)) {

	var quaternion = math.NewEmptyQuaternion()

	return func(optionalTarget *math.Vector3) (*math.Vector3) {

		result := optionalTarget
		if result == nil {
			result = math.NewEmptyVector3()
		}

		o.GetWorldQuaternion( quaternion )

		return result.Set( 0, 0, 1 ).ApplyQuaternion( quaternion )
	}
}

func (o *Object3D) buildLookAt() (func(*math.Vector3)) {
	// This routine does not support objects with rotated and/or translated parent(s)
	var m1 = math.NewMatrix4()
	return func(vector *math.Vector3) {
		m1.LookAt(vector, o.Position, o.Up)
		o.Quaternion.SetFromRotationMatrix( m1 )
	}
}
