// This conceptual change port https://github.com/lucasb-eyer/go-colorful
//
// Copyright (c) 2014 Dmitry Ponomarev
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights to use,
// copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the
// Software, and to permit persons to whom the Software is furnished to do so, subject
// to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all copies
//  or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED,
// INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR
// PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
// DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package colorful

import (
  "fmt"
  "math"
)

type Color struct {
  R, G, B, A float64
}

///////////////////////////////////////////////////////////////////////////////
/// Hex
///////////////////////////////////////////////////////////////////////////////

// Hex returns the hex "html" representation of the color, as in #ff0080.
func (c Color) HexString() string {
  // Add 0.5 for rounding
  if 1.0 == c.A {
    return fmt.Sprintf("#%02x%02x%02x",
      uint8(c.R*255.0+0.5), uint8(c.G*255.0+0.5), uint8(c.B*255.0+0.5))
  }
  return fmt.Sprintf("#%02x%02x%02x%02x",
    uint8(c.R*255.0+0.5), uint8(c.G*255.0+0.5), uint8(c.B*255.0+0.5),
    uint8(c.A*255.0+0.5))
}

// Hex parses a "html" hex color-string, either in the 3 "#f0c" or 6 "#ff1034" or 4 "#ffoc" or 8 "#ffff1034" digits form.
func Hex(scol string) (Color, error) {
  var format string
  var count int = 3

  switch len(scol) {
  case 4:
    format = "#%x%x%x"
  case 5:
    format = "#%x%x%x%x"
    count = 4
  case 7:
    format = "#%02x%02x%02x"
  case 9:
    format = "#%02x%02x%02x%02x"
    count = 4
  default:
    return Color{}, fmt.Errorf("color: %v is not a hex-color", scol)
  }

  var r, g, b, a uint8
  var n int
  var err error

  if 4 == count {
    n, err = fmt.Sscanf(scol, format, &r, &g, &b, &a)
  } else {
    n, err = fmt.Sscanf(scol, format, &r, &g, &b)
  }
  if err != nil {
    return Color{}, err
  }
  if n != count {
    return Color{}, fmt.Errorf("color: %v is not a hex-color", scol)
  }
  if count != 4 {
    a = 255
  }

  return Color{float64(r) / 255.0, float64(g) / 255.0, float64(b) / 255.0, float64(a) / 255.0}, nil
}

///////////////////////////////////////////////////////////////////////////////
/// RGB & RGBA
///////////////////////////////////////////////////////////////////////////////

func (c Color) RGBString() string {
  return fmt.Sprintf("%d, %d, %d",
    uint8(c.R*255.0+0.5), uint8(c.G*255.0+0.5), uint8(c.B*255.0+0.5))
}

func (c Color) RGBAString() string {
  return fmt.Sprintf("%d, %d, %d, %f",
    uint8(c.R*255.0+0.5), uint8(c.G*255.0+0.5), uint8(c.B*255.0+0.5), c.A)
}

func (c Color) String() string {
  return c.RGBAString()
}

func RGB(r, g, b uint8) Color {
  return Color{float64(r) / 255.0, float64(g) / 255.0, float64(b) / 255.0, 1.0}
}

func RGBA(r, g, b, a uint8) Color {
  return Color{float64(r) / 255.0, float64(g) / 255.0, float64(b) / 255.0, float64(a) / 255.0}
}

// Might come in handy sometimes to reduce boilerplate code.
func (col Color) RGB255() (r, g, b uint8) {
  r = uint8(col.R * 255.0)
  g = uint8(col.G * 255.0)
  b = uint8(col.B * 255.0)
  return
}

func (col Color) RGBA255() (r, g, b, a uint8) {
  r = uint8(col.R * 255.0)
  g = uint8(col.G * 255.0)
  b = uint8(col.B * 255.0)
  a = uint8(col.A * 255.0)
  return
}

// DistanceRgb computes the distance between two colors in RGB space.
// This is not a good measure! Rather do it in Lab space.
func (c1 Color) DistanceRgb(c2 Color) float64 {
  return math.Sqrt(sq(c1.R-c2.R) + sq(c1.G-c2.G) + sq(c1.B-c2.B))
}

