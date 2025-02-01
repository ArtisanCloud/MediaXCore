package main

import (
	"context"
	"github.com/ArtisanCloud/MediaXCore/pkg/logger"
	"github.com/ArtisanCloud/MediaXCore/pkg/logger/config"
	plugin2 "github.com/ArtisanCloud/MediaXCore/pkg/plugin"
	"github.com/ArtisanCloud/MediaXCore/pkg/plugin/core/contract"
	contract2 "github.com/ArtisanCloud/MediaXCore/pkg/plugin/examples/contract"
	"plugin"
)

func main() {
	ctx := context.Background()
	log := logger.GetLogger(nil)

	// 加载yaml配置文件
	configPlugin, err := plugin2.ReadPluginMetadata("./pkg/plugin/examples/plugin.yaml")
	if err != nil {
		panic(err)
	}
	pluginName := configPlugin.Name
	buildPath := configPlugin.BuildPath

	// 加载插件
	p, err := plugin.Open(buildPath)
	if err != nil {
		panic(err)
	}

	ptrProvider, err := plugin2.LookUpSymbol[contract.ProviderInterface](p, pluginName)
	if err != nil {
		panic(err)
	}

	examplePlugin := *ptrProvider
	// 初始化插件
	err = examplePlugin.Initialize(&ctx, &contract2.PluginConfig{
		LogConfig: config.LogConfig{
			Level:         "debug",
			Console:       true,
			UseJsonFormat: true,
			File: config.FileConfig{
				Enable: true,
				//InfoFilePath:  "./logs/info.log",
				//ErrorFilePath: "./logs/error.log",
			},
			//Loki: config.LokiConfig{},
			HttpDebug: true,
			Debug:     true,
		},
	})

	log.Info("plugin loaded name " + examplePlugin.Name(&ctx))

	res, err := examplePlugin.Publish(&ctx, &contract2.PublishRequest{})
	if err != nil {
		log.Error(err.Error())
	}
	pRes := res.(*contract2.PublishResponse)
	log.Info("published content: " + pRes.Msg)

}
