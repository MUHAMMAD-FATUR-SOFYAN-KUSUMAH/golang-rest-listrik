package middelware

import (
	"fmt"
	"golang_listrik/helper"
	"golang_listrik/model/response"
	"net"
	"net/http"
)

type authmiddelware struct {
	handler http.Handler
}

func Newmiddelware(handle http.Handler) *authmiddelware {
	return &authmiddelware{
		handler: handle,
	}
}

func (middel *authmiddelware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	ip, _, _ := net.SplitHostPort(request.RemoteAddr)
	fmt.Println(ip)
	if "RAHASIA" == request.Header.Get("x-api-key") {
		// ok
		middel.handler.ServeHTTP(writer, request)
	} else {
		// error
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		webResponse := response.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}

		helper.Encode_Json(writer, webResponse)
	}
}
