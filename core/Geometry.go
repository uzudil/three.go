package core
import (
	math3d "github.com/uzudil/three.go/math"
	"github.com/uzudil/three.go/objects"
)

type Geometry struct {
	*EventDispatcher
	Id int
	Uuid string
	Name string
	Type string
	Vertices []*math3d.Vector3
	Colors []*math3d.Color
	Faces []*Face3
	FaceVertexUvs []([]([]*math3d.Vector2))
	BoundingBox *math3d.Box3
	BoundingSphere *math3d.Sphere
	VerticesNeedUpdate bool
	ElementsNeedUpdate bool
	UvsNeedUpdate bool
	NormalsNeedUpdate bool
	ColorsNeedUpdate bool

	RotateX func(angle float64) (*Geometry)
	RotateY func(angle float64) (*Geometry)
	RotateZ func(angle float64) (*Geometry)
	Translate func(float64, float64, float64) (*Geometry)
	Scale func(float64, float64, float64) (*Geometry)
	LookAt func(*math3d.Vector3)
}

var GeometryIdCount int = 0

func NewGeometry() (*Geometry) {
	GeometryIdCount++
	g := &Geometry{
		NewEventDispatcher(),
		Id: GeometryIdCount,
		Uuid: math3d.GenerateUUID(),
		Name: "",
		Type: "Geometry",
		Vertices: make([]*math3d.Vector3, 0),
		Colors: make([]*math3d.Color, 0),
		Faces: make([]*Face3, 0),
		FaceVertexUvs: make([]([]([]*math3d.Vector2)), 0),
	}
	g.RotateX = g.buildRotateX()
	g.RotateY = g.buildRotateY()
	g.RotateZ = g.buildRotateZ()
	g.Translate = g.buildTranslate()
	g.Scale = g.buildScale()
	g.LookAt = g.buildLookAt()
	return g
}

func (g *Geometry) ApplyMatrix(matrix *math3d.Matrix4) {
	normalMatrix := math3d.NewMatrix3().GetNormalMatrix( matrix )

	for _, v := range g.Vertices {
		v.ApplyMatrix4( matrix )
	}

	for _, face := range g.Faces {
		face.Normal.ApplyMatrix3( normalMatrix ).Normalize()
		for _, vn := range face.VertexNormals {
			vn.ApplyMatrix3( normalMatrix ).Normalize()
		}
	}

	if g.BoundingBox != nil {
		g.ComputeBoundingBox()
	}

	if g.BoundingSphere != nil {
		g.ComputeBoundingSphere()
	}

	g.VerticesNeedUpdate = true
	g.NormalsNeedUpdate = true
}

func (g *Geometry) buildRotateX() (func(angle float64) (*Geometry)) {
	// rotate geometry around world x-axis
	var m1 math3d.Matrix4
	return func(angle float64) (*Geometry) {
		if m1 == nil {
			m1 = math3d.NewMatrix4()
		}
		m1.MakeRotationX( angle )
		g.ApplyMatrix( m1 )
		return g
	}
}

func (g *Geometry) buildRotateY() (func(angle float64) (*Geometry)) {
	// rotate geometry around world y-axis
	var m1 math3d.Matrix4
	return func(angle float64) (*Geometry) {
		if m1 == nil {
			m1 = math3d.NewMatrix4()
		}
		m1.MakeRotationY( angle )
		g.ApplyMatrix( m1 )
		return g
	}
}

func (g *Geometry) buildRotateZ() (func(angle float64) (*Geometry)) {
	// rotate geometry around world z-axis
	var m1 math3d.Matrix4
	return func(angle float64) (*Geometry) {
		if m1 == nil {
			m1 = math3d.NewMatrix4()
		}
		m1.MakeRotationZ( angle )
		g.ApplyMatrix( m1 )
		return g
	}
}

func (g *Geometry) buildTranslate() (func(float64, float64, float64) (*Geometry)) {
	// translate geometry
	var m1 math3d.Matrix4
	return func(x, y, z float64) (*Geometry) {
		if m1 == nil {
			m1 = math3d.NewMatrix4()
		}
		m1.MakeTranslation( x, y, z )
		g.ApplyMatrix( m1 )
		return g
	}
}

