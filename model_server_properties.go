/*
 * CLOUD API
 *
 * An enterprise-grade Infrastructure is provided as a Service (IaaS) solution that can be managed through a browser-based \"Data Center Designer\" (DCD) tool or via an easy to use API.   The API allows you to perform a variety of management tasks such as spinning up additional servers, adding volumes, adjusting networking, and so forth. It is designed to allow users to leverage the same power and flexibility found within the DCD visual tool. Both tools are consistent with their concepts and lend well to making the experience smooth and intuitive.
 *
 * API version: 5.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package ionossdk
// ServerProperties struct for ServerProperties
type ServerProperties struct {
	// A name of that resource
	Name *string `json:"name,omitempty"` /* nullable: false */
	// The total number of cores for the server
	Cores *int32 `json:"cores,omitempty"` /* nullable: false */
	// The amount of memory for the server in MB, e.g. 2048. Size must be specified in multiples of 256 MB with a minimum of 256 MB; however, if you set ramHotPlug to TRUE then you must use a minimum of 1024 MB. If you set the RAM size more than 240GB, then ramHotPlug will be set to FALSE and can not be set to TRUE unless RAM size not set to less than 240GB.
	Ram *int32 `json:"ram,omitempty"` /* nullable: false */
	// The availability zone in which the server should exist
	AvailabilityZone *string `json:"availabilityZone,omitempty"` /* nullable: false */
	// Status of the virtual Machine
	VmState *string `json:"vmState,omitempty"` /* nullable: false */
	BootCdrom *ResourceReference `json:"bootCdrom,omitempty"` /* nullable: false */
	BootVolume *ResourceReference `json:"bootVolume,omitempty"` /* nullable: false */
	// Cpu family of pserver
	CpuFamily *string `json:"cpuFamily,omitempty"` /* nullable: false */
}
