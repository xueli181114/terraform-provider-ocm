package e2e

import (

	// nolint

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	EXE "github.com/terraform-redhat/terraform-provider-rhcs/tests/utils/exec"
)

var region = "us-west-2"

var _ = Describe("TF Test", func() {
	Describe("Create cluster test", func() {
		It("TestExampleNegative", func() {

			clusterParam := &EXE.ClusterCreationArgs{
				Token:              token,
				OCMENV:             "staging",
				ClusterName:        "xuelitf",
				OperatorRolePrefix: "xueli",
				AccountRolePrefix:  "xueli",
				Replicas:           3,
				OpenshiftVersion:   "invalid",
				OIDCConfig:         "managed",
			}

			_, err := EXE.CreateMyTFCluster(clusterParam, "-auto-approve")
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).Should(ContainSubstring("version %s is not in the list", clusterParam.OpenshiftVersion))
		})
		It("TestExampleCritical", func() {
			accRolePrefix := "xueli-2"
			By("Create VPCs")
			args := &EXE.VPCVariables{
				ClusterName: "xueli",
				AWSRegion:   region,
				MultiAZ:     true,
				VPCCIDR:     "11.0.0.0/16",
			}
			priSubnets, pubSubnets, zones, err := EXE.CreateAWSVPC(args, "-auto-approve")
			Expect(err).ToNot(HaveOccurred())
			// defer DestroyAWSVPC(args, "-auto-approve")

			By("Create account-roles")
			accRoleParam := &EXE.AccountRolesArgs{
				Token:             token,
				AccountRolePrefix: accRolePrefix,
			}
			_, err = EXE.CreateMyTFAccountRoles(accRoleParam, "-auto-approve")
			Expect(err).ToNot(HaveOccurred())
			// defer DestroyMyTFAccountRoles(accRoleParam, "-auto-approve")

			By("Create Cluster")
			clusterParam := &EXE.ClusterCreationArgs{
				Token:                token,
				OCMENV:               "staging",
				ClusterName:          "xuelitf",
				OperatorRolePrefix:   "xuelitf",
				AccountRolePrefix:    accRolePrefix,
				Replicas:             3,
				AWSRegion:            region,
				AWSAvailabilityZones: zones,
				AWSSubnetIDs:         append(priSubnets, pubSubnets...),
				MultiAZ:              true,
				MachineCIDR:          args.VPCCIDR,
				OIDCConfig:           "managed",
			}

			clusterID, err := EXE.CreateMyTFCluster(clusterParam, "-auto-approve")
			defer EXE.DestroyMyTFCluster(clusterParam, "-auto-approve")
			Expect(clusterID).ToNot(BeEmpty())
			Expect(err).ToNot(HaveOccurred())

		})
	})
})
