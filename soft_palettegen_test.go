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
	"testing"
)

// This is really difficult to test, if you've got a good idea, pull request!

// Check if it returns all valid and enough colors.
func TestColorCount(t *testing.T) {
	fmt.Printf("Testing up to %v palettes may take a while...\n", 100)
	for i := 0; i < 100; i++ {
		//pal, err := SoftPaletteEx(i, SoftPaletteGenSettings{nil, 50, true})
		pal, err := SoftPalette(i)
		if err != nil {
			t.Errorf("Error: %v", err)
		}

		// Check the color count of the palette
		if len(pal) != i {
			t.Errorf("Requested %v colors but got %v", i, len(pal))
		}

		// Also check whether all colors exist in RGB space.
		for icol, col := range pal {
			if !col.IsValid() {
				t.Errorf("Color %v in palette of %v is invalid: %v", icol, len(pal), col)
			}
		}
	}
	fmt.Println("Done with that, but more tests to run.")
}

// Check if it errors-out on an impossible constraint
func TestImpossibleConstraint(t *testing.T) {
	never := func(lab ColorLab) bool { return false }

	pal, err := SoftPaletteEx(10, SoftPaletteSettings{never, 50, true})
	if err == nil || pal != nil {
		t.Error("Should error-out on impossible constraint!")
	}
}

// Check whether the constraint is respected
func TestConstraint(t *testing.T) {
	octant := func(lab ColorLab) bool { return lab.L <= 0.5 && lab.A <= 0.0 && lab.B <= 0.0 }

	pal, err := SoftPaletteEx(100, SoftPaletteSettings{octant, 50, true})
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	// Check ALL the colors!
	for icol, col := range pal {
		if !col.IsValid() {
			t.Errorf("Color %v in constrained palette is invalid: %v", icol, col)
		}

		lab := col.Lab()
		if lab.L > 0.5 || lab.A > 0.0 || lab.B > 0.0 {
			t.Errorf("Color %v in constrained palette violates the constraint: %v (lab: %v)", icol, col, [3]float64{lab.L, lab.A, lab.B})
		}
	}
}
