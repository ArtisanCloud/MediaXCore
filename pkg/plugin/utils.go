package plugin

import (
	"errors"
	"github.com/ArtisanCloud/MediaXCore/pkg/plugin/core/contract"
	"plugin"
)

// LoadPlugin 动态加载插件
func LoadPlugin(path string) (contract.Provider, error) {
	// 打开插件文件
	plug, err := plugin.Open(path)
	if err != nil {
		return nil, err
	}

	// 查找 "Provider" 符号
	symProvider, err := plug.Lookup("Provider")
	if err != nil {
		return nil, err
	}

	// 类型断言为 Provider 接口
	provider, ok := symProvider.(contract.Provider)
	if !ok {
		return nil, errors.New("invalid plugin type")
	}

	return provider, nil
}
