package plugin

import "github.com/opensec-cn/kunpeng/util"

// GoPlugin 插件接口
type GoPlugin interface {
	Init() Plugin
	Check(netloc string, meta TaskMeta) bool
	GetResult() []Plugin
}

// Regist 注册插件
func Regist(target string, plugin GoPlugin) {
	GoPlugins[target] = append(GoPlugins[target], plugin)
	var pluginInfo = plugin.Init()
	util.Logger.Println("init plugin:", pluginInfo.Name)
}
