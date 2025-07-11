package dto

type PaginatedResponse struct {
	Total      int         `json:"total"`
	Page       int         `json:"page"`
	PageSize   int         `json:"page_size"`
	TotalPages int64       `json:"total_pages"`
	Data       interface{} `json:"data"`
}
