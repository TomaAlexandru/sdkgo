/*
 * CLOUD API
 *
 * An enterprise-grade Infrastructure is provided as a Service (IaaS) solution that can be managed through a browser-based \"Data Center Designer\" (DCD) tool or via an easy to use API.   The API allows you to perform a variety of management tasks such as spinning up additional servers, adding volumes, adjusting networking, and so forth. It is designed to allow users to leverage the same power and flexibility found within the DCD visual tool. Both tools are consistent with their concepts and lend well to making the experience smooth and intuitive.
 *
 * API version: 5.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package ionossdk
// KubernetesNodePools struct for KubernetesNodePools
type KubernetesNodePools struct {
	// Unique representation for Kubernetes Node Pool as a collection on a resource.
	Id *string `json:"id,omitempty"` /* nullable: false */
	// The type of resource within a collection
	Type *string `json:"type,omitempty"` /* nullable: false */
	// URL to the collection representation (absolute path)
	Href *string `json:"href,omitempty"` /* nullable: false */
	// Array of items in that collection
	Items *[]KubernetesNodePool `json:"items,omitempty"` /* nullable: false */
}
