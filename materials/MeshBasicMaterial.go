package materials

import (
	math3d "github.com/uzudil/three.go/math"
	three "github.com/uzudil/three.go"
)

/**
 * @author mrdoob / http://mrdoob.com/
 * @author alteredq / http://alteredqualia.com/
 *
 * parameters = {
 *  color: <hex>,
 *  opacity: <float>,
 *  map: new THREE.Texture( <Image> ),
 *
 *  aoMap: new THREE.Texture( <Image> ),
 *  aoMapIntensity: <float>
 *
 *  specularMap: new THREE.Texture( <Image> ),
 *
 *  alphaMap: new THREE.Texture( <Image> ),
 *
 *  envMap: new THREE.TextureCube( [posx, negx, posy, negy, posz, negz] ),
 *  combine: THREE.Multiply,
 *  reflectivity: <float>,
 *  refractionRatio: <float>,
 *
 *  shading: THREE.SmoothShading,
 *  blending: THREE.NormalBlending,
 *  depthTest: <bool>,
 *  depthWrite: <bool>,
 *
 *  wireframe: <boolean>,
 *  wireframeLinewidth: <float>,
 *
 *  vertexColors: THREE.NoColors / THREE.VertexColors / THREE.FaceColors,
 *
 *  skinning: <bool>,
 *  morphTargets: <bool>,
 *
 *  fog: <bool>
 * }
 */

type MeshBasicMaterial struct {
	*Material
	Color *math3d.Color
	Map string
	AoMap string
	AoMapIntensity float64
	SpecularMap, AlphaMap, EnvMap string
	Combine int
	Reflectivity int
	RefractionRatio float64
	Fog bool
	Shading int
	Wireframe bool
	WireframeLinewidth int
	WireframeLinecap, WireframeLinejoin string
	VertexColors int
	Skinning bool
	MorphTargets bool
}

func NewMeshBasicMaterial(parameters map[string]interface{}) (*MeshBasicMaterial) {
	m := &MeshBasicMaterial{
		NewMaterial(),
	}
	m.Type = "MeshBasicMaterial"
	m.Color = math3d.NewColor(1.0, 1.0, 1.0) // emissive
	m.Map = nil
	m.AoMap = nil
	m.AoMapIntensity = 1.0
	m.SpecularMap = nil
	m.AlphaMap = nil
	m.EnvMap = nil
	m.Combine = three.MultiplyOperation
	m.Reflectivity = 1
	m.RefractionRatio = 0.98
	m.Fog = true
	m.Shading = three.SmoothShading
	m.Wireframe = false
	m.WireframeLinewidth = 1
	m.WireframeLinecap = "round"
	m.WireframeLinejoin = "round"
	m.VertexColors = three.NoColors
	m.Skinning = false
	m.MorphTargets = false

	m.SetValues( parameters )

	return m
}

func (m *MeshBasicMaterial) Copy(source *MeshBasicMaterial) (*MeshBasicMaterial) {
	m.Material.Copy(source.Material)
	m.Color.Copy(source.Color)
	m.Map = source.Map
	m.AoMap = source.AoMap
	m.AoMapIntensity = source.AoMapIntensity
	m.SpecularMap = source.SpecularMap
	m.AlphaMap = source.AlphaMap
	m.EnvMap = source.EnvMap
	m.Combine = source.Combine
	m.Reflectivity = source.Reflectivity
	m.RefractionRatio = source.RefractionRatio
	m.Fog = source.Fog
	m.Shading = source.Shading
	m.Wireframe = source.Wireframe
	m.WireframeLinewidth = source.WireframeLinewidth
	m.WireframeLinecap = source.WireframeLinecap
	m.WireframeLinejoin = source.WireframeLinejoin
	m.VertexColors = source.VertexColors
	m.Skinning = source.Skinning
	m.MorphTargets = source.MorphTargets
	return m
}
