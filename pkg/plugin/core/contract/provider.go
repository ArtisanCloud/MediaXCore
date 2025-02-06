package contract

import "context"

// ---------------------------------------------
// Provider 定义插件需要实现的接口
// ---------------------------------------------
type ProviderInterface interface {
	// Initialize 插件初始化方法
	Initialize(ctx *context.Context, config interface{}) error
	// Name 返回插件名称
	Name(ctx *context.Context) string
	// Publish 发布内容的方法
	Publish(ctx *context.Context, req interface{}) (interface{}, error)
}
