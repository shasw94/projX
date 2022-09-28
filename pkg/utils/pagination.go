package utils

import "math"

// IPagination abstraction for pagination data getters
type IPagination interface {
	Get() *Pagination
	GetPage() int
	GetLimit() int
}

// Pagination pagination data
type Pagination struct {
	Page  int
	Limit int
}

// Get Get pagination struct
func (p *Pagination) Get() *Pagination {
	return p
}

// GetPage get page value from pagination struct
func (p *Pagination) GetPage() int {
	if p.Page <= 0 {
		p.Page = 1
	}
	return p.Page
}

// GetLimit get limit value from pagination struct
func (p *Pagination) GetLimit() int {
	if p.Limit <= 0 {
		p.Limit = 20
	}
	return p.Limit
}

func NextPageCal(page int, totalPage int) int {
	if page == totalPage {
		return page
	}
	return page + 1
}

func PrevPageCal(page int) int {
	if page > 1 {
		return page - 1
	}
	return page
}

func TotalPage(count int64, limit int) int {
	return int(math.Ceil(float64(count) / float64(limit)))
}

func OffsetCal(page int, limit int) int {
	return (page - 1) * limit
}
