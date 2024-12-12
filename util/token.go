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

package util

import (
	"errors"
	"fmt"
)

// GenerateTokenPair generates a string suitable to be included in a marshalled JSON object.
func GenerateTokenPair(token string) (string, error) {
	if token == "" {
		return "", errors.New("empty string")
	}

	return fmt.Sprintf(`"token":"%s"`, token), nil
}

func AppendTokenToJSONBody(body []byte, token string) ([]byte, error) {
	if token == "" {
		return nil, errors.New("empty string")
	}

	return append(body[:len(body)-1], []byte(fmt.Sprintf(`,%s}`, token))...), nil
}
