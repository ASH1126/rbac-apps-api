package utils

type Respon struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func ErrorJSON(messages string) Respon {
	var respon Respon
	respon.Success = false
	respon.Message = messages
	return respon
}
