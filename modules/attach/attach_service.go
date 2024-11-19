package attach

import (
	"bnqkl/chain-cms/exception"
	"bnqkl/chain-cms/logger"
	"bnqkl/chain-cms/storage"
	"io"
	"os"
	"path/filepath"

	"gorm.io/gorm"
)

type AttachService struct {
	db  *gorm.DB
	log *logger.Logger
}

var attachService *AttachService

func NewAttachService(db *gorm.DB, log *logger.Logger) {
	attachService = &AttachService{
		db:  db,
		log: log,
	}
}

func GetAttachService() *AttachService {
	return attachService
}

func (s *AttachService) UploadBlob(req UploadBlobReq) (UploadBlobRes, error) {
	res := UploadBlobRes{}
	// 获取源文件File
	srcFile, err := req.File.Open() // 打开并返回FileHeader的相关文件
	if err != nil {
		s.log.Error(err)
		return res, exception.NewExceptionWithoutParam(exception.THE_UPLOADED_FILE_IS_INVALID)
	}
	defer srcFile.Close()
	fileName := req.GetName()
	if req.Extension != nil {
		fileName += "." + req.GetExtension()
	}
	savePath := filepath.Join(storage.GetBlobStorageDir(), fileName)
	dstFile, err := os.Create(savePath)
	if err != nil {
		s.log.Error(err)
		return res, exception.NewExceptionWithoutParam(exception.THE_UPLOADED_FILE_SAVE_FAILED)
	}
	defer dstFile.Close()
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		s.log.Error(err)
		return res, exception.NewExceptionWithoutParam(exception.THE_UPLOADED_FILE_SAVE_FAILED)
	}
	res.Url = "blob/" + fileName
	return res, nil
}
