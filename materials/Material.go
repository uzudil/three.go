package materials

import (
	math3d "github.com/uzudil/three.go/math"
	three "github.com/uzudil/three.go"
	"github.com/uzudil/three.go/core"
	"fmt"
)

type Material struct {
	*core.EventDispatcher
	Id int
	Uuid string
	Name string
	Type string
	Side int
	Opacity float64
	Transparent bool
	Blending int
	BlendSrc, BlendDst, BlendEquation int
	BlendSrcAlpha, BlendDstAlpha float64
	BlendEquationAlpha int
	DepthFunc int
	DepthTest, DepthWrite bool
	ColorWrite bool
	Precision int
	PolygonOffset bool
	PolygonOffsetFactor, PolygonOffsetUnits int
	AlphaTest int
	Overdraw int
	Visible bool
	needsUpdate bool
}

var MaterialIdCount int = 0

func NewMaterial() (*Material) {
	MaterialIdCount++
	m := &Material{
		core.NewEventDispatcher(),
		Id: MaterialIdCount,
		Uuid: math3d.GenerateUUID(),
		Name: "",
		Type: "Material",
		Opacity: 1.0,
		Transparent: false,
		Blending: three.NormalBlending,
		BlendSrc: three.SrcAlphaFactor,
		BlendDst: three.OneMinusSrcAlphaFactor,
		BlendEquation: three.AddEquation,
		BlendSrcAlpha: nil,
		BlendDstAlpha: nil,
		BlendEquationAlpha: nil,
		DepthFunc: three.LessEqualDepth,
		DepthTest: true,
		DepthWrite: true,
		ColorWrite: true,
		Precision: nil, // override the renderer's default precision for this material
		PolygonOffset: false,
		PolygonOffsetFactor: 0,
		PolygonOffsetUnits: 0,
		AlphaTest: 0,
		Overdraw: 0, // Overdrawn pixels (typically between 0 and 1) for fixing antialiasing gaps in CanvasRenderer
		Visible: true,
		needsUpdate: true,
	}
	return m
}

func (m *Material) GetNeedsUpdate() bool {
	return m.needsUpdate
}

func (m *Material) SetNeedsUpdate(value bool) {
	if value == true {
		m.Update()
	}
	m.needsUpdate = value
}

func (m *Material) SetValues(values map[string]interface{}) {
	if values == nil {
		return
	}

	for key, newValue := range values {
		if newValue == nil {
			fmt.Println("THREE.Material: '%s'' parameter is undefined.", key)
			continue
		}

		currentValue := three.GetField(m, key)
		if currentValue == nil {
			fmt.Println("THREE.%s: '%s'' is not a property of this material.", m.Type, key)
			continue;
		}

		if _, ok := m.(math3d.Color); ok {
			currentValue.(math3d.Color).Copy(newValue.(math3d.Color))
		} else if _, ok := currentValue.(math3d.Vector3); ok {
			if _, ok := newValue.(math3d.Vector3); ok {
				currentValue.(math3d.Vector3).Copy(newValue.(math3d.Vector3))
			}
		} else {
			three.SetField(m, key, newValue)
		}
	}
}

/*
	toJSON: function ( meta ) {

		var data = {
			metadata: {
				version: 4.4,
				type: 'Material',
				generator: 'Material.toJSON'
			}
		};

		// standard Material serialization
		data.uuid = m.uuid;
		data.type = m.type;
		if ( m.name !== '' ) data.name = m.name;

		if ( m.color instanceof THREE.Color ) data.color = m.color.getHex();
		if ( m.emissive instanceof THREE.Color ) data.emissive = m.emissive.getHex();
		if ( m.specular instanceof THREE.Color ) data.specular = m.specular.getHex();
		if ( m.shininess !== undefined ) data.shininess = m.shininess;

		if ( m.map instanceof THREE.Texture ) data.map = m.map.toJSON( meta ).uuid;
		if ( m.alphaMap instanceof THREE.Texture ) data.alphaMap = m.alphaMap.toJSON( meta ).uuid;
		if ( m.lightMap instanceof THREE.Texture ) data.lightMap = m.lightMap.toJSON( meta ).uuid;
		if ( m.bumpMap instanceof THREE.Texture ) {

			data.bumpMap = m.bumpMap.toJSON( meta ).uuid;
			data.bumpScale = m.bumpScale;

		}
		if ( m.normalMap instanceof THREE.Texture ) {

			data.normalMap = m.normalMap.toJSON( meta ).uuid;
			data.normalScale = m.normalScale; // Removed for now, causes issue in editor ui.js

		}
		if ( m.displacementMap instanceof THREE.Texture ) {

			data.displacementMap = m.displacementMap.toJSON( meta ).uuid;
			data.displacementScale = m.displacementScale;
			data.displacementBias = m.displacementBias;

		}
		if ( m.specularMap instanceof THREE.Texture ) data.specularMap = m.specularMap.toJSON( meta ).uuid;
		if ( m.envMap instanceof THREE.Texture ) {

			data.envMap = m.envMap.toJSON( meta ).uuid;
			data.reflectivity = m.reflectivity; // Scale behind envMap

		}

		if ( m.size !== undefined ) data.size = m.size;
		if ( m.sizeAttenuation !== undefined ) data.sizeAttenuation = m.sizeAttenuation;

		if ( m.vertexColors !== undefined && m.vertexColors !== THREE.NoColors ) data.vertexColors = m.vertexColors;
		if ( m.shading !== undefined && m.shading !== THREE.SmoothShading ) data.shading = m.shading;
		if ( m.blending !== undefined && m.blending !== THREE.NormalBlending ) data.blending = m.blending;
		if ( m.side !== undefined && m.side !== THREE.FrontSide ) data.side = m.side;

		if ( m.opacity < 1 ) data.opacity = m.opacity;
		if ( m.transparent === true ) data.transparent = m.transparent;
		if ( m.alphaTest > 0 ) data.alphaTest = m.alphaTest;
		if ( m.wireframe === true ) data.wireframe = m.wireframe;
		if ( m.wireframeLinewidth > 1 ) data.wireframeLinewidth = m.wireframeLinewidth;

		return data;

	},
*/
func (m *Material) Clone() (*Material) {
	return NewMaterial().Copy(m)
}

func (m *Material) Copy(source *Material) (*Material) {
	m.Name = source.Name

	m.Side = source.Side

	m.Opacity = source.Opacity
	m.Transparent = source.Transparent

	m.Blending = source.Blending

	m.BlendSrc = source.BlendSrc
	m.BlendDst = source.BlendDst
	m.BlendEquation = source.BlendEquation
	m.BlendSrcAlpha = source.BlendSrcAlpha
	m.BlendDstAlpha = source.BlendDstAlpha
	m.BlendEquationAlpha = source.BlendEquationAlpha

	m.DepthFunc = source.DepthFunc
	m.DepthTest = source.DepthTest
	m.DepthWrite = source.DepthWrite

	m.Precision = source.Precision

	m.PolygonOffset = source.PolygonOffset
	m.PolygonOffsetFactor = source.PolygonOffsetFactor
	m.PolygonOffsetUnits = source.PolygonOffsetUnits

	m.AlphaTest = source.AlphaTest

	m.Overdraw = source.Overdraw

	m.Visible = source.Visible

	return m
}

func (m *Material) Update() {
	m.DispatchEvent( core.NewEvent("update") )
}

func (m *Material) Dispose() {
	m.DispatchEvent( core.NewEvent("dispose") )
}
