// Copyright © 2021 Alexandre KOVAC <contact@kovacou.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package echox

import (
	"github.com/kovacou/go-convert"
	"github.com/kovacou/go-types"
	"github.com/labstack/echo/v4"
)

// ParamID returns the value of the parameter :id
func ParamID(ctx echo.Context) uint64 {
	return convert.Uint64(ctx.Param("id"))
}

// ParamDate returns the value of the parameter :date
func ParamDate(ctx echo.Context) types.Date {
	return ParseDate(ctx.Param("date"))
}