// Check for equality between colors within the tolerance Delta (1/255).
func (c1 Color) AlmostEqualRgb(c2 Color) bool {
  return math.Abs(c1.R-c2.R)+
    math.Abs(c1.G-c2.G)+
    math.Abs(c1.B-c2.B) < 3.0*Delta
}

// Checks whether the color exists in RGB space, i.e. all values are in [0..1]
func (c Color) IsValid() bool {
  return 0.0 <= c.R && c.R <= 1.0 &&
    0.0 <= c.G && c.G <= 1.0 &&
    0.0 <= c.B && c.B <= 1.0 &&
    0.0 <= c.A && c.A <= 1.0
}

// Returns Clamps the color into valid range, clamping each value to [0..1]
// If the color is valid already, this is a no-op.
func (c Color) Clamped() Color {
  return Color{clamp01(c.R), clamp01(c.G), clamp01(c.B), 1.0}
}

// You don't really want to use this, do you? Go for BlendLab, BlendLuv or BlendHcl.
func (c1 Color) BlendRgb(c2 Color, t float64) Color {
  return Color{c1.R + t*(c2.R-c1.R), c1.G + t*(c2.G-c1.G), c1.B + t*(c2.B-c1.B), 1.0}
}

func (c Color) IsEqual(col Color) bool {
  return col.R == c.R && col.G == c.G && col.B == c.B
}

func (c Color) IsEqualStrict(col Color) bool {
  return col.R == c.R && col.G == c.G && col.B == c.B && col.A == c.A
}

func (c Color) IsEqualRGB(r, g, b float64) bool {
  return c.IsEqual(Color{r, g, b, 0.0})
}

func (c Color) IsEqualRGBA(r, g, b, a float64) bool {
  return c.IsEqual(Color{r, g, b, a})
}

func (c Color) IsWebSafe() bool {
  var x, y uint8
  r, g, b := c.RGB255()

  // Red
  x, y = (r>>4)&0x0f, r&0x0f
  if 0 != x%3 || x != y {
    return false
  }

  // Green
  x, y = (g>>4)&0x0f, g&0x0f
  if 0 != x%3 || x != y {
    return false
  }

  // Blue
  x, y = (b>>4)&0x0f, b&0x0f
  if 0 != x%3 || x != y {
    return false
  }

  return true
}

///////////////////////////////////////////////////////////////////////////////
/// CMYK
///////////////////////////////////////////////////////////////////////////////

type ColorCMYK struct {
  C, M, Y, K float64
}

func (c Color) CMYK() ColorCMYK {
  // BLACK
  if c.IsEqualRGB(0.0, 0.0, 0.0) {
    return ColorCMYK{0.0, 0.0, 0.0, 1.0}
  }

  computedC := 1.0 - c.R
  computedM := 1.0 - c.G
  computedY := 1.0 - c.B

  minCMY := math.Min(computedC, math.Min(computedM, computedY))
  computedC = (computedC - minCMY) / (1.0 - minCMY)
  computedM = (computedM - minCMY) / (1.0 - minCMY)
  computedY = (computedY - minCMY) / (1.0 - minCMY)

  return ColorCMYK{computedC, computedM, computedY, minCMY}
}

func (c ColorCMYK) RGB() Color {
  r := 255.0 - round(2.55*(c.C+c.K), 0)
  g := 255.0 - round(2.55*(c.M+c.K), 0)
  b := 255.0 - round(2.55*(c.Y+c.K), 0)

  if r < 0 {
    r = 0
  }
  if g < 0 {
    g = 0
  }
  if b < 0 {
    b = 0
  }

  return Color{r, g, b, 1.0}
}

func (c ColorCMYK) String() string {
  return fmt.Sprintf("%f, %f, %f, %f", c.C, c.M, c.Y, c.K)
}

///////////////////////////////////////////////////////////////////////////////
/// HSV
///////////////////////////////////////////////////////////////////////////////
// From http://en.wikipedia.org/wiki/HSL_and_HSV
// Note that h is in [0..360] and s,v in [0..1]

