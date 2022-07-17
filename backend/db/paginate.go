package db

import (
	"fmt"
	"math"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/xdarkyne/udemy/db/models"
	"gorm.io/gorm"
)

func Paginate(c *fiber.Ctx, orm *gorm.DB, entity models.Entity) (fiber.Map, error) {
	page, pageErr := strconv.Atoi(c.Query("page", "1"))
	limit, limitErr := strconv.Atoi(c.Query("limit", "10"))
	if pageErr != nil || limitErr != nil {
		return nil, fmt.Errorf("invalid query parameters")
	}
	offset := (page - 1) * limit
	total := entity.Count(ORM)
	data := entity.Take(ORM, limit, offset)
	lastPage := math.Ceil(float64(total) / float64(limit))

	return fiber.Map{
		"data": data,
		"meta": fiber.Map{
			"total":     total,
			"page":      page,
			"last_page": lastPage,
			"limit":     limit,
		},
	}, nil
}
