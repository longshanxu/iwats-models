/*
 * @Author: wzf 1490216271@qq.com
 * @Date: 2024-07-31 14:56:44
 * @LastEditors: wzf 1490216271@qq.com
 * @LastEditTime: 2024-07-31 18:58:08
 * @FilePath: \iwats-admin\server\model\iw_message\iw_message.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
// 自动生成模板IwMessage
package iw_message

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// iwMessage表 结构体  IwMessage
type IwMessage struct {
	global.GVA_MODEL
	AccountId     int    `json:"accountId" form:"accountId" gorm:"column:account_id;comment:自注账号id;"`                   //自注账号id
	MessageId     string `json:"messageId" form:"messageId" gorm:"column:message_id;comment:当前这条信息的id;size:500;"`       //当前这条信息的id
	IsSelf        int    `json:"isSelf" form:"isSelf" gorm:"column:is_self;comment:是否自己发送;"`                            //是否自己发送
	FriendId      int    `json:"friendId" form:"friendId" gorm:"column:friend_id;comment:添加用户账号id;"`                    //添加用户账号id
	Phone         string `json:"phone" form:"phone" gorm:"column:phone;comment:手机号;size:255;"`                          //手机号
	Status        int    `json:"status" form:"status" gorm:"column:status;comment:消息状态1已读2未读;"`                         //消息状态1已读2未读
	Content       string `json:"content" form:"content" gorm:"column:content;comment:消息内容;size:255;"`                   //消息内容
	MessageStatus int    `json:"messageStatus" form:"messageStatus" gorm:"column:message_status;comment:1文本2图片3语音4视频;"` //1文本2图片3语音4视频
}

// TableName iwMessage表 IwMessage自定义表名 iw_message
func (IwMessage) TableName() string {
	return "iw_message"
}
