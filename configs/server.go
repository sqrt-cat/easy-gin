package configs

// 服务配置 防止变量污染故用函数组织
func GetServerConfig() (serverConfig map[string]string) {
	serverConfig = make(map[string]string)

	serverConfig["HOST"] = "0.0.0.0"            //监听地址
	serverConfig["PORT"] = "8080"               //监听端口
	serverConfig["VIEWS_PATTERN"] = "views/*/*" //模板路径pattern
	serverConfig["ENV"] = "debug"               // 环境模式 release/debug/test
	return
}
