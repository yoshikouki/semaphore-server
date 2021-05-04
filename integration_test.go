package main

import (
	"io/ioutil"
	"net/http"
	"os"
	"testing"
)

const testHost = "http://localhost:8686"

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestServerConnection(t *testing.T) {
	res, err := http.Get(testHost + "/ping")
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
	res, err := http.Get(testHost + "/redis/ping")
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
