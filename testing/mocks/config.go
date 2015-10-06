package mocks

import (
	"github.com/v2ray/v2ray-core/config"
)

type ConnectionConfig struct {
	ProtocolValue string
	SettingsValue interface{}
}

func (config *ConnectionConfig) Protocol() string {
	return config.ProtocolValue
}

func (config *ConnectionConfig) Settings(config.Type) interface{} {
	return config.SettingsValue
}

type Config struct {
	PortValue           uint16
	InboundConfigValue  *ConnectionConfig
	OutboundConfigValue *ConnectionConfig
}

func (config *Config) Port() uint16 {
	return config.PortValue
}

func (config *Config) InboundConfig() config.ConnectionConfig {
	return config.InboundConfigValue
}

func (config *Config) OutboundConfig() config.ConnectionConfig {
	return config.OutboundConfigValue
}
