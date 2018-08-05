// Copyright (c) OpenFaaS Author(s) 2018. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

package function

import (
	"net/http"
	"os"

	"github.com/openfaas-incubator/go-function-sdk"
)

// Handle a function invocation
func Handle(req handler.Request) (handler.Response, error) {
	var err error
	res := handler.Response{}

	if !isAuthorized(req) {
		message := "You must authorize."
		res.Body = []byte(message)
		res.StatusCode = http.StatusUnauthorized

		res.Header = http.Header{
			"WWW-Authenticate": []string{`Basic realm="Restricted"`},
		}
		return res, err
	}

	res.StatusCode = http.StatusOK
	res.Body = []byte("Authorization. OK.")

	return res, err
}

func isAuthorized(req handler.Request) bool {
	r := http.Request{}
	r.Header = http.Header{
		"Authorization": []string{req.Header.Get("Authorization")},
	}

	if username, password, ok := r.BasicAuth(); ok &&
		username == os.Getenv("BASIC_AUTH_USERNAME") &&
		password == os.Getenv("BASIC_AUTH_PASSWORD") {
		return true
	}
	return false
}
