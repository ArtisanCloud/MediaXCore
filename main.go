package main

import (
	"context"
	"github.com/ArtisanCloud/MediaXCore/pkg/logger"
	"github.com/ArtisanCloud/MediaXCore/pkg/logger/config"
	plugin2 "github.com/ArtisanCloud/MediaXCore/pkg/plugin"
	"github.com/ArtisanCloud/MediaXCore/pkg/plugin/core/contract"
	"plugin"
)

func main() {
	ctx := context.Background()
	configPlugin := &contract.PluginConfig{
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
	}
	log := logger.GetLogger(&configPlugin.LogConfig)
	//fmt.Printf("main logger address %p \n", log)
	// 加载yaml配置文件
	configPluginMetadata, err := plugin2.ReadPluginMetadata("./pkg/plugin/examples/plugin/plugin.yaml")
	if err != nil {
		panic(err)
	}
	pluginName := configPluginMetadata.Name
	buildPath := configPluginMetadata.BuildPath

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
	err = examplePlugin.Initialize(&ctx, configPlugin)

	log.Info("plugin loaded name " + examplePlugin.Name(&ctx))

	res, err := examplePlugin.Publish(&ctx, &contract.PublishRequest{})
	if err != nil {
		log.Error(err.Error())
	}
	pRes := res.(*contract.PublishResponse)
	log.Info("published content: " + pRes.Msg)

}
