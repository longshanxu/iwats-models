// 自动生成模板IwFans
package fans

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// iwFans表 结构体  IwFans
type IwFans struct {
	global.GVA_MODEL
	FansJid string `json:"fansJid" form:"fansJid" gorm:"column:fans_jid;comment:粉丝的wsID;size:255;"` //粉丝的wsID
	GroupId *int   `json:"groupId" form:"groupId" gorm:"column:group_id;comment:分组ID;"`             //分组ID
	Phone   string `json:"phone" form:"phone" gorm:"column:phone;comment:手机号;size:255;"`            //手机号
	
}

// TableName iwFans表 IwFans自定义表名 iw_fans
func (IwFans) TableName() string {
	return "iw_fans"
}
