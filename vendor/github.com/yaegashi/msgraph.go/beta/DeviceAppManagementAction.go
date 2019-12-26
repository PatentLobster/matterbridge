// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// DeviceAppManagementSyncMicrosoftStoreForBusinessAppsRequestParameter undocumented
type DeviceAppManagementSyncMicrosoftStoreForBusinessAppsRequestParameter struct {
}

//
type DeviceAppManagementSyncMicrosoftStoreForBusinessAppsRequestBuilder struct{ BaseRequestBuilder }

// SyncMicrosoftStoreForBusinessApps action undocumented
func (b *DeviceAppManagementRequestBuilder) SyncMicrosoftStoreForBusinessApps(reqObj *DeviceAppManagementSyncMicrosoftStoreForBusinessAppsRequestParameter) *DeviceAppManagementSyncMicrosoftStoreForBusinessAppsRequestBuilder {
	bb := &DeviceAppManagementSyncMicrosoftStoreForBusinessAppsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/syncMicrosoftStoreForBusinessApps"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type DeviceAppManagementSyncMicrosoftStoreForBusinessAppsRequest struct{ BaseRequest }

//
func (b *DeviceAppManagementSyncMicrosoftStoreForBusinessAppsRequestBuilder) Request() *DeviceAppManagementSyncMicrosoftStoreForBusinessAppsRequest {
	return &DeviceAppManagementSyncMicrosoftStoreForBusinessAppsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *DeviceAppManagementSyncMicrosoftStoreForBusinessAppsRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}