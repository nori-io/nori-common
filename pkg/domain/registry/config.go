/*
Copyright 2018-2020 The Nori Authors.
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

package registry

import (
	"github.com/nori-io/common/v5/pkg/domain/config"
	enum "github.com/nori-io/common/v5/pkg/domain/enum/config"
	"github.com/nori-io/common/v5/pkg/domain/meta"
)

//go:generate mockgen -destination=../mocks/registry/config_registry.go -package=mocks github.com/nori-io/common/v5/pkg/domain/registry ConfigRegistry

type (
	ConfigRegistry interface {
		Register(id meta.ID) config.Config
		PluginVariables(id meta.ID) []Variable
	}

	Variable struct {
		Name        string
		Description string
		Type        enum.Kind
	}
)
