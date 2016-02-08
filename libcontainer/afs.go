package libcontainer

import (
	"os/exec"

	"github.com/pagarme/goafs"

	"github.com/opencontainers/runc/libcontainer/configs"
)

func setupAfs(config *configs.Afs) error {
	if !config.Enabled {
		return nil
	}

	if err := goafs.Setpag(); err != nil {
		return err
	}

	// Make sure we're not holding any tickets
	if err := goafs.Unlog(); err != nil {
		return err
	}

	cmd := exec.Command(config.AklogPath, "-d", "-force", "-noprdb", "-keytab", config.KerberosKeytab, "-principal", config.KerberosPrincipal)

	cmd.Env = []string{}

	return cmd.Run()
}
