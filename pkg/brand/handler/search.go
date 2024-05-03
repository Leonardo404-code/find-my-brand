package handler

import (
	"encoding/json"
	"io"
	"net/http"

	customErr "github.com/Leonargo404-code/find-my-brand/pkg/errors"
)

// @Summary Search
// @Tags Brand
// @Param terms body handler.SearchRequest true "Brand terms to search in web"
// @Router / [post]
// @Accept json
// @Produce json
// @Success 200 {object} brand.Result
func (h *handler) Search(w http.ResponseWriter, r *http.Request) {
	terms := new(SearchRequest)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		customErr.Error(w, http.StatusInternalServerError, err)
		return
	}

	if err := json.Unmarshal(body, &terms); err != nil {
		customErr.Error(w, http.StatusBadRequest, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
