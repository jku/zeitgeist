/*
Copyright 2020 The Kubernetes Authors.

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

/* Zeitgeist is a language-agnostic dependency checker */

package main

import (
	"github.com/sirupsen/logrus"

	"sigs.k8s.io/zeitgeist/commands"
)

func main() {
	if err := commands.New().Execute(); err != nil {
		logrus.Fatalf("error during command execution: %#v", err)
	}
}
