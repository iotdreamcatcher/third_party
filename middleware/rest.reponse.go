package middleware

import (
	"encoding/json"
	"net/http"
	"third_party/response"
	"third_party/sony"
)

func CommonErrResponse(w http.ResponseWriter, r *http.Request, Code int32, v ...any) {

	msg := ""
	if len(v) > 0 {
		msg = v[0].(string)
	} else {
		msg = response.StatusText(Code)
	}

	resp := response.CommonResponse{
		Code:      Code,
		Msg:       msg,
		RequestID: sony.NextId(),
		Path:      r.RequestURI,
	}
	body, _ := json.Marshal(&resp)
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}
