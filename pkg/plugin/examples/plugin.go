package main

import (
	"context"
	"fmt"
	"github.com/ArtisanCloud/MediaXCore/pkg/logger"
	"github.com/ArtisanCloud/MediaXCore/pkg/plugin/examples/contract"
	"reflect"
)

type ExampleXPlugin struct {
	PluginName string
	Logger     *logger.Logger
}

func NewExampleXPlugin() ExampleXPlugin {
	return ExampleXPlugin{
		PluginName: "ExamplePlugin",
	}
}

func (p *ExampleXPlugin) Initialize(ctx *context.Context, arg interface{}) error {

	// parse arg to contract config.PluginConfig firstly
	c, ok := arg.(*contract.PluginConfig)
	if !ok {
		argType := reflect.TypeOf(arg)
		return fmt.Errorf("initializing %s invalid argument type %s for arg: *config.PluginConfig", p.PluginName, argType.String())
	}

	p.PluginName = "ExamplePlugin"

	p.Logger = logger.NewLogger(&c.LogConfig)

	p.Logger.InfoF("Initializing %s plugin with config base uri: %+s\n", p.PluginName, c.BaseUri)

	return nil
}

func (p *ExampleXPlugin) Name(ctx *context.Context) string {
	return p.PluginName
}

func (p *ExampleXPlugin) Publish(ctx *context.Context, arg interface{}) (interface{}, error) {
	// parse arg to contract contract2.PublishRequest firstly
	req, ok := arg.(*contract.PublishRequest)
	if !ok {
		argType := reflect.TypeOf(arg)
		return nil, fmt.Errorf("invalid argument type %s for PluginMediaX arg: *contract.PublishRequest", argType.String())
	}
	p.Logger.InfoF("Publishing %s plugin with request: %+s\n", p.PluginName, req.Content)

	result := &contract.PublishResponse{}
	result.Code = 0
	result.Msg = "MediaX Example Plugin Published Successfully"

	return result, nil
}

var PluginExample ExampleXPlugin = NewExampleXPlugin()
