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

	env := []string{"KRB5CCNAME=KEYRING:session:container-afs"}

	cmd := exec.Command("kinit", "-k", "-t", config.KerberosKeytab, config.KerberosPrincipal)
	cmd.Env = env

	if err := cmd.Run(); err != nil {
		return err
	}

	cmd = exec.Command("aklog", "-force")
	cmd.Env = env

	return cmd.Run()
}
