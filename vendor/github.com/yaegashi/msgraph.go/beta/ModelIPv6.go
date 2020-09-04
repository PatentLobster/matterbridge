// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// IPv6CidrRange undocumented
type IPv6CidrRange struct {
	// IPRange is the base model of IPv6CidrRange
	IPRange
	// CIDRAddress undocumented
	CIDRAddress *string `json:"cidrAddress,omitempty"`
}

// IPv6Range undocumented
type IPv6Range struct {
	// IPRange is the base model of IPv6Range
	IPRange
	// LowerAddress Lower address.
	LowerAddress *string `json:"lowerAddress,omitempty"`
	// UpperAddress Upper address.
	UpperAddress *string `json:"upperAddress,omitempty"`
}
