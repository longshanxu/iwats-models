/*
 * @Author: wzf 1490216271@qq.com
 * @Date: 2024-07-31 14:56:44
 * @LastEditors: wzf 1490216271@qq.com
 * @LastEditTime: 2024-07-31 16:51:11
 * @FilePath: \iwats-admin\server\model\iw_message\request\iw_message.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type IwMessageSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`

	request.PageInfo
}

type MesListReq struct {
	MyAccountId int `json:"myAccountId" form:"myAccountId"`
	request.PageInfo
}

type MesListRes struct {
	AddUserId   int    `json:"addUserId" `   //添加用户账号id
	UserName    string `json:"userName" `    //城市
	MyAccountId int    `json:"myAccountId" ` //自注账号id
	Phone       string `json:"phone"`
	ReadCount   int    `json:"readCount"`
}

// send message
type SendMessageReq struct {
	MyAccountId int    `json:"myAccountId,omitempty" form:"myAccountId"`
	Text        string `json:"text,omitempty" form:"text"`
	AddUserId   int    `json:"addUserId,omitempty" form:"addUserId"`
}

// api send message
type ApiSendMessage struct {
	From string `json:"from"`
	Text string `json:"text"`
	To   string `json:"to"`
}

type RefreshMessageReq struct {
	MyAccountId int `json:"myAccountId,omitempty" form:"myAccountId"`
	AddUserId   int `json:"addUserId,omitempty" form:"addUserId"`
}

type SenMessageRes struct {
	Data    string `json:"data"`
	Message string `json:"message"`
}
