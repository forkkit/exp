// Code generated by 'goexports gonum.org/v1/gonum/graph/simple'. DO NOT EDIT.

// Copyright ©2019 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build go1.14,!go1.15

package yaegi

import (
	"reflect"

	"gonum.org/v1/gonum/graph/simple"
)

func init() {
	Symbols["gonum.org/v1/gonum/graph/simple"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"NewDirectedGraph":           reflect.ValueOf(simple.NewDirectedGraph),
		"NewDirectedMatrix":          reflect.ValueOf(simple.NewDirectedMatrix),
		"NewDirectedMatrixFrom":      reflect.ValueOf(simple.NewDirectedMatrixFrom),
		"NewUndirectedGraph":         reflect.ValueOf(simple.NewUndirectedGraph),
		"NewUndirectedMatrix":        reflect.ValueOf(simple.NewUndirectedMatrix),
		"NewUndirectedMatrixFrom":    reflect.ValueOf(simple.NewUndirectedMatrixFrom),
		"NewWeightedDirectedGraph":   reflect.ValueOf(simple.NewWeightedDirectedGraph),
		"NewWeightedUndirectedGraph": reflect.ValueOf(simple.NewWeightedUndirectedGraph),

		// type definitions
		"DirectedGraph":           reflect.ValueOf((*simple.DirectedGraph)(nil)),
		"DirectedMatrix":          reflect.ValueOf((*simple.DirectedMatrix)(nil)),
		"Edge":                    reflect.ValueOf((*simple.Edge)(nil)),
		"Node":                    reflect.ValueOf((*simple.Node)(nil)),
		"UndirectedGraph":         reflect.ValueOf((*simple.UndirectedGraph)(nil)),
		"UndirectedMatrix":        reflect.ValueOf((*simple.UndirectedMatrix)(nil)),
		"WeightedDirectedGraph":   reflect.ValueOf((*simple.WeightedDirectedGraph)(nil)),
		"WeightedEdge":            reflect.ValueOf((*simple.WeightedEdge)(nil)),
		"WeightedUndirectedGraph": reflect.ValueOf((*simple.WeightedUndirectedGraph)(nil)),
	}
}