type ColorHsv struct {
  H, S, V float64
}

// Hsv returns the Hue [0..360], Saturation and Value [0..1] of the color.
func (col Color) Hsv() (hsv ColorHsv) {
  min := math.Min(math.Min(col.R, col.G), col.B)
  hsv.V = math.Max(math.Max(col.R, col.G), col.B)
  C := hsv.V - min

  hsv.S = 0.0
  if hsv.V != 0.0 {
    hsv.S = C / hsv.V
  }

  hsv.H = 0.0 // We use 0 instead of undefined as in wp.
  if min != hsv.V {
    if hsv.V == col.R {
      hsv.H = math.Mod((col.G-col.B)/C, 6.0)
    }
    if hsv.V == col.G {
      hsv.H = (col.B-col.R)/C + 2.0
    }
    if hsv.V == col.B {
      hsv.H = (col.R-col.G)/C + 4.0
    }
    hsv.H *= 60.0
    if hsv.H < 0.0 {
      hsv.H += 360.0
    }
  }
  return
}

// Hsv creates a new Color given a Hue in [0..360], a Saturation and a Value in [0..1]
func (c ColorHsv) Color() Color {
  Hp := c.H / 60.0
  C := c.V * c.S
  X := C * (1.0 - math.Abs(math.Mod(Hp, 2.0)-1.0))

  m := c.V - C
  r, g, b := 0.0, 0.0, 0.0

  switch {
  case 0.0 <= Hp && Hp < 1.0:
    r = C
    g = X
  case 1.0 <= Hp && Hp < 2.0:
    r = X
    g = C
  case 2.0 <= Hp && Hp < 3.0:
    g = C
    b = X
  case 3.0 <= Hp && Hp < 4.0:
    g = X
    b = C
  case 4.0 <= Hp && Hp < 5.0:
    r = X
    b = C
  case 5.0 <= Hp && Hp < 6.0:
    r = C
    b = X
  }

  return Color{m + r, m + g, m + b, 1.0}
}

// You don't really want to use this, do you? Go for BlendLab, BlendLuv or BlendHcl.
func (c1 Color) BlendHsv(c2 Color, t float64) Color {
  h1 := c1.Hsv()
  h2 := c2.Hsv()

  // We know that h are both in [0..360]
  var H float64
  if math.Abs(h2.H-h1.H) <= 180.0 {
    // Won't wrap
    H = h1.H + t*(h2.H-h1.H)
  } else if h1.H < h2.H {
    // Will wrap
    H = math.Mod(h1.H+360.0+t*(h2.H-h1.H-360.0), 360.0)
  } else {
    // Will wrap
    H = math.Mod(h2.H+360.0+t*(h1.H-h2.H-360.0), 360.0)
  }

  return ColorHsv{H, h1.S + t*(h2.S-h1.S), h1.V + t*(h2.V-h1.V)}.Color()
}

///////////////////////////////////////////////////////////////////////////////
/// Hsl
///////////////////////////////////////////////////////////////////////////////

type ColorHsl struct {
  H, S, L float64
}

func (c Color) Hsl() ColorHsl {

  max := math.Max(c.R, math.Max(c.G, c.B))
  min := math.Min(c.R, math.Min(c.G, c.B))
  var h, s, l float64 = 0.0, 0.0, (max + min) / 2

  if max != min {
    d := max - min
    if l > 0.5 {
      s = d / (2.0 - max - min)
    } else {
      s = d / (max + min)
    }
    if max == c.R {
      if c.G < c.B {
        h = (c.G-c.B)/d + 6.0
      } else {
        h = (c.G - c.B) / d
      }
    } else if max == c.G {
      h = (c.B-c.R)/d + 2.0
    } else if max == c.B {
      h = (c.R-c.G)/d + 4.0
    }
    h /= 6.0
  }

  return ColorHsl{H: h * 360.0, S: s, L: l}
}

