package tokodto

type TokoFilterRequest struct {
	Limit int64  `query:"limit"`
	Page  int64  `query:"page"`
	Name  string `query:"nama"`
}
