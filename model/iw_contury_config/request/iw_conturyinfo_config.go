package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type IwConturyinfoConfigSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`

	request.PageInfo
}

// add user
type AddUserReq struct {
	MyCountId        int    `json:"MyCountId"`
	FriendName       string `json:"friendName"`
	FriendFamilyName string `json:"friendFamilyName"`
	Phone            string `json:"phone"` //手机号带国家号
}
