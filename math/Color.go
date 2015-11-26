package math
import "math"

type Color struct {
	r, g, b float64
}

func NewDefaultColor() (*Color) {
	return NewColor(1.0, 1.0, 1.0)
}

func NewColor(r, g, b float64) (*Color) {
	return &Color{r, g, b}
}

func (c *Color) SetHex(hex int) (*Color) {
	c.r = float64( hex >> 16 & 255 ) / 255.0
	c.g = float64( hex >> 8 & 255 ) / 255.0
	c.b = float64( hex & 255 ) / 255.0
	return c
}

func (c *Color) SetRGB(r, g, b float64) (*Color) {
	c.r = r
	c.g = g
	c.b = b

	return c
}

func (c *Color) Clone() (*Color) {
	return NewColor( c.r, c.g, c.b )
}

func (c *Color) Copy(color *Color) (*Color) {
	c.r = color.r
	c.g = color.g
	c.b = color.b

	return c
}

func (c *Color) CopyDefaultGammaToLinear(color *Color) (*Color) {
	return c.CopyGammaToLinear(color, 2.0)
}

func (c *Color) CopyGammaToLinear(color *Color, gammaFactor float64) (*Color) {
	c.r = math.Pow( color.r, gammaFactor )
	c.g = math.Pow( color.g, gammaFactor )
	c.b = math.Pow( color.b, gammaFactor )

	return c
}

func (c *Color) CopyDefaultLinearToGamma(color *Color) (*Color) {
	return c.CopyLinearToGamma(color, 2.0)
}

func (c *Color) CopyLinearToGamma(color *Color, gammaFactor float64) (*Color) {
	var safeInverse float64
	if gammaFactor > 0 {
		safeInverse = 1.0 / gammaFactor
	} else {
		safeInverse = 1.0
	}

	c.r = math.Pow( color.r, safeInverse )
	c.g = math.Pow( color.g, safeInverse )
	c.b = math.Pow( color.b, safeInverse )

	return c
}

func (c *Color) ConvertGammaToLinear() (*Color) {

	r := c.r
	g := c.g
	b := c.b

	c.r = r * r
	c.g = g * g
	c.b = b * b

	return c
}

func (c *Color) ConvertLinearToGamma() (*Color) {
	c.r = math.Sqrt( c.r )
	c.g = math.Sqrt( c.g )
	c.b = math.Sqrt( c.b )

	return c
}

func (c *Color) GetHex() int {
	return int(( c.r * 255 ) << 16 ^ ( c.g * 255 ) << 8 ^ ( c.b * 255 ) << 0)
}

/*
	getHexString: function () {

		return ( '000000' + c.getHex().toString( 16 ) ).slice( - 6 );

	},

	getHSL: function ( optionalTarget ) {

		// h,s,l ranges are in 0.0 - 1.0

		var hsl = optionalTarget || { h: 0, s: 0, l: 0 };

		var r = c.r, g = c.g, b = c.b;

		var max = Math.max( r, g, b );
		var min = Math.min( r, g, b );

		var hue, saturation;
		var lightness = ( min + max ) / 2.0;

		if ( min === max ) {

			hue = 0;
			saturation = 0;

		} else {

			var delta = max - min;

			saturation = lightness <= 0.5 ? delta / ( max + min ) : delta / ( 2 - max - min );

			switch ( max ) {

				case r: hue = ( g - b ) / delta + ( g < b ? 6 : 0 ); break;
				case g: hue = ( b - r ) / delta + 2; break;
				case b: hue = ( r - g ) / delta + 4; break;

			}

			hue /= 6;

		}

		hsl.h = hue;
		hsl.s = saturation;
		hsl.l = lightness;

		return hsl;

	},

	getStyle: function () {

		return 'rgb(' + ( ( c.r * 255 ) | 0 ) + ',' + ( ( c.g * 255 ) | 0 ) + ',' + ( ( c.b * 255 ) | 0 ) + ')';

	},

	offsetHSL: function ( h, s, l ) {

		var hsl = c.getHSL();

		hsl.h += h; hsl.s += s; hsl.l += l;

		c.setHSL( hsl.h, hsl.s, hsl.l );

		return c

	},
*/

