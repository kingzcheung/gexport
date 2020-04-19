package gexport

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestJson_parseJson(t *testing.T) {
	input := []byte(`{
    "rewardable": true,
    "zfq": "zfq",
    "setting": {
        "description": "this is desc",
        "default_amount": 200
    },
    "total_rewards_count": 0,
    "reward_buyers": [
      {"foo":"bar"}
    ]
}`)
	nj := NewJson()
	as := assert.New(t)
	p, err := nj.parseJson(input)
	if err != nil {
		as.Error(err)
	}
	as.Equal(p["zfq"], "zfq")
}

func TestJson_genInline(t *testing.T) {
	input := []byte(`{
    "rewardable": true,
    "zfq": "zfq",
    "setting": {
        "description": "this is desc",
        "default_amount": 200
    },
    "total_rewards_count": 0,
    "reward_buyers": [
      {"foo":"bar"}
    ]
}`)
	nj := NewJson()
	as := assert.New(t)
	p, err := nj.parseJson(input)
	if err != nil {
		as.Error(err)
	}
	inlineStr := nj.generated(p)
	fmt.Println(inlineStr)
}

func TestJson_Parse(t *testing.T) {
	input := `{
    "rewardable": true,
    "zfq": "zfq",
    "setting": {
        "description": "this is desc",
        "default_amount": 200
    },
    "total_rewards_count": 0,
    "reward_buyers": [
      {"foo":"bar"}
    ]
}`
	nj := NewJson(false)
	as := assert.New(t)
	res, err := nj.Parse(input)
	fmt.Println(res)
	if err != nil {
		as.Error(err)
	}

	as.Equal(len(res), 3)

}

func TestJson_dimensionalReduction(t *testing.T) {
	input, err := ioutil.ReadFile("testdata/data.json")
	input = []byte(`{
    "rewardable": true,
    "zfq": "zfq",
    "setting": {
        "description": "this is desc",
        "default_amount": 200
    },
    "total_rewards_count": 0,
    "reward_buyers": [
      {"foo":"bar"}
    ]
}`)
	as := assert.New(t)
	if err != nil {
		as.Error(err)
	}
	nj := NewJson()
	res, _ := nj.parseJson(input)
	nj.dimensionalReduction(res, DefaultStructName)
	for key, result := range nj.results {
		fmt.Println(key, result)
	}

}
