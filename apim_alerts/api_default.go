/*
 * Alert API Manager
 *
 * Alert for API Manager
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apim_alerts

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

type DefaultApiApiCreateAlertForAPIManagerInstanceRequest struct {
	ctx _context.Context
	ApiService *DefaultApiService
	orgId string
	envId string
	apiVersion string
	alertCore *AlertCore
}

func (r DefaultApiApiCreateAlertForAPIManagerInstanceRequest) AlertCore(alertCore AlertCore) DefaultApiApiCreateAlertForAPIManagerInstanceRequest {
	r.alertCore = &alertCore
	return r
}

func (r DefaultApiApiCreateAlertForAPIManagerInstanceRequest) Execute() (Alert, *_nethttp.Response, error) {
	return r.ApiService.CreateAlertForAPIManagerInstanceExecute(r)
}

/*
 * CreateAlertForAPIManagerInstance Method for CreateAlertForAPIManagerInstance
 * Create a new alert for API manager
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param orgId The organization Id
 * @param envId The environment id
 * @param apiVersion The api version
 * @return DefaultApiApiCreateAlertForAPIManagerInstanceRequest
 */
func (a *DefaultApiService) CreateAlertForAPIManagerInstance(ctx _context.Context, orgId string, envId string, apiVersion string) DefaultApiApiCreateAlertForAPIManagerInstanceRequest {
	return DefaultApiApiCreateAlertForAPIManagerInstanceRequest{
		ApiService: a,
		ctx: ctx,
		orgId: orgId,
		envId: envId,
		apiVersion: apiVersion,
	}
}

/*
 * Execute executes the request
 * @return Alert
 */
func (a *DefaultApiService) CreateAlertForAPIManagerInstanceExecute(r DefaultApiApiCreateAlertForAPIManagerInstanceRequest) (Alert, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Alert
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.CreateAlertForAPIManagerInstance")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{orgId}/environments/{envId}/apis/{apiVersion}/alerts"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", _neturl.PathEscape(parameterToString(r.orgId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"envId"+"}", _neturl.PathEscape(parameterToString(r.envId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"apiVersion"+"}", _neturl.PathEscape(parameterToString(r.apiVersion, "")), -1)

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
	localVarPostBody = r.alertCore
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

type DefaultApiApiDeleteanAlertfromAPImanagerRequest struct {
	ctx _context.Context
	ApiService *DefaultApiService
	orgId string
	envId string
	apiVersion string
	alertId string
}


func (r DefaultApiApiDeleteanAlertfromAPImanagerRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DeleteanAlertfromAPImanagerExecute(r)
}

/*
 * DeleteanAlertfromAPImanager DeleteanAlertfromAPImanager
 * Delete an Alert from API manager
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param orgId The organization Id
 * @param envId the environment id
 * @param apiVersion the api version
 * @param alertId the id of the alert
 * @return DefaultApiApiDeleteanAlertfromAPImanagerRequest
 */
func (a *DefaultApiService) DeleteanAlertfromAPImanager(ctx _context.Context, orgId string, envId string, apiVersion string, alertId string) DefaultApiApiDeleteanAlertfromAPImanagerRequest {
	return DefaultApiApiDeleteanAlertfromAPImanagerRequest{
		ApiService: a,
		ctx: ctx,
		orgId: orgId,
		envId: envId,
		apiVersion: apiVersion,
		alertId: alertId,
	}
}

/*
 * Execute executes the request
 */
func (a *DefaultApiService) DeleteanAlertfromAPImanagerExecute(r DefaultApiApiDeleteanAlertfromAPImanagerRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.DeleteanAlertfromAPImanager")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{orgId}/environments/{envId}/apis/{apiVersion}/alerts/{alertId}"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", _neturl.PathEscape(parameterToString(r.orgId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"envId"+"}", _neturl.PathEscape(parameterToString(r.envId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"apiVersion"+"}", _neturl.PathEscape(parameterToString(r.apiVersion, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"alertId"+"}", _neturl.PathEscape(parameterToString(r.alertId, "")), -1)

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
	localVarHTTPHeaderAccepts := []string{"text/plain"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type DefaultApiApiModifyonealertfromAPImangerRequest struct {
	ctx _context.Context
	ApiService *DefaultApiService
	orgId string
	envId string
	apiVersion string
	alertId string
	alertCore *AlertCore
}

func (r DefaultApiApiModifyonealertfromAPImangerRequest) AlertCore(alertCore AlertCore) DefaultApiApiModifyonealertfromAPImangerRequest {
	r.alertCore = &alertCore
	return r
}

func (r DefaultApiApiModifyonealertfromAPImangerRequest) Execute() (Alert, *_nethttp.Response, error) {
	return r.ApiService.ModifyonealertfromAPImangerExecute(r)
}

/*
 * ModifyonealertfromAPImanger ModifyonealertfromAPImanger
 * Modify one alert from API manger
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param orgId The organization Id
 * @param envId the environment id
 * @param apiVersion the api version
 * @param alertId the id of the alert
 * @return DefaultApiApiModifyonealertfromAPImangerRequest
 */
func (a *DefaultApiService) ModifyonealertfromAPImanger(ctx _context.Context, orgId string, envId string, apiVersion string, alertId string) DefaultApiApiModifyonealertfromAPImangerRequest {
	return DefaultApiApiModifyonealertfromAPImangerRequest{
		ApiService: a,
		ctx: ctx,
		orgId: orgId,
		envId: envId,
		apiVersion: apiVersion,
		alertId: alertId,
	}
}

/*
 * Execute executes the request
 * @return Alert
 */
func (a *DefaultApiService) ModifyonealertfromAPImangerExecute(r DefaultApiApiModifyonealertfromAPImangerRequest) (Alert, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Alert
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.ModifyonealertfromAPImanger")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{orgId}/environments/{envId}/apis/{apiVersion}/alerts/{alertId}"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", _neturl.PathEscape(parameterToString(r.orgId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"envId"+"}", _neturl.PathEscape(parameterToString(r.envId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"apiVersion"+"}", _neturl.PathEscape(parameterToString(r.apiVersion, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"alertId"+"}", _neturl.PathEscape(parameterToString(r.alertId, "")), -1)

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
	localVarPostBody = r.alertCore
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

