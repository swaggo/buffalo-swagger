package buffaloSwagger

import (
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

	router.GET("/swagger/{doc:.*}", WrapHandler(swaggerFiles.Handler))

	w1 := performRequest("GET", "/swagger/index.html", router)
	assert.Equal(t, 200, w1.Code)

	swag.Register(swag.Name, &mockedSwag{})

	w2 := performRequest("GET", "/swagger/doc.json", router)
	assert.Equal(t, 200, w2.Code)

	w3 := performRequest("GET", "/swagger/favicon-16x16.png", router)
	assert.Equal(t, 200, w3.Code)

	w4 := performRequest("GET", "/swagger/swagger-ui.css", router)
	assert.Equal(t, 200, w4.Code)

	w5 := performRequest("GET", "/swagger/notfound", router)
	assert.Equal(t, 404, w5.Code)
}

func performRequest(method, target string, router *buffalo.App) *httptest.ResponseRecorder {
	r := httptest.NewRequest(method, target, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)
	return w
}
