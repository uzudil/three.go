package main

import three "github.com/uzudil/three.go"

var camera three.Camera
var scene three.Scene
var mesh three.Mesh
var renderer three.Renderer

const (
	WIDTH = 800
	HEIGHT = 600
)

func main() {
	camera = three.NewPerspectiveCamera( 70, WIDTH / HEIGHT, 1, 1000 )
	camera.Position.Z = 400

	scene = three.NewScene()

	geometry := three.NewBoxGeometry( 200, 200, 200 )
	material := three.NewMeshBasicMaterial( { map: texture } )
	mesh = three.NewMesh( geometry, material )
	scene.add( mesh )

	renderer = three.NewGLRenderer()
	// renderer.setPixelRatio( window.devicePixelRatio );
	// renderer.setSize( window.innerWidth, window.innerHeight );

	animate()
}

func animate() {
	mesh.Rotation.X += 0.005
	mesh.Rotation.Y += 0.01
	renderer.Render( scene, camera )
}
