package helper

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

type PathHelper struct {
	rootPath string
}

var pathHelper *PathHelper

func GetRootPath() string {
	return pathHelper.rootPath
}

func NewPathHelper(rootPath string) *PathHelper {
	return &PathHelper{
		rootPath: rootPath,
	}
}

func InitRootPath() error {
	pathHelper = NewPathHelper("")
	err := pathHelper.InitRootPath()
	return err
}

func (helper *PathHelper) InitRootPath() error {
	rootPath, err := helper.getRootPathByExecutable()
	if err != nil {
		return err
	}
	tempRootPath, err := helper.getRootPathByTmpDir()
	if err != nil {
		return err
	}
	if !strings.Contains(rootPath, tempRootPath) {
		helper.rootPath = rootPath
		return nil
	}
	rootPath, err = helper.getRootPathByCaller()
	if err != nil {
		return err
	}
	helper.rootPath = rootPath
	return nil
}

// 获取系统临时目录，兼容go run
func (helper *PathHelper) getRootPathByTmpDir() (string, error) {
	tempDir := os.Getenv("TEMP")
	if tempDir == "" {
		tempDir = os.Getenv("TMP")
	}
	rootPath, error := filepath.EvalSymlinks(tempDir)
	return rootPath, error
}

func (helper *PathHelper) getRootPathByExecutable() (string, error) {
	exePath, err := os.Executable()
	if err != nil {
		return "", err
	}
	rootPath, err := filepath.EvalSymlinks(filepath.Dir(exePath))
	return rootPath, err
}

func (helper *PathHelper) getRootPathByCaller() (string, error) {
	_, runFileName, _, ok := runtime.Caller(0)
	if !ok {
		return "", errors.New("failed to get run file name")
	}
	rootPath := filepath.Dir(filepath.Dir(runFileName))
	return rootPath, nil
}
