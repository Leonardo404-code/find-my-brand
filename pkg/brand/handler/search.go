package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Leonargo404-code/find-my-brand/pkg/brand"
	customErr "github.com/Leonargo404-code/find-my-brand/pkg/errors"
	"github.com/Leonargo404-code/find-my-brand/pkg/mail"
)

// @Summary Search
// @Tags Brand
// @Param terms body handler.SearchRequest true "Brand terms to search in web"
// @Router / [post]
// @Accept json
// @Produce json
// @Success 200 {object} brand.Result
func (h *handler) Search(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		customErr.Error(w, http.StatusInternalServerError, err)
		return
	}

	findBrandReq := new(brand.FindBrandRequest)
	if err := json.Unmarshal(body, &findBrandReq); err != nil {
		customErr.Error(
			w,
			http.StatusBadRequest,
			fmt.Errorf("error in unmarshal body: %v", err),
		)
		return
	}

	location := r.URL.Query().Get("location")

	result, err := h.brandSvc.SearchTerms(findBrandReq, location)
	if err != nil {
		customErr.Error(w, http.StatusInternalServerError, err)
		return
	}

	if err := mail.SendEmail(result); err != nil {
		customErr.Error(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
