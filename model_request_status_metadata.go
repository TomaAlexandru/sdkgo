/*
 * CLOUD API
 *
 * An enterprise-grade Infrastructure is provided as a Service (IaaS) solution that can be managed through a browser-based \"Data Center Designer\" (DCD) tool or via an easy to use API.   The API allows you to perform a variety of management tasks such as spinning up additional servers, adding volumes, adjusting networking, and so forth. It is designed to allow users to leverage the same power and flexibility found within the DCD visual tool. Both tools are consistent with their concepts and lend well to making the experience smooth and intuitive.
 *
 * API version: 5.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package ionossdk
// RequestStatusMetadata struct for RequestStatusMetadata
type RequestStatusMetadata struct {
	Status *string `json:"status,omitempty"` /* nullable: false */
	Message *string `json:"message,omitempty"` /* nullable: false */
	// Resource's Entity Tag as defined in http://www.w3.org/Protocols/rfc2616/rfc2616-sec3.html#sec3.11 . Entity Tag is also added as an 'ETag response header to requests which don't use 'depth' parameter. 
	Etag *string `json:"etag,omitempty"` /* nullable: false */
	Targets *[]RequestTarget `json:"targets,omitempty"` /* nullable: false */
}
