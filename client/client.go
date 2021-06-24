package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"
)

func main() {

	fcg, err := NewYamlConfigProvider("your path")
	if err != nil {
		panic("初始化配置失败")
	}
	err = InitApplication(WithCfgProvider(fcg))
	if err != nil {
		panic("初始化应用失败")
	}
}

func SetFuncField(val Service) {

	v := reflect.ValueOf(val) // 这是指针的反射
	ele := v.Elem() // 拿到了指针指向的结构体
	t := ele.Type() // 拿到了指针指向的结构体的类型信息

	numField := t.NumField()
	for i := 0; i < numField; i++ {
		field := t.Field(i)
		fieldValue := ele.Field(i) // 用指针指向的结构体来访问
		if fieldValue.CanSet() {
			fn := func(args []reflect.Value) (results []reflect.Value) {
				in := args[0].Interface()
				out := reflect.New(field.Type.Out(0).Elem()).Interface()
				inData, err := json.Marshal(in)

				if err != nil {
					return []reflect.Value{reflect.ValueOf(out), reflect.ValueOf(err)}
				}

				client := http.Client{}

				name := val.ServiceName()

				cfg, err := App.CfgProvider.GetServiceConfig(name)

				if err != nil {
					return []reflect.Value{reflect.ValueOf(out), reflect.ValueOf(err)}
				}
				req, err := http.NewRequest("POST", cfg.Endpoint, bytes.NewReader(inData))

				if err != nil {
					return []reflect.Value{reflect.ValueOf(out), reflect.ValueOf(err)}
				}

				req.Header.Set("Content-Type", "application/json")
				req.Header.Set("sparrow-service", name)
				req.Header.Set("sparrow-service-method", field.Name)

				resp, err := client.Do(req)

				if err != nil {
					return []reflect.Value{reflect.ValueOf(out), reflect.ValueOf(err)}
				}
				data, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					return []reflect.Value{reflect.ValueOf(out), reflect.ValueOf(err)}
				}
				err = json.Unmarshal(data, out)
				if err != nil {
					return []reflect.Value{reflect.ValueOf(out), reflect.ValueOf(err)}
				}
				return []reflect.Value{reflect.ValueOf(out), reflect.Zero(reflect.TypeOf(new(error)).Elem())}
			}
			fieldValue.Set(reflect.MakeFunc(field.Type, fn))
		}
	}
}

type Service interface {
	ServiceName() string
}
