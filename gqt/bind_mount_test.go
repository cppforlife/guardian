package gqt_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	"code.cloudfoundry.org/garden"
	"code.cloudfoundry.org/guardian/gqt/runner"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type bindMountDescriptor struct {
	garden.BindMount
	fileToCheck string
	mountpoints []string
}

func destroy(descriptor bindMountDescriptor) {
	for i := len(descriptor.mountpoints) - 1; i >= 0; i-- {
		unmount(descriptor.mountpoints[i])
	}

	Expect(os.RemoveAll(descriptor.BindMount.SrcPath)).To(Succeed())
}

var _ = FDescribe("Bind mount", func() {
	var (
		client        *runner.RunningGarden
		container     garden.Container
		containerSpec garden.ContainerSpec

		bindMounts map[string]bindMountDescriptor
	)

	BeforeEach(func() {
		bindMounts = make(map[string]bindMountDescriptor)

		// bindMounts["ro-file"] = createFileSrcBindMount("/home/alice/file/ro", garden.BindMountModeRO)
		// bindMounts["rw-file"] = createFileSrcBindMount("/home/alice/file/rw", garden.BindMountModeRW)
		//
		// bindMounts["ro-dir"] = createDirSrcBindMount("/home/alice/dir/ro", garden.BindMountModeRO)
		// bindMounts["rw-dir"] = createDirSrcBindMount("/home/alice/dir/rw", garden.BindMountModeRW)

		bindMounts["ro-mountpoint-dir"] = createMountpointDirSrcBindMount("/home/alice/ro-mountpoint-dir", garden.BindMountModeRO)
		bindMounts["rw-mountpoint-dir"] = createMountpointDirSrcBindMount("/home/alice/rw-mountpoint-dir", garden.BindMountModeRW)

		bindMounts["ro-nested-mountpoint-dir"] = createNestedMountpointDirSrcBindMount("/home/alice/ro-nested-mountpoint-dir", garden.BindMountModeRO)
		bindMounts["rw-nested-mountpoint-dir"] = createNestedMountpointDirSrcBindMount("/home/alice/rw-nested-mountpoint-dir", garden.BindMountModeRW)

		containerSpec = garden.ContainerSpec{
			BindMounts: gardenBindMounts(bindMounts),
			Network:    fmt.Sprintf("10.0.%d.0/24", GinkgoParallelNode()),
		}
	})

	JustBeforeEach(func() {
		client = runner.Start(config)

		var err error
		container, err = client.Create(containerSpec)
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		for _, desc := range bindMounts {
			destroy(desc)
		}

		Expect(client.DestroyAndStop()).To(Succeed())
	})

	Context("when the continer is privileged", func() {
		BeforeEach(func() {
			containerSpec.Privileged = true
		})

		Context("User root", func() {
			It("can read", func() {
				canRead("root", container, all(bindMounts)...)
			})

			It("can write", func() {
				canWrite("root", container, readWrite(bindMounts)...)
			})

			It("cannot write", func() {
				canNotWrite("root", container, readOnly(bindMounts)...)
			})
		})

		Context("User alice", func() {
			It("can read", func() {
				canRead("alice", container, all(bindMounts)...)
			})

			It("can write", func() {
				canWrite("alice", container, readWrite(bindMounts)...)
			})

			It("cannot write", func() {
				canNotWrite("alice", container, readOnly(bindMounts)...)
			})
		})
	})

	Context("when the container is not privileged", func() {
		BeforeEach(func() {
			containerSpec.Privileged = false
		})

		Context("User root", func() {
			It("can read", func() {
				canRead("root", container, all(bindMounts)...)
			})

			It("can write", func() {
				canWrite("root", container, readWrite(bindMounts)...)
			})

			It("cannot write", func() {
				canNotWrite("root", container, readOnly(bindMounts)...)
			})
		})

		Context("User alice", func() {
			It("can read", func() {
				canRead("alice", container, all(bindMounts)...)
			})

			It("can write", func() {
				canWrite("alice", container, readWrite(bindMounts)...)
			})

			It("cannot write", func() {
				canNotWrite("alice", container, readOnly(bindMounts)...)
			})
		})

	})

})

