package server

type ServerConfig struct {
	Host      string // 主机名
	Port      int    // 端口
	Url       string // 短链接地址
	IndexPath string // 跳转路径

	Redis string // redis连接地址，如：“127.0.0.1:6379”
}

// 默认配置参数
func DefaultServerConfig() *ServerConfig {
	cfg := ServerConfig{}
	cfg.Host = "0.0.0.0"
	cfg.Port = 9999
	cfg.Url = "http://127.0.0.1"
	cfg.IndexPath = "/"
	cfg.Redis = "127.0.0.1:6379"
	return &cfg
}
