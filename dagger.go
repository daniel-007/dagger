package dagger

import (
	"github.com/autom8ter/dagger/primitive"
	"sort"
)

var globalGraph = primitive.NewGraph()

// NodeCount returns the total number of nodes in the graph
func NodeCount() int {
	i := 0
	globalGraph.RangeNodes(func(n primitive.Node) bool {
		if n != nil {
			i++
		}
		return true
	})
	return i
}

// EdgeCount returns the total number of edges in the graph
func EdgeCount() int {
	i := 0
	globalGraph.RangeEdges(func(n *primitive.Edge) bool {
		if n != nil {
			i++
		}
		return true
	})
	return i
}

// EdgeTypes returns the types of relationships/edges/connections in the graph
func EdgeTypes() []string {
	edgeTypes := globalGraph.EdgeTypes()
	sort.Strings(edgeTypes)
	return edgeTypes
}

// NodeTypes returns the types of nodes in the graph
func NodeTypes() []string {
	nodeTypes := globalGraph.NodeTypes()
	sort.Strings(nodeTypes)
	return nodeTypes
}

// GetNode gets a node from the graph
func GetNode(id primitive.TypedID) (*Node, bool) {
	n, ok := globalGraph.GetNode(id)
	if !ok {
		return nil, false
	}
	return &Node{n}, true
}

// RangeNodeTypes iterates over nodes of a given type until the iterator returns false
func RangeNodeTypes(typ primitive.Type, fn func(n *Node) bool) {
	globalGraph.RangeNodeTypes(typ, func(n primitive.Node) bool {
		return fn(&Node{n})
	})
}

// RangeNodes iterates over all nodes until the iterator returns false
func RangeNodes(fn func(n *Node) bool) {
	globalGraph.RangeNodes(func(n primitive.Node) bool {
		return fn(&Node{n})
	})
}

// RangeEdges iterates over all edges/connections until the iterator returns false
func RangeEdges(fn func(e *Edge) bool) {
	globalGraph.RangeEdges(func(e *primitive.Edge) bool {
		this, err := edgeFrom(e)
		if err != nil {
			panic(err)
		}
		return fn(this)
	})
}

// RangeEdgeTypes iterates over edges/connections of a given type until the iterator returns false
func RangeEdgeTypes(edgeType primitive.Type, fn func(e *Edge) bool) {
	globalGraph.RangeEdgeTypes(edgeType, func(e *primitive.Edge) bool {
		this, err := edgeFrom(e)
		if err != nil {
			panic(err)
		}
		return fn(this)
	})
}

// HasNode returns true if a node with the typed ID exists in the graph
func HasNode(id primitive.TypedID) bool {
	return globalGraph.HasNode(id)
}

// Close closes the global graph instance
func Close() {
	globalGraph.Close()
}

// ForeignKey is a helper that returns a primitive.TypedID from the given type and id
func ForeignKey(typ, id string) primitive.TypedID {
	return primitive.ForeignKey(typ, id)
}
