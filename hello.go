package p

import (
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/api/option"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"strings"

	"cloud.google.com/go/firestore"
)

var namesCollection *firestore.CollectionRef

func init() {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, "gcptest-285205", option.WithCredentialsFile("gcptest-285205-62e9aabf356b.json"))
	if err != nil {
		panic(fmt.Errorf("error creating new client: %w", err))
	}

	namesCollection = client.Collection("names")

}

func GetHello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.RequestURI()
	name = strings.Trim(name, "/")

	helloStr := fmt.Sprintf("Hello, %v", name)
	_, err := w.Write([]byte(helloStr))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(fmt.Errorf("error writing to writer: %w", err).Error())
	}

}

func GetHelloFirestore(w http.ResponseWriter, r *http.Request) {
	name := r.URL.RequestURI()
	name = strings.Trim(name, "/")

	doc := namesCollection.Doc(name)
	docSnapshot, err := doc.Get(r.Context())
	if err != nil {
		if status.Code(err) == codes.NotFound {
			w.WriteHeader(http.StatusNotFound)
			return
		}
	}

	rawData, err := json.Marshal(docSnapshot.Data())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(fmt.Errorf("error writing to writer: %w", err).Error())
	}

	_, err = w.Write(rawData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(fmt.Errorf("error writing to writer: %w", err).Error())
	}

}
