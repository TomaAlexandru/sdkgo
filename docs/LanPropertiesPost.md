# LanPropertiesPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A name of that resource | [optional] 
**Public** | Pointer to **bool** | Does this LAN faces the public Internet or not | [optional] 

## Methods

### NewLanPropertiesPost

`func NewLanPropertiesPost() *LanPropertiesPost`

NewLanPropertiesPost instantiates a new LanPropertiesPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLanPropertiesPostWithDefaults

`func NewLanPropertiesPostWithDefaults() *LanPropertiesPost`

NewLanPropertiesPostWithDefaults instantiates a new LanPropertiesPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *LanPropertiesPost) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LanPropertiesPost) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LanPropertiesPost) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LanPropertiesPost) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPublic

`func (o *LanPropertiesPost) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *LanPropertiesPost) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *LanPropertiesPost) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *LanPropertiesPost) HasPublic() bool`

HasPublic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


