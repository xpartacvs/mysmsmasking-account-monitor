package config

import (
	"os"
	"strings"
	"sync"

	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

type config struct {
	schedule     string
	logMode      zerolog.Level
	msmUrl       string
	msmUser      string
	msmPass      string
	limitBalance uint
	limitPeriod  uint
	disHook      string
	disBotName   string
	disBotAva    string
	disBotMsg    string
}

type Config interface {
	ZerologLevel() zerolog.Level
	MySMSMaskingURL() string
	MySMSMaskingUser() string
	MySMSMaskingPassword() string
	BalanceLimit() uint
	GracePeriod() uint
	DishookURL() string
	DishookBotName() string
	DishookBotAvatarURL() string
	DishookBotMessage() string
	Schedule() string
}

var (
	cfg  Config
	once sync.Once
)

func (c config) Schedule() string {
	return c.schedule
}

func (c config) ZerologLevel() zerolog.Level {
	return c.logMode
}

func (c config) MySMSMaskingURL() string {
	return c.msmUrl
}

func (c config) MySMSMaskingUser() string {
	return c.msmUser
}

func (c config) MySMSMaskingPassword() string {
	return c.msmPass
}

func (c config) BalanceLimit() uint {
	return c.limitBalance
}

func (c config) GracePeriod() uint {
	return c.limitPeriod
}

func (c config) DishookURL() string {
	return c.disHook
}

func (c config) DishookBotName() string {
	return c.disBotName
}

func (c config) DishookBotAvatarURL() string {
	return c.disBotAva
}

func (c config) DishookBotMessage() string {
	return c.disBotMsg
}

func Get() Config {
	once.Do(func() {
		cfg = read()
	})
	return cfg
}

func read() Config {
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

	balance := fang.GetUint("balance.limit")
	if balance == 0 {
		balance = 100000
	}

	period := fang.GetUint("grace.period")
	if period == 0 {
		period = 7
	}

	botMsg := fang.GetString("discord.bot.message")
	if len(strings.TrimSpace(botMsg)) == 0 {
		botMsg = "Reminder akun MySMSMasking"
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

	return &config{
		schedule:     strings.TrimSpace(fang.GetString("schedule")),
		logMode:      logmode,
		msmUrl:       strings.TrimSpace(fang.GetString("mysmsmasking.url")),
		msmUser:      strings.TrimSpace(fang.GetString("mysmsmasking.user")),
		msmPass:      strings.TrimSpace(fang.GetString("mysmsmasking.password")),
		limitBalance: balance,
		limitPeriod:  period,
		disHook:      strings.TrimSpace(fang.GetString("discord.webhookurl")),
		disBotName:   strings.TrimSpace(fang.GetString("discord.bot.name")),
		disBotAva:    strings.TrimSpace(fang.GetString("discord.bot.avatarurl")),
		disBotMsg:    botMsg,
	}
}
