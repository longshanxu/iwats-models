// 自动生成模板IwUploadAccountTask
package upload_account_task

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// iwUploadAccountTask表 结构体  IwUploadAccountTask
type IwUploadAccountTask struct {
	global.GVA_MODEL
	AccountType     int    `json:"accountType" form:"accountType" gorm:"column:account_type;comment:账户类型:0.注册1.导入;size:10;"`       //账户类型:0.注册1.导入
	ConturyCode     string `json:"conturyCode" form:"conturyCode" gorm:"column:contury_code;comment:国家代码;size:10;"`                //国家代码
	ConturyLanguage string `json:"conturyLanguage" form:"conturyLanguage" gorm:"column:contury_language;comment:国家语言;size:255;"`   //国家语言
	ConturyZoneCode int    `json:"conturyZoneCode" form:"conturyZoneCode" gorm:"column:contury_zone_code;comment:国家区域号码;"`         //国家区域号码
	DevcieDeck      string `json:"devcieDeck" form:"devcieDeck" gorm:"column:devcie_deck;comment:设备类型;size:255;"`                  //设备类型
	ExecuteDetail   string `json:"executeDetail" form:"executeDetail" gorm:"column:execute_detail;comment:执行细节;size:4294967295;"`  //执行细节
	FailedCount     int    `json:"failedCount" form:"failedCount" gorm:"column:failed_count;comment:执行失败的数量;"`                     //执行失败的数量
	GroupId         int    `json:"groupId" form:"groupId" gorm:"column:group_id;comment:分组id;"`                                    //分组id
	GroupName       string `json:"groupName" form:"groupName" gorm:"column:group_name;comment:分组名称;size:255;"`                     //分组名称
	IsExecute       int    `json:"isExecute" form:"isExecute" gorm:"column:is_execute;comment:是否在执行;"`                             //是否在执行
	RegisterType    int    `json:"registerType" form:"registerType" gorm:"column:register_type;comment:注册类型:1.个人般,2.商业版;size:10;"` //注册类型:1.个人般,2.商业版
	Remark          string `json:"remark" form:"remark" gorm:"column:remark;comment:备注;"`                                          //备注
	Sort            int    `json:"sort" form:"sort" gorm:"column:sort;comment:权重;"`                                                //权重
	SuccCount       int    `json:"succCount" form:"succCount" gorm:"column:succ_count;comment:执行成功的数量;"`                           //执行成功的数量
	SystemDeck      int    `json:"systemDeck" form:"systemDeck" gorm:"column:system_deck;comment:设备系统平台;"`                         //设备系统平台
	TaskName        string `json:"taskName" form:"taskName" gorm:"column:task_name;comment:任务名称;size:255;"`                        //任务名称
	TaskStatusCode  int    `json:"taskStatusCode" form:"taskStatusCode" gorm:"column:task_status_code;comment:任务状态;"`              //任务状态
	UploadCount     int    `json:"uploadCount" form:"uploadCount" gorm:"column:upload_count;comment:注册数量;"`                        //注册数量
	UploadType      int    `json:"uploadType" form:"uploadType" gorm:"column:upload_type;comment:导入类型1.六段协议号;size:10;"`            //导入类型1.六段协议号
}

// TableName iwUploadAccountTask表 IwUploadAccountTask自定义表名 iw_upload_account_task
func (IwUploadAccountTask) TableName() string {
	return "iw_upload_account_task"
}
