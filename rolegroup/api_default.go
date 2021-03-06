/*
 * RoleGroup API
 *
 * Description of the RoleGroup API
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rolegroup

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// DefaultApiService DefaultApi service
type DefaultApiService service

type DefaultApiApiOrganizationsOrgIdRolegroupsGetRequest struct {
	ctx _context.Context
	ApiService *DefaultApiService
	orgId string
}


func (r DefaultApiApiOrganizationsOrgIdRolegroupsGetRequest) Execute() (InlineResponse200, *_nethttp.Response, error) {
	return r.ApiService.OrganizationsOrgIdRolegroupsGetExecute(r)
}

/*
 * OrganizationsOrgIdRolegroupsGet Method for OrganizationsOrgIdRolegroupsGet
 * Returns all rolegroups within the organization
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param orgId The ID of the organization in GUID format
 * @return DefaultApiApiOrganizationsOrgIdRolegroupsGetRequest
 */
func (a *DefaultApiService) OrganizationsOrgIdRolegroupsGet(ctx _context.Context, orgId string) DefaultApiApiOrganizationsOrgIdRolegroupsGetRequest {
	return DefaultApiApiOrganizationsOrgIdRolegroupsGetRequest{
		ApiService: a,
		ctx: ctx,
		orgId: orgId,
	}
}

/*
 * Execute executes the request
 * @return InlineResponse200
 */
func (a *DefaultApiService) OrganizationsOrgIdRolegroupsGetExecute(r DefaultApiApiOrganizationsOrgIdRolegroupsGetRequest) (InlineResponse200, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  InlineResponse200
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.OrganizationsOrgIdRolegroupsGet")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{orgId}/rolegroups"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", _neturl.PathEscape(parameterToString(r.orgId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type DefaultApiApiOrganizationsOrgIdRolegroupsPostRequest struct {
	ctx _context.Context
	ApiService *DefaultApiService
	orgId string
	rolegroupPostBody *RolegroupPostBody
}

func (r DefaultApiApiOrganizationsOrgIdRolegroupsPostRequest) RolegroupPostBody(rolegroupPostBody RolegroupPostBody) DefaultApiApiOrganizationsOrgIdRolegroupsPostRequest {
	r.rolegroupPostBody = &rolegroupPostBody
	return r
}

func (r DefaultApiApiOrganizationsOrgIdRolegroupsPostRequest) Execute() (Rolegroup, *_nethttp.Response, error) {
	return r.ApiService.OrganizationsOrgIdRolegroupsPostExecute(r)
}

/*
 * OrganizationsOrgIdRolegroupsPost Method for OrganizationsOrgIdRolegroupsPost
 * Creates an rolegroup. Name must be unique.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param orgId The ID of the organization in GUID format
 * @return DefaultApiApiOrganizationsOrgIdRolegroupsPostRequest
 */
func (a *DefaultApiService) OrganizationsOrgIdRolegroupsPost(ctx _context.Context, orgId string) DefaultApiApiOrganizationsOrgIdRolegroupsPostRequest {
	return DefaultApiApiOrganizationsOrgIdRolegroupsPostRequest{
		ApiService: a,
		ctx: ctx,
		orgId: orgId,
	}
}

/*
 * Execute executes the request
 * @return Rolegroup
 */
func (a *DefaultApiService) OrganizationsOrgIdRolegroupsPostExecute(r DefaultApiApiOrganizationsOrgIdRolegroupsPostRequest) (Rolegroup, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Rolegroup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.OrganizationsOrgIdRolegroupsPost")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{orgId}/rolegroups"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", _neturl.PathEscape(parameterToString(r.orgId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.rolegroupPostBody
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v InlineResponse400
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type DefaultApiApiOrganizationsOrgIdRolegroupsRolegroupIdDeleteRequest struct {
	ctx _context.Context
	ApiService *DefaultApiService
	orgId string
	rolegroupId string
}


func (r DefaultApiApiOrganizationsOrgIdRolegroupsRolegroupIdDeleteRequest) Execute() ([]string, *_nethttp.Response, error) {
	return r.ApiService.OrganizationsOrgIdRolegroupsRolegroupIdDeleteExecute(r)
}

/*
 * OrganizationsOrgIdRolegroupsRolegroupIdDelete Method for OrganizationsOrgIdRolegroupsRolegroupIdDelete
 * Delete a rolegroup
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param orgId The ID of the organization in GUID format
 * @param rolegroupId The id of a rolegroup
 * @return DefaultApiApiOrganizationsOrgIdRolegroupsRolegroupIdDeleteRequest
 */
func (a *DefaultApiService) OrganizationsOrgIdRolegroupsRolegroupIdDelete(ctx _context.Context, orgId string, rolegroupId string) DefaultApiApiOrganizationsOrgIdRolegroupsRolegroupIdDeleteRequest {
	return DefaultApiApiOrganizationsOrgIdRolegroupsRolegroupIdDeleteRequest{
		ApiService: a,
		ctx: ctx,
		orgId: orgId,
		rolegroupId: rolegroupId,
	}
}

/*
 * Execute executes the request
 * @return []string
 */
func (a *DefaultApiService) OrganizationsOrgIdRolegroupsRolegroupIdDeleteExecute(r DefaultApiApiOrganizationsOrgIdRolegroupsRolegroupIdDeleteRequest) ([]string, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.OrganizationsOrgIdRolegroupsRolegroupIdDelete")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{orgId}/rolegroups/{rolegroupId}"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", _neturl.PathEscape(parameterToString(r.orgId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"rolegroupId"+"}", _neturl.PathEscape(parameterToString(r.rolegroupId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type DefaultApiApiOrganizationsOrgIdRolegroupsRolegroupIdGetRequest struct {
	ctx _context.Context
	ApiService *DefaultApiService
	orgId string
	rolegroupId string
}


