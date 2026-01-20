package response

// APIResponse represents a standardized API response structure
type APIResponse struct {
	Data    interface{} `json:"data,omitempty"`
	Success bool        `json:"success"`
	Error   *ErrorInfo  `json:"error,omitempty"`
}

// PagingResponse represents a standardized paging API response structure
type PagingResponse struct {
	Data    *PagingData `json:"data,omitempty"`
	Success bool        `json:"success"`
	Error   *ErrorInfo  `json:"error,omitempty"`
}

// PagingData contains paginated data and metadata
type PagingData struct {
	Page         int64       `json:"page"`
	PageSize     int64       `json:"pageSize"`
	TotalElement int64       `json:"totalElement"`
	Data         interface{} `json:"data"`
}

// ErrorInfo contains error details
type ErrorInfo struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}
