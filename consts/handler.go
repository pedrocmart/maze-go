package consts

import "net/http"

const (
	HandlerStatusCodeOK                  = int64(http.StatusOK)
	HandlerStatusCreated                 = int64(http.StatusCreated)
	HandlerStatusNotFound                = int64(http.StatusNotFound)
	HandlerStatusCodeBadRequest          = int64(http.StatusBadRequest)
	HandlerStatusCodeInternalServerError = int64(http.StatusInternalServerError)

	HandlerMessageSuccess = "Success"
	HandlerMessageError   = "Error"

	HandlerSuccess = true
	HandlerFailed  = false
)
