/*
 * Codema
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.41.174.post.dev2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Status struct {
	// The Unique ID of the Request
	RequestId string `json:"request_id"`
	// The FQDN that allow to connect to your instance
	Fqdn string `json:"fqdn"`
	// The name of the deployed App
	AppName string `json:"app_name"`
	// The version of the deployed App
	AppVersion string `json:"app_version"`
	// The Current Status of the Request
	CurrentStatus string `json:"current_status"`
	// True if the Current Request is Ready to be connected and running
	Running bool `json:"running"`
	// True if the Current Request is ACL Protected
	WhitelistingActive bool `json:"whitelisting_active"`
	// Timestamp of the beginning of the Request
	StartTime string `json:"start_time"`
	// Time since the Request is up and running in seconds
	ElapsedTime int32 `json:"elapsed_time"`
	// The Last Status of the Request
	LastStatus string `json:"last_status,omitempty"`
	// True if there is an Error with the Request
	Error_ bool `json:"error"`
	// The Error detail of the Request
	ErrorDetail string                 `json:"error_detail,omitempty"`
	Ports       map[string]PortMapping `json:"ports,omitempty"`
	// The public IP
	PublicIp string `json:"public_ip"`
	// List of Active Sessions if Deployment App is Session Based
	Sessions []DeploymentSessionContext `json:"sessions,omitempty"`
	// Location related information
	Location *DeploymentLocation `json:"location,omitempty"`
	// List of tags associated with the deployment
	Tags []string `json:"tags,omitempty"`
	// The Capacity of the Deployment
	Sockets int32 `json:"sockets,omitempty"`
	// The Capacity Usage of the Deployment
	SocketsUsage int32 `json:"sockets_usage,omitempty"`
}
