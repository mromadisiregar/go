// Trying socket programming with go
// Little code but huge binary
// hemm... :|
// zamprox on Go
// mromadisiregar@gmail.com
package main 

import (
	"net"	
	"fmt"
	"io"
	"os"
)

// Main function
func main() {
	var remoteProxyAddr = os.Args[1]
	var remoteProxyPort = os.Args[2]
	// Port
	var localPort = ":8888"
	// Create sockServer
	sockServer, _ := net.Listen("tcp", localPort)
	fmt.Printf("Listen pada port %s\n", localPort)
	// Create daemon
	for {
		sockClient, _ := sockServer.Accept()
		sockProxy, _ := net.Dial("tcp", remoteProxyAddr + ":" + remoteProxyPort)
		go handleCon(sockClient, sockProxy)
		go handleCon(sockProxy, sockClient)
	}
	sockServer.Close()
}

func handleCon(c net.Conn, p net.Conn) {
	defer c.Close()
	defer p.Close()
	io.Copy(c, p)
}
