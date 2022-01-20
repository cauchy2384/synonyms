package services

import (
	"context"
	"fmt"
	"strings"
	"synonyms/internal/synonyms"
)

type Synonyms struct {
	graph Graph
}

type Graph interface {
	Nodes(ctx context.Context, word string) ([]string, error)
	Bfs(ctx context.Context, from string, levels int) ([]string, error)
	AddEdge(ctx context.Context, word, synonym string)
}

func NewSynonym(graph Graph) *Synonyms {
	return &Synonyms{
		graph: graph,
	}
}

func (s *Synonyms) GetSynonyms(ctx context.Context, word string) ([]string, error) {
	word = s.sanitizeWord(word)

	if len(word) == 0 {
		return nil, fmt.Errorf("%w: word is empty", synonyms.ErrValiadation)
	}

	// \todo move to config
	const searchDepth = 2
	synonyms, err := s.graph.Bfs(ctx, word, searchDepth)
	if err != nil {
		return nil, fmt.Errorf("nodes: %w", err)
	}

	return synonyms, nil
}

func (s *Synonyms) AddSynonyms(ctx context.Context, word, synonym string) error {
	word, synonym = s.sanitizeWord(word), s.sanitizeWord(synonym)

	if len(word) == 0 {
		return fmt.Errorf("%w: word is empty", synonyms.ErrValiadation)
	}
	if len(synonym) == 0 {
		return fmt.Errorf("%w: synonym is empty", synonyms.ErrValiadation)
	}

	s.graph.AddEdge(ctx, word, synonym)
	s.graph.AddEdge(ctx, synonym, word)

	return nil
}

func (s *Synonyms) sanitizeWord(word string) string {
	word = strings.TrimSpace(word)
	word = strings.ToLower(word)

	return word
}
