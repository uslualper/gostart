package http

import (
	"math"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Paginate struct {
	Page     int
	Limit    int
	MaxLimit int
}

func InitPaginate(c *fiber.Ctx) Paginate {
	p := new(Paginate)

	page, pageErr := strconv.Atoi(c.Query("page"))
	limit, limitErr := strconv.Atoi(c.Query("limit"))

	if pageErr != nil {
		page = 1
	}
	if limitErr != nil {
		limit = 20
	}
	if p.MaxLimit == 0 {
		p.MaxLimit = 100
	}

	p.Page = page
	p.Limit = limit

	if p.Limit > p.MaxLimit {
		p.Limit = p.MaxLimit
	}

	return *p
}

func (p *Paginate) Response(data interface{}, total int64) (response fiber.Map) {
	response = fiber.Map{
		"current_page": p.Page,
		"last_page":    math.Ceil(float64(total) / float64(p.Limit)),
		"per_page":     p.Limit,
		"total":        total,
		"data":         data,
	}
	return
}
