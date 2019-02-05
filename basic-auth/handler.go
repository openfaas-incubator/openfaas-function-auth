// Copyright (c) OpenFaaS Author(s) 2018. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

package function

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/openfaas-incubator/go-function-sdk"
)

var (
	usernameSecret string
	passwordSecret string
)

func init() {

	// optional, add error handling, or read on each request / use sync.Once()

	usernameRaw, err := ioutil.ReadFile("/var/openfaas/secrets/fn-basic-auth-username")
	if err != nil {
		panic(err)
	}
	usernameSecret = strings.TrimSpace(string(usernameRaw))

	passwordRaw, err := ioutil.ReadFile("/var/openfaas/secrets/fn-basic-auth-password")
	if err != nil {
		panic(err)
	}
	passwordSecret = strings.TrimSpace(string(passwordRaw))
}

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
		username == usernameSecret &&
		password == passwordSecret {
		return true
	}
	return false
}
