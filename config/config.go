package config

type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
	Webhook  WebhookConfig  `yaml:"webhook"`
}

type ServerConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Name     string `yaml:"name"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type WebhookConfig struct {
	Wechat    WebhookWechatConfig    `yaml:"wechat"`
	OurServer WebhookOurServerConfig `yaml:"ourserver"`
}

type WebhookWechatConfig struct {
	Key string `yaml:"key"`
}

type WebhookOurServerConfig struct {
	Referer  string `yaml:"referer"`
	URL      string `yaml:"url"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}
