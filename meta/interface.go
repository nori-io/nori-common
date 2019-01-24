/*
Copyright 2019 The Nori Authors.

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

package meta

type Interface string

const (
	Custom        Interface = "Custom"
	Auth          Interface = "Auth"
	Authorize     Interface = "Authorize"
	Cache         Interface = "Cache"
	Config        Interface = "Config"
	HTTP          Interface = "HTTP"
	HTTPTransport Interface = "HTTPTransport"
	Mail          Interface = "Mail"
	PubSub        Interface = "PubSub"
	Session       Interface = "Session"
	SQL           Interface = "SQL"
	Templates     Interface = "Templates"
	Transport     Interface = "Transport"
)

func (i Interface) Dependency(ver string) Dependency {
	return Dependency{
		ID:         "",
		Constraint: "~" + ver,
		Interface:  i,
	}
}