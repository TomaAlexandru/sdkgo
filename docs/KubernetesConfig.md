# KubernetesConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The resource&#39;s unique identifier. | [optional] [readonly] 
**Type** | Pointer to **string** | The type of object | [optional] [readonly] 
**Href** | Pointer to **string** | URL to the object representation (absolute path) | [optional] [readonly] 
**Properties** | [**KubernetesConfigProperties**](KubernetesConfigProperties.md) |  | 

## Methods

### NewKubernetesConfig

`func NewKubernetesConfig(properties KubernetesConfigProperties, ) *KubernetesConfig`

NewKubernetesConfig instantiates a new KubernetesConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesConfigWithDefaults

`func NewKubernetesConfigWithDefaults() *KubernetesConfig`

NewKubernetesConfigWithDefaults instantiates a new KubernetesConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KubernetesConfig) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KubernetesConfig) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KubernetesConfig) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *KubernetesConfig) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *KubernetesConfig) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *KubernetesConfig) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *KubernetesConfig) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *KubernetesConfig) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHref

`func (o *KubernetesConfig) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *KubernetesConfig) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *KubernetesConfig) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *KubernetesConfig) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetProperties

`func (o *KubernetesConfig) GetProperties() KubernetesConfigProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *KubernetesConfig) GetPropertiesOk() (*KubernetesConfigProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *KubernetesConfig) SetProperties(v KubernetesConfigProperties)`

SetProperties sets Properties field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


