// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package healthcheck

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

type HealthCheck struct {
	cert string `json:"cert"`
	http string `json:"http"`
}

func New() (*HealthCheck, error) {
	this := new(HealthCheck)
	// For now, this is just a constant stub.
	this.cert := "OK"
	this.http := "OK"
	return this, nil
}

// Returns false if any HealthCheck field is not set to "OK".
func (this *HealthCheck) IsHealthy() (bool) {
	if this.cert != "OK" {
		return false;
	}
	if this.http != "OK" {
		return false;
	}
	return true;
}

func (this *HealthCheck) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	content, err = json.MarshalIndent(this, "", "\t")
	if err != nil {
		http.Error(resp, err.Error(), http.StatusInternalServerError)
		return
	}

	if this.IsHealthy() {
	  resp.WriteHeader(http.StatusOK)
	} else {
	  resp.WriteHeader(http.StatusInternalServerError)
	}

	resp.Header().Set("Content-Type", "text/plain")
	resp.Header().Set("Cache-Control", "no-store")
	resp.Header().Set("X-Content-Type-Options", "nosniff")
	http.ServeContent(resp, req, "", time.Time{}, bytes.NewReader(content))
}
