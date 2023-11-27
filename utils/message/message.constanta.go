package message

import "net/http"

var (
	StatusBadRequestCode = http.StatusBadRequest
	StatusInternalServer = http.StatusInternalServerError
	StatusOk             = http.StatusOK
	StatusCreated        = http.StatusCreated
	StatusUnauthorize    = http.StatusUnauthorized
	StatusNotFound       = http.StatusNotFound
	StatusForbiden       = http.StatusForbidden
)

const (
	SUCCESS = "SUCCESS"
	FAILED  = "FAILED"

	SuccessDisplay = "Success"
	FailedDisplay  = "Failed"
)
