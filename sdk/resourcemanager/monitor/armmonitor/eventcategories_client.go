//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmonitor

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// EventCategoriesClient contains the methods for the EventCategories group.
// Don't use this type directly, use NewEventCategoriesClient() instead.
type EventCategoriesClient struct {
	internal *arm.Client
}

// NewEventCategoriesClient creates a new instance of EventCategoriesClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewEventCategoriesClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*EventCategoriesClient, error) {
	cl, err := arm.NewClient(moduleName+".EventCategoriesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &EventCategoriesClient{
		internal: cl,
	}
	return client, nil
}

// NewListPager - Get the list of available event categories supported in the Activity Logs Service. The current list includes
// the following: Administrative, Security, ServiceHealth, Alert, Recommendation, Policy.
//
// Generated from API version 2015-04-01
//   - options - EventCategoriesClientListOptions contains the optional parameters for the EventCategoriesClient.NewListPager
//     method.
func (client *EventCategoriesClient) NewListPager(options *EventCategoriesClientListOptions) *runtime.Pager[EventCategoriesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[EventCategoriesClientListResponse]{
		More: func(page EventCategoriesClientListResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *EventCategoriesClientListResponse) (EventCategoriesClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "EventCategoriesClient.NewListPager")
			req, err := client.listCreateRequest(ctx, options)
			if err != nil {
				return EventCategoriesClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return EventCategoriesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return EventCategoriesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *EventCategoriesClient) listCreateRequest(ctx context.Context, options *EventCategoriesClientListOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Insights/eventcategories"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *EventCategoriesClient) listHandleResponse(resp *http.Response) (EventCategoriesClientListResponse, error) {
	result := EventCategoriesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EventCategoryCollection); err != nil {
		return EventCategoriesClientListResponse{}, err
	}
	return result, nil
}
