// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.


// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/135ae054e304239743b5777ad8d41cb2c9091d35


package types

// ReindexNode type.
//
// https://github.com/elastic/elasticsearch-specification/blob/135ae054e304239743b5777ad8d41cb2c9091d35/specification/_global/reindex_rethrottle/types.ts#L26-L28
type ReindexNode struct {
	Attributes       map[string]string      `json:"attributes"`
	Host             Host                   `json:"host"`
	Ip               Ip                     `json:"ip"`
	Name             Name                   `json:"name"`
	Roles            *NodeRoles             `json:"roles,omitempty"`
	Tasks            map[TaskId]ReindexTask `json:"tasks"`
	TransportAddress TransportAddress       `json:"transport_address"`
}

// ReindexNodeBuilder holds ReindexNode struct and provides a builder API.
type ReindexNodeBuilder struct {
	v *ReindexNode
}

// NewReindexNode provides a builder for the ReindexNode struct.
func NewReindexNodeBuilder() *ReindexNodeBuilder {
	r := ReindexNodeBuilder{
		&ReindexNode{
			Attributes: make(map[string]string, 0),
			Tasks:      make(map[TaskId]ReindexTask, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the ReindexNode struct
func (rb *ReindexNodeBuilder) Build() ReindexNode {
	return *rb.v
}

func (rb *ReindexNodeBuilder) Attributes(value map[string]string) *ReindexNodeBuilder {
	rb.v.Attributes = value
	return rb
}

func (rb *ReindexNodeBuilder) Host(host Host) *ReindexNodeBuilder {
	rb.v.Host = host
	return rb
}

func (rb *ReindexNodeBuilder) Ip(ip Ip) *ReindexNodeBuilder {
	rb.v.Ip = ip
	return rb
}

func (rb *ReindexNodeBuilder) Name(name Name) *ReindexNodeBuilder {
	rb.v.Name = name
	return rb
}

func (rb *ReindexNodeBuilder) Roles(roles *NodeRolesBuilder) *ReindexNodeBuilder {
	v := roles.Build()
	rb.v.Roles = &v
	return rb
}

func (rb *ReindexNodeBuilder) Tasks(values map[TaskId]*ReindexTaskBuilder) *ReindexNodeBuilder {
	tmp := make(map[TaskId]ReindexTask, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Tasks = tmp
	return rb
}

func (rb *ReindexNodeBuilder) TransportAddress(transportaddress TransportAddress) *ReindexNodeBuilder {
	rb.v.TransportAddress = transportaddress
	return rb
}
