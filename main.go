package main

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetOutput(os.Stdout)

	dt := time.Now()
	log.Info("Start time: ", dt.Format("01-02-2006 15:04:05"))

	waitTime := float64(10)

	if len(os.Args) > 1 {
		waitTime, _ = strconv.ParseFloat(os.Args[1], 64)
	}

	log.Info("Wait Time: ", waitTime)

	counter := 0

	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		counter++
		newDt := time.Now()

		minutesElapsed := newDt.Sub(dt).Minutes()

		responseCode := http.StatusBadGateway

		if minutesElapsed > waitTime {
			responseCode = http.StatusOK
		}
		w.WriteHeader(responseCode)

		log.WithFields(
			log.Fields{
				"counter":        counter,
				"minutesElapsed": minutesElapsed,
				"responseCode":   responseCode,
			},
		).Info("request received")
	})

	http.ListenAndServe(":80", router)
}
