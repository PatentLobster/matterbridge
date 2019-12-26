// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// MigrationStatus undocumented
type MigrationStatus int

const (
	// MigrationStatusVReady undocumented
	MigrationStatusVReady MigrationStatus = 0
	// MigrationStatusVNeedsReview undocumented
	MigrationStatusVNeedsReview MigrationStatus = 1
	// MigrationStatusVAdditionalStepsRequired undocumented
	MigrationStatusVAdditionalStepsRequired MigrationStatus = 2
	// MigrationStatusVUnknownFutureValue undocumented
	MigrationStatusVUnknownFutureValue MigrationStatus = 3
)

// MigrationStatusPReady returns a pointer to MigrationStatusVReady
func MigrationStatusPReady() *MigrationStatus {
	v := MigrationStatusVReady
	return &v
}

// MigrationStatusPNeedsReview returns a pointer to MigrationStatusVNeedsReview
func MigrationStatusPNeedsReview() *MigrationStatus {
	v := MigrationStatusVNeedsReview
	return &v
}

// MigrationStatusPAdditionalStepsRequired returns a pointer to MigrationStatusVAdditionalStepsRequired
func MigrationStatusPAdditionalStepsRequired() *MigrationStatus {
	v := MigrationStatusVAdditionalStepsRequired
	return &v
}

// MigrationStatusPUnknownFutureValue returns a pointer to MigrationStatusVUnknownFutureValue
func MigrationStatusPUnknownFutureValue() *MigrationStatus {
	v := MigrationStatusVUnknownFutureValue
	return &v
}