package repository

import (
	"context"
	"sort"
	"testing"

	"github.com/matryer/is"
)

func TestGraphAddEdges(t *testing.T) {
	ctx := context.Background()
	g := NewGraph()

	g.AddEdge(ctx, "a", "b")
	g.AddEdge(ctx, "b", "a")
	g.AddEdge(ctx, "a", "b")
}

func TestGraphNodesDirectional(t *testing.T) {
	is := is.New(t)

	ctx := context.Background()
	g := NewGraph()

	g.AddEdge(ctx, "a", "b")

	nodes, err := g.Nodes(ctx, "a")
	is.NoErr(err)
	is.Equal(nodes, []string{"b"})

	nodes, err = g.Nodes(ctx, "b")
	is.NoErr(err)
	is.Equal(nodes, []string{})
}

func TestGraphNodesBidirectional(t *testing.T) {
	is := is.New(t)

	ctx := context.Background()
	g := NewGraph()

	g.AddEdge(ctx, "a", "b")
	g.AddEdge(ctx, "b", "a")

	nodes, err := g.Nodes(ctx, "a")
	is.NoErr(err)
	is.Equal(nodes, []string{"b"})

	nodes, err = g.Nodes(ctx, "b")
	is.NoErr(err)
	is.Equal(nodes, []string{"a"})
}

func TestGraphBfs(t *testing.T) {
	is := is.New(t)

	ctx := context.Background()
	g := NewGraph()

	g.AddEdge(ctx, "a", "b")
	g.AddEdge(ctx, "b", "a")

	g.AddEdge(ctx, "b", "c")
	g.AddEdge(ctx, "c", "b")

	nodes, err := g.Nodes(ctx, "a")
	is.NoErr(err)
	is.Equal(nodes, []string{"b"})

	nodes, err = g.Nodes(ctx, "b")
	is.NoErr(err)
	is.Equal(
		sort.StringSlice(nodes),
		sort.StringSlice([]string{"a", "c"}),
	)

	nodes, err = g.Nodes(ctx, "c")
	is.NoErr(err)
	is.Equal(nodes, []string{"b"})

	nodes, err = g.Bfs(ctx, "a", 1)
	is.NoErr(err)
	is.Equal(nodes, []string{"b"})

	nodes, err = g.Bfs(ctx, "a", 2)
	is.NoErr(err)
	is.Equal(
		sort.StringSlice(nodes),
		sort.StringSlice([]string{"b", "c"}),
	)
}
