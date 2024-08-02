/*
 * @Author: wzf 1490216271@qq.com
 * @Date: 2024-07-30 20:05:05
 * @LastEditors: wzf 1490216271@qq.com
 * @LastEditTime: 2024-07-30 23:49:46
 * @FilePath: \server\model\iw_qunfa_task\iw_qunfa_task.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
// 自动生成模板IwQunfaTask
package iw_qunfa_task

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// iwQunfaTask表 结构体  IwQunfaTask
type IwQunfaTask struct {
	global.GVA_MODEL
	AccountIds    string `json:"accountIds" form:"accountIds" gorm:"column:account_ids;comment:WS账号来源;size:4294967295;"`               //WS账号来源
	AccountNum    int    `json:"accountNum" form:"accountNum" gorm:"column:account_num;comment:账号总数;size:19;"`                         //账号总数
	CompletedNum  int    `json:"completedNum" form:"completedNum" gorm:"column:completed_num;comment:已完成数;size:19;"`                   //已完成数
	Desc          string `json:"desc" form:"desc" gorm:"column:desc;comment:描述;size:1000;"`                                            //描述
	GroupId       string `json:"groupId" form:"groupId" gorm:"column:group_id;comment:分组ID;"`                                          //分组ID
	IntervalStart int    `json:"intervalStart" form:"intervalStart" gorm:"column:interval_start;comment:间隔开始秒;"`                       //间隔开始秒
	IntervalStop  int    `json:"intervalStop" form:"intervalStop" gorm:"column:interval_stop;comment:间隔结束秒;"`                          //间隔结束秒
	Link          string `json:"link" form:"link" gorm:"column:link;comment:图片链接;size:4294967295;"`                                    //图片链接
	MsgType       int    `json:"msgType" form:"msgType" gorm:"column:msg_type;comment:消息类型1文本2图片3语音;size:10;"`                         //消息类型1文本2图片3语音
	PicContent    string `json:"picContent" form:"picContent" gorm:"column:pic_content;comment:图片内容;size:4294967295;"`                 //图片内容
	PicTilte      string `json:"picTilte" form:"picTilte" gorm:"column:pic_tilte;comment:图片标题;size:255;"`                              //图片标题
	RestMinutes   int    `json:"restMinutes" form:"restMinutes" gorm:"column:rest_minutes;comment:休息分钟;"`                              //休息分钟
	SendData      int    `json:"sendData" form:"sendData" gorm:"column:send_data;comment:群发条数数据;"`                                     //群发条数数据
	SendRule      int    `json:"sendRule" form:"sendRule" gorm:"column:send_rule;comment:群发条数规则;"`                                     //群发条数规则
	TalkInput     string `json:"talkInput" form:"talkInput" gorm:"column:talk_input;comment:话术/图片内容;size:1000;"`                       //话术/图片内容
	TaskName      string `json:"taskName" form:"taskName" gorm:"column:task_name;comment:任务名称;size:255;"`                              //任务名称
	TaskStart     string `json:"taskStart" form:"taskStart" gorm:"column:task_start;comment:任务开始时间;size:191;"`                         //任务开始时间
	TaskStatus    int    `json:"taskStatus" form:"taskStatus" gorm:"column:task_status;comment:任务状态0.初始化1.已启动未执行 2.执行中 3.执行已完成4执行失败;"` //任务状态0.初始化1.已启动未执行 2.执行中 3.执行已完成4执行失败
	TaskStop      string `json:"taskStop" form:"taskStop" gorm:"column:task_stop;comment:任务结束时间;size:191;"`                            //任务结束时间
}

// TableName iwQunfaTask表 IwQunfaTask自定义表名 iw_qunfa_task
func (IwQunfaTask) TableName() string {
	return "iw_qunfa_task"
}
