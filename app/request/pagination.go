package request

type paginationRequest struct {
	Limit  int `form:"limit"`
	Offset int `form:"offset"`
}
