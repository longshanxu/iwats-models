// 自动生成模板IwDeviceConfig
package iw_device_config

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// iwDeviceConfig表 结构体  IwDeviceConfig
type IwDeviceConfig struct {
 global.GVA_MODEL 
      Brand  string `json:"brand" form:"brand" gorm:"column:brand;comment:手机品牌;size:255;"`  //手机品牌 
      SystemDeck  *int `json:"systemDeck" form:"systemDeck" gorm:"column:system_deck;comment:1安卓2苹果3鸿蒙;"`  //1安卓2苹果3鸿蒙 
      Model  string `json:"model" form:"model" gorm:"column:model;comment:手机型号;size:255;"`  //手机型号 
      BrandVersion  string `json:"brandVersion" form:"brandVersion" gorm:"column:brand_version;comment:品牌版本;size:255;"`  //品牌版本 
}


// TableName iwDeviceConfig表 IwDeviceConfig自定义表名 iw_device_config
func (IwDeviceConfig) TableName() string {
  return "iw_device_config"
}

