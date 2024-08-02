// 自动生成模板IwGroup
package iw_group

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// iwGroup表 结构体  IwGroup
type IwGroup struct {
 global.GVA_MODEL 
      GroupType  *int `json:"groupType" form:"groupType" gorm:"column:group_type;comment:组类型1:导入2:注册3:粉丝;size:10;"`  //组类型1:导入2:注册3:粉丝 
      Name  string `json:"name" form:"name" gorm:"column:name;comment:分组名称;size:255;"`  //分组名称 
      Sort  *int `json:"sort" form:"sort" gorm:"column:sort;comment:权重;size:10;"`  //权重 
}


// TableName iwGroup表 IwGroup自定义表名 iw_group
func (IwGroup) TableName() string {
  return "iw_group"
}

