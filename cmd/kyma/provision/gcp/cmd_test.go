package gcp

import (
	"testing"

	"github.com/kyma-incubator/hydroform/types"
	"github.com/kyma-project/cli/internal/cli"
	"github.com/stretchr/testify/require"
)

// TestUninstallFlags ensures that the provided command flags are stored in the options.
func TestProvisionGCPFlags(t *testing.T) {
	o := NewOptions(&cli.Options{})
	c := NewCmd(o)

	// test default flag values
	require.Equal(t, "", o.Name, "Default value for the name flag not as expected.")
	require.Equal(t, "", o.Project, "Default value for the project flag not as expected.")
	require.Equal(t, "", o.CredentialsFile, "Default value for the credentials flag not as expected.")
	require.Equal(t, "1.14.6", o.KubernetesVersion, "Default value for the kube-version flag not as expected.")
	require.Equal(t, "europe-west3-a", o.Location, "Default value for the location flag not as expected.")
	require.Equal(t, "n1-standard-4", o.MachineType, "Default value for the type flag not as expected.")
	require.Equal(t, 30, o.DiskSizeGB, "Default value for the disk-size flag not as expected.")
	require.Equal(t, 1, o.NodeCount, "Default value for the nodes flag not as expected.")
	require.Empty(t, o.Extra, "Default value for the extra flag not as expected.")

	// test passing flags
	c.ParseFlags([]string{
		"-n", "my-cluster",
		"-p", "my-project",
		"-c", "/my/credentials/file",
		"-k", "1.16.0",
		"-l", "us-central1-c",
		"-t", "quantum-computer",
		"--disk-size", "2000",
		"--nodes", "7",
		"--extra", "VAR1=VALUE1,VAR2=VALUE2",
	})
	require.Equal(t, "my-cluster", o.Name, "The parsed value for the name flag not as expected.")
	require.Equal(t, "my-project", o.Project, "The parsed value for the project flag not as expected.")
	require.Equal(t, "/my/credentials/file", o.CredentialsFile, "The parsed value for the credentials flag not as expected.")
	require.Equal(t, "1.16.0", o.KubernetesVersion, "The parsed value for the kube-version flag not as expected.")
	require.Equal(t, "us-central1-c", o.Location, "The parsed value for the location flag not as expected.")
	require.Equal(t, "quantum-computer", o.MachineType, "The parsed value for the type flag not as expected.")
	require.Equal(t, 2000, o.DiskSizeGB, "The parsed value for the disk-size flag not as expected.")
	require.Equal(t, 7, o.NodeCount, "The parsed value for the nodes flag not as expected.")
	require.Equal(t, []string{"VAR1=VALUE1", "VAR2=VALUE2"}, o.Extra, "The parsed value for the extra flag not as expected.")
}

func TestProvisionGCPSubcommands(t *testing.T) {
	o := NewOptions(&cli.Options{})
	c := NewCmd(o)

	sub := c.Commands()

	require.Equal(t, 0, len(sub), "Number of provision gpc subcommands not as expected")
}

func TestNewCluster(t *testing.T) {
	o := &Options{
		Name:              "mega-cluster",
		KubernetesVersion: "1.16.0",
		Location:          "north-pole",
		MachineType:       "HAL",
		DiskSizeGB:        9000,
		NodeCount:         3,
	}

	c := newCluster(o)
	require.Equal(t, o.Name, c.Name, "Cluster name not as expected.")
	require.Equal(t, o.KubernetesVersion, c.KubernetesVersion, "Cluster Kubernetes version not as expected.")
	require.Equal(t, o.Location, c.Location, "Cluster location not as expected.")
	require.Equal(t, o.MachineType, c.MachineType, "Cluster machine type not as expected.")
	require.Equal(t, o.DiskSizeGB, c.DiskSizeGB, "Cluster disk size not as expected.")
	require.Equal(t, o.NodeCount, c.NodeCount, "Cluster number of nodes not as expected.")
}

func TestNewProvider(t *testing.T) {
	o := &Options{
		Project:         "cool-project",
		CredentialsFile: "/path/to/credentials",
		Extra:           []string{"VAR1=VALUE1", "VAR2=VALUE2"},
	}

	p, err := newProvider(o)
	require.NoError(t, err)

	require.Equal(t, types.GCP, p.Type, "Provider type not as expected.")
	require.Equal(t, o.Project, p.ProjectName, "Provider project name not as expected.")
	require.Equal(t, o.CredentialsFile, p.CredentialsFilePath, "Provider credentials file path not as expected.")
	require.Equal(t, map[string]interface{}{"VAR1": "VALUE1", "VAR2": "VALUE2"}, p.CustomConfigurations, "Provider extra configurations not as expected.")
}
