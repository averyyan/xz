package xzjwt

type Config struct {
	Jwt []*ConfigE `yaml:"jwt"`
}

type ConfigE struct {
	Name        string          `yaml:"name"`        // 名称
	Description string          `yaml:"description"` // 描述
	Enabled     bool            `yaml:"enabled"`     // 是否启用
	Symmetric   SymmetricConfig `yaml:"symmetric"`   // 对称加密配置
	Ignore      []string        `yaml:"ignore"`      // 跳过检测
}
