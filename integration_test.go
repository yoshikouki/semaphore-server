package main

import (
	"bytes"
	"encoding/json"
	"github.com/yoshikouki/semaphore-server/api"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
)

const testURL = "http://localhost:8686"

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestServerConnection(t *testing.T) {
	res, err := http.Get(testURL + "/semaphore/ping")
	if err != nil {
		t.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err := res.Body.Close(); err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("Server returned wrong status code: got %d want %d", res.StatusCode, http.StatusOK)
	}

	expected := "pong"
	got := string(body)
	if got != expected {
		t.Errorf("Server returned wrong body: got %s want %s", got, expected)
	}
}

func TestRedisConnection(t *testing.T) {
	res, err := http.Get(testURL + "/semaphore/redis/ping")
	if err != nil {
		t.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err := res.Body.Close(); err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("Redis returned wrong status code: got %d want %d", res.StatusCode, http.StatusOK)
	}

	expected := "PONG"
	got := string(body)
	if got != expected {
		t.Errorf("Redis returned wrong body: got %s want %s", got, expected)
	}
}

func TestLock(t *testing.T) {
	body := lockRequest(t)

	expected := api.LockIfNotExistsResponse{
		GetLocked:  "true",
		User:       "test",
	}
	var got api.LockIfNotExistsResponse
	json.Unmarshal(body, &got)
	if got.GetLocked != expected.GetLocked || got.User != expected.User {
		t.Errorf("/lock returned wrong body: got %s want %s", got, expected)
	}
}

func lockRequest(t *testing.T) []byte {
	client := &http.Client{}
	data, _ := json.Marshal(map[string]string{
		"lock_target": "org-repo-stage",
		"user":        "test",
		"ttl":         "1s",
	})

	req, err := http.NewRequest("POST", testURL+"/semaphore/lock", bytes.NewBuffer(data))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("/lock returned wrong status code: got %d want %d", res.StatusCode, http.StatusOK)
	}

	return body
}
