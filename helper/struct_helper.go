package helper

type PageReq struct {
	Page     *int `url:"page" json:"page" validate:"required,min=1"`                 // 页码
	PageSize *int `url:"pageSize" json:"pageSize" validate:"required,min=1,max=100"` // 条数
}

func (req *PageReq) GetPage() int {
	return *req.Page
}

func (req *PageReq) GetPageSize() int {
	return *req.PageSize
}

type Pagination struct {
	Page     int   `json:"page"`     // 页码
	PageSize int   `json:"pageSize"` // 条数
	Total    int64 `json:"total"`    // 总条数
	HasMore  bool  `json:"hasMore"`  // 是否还有更多数据
}

func NewPagination(page int, pageSize int, total int64) Pagination {
	curTotal := page * pageSize
	return Pagination{
		Page:     page,
		PageSize: pageSize,
		Total:    total,
		HasMore:  total > int64(curTotal),
	}
}

type IndexReq struct {
	MinIndex int `url:"minInex" validate:"required,min=0"`  // 最小索引
	MaxIndex int `url:"maxIndex" validate:"required,min=0"` // 最大索引
}
