package service

import (
	"github.com/farkramdev/auth-service/src/api"
	"github.com/labstack/echo"
	"net/http"
)

type authRequest struct {
	GrantType    grantType `json:"grant_type"`
	Username     string    `json:"username"`
	Password     string    `json:"password"`
	RefreshToken string    `json:"refresh_token"`
}

type authResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int64  `json:"expires_in"` // unit: seconds
	RefreshToken string `json:"refresh_token,omitempty"`
	UID          int64  `json:"uid"`
}

type grantType string

// Grant Type
const (
	grantTypePassword     = "password"
	grantTypeRefreshToken = "refresh_token"
)

// Auth Auth
func Auth(g *echo.Group) {
	g.POST("", authTokenHandler)
}

func authTokenHandler(c echo.Context) error {
	var body authRequest

	if err := c.Bind(&body); err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}

	if body.GrantType == grantTypePassword {

		// handle password grant type => return refresh token

		user, err := api.FindUser(body.Username, body.Password)

		// if err != nil {
		// 	log.Println(err)
		// 	return c.String(http.StatusInternalServerError, "Internal Server Error")
		// }

		// if user == nil {
		// 	// user or password wrong = unauthorized
		// 	return c.String(http.StatusUnauthorized, "Unauthorized")
		// }

		// refreshToken, err := generateRefreshToken(user.ID)

		// if err != nil {
		// 	log.Println(err)
		// 	return c.String(http.StatusInternalServerError, "Internal Server Error")
		// }

		// accessToken, err := generateAccessToken(user.ID, accessTokenDuration)

		// if err != nil {
		// 	log.Println(err)
		// 	return c.String(http.StatusInternalServerError, "Internal Server Error")
		// }

		// return c.JSON(http.StatusOK, authResponse{
		// 	accessToken,
		// 	"bearer",
		// 	int64(accessTokenDuration.Seconds()),
		// 	refreshToken,
		// 	user.ID,
		// })

	}

	// if body.GrantType == grantTypeRefreshToken {
	// 	// handle refresh token grant type => return access token

	// 	// get user id from context
	// 	claims, err := validateToken(body.RefreshToken)
	// 	if err != nil {
	// 		return c.String(http.StatusUnauthorized, "Unauthorized")
	// 	}
	// 	// verify refresh token in database
	// 	if ok, err := api.ValidateToken(body.RefreshToken, claims.ID, refreshTokenDuration); !ok {
	// 		if err != nil {
	// 			log.Println(err)
	// 			return c.String(http.StatusInternalServerError, "Internal Server Error")
	// 		}
	// 		return c.String(http.StatusUnauthorized, "Unauthorized")
	// 	}

	// 	accessToken, err := generateAccessToken(claims.ID, accessTokenDuration)
	// 	if err != nil {
	// 		log.Println(err)
	// 		return c.String(http.StatusInternalServerError, "Internal Server Error")
	// 	}
	// 	return c.JSON(http.StatusOK, authResponse{
	// 		accessToken,
	// 		"bearer",
	// 		int64(accessTokenDuration.Seconds()),
	// 		"",
	// 		claims.ID,
	// 	})
	// }

	return c.String(http.StatusUnauthorized, "Unauthorized")
}