func (c ColorHsl) Color() Color {
  r, g, b := c.L, c.L, c.L
  if c.S != 0 {
    var q float64 = 0.0
    if c.L < 0.5 {
      q = c.L * (1.0 + c.S)
    } else {
      q = c.L + c.S - c.L*c.S
    }
    p := 2.0*c.L - q
    h := c.H / 360.0
    r = hue2rgb(p, q, h+1.0/3.0)
    g = hue2rgb(p, q, h)
    b = hue2rgb(p, q, h-1.0/3.0)
  }

  return RGB(uint8(r*255), uint8(g*255), uint8(b*255))
}

///////////////////////////////////////////////////////////////////////////////
/// Linear
///////////////////////////////////////////////////////////////////////////////
// http://www.sjbrown.co.uk/2004/05/14/gamma-correct-rendering/
// http://www.brucelindbloom.com/Eqn_RGB_to_XYZ.html

func linearize(v float64) float64 {
  if v <= 0.04045 {
    return v / 12.92
  }
  return math.Pow((v+0.055)/1.055, 2.4)
}

// LinearRgb converts the color into the linear RGB space (see http://www.sjbrown.co.uk/2004/05/14/gamma-correct-rendering/).
func (c Color) LinearRgb() Color {
  return Color{linearize(c.R), linearize(c.G), linearize(c.B), c.A}
}

// FastLinearRgb is much faster than and almost as accurate as LinearRgb.
func (c Color) FastLinearRgb() Color {
  return Color{math.Pow(c.R, 2.2), math.Pow(c.G, 2.2), math.Pow(c.B, 2.2), c.A}
}

func delinearize(v float64) float64 {
  if v <= 0.0031308 {
    return 12.92 * v
  }
  return 1.055*math.Pow(v, 1.0/2.4) - 0.055
}

// LinearRgb creates an sRGB color out of the given linear RGB color (see http://www.sjbrown.co.uk/2004/05/14/gamma-correct-rendering/).
func LinearRgb(r, g, b float64) Color {
  return Color{delinearize(r), delinearize(g), delinearize(b), 1.0}
}

func (c Color) DelinearRgb() Color {
  return Color{delinearize(c.R), delinearize(c.G), delinearize(c.B), 1.0}
}

// FastLinearRgb is much faster than and almost as accurate as LinearRgb.
func FastLinearRgb(r, g, b float64) Color {
  return Color{math.Pow(r, 1.0/2.2), math.Pow(g, 1.0/2.2), math.Pow(b, 1.0/2.2), 1.0}
}

///////////////////////////////////////////////////////////////////////////////
/// XYZ
///////////////////////////////////////////////////////////////////////////////
// http://www.sjbrown.co.uk/2004/05/14/gamma-correct-rendering/

type ColorXyz struct {
  X, Y, Z float64
}

// XyzToLinearRgb converts from CIE XYZ-space to Linear RGB space.
func (c ColorXyz) LinearRgb() (r Color) {
  r.R = 3.2404542*c.X - 1.5371385*c.Y - 0.4985314*c.Z
  r.G = -0.9692660*c.X + 1.8760108*c.Y + 0.0415560*c.Z
  r.B = 0.0556434*c.X - 0.2040259*c.Y + 1.0572252*c.Z
  return
}

func (c Color) LinearRgbToXyz() (r ColorXyz) {
  r.X = 0.4124564*c.R + 0.3575761*c.G + 0.1804375*c.B
  r.Y = 0.2126729*c.R + 0.7151522*c.G + 0.0721750*c.B
  r.Z = 0.0193339*c.R + 0.1191920*c.G + 0.9503041*c.B
  return
}

func (c Color) Xyz() ColorXyz {
  return c.LinearRgb().LinearRgbToXyz()
}

func (c ColorXyz) Color() Color {
  return c.LinearRgb().DelinearRgb()
}

///////////////////////////////////////////////////////////////////////////////
/// xyY
///////////////////////////////////////////////////////////////////////////////
// http://www.brucelindbloom.com/Eqn_XYZ_to_xyY.html

type ColorXyy struct {
  X, Y, Yout float64
}

// Well, the name is bad, since it's xyY but Golang needs me to start with a
// capital letter to make the method public.
func XyzToXyy(X, Y, Z float64) (x, y, Yout float64) {
  return XyzToXyyWhiteRef(X, Y, Z, D65)
}

