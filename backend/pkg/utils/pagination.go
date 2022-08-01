package utils

import (
	"errors"
	"math"
)

type PaginationQuery struct {
	Filter	string
	Limit	float32
	Offset 	float32
	PageNum	float32
	MaxNum	float32
}

func NewPagination(filter string, limit float32, pageNum float32, itemCount float32) (*PaginationQuery, error) {

	var pages float32
	if itemCount < limit {
		pages = 1
	} else {
		pages = float32(math.Round(float64(itemCount) / float64(limit)))
	}

	if pageNum > pages {
		return nil,	errors.New("Selected page number more than available pages.")
	}

	return &PaginationQuery{Filter:filter, Limit:limit, PageNum: pageNum, MaxNum: pages, Offset: limit * (pageNum-1)}, nil
}