// Copyright (c) The OpenTofu Authors
// SPDX-License-Identifier: MPL-2.0
// Copyright (c) 2023 HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package configload

import (
	"os"
	"testing"
)

// NewLoaderForTests is a variant of NewLoader that is intended to be more
// convenient for unit tests.
//
// The loader's modules directory is a separate temporary directory created
// for each call. Along with the created loader, this function returns a
// cleanup function that should be called before the test completes in order
// to remove that temporary directory.
//
// In the case of any errors, t.Fatal (or similar) will be called to halt
// execution of the test, so the calling test does not need to handle errors
// itself.
func NewLoaderForTests(t testing.TB) (*Loader, func()) {
	t.Helper()

	modulesDir, err := os.MkdirTemp("", "tf-configs")
	if err != nil {
		t.Fatalf("failed to create temporary modules dir: %s", err)
		return nil, func() {}
	}

	cleanup := func() {
		os.RemoveAll(modulesDir)
	}

	loader, err := NewLoader(&Config{
		ModulesDir: modulesDir,
	})
	if err != nil {
		cleanup()
		t.Fatalf("failed to create config loader: %s", err)
		return nil, func() {}
	}

	return loader, cleanup
}
