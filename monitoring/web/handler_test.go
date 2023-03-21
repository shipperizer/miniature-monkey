package monitoring

import (
	"net/http"
	"net/http/httptest"
	"testing"

	chi "github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

func TestMetricsBlueprintWorks(t *testing.T) {
	bp := NewBlueprint()

	req, err := http.NewRequest(http.MethodGet, "/api/v1/metrics", nil)
	req.Header.Set("Content-Type", "application/json")

	assert := assert.New(t)

	assert.Equal(nil, err, "error should be nil")

	rr := httptest.NewRecorder()
	router := chi.NewMux()
	bp.Routes(router)
	router.ServeHTTP(rr, req)

	assert.Equal(http.StatusOK, rr.Code, "should equal")
}
