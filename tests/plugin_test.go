package tests

import (
	"github.com/ArtisanCloud/MediaXCore/pkg/plugin/core/contract"
	"testing"
)

type MockPlugin struct {
	PluginName string
}

func (m *MockPlugin) Initialize(config map[string]interface{}) error {
	m.PluginName = "MockPlugin"
	return nil
}

func (m *MockPlugin) Name() string {
	return m.PluginName
}

func (m *MockPlugin) Publish(req *contract.PublishRequest, args ...interface{}) (*contract.PublishResult, error) {
	return &contract.PublishResult{
		Status:  "success",
		Message: "Mock Publish Successful",
	}, nil
}

func TestProviderInterface(t *testing.T) {
	var provider contract.ProviderInterface // 定义接口类型的变量

	mock := &MockPlugin{}
	mock.Initialize(nil)
	provider = mock // 验证 MockPlugin 是否满足 contract.Provider 接口

	if provider.Name() != "MockPlugin" {
		t.Errorf("Expected plugin name to be 'MockPlugin', got %s", provider.Name())
	}
}
