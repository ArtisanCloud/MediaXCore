
# MediaXCore

MediaXCore 是一个支持插件化的核心模块，为内容发布和管理场景提供标准接口和基础设施。它允许开发者通过插件动态扩展功能，同时支持独立插件系统和与其他服务的集成。

---

## 项目结构

```bash
.
├── README.md                # 项目说明文档
├── go.mod                   # Go 模块管理文件
├── main.go                  # 示例主程序
├── pkg
│   └── plugin
│       ├── core
│       │   └── contract
│       │       └── pluginInterface.go   # 插件接口定义
│       ├── examples
│       │   └── examplePlugin.go         # 插件示例
│       └── utils.go                     # 插件工具函数
└── tests
└── plugin_test.go        # 单元测试
```

---

## 功能介绍

- **插件接口**：提供 `ProviderInterface`，支持实现自定义插件。
- **动态加载**：支持通过 Go 动态加载 `.so` 插件文件。
- **标准化接口**：包含 `PublishRequest` 和 `PublishResult` 数据结构，规范化内容发布流程。
- **测试用例**：内置单元测试，验证插件接口实现。

---

## 使用方法

### 1. 编译插件

运行以下命令编译插件：

```bash
go build -o pkg/plugin/examples/plugin.so -buildmode=plugin pkg/plugin/examples/plugin.go
```

### 2. 运行主程序

通过 `main.go` 加载并使用插件：

```bash
go run main.go
```

### 3. 单元测试

执行单元测试以验证接口实现：

```bash
go test ./tests
```

---

## 示例输出

运行主程序后，输出类似如下：

```
plugin loaded name: ExamplePlugin
&{success Mock Publish Successful} <nil>
```

---

## 开发者指南

- **扩展插件**：继承 `ProviderInterface` 接口并实现 `Initialize`、`Name` 和 `Publish` 方法。
- **新增功能**：在 `pkg/plugin` 模块中添加新功能或插件工具函数。

---

## 许可证

此项目基于 [MIT License](https://opensource.org/licenses/MIT)。
```