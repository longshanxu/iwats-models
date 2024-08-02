package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type IwQunfaDetailsSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	QunfaTaskId    int        `json:"qunfaTaskId" form:"qunfaTaskId"`
	AccoutIds      []int      `json:"accoutIds" form:"accoutIds"`
	Phone          string     `json:"phone" form:"phone"`

	request.PageInfo
}

type IwQunfaDetailsRes struct {
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
	AccountStatus  int    `json:"accountStatus" form:"accountStatus" gorm:"column:account_status;comment:账号状态;"`                 //账号状态
}

// 群发任务统计返回
type GetCollectRes struct {
	AccountCount    int     `json:"accountCount" form:"accountCount"`       //账号总数
	CompleteAccount int     `json:"completeAccount" form:"completeAccount"` //已完成账号
	TargetFansCount int     `json:"targetFansCount" form:"targetFansCount"` //目标粉丝书
	SuccExpandFans  int     `json:"succExpandFans" form:"succExpandFans"`   //成功推送粉丝
	AlreadyReadRate float64 `json:"alreadyReadRate" form:"alreadyReadRate"` //已经读率
	ReplyRate       float64 `json:"replyRate" form:"replyRate"`             //回复率
	AcceptMsgCount  int     `json:"acceptMsgCount" fom:"acceptMsgCount"`    //接收消息总数
}
