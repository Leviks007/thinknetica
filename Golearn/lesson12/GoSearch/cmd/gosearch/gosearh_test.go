package main

import (
	"lesson12/GoSearch/pkg/crawler"
	indexDoc "lesson12/GoSearch/pkg/index"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_handlerDoc(t *testing.T) {
	Index = indexDoc.New()

	documents := []crawler.Document{
		{ID: 1, URL: "https://example.com", Title: "Example Website", Body: "This is an example website."},
		{ID: 2, URL: "https://example.org", Title: "Another Example Website", Body: "This is another example website."},
		{ID: 3, URL: "https://example.net", Title: "Yet Another Example", Body: "Yet another example website here."},
	}

	Index.AddDocuments(documents)

	req, err := http.NewRequest("GET", "/doc", nil)
	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()

	handlerDoc(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Ожидался код статуса %v, получено %v", http.StatusOK, rec.Code)
	}

	expected := Index.StringDoc()

	if rec.Body.String() != expected {
		t.Errorf("Ожидаемый результат: %v, получено: %v", expected, rec.Body.String())
	}
}

func Test_handlerIndex(t *testing.T) {
	Index = indexDoc.New()

	documents := []crawler.Document{
		{ID: 1, URL: "https://example.com", Title: "Example Website", Body: "This is an example website."},
		{ID: 2, URL: "https://example.org", Title: "Another Example Website", Body: "This is another example website."},
		{ID: 3, URL: "https://example.net", Title: "Yet Another Example", Body: "Yet another example website here."},
	}
	// Добавляем документы в индекс
	Index.AddDocuments(documents)

	// Создаем новый запрос с методом GET и параметром "s" равным "example"
	req, err := http.NewRequest("GET", "/index", nil)
	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()

	handlerIndex(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Ожидался код статуса %v, получено %v", http.StatusOK, rec.Code)
	}

	expected := Index.String()

	if rec.Body.String() != expected {
		t.Errorf("Ожидаемый результат: %v, получено: %v", expected, rec.Body.String())
	}
}
