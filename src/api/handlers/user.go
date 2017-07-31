package handlers

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"

	"github.com/alsx/enli-task/src/api/models"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

// UserClaims contains user attributes and is extended with custom claims.
type UserClaims struct {
	Email string
	jwt.StandardClaims
}

// TODO: move to settings
var oauthConf = &oauth2.Config{
	ClientID:     "329049282165041",
	ClientSecret: "048b0ff413d0c063077ed7d47af3db55",
	RedirectURL:  "http://example.com/api/v1/fb-callback/",
	Scopes:       []string{"public_profile", "email"},
	Endpoint:     facebook.Endpoint,
}
var oauthStateString = "thisshouldberandom"

// Handler base for controller
type Handler struct{}

// UserHandler contains methods to work with user entity
var UserHandler = Handler{}

// Info get user info
func (h *Handler) Info(c echo.Context) error {
	cxtUser := c.Get("user").(*jwt.Token)
	claims := cxtUser.Claims.(*UserClaims)
	email := claims.Email
	userMgr := models.NewUserManager(c)
	user, err := userMgr.GetByEmail(email)
	if err != nil {
		return c.JSONPretty(http.StatusOK, echo.Map{"error": err}, "    ")
	}
	return c.JSONPretty(http.StatusOK, user, "    ")
}

// LogIn authenticate user in api
func (h *Handler) LogIn(c echo.Context) error {
	email := c.FormValue("email")
	passwordSha256 := fmt.Sprintf("%x", sha256.Sum256([]byte(c.FormValue("password"))))
	userMgr := models.NewUserManager(c)
	user, err := userMgr.GetByEmail(email)
	if err != nil {
		return echo.ErrUnauthorized
	}
	if passwordSha256 == user.Password {
		// Set custom claims
		claims := &UserClaims{
			user.Email,
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			},
		}
		// Create token with claims
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		// Generate encoded token and send it as response.
		secret := []byte(c.Get("secret").(string))
		t, err := token.SignedString(secret)
		if err != nil {
			return err
		}
		return c.JSONPretty(http.StatusOK, echo.Map{
			"token": t,
		}, "    ")
	}
	return echo.ErrUnauthorized
}

// SignIn register user using email and password.
func (h *Handler) SignIn(c echo.Context) error {
	email := c.FormValue("email")
	if email == "" {
		return c.JSONPretty(http.StatusBadRequest, echo.Map{"error": "Invalid email."}, "    ")
	}
	name := c.FormValue("name")
	// sha256 of password
	passwordSha256 := fmt.Sprintf("%x", sha256.Sum256([]byte(c.FormValue("password"))))
	user := models.User{
		Email:    email,
		Name:     name,
		Password: passwordSha256,
	}

	userMgr := models.NewUserManager(c)
	id, err := userMgr.Add(user)
	if err != nil {
		return c.JSONPretty(http.StatusInternalServerError, echo.Map{"error": err}, "    ")
	}

	user.ID = id
	return c.JSONPretty(http.StatusCreated, user, "    ")
}

// FacebookLogIn login user with facebook account
func (h *Handler) FacebookLogIn(c echo.Context) error {

	URL, err := url.Parse(oauthConf.Endpoint.AuthURL)
	if err != nil {
		log.Fatal("Parse: ", err)
	}
	parameters := url.Values{}
	parameters.Add("client_id", oauthConf.ClientID)
	parameters.Add("scope", strings.Join(oauthConf.Scopes, " "))
	parameters.Add("redirect_uri", oauthConf.RedirectURL)
	parameters.Add("response_type", "code")
	parameters.Add("state", oauthStateString)
	URL.RawQuery = parameters.Encode()
	url := URL.String()
	return c.JSONPretty(http.StatusOK, echo.Map{"redirect_url": url}, "    ")
}

// FacebookCallback receive token from fb.
// TODO: add validation, save name, email and token into db.
func (h *Handler) FacebookCallback(c echo.Context) error {
	code := c.FormValue("code")
	token, _ := oauthConf.Exchange(oauth2.NoContext, code)
	resp, _ := http.Get("https://graph.facebook.com/me?access_token=" + url.QueryEscape(token.AccessToken))
	response, _ := ioutil.ReadAll(resp.Body)
	return c.JSONPretty(http.StatusOK, response, "    ")
}
