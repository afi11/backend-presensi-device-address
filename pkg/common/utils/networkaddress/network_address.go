package networkaddress

import (
	"fmt"
	"net"
	"net/http"
	"strings"
)

func GetClientIPByHeaders(req *http.Request) (ip string) {

	ip = req.Header.Get("X-Forwarded-For")
	if ip == "" {
		ip = strings.Split(req.RemoteAddr, ":")[0]
	}

	ifas, err := net.Interfaces()
	if err != nil {
		return ip
	}

	for _, ifa := range ifas {
		a := ifa.HardwareAddr.String()
		if a != "" {
			fmt.Println("Mac Address : ", a)
			break
		}
	}

	return ip

}

func GetUserAgent(req *http.Request) (user_agent string) {
	user_agent = req.UserAgent()
	return user_agent
}
