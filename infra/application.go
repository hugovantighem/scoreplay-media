package infra

import (
	"context"
	"errors"
	"log"
	"myproject/api"
	"myproject/app"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ShutDownFuncs []func()

func (x ShutDownFuncs) Shutdown() {
	for i := len(x) - 1; i >= 0; i-- {
		x[i]()
	}
}

func RunApplication(conf Config) func() {

	var shutdown ShutDownFuncs = []func(){}

	mgoClient, close := InitDB(conf)
	shutdown = append(shutdown, close)

	r := gin.Default()

	filestorageDir := filepath.Join(".", "assets")
	r.Static("/assets", filestorageDir)
	err := os.MkdirAll(filestorageDir, os.ModePerm)
	if err != nil {
		log.Fatalf("cannot create asset directory: %s", err)
	}

	var fileStorage app.FileStorage = NewFileStorage(filestorageDir)
	var tagStore app.TagStore = NewTagStore(mgoClient)
	var mediaStore app.MediaStore = NewMediaStore(mgoClient)

	server := NewServer(mgoClient, fileStorage, tagStore, mediaStore)

	handler := api.NewStrictHandler(server, nil)
	api.RegisterHandlers(r, handler)

	// And we serve HTTP until the world ends.

	s := &http.Server{
		Handler: r,
		Addr:    conf.ServerAddr,
	}

	// And we serve HTTP until the world ends.
	go func() {
		if err := s.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logrus.Errorf("error: %s\n", err)
		}
	}()

	return func() {

		logrus.Println("Shutting down server...")
		shutdown.Shutdown()

		// The context is used to inform the server it has 5 seconds to finish
		// the request it is currently handling
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := s.Shutdown(ctx); err != nil {
			logrus.Fatalf("Error while shutting down Server: %s. Initiating force shutdown...", err.Error())
		} else {
			logrus.Info("Server exiting")
		}
	}
}
