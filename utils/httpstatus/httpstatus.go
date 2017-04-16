package httpstatus

import (
	"net/http"
)

// status code description
const (
	METHOD_NOT_ALLOWED    = "method not allowed"
	BAD_REQUEST           = "400 bad request"
	INTERNAL_SERVER_ERROR = "internal server error"
	FORBIDDEN             = "403 forbidden"
	NOT_FOUND             = "404 page not found"
)

// StatusMethodNotAllowed return code 405
func StatusMethodNotAllowed(w http.ResponseWriter) {
	http.Error(w, METHOD_NOT_ALLOWED, http.StatusMethodNotAllowed)
}

// StatusInternalServerError return code 500
func StatusInternalServerError(w http.ResponseWriter) {
	http.Error(w, INTERNAL_SERVER_ERROR, http.StatusInternalServerError)
}

// StatusBadRequest return code 400
func StatusBadRequest(w http.ResponseWriter) {
	http.Error(w, BAD_REQUEST, http.StatusBadRequest)
}

// StatusForbidden return code 403
func StatusForbidden(w http.ResponseWriter) {
	http.Error(w, FORBIDDEN, http.StatusForbidden)
}

// StatusNotFound return code 404
func StatusNotFound(w http.ResponseWriter) {
	http.Error(w, NOT_FOUND, http.StatusNotFound)
}
