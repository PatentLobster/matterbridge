// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// Attachment undocumented
type Attachment struct {
	// Entity is the base model of Attachment
	Entity
	// LastModifiedDateTime undocumented
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// Name undocumented
	Name *string `json:"name,omitempty"`
	// ContentType undocumented
	ContentType *string `json:"contentType,omitempty"`
	// Size undocumented
	Size *int `json:"size,omitempty"`
	// IsInline undocumented
	IsInline *bool `json:"isInline,omitempty"`
}

// AttachmentItem undocumented
type AttachmentItem struct {
	// Object is the base model of AttachmentItem
	Object
	// AttachmentType undocumented
	AttachmentType *AttachmentType `json:"attachmentType,omitempty"`
	// Name undocumented
	Name *string `json:"name,omitempty"`
	// Size undocumented
	Size *int `json:"size,omitempty"`
	// ContentType undocumented
	ContentType *string `json:"contentType,omitempty"`
	// IsInline undocumented
	IsInline *bool `json:"isInline,omitempty"`
}
