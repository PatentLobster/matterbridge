// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// MacManagedAppProtectionCollectionHasPayloadLinksRequestParameter undocumented
type MacManagedAppProtectionCollectionHasPayloadLinksRequestParameter struct {
	// PayloadIDs undocumented
	PayloadIDs []string `json:"payloadIds,omitempty"`
}

// SingleSignOnExtensionPkinitCertificate is navigation property
func (b *MacOSDeviceFeaturesConfigurationRequestBuilder) SingleSignOnExtensionPkinitCertificate() *MacOSCertificateProfileBaseRequestBuilder {
	bb := &MacOSCertificateProfileBaseRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/singleSignOnExtensionPkinitCertificate"
	return bb
}

// IdentityCertificateForClientAuthentication is navigation property
func (b *MacOSEnterpriseWiFiConfigurationRequestBuilder) IdentityCertificateForClientAuthentication() *MacOSCertificateProfileBaseRequestBuilder {
	bb := &MacOSCertificateProfileBaseRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/identityCertificateForClientAuthentication"
	return bb
}

// RootCertificateForServerValidation is navigation property
func (b *MacOSEnterpriseWiFiConfigurationRequestBuilder) RootCertificateForServerValidation() *MacOSTrustedRootCertificateRequestBuilder {
	bb := &MacOSTrustedRootCertificateRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/rootCertificateForServerValidation"
	return bb
}

// ManagedDeviceCertificateStates returns request builder for ManagedDeviceCertificateState collection
func (b *MacOSImportedPFXCertificateProfileRequestBuilder) ManagedDeviceCertificateStates() *MacOSImportedPFXCertificateProfileManagedDeviceCertificateStatesCollectionRequestBuilder {
	bb := &MacOSImportedPFXCertificateProfileManagedDeviceCertificateStatesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/managedDeviceCertificateStates"
	return bb
}

