package provider

import (
	"strings"
	"testing"

	"github.com/skevetter/devpod/pkg/config"
)

func TestToEnvironment_ProviderNameDoesNotCollideWithFlag(t *testing.T) {
	result := ToEnvironment(nil, nil, nil, nil, "my-provider")

	for _, entry := range result {
		key := strings.SplitN(entry, "=", 2)[0]
		if key == "DEVPOD_PROVIDER" {
			t.Errorf("found env var with key DEVPOD_PROVIDER which collides with the --provider CLI flag; use DEVPOD_PROVIDER_NAME instead")
		}
	}

	found := false
	for _, entry := range result {
		if entry == config.EnvProviderName+"=my-provider" {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("expected %s=my-provider in environment, but it was not found", config.EnvProviderName)
	}
}

func TestToEnvironment_EmptyProviderNameOmitted(t *testing.T) {
	result := ToEnvironment(nil, nil, nil, nil, "")

	for _, entry := range result {
		key := strings.SplitN(entry, "=", 2)[0]
		if key == config.EnvProviderName {
			t.Errorf("expected %s to be absent when providerName is empty, but found it", config.EnvProviderName)
		}
	}
}
