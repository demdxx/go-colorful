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
  "math"
)

func (c Color) ComplementaryHarm() []Color {
  return []Color{c, c.Complement()}
}

func (c Color) AnalogousHarm(count, slices uint8) []Color {
  if count < 2 {
    count = 2
  }
  if slices < 2 {
    slices = 2
  }

  i := 1
  h := c.Hsl()
  part := 360.0 / float64(slices)
  ret := make([]Color, count)
  ret[0] = c
  count--

  h.H = math.Remainder((h.H-(part*float64(1+count)*2))+720.0, 360)

  for ; count > 0; count-- {
    h.H = math.Remainder(h.H+part, 360.0)
    ret[i] = h.Color()
    i++
  }

  return ret
}

func (c Color) MonochromaticHarm(count uint8) []Color {
  if count < 2 {
    count = 2
  }

  i := 1
  h := c.Hsv()
  v := h.V
  modification := 1.0 / float64(count)
  ret := make([]Color, count)
  ret[0] = c

  for count--; count > 0; count-- {
    v += math.Remainder(v+modification, 1.0)
    ret[i] = ColorHsv{H: h.H, S: h.S, V: v}.Color()
    i++
  }

  return ret
}

func (c Color) TriadHarm() []Color {
  h := c.Hsv()
  return []Color{
    ColorHsv{H: math.Remainder(h.H+120.0, 360.0), S: h.S, V: h.V}.Color(), c,
    ColorHsv{H: math.Remainder(h.H-120.0, 360.0), S: h.S, V: h.V}.Color(),
  }
}

func (c Color) SplitComplementaryHarm() []Color {
  h := c.Hsv()
  return []Color{
    ColorHsv{H: math.Remainder(h.H+72.0, 360.0), S: h.S, V: h.V}.Color(), c,
    ColorHsv{H: math.Remainder(h.H-216.0, 360.0), S: h.S, V: h.V}.Color(),
  }
}

func (c Color) SquareHarm() []Color {
  h := c.Hsv()
  return []Color{c,
    ColorHsv{H: math.Remainder(h.H+90.0, 360.0), S: h.S, V: h.V}.Color(),
    ColorHsv{H: math.Remainder(h.H+180.0, 360.0), S: h.S, V: h.V}.Color(),
    ColorHsv{H: math.Remainder(h.H+270.0, 360.0), S: h.S, V: h.V}.Color(),
  }
}
