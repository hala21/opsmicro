package config

type Profiles interface {
	GetInclude() string
}

type defaultProfiles struct {
	Include string `json:"include"`
}

// 格式为：前缀"application-",格式为yaml,如 "application-xxx.yaml"
// 多个文件名以逗号隔开，并省略掉前缀"application-",如:dn, jpush, mysql
func (p defaultProfiles) GetInclude() string {
	return p.Include
}
