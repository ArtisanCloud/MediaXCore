package plugin

import (
	"errors"
	"fmt"
	"github.com/ArtisanCloud/MediaXCore/pkg/plugin/core/contract"
	"plugin"
)

// LoadPlugin 动态加载插件
func LoadPlugin(path string, pluginName string) (contract.ProviderInterface, error) {
	// 打开插件文件
	p, err := plugin.Open(path)
	if err != nil {
		return nil, err
	}

	// 查找 "Provider" 符号
	mediaXPlugin, err := LookUpSymbol[contract.ProviderInterface](p, pluginName)
	if err != nil {
		return nil, err
	}

	return *mediaXPlugin, nil
}

func LookUpSymbol[M any](plugin *plugin.Plugin, symbolName string) (*M, error) {
	symbol, err := plugin.Lookup(symbolName)
	if err != nil {
		return nil, err
	}
	switch symbol.(type) {
	case *M:
		return symbol.(*M), nil
	case M:
		result := symbol.(M)
		return &result, nil
	default:
		return nil, errors.New(fmt.Sprintf("unexpected type from module symbol: %T", symbol))
	}
}
