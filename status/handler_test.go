package status

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	chi "github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"

	"github.com/shipperizer/miniature-monkey/v2/types"
)

func TestStatusSucceeds(t *testing.T) {
	bp := NewBlueprint()

	req, err := http.NewRequest(http.MethodGet, "/api/v1/status", nil)
	req.Header.Set("Content-Type", "application/json")

	assert := assert.New(t)

	assert.Equal(nil, err, "error should be nil")

	rr := httptest.NewRecorder()
	router := chi.NewMux()
	bp.Routes(router)
	router.ServeHTTP(rr, req)

	resp := new(types.DataResponse)
	_ = json.Unmarshal(rr.Body.Bytes(), &resp)

	assert.Equal(http.StatusOK, rr.Code, "should equal")
	assert.Equal("ok", resp.Message, "should equal")
}
