package core

type PluginType string

const (
	Open     PluginType = "open"
	Closed   PluginType = "closed"
	External PluginType = "external"
)

// 插件描述文件结构
type PluginMetadata struct {
	Name    string                 `yaml:"name"`
	Version string                 `yaml:"version"`
	Type    PluginType             `yaml:"type"`
	Path    string                 `yaml:"path"`
	Config  map[string]interface{} `yaml:"config"`
}