func (g *Geometry) buildScale() (func(float64, float64, float64) (*Geometry)) {
	// scale geometry
	var m1 math3d.Matrix4
	return func(x, y, z float64) (*Geometry) {
		if m1 == nil {
			m1 = math3d.NewMatrix4()
		}
		m1.MakeScale( x, y, z )
		g.ApplyMatrix( m1 )
		return g
	}
}

func (g *Geometry) buildLookAt() (func(*math3d.Vector3)) {
	var obj *Object3D
	return func(vector *math3d.Vector3) {
		if obj == nil {
			obj = NewObject3D()
		}
		obj.LookAt(vector)
		obj.UpdateMatrix()
		g.ApplyMatrix(obj.Matrix)
	}
}

/*
fromBufferGeometry: function ( geometry ) {

	var scope = this;

	var indices = geometry.index !== null ? geometry.index.array : undefined;
	var attributes = geometry.attributes;

	var vertices = attributes.position.array;
	var normals = attributes.normal !== undefined ? attributes.normal.array : undefined;
	var colors = attributes.color !== undefined ? attributes.color.array : undefined;
	var uvs = attributes.uv !== undefined ? attributes.uv.array : undefined;
	var uvs2 = attributes.uv2 !== undefined ? attributes.uv2.array : undefined;

	if ( uvs2 !== undefined ) g.faceVertexUvs[ 1 ] = [];

	var tempNormals = [];
	var tempUVs = [];
	var tempUVs2 = [];

	for ( var i = 0, j = 0, k = 0; i < vertices.length; i += 3, j += 2, k += 4 ) {

		scope.vertices.push( new THREE.Vector3( vertices[ i ], vertices[ i + 1 ], vertices[ i + 2 ] ) );

		if ( normals !== undefined ) {

			tempNormals.push( new THREE.Vector3( normals[ i ], normals[ i + 1 ], normals[ i + 2 ] ) );

		}

		if ( colors !== undefined ) {

			scope.colors.push( new THREE.Color( colors[ i ], colors[ i + 1 ], colors[ i + 2 ] ) );

		}

		if ( uvs !== undefined ) {

			tempUVs.push( new THREE.Vector2( uvs[ j ], uvs[ j + 1 ] ) );

		}

		if ( uvs2 !== undefined ) {

			tempUVs2.push( new THREE.Vector2( uvs2[ j ], uvs2[ j + 1 ] ) );

		}

	}

	function addFace( a, b, c ) {

		var vertexNormals = normals !== undefined ? [ tempNormals[ a ].clone(), tempNormals[ b ].clone(), tempNormals[ c ].clone() ] : [];
		var vertexColors = colors !== undefined ? [ scope.colors[ a ].clone(), scope.colors[ b ].clone(), scope.colors[ c ].clone() ] : [];

		var face = new THREE.Face3( a, b, c, vertexNormals, vertexColors );

		scope.faces.push( face );

		if ( uvs !== undefined ) {

			scope.faceVertexUvs[ 0 ].push( [ tempUVs[ a ].clone(), tempUVs[ b ].clone(), tempUVs[ c ].clone() ] );

		}

		if ( uvs2 !== undefined ) {

			scope.faceVertexUvs[ 1 ].push( [ tempUVs2[ a ].clone(), tempUVs2[ b ].clone(), tempUVs2[ c ].clone() ] );

		}

	};

	if ( indices !== undefined ) {

		var groups = geometry.groups;

		if ( groups.length > 0 ) {

			for ( var i = 0; i < groups.length; i ++ ) {

				var group = groups[ i ];

				var start = group.start;
				var count = group.count;

				for ( var j = start, jl = start + count; j < jl; j += 3 ) {

					addFace( indices[ j ], indices[ j + 1 ], indices[ j + 2 ] );

				}

			}

		} else {

			for ( var i = 0; i < indices.length; i += 3 ) {

				addFace( indices[ i ], indices[ i + 1 ], indices[ i + 2 ] );

			}

		}

	} else {

		for ( var i = 0; i < vertices.length / 3; i += 3 ) {

			addFace( i, i + 1, i + 2 );

		}

	}

	g.computeFaceNormals();

	if ( geometry.boundingBox !== null ) {

		g.boundingBox = geometry.boundingBox.clone();

	}

	if ( geometry.boundingSphere !== null ) {

		g.boundingSphere = geometry.boundingSphere.clone();

	}

	return this;

},
*/

