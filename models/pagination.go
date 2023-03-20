package models

import (
	"math"
	"net/http"
	"strconv"
)

type Pagination struct {
	Page       int  `json:"page"`       // current page
	Size       int  `json:"size"`       // limit
	Offset     int  `json:"offset"`     // limit
	TotalPages int  `json:"totalPages"` // total page
	Total      int  `json:"total"`      // total row
	Visible    int  `json:"visible"`    // total row in current page
	Last       bool `json:"last"`       // is last page
	First      bool `json:"first"`      // is first page
}

func (pagination *Pagination) GetPagination(r *http.Request) Pagination {
	var climit, cpage int
	var int_limit int
	var int_page int
	var int_offset int
	var isall bool

	page, ok1 := r.URL.Query()["page"]
	limit, ok2 := r.URL.Query()["limit"]

	if ok2 {
		if limit[0] != "all" {
			climit, _ = strconv.Atoi(limit[0])
			isall = false
		} else {
			isall = true
		}
	}

	if !isall {
		if ok1 {
			cpage, _ = strconv.Atoi(page[0])
		}

		if !ok1 || len(page[0]) < 1 {
			int_page = 1
		} else {
			if climit < 1 {
				int_page = 1
			} else {
				int_page = cpage
			}
		}

		if !ok2 || len(limit[0]) < 1 {
			int_limit = 10
		} else {
			int_limit = climit
		}

		if int_page == 1 {
			int_offset = 0
		} else {
			int_offset = (int_page - 1) * int_limit
		}

		pagination.Page = int_page
		pagination.Size = int_limit
		pagination.Offset = int_offset
	} else {
		pagination.Page = 0
		pagination.Size = 0
		pagination.Offset = 0
	}

	return *pagination
}

func (pagination *Pagination) CreatePagination() Pagination {

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
	pagination.TotalPages = int(total_pages)
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
