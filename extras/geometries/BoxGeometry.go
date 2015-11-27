package geometries

import (
	"github.com/uzudil/three.go/core"
	math3d "github.com/uzudil/three.go/math"
	three "github.com/uzudil/three.go"
)

type BoxGeometry struct {
	*core.Geometry
	Width, Height, Depth float64
	WidthSegments, HeightSegments, DepthSegments int
	Parameters map[string]interface{}
}

func NewDefaultBoxGeometry(width, height, depth float64) (*BoxGeometry) {
	return NewBoxGeometry(width, height, depth, 1, 1, 1)
}

func NewBoxGeometry(width, height, depth float64, widthSegments, heightSegments, depthSegments int) (*BoxGeometry) {
	g := &BoxGeometry{
		core.NewGeometry(),
		Width: width,
		Height: height,
		Depth: depth,
		WidthSegments: widthSegments,
		HeightSegments: heightSegments,
		DepthSegments: depthSegments,
	}
	g.Type = "BoxGeometry"

	g.Parameters = map[string]interface{}{
		"width": width,
		"height": height,
		"depth": depth,
		"widthSegments": widthSegments,
		"heightSegments": heightSegments,
		"depthSegments": depthSegments,
	};

	scope := g

	width_half := width / 2
	height_half := height / 2
	depth_half := depth / 2

	buildPlane := func( u, v string, udir, vdir int, width, height, depth float64, materialIndex int) {

		// w, ix, iy,
		var w string
		gridX := scope.WidthSegments
		gridY := scope.HeightSegments
		width_half := width / 2
		height_half := height / 2
		offset := len(scope.Vertices)

		if ( u == "x" && v == "y" ) || ( u == "y" && v == "x" ) {
			w = "z";
		} else if ( u == "x" && v == "z" ) || ( u == "z" && v == "x" ) {
			w = "y";
			gridY = scope.DepthSegments
		} else if ( u == "z" && v == "y" ) || ( u == "y" && v == "z" ) {
			w = "x";
			gridX = scope.DepthSegments
		}

		gridX1 := gridX + 1
		gridY1 := gridY + 1
		segment_width := width / gridX
		segment_height := height / gridY
		normal := math3d.NewEmptyVector3()

		if depth > 0 {
			three.SetField(normal, w, 1)
		} else {
			three.SetField(normal, w, -1)
		}

		for iy := 0; iy < gridY1; iy ++ {
			for ix := 0; ix < gridX1; ix ++ {
				var vector = math3d.NewEmptyVector3()
				three.SetField(vector, u, ( ix * segment_width - width_half ) * udir)
				three.SetField(vector, v, ( iy * segment_height - height_half ) * vdir)
				three.SetField(vector, w, depth)
				scope.Vertices = append(scope.Vertices, vector)
			}
		}

		for iy := 0; iy < gridY; iy++ {
			for ix := 0; ix < gridX; ix++ {
				a := ix + gridX1 * iy
				b := ix + gridX1 * ( iy + 1 )
				c := ( ix + 1 ) + gridX1 * ( iy + 1 )
				d := ( ix + 1 ) + gridX1 * iy

				uva := math3d.NewVector2(ix / gridX, 1 - iy / gridY)
				uvb := math3d.NewVector2( ix / gridX, 1 - ( iy + 1 ) / gridY )
				uvc := math3d.NewVector2( ( ix + 1 ) / gridX, 1 - ( iy + 1 ) / gridY )
				uvd := math3d.NewVector2( ( ix + 1 ) / gridX, 1 - iy / gridY )

				face := core.NewDefaultFace3(a + offset, b + offset, d + offset)
				face.Normal.Copy( normal );
				face.VertexNormals = append( face.VertexNormals, normal.Clone(), normal.Clone(), normal.Clone() )
				face.MaterialIndex = materialIndex

				scope.Faces = append(scope.Faces, face )
				scope.FaceVertexUvs[ 0 ] = append(scope.FaceVertexUvs[ 0 ], []*math3d.Vector2{ uva, uvb, uvd } )

				face = core.NewDefaultFace3( b + offset, c + offset, d + offset )
				face.Normal.Copy( normal );
				face.VertexNormals = append( face.VertexNormals, normal.Clone(), normal.Clone(), normal.Clone() )
				face.MaterialIndex = materialIndex

				scope.Faces = append(scope.Faces, face )
				scope.FaceVertexUvs[ 0 ] = append(scope.FaceVertexUvs[ 0 ], []*math3d.Vector2{ uvb.Clone(), uvc, uvd.Clone() } )
			}
		}
	}

	buildPlane( "z", "y", - 1, - 1, depth, height, width_half, 0 ) // px
	buildPlane( "z", "y",   1, - 1, depth, height, - width_half, 1 ) // nx
	buildPlane( "x", "z",   1,   1, width, depth, height_half, 2 ) // py
	buildPlane( "x", "z",   1, - 1, width, depth, - height_half, 3 ) // ny
	buildPlane( "x", "y",   1, - 1, width, height, depth_half, 4 ) // pz
	buildPlane( "x", "y", - 1, - 1, width, height, - depth_half, 5 ) // nz

	g.MergeVertices()
	return g
}

func (g *BoxGeometry) Clone() (*BoxGeometry) {
	p := g.Parameters
	return &BoxGeometry{
		Width: p["width"],
		Height: p["height"],
		Depth: p["depth"],
		WidthSegments: p["widthSegments"],
		HeightSegments: p["heightSegments"],
		DepthSegments: p["depthSegments"],
	}
}

