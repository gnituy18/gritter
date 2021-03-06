package auth

import (
	"go.uber.org/zap/zapcore"
	"google.golang.org/api/oauth2/v2"
)

type Type int

const (
	TypeGoogle Type = iota + 1
)

type Info struct {
	Type Type `json:"type"`

	Google *InfoGoogle `json:"google"`
}

type InfoGoogle struct {
	AccessToken string `json:"accessToken"`
}

type Result struct {
	Type Type

	Google *ResultGoogle
}

func (r *Result) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	encoder.AddInt("type", int(r.Type))
	encoder.AddObject("google", r.Google)
	return nil
}

type ResultGoogle struct {
	UserInfo *oauth2.Userinfo
}

func (rg *ResultGoogle) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	bytes, err := rg.UserInfo.MarshalJSON()
	if err != nil {
		return err
	}

	encoder.AddString("accessToken", string(bytes))
	return nil
}
