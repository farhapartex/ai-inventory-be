package dto

type PaginatedResponse struct {
	Total      int         `json:"total"`
	Page       int         `json:"page"`
	PageSize   int         `json:"page_size"`
	TotalPages int64       `json:"total_pages"`
	Data       interface{} `json:"data"`
}

type ListQueryDTO struct {
	Page     int    `form:"page" binding:"min=1"`
	PageSize int    `form:"pageSize" binding:"min=1,max=100"`
	Status   string `form:"status"`
	Search   string `form:"search"`
	SortBy   string `form:"sortBy"`
	SortDir  string `form:"sortDir"`
}
