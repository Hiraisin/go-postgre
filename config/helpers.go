package config

import (
	"html/template"
	"io"
	"net/http"
	"time"

	"github.com/gorilla/securecookie"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

type M map[string]interface{}

type Renderer struct {
	template *template.Template
	debug    bool
	location string
}

//secureCookie
var sc = securecookie.New([]byte("very-secret"), []byte("a-lot-secret-yay"))

func NewRenderer(location string, debug bool) *Renderer {
	tpl := new(Renderer)
	tpl.location = location
	tpl.debug = debug

	tpl.ReloadTemplates()

	return tpl
}

func (t *Renderer) ReloadTemplates() {
	t.template = template.Must(template.ParseGlob(t.location))
}

func (t *Renderer) Render(
	w io.Writer,
	name string,
	data interface{},
	c echo.Context,
) error {
	if t.debug {
		t.ReloadTemplates()
	}

	return t.template.ExecuteTemplate(w, name, data)
}

//password bcrypt
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

//secureCookie
func SetCookie(c echo.Context, name string, data M) error {
	encoded, err := sc.Encode(name, data)
	if err != nil {
		return err
	}

	cookie := &http.Cookie{
		Name:     name,
		Value:    encoded,
		Path:     "/",
		Secure:   false,
		HttpOnly: true,
		Expires:  time.Now().Add(1 * time.Hour),
	}
	http.SetCookie(c.Response(), cookie)

	return nil
}

func GetCookie(c echo.Context, name string) (M, error) {
	cookie, err := c.Request().Cookie(name)
	if err == nil {

		data := M{}
		if err = sc.Decode(name, cookie.Value, &data); err == nil {
			return data, nil
		}
	}

	return nil, err
}
