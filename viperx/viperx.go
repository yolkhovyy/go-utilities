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

	if envKeyPreplacer != nil {
		vpr.SetEnvKeyReplacer(envKeyPreplacer)
	} else {
		vpr.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	}

	return Viperx{Viper: vpr}
}

func (v Viperx) Load(config any) error {
	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("read config: %w", err)
	}

	if err := v.Unmarshal(config); err != nil {
		return fmt.Errorf("unnmarshal: %w", err)
	}

	return nil
}
