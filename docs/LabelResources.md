# LabelResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique representation for Label as a collection on a resource. | [optional] [readonly] 
**Type** | Pointer to **string** | The type of resource within a collection | [optional] [readonly] 
**Href** | Pointer to **string** | URL to the collection representation (absolute path) | [optional] [readonly] 
**Items** | Pointer to [**[]LabelResource**](LabelResource.md) | Array of items in that collection | [optional] [readonly] 

## Methods

### NewLabelResources

`func NewLabelResources() *LabelResources`

NewLabelResources instantiates a new LabelResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLabelResourcesWithDefaults

`func NewLabelResourcesWithDefaults() *LabelResources`

NewLabelResourcesWithDefaults instantiates a new LabelResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LabelResources) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LabelResources) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LabelResources) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LabelResources) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *LabelResources) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LabelResources) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LabelResources) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *LabelResources) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHref

`func (o *LabelResources) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *LabelResources) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *LabelResources) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *LabelResources) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetItems

`func (o *LabelResources) GetItems() []LabelResource`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *LabelResources) GetItemsOk() (*[]LabelResource, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *LabelResources) SetItems(v []LabelResource)`

SetItems sets Items field to given value.

### HasItems

`func (o *LabelResources) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


