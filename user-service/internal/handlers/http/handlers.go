package server

import (
	"net/http"

	_ "github.com/aurindo10/internal/handlers/http/docs"
	httpSwagger "github.com/swaggo/http-swagger"
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
}

// RegisterUser handles user registration
// @Summary      Register User
// @Description  Registers a new user
// @Tags         user
// @Produce      plain
// @Success      200  {string}  string "message"
// @Router       /registeruser [get]
func (c *Handlers) RegisterUser(w http.ResponseWriter, r *http.Request) {
	res := []byte("Olá")
	w.Write(res)
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
