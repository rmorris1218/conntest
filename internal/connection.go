package internal

import (
	"fmt"
	"net"
	"time"
)

type ConnectionParams struct {
	Port     int `json:"port"`
	Protocol string `json:"protocol"`
	isOpen   bool
	message  string
}

type Endpoint struct {
	Uri              string `json:"endpoint"`
	ConnectionParams []ConnectionParams `json:"proto_ports"`
	timeoutSeconds   int
}

func (e Endpoint) TestReachability() bool {
	failures_present := false
	for _, param := range e.ConnectionParams {
		isOpen, msg := tryConnectivity(e.Uri, param.Port, param.Protocol)
		param.isOpen = isOpen
		param.message = msg
		if !isOpen {
			failures_present = true
		}
	}
	return !failures_present
}

func tryConnectivity(endpoint string, port int, protocol string) (bool, string) {
	// Need to iron out how to reliably test UDP "connections"
	if protocol == "" {
		protocol = "tcp"
	}
	hostPort := fmt.Sprintf("%s:%d", endpoint, port)
	conn, err := net.DialTimeout(protocol, hostPort, 2*time.Second)
	if err != nil {
		return false, "could not connect"
	}
	if conn != nil {
		defer conn.Close()
		return true, "connection successful"
	}
	return false, "port not open"
}
