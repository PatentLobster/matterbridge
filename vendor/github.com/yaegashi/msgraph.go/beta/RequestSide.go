// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// SideLoadingKeyRequestBuilder is request builder for SideLoadingKey
type SideLoadingKeyRequestBuilder struct{ BaseRequestBuilder }

// Request returns SideLoadingKeyRequest
func (b *SideLoadingKeyRequestBuilder) Request() *SideLoadingKeyRequest {
	return &SideLoadingKeyRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SideLoadingKeyRequest is request for SideLoadingKey
type SideLoadingKeyRequest struct{ BaseRequest }

// Get performs GET request for SideLoadingKey
func (r *SideLoadingKeyRequest) Get(ctx context.Context) (resObj *SideLoadingKey, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SideLoadingKey
func (r *SideLoadingKeyRequest) Update(ctx context.Context, reqObj *SideLoadingKey) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SideLoadingKey
func (r *SideLoadingKeyRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
