package application

import (
	"errors"
	"grpc_go/src/internal/authpb"
	"time"
)

var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

type RSharing struct {
	Jwt     AuthInterface
	Paseto0 AuthInterface
}

type AccessPayload struct {
	Id        string
	IssuedAt  time.Time
	ExpiredAt time.Time
}

func (ap *AccessPayload) GetExpiredAt() time.Time {
	return ap.ExpiredAt
}

func (ap *AccessPayload) New() *AccessPayload {
	ap.IssuedAt = time.Now()
	ap.ExpiredAt = time.Now().Add(time.Minute * 45)
	return ap
}

func (ap *AccessPayload) NonValue() *AccessPayload {
	ap.IssuedAt = time.Now()
	ap.ExpiredAt = time.Now().Add(time.Hour * 24 * 1000)
	return ap
}

func (ap *AccessPayload) SendRpc() *authpb.DeJwtSecure {
	return &authpb.DeJwtSecure{
		Id: ap.Id,
	}
}

type RefreshPayload struct {
	Id        string
	IssuedAt  time.Time
	ExpiredAt time.Time
}

func (rp *RefreshPayload) GetExpiredAt() time.Time {
	return rp.ExpiredAt
}

func (rp *RefreshPayload) New() *RefreshPayload {
	rp.IssuedAt = time.Now()
	rp.ExpiredAt = time.Now().Add(time.Hour * 24 * 7)
	return rp
}

func (rp *RefreshPayload) NonValue() *RefreshPayload {
	rp.IssuedAt = time.Now()
	rp.ExpiredAt = time.Now().Add(time.Hour * 24 * 1000)
	return rp
}

func (rp *RefreshPayload) SendRpc() *authpb.DeJwtSecure {
	return &authpb.DeJwtSecure{
		Id: rp.Id,
	}
}

type RallyPayload interface {
	GetExpiredAt() time.Time
}

func Valid[EP RallyPayload](payload EP) error {
	if time.Now().After(payload.GetExpiredAt()) {
		return ErrExpiredToken
	}
	return nil
}

type AuthInterface interface {
	CreateAccessToken(payload *AccessPayload) (string, error)
	CreateRefreshToken(payload *RefreshPayload) (string, error)
	CreateNonValueAccessToken(payload *AccessPayload) (string, error)
	CreateNonValueRefreshToken(payload *RefreshPayload) (string, error)
	VerifyAccessToken(token string) (*AccessPayload, error)
	VerifyRefreshToken(token string) (*RefreshPayload, error)
}
