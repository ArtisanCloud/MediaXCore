package main

import (
	"fmt"
	plugin2 "github.com/ArtisanCloud/MediaXCore/pkg/plugin"
	"github.com/ArtisanCloud/MediaXCore/pkg/plugin/core/contract"
	"plugin"
)

func main() {

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
	fmt.Printf("plugin loaded name :%s \n", examplePlugin.Name())
	fmt.Println(examplePlugin.Publish(&contract.PublishRequest{}))

}
