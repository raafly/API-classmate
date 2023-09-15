package exception

import (
	"net/http"

	"github.com/raafly/api-classmate/helper"
	"github.com/raafly/api-classmate/model"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, err interface{}) {
	
	if notFoundError(w, r, err) {
		return
	}

	internalServerError(w, r, err)
}

func notFoundError(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError) 
	if ok {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)

		webResponse := model.WebResponse{
			Code: http.StatusNotFound,
			Status: "Status Not Found",
			Data: exception.Error,
		}

		helper.WriteToResponseBody(w, webResponse)

		return true
	} else {
		return false
	}
}

func internalServerError(w http.ResponseWriter, r *http.Request, err interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	webResponse := model.WebResponse{
		Code: http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Data: err,
	}

	helper.WriteToResponseBody(w, webResponse)
}
