package services

import (
	"context"
	"errors"
	"synonyms/internal/synonyms"
	"testing"

	"github.com/matryer/is"
)

type mockGraph struct {
	Graph
}

func (mg *mockGraph) Nodes(ctx context.Context, word string) ([]string, error) {
	return nil, nil
}

func (mg *mockGraph) AddEdge(ctx context.Context, word, synonym string) {}

func (mg *mockGraph) Bfs(ctx context.Context, from string, levels int) ([]string, error) {
	return nil, nil
}

func TestServiceGetSynonymsOK(t *testing.T) {
	is := is.New(t)

	ctx := context.Background()
	s := NewSynonym(&mockGraph{})

	_, err := s.GetSynonyms(ctx, "a")
	is.True(errors.Is(err, nil))
}
func TestServiceGetSynonymsValidation(t *testing.T) {
	is := is.New(t)

	ctx := context.Background()
	s := NewSynonym(&mockGraph{})

	_, err := s.GetSynonyms(ctx, "")
	is.True(errors.Is(err, synonyms.ErrValiadation))
}

func TestServiceAddSynonymsOK(t *testing.T) {
	is := is.New(t)

	ctx := context.Background()
	s := NewSynonym(&mockGraph{})

	err := s.AddSynonyms(ctx, "a", "b")
	is.True(errors.Is(err, nil))

}
func TestServiceAddSynonymsValidation(t *testing.T) {
	is := is.New(t)

	ctx := context.Background()
	s := NewSynonym(&mockGraph{})

	testCases := []struct {
		word    string
		synonym string
	}{
		{"", ""},
		{"a", ""},
		{"", "a"},
		{"a", " "},
	}

	for _, tt := range testCases {
		err := s.AddSynonyms(ctx, tt.word, tt.synonym)
		is.True(errors.Is(err, synonyms.ErrValiadation))
	}
}
