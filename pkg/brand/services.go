package brand

type Services interface {
	SearchTerms(
		terms *FindBrandRequest,
		location string,
	) (*Result, error)
}
