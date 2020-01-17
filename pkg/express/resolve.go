package express

import (
	"sort"
	"strings"
	"unicode/utf8"
)

const (
	empty = ""
	slot  = " "
)

// 解析快递收货地址
func ResolveAddress(original string) Address {
	var (
		name       string
		phone      string
		addr       string
		postalCode string
	)
	addr = CleanAddress(original)
	p := Address{
		Original: original,
	}
	phone, addr = FilterPhone(addr)
	name, addr = FilterName(addr)
	postalCode, addr = FilterPostalCode(addr)

	addr = strings.TrimSpace(addr)
	p.Name = name
	p.Phone = phone
	p.Detail = addr
	p.PostalCode = postalCode
	return p
}

// 清理地址字符串中的多余字符
func CleanAddress(address string) (cleaned string) {
	cleaned = controlSymbolRegexp.ReplaceAllLiteralString(address, slot)
	cleaned = specialSymbolRegexp.ReplaceAllLiteralString(cleaned, slot)
	for _, text := range GetFilteredTextPattern() {
		cleaned = strings.ReplaceAll(cleaned, text, "")
	}
	return
}

// 获取要过滤掉的字符串
func GetFilteredTextPattern() []string {
	return []string{"详细地址", "收货地址", "收件地址", "地址", "所在地区", "地区", "姓名", "收货人", "收件人", "联系人", "收", "邮编", "联系电话", "电话", "联系人手机号码", "手机号码", "手机号"}
}

// 清理多余的空格
func removeDuplicatedSpaces(address string) string {
	return duplicatedSpaceRegexp.ReplaceAllLiteralString(address, slot)
}

// 筛选电话号码
func FilterPhone(address string) (phone string, left string) {
	left = address
	address = unifyPhonePattern(address)
	phone = phoneUnityRegexp.FindString(address)
	if len(phone) == 0 {
		return
	}
	left = removeDuplicatedSpaces(strings.ReplaceAll(left, phone, empty))
	return
}

// 格式化电话号码成连续的11位数字
func unifyPhonePattern(content string) string {
	for _, pattern := range phoneCompatibleRegexps {
		content = pattern.ReplaceAllLiteralString(content, "$1$2$3")
	}
	return content
}

// 从收货地址中筛选收货人名称
func FilterName(addr string) (name string, left string) {
	left = addr
	splits := strings.Split(addr, slot)
	if len(splits) == 0 {
		return "", addr
	}
	sort.SliceStable(splits, func(i, j int) bool {
		return len(splits[i]) < len(splits[j])
	})
	for i := range splits {
		chip := splits[i]
		charCount := utf8.RuneCountInString(chip)
		if charCount <= 0 || charCount > GetNameMaxLength() {
			continue
		}
		name = chip
	}
	if len(name) > 0 {
		left = strings.ReplaceAll(left, name, empty)
	}
	return
}

// 获取姓名的长度限制
func GetNameMaxLength() int {
	return 5
}

// 从收货地址中筛选邮政编码
func FilterPostalCode(addr string) (postalCode string, left string) {
	left = addr
	postalCode = postalCodeRegexp.FindString(addr)
	if len(postalCode) == 0 {
		return
	}
	left = strings.ReplaceAll(left, postalCode, empty)
	return
}
