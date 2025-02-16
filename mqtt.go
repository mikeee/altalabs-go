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

package altalabs

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"os"
	"time"
)

func (a *AltaClient) MqttConn() error {
	testSite := os.Getenv("SDK_ALTA_SITE")

	token := a.AuthClient.GetIDToken()
	if token == "" {
		return fmt.Errorf("token is empty")
	}

	fullurl := "wss://manage.alta.inc/mqtt?x-amz-customauthorizer-name=DeviceAuth-prod"
	fullurl += fmt.Sprintf("&token=%s", token)

	fullurl += fmt.Sprintf("&site=%s", testSite)

	fullurl += "&fe=web"

	fullurl += "&version=9b698a96"

	fullurl += fmt.Sprintf("&timestamp=%v", time.Now().UnixMilli())

	fullurl += "&seq=1" // TODO: increment seq

	opts := mqtt.NewClientOptions().
		AddBroker(fullurl).
		SetKeepAlive(5 * time.Second).
		SetPingTimeout(5 * time.
			Second)

	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	defer c.Disconnect(1)

	fmt.Println("Connected to mqtt")

	return nil
}
