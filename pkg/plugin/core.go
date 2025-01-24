package plugin

// PublishRequest 定义发布请求的结构体
type PublishRequest struct {
	Title   string
	Content string
}

// PublishResult 定义发布结果的结构体
type PublishResult struct {
	Status  string
	Message string
}

// Provider 定义插件需要实现的接口
type Provider interface {
	// Initialize 插件初始化方法
	Initialize(config map[string]interface{}) error
	// Name 返回插件名称
	Name() string
	// Publish 发布内容的方法
	Publish(req PublishRequest, args ...interface{}) (PublishResult, error)
}
