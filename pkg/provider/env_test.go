package provider

import (
	"slices"
	"strings"
	"testing"

	"github.com/skevetter/devpod/pkg/config"
)

func TestToEnvironment_ProviderNameDoesNotCollideWithFlag(t *testing.T) {
	result := ToEnvironment(nil, nil, nil, nil, "my-provider")

	for _, entry := range result {
		key := strings.SplitN(entry, "=", 2)[0]
		if key == "DEVPOD_PROVIDER" {
			t.Error(
				"found DEVPOD_PROVIDER in env; use DEVPOD_PROVIDER_NAME instead",
			)
		}
	}

	expected := config.EnvProviderName + "=my-provider"
	if !slices.Contains(result, expected) {
		t.Errorf(
			"expected %s=my-provider in environment, but it was not found",
			config.EnvProviderName,
		)
	}
}

func TestToEnvironment_EmptyProviderNameOmitted(t *testing.T) {
	result := ToEnvironment(nil, nil, nil, nil, "")

	for _, entry := range result {
		key := strings.SplitN(entry, "=", 2)[0]
		if key == config.EnvProviderName {
			t.Errorf(
				"expected %s to be absent when providerName is empty, but found it",
				config.EnvProviderName,
			)
		}
	}
}
