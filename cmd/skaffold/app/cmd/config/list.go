/*
Copyright 2018 The Skaffold Authors

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

package config

import (
	"io"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	yaml "gopkg.in/yaml.v2"
)

func NewCmdList(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all values set in the global skaffold config",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			resolveKubectlContext()
			return runList(out)
		},
	}
	AddConfigFlags(cmd)
	AddListFlags(cmd)
	return cmd
}

func runList(out io.Writer) error {
	var configYaml []byte
	if showAll {
		cfg, err := readConfig()
		if err != nil {
			return err
		}
		configYaml, err = yaml.Marshal(&cfg)
		if err != nil {
			return errors.Wrap(err, "marshaling config")
		}
	} else {
		config, err := getConfigForKubectx()
		if err != nil {
			return err
		}
		configYaml, err = yaml.Marshal(&config)
		if err != nil {
			return errors.Wrap(err, "marshaling config")
		}
	}

	out.Write(configYaml)
	return nil
}
