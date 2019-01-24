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

package interfaces

import (
	"database/sql"
)

type SQL interface {
	GetDB() *sql.DB
	GetDBName() RDBMSInfo
}

type RDBMSInfo struct {
	Name    string
	Version string
}

const (
	Oracle             = "Oracle"
	MariaDB            = "MariaDB"
	MicrosoftSQLSerevr = "MicrosoftSQLSerevr"
	MySQL              = "MySQL"
	PostgreSQL         = "PostgreSQL"
	SQLite             = "SQLite"
)