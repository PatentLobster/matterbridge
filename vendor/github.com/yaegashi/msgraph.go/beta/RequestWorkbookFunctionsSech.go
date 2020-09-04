// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

//
type WorkbookFunctionsSechRequestBuilder struct{ BaseRequestBuilder }

// Sech action undocumented
func (b *WorkbookFunctionsRequestBuilder) Sech(reqObj *WorkbookFunctionsSechRequestParameter) *WorkbookFunctionsSechRequestBuilder {
	bb := &WorkbookFunctionsSechRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/sech"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsSechRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsSechRequestBuilder) Request() *WorkbookFunctionsSechRequest {
	return &WorkbookFunctionsSechRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFunctionsSechRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}
