package index

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/money_kanri/app"
	pubsub "google.golang.org/api/pubsub/v1"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

// Req is parsed data of http request.
type Req struct {
	pubsub.PubsubMessage `json:"message"`
}

func init() {
	http.HandleFunc("/pubsub", pubsubHandler)
}

func pubsubHandler(w http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(w, "failed to read body request", http.StatusBadRequest)
		log.Errorf(ctx, "failed to read body: %v", err)
	}
	log.Debugf(ctx, "request: %v", string(body))

	request := Req{}
	if err := json.Unmarshal(body, &request); err != nil {
		http.Error(w, "failed to decode json", http.StatusBadRequest)
		log.Errorf(ctx, "failed to decode json: %v", err)
		return
	}
	log.Debugf(ctx, "pub/sub message request: %v", request.PubsubMessage.Attributes)
	event := request.PubsubMessage.Attributes["eventType"]
	if event != "OBJECT_FINALIZE" {
		return
	}
	filename := request.PubsubMessage.Attributes["objectId"]
	bucketID := request.PubsubMessage.Attributes["bucketId"]
	if err := app.GetFile(ctx, filename, bucketID); err != nil {
		log.Errorf(ctx, "failed to get file: %v", err)
		return
	}
	return
}
