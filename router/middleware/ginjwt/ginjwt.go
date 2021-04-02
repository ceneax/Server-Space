package ginjwt

import (
	"log"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"

	"xcdh.space/space/config"
	"xcdh.space/space/entity"
	"xcdh.space/space/service"
	"xcdh.space/space/utils/msg"
)

func GinJWTMiddlewareInit() (authMiddleware *jwt.GinJWTMiddleware) {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "gin jwt",
		Key:         []byte("~cW@N11Fv#8E<Yfi"),
		Timeout:     time.Hour * 24,
		MaxRefresh:  time.Hour * 720,
		IdentityKey: config.IdentityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*entity.Login); ok {
				//get claims from username
				// v.UserId = models.GetUserClaims(v.UserName)
				// jsonClaim, _ := json.Marshal(v.UserClaims)
				//maps the claims in the JWT
				return jwt.MapClaims{
					"UserId": v.UserId,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			//Set the identity
			return claims[config.IdentityKey]
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			//handles the login logic. On success LoginResponse is called, on failure Unauthorized is called
			var loginVals entity.Login
			if err := c.ShouldBind(&loginVals); err != nil {
				if err1 := c.ShouldBindJSON(&loginVals); err1 != nil {
					return "", jwt.ErrMissingLoginValues
				}
			}
			var user service.User
			user.UserId = loginVals.UserId
			user.Password = loginVals.Password

			b := user.CheckLogin()
			if b {
				return &entity.Login{
					UserId: user.UserId,
				}, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		//receives identity and handles authorization logic
		Authorizator: Authorizator,
		//handles unauthorized logic
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		RefreshResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			c.JSON(code, gin.H{
				"code":   code,
				"token":  token,
				"expire": expire.Format(time.RFC3339),
			})
		},
		// User can define own LoginResponse func.
		LoginResponse: func(c *gin.Context, code int, token string, expire time.Time) {

			c.JSON(code, gin.H{
				"code":   code,
				"token":  token,
				"expire": expire.Format(time.RFC3339),
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time valumsg. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}
	return
}

func Authorizator(data interface{}, c *gin.Context) bool {
	if data.(string) == "test" {
		return true
	}
	return true
}

//404 handler
func NoRouteHandler(c *gin.Context) {
	code := msg.PAGE_NOT_FOUND
	c.JSON(404, gin.H{"code": code, "message": msg.GetMsg(code)})
}

// LoginHandler godoc
// @Summary 登陆
// @Description 登陆
// @ID login-main-json
// @Accept  application/x-www-form-urlencoded
// @Accept  application/json
// @Produce application/json
// @Param  User body  entity.Login  true  "用户登陆信息"
// @Success 200 {string} utils.Message
// @Failure 400 {object} utils.HTTPError
// @Failure 404 {object} utils.HTTPError
// @Failure 500 {object} utils.HTTPError
// @Router /v1/login [post]
func LoginHandler(authMiddleware *jwt.GinJWTMiddleware) gin.HandlerFunc {
	return authMiddleware.LoginHandler
}

// RefreshHandler godoc
// @Summary 刷新token
// @Description 刷新token
// @ID login-refresh-toekn-json
// @Produce application/json
// @Security ApiKeyAuth
// @Success 200 {string} utils.Message
// @Failure 400 {object} utils.HTTPError
// @Failure 404 {object} utils.HTTPError
// @Failure 500 {object} utils.HTTPError
// @Router /v1/refresh_toekn [get]
func RefreshHandler(authMiddleware *jwt.GinJWTMiddleware) gin.HandlerFunc {
	return authMiddleware.RefreshHandler
}
