package scopes

import (
	"github.com/shasw94/projX/pkg/utils"
	"gorm.io/gorm"
)

type GormPager interface {
	ToPaginate() func(db *gorm.DB) *gorm.DB
}

type GormPagination struct {
	*utils.Pagination
}

func (r *GormPagination) ToPaginate() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(utils.OffsetCal(r.Pagination.GetPage(), r.Pagination.GetLimit())).Limit(r.Pagination.GetLimit())
	}
}
