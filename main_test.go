package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetCommentHttp(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "pong")
	}

	req := httptest.NewRequest("GET", "localhost:4000/api/comment/1", nil)

	w := httptest.NewRecorder()
	handler(w, req)

	if w.Code == 404 {
		t.Error("http error")
	}
	// Return value test
	if w.Code != 200 {
		t.Error("Not match")
	}

}

func TestGetArticleListHttp(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "pong")
	}

	req := httptest.NewRequest("GET", "localhost:4000/api/article/list?page_size=20&page=1", nil)

	w := httptest.NewRecorder()
	handler(w, req)

	if w.Code == 404 {
		t.Error("http error")
	}
	// Return value test
	if w.Code != 200 {
		t.Error("Not match")
	}

}
