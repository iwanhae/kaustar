package authenticator

import (
	"context"
)

type AuthenticateService struct{}

func NewAuthenticateService() *AuthenticateService {
	return &AuthenticateService{}
}

func (s *AuthenticateService) Start(ctx context.Context) {

}
