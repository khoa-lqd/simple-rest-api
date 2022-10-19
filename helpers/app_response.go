package helpers

type successRes struct {
	Data interface{} `json:"data"`
}

func NewSuccessResponse(data interface{}) *successRes {
	return &successRes{Data: data}
}
