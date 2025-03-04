/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package config

import (
	"testing"

	. "github.com/onsi/gomega"
	"gopkg.in/yaml.v3"
)

func TestTypeConfiguration_WhenYAMLWellFormed_ReturnsExpectedResult(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	yamlBytes := loadTestData(t)

	var typeConfig TypeConfiguration
	err := yaml.Unmarshal(yamlBytes, &typeConfig)
	g.Expect(err).To(Succeed())
	g.Expect(typeConfig.properties).To(HaveLen(4))

	name, ok := typeConfig.NameInNextVersion.read()
	g.Expect(name).To(Equal("Demo"))
	g.Expect(ok).To(BeTrue())

	export, ok := typeConfig.Export.read()
	g.Expect(export).To(BeTrue())
	g.Expect(ok).To(BeTrue())

	exportAs, ok := typeConfig.ExportAs.read()
	g.Expect(exportAs).To(Equal("Demo"))
	g.Expect(ok).To(BeTrue())

	azureGeneratedSecrets, ok := typeConfig.AzureGeneratedSecrets.read()
	g.Expect(azureGeneratedSecrets).To(HaveLen(2))
	g.Expect(ok).To(BeTrue())

	supportedFrom, ok := typeConfig.SupportedFrom.read()
	g.Expect(supportedFrom).To(Equal("beta.3"))
	g.Expect(ok).To(BeTrue())
}

func TestTypeConfiguration_WhenYAMLBadlyFormed_ReturnsError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	yamlBytes := loadTestData(t)

	var typeConfig TypeConfiguration
	err := yaml.Unmarshal(yamlBytes, &typeConfig)
	g.Expect(err).NotTo(Succeed())
}

func TestTypeConfiguration_WhenAzureSecretsBadlyFormed_ReturnsError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	yamlBytes := loadTestData(t)

	var typeConfig TypeConfiguration
	err := yaml.Unmarshal(yamlBytes, &typeConfig)
	g.Expect(err).NotTo(Succeed())
	g.Expect(err.Error()).To((ContainSubstring(azureGeneratedSecretsTag)))
}

/*
 * Type Rename tests
 */

func TestTypeConfiguration_TypeRename_WhenRenameConfigured_ReturnsExpectedResult(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	typeConfig := NewTypeConfiguration("Person")
	typeConfig.NameInNextVersion.Set("Address")

	name, err := typeConfig.NameInNextVersion.Lookup()

	g.Expect(name).To(Equal("Address"))
	g.Expect(err).To(Succeed())
}

func TestTypeConfiguration_TypeRename_WhenRenameNotConfigured_ReturnsExpectedResult(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	typeConfig := NewTypeConfiguration("Person")

	name, err := typeConfig.NameInNextVersion.Lookup()
	g.Expect(name).To(Equal(""))
	g.Expect(err).NotTo(Succeed())
	g.Expect(err.Error()).To(ContainSubstring(typeConfig.name))
	g.Expect(err.Error()).To(ContainSubstring(nameInNextVersionTag))
}

func TestTypeConfiguration_VerifyTypeRenameConsumed_WhenRenameUsed_ReturnsNoError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	typeConfig := NewTypeConfiguration("Person")
	typeConfig.NameInNextVersion.Set("Party")

	_, err := typeConfig.NameInNextVersion.Lookup()
	g.Expect(err).To(Succeed())
	g.Expect(typeConfig.NameInNextVersion.VerifyConsumed()).To(Succeed())
}

func TestTypeConfiguration_VerifyTypeRenameConsumed_WhenRenameUnused_ReturnsExpectedError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	typeConfig := NewTypeConfiguration("Person")
	typeConfig.NameInNextVersion.Set("Party")

	err := typeConfig.NameInNextVersion.VerifyConsumed()
	g.Expect(err).NotTo(BeNil())
	g.Expect(err.Error()).To(ContainSubstring(typeConfig.name))
}

/*
 * SupportedFrom tests
 */

func TestTypeConfiguration_LookupSupportedFrom_WhenConfigured_ReturnsExpectedResult(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	typeConfig := NewTypeConfiguration("Person")
	typeConfig.SupportedFrom.Set("beta.0")

	from, err := typeConfig.SupportedFrom.Lookup()

	g.Expect(from).To(Equal("beta.0"))
	g.Expect(err).To(Succeed())
}

func TestTypeConfiguration_LookupSupportedFrom_WhenNotConfigured_ReturnsExpectedError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	typeConfig := NewTypeConfiguration("Person")

	name, err := typeConfig.SupportedFrom.Lookup()
	g.Expect(name).To(Equal(""))
	g.Expect(err).NotTo(Succeed())
	g.Expect(err.Error()).To(ContainSubstring(typeConfig.name))
	g.Expect(err.Error()).To(ContainSubstring(supportedFromTag))
}

func TestTypeConfiguration_VerifySupportedFromConsumed_WhenConsumed_ReturnsNoError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	typeConfig := NewTypeConfiguration("Person")
	typeConfig.SupportedFrom.Set("beta.0")

	_, err := typeConfig.SupportedFrom.Lookup()
	g.Expect(err).To(Succeed())
	g.Expect(typeConfig.SupportedFrom.VerifyConsumed()).To(Succeed())
}

func TestTypeConfiguration_VerifySupportedFromConsumed_WhenNotConsumed_ReturnsExpectedError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	typeConfig := NewTypeConfiguration("Person")
	typeConfig.SupportedFrom.Set("beta.0")

	err := typeConfig.SupportedFrom.VerifyConsumed()
	g.Expect(err).NotTo(BeNil())
	g.Expect(err.Error()).To(ContainSubstring(typeConfig.name))
}
