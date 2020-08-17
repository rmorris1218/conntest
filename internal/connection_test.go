package internal

import "testing"

func TestHasEndpointConnectivity(t *testing.T) {
	cases := []struct {
		endpoint         string
		port             int
		protocol         string
		expectedResult   bool
		expeectedMessage string
	}{
		{"1.1.1.1", 53, "udp", true, "connection successful"},
		{"1.1.1.1", 533, "tcp", false, "port not open"},
		{"google.com", 443, "tcp", true, "connection successful"},
		{"google.com", 1433, "tcp", false, "port not open"},
		{"0.0.0.0", 0, "tcp", false, "count not connect"},
		{"ewttgsetgwegtw", 443, "tcp", false, "could not connect"},
		// {"google.com", 443, "udp", false, "port not open"},
	}

	for _, info := range cases {
		isOpen, _ := tryConnectivity(info.endpoint, info.port, info.protocol)
		if isOpen != info.expectedResult {
			t.Errorf("error: expected %t got %t for %s/%d to %s", info.expectedResult, isOpen, info.protocol, info.port, info.endpoint)
		}
	}
}

// func TestTestReachability(t *testing.T) {
// 	cases := []Endpoint{
// 		{uri: "1.1.1.1",
// 			connectionParams: [
// 				ConnectionParams{53, "tcp"},
// 				ConnectionParams{53, "udp"},
// 				ConnectionParams{533, "tcp"}
// 			]
// 		},
// 	}

// 	for _, info := range cases {
// 		isOpen, _ := tryConnectivity(info.endpoint, info.port, info.protocol)
// 		if isOpen != info.expectedResult {
// 			t.Errorf("error: expected %t got %t for %s:%d", info.expectedResult, isOpen, info.endpoint, info.port)
// 		}
// 	}
// }
