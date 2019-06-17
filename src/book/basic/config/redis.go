package config

import "strings"

type RedisConfig interface {
	GetEnabled() bool
	GetConn() string
	GetPassword() string
	GetDBNum() int
	GetSentinelConfig() RedisSentinelConfig
}

type RedisSentinelConfig interface {
	GetEnabled() bool
	GetMaster() string
	GetNodes() []string
}

type defaultRedisConfig struct {
	Enabled  bool          `json:"enabled"`
	Conn     string        `json:"conn"`
	Password string        `json:"password"`
	DBNum    int           `json:"dbNum"`
	TimeOut  int           `json:"timeout"`
	sentinel redisSentinel `json:"sentinel"`
}

type redisSentinel struct {
	Enable bool   `json:"enabled"`
	Master string `json:"master"`
	Nodes  string `json:"nodes"`
	nodes  []string
}

func (r defaultRedisConfig) GetEnabled() bool {
	return r.Enabled
}

func (r defaultRedisConfig) GetConn() string {
	return r.Conn
}

func (r defaultRedisConfig) GetPassword() string {
	return r.Password
}

func (r defaultRedisConfig) GetDBNum() int {
	return r.DBNum
}

func (r defaultRedisConfig) GetSentinelConfig() RedisSentinelConfig {
	return r.sentinel
}

func (s redisSentinel) GetEnabled() bool {
	return s.Enable
}
func (s redisSentinel) GetMaster() string {
	return s.Master
}

func (s redisSentinel) GetNodes() []string {
	if len(s.Nodes) > 0 {
		for _, val := range strings.Split(s.Nodes, ",") {
			val := strings.TrimSpace(val)
			s.nodes = append(s.nodes, val)
		}
	}
	return s.nodes
}