func (g *Geometry) Center() (*math3d.Vector3) {
	g.ComputeBoundingBox()
	offset := g.BoundingBox.Center().Negate()
	g.Translate(offset.X, offset.Y, offset.Z)
	return offset
}

func (g *Geometry) Normalize() (*Geometry) {
	g.ComputeBoundingSphere()

	center := g.BoundingSphere.Center;
	radius := g.BoundingSphere.Radius;

	s := 1
	if radius != 0 {
		s = 1.0 / radius
	};

	matrix := math3d.NewMatrix4()
	matrix.Set(
		s, 0, 0, - s * center.X,
		0, s, 0, - s * center.Y,
		0, 0, s, - s * center.Z,
		0, 0, 0, 1,
	)
	g.ApplyMatrix( matrix )
	return g
}

func (g *Geometry) ComputeFaceNormals() {
	cb := math3d.NewEmptyVector3()
	ab := math3d.NewEmptyVector3()
	for _, face := range g.Faces {
		var vA = g.Vertices[ face.A ]
		var vB = g.Vertices[ face.B ]
		var vC = g.Vertices[ face.C ]

		cb.SubVectors( vC, vB )
		ab.SubVectors( vA, vB )
		cb.Cross( ab )

		cb.Normalize()

		face.Normal.Copy( cb )
	}
}

