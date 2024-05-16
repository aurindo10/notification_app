package server

import (
	"database/sql"
	"net/http"
)

func NewServer(
	db *sql.DB,
) http.Handler {
	mux := http.NewServeMux()
	NewHandlers(mux).StartHandlers()
	var handler http.Handler = mux
	return handler
}
