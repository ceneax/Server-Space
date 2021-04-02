package entity

import (
	"unsafe"

	"xcdh.space/space/config"
)

type Pageinfo struct {
	Page     int    `form:"page" json:"page"`
	PageSize int    `form:"pagesize" json:"pagesize"`
	Total    int64  `form:"total" json:"total"`
	Sort     string `form:"sort" json:"sort"`
}

func (pi *Pageinfo) GetLimit(totalNow int64) (pageSize int, rowBegin int) {
	if pi.PageSize == 0 {
		pi.PageSize = config.PageSize
	}
	if pi.Page == 0 {
		pi.Page = 1
	}
	pageSize = pi.PageSize
	rowBegin = pi.PageSize * (pi.Page - 1)
	if pi.Total == 0 || pi.Page == 1 {
		return pageSize, rowBegin
	} else {
		int64rowBegin := totalNow - pi.Total + int64(rowBegin)
		rowBegin = *(*int)(unsafe.Pointer(&int64rowBegin))
		return pageSize, rowBegin
	}
}