func (c *Color) Add(color *Color) (*Color) {
	c.r += color.r
	c.g += color.g
	c.b += color.b

	return c
}

func (c *Color) AddColors(color1, color2 *Color) (*Color) {
	c.r = color1.r + color2.r
	c.g = color1.g + color2.g
	c.b = color1.b + color2.b

	return c
}

func (c *Color) AddScalar(s float64) (*Color) {
	c.r += s
	c.g += s
	c.b += s

	return c
}

func (c *Color) Multiply(color *Color) (*Color) {
	c.r *= color.r
	c.g *= color.g
	c.b *= color.b

	return c
}

func (c *Color) MultiplyScalar(s float64) (*Color) {
	c.r *= s
	c.g *= s
	c.b *= s

	return c
}

func (c *Color) Lerp(color *Color, alpha float64) (*Color) {
	c.r += ( color.r - c.r ) * alpha
	c.g += ( color.g - c.g ) * alpha
	c.b += ( color.b - c.b ) * alpha

	return c
}

func (c *Color) Equals(color *Color) bool {
	return ( c.r == color.r ) && ( c.g == color.g ) && ( c.b == color.b )
}

func (c *Color) FromArray(array []float64, offset int) (*Color) {
	c.r = array[ offset ]
	c.g = array[ offset + 1 ]
	c.b = array[ offset + 2 ]

	return c
}

func (c *Color) ToArray(array []float64, offset int) (*Color) {
	array[ offset ] = c.r
	array[ offset + 1 ] = c.g
	array[ offset + 2 ] = c.b

	return array
}


