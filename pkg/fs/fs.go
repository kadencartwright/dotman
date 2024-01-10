package fs

import (
	"errors"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
)

func makeParentDirRecursive(path string) error {
	parentDir := filepath.Dir(path)
	cmd := exec.Command("mkdir", "-p", parentDir)
	err := cmd.Run()
	return err
}
func LinkPaths(source, destination string) error {
	if pathsAreEquivalent(source, destination) {
		return nil
	}
	err := makeParentDirRecursive(destination)
	err = os.Symlink(source, destination)
	if os.IsExist(err) {
		os.RemoveAll(destination)
		return LinkPaths(source, destination)
	}
	return err
}
func Backup(path string) error {
	cd, err := os.UserCacheDir()
	if err != nil {
		return err
	}
	backupDir := filepath.Join(cd, "dotman", "backup", path)
	err = os.MkdirAll(backupDir, os.ModePerm)
	if err != nil {
		println(err.Error())
		return err
	}
	return nil
}

func CopyRecursively(source, destination string) error {
	//if both paths point to same location, return nil
	if pathsAreEquivalent(source, destination) {
		return nil
	}
	err := os.RemoveAll(filepath.Join(destination))
	if err != nil && !errors.Is(err, fs.ErrNotExist) {
		return err
	}
	cmd := exec.Command("cp", "-r", source, destination)
	err = cmd.Run()
	return err
}
func pathsAreEquivalent(path1, path2 string) bool {
	n1, err := filepath.EvalSymlinks(path1)
	if err != nil {
		return false
	}

	n2, err := filepath.EvalSymlinks(path2)
	if err != nil {
		return false
	}

	return n1 == n2
}