type DefaultApiApiOrganizationsOrgIdEnvironmentsEnvIdApisApiVersionAlertsAlertIdGetRequest struct {
	ctx _context.Context
	ApiService *DefaultApiService
	orgId string
	envId string
	apiVersion string
	alertId string
}


func (r DefaultApiApiOrganizationsOrgIdEnvironmentsEnvIdApisApiVersionAlertsAlertIdGetRequest) Execute() (Alert, *_nethttp.Response, error) {
	return r.ApiService.OrganizationsOrgIdEnvironmentsEnvIdApisApiVersionAlertsAlertIdGetExecute(r)
}

/*
 * OrganizationsOrgIdEnvironmentsEnvIdApisApiVersionAlertsAlertIdGet Getonealert
 * Get one alert
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param orgId The organization Id
 * @param envId the environment id
 * @param apiVersion the api version
 * @param alertId the id of the alert
 * @return DefaultApiApiOrganizationsOrgIdEnvironmentsEnvIdApisApiVersionAlertsAlertIdGetRequest
 */
func (a *DefaultApiService) OrganizationsOrgIdEnvironmentsEnvIdApisApiVersionAlertsAlertIdGet(ctx _context.Context, orgId string, envId string, apiVersion string, alertId string) DefaultApiApiOrganizationsOrgIdEnvironmentsEnvIdApisApiVersionAlertsAlertIdGetRequest {
	return DefaultApiApiOrganizationsOrgIdEnvironmentsEnvIdApisApiVersionAlertsAlertIdGetRequest{
		ApiService: a,
		ctx: ctx,
		orgId: orgId,
		envId: envId,
		apiVersion: apiVersion,
		alertId: alertId,
	}
}

/*
 * Execute executes the request
 * @return Alert
 */
func (a *DefaultApiService) OrganizationsOrgIdEnvironmentsEnvIdApisApiVersionAlertsAlertIdGetExecute(r DefaultApiApiOrganizationsOrgIdEnvironmentsEnvIdApisApiVersionAlertsAlertIdGetRequest) (Alert, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Alert
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.OrganizationsOrgIdEnvironmentsEnvIdApisApiVersionAlertsAlertIdGet")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{orgId}/environments/{envId}/apis/{apiVersion}/alerts/{alertId}"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", _neturl.PathEscape(parameterToString(r.orgId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"envId"+"}", _neturl.PathEscape(parameterToString(r.envId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"apiVersion"+"}", _neturl.PathEscape(parameterToString(r.apiVersion, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"alertId"+"}", _neturl.PathEscape(parameterToString(r.alertId, "")), -1)

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

type DefaultApiApiOrganizationsOrgIdEnvironmentsEnvIdApisApiVersionAlertsGetRequest struct {
	ctx _context.Context
	ApiService *DefaultApiService
	orgId string
	envId string
	apiVersion string
}


func (r DefaultApiApiOrganizationsOrgIdEnvironmentsEnvIdApisApiVersionAlertsGetRequest) Execute() ([]Alert, *_nethttp.Response, error) {
	return r.ApiService.OrganizationsOrgIdEnvironmentsEnvIdApisApiVersionAlertsGetExecute(r)
}

/*
 * OrganizationsOrgIdEnvironmentsEnvIdApisApiVersionAlertsGet GetAlertsfromAPImanager
 * Get Alerts from API manager
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param orgId The organization Id
 * @param envId The environment id
 * @param apiVersion The api version
 * @return DefaultApiApiOrganizationsOrgIdEnvironmentsEnvIdApisApiVersionAlertsGetRequest
 */
func (a *DefaultApiService) OrganizationsOrgIdEnvironmentsEnvIdApisApiVersionAlertsGet(ctx _context.Context, orgId string, envId string, apiVersion string) DefaultApiApiOrganizationsOrgIdEnvironmentsEnvIdApisApiVersionAlertsGetRequest {
	return DefaultApiApiOrganizationsOrgIdEnvironmentsEnvIdApisApiVersionAlertsGetRequest{
		ApiService: a,
		ctx: ctx,
		orgId: orgId,
		envId: envId,
		apiVersion: apiVersion,
	}
}

/*
 * Execute executes the request
 * @return []Alert
 */
func (a *DefaultApiService) OrganizationsOrgIdEnvironmentsEnvIdApisApiVersionAlertsGetExecute(r DefaultApiApiOrganizationsOrgIdEnvironmentsEnvIdApisApiVersionAlertsGetRequest) ([]Alert, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []Alert
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.OrganizationsOrgIdEnvironmentsEnvIdApisApiVersionAlertsGet")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{orgId}/environments/{envId}/apis/{apiVersion}/alerts"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", _neturl.PathEscape(parameterToString(r.orgId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"envId"+"}", _neturl.PathEscape(parameterToString(r.envId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"apiVersion"+"}", _neturl.PathEscape(parameterToString(r.apiVersion, "")), -1)

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
