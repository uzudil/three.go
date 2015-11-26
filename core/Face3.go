package core
import "github.com/uzudil/three.go/math"

type Face3 struct {
	A, B, C int
	Normal *math.Vector3
	VertexNormals []*math.Vector3
	Color *math.Color
	VertexColors []*math.Color
	MaterialIndex int
}

func NewDefaultFace3(a, b, c int) (*Face3) {
	return NewFace3(a, b, c, math.NewEmptyVector3(), math.NewDefaultColor(), 0)
}

func NewFace3(a, b, c int, normal *math.Vector3, color *math.Color, materialIndex int) (*Face3) {
	return &Face3{
		A: a,
		B: b,
		C: c,
		Normal: normal,
		VertexNormals: make([]*math.Vector3, 0),
		Color: color,
		VertexColors: make([]*math.Color, 0),
		MaterialIndex: materialIndex,
	}
}

func NewArraysFace3(a, b, c int, normals []*math.Vector3, colors []*math.Color, materialIndex int) (*Face3) {
	return &Face3{
		A: a,
		B: b,
		C: c,
		Normal: math.NewEmptyVector3(),
		VertexNormals: normals,
		Color: math.NewDefaultColor(),
		VertexColors: colors,
		MaterialIndex: materialIndex,
	}
}

func (f *Face3) Clone() (*Face3) {
	return (&Face3{}).Copy(f)
}

func (f *Face3) Copy(source *Face3) (*Face3) {
	f.A = source.A
	f.B = source.B
	f.C = source.C

	f.Normal.Copy( source.Normal )
	f.Color.Copy( source.Color )

	f.MaterialIndex = source.MaterialIndex

	// todo: this won't work if source/target have different length VertexNormals or VertexColors
	il := len(source.VertexNormals)
	for i := 0; i < il; i ++ {
		f.VertexNormals[ i ] = source.VertexNormals[ i ].Clone()
	}

	il = len(source.VertexColors)
	for i := 0; i < il; i ++ {
		f.VertexColors[ i ] = source.VertexColors[ i ].Clone()
	}

	return f
}
