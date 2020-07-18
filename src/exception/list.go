package exception

import "net/http"

var (
	BadRequest = &Exception{http.StatusBadRequest, "BAD_REQUEST", "bad request"}
	Internal   = &Exception{http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", "internal server error"}
	// Add any new error instances
)
