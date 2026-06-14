package main

import (
	"context"
	"quanteye_utils/buffer"
	"strings"
	"time"

	wechatbot "github.com/corespeed-io/wechatbot/golang"
)

type MessageHandler func(msg string) string
type UrlHandler func(url string)
type Weixin struct {
	ctx        context.Context
	bot        *wechatbot.Bot
	creds      *wechatbot.Credentials
	Handler    MessageHandler
	UrlHandler UrlHandler
	buffer     *buffer.BufferSender[string]
	last       string
}

func (w *Weixin) Init() {
	w.buffer = buffer.NewBufferSender(0, 20*time.Second, w.commitMessage)
	w.ctx = context.Background()
	w.bot = wechatbot.New(wechatbot.Options{OnQRURL: w.UrlHandler})
	w.creds, _ = w.bot.Login(w.ctx, false)

	w.bot.OnMessage(func(msg *wechatbot.IncomingMessage) {
		w.bot.SendTyping(w.ctx, msg.UserID)
		if !IsOldMessage(msg.Timestamp) {
			txt := w.Handler(msg.Text)
			if txt == "" {
				txt = w.last
			}
			w.bot.Reply(w.ctx, msg, txt)
		}

	})
}

func (w *Weixin) Start() {
	go w.bot.Run(w.ctx)
}

func (w *Weixin) Notify(msg string) {
	if msg != "" {
		w.buffer.Add(msg)
	}
}

func (w *Weixin) commitMessage(msgs []string) {
	msg := strings.Join(msgs, "\n")
	w.last = msg
	w.bot.Send(w.ctx, w.creds.UserID, msg)
}

var StartTime = time.Now()

// IsOldMessage returns true if msgTime is before the process StartTime.
// A small grace period (2 seconds) is applied to avoid race conditions
// with messages sent right at startup.
func IsOldMessage(msgTime time.Time) bool {
	return msgTime.Before(StartTime.Add(-2 * time.Second))
}
