package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type IwFriendsSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`

	request.PageInfo
}

// batch login
type BatchLoginReq struct {
	MyAccountIds []int `json:"myAccountIds"`
}

// api login
type ApiLoginReq struct {
	CC    string `json:"cc"`
	Phone string `json:"phone"`
	Sock5 string `json:"sock5"`
}

type ApiLoginRes struct {
	Data    Data   `json:"data"`
	Message string `json:"message"`
}

type Data struct {
	Status string `json:"status"`
	Reason string `json:"reason"`
	Vt     string `json:"vt"`
}

// api check out
type CheckLoginReq struct {
	CheckUser string `json:"check_user"`
	FromUser  string `json:"from_user"`
}
type CheckLoginRes struct {
	Data    bool   `json:"data"`
	Message string `json:"message"`
}

// api logout
type ApiLogoutReq struct {
	Phone string `json:"phone"`
}

type EndLogoutRes struct {
	Message string `json:"message"`
}

// add user
type AddUserReq struct {
	MyCountId        int    `json:"MyCountId"`
	FriendName       string `json:"friendName"`
	FriendFamilyName string `json:"friendFamilyName"`
	Phone            string `json:"phone"` //手机号带国家号
}

// api add user
type ApiAddUser struct {
	AddUser  string `json:"add_user"`
	FromUser string `json:"from_user"`
}

type DataInfo struct {
	Status  bool   `json:"status"`
	UserJID string `json:"user_jid"`
}

type ApiAddUserRes struct {
	Data    DataInfo `json:"data"`
	Message string   `json:"message"`
}

// check register
type CheckRegisterReq struct {
	CheckUser string `json:"check_user"`
	FromUser  string `json:"from_user"`
}

type CheckRegisterRes struct {
	Data    bool   `json:"data"`
	Message string `json:"message"`
}
