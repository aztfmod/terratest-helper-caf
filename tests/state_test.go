package caf_helper_tests

import (
	"fmt"
	"testing"

	"github.com/aztfmod/terratest-helper-caf/state"
	"github.com/stretchr/testify/assert"
)

func TestMyTest(t *testing.T){
	t.Parallel()
	fmt.Println("A Test")
	
}

func TestLandingZoneKey(t *testing.T) {
	t.Parallel()
	tfState := state.NewTerraformState(t, "gears")

	landingZoneKey := tfState.GetLandingZoneKey()

	fmt.Println(landingZoneKey)
	assert.Equal(t, "gears", landingZoneKey)
}

func TestResourceGroups(t *testing.T) {
	t.Parallel()
	tfState := state.NewTerraformState(t, "gears")
	resourceGroups := tfState.GetResourceGroups()

	for _, resourceGroup := range resourceGroups {
		name := resourceGroup.GetName()
		fmt.Println(name)
	}
}

func TestSQLServers(t *testing.T) {
	t.Parallel()
	tfState := state.NewTerraformState(t, "gears")
	sqlServers := tfState.GetSQLServers()

	for _, sqlServer := range sqlServers {
		name := sqlServer.GetName()
		fmt.Println(name)
	}
}

func TestAppInsights(t *testing.T) {
	t.Parallel()
	tfState := state.NewTerraformState(t, "gears")
	appInsights := tfState.GetAppInsights()

	for _, appInsight := range appInsights {
		id := appInsight.GetID()
		fmt.Println(id)
	}
}