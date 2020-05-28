package web

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

func Decode(in io.ReadCloser, out interface{}) error {
	payloadBytes, err := ioutil.ReadAll(in)
	if err != nil {
		return &Error{
			Err:       err,
			ErrorType: ValidationErrCode,
			Message:   "invalid payload",
		}
	}
	if err := json.Unmarshal(payloadBytes, out); err != nil {
		return &Error{
			Err:       err,
			ErrorType: ValidationErrCode,
			Message:   "invalid payload",
		}
	}

	return nil
}
