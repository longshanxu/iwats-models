package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type IwJiemaPlatformSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`

	request.PageInfo
}

type Balance struct {
	AccessBalance float64 `json:"access_balance"`
}

// 5sms
type DefaultCountry struct {
	Name   string `json:"name"`
	Iso    string `json:"iso"`
	Prefix string `json:"prefix"`
}

type DefaultOperator struct {
	Name string `json:"name"`
}

type Profile struct {
	ID                int             `json:"id"`
	Email             string          `json:"email"`
	Balance           float64         `json:"balance"`
	Rating            float64         `json:"rating"`
	DefaultCountry    DefaultCountry  `json:"default_country"`
	DefaultOperator   DefaultOperator `json:"default_operator"`
	FrozenBalance     float64         `json:"frozen_balance"`
	LastTopOrders     string          `json:"last_top_orders"`
	LastTopIdx        int             `json:"last_top_idx"`
	LastOrder         string          `json:"last_order"`
	TotalActiveOrders int             `json:"total_active_orders"`
	DidOrder          bool            `json:"did_order"`
}

// sms-man
type SmsManInfo struct {
	Balance  string `json:"balance"`
	Hold     int    `json:"hold"`
	Channels int    `json:"channels"`
}
