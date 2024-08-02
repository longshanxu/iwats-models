/*
 * @Author: wzf 1490216271@qq.com
 * @Date: 2024-07-30 20:05:05
 * @LastEditors: wzf 1490216271@qq.com
 * @LastEditTime: 2024-07-30 20:58:36
 * @FilePath: \server\model\iw_qunfa_task\request\iw_qunfa_task.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type IwQunfaTaskSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`

	request.PageInfo
}

// 参数结构体
type IwQunfaTaskQery struct {
	global.GVA_MODEL
	AccountIds    []int  `json:"accountIds" form:"accountIds"`       // WS账号来源
	AccountNum    int    `json:"accountNum" form:"accountNum"`       // 账号总数
	CompletedNum  int    `json:"completedNum" form:"completedNum"`   // 已完成数
	Desc          string `json:"desc" form:"desc"`                   // 描述
	GroupId       []int  `json:"groupId" form:"groupId"`             // 分组ID //粉丝组id
	IntervalStart int    `json:"intervalStart" form:"intervalStart"` // 间隔开始秒
	IntervalStop  int    `json:"intervalStop" form:"intervalStop"`   // 间隔结束秒
	Link          string `json:"link" form:"link"`                   // 图片链接
	MsgType       int    `json:"msgType" form:"msgType"`             // 消息类型1文本2图片3语音
	PicContent    string `json:"picContent" form:"picContent"`       // 图片内容
	PicTilte      string `json:"picTilte" form:"picTilte"`           // 图片标题
	RestMinutes   int    `json:"restMinutes" form:"restMinutes"`     // 休息分钟
	SendData      int    `json:"sendData" form:"sendData"`           // 群发条数数据
	SendRule      int    `json:"sendRule" form:"sendRule"`           // 群发条数规则
	TalkInput     string `json:"TalkInput" form:"TalkInput"`         // 话术/图片内容
	TaskName      string `json:"taskName" form:"taskName"`           // 任务名称
	TaskStart     string `json:"taskStart" form:"taskStart"`         // 任务开始时间
	TaskStatus    int    `json:"taskStatus" form:"taskStatus"`       // 任务状态0.初始化1.已启动未执行 2.执行中 3.执行已完成4执行失败
	TaskStop      string `json:"taskStop" form:"taskStop"`           // 任务结束时间
}
