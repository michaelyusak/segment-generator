package entity

const (
	CodeSuccess             = "SUCCESS"
	CodeInternalServerError = "INTERNAL_SERVER_ERROR"
	CodeBadRequest          = "BAD_REQUEST"
	CodeNotFound            = "NOT_FOUND"
)

type Response struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}
