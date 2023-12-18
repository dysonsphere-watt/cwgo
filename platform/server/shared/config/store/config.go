/*
 *
 * Copyright 2023 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package store

import (
	"github.com/cloudwego/cwgo/platform/server/shared/consts"
)

type Config struct {
	Type  string `mapstructure:"type"`
	Mysql Mysql  `mapstructure:"mysql"`
	Mongo Mongo  `mapstructure:"mongo"`
	Redis Redis  `mapstructure:"redis"`
}

func (conf *Config) SetUp() {
	conf.setDefaults()
}

func (conf *Config) setDefaults() {
	conf.Type = consts.StoreTypeMysql
}

func (conf *Config) GetStoreType() consts.StoreType {
	return consts.StoreTypeMapToNum[conf.Type]
}