func XyzToXyyWhiteRef(X, Y, Z float64, wref [3]float64) (x, y, Yout float64) {
  Yout = Y
  N := X + Y + Z
  if math.Abs(N) < 1e-14 {
    // When we have black, Bruce Lindbloom recommends to use
    // the reference white's chromacity for x and y.
    x = wref[0] / (wref[0] + wref[1] + wref[2])
    y = wref[1] / (wref[0] + wref[1] + wref[2])
  } else {
    x = X / N
    y = Y / N
  }
  return
}

func XyyToXyz(x, y, Y float64) (X, Yout, Z float64) {
  Yout = Y

  if -1e-14 < y && y < 1e-14 {
    X = 0.0
    Z = 0.0
  } else {
    X = Y / y * x
    Z = Y / y * (1.0 - x - y)
  }

  return
}

// Converts the given color to CIE xyY space using D65 as reference white.
// (Note that the reference white is only used for black input.)
// x, y and Y are in [0..1]
func (c Color) Xyy() ColorXyy {
  xyz := c.Xyz()
  x, y, Y := XyzToXyy(xyz.X, xyz.Y, xyz.Z)
  return ColorXyy{x, y, Y}
}

// Converts the given color to CIE xyY space, taking into account
// a given reference white. (i.e. the monitor's white)
// (Note that the reference white is only used for black input.)
// x, y and Y are in [0..1]
func (c Color) XyyWhiteRef(wref [3]float64) ColorXyy {
  xyz := c.Xyz()
  x, y, Yout := XyzToXyyWhiteRef(xyz.X, xyz.Y, xyz.Z, wref)
  return ColorXyy{x, y, Yout}
}

// Conver Xyy to Xyz color
func (c ColorXyy) Xyz() ColorXyz {
  x, y, z := XyyToXyz(c.X, c.Y, c.Yout)
  return ColorXyz{x, y, z}
}

// Generates a color by using data given in CIE xyY space.
// x, y and Y are in [0..1]
func (c ColorXyy) Color() Color {
  return c.Xyz().Color()
}

///////////////////////////////////////////////////////////////////////////////
/// L*a*b*
///////////////////////////////////////////////////////////////////////////////
// http://en.wikipedia.org/wiki/Lab_color_space#CIELAB-CIEXYZ_conversions
// For L*a*b*, we need to L*a*b*<->XYZ->RGB and the first one is device dependent.

type ColorLab struct {
  L, A, B float64
}

func lab_f(t float64) float64 {
  if t > 6.0/29.0*6.0/29.0*6.0/29.0 {
    return math.Cbrt(t)
  }
  return t/3.0*29.0/6.0*29.0/6.0 + 4.0/29.0
}

func (c ColorXyz) Lab() ColorLab {
  // Use D65 white as reference point by default.
  // http://www.fredmiranda.com/forum/topic/1035332
  // http://en.wikipedia.org/wiki/Standard_illuminant
  return c.LabWhiteRef(D65)
}

func (c ColorXyz) LabWhiteRef(wref [3]float64) (l ColorLab) {
  fy := lab_f(c.Y / wref[1])
  l.L = 1.16*fy - 0.16
  l.A = 5.0 * (lab_f(c.X/wref[0]) - fy)
  l.B = 2.0 * (fy - lab_f(c.Z/wref[2]))
  return
}

func lab_finv(t float64) float64 {
  if t > 6.0/29.0 {
    return t * t * t
  }
  return 3.0 * 6.0 / 29.0 * 6.0 / 29.0 * (t - 4.0/29.0)
}

func (c ColorLab) Xyz() ColorXyz {
  // D65 white (see above).
  return c.XyzWhiteRef(D65)
}

func (c ColorLab) XyzWhiteRef(wref [3]float64) (xyz ColorXyz) {
  l2 := (c.L + 0.16) / 1.16
  xyz.X = wref[0] * lab_finv(l2+c.A/5.0)
  xyz.Y = wref[1] * lab_finv(l2)
  xyz.Z = wref[2] * lab_finv(l2-c.B/2.0)
  return
}

