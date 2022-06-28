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

// SimpleValueAggregate type.
//
// https://github.com/elastic/elasticsearch-specification/blob/135ae054e304239743b5777ad8d41cb2c9091d35/specification/_types/aggregations/Aggregate.ts#L213-L214
type SimpleValueAggregate struct {
	Meta *Metadata `json:"meta,omitempty"`
	// Value The metric value. A missing value generally means that there was no data to
	// aggregate,
	// unless specified otherwise.
	Value         float64 `json:"value,omitempty"`
	ValueAsString *string `json:"value_as_string,omitempty"`
}

// SimpleValueAggregateBuilder holds SimpleValueAggregate struct and provides a builder API.
type SimpleValueAggregateBuilder struct {
	v *SimpleValueAggregate
}

// NewSimpleValueAggregate provides a builder for the SimpleValueAggregate struct.
func NewSimpleValueAggregateBuilder() *SimpleValueAggregateBuilder {
	r := SimpleValueAggregateBuilder{
		&SimpleValueAggregate{},
	}

	return &r
}

// Build finalize the chain and returns the SimpleValueAggregate struct
func (rb *SimpleValueAggregateBuilder) Build() SimpleValueAggregate {
	return *rb.v
}

func (rb *SimpleValueAggregateBuilder) Meta(meta *MetadataBuilder) *SimpleValueAggregateBuilder {
	v := meta.Build()
	rb.v.Meta = &v
	return rb
}

// Value The metric value. A missing value generally means that there was no data to
// aggregate,
// unless specified otherwise.

func (rb *SimpleValueAggregateBuilder) Value(value float64) *SimpleValueAggregateBuilder {
	rb.v.Value = value
	return rb
}

func (rb *SimpleValueAggregateBuilder) ValueAsString(valueasstring string) *SimpleValueAggregateBuilder {
	rb.v.ValueAsString = &valueasstring
	return rb
}
