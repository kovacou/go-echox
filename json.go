// Copyright © 2021 Alexandre KOVAC <contact@kovacou.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package echox

import (
	"github.com/labstack/echo"
	"github.com/wI2L/jettison"
)

// EncodeJSON sends a JSON response with status code.
func EncodeJSON(ctx echo.Context, code int, i interface{}) error {
	r := ctx.Response()
	r.Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	r.Status = code

	json, _ := jettison.MarshalOpts(i, jettison.NilSliceEmpty())
	_, err := r.Write(json)
	return err
}
