# Terratest helper for Azure Terraform SRE

> :warning: This solution, offered by the Open-Source community, will no longer receive contributions from Microsoft.

This library contains helper methods that are useful when testing infrastructure deployed with [Azure Terraform SRE.](https://github.com/aztfmod/caf-terraform-landingzones)

The  helper methods provide an easy way to extract resource information from terraform state files. This makes easy to create dynamic tests that source resource names from Terraform state.

Example:

```go

import "github.com/aztfmod/terratest-helper-caf/state"

func TestLaunchpadResourceGroupIsExists(t *testing.T) {
  t.Parallel()
  tfState := state.NewTerraformState(t, "launchpad")
  resourceGroups := tfState.GetResourceGroups()

  for _, resourceGroup := range resourceGroups {
    name := resourceGroup.GetName()
    exists := azure.ResourceGroupExists(t, name, tfState.SubscriptionID)
    assert.True(t, exists, fmt.Sprintf("Resource group (%s) does not exist", name))
  }
}

```

## Getting Started

* `import "github.com/aztfmod/terratest-helper-caf/state"`
* `go get`
* Create CAF Tests

## Azure Terraform SRE Test structure

Azure Terraform SRE Tests consist of reading resource information from the state file and then using Terratest helpers to assert the existence of resource and the correct configuration.

In order bo build a test, start with definting a new tfState object by invoking

```go
tfState := state.NewTerraformState(t, "launchpad")
```

The second parameter is the [landing zone key](https://github.com/Azure/caf-terraform-landingzones-starter/blob/starter/configuration/demo/level0/launchpad/configuration.tfvars#L4), defined in the CAF configuration for the specific landing zone.

After creating a tfState object, it is possible to call a helper function to return a specific resource type from the state file.

eg.

```go
  resourceGroups := tfState.GetResourceGroups()

  for _, resourceGroup := range resourceGroups {
    name := resourceGroup.GetName()
    fmt.Println(name)
  }
```

There are two options for running the tests.

### Option 1 - rover

Run the rover test command:

```shell
rover test \
      -b <path to tests folder> \
      -env <environment name> \
      -level <level> \
      -tfstate <name of deployed state file> \
      -d
```

### Option 2 - Debugging or cli invocation of go test

* Download your state file from azure blob storage.
* Rename the file to terraform.tfstate and place it in a known location.
* create a .env file in the tests folder with the following structure

```shell
 STATE_FILE_PATH=<path to state file>
 ENVIRONMENT=<deployed caf environment>
 ARM_SUBSCRIPTION_ID=<subscription id>
```

Note: if running through rover, you don't have to set these as rover test exports the values.

When running through go test you must specifiy the build tag and ensure that the correct state file is loaded into STATE_FILE_PATH.

eg
`go test -tags level0 -v
`

You can only test one level at a time when using option 2.

## Helper Functions

| function      | Description |
| ----------- | ----------- |
| GetResources      |  Returns a list of all resources in the state file.       |
| GetClientConfig   | Returns the CAF client config object.        |
| GetGlobalSettings   |Returns the landing zone key defined in the CAF tfvars configuration for the landing zone.        |
| GetResourceGroups   | Returns all deployed resource groups in the current state file.        |
| GetKeyVaults   |  Returns all deployed Azure KeyVaults in the current state file.        |
| GetRecoveryVaults   | Returns all deployed Recovery Vaults in the current state file.        |
| GetKeyVaultByResourceGroup   | Returns CAF KeyVault in a specific Resource Group.        |
| GetStorageAccounts   | Returns all deployed Azure Storage Accounts in the current state file.        |
| GetStorageAccountByResourceGroup   | Returns all storage accounts in a resource group.       |
| GetVNets      |  Returns all deployed VNETS in the current state file.       |
| GetAKSClusters      |  Returns all deployed AKS Clusters in the current state file.       |
| GetSQLServers      |   Returns all deployed Sql Servers in the current state file.       |
| GetSQLDBs      | Returns all deployed Sql Server Databases in the current state file.       |
| GetAppServices      |  Returns all deployed App Services in the current state file.       |
| GetAppInsights      |  Returns all deployed App Insights in the current state file.       |
| GetStorageAccountQueues      |  Returns all deployed Storage Account Queues in the current state file.      |
| GetMachineLearningWorkspaces      |  Returns all deployed Machine Learning Workspaces in the current state file.      |
| GetAzureContainerRegistries      | Returns all deployed Azure Container Registries in the current state file.      |

* Note : This list of helper functions will continue to grow. However, they are utility functions and meant to accelerate development. It is possible to write tests with only using GetResources that returns a list of all deployed resources. This can then be further filtered by desired resource type.

## Example Tests

For examples of writing Tests, please see the [Symphony Reference application tests](https://github.com/aztfmod/symphony/tree/master/tests
)
References:

* [Terratest](https://github.com/gruntwork-io/terratest)
* [Azure Terraform SRE](https://github.com/aztfmod/caf-terraform-landingzones)
