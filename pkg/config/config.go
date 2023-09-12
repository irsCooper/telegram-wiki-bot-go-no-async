package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Token    string
	Messages Messages
}

type Messages struct {
	Text
	Comands
}

type Text struct {
	Invalid_Comand string `mapstructure:"invalid_comand"`
	Bot_Is_Started string `mapstructure:"bot_is_started"`
	Hello          string `mapstructure:"hello"`
	New_Language   string `mapstructure:"new_language"`
	New_Query      string `mapstructure:"new_query"`
	Rus_Install    string `mapstructure:"rus_install"`
	Bel_Install    string `mapstructure:"bel_install"`
	En_Install     string `mapstructure:"en_install"`
	Nl_Install     string `mapstructure:"nl_install"`
	Invalid_Input  string `mapstructure:"invalid_input"`
	See_Comand     string `mapstructure:"see_comand"`
	Comand         string `mapstructure:"comand"`
}

type Comands struct {
	C_Start        string `mapstructure:"start"`
	Reset_Language string `mapstructure:"reset_language"`
	Random         string `mapstructure:"random"`
	Query          string `mapstructure:"query"`
	Help           string `mapstructure:"help"`
}

func Init() (*Config, error) {
	viper.AddConfigPath("configs")
	viper.SetConfigName("main")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	if err := viper.UnmarshalKey("messages.text", &cfg.Messages.Text); err != nil {
		return nil, err
	}

	if err := viper.UnmarshalKey("messages.comands", &cfg.Messages.Comands); err != nil {
		return nil, err
	}

	if err := parsEnv(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func parsEnv(cfg *Config) error {
	if err := viper.BindEnv("token"); err != nil {
		return err
	}

	cfg.Token = viper.GetString("token")
	return nil
}
