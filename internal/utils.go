package internal

import "reflect"

type PaginatedResponse struct {
	Data interface{} `json:"data"`
	Meta struct {
		Page       int `json:"page"`
		PageSize   int `json:"page_size"`
		Total      int `json:"total"`
		TotalPages int `json:"total_pages"`
	} `json:"meta"`
}

func Paginate(data interface{}, page, pageSize int) *PaginatedResponse {
	v := reflect.ValueOf(data)
	if v.Kind() != reflect.Slice {
		return &PaginatedResponse{Data: data}
	}

	total := v.Len()

	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	if pageSize > 100 {
		pageSize = 100
	}

	offset := (page - 1) * pageSize
	if offset > total {
		offset = total
	}

	end := offset + pageSize
	if end > total {
		end = total
	}

	totalPages := (total + pageSize - 1) / pageSize
	if totalPages < 1 {
		totalPages = 1
	}

	response := &PaginatedResponse{
		Data: v.Slice(offset, end).Interface(),
	}
	response.Meta.Page = page
	response.Meta.PageSize = pageSize
	response.Meta.Total = total
	response.Meta.TotalPages = totalPages

	return response
}
