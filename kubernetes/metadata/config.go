// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package metadata

import (
	"github.com/elastic/elastic-agent-libs/config"
)

// Config declares supported configuration for metadata generation
type Config struct {
	KubeConfig string `config:"kube_config"`

	UseRegexInclude    bool     `config:"use_regex_include"`
	UseRegexExclude    bool     `config:"use_regex_exclude"`
	IncludeLabels      []string `config:"include_labels"`
	ExcludeLabels      []string `config:"exclude_labels"`
	IncludeAnnotations []string `config:"include_annotations"`

	LabelsDedot      bool `config:"labels.dedot"`
	AnnotationsDedot bool `config:"annotations.dedot"`
}

// AddResourceMetadataConfig allows adding config for enriching additional resources
type AddResourceMetadataConfig struct {
	Node       *config.C `config:"node"`
	Namespace  *config.C `config:"namespace"`
	Deployment bool      `config:"deployment"`
	CronJob    bool      `config:"cronjob"`
}

// InitDefaults initializes the defaults for the config.
func (c *Config) InitDefaults() {
	c.LabelsDedot = true
	c.AnnotationsDedot = true
	c.UseRegexInclude = false
	c.UseRegexExclude = false
}

// Unmarshal unpacks a Config into the metagen Config
func (c *Config) Unmarshal(cfg *config.C) error {
	return cfg.Unpack(c)
}

func GetDefaultResourceMetadataConfig() *AddResourceMetadataConfig {
	metaConfig := Config{}
	metaConfig.InitDefaults()
	metaCfg, _ := config.NewConfigFrom(&metaConfig)
	return &AddResourceMetadataConfig{
		Node:       metaCfg,
		Namespace:  metaCfg,
		Deployment: false,
		CronJob:    false,
	}
}
