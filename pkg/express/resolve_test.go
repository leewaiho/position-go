package express

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestResolveAddress_SepByComma(t *testing.T) {
	var (
		name       = "李四"
		phone      = "13800000000"
		detail     = "广东省深圳市宝安区海雅缤纷城L3-303"
		postalCode = "518000"
	)
	unresolved := fmt.Sprintf("%s,%s,%s,%s", name, phone, detail, postalCode)
	resolved := ResolveAddress(unresolved)
	s := assert.New(t)
	s.Equal(name, resolved.Name, "收货人名称解析失败")
	s.Equal(phone, resolved.Phone, "收货人手机号码解析失败")
	s.Equal(detail, resolved.Detail, "收货地址详情解析失败")
	s.Equal(postalCode, resolved.PostalCode, "收货地址邮政编码解析失败")
}
