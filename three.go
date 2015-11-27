package main
import (
	"errors"
	"reflect"
	"fmt"
)


// https://developer.mozilla.org/en-US/docs/Web/API/MouseEvent.button

var MOUSE map[string]int = map[string]int{ "LEFT": 0, "MIDDLE": 1, "RIGHT": 2 }

// GL STATE CONSTANTS

var CullFaceNone int = 0
var CullFaceBack int = 1
var CullFaceFront int = 2
var CullFaceFrontBack int = 3

var FrontFaceDirectionCW int = 0
var FrontFaceDirectionCCW int = 1

// SHADOWING TYPES

var BasicShadowMap int = 0
var PCFShadowMap int = 1
var PCFSoftShadowMap int = 2

// MATERIAL CONSTANTS

// side

var FrontSide int = 0
var BackSide int = 1
var DoubleSide int = 2

// shading

var FlatShading int = 1
var SmoothShading int = 2

// colors

var NoColors int = 0
var FaceColors int = 1
var VertexColors int = 2

// blending modes

var NoBlending int = 0
var NormalBlending int = 1
var AdditiveBlending int = 2
var SubtractiveBlending int = 3
var MultiplyBlending int = 4
var CustomBlending int = 5

// custom blending equations
// (numbers start from 100 not to clash with other
// mappings to OpenGL constants defined in Texture.js)

var AddEquation int = 100
var SubtractEquation int = 101
var ReverseSubtractEquation int = 102
var MinEquation int = 103
var MaxEquation int = 104

// custom blending destination factors

var ZeroFactor int = 200
var OneFactor int = 201
var SrcColorFactor int = 202
var OneMinusSrcColorFactor int = 203
var SrcAlphaFactor int = 204
var OneMinusSrcAlphaFactor int = 205
var DstAlphaFactor int = 206
var OneMinusDstAlphaFactor int = 207

// custom blending source factors

//var ZeroFactor int = 200
//var OneFactor int = 201
//var SrcAlphaFactor int = 204
//var OneMinusSrcAlphaFactor int = 205
//var DstAlphaFactor int = 206
//var OneMinusDstAlphaFactor int = 207
var DstColorFactor int = 208
var OneMinusDstColorFactor int = 209
var SrcAlphaSaturateFactor int = 210

// depth modes

var NeverDepth int = 0
var AlwaysDepth int = 1
var LessDepth int = 2
var LessEqualDepth int = 3
var EqualDepth int = 4
var GreaterEqualDepth int = 5
var GreaterDepth int = 6
var NotEqualDepth int = 7


// TEXTURE CONSTANTS

var MultiplyOperation int = 0
var MixOperation int = 1
var AddOperation int = 2

// Mapping modes

var UVMapping int = 300

var CubeReflectionMapping int = 301
var CubeRefractionMapping int = 302

var EquirectangularReflectionMapping int = 303
var EquirectangularRefractionMapping int = 304

var SphericalReflectionMapping int = 305

// Wrapping modes

var RepeatWrapping int = 1000
var ClampToEdgeWrapping int = 1001
var MirroredRepeatWrapping int = 1002

// Filters

var NearestFilter int = 1003
var NearestMipMapNearestFilter int = 1004
var NearestMipMapLinearFilter int = 1005
var LinearFilter int = 1006
var LinearMipMapNearestFilter int = 1007
var LinearMipMapLinearFilter int = 1008

// Data types

var UnsignedByteType int = 1009
var ByteType int = 1010
var ShortType int = 1011
var UnsignedShortType int = 1012
var IntType int = 1013
var UnsignedIntType int = 1014
var FloatType int = 1015
var HalfFloatType int = 1025

// Pixel types

//var UnsignedByteType int = 1009
var UnsignedShort4444Type int = 1016
var UnsignedShort5551Type int = 1017
var UnsignedShort565Type int = 1018

// Pixel formats

var AlphaFormat int = 1019
var RGBFormat int = 1020
var RGBAFormat int = 1021
var LuminanceFormat int = 1022
var LuminanceAlphaFormat int = 1023
// var RGBEFormat handled as var RGBAFormat in shaders
var RGBEFormat int = 1024
var RGBAFormat int = 1024

// DDS / ST3C Compressed texture formats

var RGB_S3TC_DXT1_Format int = 2001
var RGBA_S3TC_DXT1_Format int = 2002
var RGBA_S3TC_DXT3_Format int = 2003
var RGBA_S3TC_DXT5_Format int = 2004


// PVRTC compressed texture formats

var RGB_PVRTC_4BPPV1_Format int = 2100
var RGB_PVRTC_2BPPV1_Format int = 2101
var RGBA_PVRTC_4BPPV1_Format int = 2102
var RGBA_PVRTC_2BPPV1_Format int = 2103

// Loop styles for AnimationAction

var LoopOnce int = 2200
var LoopRepeat int = 2201
var LoopPingPong int = 2202

// credit: http://stackoverflow.com/questions/26744873/converting-map-to-struct
func SetField(obj interface{}, name string, value interface{}) error {
	structValue := reflect.ValueOf(obj).Elem()
	structFieldValue := structValue.FieldByName(name)

	if !structFieldValue.IsValid() {
		return fmt.Errorf("No such field: %s in obj", name)
	}

	if !structFieldValue.CanSet() {
		return fmt.Errorf("Cannot set %s field value", name)
	}

	structFieldType := structFieldValue.Type()
	val := reflect.ValueOf(value)
	if structFieldType != val.Type() {
		invalidTypeError := errors.New("Provided value type didn't match obj field type")
		return invalidTypeError
	}

	structFieldValue.Set(val)
	return nil
}

func GetField(obj interface{}, name string) interface{} {
	r := reflect.ValueOf(obj)
	return reflect.Indirect(r).FieldByName(name).(interface{})
}
