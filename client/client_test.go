package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 跑一个测试
func TestSetFuncField(t *testing.T) {

	path := `..\test\client.yaml`
	ycp, _ := NewYamlConfigProvider(path)

	_ = InitApplication(WithCfgProvider(ycp))

	helloService := &hello{
	}

	SetFuncField(helloService)

	res, err := helloService.SayHello(&Input{
		Name: "golang",
	})

	assert.Nil(t, err)
	assert.Equal(t, "Hello, golang", res.Msg)
}

type hello struct {
	SayHello func(in *Input) (*Output, error)
}

func (h hello) ServiceName() string {
	return "hello"
}

type Input struct {
	Name string
}

type Output struct {
	Msg string
}
