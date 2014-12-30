// build +linux
package namespace

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"syscall"

	"github.com/docker/docker-network/Godeps/_workspace/src/golang.org/x/sys/unix"
)

// New creates new namespace at specified path
func New(path string) (*Namespace, error) {
	runtime.LockOSThread()
	if err := unix.Unshare(unix.CLONE_NEWNET); err != nil {
		return nil, err
	}
	f, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	f.Close()
	if err := unix.Mount(fmt.Sprintf("/proc/self/ns/net"), path, "bind", unix.MS_BIND, ""); err != nil {
		return nil, err
	}
	return &Namespace{Path: path}, nil
}

// Join joins to network namespace
func (n *Namespace) Join() error {
	runtime.LockOSThread()
	f, err := os.OpenFile(n.Path, os.O_RDONLY, 0)
	if err != nil {
		return fmt.Errorf("failed get network namespace fd: %v", err)
	}
	defer f.Close()
	if _, _, err := unix.RawSyscall(unix.SYS_SETNS, f.Fd(), unix.CLONE_NEWNET, 0); err != 0 {
		return err
	}
	return nil
}

// Delete physically removes namespace if it is possible
func (n *Namespace) Delete() error {
	runtime.LockOSThread()
	if err := syscall.Unmount(n.Path, syscall.MNT_DETACH); err != nil {
		return err
	}
	return os.Remove(n.Path)
}

func (n *Namespace) Exec(cmd *exec.Cmd) error {
	if err := n.Join(); err != nil {
		return err
	}
	return cmd.Run()
}
