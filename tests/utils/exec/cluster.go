package exec

import (
	"context"
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"

	CON "github.com/terraform-redhat/terraform-provider-rhcs/tests/utils/constants"
)

type ClusterCreationArgs struct {
	AccountRolePrefix    string            `json:"account_role_prefix,omitempty"`
	OCMENV               string            `json:"rhcs_environment,omitempty"`
	ClusterName          string            `json:"cluster_name,omitempty"`
	OperatorRolePrefix   string            `json:"operator_role_prefix,omitempty"`
	OpenshiftVersion     string            `json:"openshift_version,omitempty"`
	Token                string            `json:"token,omitempty"`
	URL                  string            `json:"url,omitempty"`
	AWSRegion            string            `json:"aws_region,omitempty"`
	AWSAvailabilityZones []string          `json:"aws_availability_zones,omitempty"`
	Replicas             int               `json:"replicas,omitempty"`
	ChannelGroup         string            `json:"channel_group,omitempty"`
	AWSHttpTokensState   string            `json:"aws_http_tokens_state,omitempty"`
	PrivateLink          string            `json:"private_link,omitempty"`
	AWSSubnetIDs         []string          `json:"aws_subnet_ids,omitempty"`
	ComputeMachineType   string            `json:"compute_machine_type,omitempty"`
	DefaultMPLabels      map[string]string `json:"default_mp_labels,omitempty"`
	DisableSCPChecks     bool              `json:"disable_scp_checks,omitempty"`
	MultiAZ              bool              `json:"multi_az,omitempty"`
	MachineCIDR          string            `json:"machine_cidr,omitempty"`
	OIDCConfig           string            `json:"oidc_config,omitempty"`
}

// *********************** Cluster CMS ***********************************
func CreateCluster(ctx context.Context, args ...string) (string, error) {
	runTerraformInit(ctx, CON.ClusterDir)

	runTerraformApplyWithArgs(ctx, CON.ClusterDir, args)

	getClusterIdCmd := exec.Command("terraform", "output", "-json", "cluster_id")
	getClusterIdCmd.Dir = CON.ClusterDir
	output, err := getClusterIdCmd.Output()
	if err != nil {
		return "", err
	}

	splitOutput := strings.Split(string(output), "\"")
	if len(splitOutput) <= 1 {
		return "", fmt.Errorf("got no cluster id from the output")
	}

	return splitOutput[1], nil
}

func CreateTFCluster(ctx context.Context,
	varArgs map[string]interface{}, abArgs ...string) (string, error) {
	err := runTerraformInit(ctx, CON.ROSAClassic)
	if err != nil {
		return "", err
	}

	args := combineArgs(varArgs, abArgs...)
	_, err = runTerraformApplyWithArgs(ctx, CON.ROSAClassic, args)
	if err != nil {
		return "", err
	}

	getClusterIdCmd := exec.Command("terraform", "output", "-json", "cluster_id")
	getClusterIdCmd.Dir = CON.ClusterDir
	output, err := getClusterIdCmd.Output()
	if err != nil {
		return "", err
	}

	splitOutput := strings.Split(string(output), "\"")
	if len(splitOutput) <= 1 {
		return "", fmt.Errorf("got no cluster id from the output")
	}

	return splitOutput[1], nil
}

func DestroyTFCluster(ctx context.Context,
	varArgs map[string]interface{}, abArgs ...string) error {
	err := runTerraformInit(ctx, CON.ClusterDir)
	if err != nil {
		return err
	}

	args := combineArgs(varArgs, abArgs...)
	err = runTerraformDestroyWithArgs(ctx, CON.ClusterDir, args)
	if err != nil {
		return err
	}

	getClusterIdCmd := exec.Command("terraform", "output", "-json", "cluster_id")
	getClusterIdCmd.Dir = CON.ClusterDir
	_, err = getClusterIdCmd.Output()

	return err
}

func CreateMyTFCluster(clusterArgs *ClusterCreationArgs, arg ...string) (string, error) {
	parambytes, _ := json.Marshal(clusterArgs)
	args := map[string]interface{}{}
	json.Unmarshal(parambytes, &args)
	return CreateTFCluster(context.TODO(), args, arg...)
}

func DestroyMyTFCluster(clusterArgs *ClusterCreationArgs, arg ...string) error {
	parambytes, _ := json.Marshal(clusterArgs)
	args := map[string]interface{}{}
	json.Unmarshal(parambytes, &args)
	return DestroyTFCluster(context.TODO(), args, arg...)
}
