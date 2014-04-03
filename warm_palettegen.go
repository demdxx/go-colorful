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
  "math/rand"
)

// Uses the HSV color space to generate colors with similar S,V but distributed
// evenly along their Hue. This is fast but not always pretty.
// If you've got time to spare, use Lab (the non-fast below).
func FastWarmPalette(colorsCount int) (colors []Color) {
  colors = make([]Color, colorsCount)

  for i := 0; i < colorsCount; i++ {
    colors[i] = ColorHsv{float64(i) * (360.0 / float64(colorsCount)), 0.55 + rand.Float64()*0.2, 0.35 + rand.Float64()*0.2}.Color()
  }
  return
}

func WarmPalette(colorsCount int) ([]Color, error) {
  warmy := func(lab ColorLab) bool {
    c := lab.Hcl()
    return 0.1 <= c.C && c.C <= 0.4 && 0.2 <= lab.L && lab.L <= 0.5
  }
  return SoftPaletteEx(colorsCount, SoftPaletteSettings{warmy, 50, true})
}
