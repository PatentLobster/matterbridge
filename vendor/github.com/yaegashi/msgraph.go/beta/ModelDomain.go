// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// Domain undocumented
type Domain struct {
	// Entity is the base model of Domain
	Entity
	// AuthenticationType undocumented
	AuthenticationType *string `json:"authenticationType,omitempty"`
	// AvailabilityStatus undocumented
	AvailabilityStatus *string `json:"availabilityStatus,omitempty"`
	// IsAdminManaged undocumented
	IsAdminManaged *bool `json:"isAdminManaged,omitempty"`
	// IsDefault undocumented
	IsDefault *bool `json:"isDefault,omitempty"`
	// IsInitial undocumented
	IsInitial *bool `json:"isInitial,omitempty"`
	// IsRoot undocumented
	IsRoot *bool `json:"isRoot,omitempty"`
	// IsVerified undocumented
	IsVerified *bool `json:"isVerified,omitempty"`
	// PasswordNotificationWindowInDays undocumented
	PasswordNotificationWindowInDays *int `json:"passwordNotificationWindowInDays,omitempty"`
	// PasswordValidityPeriodInDays undocumented
	PasswordValidityPeriodInDays *int `json:"passwordValidityPeriodInDays,omitempty"`
	// SupportedServices undocumented
	SupportedServices []string `json:"supportedServices,omitempty"`
	// State undocumented
	State *DomainState `json:"state,omitempty"`
	// ServiceConfigurationRecords undocumented
	ServiceConfigurationRecords []DomainDNSRecord `json:"serviceConfigurationRecords,omitempty"`
	// VerificationDNSRecords undocumented
	VerificationDNSRecords []DomainDNSRecord `json:"verificationDnsRecords,omitempty"`
	// DomainNameReferences undocumented
	DomainNameReferences []DirectoryObject `json:"domainNameReferences,omitempty"`
}

// DomainDNSCnameRecord undocumented
type DomainDNSCnameRecord struct {
	// DomainDNSRecord is the base model of DomainDNSCnameRecord
	DomainDNSRecord
	// CanonicalName undocumented
	CanonicalName *string `json:"canonicalName,omitempty"`
}

// DomainDNSMxRecord undocumented
type DomainDNSMxRecord struct {
	// DomainDNSRecord is the base model of DomainDNSMxRecord
	DomainDNSRecord
	// MailExchange undocumented
	MailExchange *string `json:"mailExchange,omitempty"`
	// Preference undocumented
	Preference *int `json:"preference,omitempty"`
}

// DomainDNSRecord undocumented
type DomainDNSRecord struct {
	// Entity is the base model of DomainDNSRecord
	Entity
	// IsOptional undocumented
	IsOptional *bool `json:"isOptional,omitempty"`
	// Label undocumented
	Label *string `json:"label,omitempty"`
	// RecordType undocumented
	RecordType *string `json:"recordType,omitempty"`
	// SupportedService undocumented
	SupportedService *string `json:"supportedService,omitempty"`
	// TTL undocumented
	TTL *int `json:"ttl,omitempty"`
}

// DomainDNSSrvRecord undocumented
type DomainDNSSrvRecord struct {
	// DomainDNSRecord is the base model of DomainDNSSrvRecord
	DomainDNSRecord
	// NameTarget undocumented
	NameTarget *string `json:"nameTarget,omitempty"`
	// Port undocumented
	Port *int `json:"port,omitempty"`
	// Priority undocumented
	Priority *int `json:"priority,omitempty"`
	// Protocol undocumented
	Protocol *string `json:"protocol,omitempty"`
	// Service undocumented
	Service *string `json:"service,omitempty"`
	// Weight undocumented
	Weight *int `json:"weight,omitempty"`
}

// DomainDNSTxtRecord undocumented
type DomainDNSTxtRecord struct {
	// DomainDNSRecord is the base model of DomainDNSTxtRecord
	DomainDNSRecord
	// Text undocumented
	Text *string `json:"text,omitempty"`
}

// DomainDNSUnavailableRecord undocumented
type DomainDNSUnavailableRecord struct {
	// DomainDNSRecord is the base model of DomainDNSUnavailableRecord
	DomainDNSRecord
	// Description undocumented
	Description *string `json:"description,omitempty"`
}

// DomainRegistrant undocumented
type DomainRegistrant struct {
	// Object is the base model of DomainRegistrant
	Object
	// CountryOrRegionCode undocumented
	CountryOrRegionCode *string `json:"countryOrRegionCode,omitempty"`
	// Organization undocumented
	Organization *string `json:"organization,omitempty"`
	// URL undocumented
	URL *string `json:"url,omitempty"`
	// Vendor undocumented
	Vendor *string `json:"vendor,omitempty"`
}

// DomainSecurityProfile undocumented
type DomainSecurityProfile struct {
	// Entity is the base model of DomainSecurityProfile
	Entity
	// ActivityGroupNames undocumented
	ActivityGroupNames []string `json:"activityGroupNames,omitempty"`
	// AzureSubscriptionID undocumented
	AzureSubscriptionID *string `json:"azureSubscriptionId,omitempty"`
	// AzureTenantID undocumented
	AzureTenantID *string `json:"azureTenantId,omitempty"`
	// CountHits undocumented
	CountHits *int `json:"countHits,omitempty"`
	// CountInOrg undocumented
	CountInOrg *int `json:"countInOrg,omitempty"`
	// DomainCategories undocumented
	DomainCategories []ReputationCategory `json:"domainCategories,omitempty"`
	// DomainRegisteredDateTime undocumented
	DomainRegisteredDateTime *time.Time `json:"domainRegisteredDateTime,omitempty"`
	// FirstSeenDateTime undocumented
	FirstSeenDateTime *time.Time `json:"firstSeenDateTime,omitempty"`
	// LastSeenDateTime undocumented
	LastSeenDateTime *time.Time `json:"lastSeenDateTime,omitempty"`
	// Name undocumented
	Name *string `json:"name,omitempty"`
	// Registrant undocumented
	Registrant *DomainRegistrant `json:"registrant,omitempty"`
	// RiskScore undocumented
	RiskScore *string `json:"riskScore,omitempty"`
	// Tags undocumented
	Tags []string `json:"tags,omitempty"`
	// VendorInformation undocumented
	VendorInformation *SecurityVendorInformation `json:"vendorInformation,omitempty"`
}

// DomainState undocumented
type DomainState struct {
	// Object is the base model of DomainState
	Object
	// Status undocumented
	Status *string `json:"status,omitempty"`
	// Operation undocumented
	Operation *string `json:"operation,omitempty"`
	// LastActionDateTime undocumented
	LastActionDateTime *time.Time `json:"lastActionDateTime,omitempty"`
}
