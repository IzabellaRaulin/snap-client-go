// +build small

/*
http://www.apache.org/licenses/LICENSE-2.0.txt


Copyright 2017 Intel Corporation

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

package tests

import (
	"testing"

	openapiclient "github.com/go-openapi/runtime/client"
	"github.com/intelsdi-x/snap-client-go/client"
	"github.com/intelsdi-x/snap-client-go/client/plugins"
	. "github.com/smartystreets/goconvey/convey"
)

func TestClient(t *testing.T) {

	Convey("Testing Client", t, func() {
		Convey("Test a client with the default config", func() {
			c := client.NewHTTPClientWithConfig(nil, &client.TransportConfig{
				Host:     client.DefaultHost,
				BasePath: client.DefaultBasePath,
				Schemes:  []string{client.DefaultSchemes[0]},
			})
			So(c, ShouldNotBeNil)
		})
		Convey("Test get a plugin", func() {
			c := client.NewHTTPClient(nil)
			_, err := c.Plugins.GetPlugin(plugins.NewGetPluginParams(), openapiclient.BasicAuth("snap", "snap"))
			So(err, ShouldNotBeNil)
		})
	})
}
