package main

import (
	"context"
	"fmt"
	"github.com/ArtisanCloud/MediaXCore/pkg/logger"
	"github.com/ArtisanCloud/MediaXCore/pkg/plugin/core/contract"
	"reflect"
)

type ExampleXPlugin1 struct {
	PluginName string
	Logger     *logger.Logger
}

func NewExampleXPlugin1() ExampleXPlugin1 {
	return ExampleXPlugin1{
		PluginName: "ExamplePlugin1",
	}
}

func (p *ExampleXPlugin1) Initialize(ctx *context.Context, arg interface{}) error {

	// parse arg to contract config.PluginConfig firstly
	c, ok := arg.(*contract.PluginConfig)
	if !ok {
		argType := reflect.TypeOf(arg)
		return fmt.Errorf("initializing %s invalid argument type %s for arg: *config.PluginConfig", p.PluginName, argType.String())
	}

	p.Logger = logger.NewLogger(&c.LogConfig)
	//fmt.Printf("plugin logger address %p \n", p.Logger)

	p.Logger.InfoF("Initializing %s plugin with config base uri: %+s\n", p.PluginName, c.BaseUri)

	return nil
}

func (p *ExampleXPlugin1) Name(ctx *context.Context) string {
	return p.PluginName
}

func (p *ExampleXPlugin1) Publish(ctx *context.Context, arg interface{}) (interface{}, error) {
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

var PluginExample1 ExampleXPlugin1 = NewExampleXPlugin1()
