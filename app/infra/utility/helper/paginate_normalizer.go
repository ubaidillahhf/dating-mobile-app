package helper

import "gorm.io/gorm"

func NormalizeAndGetDefaultPagination(page, perPage int64) (skip, limit int, defPage, defPerPage int64) {

	// page 2, perPage 10 = skip 10, limit 10
	// page 3, perPage 10 = skip 20. limit 10

	p, pp := paginationDefault(page, perPage)
	skip = int((p - 1) * pp)
	limit = int(pp)

	defPage = p
	defPerPage = pp

	return
}

func paginationDefault(page, perPage int64) (p, pp int64) {

	if page == 0 {
		p = 1
	} else {
		p = page
	}
	if perPage == 0 {
		pp = 10
	} else {
		pp = perPage
	}

	return
}

func GormPaginate(skip, limit int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		return db.Offset(skip).Limit(limit)
	}
}
