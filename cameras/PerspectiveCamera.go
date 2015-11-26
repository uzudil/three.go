package cameras
import (
	math3d "github.com/uzudil/three.go/math"
	"math"
)


type PerspectiveCamera struct {
	*Camera
	Zoom float64
	Fov, Aspect, Near, Far float64
	FullWidth, FullHeight, X, Y, Width, Height float64
}

func NewDefaultPerspectiveCamera() (*PerspectiveCamera) {
	return NewPerspectiveCamera(50.0, 1.0, 0.1, 2000.0)
}

func NewPerspectiveCamera( fov, aspect, near, far float64 ) (*PerspectiveCamera) {
	p := &PerspectiveCamera{
		NewCamera(),
		Zoom: 1.0,
	}

	p.Type = "PerspectiveCamera"

	p.Fov = fov
	p.Aspect = aspect
	p.Near = near
	p.Far = far

	p.UpdateProjectionMatrix()

	return p
}


/**
 * Uses Focal Length (in mm) to estimate and set FOV
 * 35mm (full-frame) camera is used if frame size is not specified;
 * Formula based on http://www.bobatkins.com/photography/technical/field_of_view.html
 */
func (p *PerspectiveCamera) SetLens( focalLength, frameHeight float64) {
	if frameHeight == -1 {
		frameHeight = 24.0
	};

	p.Fov = 2 * math3d.RadToDeg( math.Atan( frameHeight / ( focalLength * 2 ) ) )
	p.UpdateProjectionMatrix();
}


/**
 * Sets an offset in a larger frustum. This is useful for multi-window or
 * multi-monitor/multi-machine setups.
 *
 * For example, if you have 3x2 monitors and each monitor is 1920x1080 and
 * the monitors are in grid like this
 *
 *   +---+---+---+
 *   | A | B | C |
 *   +---+---+---+
 *   | D | E | F |
 *   +---+---+---+
 *
 * then for each monitor you would call it like this
 *
 *   var w = 1920;
 *   var h = 1080;
 *   var fullWidth = w * 3;
 *   var fullHeight = h * 2;
 *
 *   --A--
 *   camera.setOffset( fullWidth, fullHeight, w * 0, h * 0, w, h );
 *   --B--
 *   camera.setOffset( fullWidth, fullHeight, w * 1, h * 0, w, h );
 *   --C--
 *   camera.setOffset( fullWidth, fullHeight, w * 2, h * 0, w, h );
 *   --D--
 *   camera.setOffset( fullWidth, fullHeight, w * 0, h * 1, w, h );
 *   --E--
 *   camera.setOffset( fullWidth, fullHeight, w * 1, h * 1, w, h );
 *   --F--
 *   camera.setOffset( fullWidth, fullHeight, w * 2, h * 1, w, h );
 *
 *   Note there is no reason monitors have to be the same size or in a grid.
 */

func (p *PerspectiveCamera) SetViewOffset(fullWidth, fullHeight, x, y, width, height float64) {

	p.FullWidth = fullWidth
	p.FullHeight = fullHeight
	p.X = x
	p.Y = y
	p.Width = width
	p.Height = height

	p.UpdateProjectionMatrix()
}


func (p *PerspectiveCamera) UpdateProjectionMatrix() {

	var fov = math3d.RadToDeg( 2 * math.Atan( math.Tan( math3d.DegToRad( p.Fov ) * 0.5 ) / p.Zoom ) )

	if p.FullWidth > 0 {

		var aspect = p.FullWidth / p.FullHeight
		var top = math.Tan( math3d.DegToRad( fov * 0.5 ) ) * p.Near
		var bottom = - top
		var left = aspect * bottom
		var right = aspect * top
		var width = math.Abs( right - left )
		var height = math.Abs( top - bottom )

		p.ProjectionMatrix.MakeFrustum(
			left + p.X * width / p.FullWidth,
			left + ( p.X + p.Width ) * width / p.FullWidth,
			top - ( p.Y + p.Height ) * height / p.FullHeight,
			top - p.Y * height / p.FullHeight,
			p.Near,
			p.Far,
		)
	} else {
		p.ProjectionMatrix.MakePerspective( fov, p.Aspect, p.Near, p.Far )
	}
}

func (p *PerspectiveCamera) Copy(source *PerspectiveCamera) (*PerspectiveCamera) {
	p.Camera.Copy(source.Camera)

	p.Fov = source.Fov
	p.Aspect = source.Aspect
	p.Near = source.Near
	p.Far = source.Far

	p.Zoom = source.Zoom

	return p
}

/*
THREE.PerspectiveCamera.prototype.toJSON = function ( meta ) {

	var data = THREE.Object3D.prototype.toJSON.call( this, meta );

	data.object.zoom = p.zoom;
	data.object.fov = p.fov;
	data.object.aspect = p.aspect;
	data.object.near = p.near;
	data.object.far = p.far;

	return data;

};
*/