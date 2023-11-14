package models

type BaseResponseModel struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Error   interface{} `json:"error"`
}
