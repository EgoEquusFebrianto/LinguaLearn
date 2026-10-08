package service

import (
	"context"
	"errors"
	"strings"

	"github.com/EgoEquusFebrianto/LinguaLearn/internal/delivery/http/response"
	"github.com/EgoEquusFebrianto/LinguaLearn/internal/domain"
)

type BankWordService struct {
	repository domain.BankWordRepository
}

func NewBankWordServiceService(
	repository domain.BankWordRepository,
) *BankWordService {
	return &BankWordService{
		repository: repository,
	}
}

func (s *BankWordService) Search(
	ctx context.Context,
	query string,
	page int,
	limit int,
) (*response.BankWordSearchResponse, error) {
	if page < 1 {
		return nil, errors.New("Page must be positif number")
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

	return &response.BankWordSearchResponse{
		Data:       data,
		Page:       int(page),
		Limit:      int(limit),
		Total:      total,
		TotalPages: TotalPages,
	}, nil
}

func (s *BankWordService) FindByWord(
	ctx context.Context,
	word string,
) (*domain.BankWord, error) {
	word = strings.ToLower(strings.TrimSpace(word))

	return s.repository.FindByWord(
		ctx,
		word,
	)
}
