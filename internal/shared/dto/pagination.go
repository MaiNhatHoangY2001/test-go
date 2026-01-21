package dto

// PaginationInput represents common pagination query parameters
type PaginationInput struct {
	PageNum  int `form:"page" binding:"omitempty,min=1"`
	PageSize int `form:"limit" binding:"omitempty,min=1,max=100"`
}

// PaginationInfo represents pagination metadata in responses
type PaginationInfo struct {
	PageNum    int   `json:"page"`
	PageSize   int   `json:"page_size"`
	TotalItems int64 `json:"total_items"`
	TotalPages int   `json:"total_pages"`
}
