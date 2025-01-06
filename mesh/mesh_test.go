/*
 * Copyright 2024 Holger de Carne
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
 */

package mesh_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tdrn-org/go-tr064/mesh"
)

func TestMesh(t *testing.T) {
	meshlistBytes, err := os.ReadFile("testdata/meshlist.json")
	require.NoError(t, err)
	list := &mesh.List{}
	err = json.Unmarshal(meshlistBytes, list)
	require.NoError(t, err)
	connections := list.Connections()
	require.Equal(t, 21, len(connections))
}