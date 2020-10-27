/*
 * CLOUD API
 *
 * An enterprise-grade Infrastructure is provided as a Service (IaaS) solution that can be managed through a browser-based \"Data Center Designer\" (DCD) tool or via an easy to use API.   The API allows you to perform a variety of management tasks such as spinning up additional servers, adding volumes, adjusting networking, and so forth. It is designed to allow users to leverage the same power and flexibility found within the DCD visual tool. Both tools are consistent with their concepts and lend well to making the experience smooth and intuitive.
 *
 * API version: 5.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package ionossdk
// KubernetesMaintenanceWindow struct for KubernetesMaintenanceWindow
type KubernetesMaintenanceWindow struct {
	// The day of the week for a maintenance window.
	DayOfTheWeek *string `json:"dayOfTheWeek,omitempty"` /* nullable: false */
	// The time to use for a maintenance window. Accepted formats are: HH:mm:ss; HH:mm:ss\"Z\"; HH:mm:ssZ. This time may varies by 15 minutes.
	Time *string `json:"time,omitempty"` /* nullable: false */
}