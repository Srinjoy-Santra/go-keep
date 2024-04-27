package session

import (
	"errors"
	"go-keep/internal"
	"go-keep/internal/config"
	"log"
	"net/http"
	"net/url"
)

type UserPkg struct {
	config       *config.Configuration
	SessionStore *SessionStore[Session]
	auth         *internal.Authenticator
}

type Session struct {
	State       string
	AccessToken string
	Profile     map[string]interface{}
}

type UseProfile struct {
	Name    string
	Picture string
}

func NewUserPkg(conf *config.Configuration, ss *SessionStore[Session], auth *internal.Authenticator) *UserPkg {
	return &UserPkg{config: conf, SessionStore: ss, auth: auth}
}

func (pkg *UserPkg) Put(w http.ResponseWriter, r *http.Request, sess *Session) string {
	state := pkg.SessionStore.PutSession(w, r, sess, "")
	return pkg.auth.AuthCodeURL(state)
}

func (pkg *UserPkg) Verify(w http.ResponseWriter, r *http.Request) error {
	log.Println(r.URL.RawQuery)
	query := r.URL.Query()

	state := query.Get("state")
	sessions := pkg.SessionStore.sessions
	session, found := sessions[state]
	if !found {
		return errors.New("invalid state parameter")
	}

	token, err := pkg.auth.Exchange(r.Context(), query.Get("code"))
	if err != nil {
		return err
	}
	idToken, err := pkg.auth.VerifyIDToken(r.Context(), token)
	if err != nil {
		return err
	}

	var profile map[string]interface{}
	if err := idToken.Claims(&profile); err != nil {
		return err
	}

	session.AccessToken = token.AccessToken
	session.Profile = profile
	pkg.SessionStore.PutSession(w, r, session, state)

	return nil
}

func (pkg *UserPkg) Get(w http.ResponseWriter, r *http.Request) (*UseProfile, error) {
	session := pkg.SessionStore.GetSessionFromRequest(r)
	if session == nil {
		return nil, errors.New("user not authenticated")
	}
	profile := session.Profile
	log.Println(profile)

	user := UseProfile{
		Name:    profile["name"].(string),
		Picture: profile["picture"].(string),
	}

	return &user, nil
}

func (pkg *UserPkg) Remove(w http.ResponseWriter, r *http.Request) (string, error) {

	logoutUrl, err := url.Parse("https://" + pkg.config.Auth.Domain + "/v2/logout")
	if err != nil {
		return "", err
	}

	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}

	returnTo, err := url.Parse(scheme + "://" + r.Host)
	if err != nil {
		return "", err
	}

	parameters := url.Values{}
	parameters.Add("returnTo", returnTo.String())
	parameters.Add("client_id", pkg.config.Auth.ClientID)
	logoutUrl.RawQuery = parameters.Encode()

	pkg.SessionStore.DeleteSession(r)

	return logoutUrl.String(), nil

}
