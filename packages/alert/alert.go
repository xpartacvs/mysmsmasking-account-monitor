package alert

import (
	"errors"
	"fmt"
	"regexp"
	"time"

	"github.com/leekchan/accounting"
	"github.com/rs/zerolog"
	"github.com/xpartacvs/go-dishook"
)

type alert struct {
	payload dishook.Payload
	webhook string
	logger  *zerolog.Logger
}

type Alert interface {
	Send() error
	SetBotName(name string) Alert
	SetBotAvatar(url string) Alert
	SetLogger(logger *zerolog.Logger) Alert
	SetLowBalanceReminder(balance int64, limit int64) Alert
	SetExpiryReminder(expiry time.Time, limit, remaining uint) Alert
}

const (
	errMsgNoEmbed        = "alert has nothing to send"
	errMsgInvalidWebhook = "invalid webhook url"
)

func (a *alert) Send() error {
	if a.payload.Embeds == nil {
		if a.logger != nil {
			a.logger.Warn().Msg(errMsgNoEmbed)
			return nil
		}
		return errors.New(errMsgNoEmbed)
	}

	rgxUrl := regexp.MustCompile("^https?://discord.com/api/webhooks/.*")
	if !rgxUrl.MatchString(a.webhook) {
		if a.logger != nil {
			a.logger.Warn().Msg(errMsgInvalidWebhook)
			return nil
		}
		return errors.New(errMsgInvalidWebhook)
	}

	_, err := dishook.Send(a.webhook, a.payload)
	a.payload.Embeds = nil
	if err != nil {
		if a.logger != nil {
			a.logger.Warn().Msg(err.Error())
			return nil
		}
		return err
	}

	return nil
}

func (a *alert) SetLogger(logger *zerolog.Logger) Alert {
	a.logger = logger
	return a
}

func (a *alert) SetBotName(name string) Alert {
	if len(name) > 0 {
		a.payload.Username = name
	}
	return a
}

func (a *alert) SetBotAvatar(url string) Alert {
	if len(url) > 0 {
		rgxUrl := regexp.MustCompile("^https?://.*")
		if !rgxUrl.MatchString(url) {
			if a.logger != nil {
				a.logger.Warn().Msg("invalid bot avatar url")
			}
			return a
		}
		a.payload.AvatarUrl = dishook.Url(url)
	}
	return a
}

func (a *alert) SetLowBalanceReminder(balance int64, limit int64) Alert {
	ac := accounting.Accounting{
		Symbol:   "Rp",
		Thousand: ".",
		Decimal:  ",",
		Format:   "%s %v",
	}

	moneyBalance := ac.FormatMoney(balance)
	moneyMargin := ac.FormatMoney(limit)
	title := "Saldo Akun Minim"
	desc := fmt.Sprintf("Saldo kurang dari %s Segera lakukan topup atau SMS tidak bisa terkirim.", moneyMargin)

	embed := dishook.Embed{
		// Url:         "https://raja-sms.com/topupsaldo/",
		Color:       dishook.ColorWarn,
		Title:       title,
		Description: desc,
		Fields: []dishook.Field{
			{
				Name:   "Saldo Sekarang",
				Value:  moneyBalance,
				Inline: false,
			},
		},
	}

	a.payload.Embeds = nil
	a.payload.Embeds = append(a.payload.Embeds, embed)

	return a
}

func (a *alert) SetExpiryReminder(expiry time.Time, limit, remaining uint) Alert {
	title := "Mendekati Tanggal Kedaluarsa"
	desc := "Masa aktif saldo akun hampir berakhir. Segera lakukan topup, atau saldo hangus."

	embed := dishook.Embed{
		// Url:         "https://raja-sms.com/topupsaldo/",
		Color:       dishook.ColorWarn,
		Title:       title,
		Description: desc,
		Fields: []dishook.Field{
			{
				Name:   "Tanggal Kedaluarsa",
				Value:  expiry.Format("_2 Jan 2006"),
				Inline: true,
			},
			{
				Name:   "Saldo Hangus Dalam",
				Value:  fmt.Sprintf("%d hari", remaining),
				Inline: true,
			},
		},
	}

	a.payload.Embeds = nil
	a.payload.Embeds = append(a.payload.Embeds, embed)

	return a
}

func New(webhookUrl, message string) Alert {
	return &alert{
		webhook: webhookUrl,
		payload: dishook.Payload{Content: message},
	}
}
