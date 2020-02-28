// Code generated by 'goexports gonum.org/v1/gonum/dsp/window'. DO NOT EDIT.

// Copyright ©2019 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build go1.14,!go1.15

package yaegi

import (
	"reflect"

	"gonum.org/v1/gonum/dsp/window"
)

func init() {
	Symbols["gonum.org/v1/gonum/dsp/window"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"BartlettHann":           reflect.ValueOf(window.BartlettHann),
		"BartlettHannComplex":    reflect.ValueOf(window.BartlettHannComplex),
		"Blackman":               reflect.ValueOf(window.Blackman),
		"BlackmanComplex":        reflect.ValueOf(window.BlackmanComplex),
		"BlackmanHarris":         reflect.ValueOf(window.BlackmanHarris),
		"BlackmanHarrisComplex":  reflect.ValueOf(window.BlackmanHarrisComplex),
		"BlackmanNuttall":        reflect.ValueOf(window.BlackmanNuttall),
		"BlackmanNuttallComplex": reflect.ValueOf(window.BlackmanNuttallComplex),
		"FlatTop":                reflect.ValueOf(window.FlatTop),
		"FlatTopComplex":         reflect.ValueOf(window.FlatTopComplex),
		"Hamming":                reflect.ValueOf(window.Hamming),
		"HammingComplex":         reflect.ValueOf(window.HammingComplex),
		"Hann":                   reflect.ValueOf(window.Hann),
		"HannComplex":            reflect.ValueOf(window.HannComplex),
		"Lanczos":                reflect.ValueOf(window.Lanczos),
		"LanczosComplex":         reflect.ValueOf(window.LanczosComplex),
		"NewValues":              reflect.ValueOf(window.NewValues),
		"NewValuesComplex":       reflect.ValueOf(window.NewValuesComplex),
		"Nuttall":                reflect.ValueOf(window.Nuttall),
		"NuttallComplex":         reflect.ValueOf(window.NuttallComplex),
		"Rectangular":            reflect.ValueOf(window.Rectangular),
		"RectangularComplex":     reflect.ValueOf(window.RectangularComplex),
		"Sine":                   reflect.ValueOf(window.Sine),
		"SineComplex":            reflect.ValueOf(window.SineComplex),
		"Triangular":             reflect.ValueOf(window.Triangular),
		"TriangularComplex":      reflect.ValueOf(window.TriangularComplex),

		// type definitions
		"Gaussian":        reflect.ValueOf((*window.Gaussian)(nil)),
		"GaussianComplex": reflect.ValueOf((*window.GaussianComplex)(nil)),
		"Values":          reflect.ValueOf((*window.Values)(nil)),
		"ValuesComplex":   reflect.ValueOf((*window.ValuesComplex)(nil)),
	}
}
