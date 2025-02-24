// Copyright 2025 Nitro Agility S.r.l.
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
//
// SPDX-License-Identifier: Apache-2.0

package azreq

// PolicyStore represents the Policy Store.
type PolicyStore struct {
	kind string
	id   string
}

// GetKind returns the kind of the Policy Store.
func (p *PolicyStore) GetKind() string {
	return p.kind
}

// GetID returns the ID of the Policy Store.
func (p *PolicyStore) GetID() string {
	return p.id
}
