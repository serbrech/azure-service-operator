// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package customizations

import (
	v20220501 "github.com/Azure/azure-service-operator/v2/api/appconfiguration/v1api20220501"
	v20220501s "github.com/Azure/azure-service-operator/v2/api/appconfiguration/v1api20220501/storage"
	v1beta20220501 "github.com/Azure/azure-service-operator/v2/api/appconfiguration/v1beta20220501"
	v1beta20220501s "github.com/Azure/azure-service-operator/v2/api/appconfiguration/v1beta20220501/storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type ConfigurationStoreExtension struct {
}

// GetExtendedResources Returns the KubernetesResource slice for Resource versions
func (extension *ConfigurationStoreExtension) GetExtendedResources() []genruntime.KubernetesResource {
	return []genruntime.KubernetesResource{
		&v1beta20220501.ConfigurationStore{},
		&v1beta20220501s.ConfigurationStore{},
		&v20220501.ConfigurationStore{},
		&v20220501s.ConfigurationStore{}}
}
