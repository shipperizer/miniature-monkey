package monitoring

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestMetricsBlueprintWorks(t *testing.T) {
	bp := NewBlueprint()

	req, err := http.NewRequest(http.MethodGet, "/api/v1/metrics", nil)
	req.Header.Set("Content-Type", "application/json")

	assert := assert.New(t)

	assert.Equal(nil, err, "error should be nil")

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	bp.Routes(router)
	router.ServeHTTP(rr, req)

	assert.Equal(http.StatusOK, rr.Code, "should equal")
}
