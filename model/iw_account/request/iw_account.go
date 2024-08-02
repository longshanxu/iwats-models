/*
 * @Author: Wangpengxiang 17600369018@163.com
 * @Date: 2024-07-30 21:34:19
 * @LastEditors: Wangpengxiang 17600369018@163.com
 * @LastEditTime: 2024-08-01 14:26:45
 * @FilePath: \server\model\iw_account\request\iw_account.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type IwAccountSearch struct {
	StartCreatedAt    *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt      *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	AccountType       []int        `json:"accountType" form:"accountType"`
	Phone             string     `json:"phone" form:"phone"`
	IsOnline          []int      `json:"isOnline" form:"isOnline"`
	Type              []int        `json:"type" form:"type"`
	GroupId           []int      `json:"groupId" form:"groupId"`
	Ids               []int      `json:"ids" form:"ids"`
	UpdateGroupId     int        `json:"updateGroupId" form:"updateGroupId"`
	AccountId         []int      `json:"accountId" form:"accountId"`
	IpPlatformGroupId int        `json:"ipPlatformGroupId" form:"ipPlatformGroupId"`
	IpPlatformName    string     `json:"ipPlatformName" form:"ipPlatformName"`
	IpGroupId         int        `json:"ipGroupId" form:"ipGroupId"`
	AccountPhone      []string   `json:"accountPhone" form:"accountPhone"`
	DeviceData        []string   `json:"deviceData" form:"deviceData"`

	request.PageInfo
}

type AddUserReq struct {
	MyCountId        int    `json:"MyCountId"`
	FriendName       string `json:"friendName"`
	FriendFamilyName string `json:"friendFamilyName"`
	Phone            string `json:"phone"` //手机号带国家号
}
