# LanNics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The resource&#39;s unique identifier | [optional] [readonly] 
**Type** | Pointer to [**Type**](Type.md) | The type of object that has been created | [optional] 
**Href** | Pointer to **string** | URL to the object representation (absolute path) | [optional] [readonly] 
**Items** | Pointer to [**[]Nic**](Nic.md) | Array of items in that collection | [optional] [readonly] 

## Methods

### NewLanNics

`func NewLanNics() *LanNics`

NewLanNics instantiates a new LanNics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLanNicsWithDefaults

`func NewLanNicsWithDefaults() *LanNics`

NewLanNicsWithDefaults instantiates a new LanNics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LanNics) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LanNics) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LanNics) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LanNics) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *LanNics) GetType() Type`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LanNics) GetTypeOk() (*Type, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LanNics) SetType(v Type)`

SetType sets Type field to given value.

### HasType

`func (o *LanNics) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHref

`func (o *LanNics) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *LanNics) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *LanNics) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *LanNics) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetItems

`func (o *LanNics) GetItems() []Nic`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *LanNics) GetItemsOk() (*[]Nic, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *LanNics) SetItems(v []Nic)`

SetItems sets Items field to given value.

### HasItems

`func (o *LanNics) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


