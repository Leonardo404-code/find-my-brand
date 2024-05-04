package brand

type (
	Result struct {
		Ads []map[string]any `json:"ads,omitempty" example:"competing_brand_name:https://competingbrand.com.br"`
	}

	Terms struct {
		Terms []string `json:"terms,omitempty" example:"term1,term2,term3"`
	}
)
