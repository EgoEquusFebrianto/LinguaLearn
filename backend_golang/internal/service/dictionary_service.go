package service

import (
	"context"
	"strings"

	"github.com/EgoEquusFebrianto/LinguaLearn/internal/data/response"
	"github.com/EgoEquusFebrianto/LinguaLearn/internal/models"
	"github.com/EgoEquusFebrianto/LinguaLearn/internal/repository"
)

type DictionaryService struct {
	repository repository.BankWordRepository
}

func NewDictionaryService(
	repository repository.BankWordRepository,
) *DictionaryService {
	return &DictionaryService{
		repository: repository,
	}
}

func (s *DictionaryService) Search(
	ctx context.Context,
	query string,
	page int,
	limit int,
) (*response.DictionarySearchResponse, error) {
	if page < 1 {
		page = 1
	}

	if limit < 1 {
		limit = 20
	}

	if limit > 100 {
		limit = 100
	}

	offset := (page - 1) * limit

	data, err := s.repository.Search(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}

	total, err := s.repository.Count(ctx, query)
	if err != nil {
		return nil, err
	}

	TotalPages := (int(total) + limit - 1) / limit

	return &response.DictionarySearchResponse{
		Data: data,
		Page: int(page),
		Limit: int(limit),
		Total: total,
		TotalPages: TotalPages,
	}, nil
}

func (s *DictionaryService) FindByWord(
	ctx context.Context,
	word string,
) (*models.BankWord, error) {
	word = strings.ToLower(strings.TrimSpace(word))

	return s.repository.FindByWord(
		ctx,
		word,
	)
}