func (g *Geometry) ComputeVertexNormals(areaWeighted bool) {
	vertices := make([]*math3d.Vector3, len(g.Vertices))
	for v, _ := range g.Vertices {
		vertices[v] = math3d.NewEmptyVector3()
	}

	if ( areaWeighted ) {
		// vertex normals weighted by triangle areas
		// http://www.iquilezles.org/www/articles/normals/normals.htm
		var vA, vB, vC *math3d.Vector3
		cb := math3d.NewEmptyVector3()
		ab := math3d.NewEmptyVector3()
		for _, face := range g.Faces {
			vA = g.Vertices[ face.A ]
			vB = g.Vertices[ face.B ]
			vC = g.Vertices[ face.C ]

			cb.SubVectors( vC, vB )
			ab.SubVectors( vA, vB )
			cb.Cross( ab )

			vertices[ face.A ].Add( cb )
			vertices[ face.B ].Add( cb )
			vertices[ face.C ].Add( cb )
		}
	} else {
		for _, face := range g.Faces {
			vertices[ face.A ].Add( face.Normal )
			vertices[ face.B ].Add( face.Normal )
			vertices[ face.C ].Add( face.Normal )
		}
	}

	for v, _ := range g.Vertices {
		vertices[ v ].Normalize()
	}

	for _, face := range g.Faces {
		vertexNormals := face.VertexNormals

		if len(vertexNormals) == 3 {
			vertexNormals[ 0 ].Copy( vertices[ face.A ] )
			vertexNormals[ 1 ].Copy( vertices[ face.B ] )
			vertexNormals[ 2 ].Copy( vertices[ face.C ] )
		} else {
			vertexNormals[ 0 ] = vertices[ face.A ].Clone()
			vertexNormals[ 1 ] = vertices[ face.B ].Clone()
			vertexNormals[ 2 ] = vertices[ face.C ].Clone()
		}
	}
}
/*
computeMorphNormals: function () {

	var i, il, f, fl, face;

	// save original normals
	// - create temp variables on first access
	//   otherwise just copy (for faster repeated calls)

	for ( f = 0, fl = g.faces.length; f < fl; f ++ ) {

		face = g.faces[ f ];

		if ( ! face.__originalFaceNormal ) {

			face.__originalFaceNormal = face.normal.clone();

		} else {

			face.__originalFaceNormal.copy( face.normal );

		}

		if ( ! face.__originalVertexNormals ) face.__originalVertexNormals = [];

		for ( i = 0, il = face.vertexNormals.length; i < il; i ++ ) {

			if ( ! face.__originalVertexNormals[ i ] ) {

				face.__originalVertexNormals[ i ] = face.vertexNormals[ i ].clone();

			} else {

				face.__originalVertexNormals[ i ].copy( face.vertexNormals[ i ] );

			}

		}

	}

	// use temp geometry to compute face and vertex normals for each morph

	var tmpGeo = new THREE.Geometry();
	tmpGeo.faces = g.faces;

	for ( i = 0, il = g.morphTargets.length; i < il; i ++ ) {

		// create on first access

		if ( ! g.morphNormals[ i ] ) {

			g.morphNormals[ i ] = {};
			g.morphNormals[ i ].faceNormals = [];
			g.morphNormals[ i ].vertexNormals = [];

			var dstNormalsFace = g.morphNormals[ i ].faceNormals;
			var dstNormalsVertex = g.morphNormals[ i ].vertexNormals;

			var faceNormal, vertexNormals;

			for ( f = 0, fl = g.faces.length; f < fl; f ++ ) {

				faceNormal = new THREE.Vector3();
				vertexNormals = { a: new THREE.Vector3(), b: new THREE.Vector3(), c: new THREE.Vector3() };

				dstNormalsFace.push( faceNormal );
				dstNormalsVertex.push( vertexNormals );

			}

		}

		var morphNormals = g.morphNormals[ i ];

		// set vertices to morph target

		tmpGeo.vertices = g.morphTargets[ i ].vertices;

		// compute morph normals

		tmpGeo.computeFaceNormals();
		tmpGeo.computeVertexNormals();

		// store morph normals

		var faceNormal, vertexNormals;

		for ( f = 0, fl = g.faces.length; f < fl; f ++ ) {

			face = g.faces[ f ];

			faceNormal = morphNormals.faceNormals[ f ];
			vertexNormals = morphNormals.vertexNormals[ f ];

			faceNormal.copy( face.normal );

			vertexNormals.a.copy( face.vertexNormals[ 0 ] );
			vertexNormals.b.copy( face.vertexNormals[ 1 ] );
			vertexNormals.c.copy( face.vertexNormals[ 2 ] );

		}

	}

	// restore original normals

	for ( f = 0, fl = g.faces.length; f < fl; f ++ ) {

		face = g.faces[ f ];

		face.normal = face.__originalFaceNormal;
		face.vertexNormals = face.__originalVertexNormals;

	}

},

computeTangents: function () {

	console.warn( 'THREE.Geometry: .computeTangents() has been removed.' );

},

computeLineDistances: function () {

	var d = 0;
	var vertices = g.vertices;

	for ( var i = 0, il = vertices.length; i < il; i ++ ) {

		if ( i > 0 ) {

			d += vertices[ i ].distanceTo( vertices[ i - 1 ] );

		}

		g.lineDistances[ i ] = d;

	}

},
*/

func (g *Geometry) ComputeBoundingBox() {
	if g.BoundingBox == nil {
		g.BoundingBox = math3d.NewDefaultBox3()
	}
	g.BoundingBox.SetFromPoints( g.Vertices )
}

func (g *Geometry) ComputeBoundingSphere() {
	if g.BoundingSphere == nil {
		g.BoundingSphere = math3d.NewDefaultSphere()
	}
	g.BoundingSphere.SetFromPoints(g.Vertices)
}

