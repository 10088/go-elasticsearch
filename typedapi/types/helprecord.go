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
// https://github.com/elastic/elasticsearch-specification/tree/93ed2b29c9e75f49cd340f06286d6ead5965f900


package types

// HelpRecord type.
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/cat/help/types.ts#L20-L22
type HelpRecord struct {
	Endpoint string `json:"endpoint"`
}

// HelpRecordBuilder holds HelpRecord struct and provides a builder API.
type HelpRecordBuilder struct {
	v *HelpRecord
}

// NewHelpRecord provides a builder for the HelpRecord struct.
func NewHelpRecordBuilder() *HelpRecordBuilder {
	r := HelpRecordBuilder{
		&HelpRecord{},
	}

	return &r
}

// Build finalize the chain and returns the HelpRecord struct
func (rb *HelpRecordBuilder) Build() HelpRecord {
	return *rb.v
}

func (rb *HelpRecordBuilder) Endpoint(endpoint string) *HelpRecordBuilder {
	rb.v.Endpoint = endpoint
	return rb
}