package internal

type PaginatedResponse struct {
	Data interface{} `json:"data"`
	Meta struct {
		Page       int `json:"page"`
		PageSize   int `json:"page_size"`
		Total      int `json:"total"`
		TotalPages int `json:"total_pages"`
	} `json:"meta"`
}

// Paginate creates a paginated response
func Paginate(data interface{}, page, pageSize, total int) *PaginatedResponse {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	if pageSize > 100 {
		pageSize = 100
	}

	totalPages := (total + pageSize - 1) / pageSize
	if totalPages < 1 {
		totalPages = 1
	}

	response := &PaginatedResponse{
		Data: data,
	}
	response.Meta.Page = page
	response.Meta.PageSize = pageSize
	response.Meta.Total = total
	response.Meta.TotalPages = totalPages

	return response
}
