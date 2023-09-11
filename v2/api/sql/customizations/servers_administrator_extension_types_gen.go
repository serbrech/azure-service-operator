// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package customizations

import (
	v20211101 "github.com/Azure/azure-service-operator/v2/api/sql/v1api20211101"
	v20211101s "github.com/Azure/azure-service-operator/v2/api/sql/v1api20211101/storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type ServersAdministratorExtension struct {
}

// GetExtendedResources Returns the KubernetesResource slice for Resource versions
func (extension *ServersAdministratorExtension) GetExtendedResources() []genruntime.KubernetesResource {
	return []genruntime.KubernetesResource{
		&v20211101.ServersAdministrator{},
		&v20211101s.ServersAdministrator{}}
}
