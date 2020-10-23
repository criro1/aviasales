package service

import (
	"context"
	"sort"
	"strings"

	v1 "github.com/criro1/aviasales/pkg/api/v1"
)

// Service ...
type Service interface {
	Load(ctx context.Context, words []string) (err error)
	Get(ctx context.Context, word *string) (response v1.GetResponse, err error)
}

type service struct {
	m map[string]bool
}

// Load ...
func (s *service) Load(ctx context.Context, words []string) (err error) {
	for _, word := range words {
		if _, ok := s.m[word]; !ok {
			s.m[word] = true
		}
	}
	return
}

// Get ...
func (s *service) Get(ctx context.Context, word *string) (response v1.GetResponse, err error) {
	if word == nil {
		return
	}

	*word = strings.ToLower(*word)
	wordSplit := strings.Split(*word, "")
	sort.Strings(wordSplit)
	*word = strings.Join(wordSplit, "")

	for key := range s.m {
		keySplit := strings.Split(strings.ToLower(key), "")
		sort.Strings(keySplit)
		keySort := strings.Join(keySplit, "")
		if *word == keySort {
			response.Data.Anagram = append(response.Data.Anagram, key)
		}
	}
	return
}

// NewService ...
func NewService(m map[string]bool) Service {
	return &service{
		m: m,
	}
}
