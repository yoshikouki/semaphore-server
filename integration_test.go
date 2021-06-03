package main

import (
	"bytes"
	"encoding/json"
	"github.com/yoshikouki/semapi/api"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"
)

const testURL = "http://localhost:8686/semapi"

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestServerConnection(t *testing.T) {
	res, err := http.Get(testURL + "/health-check")
	if err != nil {
		t.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err := res.Body.Close(); err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("/health-check returned wrong status code: got %d want %d", res.StatusCode, http.StatusOK)
	}

	expected := "pong"
	got := string(body)
	if got != expected {
		t.Errorf("Server returned wrong body: got %s want %s", got, expected)
	}
}

type testResponse struct {
	statusCode int
	body       string
}

func TestLock(t *testing.T) {
	statusCode, body := lockRequest(t, &api.LockParams{
		Target: "org-repo-stage",
		User:   "test",
		TTL:    "1s",
	})

	expected := testResponse{
		statusCode: 200,
		body:       "OK",
	}
	got := testResponse{
		statusCode: statusCode,
		body: string(body),
	}
	if got.statusCode != expected.statusCode {
		t.Errorf("/lock returned wrong status code: got %d want %d", got.statusCode, expected.statusCode)
	}
	if !strings.Contains(got.body, expected.body) {
		t.Errorf("/lock returned wrong body: got %s want %s", got.body, expected.body)
	}
}

func TestLockAndLock(t *testing.T) {
	lockRequest(t, &api.LockParams{
		Target: "org-repo-stage",
		User:   "test",
		TTL:    "1s",
	})
	statusCode, body := lockRequest(t, &api.LockParams{
		Target: "org-repo-stage",
		User:   "test",
		TTL:    "1s",
	})

	expected := testResponse{
		statusCode: 500,
		body:       "org-repo-stage is already locked.",
	}
	got := testResponse{
		statusCode: statusCode,
		body: string(body),
	}
	if got.statusCode != expected.statusCode {
		t.Errorf("/lock returned wrong status code: got %d want %d", got.statusCode, expected.statusCode)
	}
	if !strings.Contains(got.body, expected.body) {
		t.Errorf("/lock returned wrong body: got %s want %s", got.body, expected.body)
	}
}

func TestLockAndInvalidLock(t *testing.T) {
	lockRequest(t, &api.LockParams{
		Target: "org-repo-stage",
		User:   "test",
		TTL:    "1s",
	})
	statusCode, body := lockRequest(t, &api.LockParams{
		Target: "org-repo-stage",
		User:   "InvalidUser",
		TTL:    "1s",
	})

	expected := testResponse{
		statusCode: 500,
		body:       "org-repo-stage is locked by InvalidUser.",
	}
	got := testResponse{
		statusCode: statusCode,
		body: string(body),
	}
	if got.statusCode != expected.statusCode {
		t.Errorf("/lock returned wrong status code: got %d want %d", got.statusCode, expected.statusCode)
	}
	if !strings.Contains(got.body, expected.body) {
		t.Errorf("/lock returned wrong body: got %s want %s", got.body, expected.body)
	}
}

func TestLockAndUnlock(t *testing.T) {
	lockRequest(t, &api.LockParams{
		Target: "org-repo-stage",
		User:   "test",
		TTL:    "1s",
	})
	unlockBody := unlockRequest(t, &api.UnlockParams{
		Target: "org-repo-stage",
		User:   "test",
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
		Target: "org-repo-stage",
		User:   "test",
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
	lockRequest(t, &api.LockParams{
		Target: "org-repo-stage",
		User:   "test",
		TTL:    "1s",
	})
	unlockBody := unlockRequest(t, &api.UnlockParams{
		Target: "org-repo-stage",
		User:   "InvalidUser",
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

func lockRequest(t *testing.T, params *api.LockParams) (int, []byte) {
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

	return res.StatusCode, body
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
