package webiface

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/shipperizer/miniature-monkey/types"
)

func ProcessHttpError(err error, httpStatus int, w http.ResponseWriter) {
	resp := new(types.ErrorResponse)
	resp.Status = httpStatus
	resp.Message = fmt.Sprint(err)

	w.WriteHeader(resp.Status)

	json.NewEncoder(w).Encode(resp)
}
