package models

type ResetCartRequest struct {
	CartCode string
}
type ResetCartByUserIdRequest struct {
	UserId int64
}

type ResetCartByEmailRequest struct {
	Email string
}
