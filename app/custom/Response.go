package custom

// Best way to manage this is to extend net/http package to provide a interface to customize the responses
// Then this can be injected directly off the controller

import (
	"fmt"
	"net/http"
	"encoding/json"
)

type Object interface{}

type Response struct {
	Error bool `json:"error"`
	StatusCode int `json:"code"`
	Message string `json:"message"`
	Data interface{} `json:"response"`
}

func (response *Response) Format(w http.ResponseWriter, r *http.Request, err bool, code int, o Object) {
	fmt.Println(o)
	
	// status codes
	m := make(map[int]string)

	m[200] = "OK"
	m[400] = "Bad Request."
	m[404] = "Resource Not Found."
	m[416] = "Resource Busy Please Try Again."
	m[417] = "Validation Failed. Please Verify Input."
	m[418] = "Unable to create a new record."


	var status_code int

	if (code == 416 || code == 417 || code == 418) {
		status_code = 400
	}else{
		status_code = code
	}

	resp := Response{
		Error: err,
		Data: o,
		Message: m[code],
		StatusCode: status_code,
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status_code)
    json.NewEncoder(w).Encode(resp)

}
