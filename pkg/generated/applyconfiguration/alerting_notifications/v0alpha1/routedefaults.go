// SPDX-License-Identifier: AGPL-3.0-only

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v0alpha1

// RouteDefaultsApplyConfiguration represents a declarative configuration of the RouteDefaults type for use
// with apply.
type RouteDefaultsApplyConfiguration struct {
	GroupBy        []string `json:"group_by,omitempty"`
	GroupInterval  *string  `json:"group_interval,omitempty"`
	GroupWait      *string  `json:"group_wait,omitempty"`
	Receiver       *string  `json:"receiver,omitempty"`
	RepeatInterval *string  `json:"repeat_interval,omitempty"`
}

// RouteDefaultsApplyConfiguration constructs a declarative configuration of the RouteDefaults type for use with
// apply.
func RouteDefaults() *RouteDefaultsApplyConfiguration {
	return &RouteDefaultsApplyConfiguration{}
}

// WithGroupBy adds the given value to the GroupBy field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the GroupBy field.
func (b *RouteDefaultsApplyConfiguration) WithGroupBy(values ...string) *RouteDefaultsApplyConfiguration {
	for i := range values {
		b.GroupBy = append(b.GroupBy, values[i])
	}
	return b
}

// WithGroupInterval sets the GroupInterval field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GroupInterval field is set to the value of the last call.
func (b *RouteDefaultsApplyConfiguration) WithGroupInterval(value string) *RouteDefaultsApplyConfiguration {
	b.GroupInterval = &value
	return b
}

// WithGroupWait sets the GroupWait field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GroupWait field is set to the value of the last call.
func (b *RouteDefaultsApplyConfiguration) WithGroupWait(value string) *RouteDefaultsApplyConfiguration {
	b.GroupWait = &value
	return b
}

// WithReceiver sets the Receiver field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Receiver field is set to the value of the last call.
func (b *RouteDefaultsApplyConfiguration) WithReceiver(value string) *RouteDefaultsApplyConfiguration {
	b.Receiver = &value
	return b
}

// WithRepeatInterval sets the RepeatInterval field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RepeatInterval field is set to the value of the last call.
func (b *RouteDefaultsApplyConfiguration) WithRepeatInterval(value string) *RouteDefaultsApplyConfiguration {
	b.RepeatInterval = &value
	return b
}