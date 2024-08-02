// 自动生成模板IwQunfaRecord
package iw_qunfa_record

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// iwQunfaRecord表 结构体  IwQunfaRecord
type IwQunfaRecord struct {
	global.GVA_MODEL
	FansJid        string `json:"fansJid" form:"fansJid" gorm:"column:fans_jid;comment:粉丝好友jid;size:4294967295;"`                       //粉丝好友jid
	FansJidFailed  string `json:"fansJidFailed" form:"fansJidFailed" gorm:"column:fans_jid_failed;comment:粉丝好友jid失败;size:4294967295;"`  //粉丝好友jid失败
	Phone          string `json:"phone" form:"phone" gorm:"column:phone;comment:手机号;size:255;"`                                         //手机号
	QunfaTaskId    int    `json:"qunfaTaskId" form:"qunfaTaskId" gorm:"column:qunfa_task_id;comment:群发任务表id;size:10;"`                  //群发任务表id
	SendFans       string `json:"sendFans" form:"sendFans" gorm:"column:send_fans;comment:发送成功粉丝号;size:4294967295;"`                    //发送成功粉丝号
	SendFansFailed string `json:"sendFansFailed" form:"sendFansFailed" gorm:"column:send_fans_failed;comment:发送失败粉丝号;size:4294967295;"` //发送失败粉丝号
}

// TableName iwQunfaRecord表 IwQunfaRecord自定义表名 iw_qunfa_record
func (IwQunfaRecord) TableName() string {
	return "iw_qunfa_record"
}
