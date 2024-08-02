/*
 * @Author: Wangpengxiang 17600369018@163.com
 * @Date: 2024-07-30 21:08:41
 * @LastEditors: Wangpengxiang 17600369018@163.com
 * @LastEditTime: 2024-08-01 23:04:42
 * @FilePath: \server\model\iw_register_task\iw_register_task.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
// 自动生成模板IwRegisterTask
package iw_register_task

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// iwRegisterTask表 结构体  IwRegisterTask
type IwRegisterTask struct {
	global.GVA_MODEL
	TaskName       string `json:"taskName" form:"taskName" gorm:"column:task_name;comment:任务名称;size:255;"`                       //任务名称
	AccountType    *int   `json:"accountType" form:"accountType" gorm:"column:account_type;comment:账户类型:0.注册1.导入;size:10;"`      //账户类型:0.注册1.导入
	DeviceModel    string `json:"deviceModel" form:"deviceModel" gorm:"column:device_model;comment:接码平台名称;size:255;"`            //接码平台名称
	SystemDeck     *int   `json:"systemDeck" form:"systemDeck" gorm:"column:system_deck;comment:设备系统平台;"`                        //设备系统平台
	TaskStatusCode int    `json:"taskStatusCode" form:"taskStatusCode" gorm:"column:task_status_code;comment:任务状态;"`             //任务状态
	FailedCount    *int   `json:"failedCount" form:"failedCount" gorm:"column:failed_count;comment:执行失败的数量;"`                    //执行失败的数量
	SuccCount      int    `json:"succCount" form:"succCount" gorm:"column:succ_count;comment:执行成功的数量;"`                          //执行成功的数量
	Sort           *int   `json:"sort" form:"sort" gorm:"column:sort;comment:权重;"`                                               //权重
	ConturyinfoId  *int   `json:"conturyinfoId" form:"conturyinfoId" gorm:"column:conturyinfo_id;comment:国家信息配置表id;"`            //国家信息配置表id
	RegisterCount  *int   `json:"registerCount" form:"registerCount" gorm:"column:register_count;comment:注册数量;"`                 //注册数量
	IpDeckId       *int   `json:"ipDeckId" form:"ipDeckId" gorm:"column:ip_deck_id;comment:ip平台id;"`                             //ip平台id
	SmsDeckId      *int   `json:"smsDeckId" form:"smsDeckId" gorm:"column:sms_deck_id;comment:接码平台id;"`                          //接码平台id
	IpGroupId      *int   `json:"ipGroupId" form:"ipGroupId" gorm:"column:ip_group_id;comment:ip分组id;"`                          //ip分组id
	IpGroupName    string `json:"ipGroupName" form:"ipGroupName" gorm:"column:ip_group_name;comment:ip分组名称;size:255;"`           //ip分组名称
	ExecuteDetail  string `json:"executeDetail" form:"executeDetail" gorm:"column:execute_detail;comment:执行细节;size:4294967295;"` //执行细节
	CodingName     string `json:"codingName" form:"codingName" gorm:"column:coding_name;comment:接码平台名称;size:255;"`               //接码平台名称
	IpPlatformName string `json:"ipPlatformName" form:"ipPlatformName" gorm:"column:ip_platform_name;comment:ip平台名称;size:255;"`  //ip平台名称
	IsExecute      int    `json:"isExecute" form:"isExecute" gorm:"column:is_execute;comment:是否在执行;"`                            //是否在执行
	RegisterType   *int   `json:"registerType" form:"registerType" gorm:"column:register_type;comment:注册类型;size:10;"`            //注册类型
}

// TableName iwRegisterTask表 IwRegisterTask自定义表名 iw_register_task
func (IwRegisterTask) TableName() string {
	return "iw_register_task"
}

// iwRegisterTask表 结构体  IwRegisterTask
type IwRegisterTaskRequest struct {
	global.GVA_MODEL
	TaskName       string `json:"taskName" form:"taskName" gorm:"column:task_name;comment:任务名称;size:255;"`                       //任务名称
	AccountType    *int   `json:"accountType" form:"accountType" gorm:"column:account_type;comment:账户类型:0.注册1.导入;size:10;"`      //账户类型:0.注册1.导入
	DeviceModel    string `json:"deviceModel" form:"deviceModel" gorm:"column:device_model;comment:接码平台名称;size:255;"`            //接码平台名称
	SystemDeck     *int   `json:"systemDeck" form:"systemDeck" gorm:"column:system_deck;comment:设备系统平台;"`                        //设备系统平台
	TaskStatusCode int    `json:"taskStatusCode" form:"taskStatusCode" gorm:"column:task_status_code;comment:任务状态;"`             //任务状态
	FailedCount    *int   `json:"failedCount" form:"failedCount" gorm:"column:failed_count;comment:执行失败的数量;"`                    //执行失败的数量
	SuccCount      int    `json:"succCount" form:"succCount" gorm:"column:succ_count;comment:执行成功的数量;"`                          //执行成功的数量
	Sort           *int   `json:"sort" form:"sort" gorm:"column:sort;comment:权重;"`                                               //权重
	ConturyinfoId  *int   `json:"conturyinfoId" form:"conturyinfoId" gorm:"column:conturyinfo_id;comment:国家信息配置表id;"`            //国家信息配置表id
	RegisterCount  *int   `json:"registerCount" form:"registerCount" gorm:"column:register_count;comment:注册数量;"`                 //注册数量
	IpDeckId       *int   `json:"ipDeckId" form:"ipDeckId" gorm:"column:ip_deck_id;comment:ip平台id;"`                             //ip平台id
	SmsDeckId      *int   `json:"smsDeckId" form:"smsDeckId" gorm:"column:sms_deck_id;comment:接码平台id;"`                          //接码平台id
	IpGroupId      *int   `json:"ipGroupId" form:"ipGroupId" gorm:"column:ip_group_id;comment:ip分组id;"`                          //ip分组id
	IpGroupName    string `json:"ipGroupName" form:"ipGroupName" gorm:"column:ip_group_name;comment:ip分组名称;size:255;"`           //ip分组名称
	ExecuteDetail  string `json:"executeDetail" form:"executeDetail" gorm:"column:execute_detail;comment:执行细节;size:4294967295;"` //执行细节
	CodingName     string `json:"codingName" form:"codingName" gorm:"column:coding_name;comment:接码平台名称;size:255;"`               //接码平台名称
	IpPlatformName string `json:"ipPlatformName" form:"ipPlatformName" gorm:"column:ip_platform_name;comment:ip平台名称;size:255;"`  //ip平台名称
	IsExecute      int    `json:"isExecute" form:"isExecute" gorm:"column:is_execute;comment:是否在执行;"`                            //是否在执行
	RegisterType   *int   `json:"registerType" form:"registerType" gorm:"column:register_type;comment:注册类型;size:10;"`            //注册类型

	ConturyCode     string `json:"conturyCode" form:"conturyCode" gorm:"column:contury_code;comment:国家代码;size:10;"`                //国家代码
	ConturyLanguage string `json:"conturyLanguage" form:"conturyLanguage" gorm:"column:contury_language;comment:国家语言;size:255;"`   //国家语言
	ConturyZoneCode *int   `json:"conturyZoneCode" form:"conturyZoneCode" gorm:"column:contury_zone_code;comment:国家区域号码;size:10;"` //国家区域号码
}

// TableName iwRegisterTask表 IwRegisterTask自定义表名 iw_register_task
func (IwRegisterTaskRequest) TableName() string {
	return "iw_register_task"
}
