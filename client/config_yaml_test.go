package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/fs"
	"os"
	"testing"
)

func TestNewYamlConfigProvider(t *testing.T) {
	path := `D:\workspace\go\src\geektime\sparrow\test\client.yaml`
	ycp, err := NewYamlConfigProvider(path)
	assert.Nil(t, err)

	cfg , err := ycp.GetServiceConfig("hello")
	assert.Nil(t, err)
	assert.Equal(t, "http://localhost:8080", cfg.Endpoint)
}

func TestFile(t *testing.T) {
	path := `D:\workspace\go\src\geektime\sparrow\test\client.yaml`
	// 打开一个文件，只能读
	file, _ := os.Open(path)
	fmt.Printf("%s\n", file.Name())

	// 创建了一个文件。如果文件已经存在，会被清空
	file, _ = os.Create(`D:\workspace\go\src\geektime\sparrow\test\test.txt`)

	// 写入数据，cnt 是写入数据字节数
	cnt, _ := file.WriteString("你好")
	fmt.Printf("写入字节数：%d", cnt)

	// 打开一个append only 模式的文件
	file, _ = os.OpenFile(`D:\workspace\go\src\geektime\sparrow\test\test.txt`, os.O_APPEND, fs.ModeAppend)
	_, _ = file.WriteString("hello, world")

	os.Getwd()
}

func TestMap(t *testing.T) {
	m := map[string]string {
		"A": "a",
	}

	for key, val := range m {
		fmt.Printf("%s => %s", key, val)
	}
}
