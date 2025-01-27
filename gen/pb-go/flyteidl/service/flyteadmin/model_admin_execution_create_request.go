/*
 * flyteidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package flyteadmin

// Request to launch an execution with the given project, domain and optionally name.
type AdminExecutionCreateRequest struct {
	// Name of the project the execution belongs to.
	Project string `json:"project,omitempty"`
	// Name of the domain the execution belongs to. A domain can be considered as a subset within a specific project.
	Domain string `json:"domain,omitempty"`
	Name string `json:"name,omitempty"`
	// Additional fields necessary to launch the execution.
	Spec *AdminExecutionSpec `json:"spec,omitempty"`
	// The inputs required to start the execution. All required inputs must be included in this map. If not required and not provided, defaults apply.
	Inputs *CoreLiteralMap `json:"inputs,omitempty"`
}
