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
  "sort"
)

func (c Color) Complement() Color {
  r, g, b, a := c.RGBA255()
  return RGBA(^r, ^g, ^b, a)
}

func (c Color) ComplementaryHarm() ColorSlice {
  return ColorSlice([]Color{c, c.Complement()})
}

func (c Color) AnalogousHarm(count, slices uint8) ColorSlice {
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

  h.H = math.Mod((h.H-(part*float64(1+count)*2))+720.0, 360.0)

  for ; count > 0; count-- {
    h.H = math.Mod(h.H+part, 360.0)
    ret[i] = h.Color()
    i++
  }

  sort.Sort(ColorSlice(ret))
  return ColorSlice(ret)
}

func (c Color) MonochromaticHarm(count uint8, mod float64) ColorSlice {
  if count < 2 {
    count = 2
  }

  i := 1
  m := mod / float64(count)
  ret := make([]Color, count)
  ret[0] = c

  r, g, b := (1.0-c.R)*m, (1.0-c.G)*m, (1.0-c.B)*m
  R, G, B := c.R, c.G, c.B

  for count--; count > 0; count-- {
    R += r
    G += g
    B += b
    ret[i] = Color{R: R, G: G, B: B, A: c.A}
    i++
  }

  return ColorSlice(ret)
}

func (c Color) TriadHarm() ColorSlice {
  h := c.Hsv()
  return []Color{
    ColorHsv{H: math.Mod(h.H+120.0, 360.0), S: h.S, V: h.V}.Color(), c,
    ColorHsv{H: math.Mod(h.H-120.0, 360.0), S: h.S, V: h.V}.Color(),
  }
}

func (c Color) SplitComplementaryHarm() ColorSlice {
  h := c.Hsv()
  return ColorSlice([]Color{
    ColorHsv{H: math.Mod(h.H+72.0, 360.0), S: h.S, V: h.V}.Color(), c,
    ColorHsv{H: math.Mod(h.H-216.0, 360.0), S: h.S, V: h.V}.Color(),
  })
}

func (c Color) SquareHarm() ColorSlice {
  h := c.Hsv()
  return ColorSlice([]Color{c,
    ColorHsv{H: math.Mod(h.H+90.0, 360.0), S: h.S, V: h.V}.Color(),
    ColorHsv{H: math.Mod(h.H+180.0, 360.0), S: h.S, V: h.V}.Color(),
    ColorHsv{H: math.Mod(h.H+270.0, 360.0), S: h.S, V: h.V}.Color(),
  })
}

func (c Color) TetradicHarm() ColorSlice {
  h := c.Hsv()
  return ColorSlice([]Color{c,
    ColorHsv{H: math.Mod(h.H+120.0, 360.0), S: h.S, V: h.V}.Color(),
    ColorHsv{H: math.Mod(h.H+180.0, 360.0), S: h.S, V: h.V}.Color(),
    ColorHsv{H: math.Mod(h.H+300.0, 360.0), S: h.S, V: h.V}.Color(),
  })
}
