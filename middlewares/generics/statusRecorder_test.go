package generics

import (
	"testing"

	"net/http"
	"net/http/httptest"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestStatusRecorderImplementsResponseWriter(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	r := &statusRecorder{httptest.NewRecorder(), http.StatusAccepted}

	assert := assert.New(t)
	assert.Implements((*http.ResponseWriter)(nil), r)
}

func TestStatusRecorderSetsStatus(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	r := &statusRecorder{httptest.NewRecorder(), http.StatusAccepted}

	assert := assert.New(t)

	assert.Equal(http.StatusAccepted, r.status)

	r.WriteHeader(http.StatusConflict)
	assert.Equal(http.StatusConflict, r.status)
}
