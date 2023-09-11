// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package customizations

import (
	v20180601 "github.com/Azure/azure-service-operator/v2/api/datafactory/v1api20180601"
	v20180601s "github.com/Azure/azure-service-operator/v2/api/datafactory/v1api20180601/storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type FactoryExtension struct {
}

// GetExtendedResources Returns the KubernetesResource slice for Resource versions
func (extension *FactoryExtension) GetExtendedResources() []genruntime.KubernetesResource {
	return []genruntime.KubernetesResource{
		&v20180601.Factory{},
		&v20180601s.Factory{}}
}
