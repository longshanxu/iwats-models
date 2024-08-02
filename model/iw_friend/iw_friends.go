// 自动生成模板IwFriends
package iw_friend

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// iwFriends表 结构体  IwFriends
type IwFriends struct {
 global.GVA_MODEL 
      FansJid  string `json:"fansJid" form:"fansJid" gorm:"column:fans_jid;comment:好友手机号对应ws;size:255;"`  //好友手机号对应ws 
      Name  string `json:"name" form:"name" gorm:"column:name;comment:昵称;size:255;"`  //昵称 
      AccountId  *int `json:"accountId" form:"accountId" gorm:"column:account_id;comment:自注账号id对应自注册账号的id;size:10;"`  //自注账号id对应自注册账号的id 
      Phone  string `json:"phone" form:"phone" gorm:"column:phone;comment:好友手机号;size:255;"`  //好友手机号 
}


// TableName iwFriends表 IwFriends自定义表名 iw_friends
func (IwFriends) TableName() string {
  return "iw_friends"
}