func (r DefaultApiApiOrganizationsOrgIdRolegroupsRolegroupIdGetRequest) Execute() (Rolegroup, *_nethttp.Response, error) {
	return r.ApiService.OrganizationsOrgIdRolegroupsRolegroupIdGetExecute(r)
}

/*
 * OrganizationsOrgIdRolegroupsRolegroupIdGet Method for OrganizationsOrgIdRolegroupsRolegroupIdGet
 * Retrieves a rolegroup by id
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param orgId The ID of the organization in GUID format
 * @param rolegroupId The id of rolegroup
 * @return DefaultApiApiOrganizationsOrgIdRolegroupsRolegroupIdGetRequest
 */
func (a *DefaultApiService) OrganizationsOrgIdRolegroupsRolegroupIdGet(ctx _context.Context, orgId string, rolegroupId string) DefaultApiApiOrganizationsOrgIdRolegroupsRolegroupIdGetRequest {
	return DefaultApiApiOrganizationsOrgIdRolegroupsRolegroupIdGetRequest{
		ApiService: a,
		ctx: ctx,
		orgId: orgId,
		rolegroupId: rolegroupId,
	}
}

/*
 * Execute executes the request
 * @return Rolegroup
 */
func (a *DefaultApiService) OrganizationsOrgIdRolegroupsRolegroupIdGetExecute(r DefaultApiApiOrganizationsOrgIdRolegroupsRolegroupIdGetRequest) (Rolegroup, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Rolegroup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.OrganizationsOrgIdRolegroupsRolegroupIdGet")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{orgId}/rolegroups/{rolegroupId}"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", _neturl.PathEscape(parameterToString(r.orgId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"rolegroupId"+"}", _neturl.PathEscape(parameterToString(r.rolegroupId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type DefaultApiApiOrganizationsOrgIdRolegroupsRolegroupIdPutRequest struct {
	ctx _context.Context
	ApiService *DefaultApiService
	orgId string
	rolegroupId string
	rolegroupPutBody *RolegroupPutBody
}

func (r DefaultApiApiOrganizationsOrgIdRolegroupsRolegroupIdPutRequest) RolegroupPutBody(rolegroupPutBody RolegroupPutBody) DefaultApiApiOrganizationsOrgIdRolegroupsRolegroupIdPutRequest {
	r.rolegroupPutBody = &rolegroupPutBody
	return r
}

func (r DefaultApiApiOrganizationsOrgIdRolegroupsRolegroupIdPutRequest) Execute() (Rolegroup, *_nethttp.Response, error) {
	return r.ApiService.OrganizationsOrgIdRolegroupsRolegroupIdPutExecute(r)
}

/*
 * OrganizationsOrgIdRolegroupsRolegroupIdPut Method for OrganizationsOrgIdRolegroupsRolegroupIdPut
 * Update a role gorup, implemented as a patch. Note that only the name (must be unique) and description is allowed to be updated.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param orgId The ID of the organization in GUID format
 * @param rolegroupId The id of an rolegroup
 * @return DefaultApiApiOrganizationsOrgIdRolegroupsRolegroupIdPutRequest
 */
func (a *DefaultApiService) OrganizationsOrgIdRolegroupsRolegroupIdPut(ctx _context.Context, orgId string, rolegroupId string) DefaultApiApiOrganizationsOrgIdRolegroupsRolegroupIdPutRequest {
	return DefaultApiApiOrganizationsOrgIdRolegroupsRolegroupIdPutRequest{
		ApiService: a,
		ctx: ctx,
		orgId: orgId,
		rolegroupId: rolegroupId,
	}
}

/*
 * Execute executes the request
 * @return Rolegroup
 */
func (a *DefaultApiService) OrganizationsOrgIdRolegroupsRolegroupIdPutExecute(r DefaultApiApiOrganizationsOrgIdRolegroupsRolegroupIdPutRequest) (Rolegroup, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Rolegroup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.OrganizationsOrgIdRolegroupsRolegroupIdPut")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{orgId}/rolegroups/{rolegroupId}"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", _neturl.PathEscape(parameterToString(r.orgId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"rolegroupId"+"}", _neturl.PathEscape(parameterToString(r.rolegroupId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.rolegroupPutBody
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v InlineResponse400
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
