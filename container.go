package tinycontainer

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"strconv"
	"syscall"

	"github.com/Eslam-Nawara/tinycontainer/internal/namegenerator"
)

func Run(args []string) {
	fmt.Println("Running as", os.Getpid())
	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Credential: &syscall.Credential{
			Uid: uint32(syscall.Getuid()),
			Gid: uint32(syscall.Getgid()),
		},
		Cloneflags:   syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS | syscall.CLONE_NEWUSER,
		Unshareflags: syscall.CLONE_NEWNS,
		UidMappings:  []syscall.SysProcIDMap{{ContainerID: 0, HostID: syscall.Getuid(), Size: 1}},
		GidMappings:  []syscall.SysProcIDMap{{ContainerID: 0, HostID: syscall.Getgid(), Size: 1}},
	}

	check(cmd.Run())
}

func Child(command string, args []string) {
	fmt.Println("Running as", os.Getpid())

	newCgroup()

	check(syscall.Sethostname(namegenerator.nameGenerator()))
	check(syscall.Chroot("./rootfs"))
	check(syscall.Chdir("/"))
	check(syscall.Mount("proc", "proc", "proc", 0, ""))
	check(syscall.Mount("dev", "dev", "tmpfs", 0, ""))

	cmd := exec.Command(command, args...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	check(cmd.Run())
	check(syscall.Unmount("/proc", 0))
	check(syscall.Unmount("/dev", 0))
}

func newCgroup() {
	pids := "/sys/fs/cgroup/pids"
	os.MkdirAll(pids, 0755)

	check(os.WriteFile(path.Join(pids, "pids.max"), []byte("10"), 0700))
	check(os.WriteFile(path.Join(pids, "cgroup.procs"), []byte(strconv.Itoa(os.Getpid())), 0700))
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
