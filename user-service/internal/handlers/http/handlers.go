package server

import (
	"net/http"

	"github.com/aurindo10/internal/entities"
	_ "github.com/aurindo10/internal/handlers/http/docs"
	"github.com/aurindo10/internal/services"
	"github.com/aurindo10/pkg/utils"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/gorm"
)

// @title           Notification App API
// @version         1.0
// @description     This is a sample server for a notification service.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:3000
// @BasePath  /api/v1

// @schemes   http https
type Handlers struct {
	mux *http.ServeMux
	db  *gorm.DB
}

// RegisterUser handles user registration
// @Summary      Register User
// @Description  Registers a new user
// @Tags         user
// @Produce      plain
// @Success      200  {string}  string "message"
// @Router       /registeruser [get]
func (c *Handlers) RegisterUser(w http.ResponseWriter, r *http.Request) {
	decoded, _, err := utils.DecodeValid[entities.User](r)
	if err != nil {
		utils.Encode(w, r, 400, err.Error())
	}
	res, err := services.RegisterUserService(c.db, &decoded)
	if err != nil {
		utils.Encode(w, r, 400, err.Error())
	}
	if res != nil {
		utils.Encode(w, r, 201, res)
	}
}

// StartHandlers initializes the handlers
func (c *Handlers) StartHandlers() {
	c.mux.HandleFunc("/api/v1/registeruser", c.RegisterUser)
	c.mux.Handle("/swagger/", httpSwagger.WrapHandler)
}
func NewHandlers(mux *http.ServeMux) *Handlers {
	return &Handlers{
		mux: mux,
	}
}
