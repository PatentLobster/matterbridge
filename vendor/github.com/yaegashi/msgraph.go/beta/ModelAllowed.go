// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// AllowedDataLocation undocumented
type AllowedDataLocation struct {
	// Entity is the base model of AllowedDataLocation
	Entity
	// AppID undocumented
	AppID *string `json:"appId,omitempty"`
	// Location undocumented
	Location *string `json:"location,omitempty"`
	// IsDefault undocumented
	IsDefault *bool `json:"isDefault,omitempty"`
	// Domain undocumented
	Domain *string `json:"domain,omitempty"`
}

// AllowedDataLocationInfo undocumented
type AllowedDataLocationInfo struct {
	// Object is the base model of AllowedDataLocationInfo
	Object
}
