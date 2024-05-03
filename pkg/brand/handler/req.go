package handler

type SearchRequest struct {
	Terms []string `json:"terms,omitempty" example:"term1,term2,term3"`
}
