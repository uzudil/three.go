package main

import (
	three "github.com/uzudil/three.go"
	"github.com/uzudil/three.go/cameras"
	"github.com/uzudil/three.go/scenes"
	"github.com/uzudil/three.go/extras/geometries"
	"github.com/uzudil/three.go/materials"
	"github.com/uzudil/three.go/objects"
)

var camera cameras.Camera
var scene scenes.Scene
var mesh three.Mesh
var renderer three.Renderer

const (
	WIDTH = 800
	HEIGHT = 600
)

func main() {
	camera = cameras.NewPerspectiveCamera( 70.0, float64(WIDTH / HEIGHT), 1.0, 1000.0 )
	camera.Position.Z = 400

	scene = scenes.NewScene()

	geometry := geometries.NewDefaultBoxGeometry( 200, 200, 200 )
	material := materials.NewMeshBasicMaterial(map[string]interface{}{ "color": 0xffffff })

	mesh = objects.NewMesh( geometry, material )
	scene.Add( mesh )

	renderer = three.NewGLRenderer()

	animate()
}

func animate() {
	mesh.Rotation.X += 0.005
	mesh.Rotation.Y += 0.01
	renderer.Render( scene, camera )
}
