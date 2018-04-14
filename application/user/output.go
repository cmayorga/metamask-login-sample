package user

import "github.com/m0t0k1ch1/metamask-login-sample/domain"

type GetUserOutput struct {
	Name       string `json:"name"`
	AddressHex string `json:"address"`
}

func NewGetUserOutput(user *domain.User) *GetUserOutput {
	return &GetUserOutput{
		Name:       user.Name,
		AddressHex: user.Address.Hex(),
	}
}