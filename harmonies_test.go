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
  "bytes"
  "testing"
)

const cCount = 7

var c Color

func TestInit(t *testing.T) {
  c = FastWarmColor()
}

func TestComplementaryHarm(t *testing.T) {
  printHarm(c.ComplementaryHarm(), t)
}

func TestAnalogousHarm(t *testing.T) {
  printHarm(c.AnalogousHarm(cCount, 30), t)
}

func TestMonochromaticHarm(t *testing.T) {
  printHarm(c.MonochromaticHarm(cCount), t)
}

func TestTriadHarm(t *testing.T) {
  printHarm(c.TriadHarm(), t)
}

func TestSplitComplementaryHarm(t *testing.T) {
  printHarm(c.SplitComplementaryHarm(), t)
}

func TestSquareHarm(t *testing.T) {
  printHarm(c.SquareHarm(), t)
}

func printHarm(harm []Color, t *testing.T) {
  var s bytes.Buffer
  for _, color := range harm {
    if s.Len() > 0 {
      s.WriteString(", ")
    }
    s.WriteString(color.HexString())
  }
  t.Log(s.String())
}
