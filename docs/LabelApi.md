# \LabelApi

All URIs are relative to *https://api.ionos.com/cloudapi/v5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DatacentersLabelsDelete**](LabelApi.md#DatacentersLabelsDelete) | **Delete** /datacenters/{datacenterId}/labels/{key} | Delete a Label from Data Center
[**DatacentersLabelsFindByKey**](LabelApi.md#DatacentersLabelsFindByKey) | **Get** /datacenters/{datacenterId}/labels/{key} | Retrieve a Label of Data Center
[**DatacentersLabelsGet**](LabelApi.md#DatacentersLabelsGet) | **Get** /datacenters/{datacenterId}/labels | List all Data Center Labels
[**DatacentersLabelsPost**](LabelApi.md#DatacentersLabelsPost) | **Post** /datacenters/{datacenterId}/labels | Add a Label to Data Center
[**DatacentersLabelsPut**](LabelApi.md#DatacentersLabelsPut) | **Put** /datacenters/{datacenterId}/labels/{key} | Modify a Label of Data Center
[**DatacentersServersLabelsDelete**](LabelApi.md#DatacentersServersLabelsDelete) | **Delete** /datacenters/{datacenterId}/servers/{serverId}/labels/{key} | Delete a Label from Server
[**DatacentersServersLabelsFindByKey**](LabelApi.md#DatacentersServersLabelsFindByKey) | **Get** /datacenters/{datacenterId}/servers/{serverId}/labels/{key} | Retrieve a Label of Server
[**DatacentersServersLabelsGet**](LabelApi.md#DatacentersServersLabelsGet) | **Get** /datacenters/{datacenterId}/servers/{serverId}/labels | List all Server Labels
[**DatacentersServersLabelsPost**](LabelApi.md#DatacentersServersLabelsPost) | **Post** /datacenters/{datacenterId}/servers/{serverId}/labels | Add a Label to Server
[**DatacentersServersLabelsPut**](LabelApi.md#DatacentersServersLabelsPut) | **Put** /datacenters/{datacenterId}/servers/{serverId}/labels/{key} | Modify a Label of Server
[**DatacentersVolumesLabelsDelete**](LabelApi.md#DatacentersVolumesLabelsDelete) | **Delete** /datacenters/{datacenterId}/volumes/{volumeId}/labels/{key} | Delete a Label from Volume
[**DatacentersVolumesLabelsFindByKey**](LabelApi.md#DatacentersVolumesLabelsFindByKey) | **Get** /datacenters/{datacenterId}/volumes/{volumeId}/labels/{key} | Retrieve a Label of Volume
[**DatacentersVolumesLabelsGet**](LabelApi.md#DatacentersVolumesLabelsGet) | **Get** /datacenters/{datacenterId}/volumes/{volumeId}/labels | List all Volume Labels
[**DatacentersVolumesLabelsPost**](LabelApi.md#DatacentersVolumesLabelsPost) | **Post** /datacenters/{datacenterId}/volumes/{volumeId}/labels | Add a Label to Volume
[**DatacentersVolumesLabelsPut**](LabelApi.md#DatacentersVolumesLabelsPut) | **Put** /datacenters/{datacenterId}/volumes/{volumeId}/labels/{key} | Modify a Label of Volume
[**IpblocksLabelsDelete**](LabelApi.md#IpblocksLabelsDelete) | **Delete** /ipblocks/{ipblockId}/labels/{key} | Delete a Label from IP Block
[**IpblocksLabelsFindByKey**](LabelApi.md#IpblocksLabelsFindByKey) | **Get** /ipblocks/{ipblockId}/labels/{key} | Retrieve a Label of IP Block
[**IpblocksLabelsGet**](LabelApi.md#IpblocksLabelsGet) | **Get** /ipblocks/{ipblockId}/labels | List all Ip Block Labels
[**IpblocksLabelsPost**](LabelApi.md#IpblocksLabelsPost) | **Post** /ipblocks/{ipblockId}/labels | Add a Label to IP Block
[**IpblocksLabelsPut**](LabelApi.md#IpblocksLabelsPut) | **Put** /ipblocks/{ipblockId}/labels/{key} | Modify a Label of IP Block
[**LabelsFindByLabelurn**](LabelApi.md#LabelsFindByLabelurn) | **Get** /labels/{labelurn} | Returns the label by its URN.
[**LabelsGet**](LabelApi.md#LabelsGet) | **Get** /labels | List Labels 
[**SnapshotsLabelsDelete**](LabelApi.md#SnapshotsLabelsDelete) | **Delete** /snapshots/{snapshotId}/labels/{key} | Delete a Label from Snapshot
[**SnapshotsLabelsFindByKey**](LabelApi.md#SnapshotsLabelsFindByKey) | **Get** /snapshots/{snapshotId}/labels/{key} | Retrieve a Label of Snapshot
[**SnapshotsLabelsGet**](LabelApi.md#SnapshotsLabelsGet) | **Get** /snapshots/{snapshotId}/labels | List all Snapshot Labels
[**SnapshotsLabelsPost**](LabelApi.md#SnapshotsLabelsPost) | **Post** /snapshots/{snapshotId}/labels | Add a Label to Snapshot
[**SnapshotsLabelsPut**](LabelApi.md#SnapshotsLabelsPut) | **Put** /snapshots/{snapshotId}/labels/{key} | Modify a Label of Snapshot



## DatacentersLabelsDelete

> map[string]interface{} DatacentersLabelsDelete(ctx, datacenterId, key).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()

Delete a Label from Data Center



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    datacenterId := "datacenterId_example" // string | The unique ID of the Data Center
    key := "key_example" // string | The key of the Label
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := 987 // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := 987 // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LabelApi.DatacentersLabelsDelete(context.Background(), datacenterId, key).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LabelApi.DatacentersLabelsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersLabelsDelete`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LabelApi.DatacentersLabelsDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datacenterId** | **string** | The unique ID of the Data Center | 
**key** | **string** | The key of the Label | 

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersLabelsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

**map[string]interface{}**

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatacentersLabelsFindByKey

> LabelResource DatacentersLabelsFindByKey(ctx, datacenterId, key).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()

Retrieve a Label of Data Center



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    datacenterId := "datacenterId_example" // string | The unique ID of the Data Center
    key := "key_example" // string | The key of the Label
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := 987 // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := 987 // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LabelApi.DatacentersLabelsFindByKey(context.Background(), datacenterId, key).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LabelApi.DatacentersLabelsFindByKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersLabelsFindByKey`: LabelResource
    fmt.Fprintf(os.Stdout, "Response from `LabelApi.DatacentersLabelsFindByKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datacenterId** | **string** | The unique ID of the Data Center | 
**key** | **string** | The key of the Label | 

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersLabelsFindByKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**LabelResource**](LabelResource.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatacentersLabelsGet

> LabelResources DatacentersLabelsGet(ctx, datacenterId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()

List all Data Center Labels



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    datacenterId := "datacenterId_example" // string | The unique ID of the Data Center
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := 987 // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := 987 // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LabelApi.DatacentersLabelsGet(context.Background(), datacenterId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LabelApi.DatacentersLabelsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersLabelsGet`: LabelResources
    fmt.Fprintf(os.Stdout, "Response from `LabelApi.DatacentersLabelsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datacenterId** | **string** | The unique ID of the Data Center | 

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersLabelsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**LabelResources**](LabelResources.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatacentersLabelsPost

> LabelResource DatacentersLabelsPost(ctx, datacenterId).Label(label).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()

Add a Label to Data Center



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    datacenterId := "datacenterId_example" // string | The unique ID of the Data Center
    label := *openapiclient.NewLabelResource(*openapiclient.NewLabelResourceProperties()) // LabelResource | Label to be added
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := 987 // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := 987 // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LabelApi.DatacentersLabelsPost(context.Background(), datacenterId).Label(label).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LabelApi.DatacentersLabelsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersLabelsPost`: LabelResource
    fmt.Fprintf(os.Stdout, "Response from `LabelApi.DatacentersLabelsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datacenterId** | **string** | The unique ID of the Data Center | 

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersLabelsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **label** | [**LabelResource**](LabelResource.md) | Label to be added | 
 **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**LabelResource**](LabelResource.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatacentersLabelsPut

> LabelResource DatacentersLabelsPut(ctx, datacenterId, key).Label(label).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()

Modify a Label of Data Center



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    datacenterId := "datacenterId_example" // string | The unique ID of the Data Center
    key := "key_example" // string | The key of the Label
    label := *openapiclient.NewLabelResource(*openapiclient.NewLabelResourceProperties()) // LabelResource | Modified Label
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := 987 // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := 987 // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LabelApi.DatacentersLabelsPut(context.Background(), datacenterId, key).Label(label).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LabelApi.DatacentersLabelsPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersLabelsPut`: LabelResource
    fmt.Fprintf(os.Stdout, "Response from `LabelApi.DatacentersLabelsPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datacenterId** | **string** | The unique ID of the Data Center | 
**key** | **string** | The key of the Label | 

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersLabelsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **label** | [**LabelResource**](LabelResource.md) | Modified Label | 
 **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**LabelResource**](LabelResource.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatacentersServersLabelsDelete

> map[string]interface{} DatacentersServersLabelsDelete(ctx, datacenterId, serverId, key).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()

Delete a Label from Server



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    datacenterId := "datacenterId_example" // string | The unique ID of the Datacenter
    serverId := "serverId_example" // string | The unique ID of the Server
    key := "key_example" // string | The key of the Label
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := 987 // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := 987 // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LabelApi.DatacentersServersLabelsDelete(context.Background(), datacenterId, serverId, key).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LabelApi.DatacentersServersLabelsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersServersLabelsDelete`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LabelApi.DatacentersServersLabelsDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datacenterId** | **string** | The unique ID of the Datacenter | 
**serverId** | **string** | The unique ID of the Server | 
**key** | **string** | The key of the Label | 

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersServersLabelsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

**map[string]interface{}**

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatacentersServersLabelsFindByKey

> LabelResource DatacentersServersLabelsFindByKey(ctx, datacenterId, serverId, key).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()

Retrieve a Label of Server



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    datacenterId := "datacenterId_example" // string | The unique ID of the Datacenter
    serverId := "serverId_example" // string | The unique ID of the Server
    key := "key_example" // string | The key of the Label
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := 987 // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := 987 // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LabelApi.DatacentersServersLabelsFindByKey(context.Background(), datacenterId, serverId, key).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LabelApi.DatacentersServersLabelsFindByKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersServersLabelsFindByKey`: LabelResource
    fmt.Fprintf(os.Stdout, "Response from `LabelApi.DatacentersServersLabelsFindByKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datacenterId** | **string** | The unique ID of the Datacenter | 
**serverId** | **string** | The unique ID of the Server | 
**key** | **string** | The key of the Label | 

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersServersLabelsFindByKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**LabelResource**](LabelResource.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatacentersServersLabelsGet

> LabelResources DatacentersServersLabelsGet(ctx, datacenterId, serverId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()

List all Server Labels



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    datacenterId := "datacenterId_example" // string | The unique ID of the Datacenter
    serverId := "serverId_example" // string | The unique ID of the Server
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := 987 // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := 987 // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LabelApi.DatacentersServersLabelsGet(context.Background(), datacenterId, serverId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LabelApi.DatacentersServersLabelsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersServersLabelsGet`: LabelResources
    fmt.Fprintf(os.Stdout, "Response from `LabelApi.DatacentersServersLabelsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datacenterId** | **string** | The unique ID of the Datacenter | 
**serverId** | **string** | The unique ID of the Server | 

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersServersLabelsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**LabelResources**](LabelResources.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatacentersServersLabelsPost

> LabelResource DatacentersServersLabelsPost(ctx, datacenterId, serverId).Label(label).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()

Add a Label to Server



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    datacenterId := "datacenterId_example" // string | The unique ID of the Datacenter
    serverId := "serverId_example" // string | The unique ID of the Server
    label :=  // LabelResource | Label to be added
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := 987 // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := 987 // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LabelApi.DatacentersServersLabelsPost(context.Background(), datacenterId, serverId).Label(label).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LabelApi.DatacentersServersLabelsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersServersLabelsPost`: LabelResource
    fmt.Fprintf(os.Stdout, "Response from `LabelApi.DatacentersServersLabelsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datacenterId** | **string** | The unique ID of the Datacenter | 
**serverId** | **string** | The unique ID of the Server | 

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersServersLabelsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **label** | [**LabelResource**](LabelResource.md) | Label to be added | 
 **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**LabelResource**](LabelResource.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatacentersServersLabelsPut

> LabelResource DatacentersServersLabelsPut(ctx, datacenterId, serverId, key).Label(label).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()

Modify a Label of Server



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    datacenterId := "datacenterId_example" // string | The unique ID of the Datacenter
    serverId := "serverId_example" // string | The unique ID of the Server
    key := "key_example" // string | The key of the Label
    label :=  // LabelResource | Modified Label
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := 987 // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := 987 // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LabelApi.DatacentersServersLabelsPut(context.Background(), datacenterId, serverId, key).Label(label).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LabelApi.DatacentersServersLabelsPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersServersLabelsPut`: LabelResource
    fmt.Fprintf(os.Stdout, "Response from `LabelApi.DatacentersServersLabelsPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datacenterId** | **string** | The unique ID of the Datacenter | 
**serverId** | **string** | The unique ID of the Server | 
**key** | **string** | The key of the Label | 

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersServersLabelsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **label** | [**LabelResource**](LabelResource.md) | Modified Label | 
 **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**LabelResource**](LabelResource.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatacentersVolumesLabelsDelete

> map[string]interface{} DatacentersVolumesLabelsDelete(ctx, datacenterId, volumeId, key).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()

Delete a Label from Volume



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    datacenterId := "datacenterId_example" // string | The unique ID of the Datacenter
    volumeId := "volumeId_example" // string | The unique ID of the Volume
    key := "key_example" // string | The key of the Label
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := 987 // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := 987 // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LabelApi.DatacentersVolumesLabelsDelete(context.Background(), datacenterId, volumeId, key).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LabelApi.DatacentersVolumesLabelsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersVolumesLabelsDelete`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LabelApi.DatacentersVolumesLabelsDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datacenterId** | **string** | The unique ID of the Datacenter | 
**volumeId** | **string** | The unique ID of the Volume | 
**key** | **string** | The key of the Label | 

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersVolumesLabelsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

**map[string]interface{}**

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatacentersVolumesLabelsFindByKey

> LabelResource DatacentersVolumesLabelsFindByKey(ctx, datacenterId, volumeId, key).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()

Retrieve a Label of Volume



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    datacenterId := "datacenterId_example" // string | The unique ID of the Datacenter
    volumeId := "volumeId_example" // string | The unique ID of the Volume
    key := "key_example" // string | The key of the Label
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := 987 // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := 987 // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LabelApi.DatacentersVolumesLabelsFindByKey(context.Background(), datacenterId, volumeId, key).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LabelApi.DatacentersVolumesLabelsFindByKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersVolumesLabelsFindByKey`: LabelResource
    fmt.Fprintf(os.Stdout, "Response from `LabelApi.DatacentersVolumesLabelsFindByKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datacenterId** | **string** | The unique ID of the Datacenter | 
**volumeId** | **string** | The unique ID of the Volume | 
**key** | **string** | The key of the Label | 

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersVolumesLabelsFindByKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**LabelResource**](LabelResource.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatacentersVolumesLabelsGet

> LabelResources DatacentersVolumesLabelsGet(ctx, datacenterId, volumeId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()

List all Volume Labels



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    datacenterId := "datacenterId_example" // string | The unique ID of the Datacenter
    volumeId := "volumeId_example" // string | The unique ID of the Volume
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := 987 // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := 987 // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LabelApi.DatacentersVolumesLabelsGet(context.Background(), datacenterId, volumeId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LabelApi.DatacentersVolumesLabelsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersVolumesLabelsGet`: LabelResources
    fmt.Fprintf(os.Stdout, "Response from `LabelApi.DatacentersVolumesLabelsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datacenterId** | **string** | The unique ID of the Datacenter | 
**volumeId** | **string** | The unique ID of the Volume | 

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersVolumesLabelsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**LabelResources**](LabelResources.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatacentersVolumesLabelsPost

> LabelResource DatacentersVolumesLabelsPost(ctx, datacenterId, volumeId).Label(label).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()

Add a Label to Volume



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    datacenterId := "datacenterId_example" // string | The unique ID of the Datacenter
    volumeId := "volumeId_example" // string | The unique ID of the Volume
    label :=  // LabelResource | Label to be added
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := 987 // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := 987 // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LabelApi.DatacentersVolumesLabelsPost(context.Background(), datacenterId, volumeId).Label(label).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LabelApi.DatacentersVolumesLabelsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersVolumesLabelsPost`: LabelResource
    fmt.Fprintf(os.Stdout, "Response from `LabelApi.DatacentersVolumesLabelsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datacenterId** | **string** | The unique ID of the Datacenter | 
**volumeId** | **string** | The unique ID of the Volume | 

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersVolumesLabelsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **label** | [**LabelResource**](LabelResource.md) | Label to be added | 
 **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**LabelResource**](LabelResource.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatacentersVolumesLabelsPut

> LabelResource DatacentersVolumesLabelsPut(ctx, datacenterId, volumeId, key).Label(label).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()

Modify a Label of Volume



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    datacenterId := "datacenterId_example" // string | The unique ID of the Datacenter
    volumeId := "volumeId_example" // string | The unique ID of the Volume
    key := "key_example" // string | The key of the Label
    label :=  // LabelResource | Modified Label
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := 987 // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := 987 // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LabelApi.DatacentersVolumesLabelsPut(context.Background(), datacenterId, volumeId, key).Label(label).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LabelApi.DatacentersVolumesLabelsPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersVolumesLabelsPut`: LabelResource
    fmt.Fprintf(os.Stdout, "Response from `LabelApi.DatacentersVolumesLabelsPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datacenterId** | **string** | The unique ID of the Datacenter | 
**volumeId** | **string** | The unique ID of the Volume | 
**key** | **string** | The key of the Label | 

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersVolumesLabelsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **label** | [**LabelResource**](LabelResource.md) | Modified Label | 
 **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**LabelResource**](LabelResource.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IpblocksLabelsDelete

> map[string]interface{} IpblocksLabelsDelete(ctx, ipblockId, key).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()

Delete a Label from IP Block



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ipblockId := "ipblockId_example" // string | The unique ID of the Ip Block
    key := "key_example" // string | The key of the Label
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := 987 // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := 987 // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LabelApi.IpblocksLabelsDelete(context.Background(), ipblockId, key).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LabelApi.IpblocksLabelsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IpblocksLabelsDelete`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LabelApi.IpblocksLabelsDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ipblockId** | **string** | The unique ID of the Ip Block | 
**key** | **string** | The key of the Label | 

### Other Parameters

Other parameters are passed through a pointer to a apiIpblocksLabelsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

**map[string]interface{}**

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IpblocksLabelsFindByKey

> LabelResource IpblocksLabelsFindByKey(ctx, ipblockId, key).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()

Retrieve a Label of IP Block



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ipblockId := "ipblockId_example" // string | The unique ID of the Ip Block
    key := "key_example" // string | The key of the Label
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := 987 // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := 987 // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LabelApi.IpblocksLabelsFindByKey(context.Background(), ipblockId, key).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LabelApi.IpblocksLabelsFindByKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IpblocksLabelsFindByKey`: LabelResource
    fmt.Fprintf(os.Stdout, "Response from `LabelApi.IpblocksLabelsFindByKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ipblockId** | **string** | The unique ID of the Ip Block | 
**key** | **string** | The key of the Label | 

### Other Parameters

Other parameters are passed through a pointer to a apiIpblocksLabelsFindByKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**LabelResource**](LabelResource.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IpblocksLabelsGet

> LabelResources IpblocksLabelsGet(ctx, ipblockId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()

List all Ip Block Labels



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ipblockId := "ipblockId_example" // string | The unique ID of the Ip Block
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := 987 // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := 987 // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LabelApi.IpblocksLabelsGet(context.Background(), ipblockId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LabelApi.IpblocksLabelsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IpblocksLabelsGet`: LabelResources
    fmt.Fprintf(os.Stdout, "Response from `LabelApi.IpblocksLabelsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ipblockId** | **string** | The unique ID of the Ip Block | 

### Other Parameters

Other parameters are passed through a pointer to a apiIpblocksLabelsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**LabelResources**](LabelResources.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IpblocksLabelsPost

> LabelResource IpblocksLabelsPost(ctx, ipblockId).Label(label).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()

Add a Label to IP Block



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ipblockId := "ipblockId_example" // string | The unique ID of the Ip Block
    label :=  // LabelResource | Label to be added
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := 987 // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := 987 // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LabelApi.IpblocksLabelsPost(context.Background(), ipblockId).Label(label).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LabelApi.IpblocksLabelsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IpblocksLabelsPost`: LabelResource
    fmt.Fprintf(os.Stdout, "Response from `LabelApi.IpblocksLabelsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ipblockId** | **string** | The unique ID of the Ip Block | 

### Other Parameters

Other parameters are passed through a pointer to a apiIpblocksLabelsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **label** | [**LabelResource**](LabelResource.md) | Label to be added | 
 **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**LabelResource**](LabelResource.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IpblocksLabelsPut

> LabelResource IpblocksLabelsPut(ctx, ipblockId, key).Label(label).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()

Modify a Label of IP Block



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ipblockId := "ipblockId_example" // string | The unique ID of the Ip Block
    key := "key_example" // string | The key of the Label
    label :=  // LabelResource | Modified Label
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := 987 // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := 987 // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LabelApi.IpblocksLabelsPut(context.Background(), ipblockId, key).Label(label).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LabelApi.IpblocksLabelsPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IpblocksLabelsPut`: LabelResource
    fmt.Fprintf(os.Stdout, "Response from `LabelApi.IpblocksLabelsPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ipblockId** | **string** | The unique ID of the Ip Block | 
**key** | **string** | The key of the Label | 

### Other Parameters

Other parameters are passed through a pointer to a apiIpblocksLabelsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **label** | [**LabelResource**](LabelResource.md) | Modified Label | 
 **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**LabelResource**](LabelResource.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LabelsFindByLabelurn

> Label LabelsFindByLabelurn(ctx, labelurn).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()

Returns the label by its URN.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    labelurn := "labelurn_example" // string | The URN representing the unique ID of the label. A URN is for uniqueness of a Label and composed using urn:label:<resource_type>:<resource_uuid>:<key>
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := 987 // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := 987 // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LabelApi.LabelsFindByLabelurn(context.Background(), labelurn).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LabelApi.LabelsFindByLabelurn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LabelsFindByLabelurn`: Label
    fmt.Fprintf(os.Stdout, "Response from `LabelApi.LabelsFindByLabelurn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**labelurn** | **string** | The URN representing the unique ID of the label. A URN is for uniqueness of a Label and composed using urn:label:&lt;resource_type&gt;:&lt;resource_uuid&gt;:&lt;key&gt; | 

### Other Parameters

Other parameters are passed through a pointer to a apiLabelsFindByLabelurnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**Label**](Label.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LabelsGet

> Labels LabelsGet(ctx).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()

List Labels 



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := 987 // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := 987 // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LabelApi.LabelsGet(context.Background()).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LabelApi.LabelsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LabelsGet`: Labels
    fmt.Fprintf(os.Stdout, "Response from `LabelApi.LabelsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLabelsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**Labels**](Labels.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SnapshotsLabelsDelete

> map[string]interface{} SnapshotsLabelsDelete(ctx, snapshotId, key).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()

Delete a Label from Snapshot



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    snapshotId := "snapshotId_example" // string | The unique ID of the Snapshot
    key := "key_example" // string | The key of the Label
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := 987 // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := 987 // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LabelApi.SnapshotsLabelsDelete(context.Background(), snapshotId, key).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LabelApi.SnapshotsLabelsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SnapshotsLabelsDelete`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LabelApi.SnapshotsLabelsDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**snapshotId** | **string** | The unique ID of the Snapshot | 
**key** | **string** | The key of the Label | 

### Other Parameters

Other parameters are passed through a pointer to a apiSnapshotsLabelsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

**map[string]interface{}**

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SnapshotsLabelsFindByKey

> LabelResource SnapshotsLabelsFindByKey(ctx, snapshotId, key).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()

Retrieve a Label of Snapshot



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    snapshotId := "snapshotId_example" // string | The unique ID of the Snapshot
    key := "key_example" // string | The key of the Label
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := 987 // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := 987 // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LabelApi.SnapshotsLabelsFindByKey(context.Background(), snapshotId, key).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LabelApi.SnapshotsLabelsFindByKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SnapshotsLabelsFindByKey`: LabelResource
    fmt.Fprintf(os.Stdout, "Response from `LabelApi.SnapshotsLabelsFindByKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**snapshotId** | **string** | The unique ID of the Snapshot | 
**key** | **string** | The key of the Label | 

### Other Parameters

Other parameters are passed through a pointer to a apiSnapshotsLabelsFindByKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**LabelResource**](LabelResource.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SnapshotsLabelsGet

> LabelResources SnapshotsLabelsGet(ctx, snapshotId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()

List all Snapshot Labels



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    snapshotId := "snapshotId_example" // string | The unique ID of the Snapshot
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := 987 // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := 987 // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LabelApi.SnapshotsLabelsGet(context.Background(), snapshotId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LabelApi.SnapshotsLabelsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SnapshotsLabelsGet`: LabelResources
    fmt.Fprintf(os.Stdout, "Response from `LabelApi.SnapshotsLabelsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**snapshotId** | **string** | The unique ID of the Snapshot | 

### Other Parameters

Other parameters are passed through a pointer to a apiSnapshotsLabelsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**LabelResources**](LabelResources.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SnapshotsLabelsPost

> LabelResource SnapshotsLabelsPost(ctx, snapshotId).Label(label).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()

Add a Label to Snapshot



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    snapshotId := "snapshotId_example" // string | The unique ID of the Snapshot
    label :=  // LabelResource | Label to be added
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := 987 // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := 987 // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LabelApi.SnapshotsLabelsPost(context.Background(), snapshotId).Label(label).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LabelApi.SnapshotsLabelsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SnapshotsLabelsPost`: LabelResource
    fmt.Fprintf(os.Stdout, "Response from `LabelApi.SnapshotsLabelsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**snapshotId** | **string** | The unique ID of the Snapshot | 

### Other Parameters

Other parameters are passed through a pointer to a apiSnapshotsLabelsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **label** | [**LabelResource**](LabelResource.md) | Label to be added | 
 **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**LabelResource**](LabelResource.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SnapshotsLabelsPut

> LabelResource SnapshotsLabelsPut(ctx, snapshotId, key).Label(label).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()

Modify a Label of Snapshot



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    snapshotId := "snapshotId_example" // string | The unique ID of the Snapshot
    key := "key_example" // string | The key of the Label
    label :=  // LabelResource | Modified Label
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := 987 // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := 987 // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LabelApi.SnapshotsLabelsPut(context.Background(), snapshotId, key).Label(label).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LabelApi.SnapshotsLabelsPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SnapshotsLabelsPut`: LabelResource
    fmt.Fprintf(os.Stdout, "Response from `LabelApi.SnapshotsLabelsPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**snapshotId** | **string** | The unique ID of the Snapshot | 
**key** | **string** | The key of the Label | 

### Other Parameters

Other parameters are passed through a pointer to a apiSnapshotsLabelsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **label** | [**LabelResource**](LabelResource.md) | Modified Label | 
 **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**LabelResource**](LabelResource.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

