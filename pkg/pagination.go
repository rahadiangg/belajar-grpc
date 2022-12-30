package pkg

import (
	"math"

	"github.com/rahadiangg/belajar-grpc/model"
	"gorm.io/gorm"
)

func Pagination(sql *gorm.DB, page int64, pagination *model.Pagination) (offset, limit int64) {
	var total int64
	limit = 1
	sql.Count(&total)
	if page == 1 {
		offset = 0
	} else {
		offset = (page - 1) * limit
	}

	pagination.Total = uint64(total)
	pagination.PerPage = uint64(limit)
	pagination.CurrentPage = uint64(page)
	pagination.LastPage = uint64(math.Ceil(float64(total) / float64(limit)))

	return
}
