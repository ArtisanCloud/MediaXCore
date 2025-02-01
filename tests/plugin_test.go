package tests

import (
	"context"
	"fmt"
	"github.com/ArtisanCloud/MediaXCore/pkg/plugin/core/contract"
	"reflect"
	"testing"
)

type MockPlugin struct {
	PluginName string
}

func (m *MockPlugin) Initialize(ctx *context.Context, arg interface{}) error {
	m.PluginName = "MockPlugin"
	return nil
}

func (m *MockPlugin) Name(ctx *context.Context) string {
	return m.PluginName
}

func (m *MockPlugin) Publish(ctx *context.Context, arg interface{}) (interface{}, error) {
	// parse arg to contract contract2.PublishRequest firstly
	req, ok := arg.(*contract.PublishRequest)
	if !ok {
		argType := reflect.TypeOf(arg)
		return nil, fmt.Errorf("invalid argument type %s for PluginMediaX arg: *contract.PublishRequest", argType.String())
	}
	fmt.Printf("Publishing %s plugin with request: %+s\n", m.PluginName, req.Content)

	result := &contract.PublishResponse{}
	result.Code = 0
	result.Msg = "Mock Plugin Published Successfully"

	return result, nil

}

func TestProviderInterface(t *testing.T) {
	var provider contract.ProviderInterface // 定义接口类型的变量
	ctx := context.Background()

	mock := &MockPlugin{}
	mock.Initialize(&ctx, nil)
	provider = mock // 验证 MockPlugin 是否满足 contract.Provider 接口

	if provider.Name(&ctx) != "MockPlugin" {
		t.Errorf("Expected plugin name to be 'MockPlugin', got %s", provider.Name(&ctx))
	}
}
