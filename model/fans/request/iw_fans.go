package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type IwFansSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	FansGroupId    int        `json:"fansGroupId" form:"fansGroupId"`
	Phone          string     `json:"phone" form:"phone"`
	IDs     []uint64  `json:"IDs" form:"IDs" gorm:"column:ids;comment:粉丝的ID;size:255;"`                //粉丝的ID

	request.PageInfo
}