/*
THREE.ColorKeywords = { 'aliceblue': 0xF0F8FF, 'antiquewhite': 0xFAEBD7, 'aqua': 0x00FFFF, 'aquamarine': 0x7FFFD4, 'azure': 0xF0FFFF,
'beige': 0xF5F5DC, 'bisque': 0xFFE4C4, 'black': 0x000000, 'blanchedalmond': 0xFFEBCD, 'blue': 0x0000FF, 'blueviolet': 0x8A2BE2,
'brown': 0xA52A2A, 'burlywood': 0xDEB887, 'cadetblue': 0x5F9EA0, 'chartreuse': 0x7FFF00, 'chocolate': 0xD2691E, 'coral': 0xFF7F50,
'cornflowerblue': 0x6495ED, 'cornsilk': 0xFFF8DC, 'crimson': 0xDC143C, 'cyan': 0x00FFFF, 'darkblue': 0x00008B, 'darkcyan': 0x008B8B,
'darkgoldenrod': 0xB8860B, 'darkgray': 0xA9A9A9, 'darkgreen': 0x006400, 'darkgrey': 0xA9A9A9, 'darkkhaki': 0xBDB76B, 'darkmagenta': 0x8B008B,
'darkolivegreen': 0x556B2F, 'darkorange': 0xFF8C00, 'darkorchid': 0x9932CC, 'darkred': 0x8B0000, 'darksalmon': 0xE9967A, 'darkseagreen': 0x8FBC8F,
'darkslateblue': 0x483D8B, 'darkslategray': 0x2F4F4F, 'darkslategrey': 0x2F4F4F, 'darkturquoise': 0x00CED1, 'darkviolet': 0x9400D3,
'deeppink': 0xFF1493, 'deepskyblue': 0x00BFFF, 'dimgray': 0x696969, 'dimgrey': 0x696969, 'dodgerblue': 0x1E90FF, 'firebrick': 0xB22222,
'floralwhite': 0xFFFAF0, 'forestgreen': 0x228B22, 'fuchsia': 0xFF00FF, 'gainsboro': 0xDCDCDC, 'ghostwhite': 0xF8F8FF, 'gold': 0xFFD700,
'goldenrod': 0xDAA520, 'gray': 0x808080, 'green': 0x008000, 'greenyellow': 0xADFF2F, 'grey': 0x808080, 'honeydew': 0xF0FFF0, 'hotpink': 0xFF69B4,
'indianred': 0xCD5C5C, 'indigo': 0x4B0082, 'ivory': 0xFFFFF0, 'khaki': 0xF0E68C, 'lavender': 0xE6E6FA, 'lavenderblush': 0xFFF0F5, 'lawngreen': 0x7CFC00,
'lemonchiffon': 0xFFFACD, 'lightblue': 0xADD8E6, 'lightcoral': 0xF08080, 'lightcyan': 0xE0FFFF, 'lightgoldenrodyellow': 0xFAFAD2, 'lightgray': 0xD3D3D3,
'lightgreen': 0x90EE90, 'lightgrey': 0xD3D3D3, 'lightpink': 0xFFB6C1, 'lightsalmon': 0xFFA07A, 'lightseagreen': 0x20B2AA, 'lightskyblue': 0x87CEFA,
'lightslategray': 0x778899, 'lightslategrey': 0x778899, 'lightsteelblue': 0xB0C4DE, 'lightyellow': 0xFFFFE0, 'lime': 0x00FF00, 'limegreen': 0x32CD32,
'linen': 0xFAF0E6, 'magenta': 0xFF00FF, 'maroon': 0x800000, 'mediumaquamarine': 0x66CDAA, 'mediumblue': 0x0000CD, 'mediumorchid': 0xBA55D3,
'mediumpurple': 0x9370DB, 'mediumseagreen': 0x3CB371, 'mediumslateblue': 0x7B68EE, 'mediumspringgreen': 0x00FA9A, 'mediumturquoise': 0x48D1CC,
'mediumvioletred': 0xC71585, 'midnightblue': 0x191970, 'mintcream': 0xF5FFFA, 'mistyrose': 0xFFE4E1, 'moccasin': 0xFFE4B5, 'navajowhite': 0xFFDEAD,
'navy': 0x000080, 'oldlace': 0xFDF5E6, 'olive': 0x808000, 'olivedrab': 0x6B8E23, 'orange': 0xFFA500, 'orangered': 0xFF4500, 'orchid': 0xDA70D6,
'palegoldenrod': 0xEEE8AA, 'palegreen': 0x98FB98, 'paleturquoise': 0xAFEEEE, 'palevioletred': 0xDB7093, 'papayawhip': 0xFFEFD5, 'peachpuff': 0xFFDAB9,
'peru': 0xCD853F, 'pink': 0xFFC0CB, 'plum': 0xDDA0DD, 'powderblue': 0xB0E0E6, 'purple': 0x800080, 'red': 0xFF0000, 'rosybrown': 0xBC8F8F,
'royalblue': 0x4169E1, 'saddlebrown': 0x8B4513, 'salmon': 0xFA8072, 'sandybrown': 0xF4A460, 'seagreen': 0x2E8B57, 'seashell': 0xFFF5EE,
'sienna': 0xA0522D, 'silver': 0xC0C0C0, 'skyblue': 0x87CEEB, 'slateblue': 0x6A5ACD, 'slategray': 0x708090, 'slategrey': 0x708090, 'snow': 0xFFFAFA,
'springgreen': 0x00FF7F, 'steelblue': 0x4682B4, 'tan': 0xD2B48C, 'teal': 0x008080, 'thistle': 0xD8BFD8, 'tomato': 0xFF6347, 'turquoise': 0x40E0D0,
'violet': 0xEE82EE, 'wheat': 0xF5DEB3, 'white': 0xFFFFFF, 'whitesmoke': 0xF5F5F5, 'yellow': 0xFFFF00, 'yellowgreen': 0x9ACD32 };
*/