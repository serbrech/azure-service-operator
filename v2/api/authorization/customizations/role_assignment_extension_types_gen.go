// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package customizations

import (
	v20200801p "github.com/Azure/azure-service-operator/v2/api/authorization/v1api20200801preview"
	v20200801ps "github.com/Azure/azure-service-operator/v2/api/authorization/v1api20200801preview/storage"
	v1beta20200801p "github.com/Azure/azure-service-operator/v2/api/authorization/v1beta20200801preview"
	v1beta20200801ps "github.com/Azure/azure-service-operator/v2/api/authorization/v1beta20200801preview/storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type RoleAssignmentExtension struct {
}

// GetExtendedResources Returns the KubernetesResource slice for Resource versions
func (extension *RoleAssignmentExtension) GetExtendedResources() []genruntime.KubernetesResource {
	return []genruntime.KubernetesResource{
		&v1beta20200801p.RoleAssignment{},
		&v1beta20200801ps.RoleAssignment{},
		&v20200801p.RoleAssignment{},
		&v20200801ps.RoleAssignment{}}
}
