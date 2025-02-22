package upload_account

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Name struct {
	First string `json:"first"`
	Last  string `json:"last"`
}

type User struct {
	Name Name `json:"name"`
}

type Response struct {
	Results []User `json:"results"`
}

// GenerateRandomName 获取deviceName
func GenerateRandomName(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	var response Response
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return ""
	}

	if len(response.Results) < 0 {
		return ""
	}

	name := response.Results[0].Name
	return fmt.Sprintf("%s %s", name.First, name.Last)
}

var countryCodeMap = map[string]string{
	"1":   "美国",
	"7":   "俄罗斯",
	"20":  "埃及",
	"27":  "南非",
	"30":  "希腊",
	"31":  "荷兰",
	"32":  "比利时",
	"33":  "法国",
	"34":  "西班牙",
	"36":  "匈牙利",
	"39":  "意大利",
	"40":  "罗马尼亚",
	"41":  "瑞士",
	"43":  "奥地利",
	"44":  "英国",
	"45":  "丹麦",
	"46":  "瑞典",
	"47":  "挪威",
	"48":  "波兰",
	"49":  "德国",
	"51":  "秘鲁",
	"52":  "墨西哥",
	"53":  "古巴",
	"54":  "阿根廷",
	"55":  "巴西",
	"56":  "智利",
	"57":  "哥伦比亚",
	"58":  "委内瑞拉",
	"60":  "马来西亚",
	"61":  "澳大利亚",
	"62":  "印度尼西亚",
	"63":  "菲律宾",
	"64":  "新西兰",
	"65":  "新加坡",
	"66":  "泰国",
	"81":  "日本",
	"82":  "韩国",
	"84":  "越南",
	"86":  "中国",
	"90":  "土耳其",
	"91":  "印度",
	"92":  "巴基斯坦",
	"93":  "阿富汗",
	"94":  "斯里兰卡",
	"95":  "缅甸",
	"98":  "伊朗",
	"211": "南苏丹",
	"212": "摩洛哥",
	"213": "阿尔及利亚",
	"216": "突尼斯",
	"218": "利比亚",
	"220": "冈比亚",
	"221": "塞内加尔",
	"222": "毛里塔尼亚",
	"223": "马里",
	"224": "几内亚",
	"225": "科特迪瓦",
	"226": "布基纳法索",
	"227": "尼日尔",
	"228": "多哥",
	"229": "贝宁",
	"230": "毛里求斯",
	"231": "利比里亚",
	"232": "塞拉利昂",
	"233": "加纳",
	"234": "尼日利亚",
	"235": "乍得",
	"236": "中非共和国",
	"237": "喀麦隆",
	"238": "佛得角",
	"239": "圣多美和普林西比",
	"240": "赤道几内亚",
	"241": "加蓬",
	"242": "刚果",
	"243": "刚果（金）",
	"244": "安哥拉",
	"245": "几内亚比绍",
	"246": "英属印度洋领地",
	"247": "阿森松岛",
	"248": "塞舌尔",
	"249": "苏丹",
	"250": "卢旺达",
	"251": "埃塞俄比亚",
	"252": "索马里",
	"253": "吉布提",
	"254": "肯尼亚",
	"255": "坦桑尼亚",
	"256": "乌干达",
	"257": "布隆迪",
	"258": "莫桑比克",
	"260": "赞比亚",
	"261": "马达加斯加",
	"262": "留尼汪",
	"263": "津巴布韦",
	"264": "纳米比亚",
	"265": "马拉维",
	"266": "莱索托",
	"267": "博茨瓦纳",
	"268": "斯威士兰",
	"269": "科摩罗",
	"290": "圣赫勒拿",
	"291": "厄立特里亚",
	"297": "阿鲁巴",
	"298": "法罗群岛",
	"299": "格陵兰",
	"350": "直布罗陀",
	"351": "葡萄牙",
	"352": "卢森堡",
	"353": "爱尔兰",
	"354": "冰岛",
	"355": "阿尔巴尼亚",
	"356": "马耳他",
	"357": "塞浦路斯",
	"358": "芬兰",
	"359": "保加利亚",
	"370": "立陶宛",
	"371": "拉脱维亚",
	"372": "爱沙尼亚",
	"373": "摩尔多瓦",
	"374": "亚美尼亚",
	"375": "白俄罗斯",
	"376": "安道尔",
	"377": "摩纳哥",
	"378": "圣马力诺",
	"379": "梵蒂冈",
	"380": "乌克兰",
	"381": "塞尔维亚",
	"382": "黑山",
	"383": "科索沃",
	"385": "克罗地亚",
	"386": "斯洛文尼亚",
	"387": "波斯尼亚和黑塞哥维那",
	"389": "北马其顿",
	"420": "捷克",
	"421": "斯洛伐克",
	"423": "列支敦士登",
	"500": "福克兰群岛",
	"501": "伯利兹",
	"502": "危地马拉",
	"503": "萨尔瓦多",
	"504": "洪都拉斯",
	"505": "尼加拉瓜",
	"506": "哥斯达黎加",
	"507": "巴拿马",
	"508": "圣皮埃尔和密克隆",
	"509": "海地",
	"590": "瓜德罗普",
	"591": "玻利维亚",
	"592": "圭亚那",
	"593": "厄瓜多尔",
	"594": "法属圭亚那",
	"595": "巴拉圭",
	"596": "马提尼克",
	"597": "苏里南",
	"598": "乌拉圭",
	"599": "荷属安的列斯",
	"670": "东帝汶",
	"672": "诺福克岛",
	"673": "文莱",
	"674": "瑙鲁",
	"675": "巴布亚新几内亚",
	"676": "汤加",
	"677": "所罗门群岛",
	"678": "瓦努阿图",
	"679": "斐济",
	"680": "帕劳",
	"681": "瓦利斯和富图纳",
	"682": "库克群岛",
	"683": "纽埃",
	"685": "萨摩亚",
	"686": "基里巴斯",
	"687": "新喀里多尼亚",
	"688": "图瓦卢",
	"689": "法属波利尼西亚",
	"690": "托克劳",
	"691": "密克罗尼西亚",
	"692": "马绍尔群岛",
	"850": "朝鲜",
	"852": "香港",
	"853": "澳门",
	"855": "柬埔寨",
	"856": "老挝",
	"880": "孟加拉国",
	"886": "台湾",
	"960": "马尔代夫",
	"961": "黎巴嫩",
	"962": "约旦",
	"963": "叙利亚",
	"964": "伊拉克",
	"965": "科威特",
	"966": "沙特阿拉伯",
	"967": "也门",
	"968": "阿曼",
	"970": "巴勒斯坦",
	"971": "阿联酋",
	"972": "以色列",
	"973": "巴林",
	"974": "卡塔尔",
	"975": "不丹",
	"976": "蒙古",
	"977": "尼泊尔",
	"992": "塔吉克斯坦",
	"993": "土库曼斯坦",
	"994": "阿塞拜疆",
	"995": "格鲁吉亚",
	"996": "吉尔吉斯斯坦",
	"998": "乌兹别克斯坦",
}

// GetCountryByPhone 通过手机号码获取国家
func GetCountryByPhone(phone string) (string, string, error) {
	// 去掉手机号码中的空格和其他非数字字符
	phone = strings.TrimSpace(phone)
	phone = strings.ReplaceAll(phone, "+", "")
	phone = strings.ReplaceAll(phone, " ", "")
	phone = strings.ReplaceAll(phone, "-", "")
	phone = strings.ReplaceAll(phone, "(", "")
	phone = strings.ReplaceAll(phone, ")", "")

	// 遍历国家代码映射，查找匹配的前缀
	for code, country := range countryCodeMap {
		if strings.HasPrefix(phone, code) {

			return code, country, nil
		}
	}

	return "", "", fmt.Errorf("找不到对应的国家")
}
