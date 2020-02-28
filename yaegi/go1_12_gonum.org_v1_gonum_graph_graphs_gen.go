// Code generated by 'goexports gonum.org/v1/gonum/graph/graphs/gen'. DO NOT EDIT.

// Copyright ©2019 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build go1.12,!go1.13

package yaegi

import (
	"reflect"

	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/graphs/gen"
)

func init() {
	Symbols["gonum.org/v1/gonum/graph/graphs/gen"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"BipartitePowerLaw":          reflect.ValueOf(gen.BipartitePowerLaw),
		"Duplication":                reflect.ValueOf(gen.Duplication),
		"Gnm":                        reflect.ValueOf(gen.Gnm),
		"Gnp":                        reflect.ValueOf(gen.Gnp),
		"NavigableSmallWorld":        reflect.ValueOf(gen.NavigableSmallWorld),
		"PowerLaw":                   reflect.ValueOf(gen.PowerLaw),
		"PreferentialAttachment":     reflect.ValueOf(gen.PreferentialAttachment),
		"SmallWorldsBB":              reflect.ValueOf(gen.SmallWorldsBB),
		"TunableClusteringScaleFree": reflect.ValueOf(gen.TunableClusteringScaleFree),

		// type definitions
		"GraphBuilder":      reflect.ValueOf((*gen.GraphBuilder)(nil)),
		"UndirectedMutator": reflect.ValueOf((*gen.UndirectedMutator)(nil)),

		// interface wrapper definitions
		"_GraphBuilder":      reflect.ValueOf((*_gonum_org_v1_gonum_graph_graphs_gen_GraphBuilder)(nil)),
		"_UndirectedMutator": reflect.ValueOf((*_gonum_org_v1_gonum_graph_graphs_gen_UndirectedMutator)(nil)),
	}
}

// _gonum_org_v1_gonum_graph_graphs_gen_GraphBuilder is an interface wrapper for GraphBuilder type
type _gonum_org_v1_gonum_graph_graphs_gen_GraphBuilder struct {
	WAddNode        func(a0 graph.Node)
	WHasEdgeBetween func(xid int64, yid int64) bool
	WNewEdge        func(from graph.Node, to graph.Node) graph.Edge
	WNewNode        func() graph.Node
	WSetEdge        func(e graph.Edge)
}

func (W _gonum_org_v1_gonum_graph_graphs_gen_GraphBuilder) AddNode(a0 graph.Node) { W.WAddNode(a0) }
func (W _gonum_org_v1_gonum_graph_graphs_gen_GraphBuilder) HasEdgeBetween(xid int64, yid int64) bool {
	return W.WHasEdgeBetween(xid, yid)
}
func (W _gonum_org_v1_gonum_graph_graphs_gen_GraphBuilder) NewEdge(from graph.Node, to graph.Node) graph.Edge {
	return W.WNewEdge(from, to)
}
func (W _gonum_org_v1_gonum_graph_graphs_gen_GraphBuilder) NewNode() graph.Node  { return W.WNewNode() }
func (W _gonum_org_v1_gonum_graph_graphs_gen_GraphBuilder) SetEdge(e graph.Edge) { W.WSetEdge(e) }

// _gonum_org_v1_gonum_graph_graphs_gen_UndirectedMutator is an interface wrapper for UndirectedMutator type
type _gonum_org_v1_gonum_graph_graphs_gen_UndirectedMutator struct {
	WAddNode        func(a0 graph.Node)
	WEdge           func(uid int64, vid int64) graph.Edge
	WEdgeBetween    func(xid int64, yid int64) graph.Edge
	WFrom           func(id int64) graph.Nodes
	WHasEdgeBetween func(xid int64, yid int64) bool
	WNewEdge        func(from graph.Node, to graph.Node) graph.Edge
	WNewNode        func() graph.Node
	WNode           func(id int64) graph.Node
	WNodes          func() graph.Nodes
	WRemoveEdge     func(fid int64, tid int64)
	WSetEdge        func(e graph.Edge)
}

func (W _gonum_org_v1_gonum_graph_graphs_gen_UndirectedMutator) AddNode(a0 graph.Node) { W.WAddNode(a0) }
func (W _gonum_org_v1_gonum_graph_graphs_gen_UndirectedMutator) Edge(uid int64, vid int64) graph.Edge {
	return W.WEdge(uid, vid)
}
func (W _gonum_org_v1_gonum_graph_graphs_gen_UndirectedMutator) EdgeBetween(xid int64, yid int64) graph.Edge {
	return W.WEdgeBetween(xid, yid)
}
func (W _gonum_org_v1_gonum_graph_graphs_gen_UndirectedMutator) From(id int64) graph.Nodes {
	return W.WFrom(id)
}
func (W _gonum_org_v1_gonum_graph_graphs_gen_UndirectedMutator) HasEdgeBetween(xid int64, yid int64) bool {
	return W.WHasEdgeBetween(xid, yid)
}
func (W _gonum_org_v1_gonum_graph_graphs_gen_UndirectedMutator) NewEdge(from graph.Node, to graph.Node) graph.Edge {
	return W.WNewEdge(from, to)
}
func (W _gonum_org_v1_gonum_graph_graphs_gen_UndirectedMutator) NewNode() graph.Node {
	return W.WNewNode()
}
func (W _gonum_org_v1_gonum_graph_graphs_gen_UndirectedMutator) Node(id int64) graph.Node {
	return W.WNode(id)
}
func (W _gonum_org_v1_gonum_graph_graphs_gen_UndirectedMutator) Nodes() graph.Nodes {
	return W.WNodes()
}
func (W _gonum_org_v1_gonum_graph_graphs_gen_UndirectedMutator) RemoveEdge(fid int64, tid int64) {
	W.WRemoveEdge(fid, tid)
}
func (W _gonum_org_v1_gonum_graph_graphs_gen_UndirectedMutator) SetEdge(e graph.Edge) { W.WSetEdge(e) }