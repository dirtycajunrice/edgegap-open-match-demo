/*
 * Codema
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.41.174.post.dev2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type AppVersions struct {
	Versions   []AppVersion `json:"versions,omitempty"`
	TotalCount int32        `json:"total_count,omitempty"`
}