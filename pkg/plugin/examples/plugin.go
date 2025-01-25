package main

import (
	"fmt"
	"github.com/ArtisanCloud/MediaXCore/pkg/plugin/core/contract"
)

type ExampleXPlugin struct {
	PluginName string
}

func NewExampleXPlugin() ExampleXPlugin {
	return ExampleXPlugin{
		PluginName: "ExamplePlugin",
	}
}

func (p *ExampleXPlugin) Initialize(config map[string]interface{}) error {
	p.PluginName = "ExamplePlugin"
	return nil
}

func (p *ExampleXPlugin) Name() string {
	return p.PluginName
}

func (p *ExampleXPlugin) Publish(req *contract.PublishRequest, args ...interface{}) (*contract.PublishResult, error) {
	fmt.Printf("Publishing: Title=%s, Content=%s\n", req.Title, req.Content)
	return &contract.PublishResult{
		Status:  "success",
		Message: "Published Successfully",
	}, nil
}

var ExamplePlugin ExampleXPlugin = NewExampleXPlugin()
