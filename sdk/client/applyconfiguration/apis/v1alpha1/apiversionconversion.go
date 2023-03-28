/*
Copyright The KCP Authors.

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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

// APIVersionConversionApplyConfiguration represents an declarative configuration of the APIVersionConversion type for use
// with apply.
type APIVersionConversionApplyConfiguration struct {
	From     *string                               `json:"from,omitempty"`
	To       *string                               `json:"to,omitempty"`
	Rules    []APIConversionRuleApplyConfiguration `json:"rules,omitempty"`
	Preserve []string                              `json:"preserve,omitempty"`
}

// APIVersionConversionApplyConfiguration constructs an declarative configuration of the APIVersionConversion type for use with
// apply.
func APIVersionConversion() *APIVersionConversionApplyConfiguration {
	return &APIVersionConversionApplyConfiguration{}
}

// WithFrom sets the From field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the From field is set to the value of the last call.
func (b *APIVersionConversionApplyConfiguration) WithFrom(value string) *APIVersionConversionApplyConfiguration {
	b.From = &value
	return b
}

// WithTo sets the To field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the To field is set to the value of the last call.
func (b *APIVersionConversionApplyConfiguration) WithTo(value string) *APIVersionConversionApplyConfiguration {
	b.To = &value
	return b
}

// WithRules adds the given value to the Rules field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Rules field.
func (b *APIVersionConversionApplyConfiguration) WithRules(values ...*APIConversionRuleApplyConfiguration) *APIVersionConversionApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithRules")
		}
		b.Rules = append(b.Rules, *values[i])
	}
	return b
}

// WithPreserve adds the given value to the Preserve field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Preserve field.
func (b *APIVersionConversionApplyConfiguration) WithPreserve(values ...string) *APIVersionConversionApplyConfiguration {
	for i := range values {
		b.Preserve = append(b.Preserve, values[i])
	}
	return b
}
