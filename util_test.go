package gexport

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCapitalize(t *testing.T) {
	as := assert.New(t)
	as.Equal(Capitalize("abc"), "Abc")
	as.Equal(Capitalize("_abc"), "_abc")
}

func TestNameCamelCase(t *testing.T) {
	as := assert.New(t)
	as.Equal(NameCamelCase("str_to_byte"), "StrToByte")
	as.Equal(NameCamelCase("_str_to_byte"), "StrToByte")
	as.Equal(NameCamelCase("iL"), "IL")
	as.Equal(NameCamelCase("reward_buyers"), "RewardBuyers")
}

//func TestNewJson(t *testing.T) {
//	fmt.Println(path.Dir("./"))
//}
