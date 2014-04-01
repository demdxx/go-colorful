// Various ways to generate single random colors
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

// Creates a random dark, "warm" color through a restricted HSV space.
func FastWarmColor() Color {
	return ColorHsv{
		rand.Float64() * 360.0,
		rand.Float64()*0.3 + 0.5,
		rand.Float64()*0.3 + 0.3}.Color()
}

// Creates a random dark, "warm" color through restricted HCL space.
// This is slower than FastWarmColor but will likely give you colors which have
// the same "warmness" if you run it many times.
func WarmColor() (c Color) {
	for c = randomWarm(); !c.IsValid(); c = randomWarm() {
		// DUMMY LOOP BODY ...
	}
	return
}

func randomWarm() Color {
	return ColorHcl{
		rand.Float64() * 360.0,
		rand.Float64()*0.3 + 0.1,
		rand.Float64()*0.3 + 0.2}.Color()
}

// Creates a random bright, "pimpy" color through a restricted HSV space.
func FastHappyColor() Color {
	return ColorHsv{
		rand.Float64() * 360.0,
		rand.Float64()*0.3 + 0.7,
		rand.Float64()*0.3 + 0.6}.Color()
}

// Creates a random bright, "pimpy" color through restricted HCL space.
// This is slower than FastHappyColor but will likely give you colors which
// have the same "brightness" if you run it many times.
func HappyColor() (c Color) {
	for c = randomPimp(); !c.IsValid(); c = randomPimp() {
		// DUMMY LOOP BODY ...
	}
	return
}

func randomPimp() Color {
	return ColorHcl{
		rand.Float64() * 360.0,
		rand.Float64()*0.3 + 0.5,
		rand.Float64()*0.3 + 0.5}.Color()
}
