// Code generated by 'goexports gonum.org/v1/gonum/blas/blas32'. DO NOT EDIT.

// Copyright ©2019 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build go1.14,!go1.15

package yaegi

import (
	"reflect"

	"gonum.org/v1/gonum/blas/blas32"
)

func init() {
	Symbols["gonum.org/v1/gonum/blas/blas32"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Asum":           reflect.ValueOf(blas32.Asum),
		"Axpy":           reflect.ValueOf(blas32.Axpy),
		"Copy":           reflect.ValueOf(blas32.Copy),
		"DDot":           reflect.ValueOf(blas32.DDot),
		"Dot":            reflect.ValueOf(blas32.Dot),
		"Gbmv":           reflect.ValueOf(blas32.Gbmv),
		"Gemm":           reflect.ValueOf(blas32.Gemm),
		"Gemv":           reflect.ValueOf(blas32.Gemv),
		"Ger":            reflect.ValueOf(blas32.Ger),
		"Iamax":          reflect.ValueOf(blas32.Iamax),
		"Implementation": reflect.ValueOf(blas32.Implementation),
		"Nrm2":           reflect.ValueOf(blas32.Nrm2),
		"Rot":            reflect.ValueOf(blas32.Rot),
		"Rotg":           reflect.ValueOf(blas32.Rotg),
		"Rotm":           reflect.ValueOf(blas32.Rotm),
		"Rotmg":          reflect.ValueOf(blas32.Rotmg),
		"SDDot":          reflect.ValueOf(blas32.SDDot),
		"Sbmv":           reflect.ValueOf(blas32.Sbmv),
		"Scal":           reflect.ValueOf(blas32.Scal),
		"Spmv":           reflect.ValueOf(blas32.Spmv),
		"Spr":            reflect.ValueOf(blas32.Spr),
		"Spr2":           reflect.ValueOf(blas32.Spr2),
		"Swap":           reflect.ValueOf(blas32.Swap),
		"Symm":           reflect.ValueOf(blas32.Symm),
		"Symv":           reflect.ValueOf(blas32.Symv),
		"Syr":            reflect.ValueOf(blas32.Syr),
		"Syr2":           reflect.ValueOf(blas32.Syr2),
		"Syr2k":          reflect.ValueOf(blas32.Syr2k),
		"Syrk":           reflect.ValueOf(blas32.Syrk),
		"Tbmv":           reflect.ValueOf(blas32.Tbmv),
		"Tbsv":           reflect.ValueOf(blas32.Tbsv),
		"Tpmv":           reflect.ValueOf(blas32.Tpmv),
		"Tpsv":           reflect.ValueOf(blas32.Tpsv),
		"Trmm":           reflect.ValueOf(blas32.Trmm),
		"Trmv":           reflect.ValueOf(blas32.Trmv),
		"Trsm":           reflect.ValueOf(blas32.Trsm),
		"Trsv":           reflect.ValueOf(blas32.Trsv),
		"Use":            reflect.ValueOf(blas32.Use),

		// type definitions
		"Band":               reflect.ValueOf((*blas32.Band)(nil)),
		"BandCols":           reflect.ValueOf((*blas32.BandCols)(nil)),
		"General":            reflect.ValueOf((*blas32.General)(nil)),
		"GeneralCols":        reflect.ValueOf((*blas32.GeneralCols)(nil)),
		"Symmetric":          reflect.ValueOf((*blas32.Symmetric)(nil)),
		"SymmetricBand":      reflect.ValueOf((*blas32.SymmetricBand)(nil)),
		"SymmetricBandCols":  reflect.ValueOf((*blas32.SymmetricBandCols)(nil)),
		"SymmetricCols":      reflect.ValueOf((*blas32.SymmetricCols)(nil)),
		"SymmetricPacked":    reflect.ValueOf((*blas32.SymmetricPacked)(nil)),
		"Triangular":         reflect.ValueOf((*blas32.Triangular)(nil)),
		"TriangularBand":     reflect.ValueOf((*blas32.TriangularBand)(nil)),
		"TriangularBandCols": reflect.ValueOf((*blas32.TriangularBandCols)(nil)),
		"TriangularCols":     reflect.ValueOf((*blas32.TriangularCols)(nil)),
		"TriangularPacked":   reflect.ValueOf((*blas32.TriangularPacked)(nil)),
		"Vector":             reflect.ValueOf((*blas32.Vector)(nil)),
	}
}
