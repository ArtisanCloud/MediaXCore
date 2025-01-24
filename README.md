# MediaXCore

MediaXCore 是一个核心模块，提供了插件化的基础设施和标准接口，允许开发者定义和扩展功能，适用于各种场景中的内容发布和管理。

---

## 功能简介

- 提供插件 `Provider` 接口，支持动态扩展。
- 标准化的内容发布功能（`Publish`）。
- 提供基础结构体 `PublishRequest` 和 `PublishResult`。
- 可用于实现独立插件系统和与其他服务集成。

---

## 目录结构

```
MediaXCore
├── LICENSE
├── README.md           # 项目说明文件
├── go.mod              # Go 模块管理文件
├── go.sum              # 依赖项锁定文件
├── pkg                 # 核心功能包
│   ├── plugin          # 插件接口与定义
│   │   ├── core.go     # Provider 接口定义及相关结构体
├── tests               # 单元测试
│   ├── plugin_test.go  # 针对插件接口的测试
```

---

## 快速开始

### 1. 克隆项目

```bash
git clone https://github.com/ArtisanCloud/MediaXCore.git
cd MediaXCore
```

### 2. 安装依赖

```bash
go mod tidy
```

### 3. 运行单元测试

确保一切功能正常。

```bash
cd tests
go test -v
```

测试示例输出：

```
=== RUN   TestProviderInterface
--- PASS: TestProviderInterface (0.00s)
=== RUN   TestProviderPublish
--- PASS: TestProviderPublish (0.00s)
PASS
ok      github.com/ArtisanCloud/MediaXCore/tests        0.619s
```

---

## 如何实现插件

开发者可以按照 `Provider` 接口的定义，实现自定义插件。

### 示例

**定义插件：**

```go
package main

import (
	"fmt"
	"github.com/ArtisanCloud/MediaXCore/pkg/plugin"
)

type MyPlugin struct {
	PluginName string
}

func (p *MyPlugin) Initialize(config map[string]interface{}) error {
	p.PluginName = "MyPlugin"
	return nil
}

func (p *MyPlugin) Name() string {
	return p.PluginName
}

func (p *MyPlugin) Publish(req plugin.PublishRequest, args ...interface{}) (plugin.PublishResult, error) {
	fmt.Printf("Publishing: %s - %s\n", req.Title, req.Content)
	return plugin.PublishResult{
		Status:  "success",
		Message: "Content Published Successfully",
	}, nil
}

func main() {
	plugin := &MyPlugin{}
	plugin.Initialize(nil)

	result, _ := plugin.Publish(plugin.PublishRequest{
		Title:   "Example Title",
		Content: "Example Content",
	})
	fmt.Println(result)
}
```

---

## 贡献代码

欢迎对 MediaXCore 进行改进或扩展！

### 提交代码步骤

1. Fork 本仓库。
2. 创建新分支：`git checkout -b feature-xxx`。
3. 提交更改：`git commit -m 'Add xxx feature'`。
4. 推送分支：`git push origin feature-xxx`。
5. 创建 Pull Request。

---

## 联系我们

如有问题，请通过以下方式联系我们：

- 作者邮箱：[matrix-x@artisan-cloud.com](matrix-x@artisan-cloud.com)
- 项目地址：[https://github.com/ArtisanCloud/MediaXCore](https://github.com/ArtisanCloud/MediaXCore)

---

## License

本项目遵循 [MIT License](./LICENSE) 开源协议。
