package user

import (
	"go-keep/cmd/api"
	"go-keep/pkg/session"
	"html/template"
	"net/http"
)

type UserService struct {
	pkg *session.UserPkg
}

func NewUserService(pkg api.Packager) *UserService {
	userPkg := pkg.NewUserPkg()
	return &UserService{userPkg}
}

func (u *UserService) login(w http.ResponseWriter, r *http.Request) {

	sess := &session.Session{}

	redirectUrl := u.pkg.Put(w, r, sess)

	http.Redirect(w, r, redirectUrl, http.StatusTemporaryRedirect)
}

func (u *UserService) logout(w http.ResponseWriter, r *http.Request) {
	redirectUrl, err := u.pkg.Remove(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	http.Redirect(w, r, redirectUrl, http.StatusTemporaryRedirect)
}

func (u *UserService) user(w http.ResponseWriter, r *http.Request) {
	user := u.pkg.Get(w, r)
	t, _ := template.New("foo").Parse(userTemplate)
	t.Execute(w, user)

}

func (u *UserService) callback(w http.ResponseWriter, r *http.Request) {
	u.pkg.Verify(w, r)

	redirectUrl := "/user"
	http.Redirect(w, r, redirectUrl, http.StatusTemporaryRedirect)
}

var userTemplate = `
<p>Name: {{.Name}}</p>
<p>AvatarURL: {{.Picture}} <img src="{{.AvatarURL}}"></p>
`
