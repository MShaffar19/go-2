/*
 * Kubernetes
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v1.10.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

// ExternalMetricStatus indicates the current value of a global metric not associated with any Kubernetes object.
type V2beta1ExternalMetricStatus struct {

	// currentAverageValue is the current value of metric averaged over autoscaled pods.
	CurrentAverageValue string `json:"currentAverageValue,omitempty"`

	// currentValue is the current value of the metric (as a quantity)
	CurrentValue string `json:"currentValue"`

	// metricName is the name of a metric used for autoscaling in metric system.
	MetricName string `json:"metricName"`

	// metricSelector is used to identify a specific time series within a given metric.
	MetricSelector *V1LabelSelector `json:"metricSelector,omitempty"`
}
