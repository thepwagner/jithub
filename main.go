package main

import (
	"errors"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/thepwagner/jithub/api"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	h := api.NewHandler()

	srv := &http.Server{
		Addr:    "0.0.0.0:9666",
		Handler: h,
	}

	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		logrus.WithError(err).Fatal("server errorz")
	}
}
