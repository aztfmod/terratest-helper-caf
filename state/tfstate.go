package state

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/logger"
	terraform "github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/require"
)

type Resource = map[string]interface{}
type AzureResource struct {
	Resource Resource
}

type ResourceGroups = map[string](AzureResource)
type KeyVaults = map[string](AzureResource)
type StorageAccounts = map[string](AzureResource)
type RecoveryVaults = map[string](AzureResource)
type VNets = map[string](AzureResource)
type AKSClusters = map[string](AzureResource)
type SQLServers = map[string](AzureResource)
type SQLDBs = map[string](AzureResource)
type AppServices = map[string](AzureResource)
type AppInsights = map[string](AzureResource)
type StorageAccountQueues = map[string](AzureResource)
type RandomStrings = map[string](AzureResource)
type MachineLearningWorkspaces = map[string](AzureResource)
type AzureContainerRegistries = map[string](AzureResource)

type TerraFormState struct {
	Objects        Resource
	SubscriptionID string
	Environment    string
	Key            string
}

var TfState TerraFormState

func ReadTerraformOutputJson(t *testing.T, Options *terraform.Options, key string) string {
	args := []string{"output", "-no-color", "-json"}
	args = append(args, key)

	Options.Logger = logger.Discard
	output, err := terraform.RunTerraformCommandAndGetStdoutE(t, Options, args...)

	require.NoError(t, err)

	//return the output
	return output
}

func NewTerraformState(t *testing.T, key string) *TerraFormState {
	os.Unsetenv("TF_DATA_DIR")
	tfState := new(TerraFormState)
	options := &terraform.Options{
		TerraformDir: os.Getenv("STATE_FILE_PATH"),
	}
	outputJson := ReadTerraformOutputJson(t, options, "objects")
	json.Unmarshal([]byte(outputJson), &tfState.Objects)
	tfState.Key = key
	tfState.SubscriptionID = os.Getenv("ARM_SUBSCRIPTION_ID")
	tfState.Environment = os.Getenv("ENVIRONMENT")
	return tfState
}

func (tfState TerraFormState) GetResources() Resource {
	resourceList := tfState.Objects[tfState.Key].(Resource)
	return resourceList
}

func (tfState TerraFormState) GetClientConfig() Resource {
	resourceList := tfState.GetResources()
	client_config := resourceList["client_config"].(Resource)
	return client_config
}

func (tfState TerraFormState) GetGlobalSettings() Resource {
	resourceList := tfState.GetResources()
	client_config := resourceList["global_settings"].(Resource)
	return client_config
}
func (tfState TerraFormState) GetLandingZoneKey() string {
	client_config := tfState.GetClientConfig()
	landing_zone_key := client_config["landingzone_key"]
	return landing_zone_key.(string)
}

func (tfState TerraFormState) GetResourceGroups() ResourceGroups {
	resourceList := tfState.Objects[tfState.Key].(Resource)
	resourceGroups := resourceList["resource_groups"].(Resource)
	var m ResourceGroups = make(ResourceGroups)
	for key, item := range resourceGroups {
		rg := item.(Resource)
		m[key] = *NewAzureResource(rg)
	}
	return m
}

func (tfState TerraFormState) GetKeyVaults() KeyVaults {
	resourceList := tfState.Objects[tfState.Key].(Resource)
	keyVaults := resourceList["keyvaults"].(Resource)

	var m KeyVaults
	m = make(KeyVaults)
	for key, item := range keyVaults {
		kv := item.(Resource)
		m[key] = *NewAzureResource(kv)
	}
	return m
}

func (tfState TerraFormState) GetRecoveryVaults() RecoveryVaults {
	resourceList := tfState.Objects[tfState.Key].(Resource)
	recoveryVaults := resourceList["recovery_vaults"].(Resource)

	var m RecoveryVaults
	m = make(RecoveryVaults)
	for key, item := range recoveryVaults {
		rv := item.(Resource)
		m[key] = *NewAzureResource(rv)
	}
	return m
}

func (tfState TerraFormState) GetKeyVaultByResourceGroup(resourceGroup string) (AzureResource, error) {
	keyVaults := tfState.GetKeyVaults()
	for _, keyVault := range keyVaults {
		id := keyVault.Resource["id"].(string)
		searchString := fmt.Sprintf("resourceGroups/%s", resourceGroup)
		if strings.Contains(id, searchString) {
			return keyVault, nil
		}
	}
	return *NewAzureResource(nil), errors.New("Keyvault not found")
}

func (tfState TerraFormState) GetStorageAccounts() KeyVaults {
	resourceList := tfState.Objects[tfState.Key].(Resource)
	storageAccounts := resourceList["storage_accounts"].(Resource)

	var m StorageAccounts
	m = make(StorageAccounts)
	for key, resourceGroup := range storageAccounts {
		sa := resourceGroup.(Resource)
		m[key] = *NewAzureResource(sa)
	}
	return m
}

