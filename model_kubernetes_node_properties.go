/*
 * CLOUD API
 *
 * An enterprise-grade Infrastructure is provided as a Service (IaaS) solution that can be managed through a browser-based \"Data Center Designer\" (DCD) tool or via an easy to use API.   The API allows you to perform a variety of management tasks such as spinning up additional servers, adding volumes, adjusting networking, and so forth. It is designed to allow users to leverage the same power and flexibility found within the DCD visual tool. Both tools are consistent with their concepts and lend well to making the experience smooth and intuitive.
 *
 * API version: 5.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package ionossdk
// KubernetesNodeProperties struct for KubernetesNodeProperties
type KubernetesNodeProperties struct {
	// A Kubernetes Node Name.
	Name *string `json:"name,omitempty"` /* nullable: false */
	// A valid public IP.
	PublicIP *string `json:"publicIP,omitempty"` /* nullable: false */
	// The kubernetes version in which a nodepool is running. This imposes restrictions on what kubernetes versions can be run in a cluster's nodepools. Additionally, not all kubernetes versions are viable upgrade targets for all prior versions.
	K8sVersion *string `json:"k8sVersion,omitempty"` /* nullable: false */
}
