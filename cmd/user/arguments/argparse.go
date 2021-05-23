package arguments

import "flag"

type ArgConfig struct {
	Port int
}

func ArgParse() ArgConfig {
	var argConfig ArgConfig

	flag.IntVar(&argConfig.Port, "port", 8080, "server listen port")
	// 解析参数
	flag.Parse()

	return argConfig
}
