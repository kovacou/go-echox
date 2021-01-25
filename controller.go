// Copyright © 2021 Alexandre KOVAC <contact@kovacou.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package echox

import "net/http"

// Index is a wrapper that returns a list of values as JSON.
func Index(values interface{}, n int) OK {
	resp := OK{CodeIfEmpty: http.StatusNoContent}
	if n > 0 {
		resp.Data = values
	}
	return resp
}

// Show is a wrapper that returns a value as JSON.
func Show(val interface{}, err error) OK {
	resp := OK{CodeIfEmpty: http.StatusNotFound}
	if err == nil {
		resp.Data = val
	}
	return resp
}
