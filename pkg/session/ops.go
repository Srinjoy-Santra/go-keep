package session

import (
	"go-keep/internal"
	"go-keep/internal/config"
	"log"
	"net/http"
)

type UserPkg struct {
	config *config.Configuration
	ss     *SessionStore[Session]
	auth   *internal.Authenticator
}

type Session struct {
	Role        string
	Username    string
	AccessToken string
	Profile     map[string]interface{}
}

type UseProfile struct {
	Name    string
	Picture string
}

func NewUserPkg(conf *config.Configuration, ss *SessionStore[Session], auth *internal.Authenticator) *UserPkg {
	return &UserPkg{config: conf, ss: ss, auth: auth}
}

func (pkg *UserPkg) Put(w http.ResponseWriter, r *http.Request, sess *Session) string {
	state := pkg.ss.PutSession(w, r, sess)
	return pkg.auth.AuthCodeURL(state)
}

func (pkg *UserPkg) Verify(w http.ResponseWriter, r *http.Request) error {
	log.Println(r.URL.RawQuery)
	query := r.URL.Query()

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

	session := pkg.ss.GetSessionFromCtx(r)
	session.AccessToken = token.AccessToken
	session.Profile = profile

	return nil
}

func (pkg *UserPkg) Get(w http.ResponseWriter, r *http.Request) UseProfile {
	session := pkg.ss.GetSessionFromCtx(r)
	profile := session.Profile
	log.Println(profile)

	user := UseProfile{
		Name:    profile["name"].(string),
		Picture: profile["picture"].(string),
	}

	return user
}
