// 自动生成模板IwIpGroup
package ip_group

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// iwIpGroup表 结构体  IwIpGroup
type IwIpGroup struct {
 global.GVA_MODEL 
      IpGroupName  string `json:"ipGroupName" form:"ipGroupName" gorm:"column:ip_group_name;comment:分组名称;size:255;"`  //分组名称 
      IpPlatformId  *int `json:"ipPlatformId" form:"ipPlatformId" gorm:"column:ip_platform_id;comment:;size:10;"`  //ipPlatformId字段 
      Sort  *int `json:"sort" form:"sort" gorm:"column:sort;comment:权重;size:10;"`  //权重 
}


// TableName iwIpGroup表 IwIpGroup自定义表名 iw_ip_group
func (IwIpGroup) TableName() string {
  return "iw_ip_group"
}

