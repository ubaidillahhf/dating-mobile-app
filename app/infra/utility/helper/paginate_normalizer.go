package helper

import "gorm.io/gorm"

func pagePerPageConv(page, perPage int) (skip, limit int) {
	var tempPage = int(page)
	if tempPage == 0 {
		tempPage = 1
	}

	tempPage = tempPage - 1

	skipConv := tempPage * perPage

	return skipConv, perPage

}

func Paginate(page, perPage int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		skip, limit := pagePerPageConv(page, perPage)

		return db.Offset(skip).Limit(limit)
	}
}
