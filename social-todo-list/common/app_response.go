package common

type successResponse struct {
	Data   interface{} `json:"data"`
	Paging interface{} `json:"paging, omitempty"`
	Filter interface{} `filter, omitempty`
}

func NewSuccessResponse(data, paging, filter interface{}) *successResponse {
	return &successResponse{Data: data, Paging: paging, Filter: filter}
}
func SimpleSuccessResponse(date interface{}) *successResponse {
	return NewSuccessResponse(date, nil, nil)
}
