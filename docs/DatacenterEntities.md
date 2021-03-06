# DatacenterEntities

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Servers** | Pointer to [**Servers**](Servers.md) |  | [optional] 
**Volumes** | Pointer to [**Volumes**](Volumes.md) |  | [optional] 
**Loadbalancers** | Pointer to [**Loadbalancers**](Loadbalancers.md) |  | [optional] 
**Lans** | Pointer to [**Lans**](Lans.md) |  | [optional] 

## Methods

### NewDatacenterEntities

`func NewDatacenterEntities() *DatacenterEntities`

NewDatacenterEntities instantiates a new DatacenterEntities object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatacenterEntitiesWithDefaults

`func NewDatacenterEntitiesWithDefaults() *DatacenterEntities`

NewDatacenterEntitiesWithDefaults instantiates a new DatacenterEntities object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServers

`func (o *DatacenterEntities) GetServers() Servers`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *DatacenterEntities) GetServersOk() (*Servers, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *DatacenterEntities) SetServers(v Servers)`

SetServers sets Servers field to given value.

### HasServers

`func (o *DatacenterEntities) HasServers() bool`

HasServers returns a boolean if a field has been set.

### GetVolumes

`func (o *DatacenterEntities) GetVolumes() Volumes`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *DatacenterEntities) GetVolumesOk() (*Volumes, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *DatacenterEntities) SetVolumes(v Volumes)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *DatacenterEntities) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### GetLoadbalancers

`func (o *DatacenterEntities) GetLoadbalancers() Loadbalancers`

GetLoadbalancers returns the Loadbalancers field if non-nil, zero value otherwise.

### GetLoadbalancersOk

`func (o *DatacenterEntities) GetLoadbalancersOk() (*Loadbalancers, bool)`

GetLoadbalancersOk returns a tuple with the Loadbalancers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadbalancers

`func (o *DatacenterEntities) SetLoadbalancers(v Loadbalancers)`

SetLoadbalancers sets Loadbalancers field to given value.

### HasLoadbalancers

`func (o *DatacenterEntities) HasLoadbalancers() bool`

HasLoadbalancers returns a boolean if a field has been set.

### GetLans

`func (o *DatacenterEntities) GetLans() Lans`

GetLans returns the Lans field if non-nil, zero value otherwise.

### GetLansOk

`func (o *DatacenterEntities) GetLansOk() (*Lans, bool)`

GetLansOk returns a tuple with the Lans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLans

`func (o *DatacenterEntities) SetLans(v Lans)`

SetLans sets Lans field to given value.

### HasLans

`func (o *DatacenterEntities) HasLans() bool`

HasLans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


