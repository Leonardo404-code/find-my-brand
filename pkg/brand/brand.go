package brand

type (
	Result struct {
		Ads []Ads `json:"ads,omitempty" example:"competing_brand_name:https://competingbrand.com.br"`
	}

	Ads struct {
		Position int    `json:"position,omitempty"`
		Title    string `json:"title,omitempty"`
		Domain   string `json:"domain,omitempty"`
		Link     string `json:"link,omitempty"`
	}

	FindBrandRequest struct {
		Terms []string `json:"terms,omitempty" example:"term1,term2,term3"`
	}
)
