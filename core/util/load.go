package xzutil

import (
	"embed"
	"os"
	"path/filepath"

	"github.com/mitchellh/mapstructure"
	"gopkg.in/yaml.v2"
)

// 读取yaml文件 默认文件boot.yaml
func ReadYAML(configPath string) ([]byte, error) {
	// 判断是否为绝对路径
	if !filepath.IsAbs(configPath) {
		wd, _ := os.Getwd()
		configPath = filepath.Join(wd, configPath)
	}

	return os.ReadFile(configPath)
}

func ReadFile(filePath string, fs *embed.FS) ([]byte, error) {
	if fs != nil {
		data, err := fs.ReadFile(filePath)
		if err != nil {
			return nil, err
		}
		return data, nil
	}

	wd, _ := os.Getwd()

	if !filepath.IsAbs(filePath) {
		filePath = filepath.Join(wd, filePath)
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return data, nil
	}
	return data, nil
}

// 序列化配置
func UnmarshalConfigYAML(raw []byte, config any) error {
	originalBootM := make(map[any]any)
	if err := yaml.Unmarshal(raw, &originalBootM); err != nil {
		return err
	}
	if err := mapstructure.Decode(originalBootM, config); err != nil {
		return err
	}
	return nil

}