func unmount(mountpoint string) {
	cmd := exec.Command("umount", "-f", mountpoint)
	output, err := cmd.CombinedOutput()
	if len(output) > 0 {
		fmt.Printf("Command: umount -f [%s]\n%v", mountpoint, string(output))
	}
	Expect(err).NotTo(HaveOccurred())
}

func createFileSrcBindMount(dstPath string, mode garden.BindMountMode) bindMountDescriptor {
	file, err := ioutil.TempFile("", fmt.Sprintf("file-%s-", humanise(mode)))
	Expect(err).NotTo(HaveOccurred())
	defer file.Close()

	return bindMountDescriptor{
		BindMount: garden.BindMount{
			SrcPath: file.Name(),
			DstPath: dstPath,
			Mode:    mode,
		},
		fileToCheck: file.Name(),
	}
}

func createDirSrcBindMount(dstPath string, mode garden.BindMountMode) bindMountDescriptor {
	dir, err := ioutil.TempDir("", fmt.Sprintf("dir-%s-", humanise(mode)))
	Expect(err).NotTo(HaveOccurred())
	Expect(os.Chown(dir, 0, 0)).To(Succeed())
	Expect(os.Chmod(dir, 0755)).To(Succeed())
	filePath := filepath.Join(dir, "testfile")
	Expect(ioutil.WriteFile(filePath, []byte{}, os.ModePerm)).To(Succeed())
	// Expect(os.Chmod(filePath, 0777)).To(Succeed())

	return bindMountDescriptor{
		BindMount: garden.BindMount{
			SrcPath: dir,
			DstPath: dstPath,
			Mode:    mode,
		},
		fileToCheck: filepath.Join(dstPath, "testfile"),
	}
}

func createMountpointDirSrcBindMount(dstPath string, mode garden.BindMountMode) bindMountDescriptor {
	desc := createDirSrcBindMount(dstPath, mode)

	var cmd *exec.Cmd
	cmd = exec.Command("mount", "--bind", desc.BindMount.SrcPath, desc.BindMount.SrcPath)
	Expect(cmd.Run()).To(Succeed())

	cmd = exec.Command("mount", "--make-shared", desc.BindMount.SrcPath)
	Expect(cmd.Run()).To(Succeed())

	desc.mountpoints = []string{desc.BindMount.SrcPath}

	return desc
}

func createNestedMountpointDirSrcBindMount(dstPath string, mode garden.BindMountMode) bindMountDescriptor {
	desc := createMountpointDirSrcBindMount(dstPath, mode)

	nestedBindPath := filepath.Join(desc.SrcPath, "nested-bind")
	Expect(os.MkdirAll(nestedBindPath, os.FileMode(0755))).To(Succeed())

	cmd := exec.Command("mount", "-t", "tmpfs", "tmpfs", nestedBindPath)
	Expect(cmd.Run()).To(Succeed())

	filePath := filepath.Join(nestedBindPath, "nested-file")
	Expect(ioutil.WriteFile(filePath, []byte{}, os.ModePerm)).To(Succeed())
	Expect(os.Chmod(filePath, 0777)).To(Succeed())

	desc.mountpoints = append(desc.mountpoints, nestedBindPath)
	desc.BindMount.SrcPath = nestedBindPath
	desc.fileToCheck = filepath.Join(desc.BindMount.DstPath, "nested-file")

	return desc
}

func gardenBindMounts(bindMounts map[string]bindMountDescriptor) []garden.BindMount {
	mounts := []garden.BindMount{}
	for _, v := range bindMounts {
		mounts = append(mounts, v.BindMount)
	}

	return mounts
}

// TODO: maybe we need canTouch/canNotTouch as well to make sure we can create new files in the bind mount

func canRead(user string, container garden.Container, descriptors ...bindMountDescriptor) {
	for _, d := range descriptors {
		Expect(containerReadFile(container, d.fileToCheck, user)).To(Succeed())
	}
}

