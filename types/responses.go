package types

type BaseResponse struct {
	Message string `json:"message"`
	Page    int    `json:"page,omitempty"`
	Size    int    `json:"page_size,omitempty"`
}

type DataResponse struct {
	BaseResponse
	Data interface{} `json:"data"`
}

type ErrorResponse struct {
	BaseResponse
	Status int `json:"status"`
}
