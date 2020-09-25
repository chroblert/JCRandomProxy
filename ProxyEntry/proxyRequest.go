package ProxyEntry

import (
	"net"
	// "fmt"
)

// Forward all requests from r to w
func proxyRequest(r net.Conn, w net.Conn) {
	defer r.Close()
	defer w.Close()

	var buffer = make([]byte, 4096000)
	for {
		n, err := r.Read(buffer)
		if err != nil {
			// fmt.Printf("Unable to read from input, error: %s\n", err.Error())
			break
		}
		n, err = w.Write(buffer[:n])
		if err != nil {
			// fmt.Printf("Unable to write to output, error: %s\n", err.Error())
			break
		}
	}
}
