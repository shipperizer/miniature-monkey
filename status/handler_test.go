package status

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"

	"github.com/shipperizer/miniature-monkey/types"
)

func TestStatusSucceeds(t *testing.T) {
	bp := NewBlueprint()

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/status"), nil)
	req.Header.Set("Content-Type", "application/json")

	assert := assert.New(t)

	assert.Equal(nil, err, "error should be nil")

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	bp.Routes(router)
	router.ServeHTTP(rr, req)

	resp := new(types.DataResponse)
	_ = json.Unmarshal(rr.Body.Bytes(), &resp)

	assert.Equal(http.StatusOK, rr.Code, "should equal")
	assert.Equal("ok", resp.Message, "should equal")
}
