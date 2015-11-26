package scenes
import "github.com/uzudil/three.go/core"

type Scene struct {
	*core.Object3D
	AutoUpdate bool
}

func NewScene() (*Scene) {
	scene := &Scene{
		core.NewObject3D(),
		true,
	}
	scene.Type = "Scene"
//	this.fog = null;
//	this.overrideMaterial = null;
	return scene
}

func (scene *Scene) Copy(source *Scene) (*Scene) {
	scene.Object3D.Copy(source.Object3D)

//	if ( source.fog !== null ) this.fog = source.fog.clone();
//	if ( source.overrideMaterial !== null ) this.overrideMaterial = source.overrideMaterial.clone();

	scene.AutoUpdate = source.AutoUpdate
	scene.MatrixAutoUpdate = source.MatrixAutoUpdate

	return scene
}
