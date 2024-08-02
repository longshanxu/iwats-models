// 自动生成模板IwConturyinfoConfig
package iw_contury_config

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// iwConturyinfoConfig表 结构体  IwConturyinfoConfig
type IwConturyinfoConfig struct {
 global.GVA_MODEL 
      ConturyName  string `json:"conturyName" form:"conturyName" gorm:"column:contury_name;comment:国家名称;size:255;"`  //国家名称 
      ConturyCode  string `json:"conturyCode" form:"conturyCode" gorm:"column:contury_code;comment:国家英文代码;size:255;"`  //国家英文代码 
      SmsConturyCode  *int `json:"smsConturyCode" form:"smsConturyCode" gorm:"column:sms_contury_code;comment:俄罗斯接码平台国家代码;"`  //俄罗斯接码平台国家代码 
      ConturyLanguage  string `json:"conturyLanguage" form:"conturyLanguage" gorm:"column:contury_language;comment:国家语言;size:255;"`  //国家语言 
      ConturyZoneNo  *int `json:"conturyZoneNo" form:"conturyZoneNo" gorm:"column:contury_zone_no;comment:国家区号;"`  //国家区号 
      Sort  *int `json:"sort" form:"sort" gorm:"column:sort;comment:权重;"`  //权重 
}


// TableName iwConturyinfoConfig表 IwConturyinfoConfig自定义表名 iw_conturyinfo_config
func (IwConturyinfoConfig) TableName() string {
  return "iw_conturyinfo_config"
}

