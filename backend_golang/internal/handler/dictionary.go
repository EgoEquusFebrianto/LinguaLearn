package handler

import (
	"net/http"
	"strconv"

	"github.com/EgoEquusFebrianto/LinguaLearn/internal/helper"
	"github.com/EgoEquusFebrianto/LinguaLearn/internal/repository"
	"github.com/EgoEquusFebrianto/LinguaLearn/internal/service"
)

type DictionaryHandler struct {
	service *service.DictionaryService
}

func NewDictionaryHandler(
	service *service.DictionaryService,
) *DictionaryHandler {
	return &DictionaryHandler{
		service: service,
	}
}

func (h *DictionaryHandler) Search(
	w http.ResponseWriter,
	r *http.Request,
) {
	query := r.URL.Query().Get("q")

	page, _ := strconv.Atoi(
		r.URL.Query().Get("page"),
	)

	limit, _ := strconv.Atoi(
		r.URL.Query().Get("limit"),
	)

	result, err := h.service.Search(
		r.Context(),
		query,
		page,
		limit,
	)

	if err != nil {
		helper.Error(
			w,
			http.StatusInternalServerError,
			"failed to search dictionary",
		)
		return
	}

	helper.JSON(
		w,
		http.StatusOK,
		result,
	)
}

func (h *DictionaryHandler) FindByWord(
	w http.ResponseWriter,
	r *http.Request,
) {
	word := r.URL.Query().Get("word")

	result, err := h.service.FindByWord(
		r.Context(),
		word,
	)

	if err != nil {
		if err == repository.ErrBankWordNotFound {
			helper.Error(
				w,
				http.StatusNotFound,
				"word not found",
			)
			return
		}

		helper.Error(
			w,
			http.StatusInternalServerError,
			"failed to get word",
		)
		return
	}

	helper.JSON(
		w,
		http.StatusOK,
		result,
	)
}