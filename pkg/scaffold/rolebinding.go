// Copyright 2018 The Operator-SDK Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package scaffold

import (
	"path/filepath"

	"github.com/operator-framework/operator-sdk/pkg/scaffold/input"
)

type RoleBinding struct {
	input.Input
}

func (s *RoleBinding) GetInput() (input.Input, error) {
	if s.Path == "" {
		s.Path = filepath.Join(deployDir, roleBindingYamlFile)
	}
	s.TemplateBody = roleBindingTemplate
	return s.Input, nil
}

const roleBindingTemplate = `kind: RoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: {{.ProjectName}}
subjects:
- kind: ServiceAccount
  name: {{.ProjectName}}
roleRef:
  kind: Role
  name: {{.ProjectName}}
  apiGroup: rbac.authorization.k8s.io
`
