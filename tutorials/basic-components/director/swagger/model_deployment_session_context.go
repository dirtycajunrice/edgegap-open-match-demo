/*
 * Codema
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.41.174.post.dev2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type DeploymentSessionContext struct {
	// Unique UUID
	SessionId string `json:"session_id"`
	// Current status of the session
	Status string `json:"status"`
	// If the session is linked to a Ready deployment
	Ready bool `json:"ready"`
	// If the session is linked to a deployment
	Linked bool `json:"linked"`
	// Type of session created
	Kind string `json:"kind"`
	// Count of user this session currently have
	UserCount int32 `json:"user_count"`
}