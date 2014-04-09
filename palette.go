package colorful

import (
  "math"
)

func (c Color) PaletteTo(targ Color, count int) []Color {
  if count < 3 {
    count = 3
  }

  inc_R := (c.R - targ.R) / float64(count)
  inc_G := (c.G - targ.G) / float64(count)
  inc_B := (c.B - targ.B) / float64(count)

  colors := make([]Color, count)

  for i := 0; i < count; i++ {
    colors[i] = Color{
      R: c.R + inc_R*float64(i),
      G: c.G + inc_G*float64(i),
      B: c.B + inc_B*float64(i),
      A: 1.0,
    }
  }

  return colors
}

///////////////////////////////////////////////////////////////////////////////
/// HSL colors
///////////////////////////////////////////////////////////////////////////////

func (c ColorHsl) PaletteToLightness(l float64, count int) []ColorHsl {
  if count < 3 {
    count = 3
  }

  offset := (l - c.L) / float64(count)
  offSum := c.L + offset

  colors := make([]ColorHsl, count)

  for i := 0; i < count; i++ {
    colors[i] = ColorHsl{H: c.H, S: c.S, L: offSum}
    offSum += offset
  }

  return colors
}

func (c ColorHsl) PaletteToLightnessFor(min, max float64, count int) []ColorHsl {
  if math.Abs(max-c.L) > math.Abs(c.L-min) {
    return c.PaletteToLightness(max, count)
  }
  return c.PaletteToLightness(min, count)
}

func (c ColorHsl) PaletteToMaxLightness(count int) []ColorHsl {
  return c.PaletteToLightnessFor(0.0, 1.0, count)
}
