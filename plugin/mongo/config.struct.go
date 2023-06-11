package xzmongo

type Config struct {
	Mongo []*ConfigE `yaml:"mongo"`
}

type ConfigE struct {
	Name               string            `yaml:"name"`
	Enabled            bool              `yaml:"enabled"`
	Description        string            `yaml:"description"`
	PingTimeout        int               `yaml:"pingTimeout"`
	CertEntry          string            `yaml:"certEntry"`
	InsecureSkipVerify bool              `yaml:"insecureSkipVerify"`
	Auth               *ConfigAuth       `yaml:"auth"`
	Hosts              []string          `yaml:"hosts"`
	ReplicaSet         string            `yaml:"replicaSet"`
	Database           []*ConfigDatabase `yaml:"database"`
}

type ConfigAuth struct {
	Mechanism           string            `yaml:"mechanism"`
	MechanismProperties map[string]string `yaml:"mechanismProperties"`
	Source              string            `yaml:"source"`
	Username            string            `yaml:"username"`
	Password            string            `yaml:"password"`
	PasswordSet         bool              `yaml:"passwordSet"`
}

type ConfigDatabase struct {
	Name string `yaml:"name"`
}
