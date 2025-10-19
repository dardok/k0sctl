package linux

import (
	"strings"

	"github.com/k0sproject/k0sctl/configurer"
	"github.com/k0sproject/rig"
	"github.com/k0sproject/rig/exec"
	"github.com/k0sproject/rig/os"
	"github.com/k0sproject/rig/os/registry"
)

// BaseLinux for tricking go interfaces
type BaseLinux struct {
	configurer.Linux
}

// Gentoo provides OS support for Gentoo Linux
type Gentoo struct {
	os.Linux
	BaseLinux
}

func init() {
	registry.RegisterOSModule(
		func(os rig.OSVersion) bool {
			return os.ID == "gentoo"
		},
		func() interface{} {
			return &Gentoo{}
		},
	)
}

// InstallPackage installs packages via slackpkg
func (l *Gentoo) InstallPackage(h os.Host, pkg ...string) error {
	return h.Execf("emerge %s", strings.Join(pkg, " "), exec.Sudo(h))
}

func (l *Gentoo) Prepare(h os.Host) error {
	return l.InstallPackage(h, "sys-apps/findutils", "sys-apps/coreutils")
}
