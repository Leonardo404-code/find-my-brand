package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Leonargo404-code/find-my-brand/pkg/brand"
	customErr "github.com/Leonargo404-code/find-my-brand/pkg/errors"
	"github.com/Leonargo404-code/find-my-brand/pkg/mail"
	"github.com/badoux/checkmail"
)

// @Summary Search
// @Tags Brand
// @Param terms body handler.SearchRequest true "Brand terms to search in web"
// @Router / [post]
// @Accept json
// @Produce json
// @Success 200 {object} brand.Result
func (h *handler) Search(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

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

	if err := checkmail.ValidateFormat(findBrandReq.Email); err != nil {
		customErr.Error(
			w,
			http.StatusBadRequest,
			fmt.Errorf("invalid e-mail: %v", err),
		)
	}

	location := r.URL.Query().Get("location")
	if location == "" {
		location = "Brazil"
	}

	result, err := h.brandSvc.SearchTerms(findBrandReq, location)
	if err != nil {
		customErr.Error(w, http.StatusInternalServerError, err)
		return
	}

	if len(result.Ads) == 0 {
		customErr.Error(w, http.StatusNotFound, fmt.Errorf("results not found"))
		return
	}

	if err := mail.SendEmail(result, findBrandReq.Email); err != nil {
		customErr.Error(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