func (c ColorLab) LabWhiteRef(wref [3]float64) ColorLab {
  return c.Xyz().LabWhiteRef(wref)
}

// Converts the given color to CIE L*a*b* space using D65 as reference white.
func (c Color) Lab() ColorLab {
  return c.Xyz().Lab()
}

// Converts the given color to CIE L*a*b* space, taking into account
// a given reference white. (i.e. the monitor's white)
func (c Color) LabWhiteRef(wref [3]float64) ColorLab {
  return c.Xyz().LabWhiteRef(wref)
}

// Generates a color by using data given in CIE L*a*b* space using D65 as reference white.
func (c ColorLab) Color() Color {
  return c.Xyz().Color()
}

// Generates a color by using data given in CIE L*a*b* space, taking
// into account a given reference white. (i.e. the monitor's white)
func (c ColorLab) WhiteRef(wref [3]float64) Color {
  return c.XyzWhiteRef(wref).Color()
}

// DistanceLab is a good measure of visual similarity between two colors!
// A result of 0 would mean identical colors, while a result of 1 or higher
// means the colors differ a lot.
func (c1 Color) DistanceLab(c2 Color) float64 {
  l1 := c1.Lab()
  l2 := c2.Lab()
  return math.Sqrt(sq(l1.L-l2.L) + sq(l1.A-l2.A) + sq(l1.B-l2.B))
}

// BlendLab blends two colors in the L*a*b* color-space, which should result in a smoother blend.
// t == 0 results in c1, t == 1 results in c2
func (c1 Color) BlendLab(c2 Color, t float64) Color {
  l1 := c1.Lab()
  l2 := c2.Lab()
  return ColorLab{
    l1.L + t*(l2.L-l1.L),
    l1.A + t*(l2.A-l1.A),
    l1.B + t*(l2.B-l1.B)}.Color()
}

const LAB_DELTA = 1e-6

func (lab1 ColorLab) Eq(lab2 ColorLab) bool {
  return math.Abs(lab1.L-lab2.L) < LAB_DELTA &&
    math.Abs(lab1.A-lab2.A) < LAB_DELTA &&
    math.Abs(lab1.B-lab2.B) < LAB_DELTA
}

// That's faster than using colorful's DistanceLab since we would have to
// convert back and forth for that. Here is no conversion.
func (lab1 ColorLab) Dist(lab2 ColorLab) float64 {
  return math.Sqrt(sq(lab1.L-lab2.L) + sq(lab1.A-lab2.A) + sq(lab1.B-lab2.B))
}

///////////////////////////////////////////////////////////////////////////////
/// L*u*v*
///////////////////////////////////////////////////////////////////////////////
// http://en.wikipedia.org/wiki/CIELUV#XYZ_.E2.86.92_CIELUV_and_CIELUV_.E2.86.92_XYZ_conversions
// For L*u*v*, we need to L*u*v*<->XYZ<->RGB and the first one is device dependent.

type ColorLuv struct {
  L, U, V float64
}

func (c ColorXyz) Luv() ColorLuv {
  // Use D65 white as reference point by default.
  // http://www.fredmiranda.com/forum/topic/1035332
  // http://en.wikipedia.org/wiki/Standard_illuminant
  return c.LuvWhiteRef(D65)
}

func (c ColorXyz) LuvWhiteRef(wref [3]float64) (l ColorLuv) {
  if c.Y/wref[1] <= 6.0/29.0*6.0/29.0*6.0/29.0 {
    l.L = c.Y / wref[1] * 29.0 / 3.0 * 29.0 / 3.0 * 29.0 / 3.0
  } else {
    l.L = 1.16*math.Cbrt(c.Y/wref[1]) - 0.16
  }
  ubis, vbis := xyz_to_uv(c.X, c.Y, c.Z)
  un, vn := xyz_to_uv(wref[0], wref[1], wref[2])
  l.U = 13.0 * l.L * (ubis - un)
  l.V = 13.0 * l.L * (vbis - vn)
  return
}

