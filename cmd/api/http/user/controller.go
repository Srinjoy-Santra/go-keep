package user

import (
	"go-keep/cmd/api"
	"go-keep/pkg/session"
	"html/template"
	"net/http"
)

type UserService struct {
	pkg api.Packager
}

func NewUserService(pkg api.Packager) *UserService {
	return &UserService{pkg}
}

func (u *UserService) login(w http.ResponseWriter, r *http.Request) {

	sess := &session.Session{}

	userPkg := u.pkg.NewUserPkg()
	redirectUrl := userPkg.Put(w, r, sess)

	http.Redirect(w, r, redirectUrl, http.StatusTemporaryRedirect)
}

func (u *UserService) logout(w http.ResponseWriter, r *http.Request) {
	userPkg := u.pkg.NewUserPkg()

	redirectUrl, err := userPkg.Remove(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	http.Redirect(w, r, redirectUrl, http.StatusTemporaryRedirect)
}

func (u *UserService) user(w http.ResponseWriter, r *http.Request) {
	userPkg := u.pkg.NewUserPkg()
	user := userPkg.Get(w, r)
	t, _ := template.New("foo").Parse(userTemplate)
	t.Execute(w, user)

}

func (u *UserService) callback(w http.ResponseWriter, r *http.Request) {
	userPkg := u.pkg.NewUserPkg()
	userPkg.Verify(w, r)

	redirectUrl := "/user"
	http.Redirect(w, r, redirectUrl, http.StatusTemporaryRedirect)
}

var userTemplate = `
<p>Name: {{.Name}}</p>
<p>AvatarURL: {{.Picture}} <img src="{{.AvatarURL}}"></p>
`
