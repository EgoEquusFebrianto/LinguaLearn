package handler

import (
	"net/http"
	"strconv"

	"github.com/EgoEquusFebrianto/LinguaLearn/internal/repository"
	"github.com/EgoEquusFebrianto/LinguaLearn/internal/service"
	"github.com/EgoEquusFebrianto/LinguaLearn/internal/utils"
)

type BankWordHandler struct {
	service *service.BankWordService
}

func NewBankWordHandler(
	service *service.BankWordService,
) *BankWordHandler {
	return &BankWordHandler{
		service: service,
	}
}

func (h *BankWordHandler) Search(
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
		utils.Error(
			w,
			http.StatusInternalServerError,
			"failed to search dictionary",
		)
		return
	}

	utils.JSON(
		w,
		http.StatusOK,
		result,
	)
}

func (h *BankWordHandler) FindByWord(
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
			utils.Error(
				w,
				http.StatusNotFound,
				"word not found",
			)
			return
		}

		utils.Error(
			w,
			http.StatusInternalServerError,
			"failed to get word",
		)
		return
	}

	utils.JSON(
		w,
		http.StatusOK,
		result,
	)
}