func (tfState TerraFormState) GetStorageAccountByResourceGroup(resourceGroup string) (AzureResource, error) {
	storageAccounts := tfState.GetStorageAccounts()
	for _, storageAccount := range storageAccounts {
		id := storageAccount.Resource["id"].(string)
		searchString := fmt.Sprintf("resourceGroups/%s", resourceGroup)
		if strings.Contains(id, searchString) {
			return storageAccount, nil
		}
	}
	return *NewAzureResource(nil), errors.New("Storage Account not found")
}

func (tfState TerraFormState) GetVNets() VNets {
	resourceList := tfState.Objects[tfState.Key].(Resource)
	vNets := resourceList["vnets"].(Resource)
	var m VNets = make(VNets)
	for key, vNet := range vNets {
		vn := vNet.(Resource)
		m[key] = *NewAzureResource(vn)
	}
	return m
}

func (tfState TerraFormState) GetAKSClusters() AKSClusters {
	resourceList := tfState.Objects[tfState.Key].(Resource)
	clusters := resourceList["aks_clusters"].(Resource)
	azureResourceClusters := make(AKSClusters)
	for key, item := range clusters {
		rv := item.(Resource)
		azureResourceClusters[key] = *NewAzureResource(rv)
	}
	return azureResourceClusters
}

func (tfState TerraFormState) GetSQLServers() SQLServers {
	resourceList := tfState.Objects[tfState.Key].(Resource)
	sqlServers := resourceList["mssql_servers"].(Resource)
	var m = make(SQLServers)
	for key, item := range sqlServers {
		server := item.(Resource)
		m[key] = *NewAzureResource(server)
	}
	return m
}

func (tfState TerraFormState) GetSQLDBs() SQLDBs {
	resourceList := tfState.Objects[tfState.Key].(Resource)
	sqldbs := resourceList["mssql_databases"].(Resource)
	var m = make(SQLDBs)
	for key, item := range sqldbs {
		server := item.(Resource)
		m[key] = *NewAzureResource(server)
	}
	return m
}

func (tfState TerraFormState) GetAppServices() AppServices {
	resourceList := tfState.Objects[tfState.Key].(Resource)
	appServices := resourceList["app_services"].(Resource)
	var m = make(AppServices)
	for key, item := range appServices {
		service := item.(Resource)
		m[key] = *NewAzureResource(service)
	}
	return m
}

func (tfState TerraFormState) GetAppInsights() AppInsights {
	resourceList := tfState.Objects[tfState.Key].(Resource)
	appInsights := resourceList["application_insights"].(Resource)
	var m = make(AppInsights)
	for key, item := range appInsights {
		service := item.(Resource)
		m[key] = *NewAzureResource(service)
	}
	return m
}

func (tfState TerraFormState) GetStorageAccountQueues() StorageAccountQueues {
	resourceList := tfState.Objects[tfState.Key].(Resource)
	storageAccountQueues := resourceList["storage_account_queues"].(Resource)
	var m = make(StorageAccountQueues)
	for key, item := range storageAccountQueues {
		service := item.(Resource)
		m[key] = *NewAzureResource(service)
	}
	return m
}

func (tfState TerraFormState) GetRandomStrings() RandomStrings {
	resourceList := tfState.Objects[tfState.Key].(Resource)
	randomStrings := resourceList["random_strings"].(Resource)
	var m = make(RandomStrings)
	for key, item := range randomStrings {
		service := item.(Resource)
		m[key] = *NewAzureResource(service)
	}
	return m
}

func (tfState TerraFormState) GetMachineLearningWorkspaces() MachineLearningWorkspaces {
	resourceList := tfState.Objects[tfState.Key].(Resource)
	machineLearningWorkspaces := resourceList["machine_learning_workspaces"].(Resource)
	var m = make(MachineLearningWorkspaces)
	for key, item := range machineLearningWorkspaces {
		service := item.(Resource)
		m[key] = *NewAzureResource(service)
	}
	return m
}

func (tfState TerraFormState) GetAzureContainerRegistries() AzureContainerRegistries {
	resourceList := tfState.Objects[tfState.Key].(Resource)
	azureContainerRegistries := resourceList["azure_container_registries"].(Resource)
	var m = make(AzureContainerRegistries)
	for key, item := range azureContainerRegistries {
		service := item.(Resource)
		m[key] = *NewAzureResource(service)
	}
	return m
}

func NewAzureResource(resource Resource) *AzureResource {
	azureResource := new(AzureResource)
	azureResource.Resource = resource
	return azureResource
}

func (r AzureResource) GetString(key string) string {
	return r.Resource[key].(string)
}

func (r AzureResource) GetName() string {
	return r.GetString("name")
}

func (r AzureResource) GetResource(key string) Resource {
	return r.Resource[key].(Resource)
}

func (r AzureResource) GetTags() Resource {
	return r.GetResource("tags")
}

func (r AzureResource) GetLevel() string {
	tags := r.GetTags()
	return tags["level"].(string)
}

func (r AzureResource) GetHostName() string {
	return r.GetString("default_site_hostname")
}

func (r AzureResource) GetID() string {
	return r.GetString("id")
}