// MacOSImportedPFXCertificateProfileManagedDeviceCertificateStatesCollectionRequestBuilder is request builder for ManagedDeviceCertificateState collection
type MacOSImportedPFXCertificateProfileManagedDeviceCertificateStatesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ManagedDeviceCertificateState collection
func (b *MacOSImportedPFXCertificateProfileManagedDeviceCertificateStatesCollectionRequestBuilder) Request() *MacOSImportedPFXCertificateProfileManagedDeviceCertificateStatesCollectionRequest {
	return &MacOSImportedPFXCertificateProfileManagedDeviceCertificateStatesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ManagedDeviceCertificateState item
func (b *MacOSImportedPFXCertificateProfileManagedDeviceCertificateStatesCollectionRequestBuilder) ID(id string) *ManagedDeviceCertificateStateRequestBuilder {
	bb := &ManagedDeviceCertificateStateRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// MacOSImportedPFXCertificateProfileManagedDeviceCertificateStatesCollectionRequest is request for ManagedDeviceCertificateState collection
type MacOSImportedPFXCertificateProfileManagedDeviceCertificateStatesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ManagedDeviceCertificateState collection
func (r *MacOSImportedPFXCertificateProfileManagedDeviceCertificateStatesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]ManagedDeviceCertificateState, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []ManagedDeviceCertificateState
	for {
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			res.Body.Close()
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []ManagedDeviceCertificateState
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		res.Body.Close()
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if n >= 0 {
			n--
		}
		if n == 0 || len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// GetN performs GET request for ManagedDeviceCertificateState collection, max N pages
func (r *MacOSImportedPFXCertificateProfileManagedDeviceCertificateStatesCollectionRequest) GetN(ctx context.Context, n int) ([]ManagedDeviceCertificateState, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for ManagedDeviceCertificateState collection
func (r *MacOSImportedPFXCertificateProfileManagedDeviceCertificateStatesCollectionRequest) Get(ctx context.Context) ([]ManagedDeviceCertificateState, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for ManagedDeviceCertificateState collection
func (r *MacOSImportedPFXCertificateProfileManagedDeviceCertificateStatesCollectionRequest) Add(ctx context.Context, reqObj *ManagedDeviceCertificateState) (resObj *ManagedDeviceCertificateState, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// ManagedDeviceCertificateStates returns request builder for ManagedDeviceCertificateState collection
func (b *MacOSPkcsCertificateProfileRequestBuilder) ManagedDeviceCertificateStates() *MacOSPkcsCertificateProfileManagedDeviceCertificateStatesCollectionRequestBuilder {
	bb := &MacOSPkcsCertificateProfileManagedDeviceCertificateStatesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/managedDeviceCertificateStates"
	return bb
}

// MacOSPkcsCertificateProfileManagedDeviceCertificateStatesCollectionRequestBuilder is request builder for ManagedDeviceCertificateState collection
type MacOSPkcsCertificateProfileManagedDeviceCertificateStatesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ManagedDeviceCertificateState collection
func (b *MacOSPkcsCertificateProfileManagedDeviceCertificateStatesCollectionRequestBuilder) Request() *MacOSPkcsCertificateProfileManagedDeviceCertificateStatesCollectionRequest {
	return &MacOSPkcsCertificateProfileManagedDeviceCertificateStatesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ManagedDeviceCertificateState item
func (b *MacOSPkcsCertificateProfileManagedDeviceCertificateStatesCollectionRequestBuilder) ID(id string) *ManagedDeviceCertificateStateRequestBuilder {
	bb := &ManagedDeviceCertificateStateRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// MacOSPkcsCertificateProfileManagedDeviceCertificateStatesCollectionRequest is request for ManagedDeviceCertificateState collection
type MacOSPkcsCertificateProfileManagedDeviceCertificateStatesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ManagedDeviceCertificateState collection
func (r *MacOSPkcsCertificateProfileManagedDeviceCertificateStatesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]ManagedDeviceCertificateState, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []ManagedDeviceCertificateState
	for {
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			res.Body.Close()
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []ManagedDeviceCertificateState
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		res.Body.Close()
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if n >= 0 {
			n--
		}
		if n == 0 || len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// GetN performs GET request for ManagedDeviceCertificateState collection, max N pages
func (r *MacOSPkcsCertificateProfileManagedDeviceCertificateStatesCollectionRequest) GetN(ctx context.Context, n int) ([]ManagedDeviceCertificateState, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for ManagedDeviceCertificateState collection
func (r *MacOSPkcsCertificateProfileManagedDeviceCertificateStatesCollectionRequest) Get(ctx context.Context) ([]ManagedDeviceCertificateState, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for ManagedDeviceCertificateState collection
func (r *MacOSPkcsCertificateProfileManagedDeviceCertificateStatesCollectionRequest) Add(ctx context.Context, reqObj *ManagedDeviceCertificateState) (resObj *ManagedDeviceCertificateState, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// ManagedDeviceCertificateStates returns request builder for ManagedDeviceCertificateState collection
func (b *MacOSScepCertificateProfileRequestBuilder) ManagedDeviceCertificateStates() *MacOSScepCertificateProfileManagedDeviceCertificateStatesCollectionRequestBuilder {
	bb := &MacOSScepCertificateProfileManagedDeviceCertificateStatesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/managedDeviceCertificateStates"
	return bb
}

// MacOSScepCertificateProfileManagedDeviceCertificateStatesCollectionRequestBuilder is request builder for ManagedDeviceCertificateState collection
type MacOSScepCertificateProfileManagedDeviceCertificateStatesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ManagedDeviceCertificateState collection
func (b *MacOSScepCertificateProfileManagedDeviceCertificateStatesCollectionRequestBuilder) Request() *MacOSScepCertificateProfileManagedDeviceCertificateStatesCollectionRequest {
	return &MacOSScepCertificateProfileManagedDeviceCertificateStatesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ManagedDeviceCertificateState item
func (b *MacOSScepCertificateProfileManagedDeviceCertificateStatesCollectionRequestBuilder) ID(id string) *ManagedDeviceCertificateStateRequestBuilder {
	bb := &ManagedDeviceCertificateStateRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// MacOSScepCertificateProfileManagedDeviceCertificateStatesCollectionRequest is request for ManagedDeviceCertificateState collection
type MacOSScepCertificateProfileManagedDeviceCertificateStatesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ManagedDeviceCertificateState collection
func (r *MacOSScepCertificateProfileManagedDeviceCertificateStatesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]ManagedDeviceCertificateState, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []ManagedDeviceCertificateState
	for {
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			res.Body.Close()
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []ManagedDeviceCertificateState
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		res.Body.Close()
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if n >= 0 {
			n--
		}
		if n == 0 || len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// GetN performs GET request for ManagedDeviceCertificateState collection, max N pages
func (r *MacOSScepCertificateProfileManagedDeviceCertificateStatesCollectionRequest) GetN(ctx context.Context, n int) ([]ManagedDeviceCertificateState, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for ManagedDeviceCertificateState collection
func (r *MacOSScepCertificateProfileManagedDeviceCertificateStatesCollectionRequest) Get(ctx context.Context) ([]ManagedDeviceCertificateState, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for ManagedDeviceCertificateState collection
func (r *MacOSScepCertificateProfileManagedDeviceCertificateStatesCollectionRequest) Add(ctx context.Context, reqObj *ManagedDeviceCertificateState) (resObj *ManagedDeviceCertificateState, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// RootCertificate is navigation property
func (b *MacOSScepCertificateProfileRequestBuilder) RootCertificate() *MacOSTrustedRootCertificateRequestBuilder {
	bb := &MacOSTrustedRootCertificateRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/rootCertificate"
	return bb
}

// IdentityCertificate is navigation property
func (b *MacOSVpnConfigurationRequestBuilder) IdentityCertificate() *MacOSCertificateProfileBaseRequestBuilder {
	bb := &MacOSCertificateProfileBaseRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/identityCertificate"
	return bb
}

// IdentityCertificateForClientAuthentication is navigation property
func (b *MacOSWiredNetworkConfigurationRequestBuilder) IdentityCertificateForClientAuthentication() *MacOSCertificateProfileBaseRequestBuilder {
	bb := &MacOSCertificateProfileBaseRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/identityCertificateForClientAuthentication"
	return bb
}

// RootCertificateForServerValidation is navigation property
func (b *MacOSWiredNetworkConfigurationRequestBuilder) RootCertificateForServerValidation() *MacOSTrustedRootCertificateRequestBuilder {
	bb := &MacOSTrustedRootCertificateRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/rootCertificateForServerValidation"
	return bb
}

// AssignedLicenses returns request builder for MacOsVPPAppAssignedLicense collection
func (b *MacOsVPPAppRequestBuilder) AssignedLicenses() *MacOsVPPAppAssignedLicensesCollectionRequestBuilder {
	bb := &MacOsVPPAppAssignedLicensesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/assignedLicenses"
	return bb
}

// MacOsVPPAppAssignedLicensesCollectionRequestBuilder is request builder for MacOsVPPAppAssignedLicense collection
type MacOsVPPAppAssignedLicensesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for MacOsVPPAppAssignedLicense collection
func (b *MacOsVPPAppAssignedLicensesCollectionRequestBuilder) Request() *MacOsVPPAppAssignedLicensesCollectionRequest {
	return &MacOsVPPAppAssignedLicensesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for MacOsVPPAppAssignedLicense item
func (b *MacOsVPPAppAssignedLicensesCollectionRequestBuilder) ID(id string) *MacOsVPPAppAssignedLicenseRequestBuilder {
	bb := &MacOsVPPAppAssignedLicenseRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// MacOsVPPAppAssignedLicensesCollectionRequest is request for MacOsVPPAppAssignedLicense collection
type MacOsVPPAppAssignedLicensesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for MacOsVPPAppAssignedLicense collection
func (r *MacOsVPPAppAssignedLicensesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]MacOsVPPAppAssignedLicense, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []MacOsVPPAppAssignedLicense
	for {
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			res.Body.Close()
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []MacOsVPPAppAssignedLicense
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		res.Body.Close()
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if n >= 0 {
			n--
		}
		if n == 0 || len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// GetN performs GET request for MacOsVPPAppAssignedLicense collection, max N pages
func (r *MacOsVPPAppAssignedLicensesCollectionRequest) GetN(ctx context.Context, n int) ([]MacOsVPPAppAssignedLicense, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for MacOsVPPAppAssignedLicense collection
func (r *MacOsVPPAppAssignedLicensesCollectionRequest) Get(ctx context.Context) ([]MacOsVPPAppAssignedLicense, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for MacOsVPPAppAssignedLicense collection
func (r *MacOsVPPAppAssignedLicensesCollectionRequest) Add(ctx context.Context, reqObj *MacOsVPPAppAssignedLicense) (resObj *MacOsVPPAppAssignedLicense, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}
