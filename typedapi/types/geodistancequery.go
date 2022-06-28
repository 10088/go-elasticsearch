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

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/geodistancetype"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/geovalidationmethod"
)

// GeoDistanceQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/135ae054e304239743b5777ad8d41cb2c9091d35/specification/_types/query_dsl/geo.ts#L48-L57
type GeoDistanceQuery struct {
	Boost            *float32                                 `json:"boost,omitempty"`
	Distance         *Distance                                `json:"distance,omitempty"`
	DistanceType     *geodistancetype.GeoDistanceType         `json:"distance_type,omitempty"`
	GeoDistanceQuery map[Field]GeoLocation                    `json:"GeoDistanceQuery,omitempty"`
	QueryName_       *string                                  `json:"_name,omitempty"`
	ValidationMethod *geovalidationmethod.GeoValidationMethod `json:"validation_method,omitempty"`
}

// GeoDistanceQueryBuilder holds GeoDistanceQuery struct and provides a builder API.
type GeoDistanceQueryBuilder struct {
	v *GeoDistanceQuery
}

// NewGeoDistanceQuery provides a builder for the GeoDistanceQuery struct.
func NewGeoDistanceQueryBuilder() *GeoDistanceQueryBuilder {
	r := GeoDistanceQueryBuilder{
		&GeoDistanceQuery{
			GeoDistanceQuery: make(map[Field]GeoLocation, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the GeoDistanceQuery struct
func (rb *GeoDistanceQueryBuilder) Build() GeoDistanceQuery {
	return *rb.v
}

func (rb *GeoDistanceQueryBuilder) Boost(boost float32) *GeoDistanceQueryBuilder {
	rb.v.Boost = &boost
	return rb
}

func (rb *GeoDistanceQueryBuilder) Distance(distance Distance) *GeoDistanceQueryBuilder {
	rb.v.Distance = &distance
	return rb
}

func (rb *GeoDistanceQueryBuilder) DistanceType(distancetype geodistancetype.GeoDistanceType) *GeoDistanceQueryBuilder {
	rb.v.DistanceType = &distancetype
	return rb
}

func (rb *GeoDistanceQueryBuilder) GeoDistanceQuery(values map[Field]*GeoLocationBuilder) *GeoDistanceQueryBuilder {
	tmp := make(map[Field]GeoLocation, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.GeoDistanceQuery = tmp
	return rb
}

func (rb *GeoDistanceQueryBuilder) QueryName_(queryname_ string) *GeoDistanceQueryBuilder {
	rb.v.QueryName_ = &queryname_
	return rb
}

func (rb *GeoDistanceQueryBuilder) ValidationMethod(validationmethod geovalidationmethod.GeoValidationMethod) *GeoDistanceQueryBuilder {
	rb.v.ValidationMethod = &validationmethod
	return rb
}
