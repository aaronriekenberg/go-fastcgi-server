package server

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/fcgi"
	"os"
	"strconv"
	"syscall"

	"github.com/aaronriekenberg/go-fastcgi/config"
)

func createListener(
	serverConfiguration *config.ServerConfiguration,
) (net.Listener, error) {

	umask, err := strconv.ParseInt(serverConfiguration.UmaskOctal, 8, 0)
	if err != nil {
		return nil, fmt.Errorf("strconv.ParseInt err = %w", err)
	}
	umaskInt := int(umask)
	log.Printf("umaskInt = %03O", umaskInt)

	os.Remove(serverConfiguration.UnixSocketPath)

	// needed so group www has rwx permission on the socket.
	previousUmask := syscall.Umask(umaskInt)
	log.Printf("previousUmask = %03O", previousUmask)
	defer syscall.Umask(previousUmask)

	listener, err := net.Listen("unix", serverConfiguration.UnixSocketPath)
	if err != nil {
		return nil, fmt.Errorf("net.Listen err = %w", err)
	}

	return listener, nil
}

func RunServer(
	serverConfiguration *config.ServerConfiguration,
	serveHandler http.Handler,
) error {

	log.Printf("begin RunServer UnixSocketPath = %q UmaskOctal = %q",
		serverConfiguration.UnixSocketPath,
		serverConfiguration.UmaskOctal,
	)

	listener, err := createListener(serverConfiguration)
	if err != nil {
		return fmt.Errorf("createListener err = %w", err)
	}

	log.Printf("before fcgi.Serve")
	return fcgi.Serve(listener, serveHandler)
}
