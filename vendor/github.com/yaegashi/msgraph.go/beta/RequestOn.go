// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// OnPremisesAgentRequestBuilder is request builder for OnPremisesAgent
type OnPremisesAgentRequestBuilder struct{ BaseRequestBuilder }

// Request returns OnPremisesAgentRequest
func (b *OnPremisesAgentRequestBuilder) Request() *OnPremisesAgentRequest {
	return &OnPremisesAgentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// OnPremisesAgentRequest is request for OnPremisesAgent
type OnPremisesAgentRequest struct{ BaseRequest }

// Get performs GET request for OnPremisesAgent
func (r *OnPremisesAgentRequest) Get(ctx context.Context) (resObj *OnPremisesAgent, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for OnPremisesAgent
func (r *OnPremisesAgentRequest) Update(ctx context.Context, reqObj *OnPremisesAgent) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for OnPremisesAgent
func (r *OnPremisesAgentRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// OnPremisesAgentGroupRequestBuilder is request builder for OnPremisesAgentGroup
type OnPremisesAgentGroupRequestBuilder struct{ BaseRequestBuilder }

// Request returns OnPremisesAgentGroupRequest
func (b *OnPremisesAgentGroupRequestBuilder) Request() *OnPremisesAgentGroupRequest {
	return &OnPremisesAgentGroupRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// OnPremisesAgentGroupRequest is request for OnPremisesAgentGroup
type OnPremisesAgentGroupRequest struct{ BaseRequest }

// Get performs GET request for OnPremisesAgentGroup
func (r *OnPremisesAgentGroupRequest) Get(ctx context.Context) (resObj *OnPremisesAgentGroup, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for OnPremisesAgentGroup
func (r *OnPremisesAgentGroupRequest) Update(ctx context.Context, reqObj *OnPremisesAgentGroup) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for OnPremisesAgentGroup
func (r *OnPremisesAgentGroupRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// OnPremisesConditionalAccessSettingsRequestBuilder is request builder for OnPremisesConditionalAccessSettings
type OnPremisesConditionalAccessSettingsRequestBuilder struct{ BaseRequestBuilder }

// Request returns OnPremisesConditionalAccessSettingsRequest
func (b *OnPremisesConditionalAccessSettingsRequestBuilder) Request() *OnPremisesConditionalAccessSettingsRequest {
	return &OnPremisesConditionalAccessSettingsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// OnPremisesConditionalAccessSettingsRequest is request for OnPremisesConditionalAccessSettings
type OnPremisesConditionalAccessSettingsRequest struct{ BaseRequest }

// Get performs GET request for OnPremisesConditionalAccessSettings
func (r *OnPremisesConditionalAccessSettingsRequest) Get(ctx context.Context) (resObj *OnPremisesConditionalAccessSettings, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for OnPremisesConditionalAccessSettings
func (r *OnPremisesConditionalAccessSettingsRequest) Update(ctx context.Context, reqObj *OnPremisesConditionalAccessSettings) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for OnPremisesConditionalAccessSettings
func (r *OnPremisesConditionalAccessSettingsRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// OnPremisesPublishingProfileRequestBuilder is request builder for OnPremisesPublishingProfile
type OnPremisesPublishingProfileRequestBuilder struct{ BaseRequestBuilder }

// Request returns OnPremisesPublishingProfileRequest
func (b *OnPremisesPublishingProfileRequestBuilder) Request() *OnPremisesPublishingProfileRequest {
	return &OnPremisesPublishingProfileRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// OnPremisesPublishingProfileRequest is request for OnPremisesPublishingProfile
type OnPremisesPublishingProfileRequest struct{ BaseRequest }

// Get performs GET request for OnPremisesPublishingProfile
func (r *OnPremisesPublishingProfileRequest) Get(ctx context.Context) (resObj *OnPremisesPublishingProfile, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for OnPremisesPublishingProfile
func (r *OnPremisesPublishingProfileRequest) Update(ctx context.Context, reqObj *OnPremisesPublishingProfile) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for OnPremisesPublishingProfile
func (r *OnPremisesPublishingProfileRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
