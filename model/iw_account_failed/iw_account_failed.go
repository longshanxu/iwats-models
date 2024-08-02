/*
 * @Author: Wangpengxiang 17600369018@163.com
 * @Date: 2024-08-02 02:29:58
 * @LastEditors: Wangpengxiang 17600369018@163.com
 * @LastEditTime: 2024-08-02 02:30:26
 * @FilePath: \server\model\iw_account_failed\iw_account_failed.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */

/*
 * @Author: Wangpengxiang 17600369018@163.com
 * @Date: 2024-07-30 21:34:19
 * @LastEditors: Wangpengxiang 17600369018@163.com
 * @LastEditTime: 2024-08-01 11:02:40
 * @FilePath: \server\model\iw_account\iw_account.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
// 自动生成模板IwAccount
package iw_account_failed

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// iwAccount表 结构体  IwAccount
type IwAccountResponse struct {
	global.GVA_MODEL
	AccountType              *int   `json:"accountType" form:"accountType" gorm:"column:account_type;comment:账号来源:1.自注账户，2.导入账户;size:10;"`                             //账号来源:1.自注账户，2.导入账户
	GroupId                  *int   `json:"groupId" form:"groupId" gorm:"column:group_id;comment:分组ID;"`                                                               //分组ID
	IpConturyinfoId          *int   `json:"ipConturyinfoId" form:"ipConturyinfoId" gorm:"column:ip_conturyinfo_id;comment:ip注册国家信息;"`                                  //ip注册国家信息
	IpId                     *int   `json:"ipId" form:"ipId" gorm:"column:ip_id;comment:ipID;"`                                                                        //ipID
	IsOnline                 *int   `json:"isOnline" form:"isOnline" gorm:"column:is_online;comment:在线状态:1.离线，2在线3.封号;"`                                               //在线状态:1.离线，2在线3.封号
	Nickname                 string `json:"nickname" form:"nickname" gorm:"column:nickname;comment:用户昵称;size:255;"`                                                    //用户昵称
	Phone                    string `json:"phone" form:"phone" gorm:"column:phone;comment:手机号;size:255;"`                                                              //手机号
	PhoneConturyinfoId       *int   `json:"phoneConturyinfoId" form:"phoneConturyinfoId" gorm:"column:phone_conturyinfo_id;comment:手机国家信息id;"`                         //手机国家信息id
	RegisteriwIpConfigStatus *int   `json:"registeriwIpConfigStatus" form:"registeriwIpConfigStatus" gorm:"column:registeriw_ip_config_status;comment:1.注册成功，2.注册失败;"` //1.注册成功，2.注册失败
	Type                     *int   `json:"type" form:"type" gorm:"column:type;comment:类型：1.个人版，2.企业版;"`                                                               //类型：1.个人版，2.企业版

	IpCity       string `json:"ipCity" form:"ipCity" gorm:"column:ip_city;comment:ip城市;size:255;"`                    //ip城市
	IpCountry    string `json:"ipCountry" form:"ipCountry" gorm:"column:ip_country;comment:ip国家;size:255;"`           //ip国家
	Ip           string `json:"ip" form:"ip" gorm:"column:ip;comment:ip地址;size:255;"`                                 //ip地址
	Socks5       string `json:"socks5" form:"socks5" gorm:"column:socks5;comment:socks5;size:255;"`                   // socks5
	IpName       string `json:"ipName" form:"ipName" gorm:"column:ip_name;comment:ip;size:1255;"`                     //ip名称
	PhoneCc      *int   `json:"phoneCc" form:"phoneCc" gorm:"column:phone_cc;comment:手机国家代码;size:255;"`               //手机国家代码
	PhoneCountry string `json:"phoneCountry" form:"phoneCountry" gorm:"column:phone_country;comment:手机国家;size:255;"`  //手机国家
	IsInitStatus int    `json:"isInitStatus" form:"isInitStatus" gorm:"column:is_init_status;comment:是否初始化;size:10;"` // 是否初始化 1.未初始化，2.初始化中，3.已初始化
	GroupName    string `json:"groupName" form:"groupName" gorm:"column:group_name;comment:分组名称;size:255;"`           //分组名称
}

// TableName iwAccount表 IwAccount自定义表名 iw_account
func (IwAccountResponse) TableName() string {
	return "iw_account_failed"
}
