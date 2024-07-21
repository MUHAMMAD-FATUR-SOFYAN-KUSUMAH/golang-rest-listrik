package middelware

import (
	"golang_listrik/helper"
	"golang_listrik/model/response"
	"net"
	"net/http"
	"time"
)

var client = make(map[string]int)

var ClientTimeout = make(map[string]time.Time)

type authmiddelware struct {
	handler http.Handler
}

func addclient(ip string) {
	client[ip] = 0
}

func addcounter(ip string) {
	client[ip]++
}

func Newmiddelware(handle http.Handler) *authmiddelware {
	return &authmiddelware{
		handler: handle,
	}
}

func (middel *authmiddelware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	ip, _, _ := net.SplitHostPort(request.RemoteAddr)
	_, counter := client[ip]
	if !counter {
		addclient(ip)
	}

	if "RAHASIA" == request.Header.Get("x-api-key") {
		if client[ip] >= 50 {

			_, counter := ClientTimeout[ip]

			if !counter {
				startTime := time.Now()
				targetTime := startTime.Add(1 * time.Minute)
				ClientTimeout[ip] = targetTime
			}

			if time.Now().After(ClientTimeout[ip]) {
				client[ip] = 0
				delete(ClientTimeout, ip)
				middel.handler.ServeHTTP(writer, request)
				return
			}

			writer.Header().Set("Content-Type", "application/json")
			writer.WriteHeader(http.StatusTooManyRequests)

			webResponse := response.WebResponse{
				Code:   http.StatusTooManyRequests,
				Status: "limit request",
				Data:   "too many request wait 5 minute",
			}

			helper.Encode_Json(writer, webResponse)
			return
		}
		middel.handler.ServeHTTP(writer, request)
		addcounter(ip)
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
