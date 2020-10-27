/*
 * CLOUD API
 *
 * An enterprise-grade Infrastructure is provided as a Service (IaaS) solution that can be managed through a browser-based \"Data Center Designer\" (DCD) tool or via an easy to use API.   The API allows you to perform a variety of management tasks such as spinning up additional servers, adding volumes, adjusting networking, and so forth. It is designed to allow users to leverage the same power and flexibility found within the DCD visual tool. Both tools are consistent with their concepts and lend well to making the experience smooth and intuitive.
 *
 * API version: 5.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package ionossdk
// ResourceGroups Resources assigned to this group.
type ResourceGroups struct {
	// The resource's unique identifier
	Id *string `json:"id,omitempty"` /* nullable: false */
	// The type of the resource
	Type *Type `json:"type,omitempty"` /* nullable: false */
	// URL to the object representation (absolute path)
	Href *string `json:"href,omitempty"` /* nullable: false */
	// Array of items in that collection
	Items *[]Resource `json:"items,omitempty"` /* nullable: false */
}
