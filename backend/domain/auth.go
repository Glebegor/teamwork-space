package domain

import (
	"context"
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation"
)

const AuthCollection = "users"

var passwordRule = []validation.Rule{
	validation.Required,
	validation.Length(8, 1000),
	validation.Match(regexp.MustCompile("^\\S+$")).Error("cannot contain whitespaces"),
}

type Reg struct {
	Username string `form:"username" json:"username"`
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
}

func (r Reg) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Username, validation.Required, validation.Length(4, 128)),
		validation.Field(&r.Email, validation.Required, validation.Length(4, 256)),
		validation.Field(&r.Password, passwordRule...),
	)
}

type Login struct {
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
}
type Refresh struct {
	RefreshToken string `form:"refreshToken" json:"refreshToken"`
}

type LoginResponse struct {
	AccessToken  string `form:"accessToken" json:"accessToken"`
	RefreshToken string `form:"refreshToken" json:"refreshToken"`
}
type RefreshTokenResponse struct {
	AccessToken  string `form:"accessToken" json:"accessToken"`
	RefreshToken string `form:"refreshToken" json:"refreshToken"`
}

type AuthRepository interface {
	Create(c context.Context, input User) error
	GetByEmail(c context.Context, email string) (User, error)
	GetById(c context.Context, id string) (User, error)
}
type AuthUsecase interface {
	GetByEmail(c context.Context, email string) (User, error)
	Register(c context.Context, input User) error
	CreateAccessToken(user *User, secret string, expiry int) (string, error)
	CreateRefreshToken(user *User, secret string, expiry int) (string, error)
	GetIdFromRefreshToken(requestToken string, secret string) (string, error)
	GetUserById(c context.Context, id string) (User, error)
	EncryptPassword(password string, secret string) (string, error)
}