// For this part, we do as R's graphics.hcl does, not as wikipedia does.
// Or is it the same?
func xyz_to_uv(x, y, z float64) (u, v float64) {
  denom := x + 15.0*y + 3.0*z
  if denom == 0.0 {
    u, v = 0.0, 0.0
  } else {
    u = 4.0 * x / denom
    v = 9.0 * y / denom
  }
  return
}

func (c ColorLuv) Xyz() ColorXyz {
  // D65 white (see above).
  return c.XyzWhiteRef(D65)
}

func (c ColorLuv) XyzWhiteRef(wref [3]float64) (xyz ColorXyz) {
  //xyz.Y = wref[1] * lab_finv((c.L + 0.16) / 1.16)
  if c.L <= 0.08 {
    xyz.Y = wref[1] * c.L * 100.0 * 3.0 / 29.0 * 3.0 / 29.0 * 3.0 / 29.0
  } else {
    xyz.Y = wref[1] * cub((c.L+0.16)/1.16)
  }
  un, vn := xyz_to_uv(wref[0], wref[1], wref[2])
  if c.L != 0.0 {
    ubis := c.U/(13.0*c.L) + un
    vbis := c.V/(13.0*c.L) + vn
    xyz.X = xyz.Y * 9.0 * ubis / (4.0 * vbis)
    xyz.Z = xyz.Y * (12.0 - 3.0*ubis - 20.0*vbis) / (4.0 * vbis)
  } else {
    xyz.X, xyz.Y = 0.0, 0.0
  }
  return
}

// Converts the given color to CIE L*u*v* space using D65 as reference white.
// L* is in [0..1] and both u* and v* are in about [-1..1]
func (c Color) Luv() ColorLuv {
  return c.Xyz().Luv()
}

// Converts the given color to CIE L*u*v* space, taking into account
// a given reference white. (i.e. the monitor's white)
// L* is in [0..1] and both u* and v* are in about [-1..1]
func (c Color) LuvWhiteRef(wref [3]float64) ColorLuv {
  return c.Xyz().LuvWhiteRef(wref)
}

// Generates a color by using data given in CIE L*u*v* space using D65 as reference white.
// L* is in [0..1] and both u* and v* are in about [-1..1]
func (c ColorLuv) Color() Color {
  return c.Xyz().Color()
}

// Generates a color by using data given in CIE L*u*v* space, taking
// into account a given reference white. (i.e. the monitor's white)
// L* is in [0..1] and both u* and v* are in about [-1..1]

func (c ColorLuv) WhiteRef(wref [3]float64) ColorLuv {
  return c.XyzWhiteRef(wref).Luv()
}

// DistanceLuv is a good measure of visual similarity between two colors!
// A result of 0 would mean identical colors, while a result of 1 or higher
// means the colors differ a lot.
func (c1 Color) DistanceLuv(c2 Color) float64 {
  l1 := c1.Luv()
  l2 := c2.Luv()
  return math.Sqrt(sq(l1.L-l2.L) + sq(l1.U-l2.U) + sq(l1.V-l2.V))
}

// BlendLuv blends two colors in the CIE-L*u*v* color-space, which should result in a smoother blend.
// t == 0 results in c1, t == 1 results in c2
func (c1 Color) BlendLuv(c2 Color, t float64) Color {
  l1 := c1.Luv()
  l2 := c2.Luv()
  return ColorLuv{
    l1.L + t*(l2.L-l1.L),
    l1.U + t*(l2.U-l1.U),
    l1.V + t*(l2.V-l1.V)}.Color()
}

///////////////////////////////////////////////////////////////////////////////
/// HCL
///////////////////////////////////////////////////////////////////////////////
// HCL is nothing else than L*a*b* in cylindrical coordinates!
// (this was wrong on English wikipedia, I fixed it, let's hope the fix stays.)
// But it is widely popular since it is a "correct HSV"
// http://www.hunterlab.com/appnotes/an09_96a.pdf

type ColorHcl struct {
  H, C, L float64
}

// Converts the given color to HCL space using D65 as reference white.
// H values are in [0..360], C and L values are in [0..1] although C can overshoot 1.0
func (col Color) Hcl() ColorHcl {
  return col.HclWhiteRef(D65)
}

