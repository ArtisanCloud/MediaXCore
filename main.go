package main

import (
	"fmt"
	"github.com/ArtisanCloud/MediaXCore/pkg/plugin/core/contract"
	"github.com/ArtisanCloud/MediaXCore/pkg/plugin/examples"
)

func main() {

	// 导出插件实例
	var examplePlugin contract.Provider = &examples.ExamplePlugin{}

	_ = examplePlugin.Initialize(nil)
	result, err := examplePlugin.Publish(contract.PublishRequest{})
	if err != nil {
		panic(err)
	}

	fmt.Println(result.Status)
	fmt.Println(result.Message)

}
