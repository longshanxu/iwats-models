/*
 * @Author: Wangpengxiang 17600369018@163.com
 * @Date: 2024-07-30 15:44:30
 * @LastEditors: Wangpengxiang 17600369018@163.com
 * @LastEditTime: 2024-08-01 18:37:48
 * @FilePath: \server\model\ip_config\iw_ip_config.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
// 自动生成模板IwIpConfig
package ip_config

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// iwIpConfig表 结构体  IwIpConfig
type IwIpConfig struct {
	global.GVA_MODEL
	City       string `json:"city" form:"city" gorm:"column:city;comment:城市;size:255;"`                      //城市
	CmiId      string `json:"cmiId" form:"cmiId" gorm:"column:cmiId;comment:账号分类的channel个人号码;size:191;"`     //账号分类的channel个人号码
	Country    string `json:"country" form:"country" gorm:"column:country;comment:国家地区名称;size:255;"`         //国家地区名称
	Domain     string `json:"domain" form:"domain" gorm:"column:domain;comment:域名;size:255;"`                //域名
	EndTime    string `json:"endTime" form:"endTime" gorm:"column:endTime;comment:到期时间;size:191;"`           //到期时间
	Ip         string `json:"ip" form:"ip" gorm:"column:ip;comment:ip地址;size:255;"`                          //ip地址
	IpAccount  string `json:"ipAccount" form:"ipAccount" gorm:"column:ip_account;comment:ip账号;size:255;"`    //ip账号
	IpGroupId  *int   `json:"ipGroupId" form:"ipGroupId" gorm:"column:ip_group_id;comment:分组id;size:10;"`    //分组id
	IpPassword string `json:"ipPassword" form:"ipPassword" gorm:"column:ip_password;comment:ip密码;size:255;"` //ip密码
	Port       string `json:"port" form:"port" gorm:"column:port;comment:端口;size:255;"`                      //端口
	Remark     string `json:"remark" form:"remark" gorm:"column:remark;comment:评论;size:255;"`                //评论
	Type       string `json:"type" form:"type" gorm:"column:type;comment:类型;size:255;"`                      //类型
}

// TableName iwIpConfig表 IwIpConfig自定义表名 iw_ip_config
func (IwIpConfig) TableName() string {
	return "iw_ip_config"
}

type RequestData struct {
	global.GVA_MODEL
	Key        string `json:"key"`        // 账号名(IPIPGO官网的登陆账号名,用户唯一标识)
	Sign       string `json:"sign"`       // 签名(IPIPGO官网)
	Expiry     string `json:"expiry"`     // 是否过期
	AddressStr string `json:"addressStr"` // 国家地区名称
	IpGroupId  int    `json:"ipGroupId"`  // 外键分组ID 	// 文件路径
}


// iwIpConfig表 结构体  IwIpConfig
type IwIpConfigResponse struct {
	global.GVA_MODEL
	City       string `json:"city" form:"city" gorm:"column:city;comment:城市;size:255;"`                      //城市
	CmiId      string `json:"cmiId" form:"cmiId" gorm:"column:cmiId;comment:账号分类的channel个人号码;size:191;"`     //账号分类的channel个人号码
	Country    string `json:"country" form:"country" gorm:"column:country;comment:国家地区名称;size:255;"`         //国家地区名称
	Domain     string `json:"domain" form:"domain" gorm:"column:domain;comment:域名;size:255;"`                //域名
	EndTime    string `json:"endTime" form:"endTime" gorm:"column:endTime;comment:到期时间;size:191;"`           //到期时间
	Ip         string `json:"ip" form:"ip" gorm:"column:ip;comment:ip地址;size:255;"`                          //ip地址
	IpAccount  string `json:"ipAccount" form:"ipAccount" gorm:"column:ip_account;comment:ip账号;size:255;"`    //ip账号
	IpGroupId  *int   `json:"ipGroupId" form:"ipGroupId" gorm:"column:ip_group_id;comment:分组id;size:10;"`    //分组id
	IpPassword string `json:"ipPassword" form:"ipPassword" gorm:"column:ip_password;comment:ip密码;size:255;"` //ip密码
	Port       string `json:"port" form:"port" gorm:"column:port;comment:端口;size:255;"`                      //端口
	Remark     string `json:"remark" form:"remark" gorm:"column:remark;comment:评论;size:255;"`                //评论
	Type       string `json:"type" form:"type" gorm:"column:type;comment:类型;size:255;"`                      //类型
	Socks5    string `json:"socks5" form:"socks5" gorm:"column:socks5;comment:socks5;size:255;"`              //socks5
}

// TableName iwIpConfig表 IwIpConfig自定义表名 iw_ip_config
func (IwIpConfigResponse) TableName() string {
	return "iw_ip_config"
}