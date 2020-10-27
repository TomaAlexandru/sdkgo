# Datacenters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The resource&#39;s unique identifier | [optional] [readonly] 
**Type** | Pointer to [**Type**](Type.md) | The type of object that has been created | [optional] 
**Href** | Pointer to **string** | URL to the object representation (absolute path) | [optional] [readonly] 
**Items** | Pointer to [**[]Datacenter**](Datacenter.md) | Array of items in that collection | [optional] [readonly] 

## Methods

### NewDatacenters

`func NewDatacenters() *Datacenters`

NewDatacenters instantiates a new Datacenters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatacentersWithDefaults

`func NewDatacentersWithDefaults() *Datacenters`

NewDatacentersWithDefaults instantiates a new Datacenters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Datacenters) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Datacenters) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Datacenters) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Datacenters) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *Datacenters) GetType() Type`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Datacenters) GetTypeOk() (*Type, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Datacenters) SetType(v Type)`

SetType sets Type field to given value.

### HasType

`func (o *Datacenters) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHref

`func (o *Datacenters) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *Datacenters) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *Datacenters) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *Datacenters) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetItems

`func (o *Datacenters) GetItems() []Datacenter`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *Datacenters) GetItemsOk() (*[]Datacenter, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *Datacenters) SetItems(v []Datacenter)`

SetItems sets Items field to given value.

### HasItems

`func (o *Datacenters) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


