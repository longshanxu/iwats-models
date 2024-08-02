// 自动生成模板IwJiemaPlatform
package jiema_platform

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// iwJiemaPlatform表 结构体  IwJiemaPlatform
type IwJiemaPlatform struct {
 global.GVA_MODEL 
      Name  string `json:"name" form:"name" gorm:"column:name;comment:平台名称;size:255;"`  //平台名称 
      SmsDeckUrl  string `json:"smsDeckUrl" form:"smsDeckUrl" gorm:"column:sms_deck_url;comment:接码平台地址;size:500;"`  //接码平台地址 
      SecretKey  string `json:"secretKey" form:"secretKey" gorm:"column:secret_key;comment:api密钥;size:1000;"`  //api密钥 
      Balance  *int `json:"balance" form:"balance" gorm:"column:balance;comment:余额;size:19;"`  //余额 
      Sort  *int `json:"sort" form:"sort" gorm:"column:sort;comment:权重;"`  //权重 
}


// TableName iwJiemaPlatform表 IwJiemaPlatform自定义表名 iw_jiema_platform
func (IwJiemaPlatform) TableName() string {
  return "iw_jiema_platform"
}

