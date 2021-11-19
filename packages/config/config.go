package config

import (
	"os"
	"strings"
	"sync"

	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

type Config struct {
	schedule     string
	logMode      zerolog.Level
	msmUser      string
	msmPass      string
	limitBalance int64
	limitPeriod  uint
	disHook      string
	disBotName   string
	disBotAva    string
	disBotMsg    string
}

var (
	cfg  *Config
	once sync.Once
)

func (c Config) Schedule() string {
	return c.schedule
}

func (c Config) ZerologLevel() zerolog.Level {
	return c.logMode
}

func (c Config) MySMSMaskingUser() string {
	return c.msmUser
}

func (c Config) MySMSMaskingPassword() string {
	return c.msmPass
}

func (c Config) BalanceLimit() int64 {
	return c.limitBalance
}

func (c Config) GracePeriod() uint {
	return c.limitPeriod
}

func (c Config) DishookURL() string {
	return c.disHook
}

func (c Config) DishookBotName() string {
	return c.disBotName
}

func (c Config) DishookBotAvatarURL() string {
	return c.disBotAva
}

func (c Config) DishookBotMessage() string {
	return c.disBotMsg
}

func Get() *Config {
	once.Do(func() {
		cfg = read()
	})
	return cfg
}

func read() *Config {
	fang := viper.New()

	fang.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	fang.AutomaticEnv()

	fang.SetConfigName("mysmsmasking-account-monitor")
	fang.SetConfigType("yml")
	fang.AddConfigPath(".")

	value, available := os.LookupEnv("CONFIGDIR_PATH")
	if available {
		fang.AddConfigPath(value)
	}

	_ = fang.ReadInConfig()

	balance := fang.GetInt64("balance.limit")
	if balance == 0 {
		balance = 300000
	}

	period := fang.GetUint("grace.period")
	if period == 0 {
		period = 14
	}

	var logmode zerolog.Level
	switch fang.GetString("logmode") {
	case "debug":
		logmode = zerolog.DebugLevel
	case "info":
		logmode = zerolog.InfoLevel
	case "warn":
		logmode = zerolog.WarnLevel
	case "error":
		logmode = zerolog.ErrorLevel
	default:
		logmode = zerolog.Disabled
	}

	return &Config{
		schedule:     setDefaultString(fang.GetString("schedule"), "0 0 * * *", true),
		logMode:      logmode,
		msmUser:      setDefaultString(fang.GetString("mysmsmasking.user"), "", true),
		msmPass:      setDefaultString(fang.GetString("mysmsmasking.password"), "", true),
		limitBalance: balance,
		limitPeriod:  period,
		disHook:      setDefaultString(fang.GetString("discord.webhookurl"), "", true),
		disBotName:   setDefaultString(fang.GetString("discord.bot.name"), "MySMSMasking Monitor", true),
		disBotAva:    setDefaultString(fang.GetString("discord.bot.avatarurl"), "https://www.seekpng.com/png/small/139-1394319_sms-icon-smoking-signs-to-print.png", true),
		disBotMsg:    setDefaultString(fang.GetString("discord.bot.message"), "Reminder akun MySMSMasking", true),
	}
}

func setDefaultString(value, fallback string, trimSpace bool) string {
	if trimSpace {
		value = strings.TrimSpace(value)
	}
	if len(value) <= 0 {
		return fallback
	}
	return value
}
