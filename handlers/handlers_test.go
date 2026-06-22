package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"int-cont/store"

	"github.com/gorilla/mux"
)

func setup() {
	store.Default = store.NewMockStore()
}

func TestPushAndPop(t *testing.T) {
	setup()

	// Push
	body := bytes.NewBufferString(`{"value": "hola"}`)
	req := httptest.NewRequest(http.MethodPost, "/stack/test/push", body)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/stack/{name}/push", Push).Methods("POST")
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Push esperaba 200, got %d", w.Code)
	}

	// Pop
	req = httptest.NewRequest(http.MethodPost, "/stack/test/pop", nil)
	w = httptest.NewRecorder()
	router2 := mux.NewRouter()
	router2.HandleFunc("/stack/{name}/pop", Pop).Methods("POST")
	router2.ServeHTTP(w, req)

	var result map[string]string
	json.NewDecoder(w.Body).Decode(&result)
	if result["value"] != "hola" {
		t.Fatalf("Pop esperaba 'hola', got '%s'", result["value"])
	}
}

func TestEnqueueAndDequeue(t *testing.T) {
	setup()

	body := bytes.NewBufferString(`{"value": "primero"}`)
	req := httptest.NewRequest(http.MethodPost, "/queue/test/enqueue", body)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/queue/{name}/enqueue", Enqueue).Methods("POST")
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Enqueue esperaba 200, got %d", w.Code)
	}

	req = httptest.NewRequest(http.MethodPost, "/queue/test/dequeue", nil)
	w = httptest.NewRecorder()
	router2 := mux.NewRouter()
	router2.HandleFunc("/queue/{name}/dequeue", Dequeue).Methods("POST")
	router2.ServeHTTP(w, req)

	var result map[string]string
	json.NewDecoder(w.Body).Decode(&result)
	if result["value"] != "primero" {
		t.Fatalf("Dequeue esperaba 'primero', got '%s'", result["value"])
	}
}

func TestAppendAndRemove(t *testing.T) {
	setup()

	body := bytes.NewBufferString(`{"value": "nodo1"}`)
	req := httptest.NewRequest(http.MethodPost, "/list/test/append", body)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/list/{name}/append", Append).Methods("POST")
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Append esperaba 200, got %d", w.Code)
	}

	req = httptest.NewRequest(http.MethodPost, "/list/test/remove", nil)
	w = httptest.NewRecorder()
	router2 := mux.NewRouter()
	router2.HandleFunc("/list/{name}/remove", Remove).Methods("POST")
	router2.ServeHTTP(w, req)

	var result map[string]string
	json.NewDecoder(w.Body).Decode(&result)
	if result["value"] != "nodo1" {
		t.Fatalf("Remove esperaba 'nodo1', got '%s'", result["value"])
	}
}
