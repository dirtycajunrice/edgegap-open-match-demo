/*
 * Codema
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.41.174.post.dev2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type TotalMetricsModel struct {
	ReceiveTotal   *MetricsModel `json:"receive_total,omitempty"`
	TransmitTotal  *MetricsModel `json:"transmit_total,omitempty"`
	DiskReadTotal  *MetricsModel `json:"disk_read_total,omitempty"`
	DiskWriteTotal *MetricsModel `json:"disk_write_total,omitempty"`
}
