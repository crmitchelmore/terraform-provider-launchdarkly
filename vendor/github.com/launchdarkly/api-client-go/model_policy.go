/*
 * LaunchDarkly REST API
 *
 * Build custom integrations with the LaunchDarkly REST API
 *
 * API version: 2.0.15
 * Contact: support@launchdarkly.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package ldapi

type Policy struct {
	Resources []string `json:"resources,omitempty"`
	Actions []string `json:"actions,omitempty"`
	// Effect of the policy - allow or deny.
	Effect string `json:"effect,omitempty"`
}
