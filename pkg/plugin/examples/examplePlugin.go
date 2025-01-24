package examples

import (
	"fmt"
	"github.com/ArtisanCloud/MediaXCore/pkg/plugin/core/contract"
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

func (p *ExamplePlugin) Publish(req contract.PublishRequest, args ...interface{}) (contract.PublishResult, error) {
	fmt.Printf("Publishing: Title=%s, Content=%s\n", req.Title, req.Content)
	return contract.PublishResult{
		Status:  "success",
		Message: "Published Successfully",
	}, nil
}
