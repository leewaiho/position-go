package express

import "regexp"

var (
	phoneCompatibleRegexps = GetPhoneCompatibleRegexps()
	phoneUnityRegexp       = regexp.MustCompile(GetPhoneUnityPattern())
	controlSymbolRegexp    = regexp.MustCompile(GetControlSymbolPattern())
	specialSymbolRegexp    = regexp.MustCompile(GetSpecialSymbolPattern())
	duplicatedSpaceRegexp  = regexp.MustCompile(GetDuplicatedSpacePattern())
	postalCodeRegexp       = regexp.MustCompile(GetPostalCodePattern())
)

// 获取模式: 手机兼容格式
func GetPhoneCompatibleRegexps() []*regexp.Regexp {
	return []*regexp.Regexp{
		regexp.MustCompile("(\\d{3})-(\\d{4})-(\\d{4})"),
		regexp.MustCompile("(\\d{3}) (\\d{4}) (\\d{4})"),
	}
}

// 获取模式: 手机统一格式
func GetPhoneUnityPattern() string {
	return "(\\d{7,12})|(\\d{3,4}-\\d{6,8})|(86-[1][0-9]{10})|(86[1][0-9]{10})|([1][0-9]{10})"
}

// 获取模式: 控制字符
func GetControlSymbolPattern() string {
	return "[\r\n|\n|\t]"
}

// 获取模式: 特殊须过滤字符
func GetSpecialSymbolPattern() string {
	return "[`~!@#$^&*()=|{}':;,\\[\\].<>/?！￥…（）—【】‘；：”“’。，、？]"
}

// 获取模式: 多个空格
func GetDuplicatedSpacePattern() string {
	return " {2,}"
}

// 获取模式: 邮政编码
func GetPostalCodePattern() string {
	return "\\d{6}"
}
