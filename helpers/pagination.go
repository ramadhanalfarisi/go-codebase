package helpers

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
	var cLimit, cPage int
	var intLimit int
	var intPage int
	var intOffset int
	var isAll bool

	page, ok1 := r.URL.Query()["page"]
	limit, ok2 := r.URL.Query()["limit"]

	if ok2 {
		if limit[0] != "all" {
			cLimit, _ = strconv.Atoi(limit[0])
			isAll = false
		} else {
			isAll = true
		}
	}

	if !isAll {
		if ok1 {
			cPage, _ = strconv.Atoi(page[0])
		}

		if !ok1 || len(page[0]) < 1 {
			intPage = 1
		} else {
			if cLimit < 1 {
				intPage = 1
			} else {
				intPage = cPage
			}
		}

		if !ok2 || len(limit[0]) < 1 {
			intLimit = 10
		} else {
			intLimit = cLimit
		}

		if intPage == 1 {
			intOffset = 0
		} else {
			intOffset = (intPage - 1) * intLimit
		}

		pagination.Page = intPage
		pagination.Size = intLimit
		pagination.Offset = intOffset
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
