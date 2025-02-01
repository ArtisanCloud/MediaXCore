plugin.all.build: plugin.example1.build plugin.example2.build

.PHONY: plugin.example1.build
plugin.example1.build:
	@echo "正在构建Example插件..."
	go build -o pkg/plugin/examples/plugins/pluginExample1.so -buildmode=plugin pkg/plugin/examples/src/pluginExample1.go
	@echo "Example插件构建完成"

.PHONY: plugin.example2.build
plugin.example2.build:
	@echo "正在构建Example插件..."
	go build -o pkg/plugin/examples/plugins/pluginExample2.so -buildmode=plugin pkg/plugin/examples/src/pluginExample2.go
	@echo "Example插件构建完成"