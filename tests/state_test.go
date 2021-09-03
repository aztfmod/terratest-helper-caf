package caf_helper_tests

import (
	//"fmt"
	"testing"

	"github.com/aztfmod/terratest-helper-caf/state"
	"github.com/stretchr/testify/assert"
)

func TestLandingZoneKey(t *testing.T) {
	t.Parallel()
	tfState := state.NewTerraformState(t, "gears")

	landingZoneKey := tfState.GetLandingZoneKey()

	assert.Equal(t, "gears", landingZoneKey)
}

func TestResourceGroups(t *testing.T) {
	t.Parallel()
	tfState := state.NewTerraformState(t, "gears")
	resourceGroups := tfState.GetResourceGroups()

	assert.NotEmpty(t,resourceGroups)
}

func TestSQLServers(t *testing.T) {
	t.Parallel()
	tfState := state.NewTerraformState(t, "gears")
	sqlServers := tfState.GetSQLServers()

	assert.NotEmpty(t,sqlServers)
}

func TestAppInsights(t *testing.T) {
	t.Parallel()
	tfState := state.NewTerraformState(t, "gears")
	appInsights := tfState.GetAppInsights()

	assert.NotEmpty(t,appInsights)
}

func TestStorageAccountQueues(t *testing.T) {
	t.Parallel()
	tfState := state.NewTerraformState(t, "gears")
	storageAccountQueues := tfState.GetStorageAccountQueues()

	assert.NotEmpty(t,storageAccountQueues)
}

func TestRandomStrings(t *testing.T) {
	t.Parallel()
	tfState := state.NewTerraformState(t, "gears")
	randomStrings := tfState.GetRandomStrings()

	assert.NotEmpty(t,randomStrings)
}