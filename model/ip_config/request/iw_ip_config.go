/*
 * @Author: Wangpengxiang 17600369018@163.com
 * @Date: 2024-07-30 15:44:30
 * @LastEditors: Wangpengxiang 17600369018@163.com
 * @LastEditTime: 2024-07-30 20:34:22
 * @FilePath: \iwats-admin\server\model\ip_config\request\iw_ip_config.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type IwIpConfigSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	IpGroupId      int        `json:"ipGroupId" form:"ipGroupId"`

	request.PageInfo
}
