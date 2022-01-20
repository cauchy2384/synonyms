package repository

import (
	"context"
	"fmt"
	"sync"
	"synonyms/internal/synonyms"
)

type Graph struct {
	mu sync.Mutex
	m  map[string]map[string]struct{}
}

func NewGraph() *Graph {
	return &Graph{
		mu: sync.Mutex{},
		m:  make(map[string]map[string]struct{}),
	}
}

func (g *Graph) AddEdge(_ context.Context, from, to string) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// create nodes if not exist
	if _, ok := g.m[from]; !ok {
		g.m[from] = make(map[string]struct{})
	}
	if _, ok := g.m[to]; !ok {
		g.m[to] = make(map[string]struct{})
	}

	// add edge from -> to
	g.m[from][to] = struct{}{}
}

func (g *Graph) Nodes(_ context.Context, from string) ([]string, error) {
	g.mu.Lock()
	defer g.mu.Unlock()

	if _, ok := g.m[from]; !ok {
		return nil, fmt.Errorf("%w: %s", synonyms.ErrNotFound, from)
	}

	nodes := make([]string, 0, len(g.m[from]))
	for node := range g.m[from] {
		nodes = append(nodes, node)
	}

	return nodes, nil
}

func (g *Graph) Bfs(ctx context.Context, from string, levels int) ([]string, error) {
	g.mu.Lock()
	defer g.mu.Unlock()

	visited := make(map[string]struct{})

	return g.bfs(ctx, from, 0, levels, visited)
}

func (g *Graph) bfs(ctx context.Context, from string,
	level, maxLevel int, visited map[string]struct{},
) ([]string, error) {

	if level >= maxLevel {
		return nil, nil
	}

	if _, ok := g.m[from]; !ok {
		return nil, nil
	}

	visited[from] = struct{}{}

	nodes := make([]string, 0, len(g.m[from]))
	for node := range g.m[from] {

		if _, ok := visited[node]; ok {
			continue
		}
		nodes = append(nodes, node)

		nn, err := g.bfs(ctx, node, level+1, maxLevel, visited)
		if err != nil {
			return nil, err
		}
		nodes = append(nodes, nn...)
	}

	return nodes, nil
}
