package security

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// DeviceSecurityGroupsClient is the API spec for Microsoft.Security (Azure Security Center) resource provider
type DeviceSecurityGroupsClient struct {
	BaseClient
}

// NewDeviceSecurityGroupsClient creates an instance of the DeviceSecurityGroupsClient client.
func NewDeviceSecurityGroupsClient(subscriptionID string, ascLocation string) DeviceSecurityGroupsClient {
	return NewDeviceSecurityGroupsClientWithBaseURI(DefaultBaseURI, subscriptionID, ascLocation)
}

// NewDeviceSecurityGroupsClientWithBaseURI creates an instance of the DeviceSecurityGroupsClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewDeviceSecurityGroupsClientWithBaseURI(baseURI string, subscriptionID string, ascLocation string) DeviceSecurityGroupsClient {
	return DeviceSecurityGroupsClient{NewWithBaseURI(baseURI, subscriptionID, ascLocation)}
}

// CreateOrUpdate creates or updates the device security group on a specified IoT hub resource.
// Parameters:
// resourceID - the identifier of the resource.
// deviceSecurityGroupName - the name of the security group. Please notice that the name is case insensitive.
// deviceSecurityGroup - security group object.
func (client DeviceSecurityGroupsClient) CreateOrUpdate(ctx context.Context, resourceID string, deviceSecurityGroupName string, deviceSecurityGroup DeviceSecurityGroup) (result DeviceSecurityGroup, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DeviceSecurityGroupsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, resourceID, deviceSecurityGroupName, deviceSecurityGroup)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.DeviceSecurityGroupsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "security.DeviceSecurityGroupsClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.DeviceSecurityGroupsClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client DeviceSecurityGroupsClient) CreateOrUpdatePreparer(ctx context.Context, resourceID string, deviceSecurityGroupName string, deviceSecurityGroup DeviceSecurityGroup) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"deviceSecurityGroupName": autorest.Encode("path", deviceSecurityGroupName),
		"resourceId":              resourceID,
	}

	const APIVersion = "2017-08-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{resourceId}/providers/Microsoft.Security/deviceSecurityGroups/{deviceSecurityGroupName}", pathParameters),
		autorest.WithJSON(deviceSecurityGroup),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client DeviceSecurityGroupsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client DeviceSecurityGroupsClient) CreateOrUpdateResponder(resp *http.Response) (result DeviceSecurityGroup, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the security group
// Parameters:
// resourceID - the identifier of the resource.
// deviceSecurityGroupName - the name of the security group. Please notice that the name is case insensitive.
func (client DeviceSecurityGroupsClient) Delete(ctx context.Context, resourceID string, deviceSecurityGroupName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DeviceSecurityGroupsClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceID, deviceSecurityGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.DeviceSecurityGroupsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "security.DeviceSecurityGroupsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.DeviceSecurityGroupsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client DeviceSecurityGroupsClient) DeletePreparer(ctx context.Context, resourceID string, deviceSecurityGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"deviceSecurityGroupName": autorest.Encode("path", deviceSecurityGroupName),
		"resourceId":              resourceID,
	}

	const APIVersion = "2017-08-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{resourceId}/providers/Microsoft.Security/deviceSecurityGroups/{deviceSecurityGroupName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client DeviceSecurityGroupsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client DeviceSecurityGroupsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the device security group for the specified IoT hub resource.
// Parameters:
// resourceID - the identifier of the resource.
// deviceSecurityGroupName - the name of the security group. Please notice that the name is case insensitive.
func (client DeviceSecurityGroupsClient) Get(ctx context.Context, resourceID string, deviceSecurityGroupName string) (result DeviceSecurityGroup, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DeviceSecurityGroupsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceID, deviceSecurityGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.DeviceSecurityGroupsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "security.DeviceSecurityGroupsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.DeviceSecurityGroupsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client DeviceSecurityGroupsClient) GetPreparer(ctx context.Context, resourceID string, deviceSecurityGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"deviceSecurityGroupName": autorest.Encode("path", deviceSecurityGroupName),
		"resourceId":              resourceID,
	}

	const APIVersion = "2017-08-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{resourceId}/providers/Microsoft.Security/deviceSecurityGroups/{deviceSecurityGroupName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client DeviceSecurityGroupsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client DeviceSecurityGroupsClient) GetResponder(resp *http.Response) (result DeviceSecurityGroup, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets the list of device security groups for the specified IoT hub resource.
// Parameters:
// resourceID - the identifier of the resource.
func (client DeviceSecurityGroupsClient) List(ctx context.Context, resourceID string) (result DeviceSecurityGroupListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DeviceSecurityGroupsClient.List")
		defer func() {
			sc := -1
			if result.dsgl.Response.Response != nil {
				sc = result.dsgl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.DeviceSecurityGroupsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.dsgl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "security.DeviceSecurityGroupsClient", "List", resp, "Failure sending request")
		return
	}

	result.dsgl, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.DeviceSecurityGroupsClient", "List", resp, "Failure responding to request")
	}
	if result.dsgl.hasNextLink() && result.dsgl.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListPreparer prepares the List request.
func (client DeviceSecurityGroupsClient) ListPreparer(ctx context.Context, resourceID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceId": resourceID,
	}

	const APIVersion = "2017-08-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{resourceId}/providers/Microsoft.Security/deviceSecurityGroups", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client DeviceSecurityGroupsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client DeviceSecurityGroupsClient) ListResponder(resp *http.Response) (result DeviceSecurityGroupList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client DeviceSecurityGroupsClient) listNextResults(ctx context.Context, lastResults DeviceSecurityGroupList) (result DeviceSecurityGroupList, err error) {
	req, err := lastResults.deviceSecurityGroupListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "security.DeviceSecurityGroupsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "security.DeviceSecurityGroupsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.DeviceSecurityGroupsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client DeviceSecurityGroupsClient) ListComplete(ctx context.Context, resourceID string) (result DeviceSecurityGroupListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DeviceSecurityGroupsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceID)
	return
}
