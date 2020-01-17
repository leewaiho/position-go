package express

import "regexp"

var (
	phoneCompatibleRegexps = getPhoneCompatibleRegexps()
	phoneUnityRegexp       = regexp.MustCompile(getPhoneUnityPattern())
	controlSymbolRegexp    = regexp.MustCompile(getControlSymbolPattern())
	specialSymbolRegexp    = regexp.MustCompile(getSpecialSymbolPattern())
	duplicatedSpaceRegexp  = regexp.MustCompile(getDuplicatedSpacePattern())
	postalCodeRegexp       = regexp.MustCompile(getPostalCodePattern())
	nameRegexp             = regexp.MustCompile(getNamePattern())
)

// 获取模式: 手机兼容格式
func getPhoneCompatibleRegexps() []*regexp.Regexp {
	return []*regexp.Regexp{
		regexp.MustCompile("(\\d{3})-(\\d{4})-(\\d{4})"),
		regexp.MustCompile("(\\d{3}) (\\d{4}) (\\d{4})"),
	}
}

// 获取模式: 手机统一格式
func getPhoneUnityPattern() string {
	return "(\\d{7,12})|(\\d{3,4}-\\d{6,8})|(86-[1][0-9]{10})|(86[1][0-9]{10})|([1][0-9]{10})"
}

// 获取模式: 控制字符
func getControlSymbolPattern() string {
	return "[\r\n|\n|\t]"
}

// 获取模式: 特殊须过滤字符
func getSpecialSymbolPattern() string {
	return "[`~!@#$^&*()=|{}':;,\\[\\].<>/?！￥…（）—【】‘；：”“’。，、？]"
}

// 获取模式: 多个空格
func getDuplicatedSpacePattern() string {
	return " {2,}"
}

// 获取模式: 邮政编码
func getPostalCodePattern() string {
	return "\\d{6}"
}

// 获取要清理的字符串数组
func getCleanedTexts() []string {
	return []string{"详细地址", "收货地址", "收件地址", "地址", "所在地区", "地区", "姓名", "收货人", "收件人", "联系人", "收", "邮编", "联系电话", "电话", "联系人手机号码", "手机号码", "手机号"}
}

// 获取姓名的长度限制
func getNameMaxLength() int {
	return 5
}

// 获取模式: 中文名或者英文名
func getNamePattern() string {
	return "[\u4E00-\u9FA5|a-zA-Z]+"
}