func (g *Geometry) Merge(geometry *Geometry, matrix *math3d.Matrix4, materialIndexOffset int) {
	var normalMatrix *math3d.Matrix3
	vertexOffset := len(g.Vertices)
	vertices1 := g.Vertices
	vertices2 := geometry.Vertices
	faces1 := g.Faces
	faces2 := geometry.Faces
	uvs1 := g.FaceVertexUvs[ 0 ]
	uvs2 := geometry.FaceVertexUvs[ 0 ]

	if materialIndexOffset == nil {
		materialIndexOffset = 0
	}

	if matrix != nil {
		normalMatrix = math3d.NewMatrix3().GetNormalMatrix( matrix )
	}

	// vertices
	for _, vertex := range vertices2 {
		vertexCopy := vertex.Clone()
		if matrix != nil {
			vertexCopy.ApplyMatrix4( matrix )
		}
		vertices1 = append(vertices1, vertexCopy )
	}

	// faces
	for _, face := range faces2 {

		faceVertexNormals := face.VertexNormals
		faceVertexColors := face.VertexColors

		faceCopy := NewDefaultFace3( face.A + vertexOffset, face.B + vertexOffset, face.C + vertexOffset )
		faceCopy.Normal.Copy( face.Normal )

		if normalMatrix != nil {
			faceCopy.Normal.ApplyMatrix3( normalMatrix ).Normalize()
		}

		for _, fn := range faceVertexNormals {
			normal := fn.Clone()
			if normalMatrix != nil {
				normal.ApplyMatrix3( normalMatrix ).Normalize()
			}
			faceCopy.VertexNormals = append(faceCopy.VertexNormals, normal )
		}

		faceCopy.Color.Copy( face.Color )
		for _, color := range faceVertexColors {
			faceCopy.VertexColors = append(faceCopy.VertexColors, color.Clone() )
		}

		faceCopy.MaterialIndex = face.MaterialIndex + materialIndexOffset
		faces1 = append(faces1, faceCopy )
	}

	// uvs
	for _, uv := range uvs2 {
		if uv == nil {
			continue;
		}
		uvCopy := make([]*math3d.Vector2, len(uv))
		for _, u := range uv {
			uvCopy = append(uvCopy, u.Clone())
		}
		uvs1 = append(uvs1, uvCopy)
	}
}

func (g *Geometry) MergeMesh(mesh *objects.Mesh) {
	mesh.MatrixAutoUpdate && mesh.UpdateMatrix()
	g.Merge( mesh.Geometry, mesh.Matrix, 0 )
}

/*
 * Checks for duplicate vertices with hashmap.
 * Duplicated vertices are removed
 * and faces' vertices are updated.
 */
