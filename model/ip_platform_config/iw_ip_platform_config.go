// 自动生成模板IwIpPlatformConfig
package ip_platform_config

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// iwIpPlatformConfig表 结构体  IwIpPlatformConfig
type IwIpPlatformConfig struct {
 global.GVA_MODEL 
      Key  string `json:"key" form:"key" gorm:"column:key;comment:账号名;size:255;"`  //账号名 
      Platform  string `json:"platform" form:"platform" gorm:"column:platform;comment:平台名称;size:255;"`  //平台名称 
      Sign  string `json:"sign" form:"sign" gorm:"column:sign;comment:签名;size:255;"`  //签名 
      Type  string `json:"type" form:"type" gorm:"column:type;comment:平台类型;size:255;"`  //平台类型 
}


// TableName iwIpPlatformConfig表 IwIpPlatformConfig自定义表名 iw_ip_platform_config
func (IwIpPlatformConfig) TableName() string {
  return "iw_ip_platform_config"
}

