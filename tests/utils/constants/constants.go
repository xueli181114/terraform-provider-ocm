package constants

import (
	"os"
	"path"
	"strings"
)

func initDIR() string {
	if os.Getenv("MANIFESTS_DIR") != "" {
		return os.Getenv("MANIFESTS_DIR")
	}
	currentDir, _ := os.Getwd()
	return path.Join(strings.SplitAfter(currentDir, "tests")[0], "tf-manifests")
}

var configrationDir = initDIR()

// Provider dirs' name definition
const (
	AWSProviderDIR   = "aws"
	AZUREProviderDIR = "azure"
	RHCSProviderDIR  = "rhcs"
)

// Dirs of aws provider
var (
	AccountRolesDir  = path.Join(configrationDir, AWSProviderDIR, "account-roles")
	OperatorRolesDir = path.Join(configrationDir, AWSProviderDIR, "operator-roles")
	AWSVPCDir        = path.Join(configrationDir, AWSProviderDIR, "vpc")
)

// Dirs of rhcs provider
var (
	ClusterDir     = path.Join(configrationDir, RHCSProviderDIR, "clusters")
	IDPsDir        = path.Join(configrationDir, RHCSProviderDIR, "idps")
	MachinePoolDir = path.Join(configrationDir, RHCSProviderDIR, "machine-pools")
)

// Dirs of different types of clusters
var (
	ROSAClassic = path.Join(ClusterDir, "rosa-classic")
	OSDCCS      = path.Join(ClusterDir, "osd-ccs")
)

// Supports abs and relatives
func GrantClusterManifestDir(manifestDir string) string {
	var targetDir string
	if strings.Contains(manifestDir, ClusterDir) {
		targetDir = manifestDir
	} else {
		targetDir = path.Join(ClusterDir, manifestDir)
	}
	return targetDir
}

// Dirs of azure provider
// Just a placeholder

const (
	DefaultAWSRegion = "us-east-2"
	TokenENVName     = "RHCS_TOKEN"
)
