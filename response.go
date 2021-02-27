// Copyright © 2021 Alexandre KOVAC <contact@kovacou.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package echox

import (
	"net/http"

	"github.com/kovacou/go-types"
	"github.com/labstack/echo"
)

// OK is the representation of the response wrapper.
type OK struct {
	Code        int         `json:"code"`
	CodeIfEmpty int         `json:"-"`
	Meta        types.Map   `json:"meta"`
	Data        interface{} `json:"data"`
}

// JSON returns json encoding of OK.
func (ok OK) JSON(ctx echo.Context) error {
	if ok.Code == 0 {
		ok.Code = http.StatusOK
	}

	if ok.Data == nil {
		if ok.CodeIfEmpty > 0 {
			ok.Code = ok.CodeIfEmpty
		}

		if ok.Code > 300 {
			return ErrorWrapper(ok.Code, http.StatusText(ok.Code)).JSON(ctx)
		}

		ok.Data = http.StatusText(ok.Code)
	}

	return EncodeJSON(ctx, ok.Code, ok)
}

// Error is the representation of the error wrapper.
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// JSON returns json encoding of Error.
func (err Error) JSON(ctx echo.Context) error {
	if err.Code == 0 {
		err.Code = http.StatusInternalServerError
	}

	if err.Message == "" {
		err.Message = http.StatusText(err.Code)
	}

	return EncodeJSON(ctx, err.Code, err)
}

// Error
func (err *Error) Error() string {
	return err.Message
}

// ErrorWrapper is wrapping code & error message.
func ErrorWrapper(code int, message string) Error {
	return Error{
		Code:    code,
		Message: message,
	}
}
