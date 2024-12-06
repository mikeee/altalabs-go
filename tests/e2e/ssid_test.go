//go:build e2e

/*
Copyright 2024 Mike Nguyen (mikeee) <hey@mike.ee>
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package e2e

import (
	"github.com/mikeee/altalabs-go"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_SSID(t *testing.T) {
	client, err := altalabs.NewAltaClient(os.Getenv("SDK_ALTA_USER"), os.Getenv("SDK_ALTA_PASS"))
	if err != nil {
		panic(err)
	}

	ssidList, err := client.ListSSID()
	if err != nil {
		panic(err)
	}

	t.Run("ListSSID should return a list of SSIDs", func(t *testing.T) {
		assert.NotEmpty(t, ssidList)
	})

	// TODO: Test get/add/update/delete methods
}