func canNotRead(user string, container garden.Container, descriptors ...bindMountDescriptor) {
	for _, d := range descriptors {
		Expect(containerReadFile(container, d.fileToCheck, user)).NotTo(Succeed())
	}
}

func canWrite(user string, container garden.Container, descriptors ...bindMountDescriptor) {
	for _, d := range descriptors {
		if isDir(d.BindMount.SrcPath) {
			Expect(touchFile(container, filepath.Dir(d.fileToCheck), user)).To(Succeed())
		} else {
			Expect(writeFile(container, d.fileToCheck, user)).To(Succeed())
		}
	}
}

func canNotWrite(user string, container garden.Container, descriptors ...bindMountDescriptor) {
	for _, d := range descriptors {
		if isDir(d.BindMount.SrcPath) {
			Expect(touchFile(container, filepath.Dir(d.fileToCheck), user)).NotTo(Succeed())
		} else {
			Expect(writeFile(container, d.fileToCheck, user)).NotTo(Succeed())
		}
	}
}

func all(bindMounts map[string]bindMountDescriptor) []bindMountDescriptor {
	descs := []bindMountDescriptor{}
	for _, v := range bindMounts {
		descs = append(descs, v)
	}

	return descs
}

func readWrite(bindMounts map[string]bindMountDescriptor) []bindMountDescriptor {
	rw := []bindMountDescriptor{}
	for _, v := range bindMounts {
		if v.BindMount.Mode == garden.BindMountModeRW {
			rw = append(rw, v)
		}
	}

	return rw
}

func readOnly(bindMounts map[string]bindMountDescriptor) []bindMountDescriptor {
	ro := []bindMountDescriptor{}
	for _, v := range bindMounts {
		if v.BindMount.Mode == garden.BindMountModeRO {
			ro = append(ro, v)
		}
	}

	return ro
}

func containerReadFile(container garden.Container, filePath, user string) error {
	process, err := container.Run(garden.ProcessSpec{
		Path: "cat",
		Args: []string{filePath},
		User: user,
	}, ginkgoIO)
	Expect(err).ToNot(HaveOccurred())

	exitCode, err := process.Wait()
	Expect(err).ToNot(HaveOccurred())
	if exitCode != 0 {
		return fmt.Errorf("Could not read file %s in container %s, exit code %d", filePath, container.Handle(), exitCode)
	}

	return nil
}

func writeFile(container garden.Container, dstPath, user string) error {
	process, err := container.Run(garden.ProcessSpec{
		Path: "/bin/sh",
		Args: []string{"-c", fmt.Sprintf("echo i-can-write > %s", dstPath)},
		User: user,
	}, ginkgoIO)
	Expect(err).ToNot(HaveOccurred())

	exitCode, err := process.Wait()
	Expect(err).ToNot(HaveOccurred())
	if exitCode != 0 {
		return fmt.Errorf("Could not write to file %s in container %s, exit code %d", dstPath, container.Handle(), exitCode)
	}

	return nil
}

func touchFile(container garden.Container, dstDir, user string) error {
	fileToTouch := filepath.Join(dstDir, "can-touch-this")
	process, err := container.Run(garden.ProcessSpec{
		Path: "touch",
		Args: []string{fileToTouch},
		User: user,
	}, ginkgoIO)
	Expect(err).ToNot(HaveOccurred())

	exitCode, err := process.Wait()
	Expect(err).ToNot(HaveOccurred())
	if exitCode != 0 {
		return fmt.Errorf("Could not touch file %s in container %s, exit code %d", fileToTouch, container.Handle(), exitCode)
	}

	return nil
}

func humanise(mode garden.BindMountMode) string {
	if mode == garden.BindMountModeRW {
		return "rw"
	}
	return "ro"
}

func isDir(path string) bool {
	stat, err := os.Stat(path)
	Expect(err).NotTo(HaveOccurred())
	return stat.IsDir()
}
