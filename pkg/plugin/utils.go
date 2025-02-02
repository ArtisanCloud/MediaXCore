package plugin

import (
	"errors"
	"fmt"
	"github.com/ArtisanCloud/MediaXCore/pkg/plugin/core"
	"github.com/ArtisanCloud/MediaXCore/pkg/plugin/core/contract"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"plugin"
)

// LoadConfig 加载yaml配置文件

// 读取插件描述文件
func ReadPluginBundleMetadata(pluginsFilePath string) (*core.PluginsMetadata, error) {
	// 检查路径是否存在
	if _, err := os.Stat(pluginsFilePath); os.IsNotExist(err) {
		return nil, fmt.Errorf("file does not exist: %s", pluginsFilePath)
	}

	// 读取文件
	data, err := os.ReadFile(pluginsFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %v", err)
	}

	metadatas := &core.PluginsMetadata{}
	if err = yaml.Unmarshal(data, metadatas); err != nil {
		return nil, fmt.Errorf("failed to decode plugin description file %s: %v", pluginsFilePath, err)
	}

	// 增加验证逻辑
	if metadatas.Name == "" || metadatas.Type == "" {
		return nil, fmt.Errorf("plugin metadata is incomplete or invalid in %s", pluginsFilePath)
	}

	return metadatas, nil
}

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

// 检查是否为 .so 文件
func IsSOFile(path string) bool {
	ext := filepath.Ext(path)
	return ext == ".so" || ext == ".dylib" || ext == ".dll" // 扩展支持
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
