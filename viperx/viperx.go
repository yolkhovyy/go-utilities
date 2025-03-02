package viperx

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type Viperx struct {
	*viper.Viper
}

func New(configFile string, prefix string, envKeyPreplacer *strings.Replacer) Viperx {
	vpr := viper.New()
	vpr.SetConfigFile(configFile)
	vpr.SetEnvPrefix(prefix)
	vpr.AutomaticEnv()

	if envKeyPreplacer == nil {
		envKeyPreplacer = strings.NewReplacer(".", "_", "-", "_")
	}

	vpr.SetEnvKeyReplacer(envKeyPreplacer)

	return Viperx{Viper: vpr}
}

func (v Viperx) SetDefaults(defaults map[string]any) {
	for key, value := range defaults {
		v.SetDefault(key, value)
	}
}

func (v Viperx) Load(config any, opts ...viper.DecoderConfigOption) error {
	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("read config: %w", err)
	}

	if err := v.Unmarshal(config, opts...); err != nil {
		return fmt.Errorf("unmarshal: %w", err)
	}

	return nil
}
