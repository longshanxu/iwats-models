// 自动生成模板IwQunfaDetails
package iw_qunfa_detail

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// iwQunfaDetails表 结构体  IwQunfaDetails
type IwQunfaDetails struct {
	global.GVA_MODEL
	CompletedNum   int    `json:"completedNum" form:"completedNum" gorm:"column:completed_num;comment:完成数;size:19;"`             //完成数
	ExecuteDetail  string `json:"executeDetail" form:"executeDetail" gorm:"column:execute_detail;comment:执行细节;size:4294967295;"` //执行细节
	FansReceiveNum int    `json:"fansReceiveNum" form:"fansReceiveNum" gorm:"column:fans_receive_num;comment:粉丝接收账号消息总数;"`       //粉丝接收账号消息总数
	Phone          string `json:"phone" form:"phone" gorm:"column:phone;comment:手机号;size:255;"`                                  //手机号
	QunfaTaskId    int    `json:"qunfaTaskId" form:"qunfaTaskId" gorm:"column:qunfa_task_id;comment:任务ID;"`                      //任务ID
	ReadRate       int    `json:"readRate" form:"readRate" gorm:"column:read_rate;comment:已读率;size:10;"`                         //已读率
	ReplyRate      int    `json:"replyRate" form:"replyRate" gorm:"column:reply_rate;comment:回复率;size:10;"`                      //回复率
	SendNum        int    `json:"sendNum" form:"sendNum" gorm:"column:send_num;comment:群发数;size:19;"`                            //群发数
	TaskStatus     int    `json:"taskStatus" form:"taskStatus" gorm:"column:task_status;comment:任务状态1.未执行 2.执行中 3.已完4失败;"`       //任务状态1.未执行 2.执行中 3.已完4失败
}

// TableName iwQunfaDetails表 IwQunfaDetails自定义表名 iw_qunfa_details
func (IwQunfaDetails) TableName() string {
	return "iw_qunfa_details"
}
