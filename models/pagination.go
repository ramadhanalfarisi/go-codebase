package models

import (
	"math"
	"net/http"
	"strconv"
)

type Pagination struct {
	Page       int64 `json:"page"`       // current page
	Size       int64 `json:"size"`       // limit
	Offset     int64 `json:"offset"`     // limit
	TotalPages int64 `json:"totalPages"` // total page
	Total      int64 `json:"total"`      // total row
	Visible    int64 `json:"visible"`    // total row in current page
	Last       bool  `json:"last"`       // is last page
	First      bool  `json:"first"`      // is first page
}

func (pagination *Pagination) GetPagination(r *http.Request) Pagination {
	var climit, cpage int
	var int_limit int64
	var int_page int64
	var int_offset int64

	page, ok1 := r.URL.Query()["page"]
	limit, ok2 := r.URL.Query()["limit"]

	if ok1 {
		cpage, _ = strconv.Atoi(page[0])
	}
	if ok2 {
		climit, _ = strconv.Atoi(limit[0])
	}

	if !ok1 || len(page[0]) < 1 {
		int_page = 1
	} else {
		if int64(climit) < 1 {
			int_page = 1
		} else {
			int_page = int64(cpage)
		}
	}

	if !ok2 || len(limit[0]) < 1 {
		int_limit = 10
	} else {
		int_limit = int64(climit)
	}

	if int_page == 1 {
		int_offset = 0
	} else {
		int_offset = (int64(int_page) - 1) * int64(int_limit)
	}

	pagination.Page = int64(int_page)
	pagination.Size = int64(int_limit)
	pagination.Offset = int64(int_offset)
	return *pagination
}

func (pagination *Pagination) CreatePagination(r *http.Request) Pagination {

	if pagination.Total <= pagination.Size {
		pagination.Visible = pagination.Total
	} else if pagination.Total > pagination.Size {
		current_total := pagination.Page * pagination.Size
		if pagination.Total > current_total {
			pagination.Visible = pagination.Size
		} else {
			mod_total := pagination.Total % pagination.Size
			pagination.Visible = mod_total
		}
	}
	total_pages := math.Ceil(float64(pagination.Total / pagination.Size))
	pagination.TotalPages = int64(total_pages)
	if pagination.Page == 1 {
		pagination.First = true
		pagination.Last = false
	} else if pagination.Page == pagination.TotalPages {
		pagination.First = false
		pagination.Last = true
	} else {
		pagination.First = false
		pagination.Last = false
	}
	return *pagination
}
