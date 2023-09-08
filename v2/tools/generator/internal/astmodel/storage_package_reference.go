/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

const (
	StorageSubPackage = "storage"
)

// MakeStoragePackageReference creates a new storage package reference from a local package reference
func MakeStoragePackageReference(local LocalPackageReference) InternalPackageReference {
	return MakeSubPackageReference(StorageSubPackage, local)
}

// IsStoragePackageReference returns true if the reference is to a storage package
func IsStoragePackageReference(reference PackageReference) bool {
	if sub, ok := reference.(SubPackageReference); ok {
		if sub.name == StorageSubPackage {
			return true
		}

		return IsStoragePackageReference(sub.parent)
	}

	return false
}
