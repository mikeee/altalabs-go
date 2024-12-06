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
	"fmt"
	"net/url"
	"reflect"
	"strings"
)

// StructToParams converts a struct to encoded params for appending to URLs as part of paths.
func StructToParams(data interface{}) string {
	values := url.Values{}
	v := reflect.ValueOf(data)
	t := reflect.TypeOf(data)

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)
		tag := fieldType.Tag.Get("json")
		if tag == "" {
			tag = strings.ToLower(fieldType.Name)
		}
		values.Set(tag, fmt.Sprintf("%v", field.Interface()))
	}

	return values.Encode()
}
