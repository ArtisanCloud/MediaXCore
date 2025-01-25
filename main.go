package main

import (
	"fmt"
	plugin2 "github.com/ArtisanCloud/MediaXCore/pkg/plugin"
	"github.com/ArtisanCloud/MediaXCore/pkg/plugin/core/contract"
	"plugin"
)

func main() {

	// 导出插件实例
	// 加载插件
	p, err := plugin.Open("./pkg/plugin/examples/plugin.so")
	if err != nil {
		panic(err)
	}

	ptrProvider, err := plugin2.LookUpSymbol[contract.ProviderInterface](p, "ExamplePlugin")
	if err != nil {
		panic(err)
	}
	examplePlugin := *ptrProvider
	fmt.Printf("plugin loaded name :%s \n", examplePlugin.Name())
	fmt.Println(examplePlugin.Publish(&contract.PublishRequest{}))

}
