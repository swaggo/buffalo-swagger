package buffaloSwagger

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/x/sessions"
	"github.com/stretchr/testify/assert"
	"github.com/swaggo/buffalo-swagger/swaggerFiles"
	"github.com/swaggo/swag"
)

type mockedSwag struct{}

func (s *mockedSwag) ReadDoc() string {
	return `{
}`
}

func TestWrapHandler(t *testing.T) {
	router := buffalo.New(buffalo.Options{
		SessionStore: sessions.Null{},
		SessionName:  "_example_session",
	})

	router.ANY("/swagger/{doc:.*}", WrapHandler(swaggerFiles.Handler, InstanceName("")))

	w1 := performRequest(http.MethodGet, "/swagger/index.html", router)
	assert.Equal(t, http.StatusOK, w1.Code)
	assert.Equal(t, w1.Header()["Content-Type"][0], "text/html; charset=utf-8")
	w1BodyBytes, _ := ioutil.ReadAll(w1.Body)
	assert.Contains(t, string(w1BodyBytes), "doc.json")

	assert.Equal(t, http.StatusInternalServerError, performRequest(http.MethodGet, "/swagger/doc.json", router).Code)

	swag.Register(swag.Name, &mockedSwag{})

	w2 := performRequest(http.MethodGet, "/swagger/doc.json", router)
	assert.Equal(t, http.StatusOK, w2.Code)
	assert.Equal(t, w2.Header()["Content-Type"][0], "application/json; charset=utf-8")

	w3 := performRequest(http.MethodGet, "/swagger/favicon-16x16.png", router)
	assert.Equal(t, http.StatusOK, w3.Code)
	assert.Equal(t, w3.Header()["Content-Type"][0], "image/png")

	w4 := performRequest(http.MethodGet, "/swagger/swagger-ui.css", router)
	assert.Equal(t, http.StatusOK, w4.Code)
	assert.Equal(t, w4.Header()["Content-Type"][0], "text/css; charset=utf-8")

	w5 := performRequest(http.MethodGet, "/swagger/swagger-ui-bundle.js", router)
	assert.Equal(t, http.StatusOK, w5.Code)
	assert.Equal(t, w5.Header()["Content-Type"][0], "application/javascript")

	assert.Equal(t, http.StatusNotFound, performRequest(http.MethodGet, "/swagger/notfound", router).Code)

	assert.Equal(t, http.StatusMethodNotAllowed, performRequest(http.MethodPost, "/swagger/index.html", router).Code)

	assert.Equal(t, http.StatusMethodNotAllowed, performRequest(http.MethodPut, "/swagger/index.html", router).Code)

}

func TestCustomWrapHandler(t *testing.T) {
	router := buffalo.New(buffalo.Options{
		SessionStore: sessions.Null{},
		SessionName:  "_example_session",
	})

	router.GET("/swagger/{doc:.*}", WrapHandler(swaggerFiles.Handler, InstanceName("custom")))

	w1 := performRequest(http.MethodGet, "/swagger/index.html", router)
	assert.Equal(t, http.StatusOK, w1.Code)
	assert.Equal(t, w1.Header()["Content-Type"][0], "text/html; charset=utf-8")

	swag.Register("custom", &mockedSwag{})

	w2 := performRequest(http.MethodGet, "/swagger/doc.json", router)
	assert.Equal(t, http.StatusOK, w2.Code)
	assert.Equal(t, w2.Header()["Content-Type"][0], "application/json; charset=utf-8")

}

func performRequest(method, target string, router *buffalo.App) *httptest.ResponseRecorder {
	r := httptest.NewRequest(method, target, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)

	return w
}

func TestURL(t *testing.T) {
	cfg := Config{}

	expected := "https://github.com/swaggo/http-swagger"
	configFunc := URL(expected)
	configFunc(&cfg)
	assert.Equal(t, expected, cfg.URL)
}

func TestDocExpansion(t *testing.T) {
	var cfg Config

	expected := "list"
	configFunc := DocExpansion(expected)
	configFunc(&cfg)
	assert.Equal(t, expected, cfg.DocExpansion)

	expected = "full"
	configFunc = DocExpansion(expected)
	configFunc(&cfg)
	assert.Equal(t, expected, cfg.DocExpansion)

	expected = "none"
	configFunc = DocExpansion(expected)
	configFunc(&cfg)
	assert.Equal(t, expected, cfg.DocExpansion)
}

func TestDeepLinking(t *testing.T) {
	var cfg Config
	assert.Equal(t, false, cfg.DeepLinking)

	configFunc := DeepLinking(true)
	configFunc(&cfg)
	assert.Equal(t, true, cfg.DeepLinking)

	configFunc = DeepLinking(false)
	configFunc(&cfg)
	assert.Equal(t, false, cfg.DeepLinking)

}

func TestDefaultModelsExpandDepth(t *testing.T) {
	var cfg Config

	assert.Equal(t, 0, cfg.DefaultModelsExpandDepth)

	expected := -1
	configFunc := DefaultModelsExpandDepth(expected)
	configFunc(&cfg)
	assert.Equal(t, expected, cfg.DefaultModelsExpandDepth)

	expected = 1
	configFunc = DefaultModelsExpandDepth(expected)
	configFunc(&cfg)
	assert.Equal(t, expected, cfg.DefaultModelsExpandDepth)
}

func TestInstanceName(t *testing.T) {
	var cfg Config

	assert.Equal(t, "", cfg.InstanceName)

	expected := swag.Name
	configFunc := InstanceName(expected)
	configFunc(&cfg)
	assert.Equal(t, expected, cfg.InstanceName)

	expected = "custom_name"
	configFunc = InstanceName(expected)
	configFunc(&cfg)
	assert.Equal(t, expected, cfg.InstanceName)
}

func TestPersistAuthorization(t *testing.T) {
	var cfg Config
	assert.Equal(t, false, cfg.PersistAuthorization)

	configFunc := PersistAuthorization(true)
	configFunc(&cfg)
	assert.Equal(t, true, cfg.PersistAuthorization)

	configFunc = PersistAuthorization(false)
	configFunc(&cfg)
	assert.Equal(t, false, cfg.PersistAuthorization)
}
