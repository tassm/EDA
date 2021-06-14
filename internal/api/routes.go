package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/TasSM/EDA/internal/models"
	"github.com/gorilla/schema"
	"github.com/nats-io/nats.go"
)

var decoder = schema.NewDecoder()

type NATSEventRequest struct {
	Count   uint16
	DelayMs uint16
}

func NatsPubSubHandler(nc *nats.Conn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var natsRequest NATSEventRequest
		err := decoder.Decode(&natsRequest, r.URL.Query())
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("Bad Request: %s", err.Error())))
			return
		}
		go models.GenerateNATSEventsPubSub(nc, natsRequest.Count, natsRequest.DelayMs)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(natsRequest)
	}
}
