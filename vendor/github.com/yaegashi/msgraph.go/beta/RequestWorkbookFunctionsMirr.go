// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

//
type WorkbookFunctionsMirrRequestBuilder struct{ BaseRequestBuilder }

// Mirr action undocumented
func (b *WorkbookFunctionsRequestBuilder) Mirr(reqObj *WorkbookFunctionsMirrRequestParameter) *WorkbookFunctionsMirrRequestBuilder {
	bb := &WorkbookFunctionsMirrRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/mirr"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsMirrRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsMirrRequestBuilder) Request() *WorkbookFunctionsMirrRequest {
	return &WorkbookFunctionsMirrRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFunctionsMirrRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}
