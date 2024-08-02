/*
 * @Author: Wangpengxiang 17600369018@163.com
 * @Date: 2024-08-01 14:52:28
 * @LastEditors: Wangpengxiang 17600369018@163.com
 * @LastEditTime: 2024-08-01 14:52:46
 * @FilePath: \server\model\iw_sim_config\iw_sim_config.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */

package iw_sim_config

import "gorm.io/gorm"

type IwSimConfig struct {
	gorm.Model
	CountryZoneCode string `json:"countryZoneCode" form:"countryZoneCode" gorm:"column:country_zone_code;comment:国家区域代码;size:255;"` //国家区域代码
	Country         string `json:"country" form:"country" gorm:"column:country;comment:国家;size:100;"`                               //国家
	Mcc             string `json:"mcc" form:"mcc" gorm:"column:mcc;comment:;size:10;"`                                              //mcc字段
	Mnc             string `json:"mnc" form:"mnc" gorm:"column:mnc;comment:;size:10;"`                                              //mnc字段
	Brand           string `json:"brand" form:"brand" gorm:"column:brand;comment:运营商品牌;size:100;"`                                  //运营商品牌
	Operator        string `json:"operator" form:"operator" gorm:"column:operator;comment:;size:100;"`                              //operator字段
	Status          string `json:"status" form:"status" gorm:"column:status;comment:使用状态;size:50;"`                                 //使用状态
	Frequency       string `json:"frequency" form:"frequency" gorm:"column:frequency;comment:频段;size:100;"`         
}

// TableName nvMessageList表 NvMessageList自定义表名 iw_sim_config
func (IwSimConfig) TableName() string {
	return "iw_sim_config"
}
