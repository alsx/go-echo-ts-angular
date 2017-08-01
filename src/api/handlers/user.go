package handlers

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/alsx/enli-task/src/api/models"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

// UserClaims contains user attributes and is extended with custom claims.
type UserClaims struct {
	Email      string
	FacebookID string
	jwt.StandardClaims
}

// Handler base for controller
type Handler struct{}

// UserHandler contains methods to work with user entity
var UserHandler = Handler{}

// Info get user info
func (h *Handler) Info(c echo.Context) error {
	cxtUser := c.Get("user").(*jwt.Token)
	claims := cxtUser.Claims.(*UserClaims)
	email := claims.Email
	facebookID := claims.FacebookID
	userMgr := models.NewUserManager(c)
	user, err := userMgr.GetByEmailOrFacebookID(email, facebookID)
	if err != nil {
		return c.JSONPretty(http.StatusOK, echo.Map{"error": err}, "    ")
	}
	return c.JSONPretty(http.StatusOK, user, "    ")
}

// SignIn authenticate user in api
func (h *Handler) SignIn(c echo.Context) error {
	data := struct{ Email, Password string }{}
	err := json.NewDecoder(c.Request().Body).Decode(&data)
	passwordSha256 := fmt.Sprintf("%x", sha256.Sum256([]byte(data.Password)))
	userMgr := models.NewUserManager(c)
	user, err := userMgr.GetByEmailOrFacebookID(data.Email, "")
	if err != nil {
		return c.JSONPretty(http.StatusInternalServerError, echo.Map{"error": err}, "    ")
	}
	if passwordSha256 == user.Password {
		secret := []byte(c.Get("secret").(string))
		t := h.jwtToken(user.Email, user.FacebookID, secret)
		return c.JSONPretty(http.StatusOK, echo.Map{"token": t}, "    ")
	}
	return echo.ErrUnauthorized
}

func (h *Handler) jwtToken(email, fbID string, secret []byte) string {
	// Set custom claims
	claims := &UserClaims{
		email,
		fbID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}
	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, _ := token.SignedString(secret)
	return t
}

// SignUp register user using email and password.
func (h *Handler) SignUp(c echo.Context) error {
	data := struct{ Name, Email, Password string }{}
	err := json.NewDecoder(c.Request().Body).Decode(&data)
	if data.Email == "" {
		return c.JSONPretty(http.StatusBadRequest, echo.Map{"error": "Invalid email."}, "    ")
	}
	// sha256 of password
	passwordSha256 := fmt.Sprintf("%x", sha256.Sum256([]byte(data.Password)))
	user := models.User{
		Email:    data.Email,
		Name:     data.Name,
		Password: passwordSha256,
	}

	userMgr := models.NewUserManager(c)
	_, err = userMgr.Add(user)
	if err != nil {
		return c.JSONPretty(http.StatusInternalServerError, echo.Map{"error": err}, "    ")
	}
	secret := []byte(c.Get("secret").(string))
	t := h.jwtToken(user.Email, user.FacebookID, secret)
	return c.JSONPretty(http.StatusOK, echo.Map{"token": t}, "    ")
}

// FacebookSignUp login user with facebook account
func (h *Handler) FacebookSignUp(c echo.Context) error {
	data := struct{ FacebookToken string }{}
	err := json.NewDecoder(c.Request().Body).Decode(&data)
	if err != nil {
		return c.JSONPretty(http.StatusInternalServerError, echo.Map{"error": err}, "    ")
	}

	resp, _ := http.Get("https://graph.facebook.com/me?access_token=" + url.QueryEscape(data.FacebookToken))
	userData := struct{ ID, Name string }{}
	err = json.NewDecoder(resp.Body).Decode(&userData)
	if err != nil {
		return c.JSONPretty(http.StatusInternalServerError, echo.Map{"error": err}, "    ")
	}
	c.Logger().Debugf("userData: %v", userData)
	user := models.User{
		Name:          userData.Name,
		FacebookID:    userData.ID,
		FacebookToken: data.FacebookToken,
	}

	userMgr := models.NewUserManager(c)

	if _, err := userMgr.GetByEmailOrFacebookID("", userData.ID); err == nil {
		_ = userMgr.Update(user)
	} else if _, err = userMgr.Add(user); err != nil {
		return c.JSONPretty(http.StatusInternalServerError, echo.Map{"error": err}, "    ")
	}

	secret := []byte(c.Get("secret").(string))
	t := h.jwtToken("", userData.ID, secret)
	return c.JSONPretty(http.StatusOK, echo.Map{"token": t}, "    ")
}
