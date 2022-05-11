package middlewares

import (
	"net/http"
	"restapi-golang/helper"
	"restapi-golang/models/web"
)

type AuthMiddlewares struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddlewares {
	return &AuthMiddlewares{
		Handler: handler,
	}
}

func (middlewares *AuthMiddlewares) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if "RAHASIA" == request.Header.Get("X-API-Key") {
		middlewares.Handler.ServeHTTP(writer, request)
	} else {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "Unauthorized",
		}

		helper.WriteResponseBody(writer, webResponse)
	}
}
