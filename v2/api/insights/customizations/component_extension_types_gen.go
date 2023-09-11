// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package customizations

import (
	v20200202 "github.com/Azure/azure-service-operator/v2/api/insights/v1api20200202"
	v20200202s "github.com/Azure/azure-service-operator/v2/api/insights/v1api20200202/storage"
	v1beta20200202 "github.com/Azure/azure-service-operator/v2/api/insights/v1beta20200202"
	v1beta20200202s "github.com/Azure/azure-service-operator/v2/api/insights/v1beta20200202/storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type ComponentExtension struct {
}

// GetExtendedResources Returns the KubernetesResource slice for Resource versions
func (extension *ComponentExtension) GetExtendedResources() []genruntime.KubernetesResource {
	return []genruntime.KubernetesResource{
		&v1beta20200202.Component{},
		&v1beta20200202s.Component{},
		&v20200202.Component{},
		&v20200202s.Component{}}
}
