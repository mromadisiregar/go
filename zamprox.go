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

func usage(args []string) {
	fmt.Println("Usage:\n")
	fmt.Println(args[0] + " <RemoteProxyIP> <RemoteProxyPort>\n")
}

// Main function
func main() {
	if len(os.Args) < 3 {
		go usage(os.Args)
		os.Exit(3)
	}
	// Set arguments to variables
	var remoteProxyAddr = os.Args[1]
	var remoteProxyPort = os.Args[2]
	// Port
	var localPort = ":8888"
	// Create sockServer
	sockServer, errListener := net.Listen("tcp", localPort)
	if errListener != nil {
		fmt.Printf("Gagal listen : %+v\n", errListener)
		os.Exit(1)
	}
	fmt.Printf("Listen pada port %s\n", localPort)
	// Create daemon
	for {
		sockClient, errAsep := sockServer.Accept()
		if errAsep != nil {
			fmt.Println("Asep error, lanjut...")
			continue
		}
		// Print client address and port
		fmt.Printf("Menerima koneksi dari %+v\n", sockClient.RemoteAddr())
		sockProxy, _ := net.Dial("tcp", remoteProxyAddr + ":" + remoteProxyPort)
		// Dialing connection to remote proxy
		fmt.Println("Konek ke remote proxy ["+ remoteProxyAddr + ":" + remoteProxyPort +"]")
		// Let go handle everythings
		go handleCon(sockClient, sockProxy)
		fmt.Println("Selesai kopi data dari client ke remote proxy.")
		go handleCon(sockProxy, sockClient)
		fmt.Println("Selesai kopi data dari remote proxy ke client.")
	}
	sockServer.Close()
}

func handleCon(c net.Conn, p net.Conn) {
	// Just simple copy io on socket
	defer c.Close()
	defer p.Close()
	io.Copy(c, p)
}
