package server

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/phoihos/gosim/route"
)

type shutdownHandler struct {
	cancel context.CancelFunc
}

func (sh *shutdownHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := fmt.Sprintf("Sent shutdown message to Server(pid:%d).", os.Getpid())
	io.WriteString(w, response)

	sh.cancel()
}

// Configuration is used  to configure server
type Configuration struct {
	Port         string
	ShutdownPath string
}

// Run server application
func Run(conf *Configuration) {
	ctxSignal, cancelSignal := context.WithCancel(context.Background())
	defer cancelSignal()

	mux := route.BuildServeMux()

	if len(conf.ShutdownPath) > 1 ||
		(len(conf.ShutdownPath) == 1 && conf.ShutdownPath[0] != '/') {
		mux.Handle(conf.ShutdownPath, &shutdownHandler{cancel: cancelSignal})
	}

	server := &http.Server{Addr: ":" + conf.Port, Handler: mux}

	go func() {
		log.Printf("Server listening on port %s.", conf.Port)

		if err := server.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				log.Fatal(err)
			}
		}
	}()

	chanSingal := make(chan os.Signal, 1)
	signal.Notify(chanSingal, syscall.SIGKILL, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)

	select {
	case <-chanSingal:
		log.Println("Shutdown requested from SIGNAL.")
	case <-ctxSignal.Done():
		log.Println("Shutdown requested from HTTP.")
	}

	ctxTimeout, cancelTimeout := context.WithTimeout(context.Background(), time.Second*30)
	defer cancelTimeout()

	if err := server.Shutdown(ctxTimeout); err != nil {
		log.Fatal(err)
	}

	log.Println("Server has shutdown gracefully.")
}
