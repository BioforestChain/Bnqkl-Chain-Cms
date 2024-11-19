package attach

import "mime/multipart"

// 上传 blob
//
// method: post
//
// path: /api/attach/upload/blob
type UploadBlobReq struct {
	Name      *string               `form:"name" json:"name" binding:"required,min=1,max=255"`                               // blob名称
	Extension *string               `form:"extension,omitempty" json:"extension,omitempty" binding:"omitempty,min=1,max=30"` // blob扩展名
	File      *multipart.FileHeader `form:"file" type:"blob" json:"file" binding:"required"`                                 // blob
}

func (req *UploadBlobReq) GetName() string {
	if req.Name == nil {
		return ""
	}
	return *req.Name
}

func (req *UploadBlobReq) GetExtension() string {
	if req.Extension == nil {
		return ""
	}
	return *req.Extension
}

type UploadBlobRes struct {
	Url string `json:"url"` // 文件路径
}
