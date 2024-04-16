package internal

import (
	"context"
	"errors"
	"go-keep/internal/config"

	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
)

// Authenticator is used to authenticate our users.
type Authenticator struct {
	*oidc.Provider
	oauth2.Config
}

// NewAuth instantiates the *Authenticator.
func NewAuth(conf *config.Configuration) (*Authenticator, error) {
	provider, err := oidc.NewProvider(
		context.Background(),
		"https://"+conf.Auth.Domain+"/",
	)
	if err != nil {
		return nil, err
	}

	oauth2conf := oauth2.Config{
		ClientID:     conf.Auth.ClientID,
		ClientSecret: conf.Auth.ClientSecret,
		RedirectURL:  conf.Auth.CallbackURL,
		Endpoint:     provider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "profile"},
	}

	return &Authenticator{
		Provider: provider,
		Config:   oauth2conf,
	}, nil
}

// VerifyIDToken verifies that an *oauth2.Token is a valid *oidc.IDToken.
func (a *Authenticator) VerifyIDToken(ctx context.Context, token *oauth2.Token) (*oidc.IDToken, error) {
	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		return nil, errors.New("no id_token field in oauth2 token")
	}

	oidcConfig := &oidc.Config{
		ClientID: a.ClientID,
	}

	return a.Verifier(oidcConfig).Verify(ctx, rawIDToken)
}