/*
func (g *Geometry) MergeVertices() int {

	var verticesMap = {}; // Hashmap for looking up vertices by position coordinates (and making sure they are unique)
	var unique = [], changes = [];

	var v, key;
	var precisionPoints = 4; // number of decimal points, e.g. 4 for epsilon of 0.0001
	var precision = Math.pow( 10, precisionPoints );
	var i, il, face;
	var indices, j, jl;

	for ( i = 0, il = g.vertices.length; i < il; i ++ ) {

		v = g.vertices[ i ];
		key = Math.round( v.x * precision ) + '_' + Math.round( v.y * precision ) + '_' + Math.round( v.z * precision );

		if ( verticesMap[ key ] === undefined ) {

			verticesMap[ key ] = i;
			unique.push( g.vertices[ i ] );
			changes[ i ] = unique.length - 1;

		} else {

			//console.log('Duplicate vertex found. ', i, ' could be using ', verticesMap[key]);
			changes[ i ] = changes[ verticesMap[ key ] ];

		}

	}


	// if faces are completely degenerate after merging vertices, we
	// have to remove them from the geometry.
	var faceIndicesToRemove = [];

	for ( i = 0, il = g.faces.length; i < il; i ++ ) {

		face = g.faces[ i ];

		face.a = changes[ face.a ];
		face.b = changes[ face.b ];
		face.c = changes[ face.c ];

		indices = [ face.a, face.b, face.c ];

		var dupIndex = - 1;

		// if any duplicate vertices are found in a Face3
		// we have to remove the face as nothing can be saved
		for ( var n = 0; n < 3; n ++ ) {

			if ( indices[ n ] === indices[ ( n + 1 ) % 3 ] ) {

				dupIndex = n;
				faceIndicesToRemove.push( i );
				break;

			}

		}

	}

	for ( i = faceIndicesToRemove.length - 1; i >= 0; i -- ) {

		var idx = faceIndicesToRemove[ i ];

		g.faces.splice( idx, 1 );

		for ( j = 0, jl = g.faceVertexUvs.length; j < jl; j ++ ) {

			g.faceVertexUvs[ j ].splice( idx, 1 );

		}

	}

	// Use unique set of vertices

	var diff = g.vertices.length - unique.length;
	g.vertices = unique;
	return diff;

},

sortFacesByMaterialIndex: function () {

	var faces = g.faces;
	var length = faces.length;

	// tag faces

	for ( var i = 0; i < length; i ++ ) {

		faces[ i ]._id = i;

	}

	// sort faces

	function materialIndexSort( a, b ) {

		return a.materialIndex - b.materialIndex;

	}

	faces.sort( materialIndexSort );

	// sort uvs

	var uvs1 = g.faceVertexUvs[ 0 ];
	var uvs2 = g.faceVertexUvs[ 1 ];

	var newUvs1, newUvs2;

	if ( uvs1 && uvs1.length === length ) newUvs1 = [];
	if ( uvs2 && uvs2.length === length ) newUvs2 = [];

	for ( var i = 0; i < length; i ++ ) {

		var id = faces[ i ]._id;

		if ( newUvs1 ) newUvs1.push( uvs1[ id ] );
		if ( newUvs2 ) newUvs2.push( uvs2[ id ] );

	}

	if ( newUvs1 ) g.faceVertexUvs[ 0 ] = newUvs1;
	if ( newUvs2 ) g.faceVertexUvs[ 1 ] = newUvs2;

},

toJSON: function () {

	var data = {
		metadata: {
			version: 4.4,
			type: 'Geometry',
			generator: 'Geometry.toJSON'
		}
	};

	// standard Geometry serialization

	data.uuid = g.uuid;
	data.type = g.type;
	if ( g.name !== '' ) data.name = g.name;

	if ( g.parameters !== undefined ) {

		var parameters = g.parameters;

		for ( var key in parameters ) {

			if ( parameters[ key ] !== undefined ) data[ key ] = parameters[ key ];

		}

		return data;

	}

	var vertices = [];

	for ( var i = 0; i < g.vertices.length; i ++ ) {

		var vertex = g.vertices[ i ];
		vertices.push( vertex.x, vertex.y, vertex.z );

	}

	var faces = [];
	var normals = [];
	var normalsHash = {};
	var colors = [];
	var colorsHash = {};
	var uvs = [];
	var uvsHash = {};

	for ( var i = 0; i < g.faces.length; i ++ ) {

		var face = g.faces[ i ];

		var hasMaterial = false; // face.materialIndex !== undefined;
		var hasFaceUv = false; // deprecated
		var hasFaceVertexUv = g.faceVertexUvs[ 0 ][ i ] !== undefined;
		var hasFaceNormal = face.normal.length() > 0;
		var hasFaceVertexNormal = face.vertexNormals.length > 0;
		var hasFaceColor = face.color.r !== 1 || face.color.g !== 1 || face.color.b !== 1;
		var hasFaceVertexColor = face.vertexColors.length > 0;

		var faceType = 0;

		faceType = setBit( faceType, 0, 0 );
		faceType = setBit( faceType, 1, hasMaterial );
		faceType = setBit( faceType, 2, hasFaceUv );
		faceType = setBit( faceType, 3, hasFaceVertexUv );
		faceType = setBit( faceType, 4, hasFaceNormal );
		faceType = setBit( faceType, 5, hasFaceVertexNormal );
		faceType = setBit( faceType, 6, hasFaceColor );
		faceType = setBit( faceType, 7, hasFaceVertexColor );

		faces.push( faceType );
		faces.push( face.a, face.b, face.c );

		if ( hasFaceVertexUv ) {

			var faceVertexUvs = g.faceVertexUvs[ 0 ][ i ];

			faces.push(
				getUvIndex( faceVertexUvs[ 0 ] ),
				getUvIndex( faceVertexUvs[ 1 ] ),
				getUvIndex( faceVertexUvs[ 2 ] )
			);

		}

		if ( hasFaceNormal ) {

			faces.push( getNormalIndex( face.normal ) );

		}

		if ( hasFaceVertexNormal ) {

			var vertexNormals = face.vertexNormals;

			faces.push(
				getNormalIndex( vertexNormals[ 0 ] ),
				getNormalIndex( vertexNormals[ 1 ] ),
				getNormalIndex( vertexNormals[ 2 ] )
			);

		}

		if ( hasFaceColor ) {

			faces.push( getColorIndex( face.color ) );

		}

		if ( hasFaceVertexColor ) {

			var vertexColors = face.vertexColors;

			faces.push(
				getColorIndex( vertexColors[ 0 ] ),
				getColorIndex( vertexColors[ 1 ] ),
				getColorIndex( vertexColors[ 2 ] )
			);

		}

	}

	function setBit( value, position, enabled ) {

		return enabled ? value | ( 1 << position ) : value & ( ~ ( 1 << position ) );

	}

	function getNormalIndex( normal ) {

		var hash = normal.x.toString() + normal.y.toString() + normal.z.toString();

		if ( normalsHash[ hash ] !== undefined ) {

			return normalsHash[ hash ];

		}

		normalsHash[ hash ] = normals.length / 3;
		normals.push( normal.x, normal.y, normal.z );

		return normalsHash[ hash ];

	}

	function getColorIndex( color ) {

		var hash = color.r.toString() + color.g.toString() + color.b.toString();

		if ( colorsHash[ hash ] !== undefined ) {

			return colorsHash[ hash ];

		}

		colorsHash[ hash ] = colors.length;
		colors.push( color.getHex() );

		return colorsHash[ hash ];

	}

	function getUvIndex( uv ) {

		var hash = uv.x.toString() + uv.y.toString();

		if ( uvsHash[ hash ] !== undefined ) {

			return uvsHash[ hash ];

		}

		uvsHash[ hash ] = uvs.length / 2;
		uvs.push( uv.x, uv.y );

		return uvsHash[ hash ];

	}

	data.data = {};

	data.data.vertices = vertices;
	data.data.normals = normals;
	if ( colors.length > 0 ) data.data.colors = colors;
	if ( uvs.length > 0 ) data.data.uvs = [ uvs ]; // temporal backward compatibility
	data.data.faces = faces;

	return data;

},
*/
func (g *Geometry) Clone() (*Geometry) {
	return NewGeometry().Copy(g)
}

func (g *Geometry) Copy(source *Geometry) (*Geometry) {
	g.Vertices = make([]*math3d.Vector3, len(source.Vertices))
	g.Faces = make([]*Face3, len(source.Faces))
	g.FaceVertexUvs = make([]([]*math3d.Vector2), 0)

	vertices := source.Vertices

	for _, v := range vertices {
		g.Vertices = append(g.Vertices, v)
	}

	faces := source.Faces;
	for _, face := range faces {
		g.Faces = append(g.Faces, face)
	}

	for i, faceVertexUvs := range source.FaceVertexUvs {
		if g.FaceVertexUvs[ i ] == nil {
			g.FaceVertexUvs[ i ] = make([]([]*math3d.Vector2), 0)
		}
		for _, uvs := range faceVertexUvs {
			uvsCopy := make([]*math3d.Vector2, 0);
			for _, uv := range uvs {
				uvsCopy = append(uvsCopy, uv.Clone() )
			}
			g.FaceVertexUvs[ i ] = append(g.FaceVertexUvs[i], uvsCopy)
		}
	}
	return g
}

func (g *Geometry) Dispose() {
	g.DispatchEvent(NewEvent("dispose"))
}
