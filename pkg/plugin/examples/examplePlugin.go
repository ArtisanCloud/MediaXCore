package main

import (
	"fmt"

	"github.com/ArtisanCloud/MediaXCore/pkg/plugin"
)

type ExamplePlugin struct {
	PluginName string
}

func (p *ExamplePlugin) Initialize(config map[string]interface{}) error {
	p.PluginName = "ExamplePlugin"
	return nil
}

func (p *ExamplePlugin) Name() string {
	return p.PluginName
}

func (p *ExamplePlugin) Publish(req plugin.PublishRequest, args ...interface{}) (plugin.PublishResult, error) {
	fmt.Printf("Publishing: Title=%s, Content=%s\n", req.Title, req.Content)
	return plugin.PublishResult{
		Status:  "success",
		Message: "Published Successfully",
	}, nil
}

// 导出插件实例
var Provider plugin.Provider = &ExamplePlugin{}
