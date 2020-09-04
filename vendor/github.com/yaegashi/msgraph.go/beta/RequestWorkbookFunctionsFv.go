// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

//
type WorkbookFunctionsFvRequestBuilder struct{ BaseRequestBuilder }

// Fv action undocumented
func (b *WorkbookFunctionsRequestBuilder) Fv(reqObj *WorkbookFunctionsFvRequestParameter) *WorkbookFunctionsFvRequestBuilder {
	bb := &WorkbookFunctionsFvRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/fv"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsFvRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsFvRequestBuilder) Request() *WorkbookFunctionsFvRequest {
	return &WorkbookFunctionsFvRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFunctionsFvRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}
