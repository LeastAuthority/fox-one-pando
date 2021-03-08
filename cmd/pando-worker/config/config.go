package config

import (
	"github.com/fox-one/mixin-sdk-go"
	"github.com/fox-one/pkg/store/db"
	jsoniter "github.com/json-iterator/go"
	"github.com/shopspring/decimal"
	"github.com/spf13/viper"
	"golang.org/x/text/language"
)

type (
	Config struct {
		DB    db.Config `json:"db"`
		Dapp  Dapp      `json:"dapp"`
		Group Group     `json:"group,omitempty"`
		I18n  I18n      `json:"i18n,omitempty"`
	}

	Dapp struct {
		mixin.Keystore
		Pin string `json:"pin"`
	}

	Vote struct {
		Asset  string          `json:"asset,omitempty"`
		Amount decimal.Decimal `json:"amount,omitempty"`
	}

	Group struct {
		// 节点管理员 mixin id
		Admins []string `json:"admins,omitempty"`
		// 节点用于签名的私钥
		SignKey string `json:"sign_key,omitempty"`
		// 节点共享的用户解密的私钥
		PrivateKey string   `json:"private_key,omitempty"`
		Members    []string `json:"members,omitempty"`
		Threshold  uint8    `json:"threshold,omitempty"`

		Vote Vote `json:"vote,omitempty"`
	}

	I18n struct {
		Path string `json:"path,omitempty"`
		// default language
		Language string `json:"language,omitempty"`
	}
)

// Viperion load config by viper
func Viperion(cfgFile string) (*Config, error) {
	v := viper.New()

	if cfgFile != "" {
		v.SetConfigFile(cfgFile)
	} else {
		v.SetConfigType("yaml")
		v.SetConfigName("config")
		v.AddConfigPath("/etc/pando/worker")
		v.AddConfigPath("$HOME/.pando/worker")
		v.AddConfigPath(".")
	}

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	b, err := jsoniter.Marshal(v.AllSettings())
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := jsoniter.Unmarshal(b, &cfg); err != nil {
		return nil, err
	}

	defaultVote(&cfg)
	defaultI18n(&cfg)

	return &cfg, nil
}

func defaultVote(cfg *Config) {
	if cfg.Group.Vote.Asset == "" {
		cfg.Group.Vote.Asset = "965e5c6e-434c-3fa9-b780-c50f43cd955c" // cnb
	}

	if cfg.Group.Vote.Amount.IsZero() {
		cfg.Group.Vote.Amount = decimal.NewFromInt(1)
	}
}

func defaultI18n(cfg *Config) {
	if cfg.I18n.Path == "" {
		cfg.I18n.Path = "./assets/i18n"
	}

	if cfg.I18n.Language == "" {
		cfg.I18n.Language = language.Chinese.String()
	}
}
