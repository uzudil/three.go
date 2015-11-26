package math
import (
	"math"
	"strings"
	"math/rand"
)

func Clamp( value, min, max float64) float64 {
	return math.Max( min, math.Min( max, value ) );
}

// credit: https://gist.github.com/siddontang/1806573b9a8574989ccb
// see http://www.gnu.org/software/libc/manual/html_node/Rounding.html
func Round(x float64) float64 {
	v, frac := math.Modf(x)
	if x > 0.0 {
		if frac > 0.5 || (frac == 0.5 && uint64(v)%2 != 0) {
			v += 1.0
		}
	} else {
		if frac < -0.5 || (frac == -0.5 && uint64(v)%2 != 0) {
			v -= 1.0
		}
	}

	return v
}

func GenerateUUID() string {

	// http://www.broofa.com/Tools/Math.uuid.htm

	chars := strings.Split("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz", "")
	uuid := make([]string, 36)
	rnd := 0
	var r int

	for i := 0; i < 36; i ++ {
		if i == 8 || i == 13 || i == 18 || i == 23 {
			uuid[ i ] = "-"
		} else if i == 14 {
			uuid[ i ] = "4"
		} else {
			if rnd <= 0x02 {
				rnd = 0x2000000 + int(math.Trunc( rand.Float64() * float64(0x1000000) ))
			}
			r = rnd & 0xf
			rnd = rnd >> 4
			if i == 19 {
				uuid[ i ] = chars[( r & 0x3 ) | 0x8]
			} else {
				uuid[ i ] = chars[r]
			}
		}
	}
	return strings.Join(uuid, "")
}

var DegToRad func(float64) float64 = func() (func(float64) float64) {
	degreeToRadiansFactor := math.Pi / 180.0
	return func ( degrees float64) float64 {
		return degrees * degreeToRadiansFactor
	}
}()

var RadToDeg func(float64) float64 = func() (func(float64) float64) {
	radianToDegreesFactor := 180 / math.Pi
	return func ( radians float64) float64 {
		return radians * radianToDegreesFactor
	}
}()

