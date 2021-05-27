package main

import (
	"bytes"
	"encoding/json"
	"github.com/yoshikouki/semapi/api"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
)

const testURL = "http://localhost:8686/semaphore"

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestServerConnection(t *testing.T) {
	res, err := http.Get(testURL + "/ping")
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
	res, err := http.Get(testURL + "/redis/ping")
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
	body := lockRequest(t, &api.LockIfNotExistsParams{
		LockTarget: "org-repo-stage",
		User:       "test",
		TTL:        "1s",
	})

	expected := api.LockIfNotExistsResponse{
		GetLocked: "true",
		User:      "test",
	}
	var got api.LockIfNotExistsResponse
	json.Unmarshal(body, &got)
	if got.GetLocked != expected.GetLocked || got.User != expected.User {
		t.Errorf("/lock returned wrong body: got %s want %s", got, expected)
	}
}

func TestLockAndLock(t *testing.T) {
	lockRequest(t, &api.LockIfNotExistsParams{
		LockTarget: "org-repo-stage",
		User:       "test",
		TTL:        "1s",
	})
	body := lockRequest(t, &api.LockIfNotExistsParams{
		LockTarget: "org-repo-stage",
		User:       "test",
		TTL:        "1s",
	})

	expected := api.LockIfNotExistsResponse{
		GetLocked: "true",
		User:      "test",
	}
	var got api.LockIfNotExistsResponse
	json.Unmarshal(body, &got)
	if got.GetLocked != expected.GetLocked || got.User != expected.User {
		t.Errorf("/lock returned wrong body: got %s want %s", got, expected)
	}
}

func TestLockAndInvalidLock(t *testing.T) {
	lockRequest(t, &api.LockIfNotExistsParams{
		LockTarget: "org-repo-stage",
		User:       "test",
		TTL:        "1s",
	})
	body := lockRequest(t, &api.LockIfNotExistsParams{
		LockTarget: "org-repo-stage",
		User:       "InvalidUser",
		TTL:        "1s",
	})

	expected := api.LockIfNotExistsResponse{
		GetLocked: "false",
		User:      "test",
	}
	var got api.LockIfNotExistsResponse
	json.Unmarshal(body, &got)
	if got.GetLocked != expected.GetLocked || got.User != expected.User {
		t.Errorf("/lock returned wrong body: got %s want %s", got, expected)
	}
}

func TestLockAndUnlock(t *testing.T) {
	lockRequest(t, &api.LockIfNotExistsParams{
		LockTarget: "org-repo-stage",
		User:       "test",
		TTL:        "1s",
	})
	unlockBody := unlockRequest(t, &api.UnlockParams{
		UnlockTarget: "org-repo-stage",
		User:         "test",
	})

	expected := api.UnlockResponse{
		GetUnlock: "true",
		Message:   "",
	}
	var got api.UnlockResponse
	json.Unmarshal(unlockBody, &got)
	if got.GetUnlock != expected.GetUnlock || got.Message != expected.Message {
		t.Errorf("/lock returned wrong body: got %s want %s", got, expected)
	}
}

func TestInvalidUnlock(t *testing.T) {
	unlockBody := unlockRequest(t, &api.UnlockParams{
		UnlockTarget: "org-repo-stage",
		User:         "test",
	})

	expected := api.UnlockResponse{
		GetUnlock: "false",
		Message:   "org-repo-stage haven't locked",
	}
	var got api.UnlockResponse
	json.Unmarshal(unlockBody, &got)
	if got.GetUnlock != expected.GetUnlock || got.Message != expected.Message {
		t.Errorf("/lock returned wrong body: got %s want %s", got, expected)
	}
}

func TestLockAndInvalidUnlock(t *testing.T) {
	lockRequest(t, &api.LockIfNotExistsParams{
		LockTarget: "org-repo-stage",
		User:       "test",
		TTL:        "1s",
	})
	unlockBody := unlockRequest(t, &api.UnlockParams{
		UnlockTarget: "org-repo-stage",
		User:         "InvalidUser",
	})

	expected := api.UnlockResponse{
		GetUnlock: "false",
		Message:   "org-repo-stage don't release lock, because lock owner isn't InvalidUser",
	}
	var got api.UnlockResponse
	json.Unmarshal(unlockBody, &got)
	if got.GetUnlock != expected.GetUnlock || got.Message != expected.Message {
		t.Errorf("/lock returned wrong body: got %s want %s", got, expected)
	}
}

func lockRequest(t *testing.T, params *api.LockIfNotExistsParams) []byte {
	client := &http.Client{}
	data, _ := json.Marshal(params)

	req, err := http.NewRequest("POST", testURL+"/lock", bytes.NewBuffer(data))
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

func unlockRequest(t *testing.T, params *api.UnlockParams) []byte {
	client := &http.Client{}
	data, _ := json.Marshal(params)

	req, err := http.NewRequest("POST", testURL+"/unlock", bytes.NewBuffer(data))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("/unlock returned wrong status code: got %d want %d", res.StatusCode, http.StatusOK)
	}

	return body
}
