package xzcert

type Config struct {
	Cert []*ConfigE `yaml:"cert"`
}

type ConfigE struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Enabled     bool   `yaml:"enabled"`
	CAPath      string `yaml:"caPath"`
	CertPemPath string `yaml:"certPemPath"`
	KeyPemPath  string `yaml:"keyPemPath"`
}
