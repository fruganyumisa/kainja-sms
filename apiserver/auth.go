// Copyright 2015 sms-api-server authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package apiserver

import (
	"net/http"
)

// auth provides an authentication handler. Hook up your own stuff here. cool
func auth(f http.HandlerFunc) http.HandlerFunc {
	//log.Printf("Handler returned after authentication: %v", )
	return f
}
