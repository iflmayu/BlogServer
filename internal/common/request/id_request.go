package request

type IDRequest struct {
	ID uint `uri:"id" binding:"required"`
}
