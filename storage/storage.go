package storage

import (
	"bnqkl/chain-cms/config"
	"bnqkl/chain-cms/helper"
	"bnqkl/chain-cms/logger"
	"os"
	"path/filepath"
)

var (
	blobStorageDir = ""
)

func InitStorage(log *logger.Logger) error {
	err := initAttachStorage(log)
	if err != nil {
		return err
	}
	return nil
}

func initDir(dirPath string) error {
	return os.MkdirAll(dirPath, 0777)
}

func initAttachStorage(log *logger.Logger) error {
	rootPath := helper.GetRootPath()
	config := config.GetConfig()
	blobStorageDir = filepath.Join(rootPath, "attach", config.Attach.Blob)
	err := initDir(blobStorageDir)
	if err != nil {
		return err
	}
	log.Info("init attach storage success")
	return nil
}

func GetBlobStorageDir() string {
	return blobStorageDir
}
