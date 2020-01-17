package express

import (
	"sort"
	"strings"
	"unicode/utf8"
)

// 解析快递收货地址
func ResolveAddress(original string) ResolvedAddress {
	var (
		name       string
		phone      string
		addr       string
		postalCode string
	)
	addr = cleanAddress(original)
	p := ResolvedAddress{
		Original: original,
	}
	phone, addr = filterPhone(addr)
	name, addr = filterName(addr)
	postalCode, addr = filterPostalCode(addr)

	addr = strings.TrimSpace(addr)
	p.Name = name
	p.Phone = phone
	p.Detail = addr
	p.PostalCode = postalCode
	return p
}

// 清理地址字符串中的多余字符
func cleanAddress(address string) (cleaned string) {
	cleaned = controlSymbolRegexp.ReplaceAllLiteralString(address, slot)
	cleaned = specialSymbolRegexp.ReplaceAllLiteralString(cleaned, slot)
	for _, text := range getCleanedTexts() {
		cleaned = strings.ReplaceAll(cleaned, text, "")
	}
	cleaned = cleanDuplicatedSpaces(cleaned)
	return
}

// 筛选电话号码
func filterPhone(address string) (phone string, left string) {
	left = address
	address = unifyPhone(address)
	phone = phoneUnityRegexp.FindString(address)
	if len(phone) == 0 {
		return
	}
	left = deleteSubString(left, phone)
	return
}

// 从收货地址中筛选收货人名称
func filterName(addr string) (name string, left string) {
	left = addr
	splits := strings.Split(addr, slot)
	if len(splits) == 0 {
		return "", addr
	}
	sort.SliceStable(splits, func(i, j int) bool {
		return len(splits[i]) < len(splits[j])
	})
	for i := range splits {
		s := splits[i]
		if !maybeName(s) {
			continue
		}
		name = s
	}
	if len(name) == 0 {
		return
	}
	left = deleteSubString(left, name)
	return
}

// 可能是收货人名称
func maybeName(chip string) bool {
	lenOfChip := utf8.RuneCountInString(chip)
	return lenOfChip > 0 && lenOfChip <= getNameMaxLength() && nameRegexp.MatchString(chip)
}

// 从收货地址中筛选邮政编码
func filterPostalCode(addr string) (postalCode string, left string) {
	left = addr
	postalCode = postalCodeRegexp.FindString(addr)
	if len(postalCode) == 0 {
		return
	}
	left = deleteSubString(left, postalCode)
	return
}
