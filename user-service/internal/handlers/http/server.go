package server

import (
	"net/http"

	"gorm.io/gorm"
)

func NewServer(
	db *gorm.DB,
) http.Handler {
	mux := http.NewServeMux()
	NewHandlers(mux).StartHandlers()
	var handler http.Handler = mux
	return handler
}