func (c ColorLab) Hcl() ColorHcl {
  return LabToHcl(c.L, c.A, c.B)
}

func (c ColorHcl) Color() Color {
  return c.WhiteRef(D65)
}

func LabToHcl(L, a, b float64) (hcl ColorHcl) {
  // Oops, floating point workaround necessary if a ~= b and both are very small (i.e. almost zero).
  if math.Abs(b-a) > 1e-4 && math.Abs(a) > 1e-4 {
    hcl.H = math.Mod(57.29577951308232087721*math.Atan2(b, a)+360.0, 360.0) // Rad2Deg
  } else {
    hcl.H = 0.0
  }
  hcl.C = math.Sqrt(sq(a) + sq(b))
  hcl.L = L
  return
}

// Converts the given color to HCL space, taking into account
// a given reference white. (i.e. the monitor's white)
// H values are in [0..360], C and L values are in [0..1]
func (col Color) HclWhiteRef(wref [3]float64) ColorHcl {
  return col.LabWhiteRef(wref).Hcl()
}

func (c ColorHcl) Lab() ColorLab {
  H := 0.01745329251994329576 * c.H // Deg2Rad
  a := c.C * math.Cos(H)
  b := c.C * math.Sin(H)
  return ColorLab{c.L, a, b}
}

// Generates a color by using data given in HCL space, taking
// into account a given reference white. (i.e. the monitor's white)
// H values are in [0..360], C and L values are in [0..1]
func (c ColorHcl) WhiteRef(wref [3]float64) Color {
  return c.Lab().WhiteRef(wref)
}

// BlendHcl blends two colors in the CIE-L*C*h° color-space, which should result in a smoother blend.
// t == 0 results in c1, t == 1 results in c2
func (col1 Color) BlendHcl(col2 Color, t float64) Color {
  hcl1 := col1.Hcl()
  hcl2 := col2.Hcl()

  // We know that h are both in [0..360]
  var H float64
  if math.Abs(hcl2.H-hcl1.H) <= 180.0 {
    // Won't wrap
    H = hcl1.H + t*(hcl2.H-hcl1.H)
  } else if hcl1.H < hcl2.H {
    // Will wrap
    H = math.Mod(hcl1.H+360.0+t*(hcl2.H-hcl1.H-360.0), 360.0)
  } else {
    // Will wrap
    H = math.Mod(hcl2.H+360.0+t*(hcl1.H-hcl2.H-360.0), 360.0)
  }

  return ColorHcl{H, hcl1.C + t*(hcl2.C-hcl1.C), hcl1.L + t*(hcl2.L-hcl1.L)}.Color()
}

///////////////////////////////////////////////////////////////////////////////
/// Helpers RGBA color
///////////////////////////////////////////////////////////////////////////////

// This is the tolerance used when comparing colors using AlmostEqualRgb.
const Delta = 1.0 / 255.0

// This is the default reference white point.
var D65 = [3]float64{0.95047, 1.00000, 1.08883}

// And another one.
var D50 = [3]float64{0.96422, 1.00000, 0.82521}

func round(val float64, prec int) float64 {

  var rounder float64
  intermed := val * math.Pow(10, float64(prec))

  if val >= 0.5 {
    rounder = math.Ceil(intermed)
  } else {
    rounder = math.Floor(intermed)
  }
  return rounder / math.Pow(10, float64(prec))
}

func clamp01(v float64) float64 {
  return math.Max(0.0, math.Min(v, 1.0))
}

func sq(v float64) float64 {
  return v * v
}

func cub(v float64) float64 {
  return v * v * v
}

func hue2rgb(p, q, t float64) float64 {
  if t < 0.0 {
    t += 1.0
  }
  if t > 1.0 {
    t -= 1.0
  }
  if t < 1.0/6.0 {
    return p + (q-p)*6.0*t
  }
  if t < 1.0/2.0 {
    return q
  }
  if t < 2.0/3.0 {
    return p + (q-p)*(2.0/3.0-t)*6.0
  }
  return p
}
