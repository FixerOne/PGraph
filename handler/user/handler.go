package user

import (
	controller "pgraph/controller/user"
	repository "pgraph/repository/user"
	service "pgraph/service/user"
	"time"

	"net/http"
	"pgraph/entity"

	"github.com/dgrijalva/jwt-go"
	gin "github.com/gin-gonic/gin"

	headers "pgraph/handler/headers"
)

var (
	Repository repository.Repository        = repository.New()
	Service    service.Service              = service.New(Repository)
	Controller controller.CompanyController = controller.New(Service)
)

//Handler interface
type Handler interface {
	StartHandlers()
}

type handler struct {
	server *gin.Engine
}

//New constructor
func New(server *gin.Engine) Handler {
	return &handler{
		server: server,
	}
}

//StartHandlers function
func (h *handler) StartHandlers() {

	h.server.GET("/user/GetAll", func(c *gin.Context) {
		headers.SetUpHeaders(c)
		c.JSON(200, Controller.FindAll())
	})

	h.server.GET("/user/GetAllByType/:usertype", func(c *gin.Context) {
		headers.SetUpHeaders(c)
		usertype := c.Param("usertype")
		c.JSON(200, Controller.FindAllByType(usertype))
	})

	h.server.GET("/user/GetByCompany/:id", func(c *gin.Context) {
		headers.SetUpHeaders(c)
		id := c.Param("id")
		c.JSON(200, Controller.FindByCompanyID(id))
	})

	h.server.OPTIONS("/login", func(c *gin.Context) {
		headers.SetUpHeaders(c)
		c.Writer.WriteHeader(200)
	})

	h.server.POST("/login", func(c *gin.Context) {
		headers.SetUpHeaders(c)
		var data = Controller.Login(c)

		if data.ID == 0 {
			c.Writer.WriteHeader(http.StatusUnauthorized)
			return
		}

		expirationTime := time.Now().Add(5 * time.Minute)

		claims := &entity.Claims{
			Username: data.Mail,
			StandardClaims: jwt.StandardClaims{
				// In JWT, the expiry time is expressed as unix milliseconds
				ExpiresAt: expirationTime.Unix(),
			},
		}

		var jwtKey = []byte("my_secret_key")

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(jwtKey)

		if err != nil {
			// If there is an error in creating the JWT return an internal server error
			c.Writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		http.SetCookie(c.Writer, &http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: expirationTime,
		})

		data.Token = tokenString

		c.JSON(200, data)
		//c.String(200, tokenString)

	})

}
