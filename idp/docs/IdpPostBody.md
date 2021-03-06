# IdpPostBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**IdpPostBodyType**](IdpPostBodyType.md) |  | [optional] 
**OidcProvider** | Pointer to [**OidcProvider1**](OidcProvider1.md) |  | [optional] 
**AllowUntrustedCertificates** | Pointer to **bool** |  | [optional] 
**Saml** | Pointer to [**Saml1**](Saml1.md) |  | [optional] 
**ServiceProvider** | Pointer to [**ServiceProvider1**](ServiceProvider1.md) |  | [optional] 

## Methods

### NewIdpPostBody

`func NewIdpPostBody() *IdpPostBody`

NewIdpPostBody instantiates a new IdpPostBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdpPostBodyWithDefaults

`func NewIdpPostBodyWithDefaults() *IdpPostBody`

NewIdpPostBodyWithDefaults instantiates a new IdpPostBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *IdpPostBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IdpPostBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IdpPostBody) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IdpPostBody) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *IdpPostBody) GetType() IdpPostBodyType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IdpPostBody) GetTypeOk() (*IdpPostBodyType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IdpPostBody) SetType(v IdpPostBodyType)`

SetType sets Type field to given value.

### HasType

`func (o *IdpPostBody) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOidcProvider

`func (o *IdpPostBody) GetOidcProvider() OidcProvider1`

GetOidcProvider returns the OidcProvider field if non-nil, zero value otherwise.

### GetOidcProviderOk

`func (o *IdpPostBody) GetOidcProviderOk() (*OidcProvider1, bool)`

GetOidcProviderOk returns a tuple with the OidcProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcProvider

`func (o *IdpPostBody) SetOidcProvider(v OidcProvider1)`

SetOidcProvider sets OidcProvider field to given value.

### HasOidcProvider

`func (o *IdpPostBody) HasOidcProvider() bool`

HasOidcProvider returns a boolean if a field has been set.

### GetAllowUntrustedCertificates

`func (o *IdpPostBody) GetAllowUntrustedCertificates() bool`

GetAllowUntrustedCertificates returns the AllowUntrustedCertificates field if non-nil, zero value otherwise.

### GetAllowUntrustedCertificatesOk

`func (o *IdpPostBody) GetAllowUntrustedCertificatesOk() (*bool, bool)`

GetAllowUntrustedCertificatesOk returns a tuple with the AllowUntrustedCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUntrustedCertificates

`func (o *IdpPostBody) SetAllowUntrustedCertificates(v bool)`

SetAllowUntrustedCertificates sets AllowUntrustedCertificates field to given value.

### HasAllowUntrustedCertificates

`func (o *IdpPostBody) HasAllowUntrustedCertificates() bool`

HasAllowUntrustedCertificates returns a boolean if a field has been set.

### GetSaml

`func (o *IdpPostBody) GetSaml() Saml1`

GetSaml returns the Saml field if non-nil, zero value otherwise.

### GetSamlOk

`func (o *IdpPostBody) GetSamlOk() (*Saml1, bool)`

GetSamlOk returns a tuple with the Saml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaml

`func (o *IdpPostBody) SetSaml(v Saml1)`

SetSaml sets Saml field to given value.

### HasSaml

`func (o *IdpPostBody) HasSaml() bool`

HasSaml returns a boolean if a field has been set.

### GetServiceProvider

`func (o *IdpPostBody) GetServiceProvider() ServiceProvider1`

GetServiceProvider returns the ServiceProvider field if non-nil, zero value otherwise.

### GetServiceProviderOk

`func (o *IdpPostBody) GetServiceProviderOk() (*ServiceProvider1, bool)`

GetServiceProviderOk returns a tuple with the ServiceProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceProvider

`func (o *IdpPostBody) SetServiceProvider(v ServiceProvider1)`

SetServiceProvider sets ServiceProvider field to given value.

### HasServiceProvider

`func (o *IdpPostBody) HasServiceProvider() bool`

HasServiceProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


