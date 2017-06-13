package services

import (
	"time"

	"github.com/FlorentinDUBOIS/bouncer/src/provider/api"
	"github.com/badoux/checkmail"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	jwt "gopkg.in/appleboy/gin-jwt.v2"
)

// JWTMiddleware to secure the API
var JWTMiddleware *jwt.GinJWTMiddleware

// InitJWT middleware
func InitJWT() {
	JWTMiddleware = &jwt.GinJWTMiddleware{
		Realm:      viper.GetString("jwt-realm"),
		Key:        []byte(viper.GetString("jwt-secret")),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
		Authenticator: func(pUserId string, pPassword string, pContext *gin.Context) (string, bool) {
			var user *api.User
			if err := checkmail.ValidateFormat(pUserId); err != nil {
				user = apiUserFactory.FromPostgres(userRepository.FindByUUID(pUserId))
			} else {
				user = apiUserFactory.FromPostgres(userRepository.FindByEmail(pUserId))
			}

			if user == nil {
				return pUserId, false
			}

			if !user.CheckPassword(pPassword) {
				return pUserId, false
			}

			return *user.ID, true
		},

		Authorizator: func(pUserId string, pContext *gin.Context) bool {
			return true
		},

		Unauthorized: func(pContext *gin.Context, pCode int, pMessage string) {
			pContext.JSON(pCode, gin.H{
				"code":    pCode,
				"message": pMessage,
			})
		},

		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		TokenLookup: "header:Authorization",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "JWT",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	}
}
