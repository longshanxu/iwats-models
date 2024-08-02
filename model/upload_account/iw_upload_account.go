/*
 * @Author: Wangpengxiang 17600369018@163.com
 * @Date: 2024-07-30 21:17:56
 * @LastEditors: Wangpengxiang 17600369018@163.com
 * @LastEditTime: 2024-07-31 22:23:37
 * @FilePath: \server\model\upload_account\iw_upload_account.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
// 自动生成模板IwUploadAccount
package upload_account

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// iwUploadAccount表 结构体  IwUploadAccount
type IwUploadAccount struct {
	global.GVA_MODEL
	Brand               string `json:"brand" form:"brand" gorm:"column:brand;comment:手机品牌;size:255;"`                                                      //手机品牌
	CountryCode         string `json:"countryCode" form:"countryCode" gorm:"column:country_code;comment:国家代码;size:255;"`                                   //国家代码
	CountryLanguage     string `json:"countryLanguage" form:"countryLanguage" gorm:"column:country_language;comment:国家语言;size:255;"`                       //国家语言
	CountryName         string `json:"countryName" form:"countryName" gorm:"column:country_name;comment:ip国家名称;size:255;"`                                 //ip国家名称
	CountryZoneCode     *int   `json:"countryZoneCode" form:"countryZoneCode" gorm:"column:country_zone_code;comment:国家区号;"`                               //国家区号
	DeviceName          string `json:"deviceName" form:"deviceName" gorm:"column:device_name;comment:设备名称;size:255;"`                                      //设备名称
	GroupId             *int   `json:"groupId" form:"groupId" gorm:"column:group_id;comment:分组id;size:10;"`                                                //分组id
	GroupName           string `json:"groupName" form:"groupName" gorm:"column:group_name;comment:分组名称;size:255;"`                                         //分组名称
	Ip                  string `json:"ip" form:"ip" gorm:"column:ip;comment:ip;size:255;"`                                                                 //ip
	IpCity              string `json:"ipCity" form:"ipCity" gorm:"column:ip_city;comment:ip城市;size:255;"`                                                  //ip城市
	IpName              string `json:"ipName" form:"ipName" gorm:"column:ip_name;comment:ip平台;size:1255;"`                                                 //ip平台
	IsInitStatus        int    `json:"isInitStatus" form:"isInitStatus" gorm:"column:is_init_status;comment:准备状态:1.未绑定设备，2.绑定设备中，3.未绑定IP，4.准备就绪;size:10;"` //准备状态:1.未绑定设备，2.绑定设备中，3.未绑定IP，4.准备就绪
	IsOnline            *int   `json:"isOnline" form:"isOnline" gorm:"column:is_online;comment:在线状态:1.离线，2在线3.封号;"`                                        //在线状态:1.离线，2在线3.封号
	Lc                  string `json:"lc" form:"lc" gorm:"column:lc;comment:国家代码;size:255;"`                                                               //国家代码
	Lg                  string `json:"lg" form:"lg" gorm:"column:lg;comment:国家语言;size:255;"`                                                               //国家语言
	Mcc                 string `json:"mcc" form:"mcc" gorm:"column:mcc;comment:;size:255;"`                                                                //mcc字段
	MessagePrivateKey   string `json:"messagePrivateKey" form:"messagePrivateKey" gorm:"column:message_private_key;comment:消息私钥;"`                         //消息私钥
	MessagePublicKey    string `json:"messagePublicKey" form:"messagePublicKey" gorm:"column:message_public_key;comment:消息公钥;"`                            //消息公钥
	Mnc                 string `json:"mnc" form:"mnc" gorm:"column:mnc;comment:;size:255;"`                                                                //mnc字段
	Model               string `json:"model" form:"model" gorm:"column:model;comment:设备型号;size:255;"`                                                      //设备型号
	NetworkOperatorName string `json:"networkOperatorName" form:"networkOperatorName" gorm:"column:network_operator_name;comment:网络运营商名称;size:255;"`       //网络运营商名称
	Phone               string `json:"phone" form:"phone" gorm:"column:phone;comment:账号;size:255;"`                                                        //账号
	PhoneArea           string `json:"phoneArea" form:"phoneArea" gorm:"column:phone_area;comment:手机号区号;size:255;"`                                        //手机号区号
	PhoneCountry        string `json:"phoneCountry" form:"phoneCountry" gorm:"column:phone_country;comment:手机号国家;size:255;"`                               //手机号国家
	PhoneId             string `json:"phoneId" form:"phoneId" gorm:"column:phone_id;comment:号码id;"`                                                        //号码id
	PrivateKey          string `json:"privateKey" form:"privateKey" gorm:"column:private_key;comment:私钥;"`                                                 //私钥
	PublicKey           string `json:"publicKey" form:"publicKey" gorm:"column:public_key;comment:公钥;"`                                                    //公钥
	PushName            string `json:"pushName" form:"pushName" gorm:"column:push_name;comment:;size:255;"`                                                //pushName字段
	RegisterType        *int   `json:"registerType" form:"registerType" gorm:"column:register_type;comment:注册类型:1.个人般,2.商业版;size:10;"`                     //注册类型:1.个人般,2.商业版
	RegusterStatus      *int   `json:"regusterStatus" form:"regusterStatus" gorm:"column:reguster_status;comment:1.注册成功，2.注册失败;"`                          //1.注册成功，2.注册失败
	Socks5              string `json:"socks5" form:"socks5" gorm:"column:socks5;comment:socks5;size:255;"`                                                 //socks5
	TaskAllId           string `json:"taskAllId" form:"taskAllId" gorm:"column:task_all_id;comment:任务执行ID;size:10;"`                                       //任务执行ID
	TaskId              *int   `json:"taskId" form:"taskId" gorm:"column:task_id;comment:任务ID;size:10;"`                                                   //任务ID
	TaskName            string `json:"taskName" form:"taskName" gorm:"column:task_name;comment:任务名称;size:255;"`                                            //任务名称
}

// TableName iwUploadAccount表 IwUploadAccount自定义表名 iw_upload_account
func (IwUploadAccount) TableName() string {
	return "iw_upload_account"
}
