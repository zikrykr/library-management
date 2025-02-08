package pkg

import (
	"math"
)

const maxLimit, defaultLimit int = 1000, 10

type Pagination struct {
	CurrentPage     int64  `json:"current_page"`
	CurrentElements int64  `json:"current_elements"`
	TotalPages      int64  `json:"total_pages"`
	TotalElements   int64  `json:"total_elements"`
	SortBy          string `json:"sort_by"`
}

func ValidateLimit(l int) int {
	if l < 1 {
		return defaultLimit
	}

	if l > maxLimit {
		return maxLimit
	}
	return l
}

func ValidatePage(p int) int {
	if p < 1 {
		return 1
	}
	return p
}

func ValidatePagination(page, limit int) (int, int) {
	return ValidatePage(page), ValidateLimit(limit)
}

func CalculatePagination(totalRecords, page, pageLimit int64, sortBy string) *Pagination {
	if totalRecords == 0 {
		return &Pagination{
			CurrentPage: 1,
			TotalPages:  1,
		}
	}

	var curEl int64
	curEl = pageLimit
	if totalRecords < pageLimit {
		curEl = totalRecords
	}

	return &Pagination{
		CurrentPage:     page,
		CurrentElements: curEl,
		TotalPages:      int64(math.Ceil(float64(totalRecords) / float64(pageLimit))),
		TotalElements:   totalRecords,
		SortBy:          sortBy,
	}
}
