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
  "testing"
  "time"
)

// This is really difficult to test, if you've got a good idea, pull request!

// Check if it returns all valid colors.
func TestColorValidity(t *testing.T) {
  test(t, false)
}

// Checl and print table of colors
func TestColorPrint(t *testing.T) {
  test(t, true)
}

func test(t *testing.T, printIt bool) {
  seed := time.Now().UTC().UnixNano()
  rand.Seed(seed)

  for i := 0; i < 100; i++ {
    c := ""

    if col := WarmColor(); !col.IsValid() {
      t.Errorf("Warm color %v is not valid! Seed was: %v", col, seed)
    } else if printIt {
      c += col.HexString() + " / "
    }

    if col := FastWarmColor(); !col.IsValid() {
      t.Errorf("Fast warm color %v is not valid! Seed was: %v", col, seed)
    } else if printIt {
      c += col.HexString() + " / "
    }

    if col := HappyColor(); !col.IsValid() {
      t.Errorf("Happy color %v is not valid! Seed was: %v", col, seed)
    } else if printIt {
      c += col.HexString() + " / "
    }

    if col := FastHappyColor(); !col.IsValid() {
      t.Errorf("Fast happy color %v is not valid! Seed was: %v", col, seed)
    } else if printIt {
      c += col.HexString()
    }
    if printIt {
      t.Logf("%3d) %v", i+1, c)
    }
  }
}
