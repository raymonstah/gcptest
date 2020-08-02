package p_test

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"net/http/httptrace"
	"testing"
	"time"

	p "github.com/raymonstah/gcptest"
	"github.com/tj/assert"
)

func TestHello(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/raymond", nil)
	w := httptest.NewRecorder()

	p.GetHello(w, r)

	assert.Equal(t, "Hello, raymond", w.Body.String())
}

func TestInvoke(t *testing.T) {
	baseURL := "https://us-west2-gcptest-285205.cloudfunctions.net/sayhello/"
	var start time.Time

	trace := &httptrace.ClientTrace{}

	totalDuration := int64(0)
	count := 100
	for i := 0; i < count; i++ {
		randomString := RandStringRunes(10)
		req, err := http.NewRequest("GET", baseURL+randomString, nil)
		assert.Nil(t, err)
		req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
		start = time.Now()
		if _, err := http.DefaultTransport.RoundTrip(req); err != nil {
			log.Fatal(err)
		}
		duration := time.Since(start)
		fmt.Printf("Total time: %v\n", duration)
		totalDuration += duration.Milliseconds()
	}

	averageDuration := totalDuration / int64(count)
	fmt.Println("average time in ms: ", averageDuration)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func TestInvokeFirestore(t *testing.T) {
	baseURL := "https://us-central1-gcptest-285205.cloudfunctions.net/GetHelloFirestore/tommy"
	var start time.Time

	trace := &httptrace.ClientTrace{}

	totalDuration := int64(0)
	count := 10
	for i := 0; i < count; i++ {
		req, err := http.NewRequest("GET", baseURL, nil)
		assert.Nil(t, err)
		req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
		start = time.Now()
		if _, err := http.DefaultTransport.RoundTrip(req); err != nil {
			log.Fatal(err)
		}
		duration := time.Since(start)
		fmt.Printf("Total time: %v\n", duration)
		totalDuration += duration.Milliseconds()
	}

	averageDuration := totalDuration / int64(count)
	fmt.Println("average time in ms: ", averageDuration)
}
