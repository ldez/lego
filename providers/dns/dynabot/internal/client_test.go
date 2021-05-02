package internal

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClient_SetDNS(t *testing.T) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)

	client := NewClient("secret")
	client.baseURL = server.URL

	mux.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		data, err := ioutil.ReadFile(fmt.Sprintf("./fixtures/%s.xml", commandSetDNS))
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		rw.Header().Set("Content-Type", "application/xml")
		_, _ = rw.Write(data)
	})

	_, err := client.SetDNS("example.com", "www", "TXT", "content", 300)
	require.NoError(t, err)
}

func TestClient_SetDNS_error_key(t *testing.T) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)

	client := NewClient("secret")
	client.baseURL = server.URL

	mux.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		data, err := ioutil.ReadFile(fmt.Sprintf("./fixtures/%s.xml", "error"))
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		rw.Header().Set("Content-Type", "application/xml")
		_, _ = rw.Write(data)
	})

	_, err := client.SetDNS("example.com", "www", "TXT", "content", 300)
	require.Error(t, err)
}

func TestClient_SetDNS_error(t *testing.T) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)

	client := NewClient("secret")
	client.baseURL = server.URL

	mux.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		data, err := ioutil.ReadFile(fmt.Sprintf("./fixtures/%s.xml", commandSetDNS+"_error"))
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		rw.Header().Set("Content-Type", "application/xml")
		_, _ = rw.Write(data)
	})

	_, err := client.SetDNS("example.com", "www", "TXT", "content", 300)
	require.Error(t, err)
}

func TestClient_IsProcessing(t *testing.T) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)

	client := NewClient("secret")
	client.baseURL = server.URL

	mux.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		data, err := ioutil.ReadFile(fmt.Sprintf("./fixtures/%s.xml", commandIsProcessing))
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		rw.Header().Set("Content-Type", "application/xml")
		_, _ = rw.Write(data)
	})

	ok, err := client.IsProcessing()
	require.NoError(t, err)

	assert.True(t, ok)
}

func TestClient_IsProcessing_error(t *testing.T) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)

	client := NewClient("secret")
	client.baseURL = server.URL

	mux.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		data, err := ioutil.ReadFile(fmt.Sprintf("./fixtures/%s.xml", "error"))
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		rw.Header().Set("Content-Type", "application/xml")
		_, _ = rw.Write(data)
	})

	ok, err := client.IsProcessing()
	require.Error(t, err)

	assert.False(t, ok)
}
