package types

// BaseResponse
type BaseResponse struct {
	Message string `json:"message"`
	Page    int    `json:"page,omitempty"`
	Size    int    `json:"page_size,omitempty"`
}

// DataResponse
type DataResponse struct {
	BaseResponse
	Data interface{} `json:"data"`
}

// ErrorResponse
type ErrorResponse struct {
	BaseResponse
	Status int `json:"status"`
}
