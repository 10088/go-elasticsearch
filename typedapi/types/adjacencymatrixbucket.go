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

// AdjacencyMatrixBucket type.
//
// https://github.com/elastic/elasticsearch-specification/blob/135ae054e304239743b5777ad8d41cb2c9091d35/specification/_types/aggregations/Aggregate.ts#L556-L556
type AdjacencyMatrixBucket struct {
	Aggregations map[AggregateName]Aggregate `json:"aggregations,omitempty"`
	DocCount     int64                       `json:"doc_count"`
}

// AdjacencyMatrixBucketBuilder holds AdjacencyMatrixBucket struct and provides a builder API.
type AdjacencyMatrixBucketBuilder struct {
	v *AdjacencyMatrixBucket
}

// NewAdjacencyMatrixBucket provides a builder for the AdjacencyMatrixBucket struct.
func NewAdjacencyMatrixBucketBuilder() *AdjacencyMatrixBucketBuilder {
	r := AdjacencyMatrixBucketBuilder{
		&AdjacencyMatrixBucket{
			Aggregations: make(map[AggregateName]Aggregate, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the AdjacencyMatrixBucket struct
func (rb *AdjacencyMatrixBucketBuilder) Build() AdjacencyMatrixBucket {
	return *rb.v
}

func (rb *AdjacencyMatrixBucketBuilder) Aggregations(values map[AggregateName]*AggregateBuilder) *AdjacencyMatrixBucketBuilder {
	tmp := make(map[AggregateName]Aggregate, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Aggregations = tmp
	return rb
}

func (rb *AdjacencyMatrixBucketBuilder) DocCount(doccount int64) *AdjacencyMatrixBucketBuilder {
	rb.v.DocCount = doccount
	return rb
}
