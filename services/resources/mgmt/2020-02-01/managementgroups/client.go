// Package managementgroups implements the Azure ARM Managementgroups service API version 2020-02-01.
//
// The Azure Management Groups API enables consolidation of multiple
// subscriptions/resources into an organizational hierarchy and centrally
// manage access control, policies, alerting and reporting for those resources.
//
package managementgroups

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

const (
	// DefaultBaseURI is the default URI used for the service Managementgroups
	DefaultBaseURI = "https://management.azure.com"
)

// BaseClient is the base client for Managementgroups.
type BaseClient struct {
	autorest.Client
	BaseURI           string
	OperationResultID string
	Skip              *int32
	Top               *int32
	Skiptoken         string
}

// New creates an instance of the BaseClient client.
func New(operationResultID string, skip *int32, top *int32, skiptoken string) BaseClient {
	return NewWithBaseURI(DefaultBaseURI, operationResultID, skip, top, skiptoken)
}

// NewWithBaseURI creates an instance of the BaseClient client using a custom endpoint.  Use this when interacting with
// an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewWithBaseURI(baseURI string, operationResultID string, skip *int32, top *int32, skiptoken string) BaseClient {
	return BaseClient{
		Client:            autorest.NewClientWithUserAgent(UserAgent()),
		BaseURI:           baseURI,
		OperationResultID: operationResultID,
		Skip:              skip,
		Top:               top,
		Skiptoken:         skiptoken,
	}
}

// CheckNameAvailability checks if the specified management group name is valid and unique
// Parameters:
// checkNameAvailabilityRequest - management group name availability check parameters.
func (client BaseClient) CheckNameAvailability(ctx context.Context, checkNameAvailabilityRequest CheckNameAvailabilityRequest) (result CheckNameAvailabilityResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.CheckNameAvailability")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CheckNameAvailabilityPreparer(ctx, checkNameAvailabilityRequest)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managementgroups.BaseClient", "CheckNameAvailability", nil, "Failure preparing request")
		return
	}

	resp, err := client.CheckNameAvailabilitySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "managementgroups.BaseClient", "CheckNameAvailability", resp, "Failure sending request")
		return
	}

	result, err = client.CheckNameAvailabilityResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managementgroups.BaseClient", "CheckNameAvailability", resp, "Failure responding to request")
		return
	}

	return
}

// CheckNameAvailabilityPreparer prepares the CheckNameAvailability request.
func (client BaseClient) CheckNameAvailabilityPreparer(ctx context.Context, checkNameAvailabilityRequest CheckNameAvailabilityRequest) (*http.Request, error) {
	const APIVersion = "2020-02-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.Management/checkNameAvailability"),
		autorest.WithJSON(checkNameAvailabilityRequest),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CheckNameAvailabilitySender sends the CheckNameAvailability request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) CheckNameAvailabilitySender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CheckNameAvailabilityResponder handles the response to the CheckNameAvailability request. The method always
// closes the http.Response Body.
func (client BaseClient) CheckNameAvailabilityResponder(resp *http.Response) (result CheckNameAvailabilityResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// StartTenantBackfill starts backfilling subscriptions for the Tenant.
func (client BaseClient) StartTenantBackfill(ctx context.Context) (result TenantBackfillStatusResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.StartTenantBackfill")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.StartTenantBackfillPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managementgroups.BaseClient", "StartTenantBackfill", nil, "Failure preparing request")
		return
	}

	resp, err := client.StartTenantBackfillSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "managementgroups.BaseClient", "StartTenantBackfill", resp, "Failure sending request")
		return
	}

	result, err = client.StartTenantBackfillResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managementgroups.BaseClient", "StartTenantBackfill", resp, "Failure responding to request")
		return
	}

	return
}

// StartTenantBackfillPreparer prepares the StartTenantBackfill request.
func (client BaseClient) StartTenantBackfillPreparer(ctx context.Context) (*http.Request, error) {
	const APIVersion = "2020-02-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.Management/startTenantBackfill"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// StartTenantBackfillSender sends the StartTenantBackfill request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) StartTenantBackfillSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// StartTenantBackfillResponder handles the response to the StartTenantBackfill request. The method always
// closes the http.Response Body.
func (client BaseClient) StartTenantBackfillResponder(resp *http.Response) (result TenantBackfillStatusResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// TenantBackfillStatus gets tenant backfill status
func (client BaseClient) TenantBackfillStatus(ctx context.Context) (result TenantBackfillStatusResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.TenantBackfillStatus")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.TenantBackfillStatusPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managementgroups.BaseClient", "TenantBackfillStatus", nil, "Failure preparing request")
		return
	}

	resp, err := client.TenantBackfillStatusSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "managementgroups.BaseClient", "TenantBackfillStatus", resp, "Failure sending request")
		return
	}

	result, err = client.TenantBackfillStatusResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managementgroups.BaseClient", "TenantBackfillStatus", resp, "Failure responding to request")
		return
	}

	return
}

// TenantBackfillStatusPreparer prepares the TenantBackfillStatus request.
func (client BaseClient) TenantBackfillStatusPreparer(ctx context.Context) (*http.Request, error) {
	const APIVersion = "2020-02-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.Management/tenantBackfillStatus"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// TenantBackfillStatusSender sends the TenantBackfillStatus request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) TenantBackfillStatusSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// TenantBackfillStatusResponder handles the response to the TenantBackfillStatus request. The method always
// closes the http.Response Body.
func (client BaseClient) TenantBackfillStatusResponder(resp *http.Response) (result TenantBackfillStatusResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
