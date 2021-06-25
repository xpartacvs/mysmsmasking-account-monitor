package worker

import (
	"mysmsmasking-account-monitor/packages/alert"
	"mysmsmasking-account-monitor/packages/config"
	"mysmsmasking-account-monitor/packages/logger"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/xpartacvs/go-mysmsmasking"
)

func Start() error {
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		logger.Log().Err(err)
		loc = time.Local
	}

	cron := gocron.NewScheduler(loc)
	_, err = cron.Cron(config.Get().Schedule()).Do(do)
	if err != nil {
		logger.Log().Err(err)
		return err
	}
	cron.StartBlocking()

	return nil
}

func do() {
	sms := mysmsmasking.New(
		config.Get().MySMSMaskingURL(),
		config.Get().MySMSMaskingUser(),
		config.Get().MySMSMaskingPassword(),
	)

	acc, err := sms.GetAccountInfo()
	if err != nil {
		logger.Log().Err(err)
		return
	}

	if acc.Balance <= uint32(config.Get().BalanceLimit()) {
		notif := alert.New(
			config.Get().DishookURL(),
			config.Get().DishookBotMessage(),
		)

		err := notif.SetLogger(logger.Log()).
			SetBotName(config.Get().DishookBotName()).
			SetBotAvatar(config.Get().DishookBotAvatarURL()).
			SetLowBalanceReminder(acc.Balance, config.Get().BalanceLimit()).
			Send()
		if err != nil {
			logger.Log().Err(err)
		}
		return
	}

	remaining := uint(time.Until(acc.Expiry).Hours() / 24)
	if remaining <= uint(config.Get().GracePeriod()) {
		notif := alert.New(
			config.Get().DishookURL(),
			config.Get().DishookBotMessage(),
		)

		err := notif.SetLogger(logger.Log()).
			SetBotName(config.Get().DishookBotName()).
			SetBotAvatar(config.Get().DishookBotAvatarURL()).
			SetExpiryReminder(acc.Expiry, config.Get().GracePeriod(), remaining).
			Send()
		if err != nil {
			logger.Log().Err(err)
		}
	}

}
