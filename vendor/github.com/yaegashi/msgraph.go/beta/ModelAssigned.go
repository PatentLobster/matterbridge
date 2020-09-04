// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// AssignedLabel undocumented
type AssignedLabel struct {
	// Object is the base model of AssignedLabel
	Object
	// LabelID undocumented
	LabelID *string `json:"labelId,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
}

// AssignedLicense undocumented
type AssignedLicense struct {
	// Object is the base model of AssignedLicense
	Object
	// DisabledPlans undocumented
	DisabledPlans []UUID `json:"disabledPlans,omitempty"`
	// SKUID undocumented
	SKUID *UUID `json:"skuId,omitempty"`
}

// AssignedPlan undocumented
type AssignedPlan struct {
	// Object is the base model of AssignedPlan
	Object
	// AssignedDateTime undocumented
	AssignedDateTime *time.Time `json:"assignedDateTime,omitempty"`
	// CapabilityStatus undocumented
	CapabilityStatus *string `json:"capabilityStatus,omitempty"`
	// Service undocumented
	Service *string `json:"service,omitempty"`
	// ServicePlanID undocumented
	ServicePlanID *UUID `json:"servicePlanId,omitempty"`
}
