# 单独构建插件文件，传入模块名
.PHONY: plugin.build
plugin.build:
	@echo "正在构建Example插件..."
	go build -o pkg/plugin/examples/plugin/plugin.so -buildmode=plugin pkg/plugin/examples/src/plugin.go
	@echo "Example插件构建完成"