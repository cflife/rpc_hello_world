package main

import "sync"

type application struct {
	CfgProvider ConfigProvider
}

// App 必须要显示使用 InitApplication 来创建
var App *application
var appOnce sync.Once

type AppOption func(app *application) error

func InitApplication(opts ... AppOption) error {
	var err error
	appOnce.Do(func() {
		App = &application{
			// 默认实现
			CfgProvider: NewInMemoryConfigProvider(),
		}
		for _, opt := range opts {
			err = opt(App)
			if err != nil {
				return
			}
		}
	})
	return err
}

func WithCfgProvider(cfp ConfigProvider) AppOption {
	return func(app *application) error {
		app.CfgProvider = cfp
		return nil
	}
}