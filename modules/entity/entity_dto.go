package entity

import "bnqkl/chain-cms/database/model"

// 新增 entity
//
// method: post
//
// path: /api/entity/add
type AddEntityReq struct {
	Possessor        *string `form:"possessor" json:"possessor" binding:"required,min=1,max=255"`                         // 拥有者
	ChainName        *string `form:"chainName" json:"chainName" binding:"required,min=1,max=30"`                          // 链名
	ChainMagic       *string `form:"chainMagic" json:"chainMagic" binding:"required,min=1,max=30"`                        // 链网络标识
	FactoryName      *string `form:"factoryName,omitempty" json:"factoryName,omitempty" binding:"omitempty,min=1,max=50"` // 模板名称
	FactoryId        *string `form:"factoryId" json:"factoryId" binding:"required,min=1,max=50"`                          // 模板ID
	EntityId         *string `form:"entityId" json:"entityId" binding:"required,min=1,max=50"`                            // 非同质资产ID
	TaxCollector     *string `form:"taxCollector" json:"taxCollector" binding:"required,min=1,max=255"`                   // 收税人
	TaxAssetPrealnum *string `form:"taxAssetPrealnum" json:"taxAssetPrealnum" binding:"required,min=1,max=50"`            // 缴纳数量
	Type             *int    `form:"type" json:"type" binding:"required,min=1,max=2"`                                     // 类型(1:普通,2:限量,默认:1)
	Hash             *string `form:"hash" json:"hash" binding:"min=1,max=100"`                                            // blob哈希
	Extension        *string `form:"extension,omitempty" json:"extension,omitempty" binding:"omitempty,min=1,max=50"`     // blob扩展名
}

func (req *AddEntityReq) GetPossessor() string {
	if req.Possessor == nil {
		return ""
	}
	return *req.Possessor
}
func (req *AddEntityReq) GetChainName() string {
	if req.ChainName == nil {
		return ""
	}
	return *req.ChainName
}
func (req *AddEntityReq) GetChainMagic() string {
	if req.ChainMagic == nil {
		return ""
	}
	return *req.ChainMagic
}
func (req *AddEntityReq) GetFactoryName() string {
	if req.FactoryName == nil {
		return ""
	}
	return *req.FactoryName
}
func (req *AddEntityReq) GetFactoryId() string {
	if req.FactoryId == nil {
		return ""
	}
	return *req.FactoryId
}
func (req *AddEntityReq) GetEntityId() string {
	if req.EntityId == nil {
		return ""
	}
	return *req.EntityId
}
func (req *AddEntityReq) GetTaxCollector() string {
	if req.TaxCollector == nil {
		return ""
	}
	return *req.TaxCollector
}
func (req *AddEntityReq) GetTaxAssetPrealnum() string {
	if req.TaxAssetPrealnum == nil {
		return ""
	}
	return *req.TaxAssetPrealnum
}
func (req *AddEntityReq) GetType() int {
	if req.Type == nil {
		return 0
	}
	return *req.Type
}
func (req *AddEntityReq) GetHash() string {
	if req.Hash == nil {
		return ""
	}
	return *req.Hash
}
func (req *AddEntityReq) GetExtension() string {
	if req.Extension == nil {
		return ""
	}
	return *req.Extension
}

type AddEntityRes bool

type EntityStruct struct {
	EntityId         *string `form:"entityId" json:"entityId" binding:"required,min=1,max=50"`                               // 非同质资产ID
	TaxCollector     *string `form:"taxCollector,omitempty" json:"taxCollector,omitempty" binding:"omitempty,min=1,max=255"` // 收税人
	TaxAssetPrealnum *string `form:"taxAssetPrealnum" json:"taxAssetPrealnum" binding:"required,min=1,max=50"`               // 缴纳数量
}

func (req *EntityStruct) GetEntityId() string {
	if req.EntityId == nil {
		return ""
	}
	return *req.EntityId
}
func (req *EntityStruct) GetTaxCollector() string {
	if req.TaxCollector == nil {
		return ""
	}
	return *req.TaxCollector
}
func (req *EntityStruct) GetTaxAssetPrealnum() string {
	if req.TaxAssetPrealnum == nil {
		return ""
	}
	return *req.TaxAssetPrealnum
}

// 批量新增 entity
//
// method: post
//
// path: /api/entity/add/multi
type AddEntityMultiReq struct {
	Possessor    *string         `form:"possessor" json:"possessor" binding:"required,min=1,max=255"`                         // 拥有者
	ChainName    *string         `form:"chainName" json:"chainName" binding:"required,min=1,max=30"`                          // 链名
	ChainMagic   *string         `form:"chainMagic" json:"chainMagic" binding:"required,min=1,max=30"`                        // 链网络标识
	FactoryName  *string         `form:"factoryName,omitempty" json:"factoryName,omitempty" binding:"omitempty,min=1,max=50"` // 模板名称
	FactoryId    *string         `form:"factoryId" json:"factoryId" binding:"required,min=1,max=50"`                          // 模板ID
	TaxCollector *string         `form:"taxCollector" json:"taxCollector" binding:"required,min=1,max=255"`                   // 收税人
	Entities     *[]EntityStruct `form:"entities" json:"entities" binding:"required,gt=0,lt=6000,dive"`                       // 非同质资产列表
	Type         *int            `form:"type" json:"type" binding:"required,min=1,max=2"`                                     // 类型(1:普通,2:限量,默认:1)
	Hash         *string         `form:"hash" json:"hash" binding:"min=1,max=100"`                                            // blob哈希
	Extension    *string         `form:"extension,omitempty" json:"extension,omitempty" binding:"omitempty,min=1,max=50"`     // blob扩展名
}

func (req *AddEntityMultiReq) GetPossessor() string {
	if req.Possessor == nil {
		return ""
	}
	return *req.Possessor
}
func (req *AddEntityMultiReq) GetChainName() string {
	if req.ChainName == nil {
		return ""
	}
	return *req.ChainName
}
func (req *AddEntityMultiReq) GetChainMagic() string {
	if req.ChainMagic == nil {
		return ""
	}
	return *req.ChainMagic
}
func (req *AddEntityMultiReq) GetFactoryName() string {
	if req.FactoryName == nil {
		return ""
	}
	return *req.FactoryName
}
func (req *AddEntityMultiReq) GetFactoryId() string {
	if req.FactoryId == nil {
		return ""
	}
	return *req.FactoryId
}
func (req *AddEntityMultiReq) GetTaxCollector() string {
	if req.TaxCollector == nil {
		return ""
	}
	return *req.TaxCollector
}
func (req *AddEntityMultiReq) GetEntities() []EntityStruct {
	if req.Entities == nil {
		return []EntityStruct{}
	}
	return *req.Entities
}
func (req *AddEntityMultiReq) GetType() int {
	if req.Type == nil {
		return 0
	}
	return *req.Type
}
func (req *AddEntityMultiReq) GetHash() string {
	if req.Hash == nil {
		return ""
	}
	return *req.Hash
}
func (req *AddEntityMultiReq) GetExtension() string {
	if req.Extension == nil {
		return ""
	}
	return *req.Extension
}

type AddEntityMultiRes bool

// 更新 entity
//
// method: post
//
// path: /api/entity/update
type UpdateEntityReq struct {
	ChainName   *string `form:"chainName" json:"chainName" binding:"required,min=1,max=30"`                          // 链名
	ChainMagic  *string `form:"chainMagic" json:"chainMagic" binding:"required,min=1,max=30"`                        // 链网络标识
	FactoryId   *string `form:"factoryId" json:"factoryId" binding:"required,min=1,max=50"`                          // 模板ID
	EntityId    *string `form:"entityId" json:"entityId" binding:"required,min=1,max=50"`                            // 非同质资产ID
	Possessor   *string `form:"possessor,omitempty" json:"possessor,omitempty" binding:"omitempty,min=1,max=255"`    // 拥有者
	FactoryName *string `form:"factoryName,omitempty" json:"factoryName,omitempty" binding:"omitempty,min=1,max=50"` // 模板名称
}

func (req *UpdateEntityReq) GetChainName() string {
	if req.ChainName == nil {
		return ""
	}
	return *req.ChainName
}
func (req *UpdateEntityReq) GetChainMagic() string {
	if req.ChainMagic == nil {
		return ""
	}
	return *req.ChainMagic
}
func (req *UpdateEntityReq) GetFactoryId() string {
	if req.FactoryId == nil {
		return ""
	}
	return *req.FactoryId
}
func (req *UpdateEntityReq) GetEntityId() string {
	if req.EntityId == nil {
		return ""
	}
	return *req.EntityId
}
func (req *UpdateEntityReq) GetPossessor() string {
	if req.Possessor == nil {
		return ""
	}
	return *req.Possessor
}
func (req *UpdateEntityReq) GetFactoryName() string {
	if req.FactoryName == nil {
		return ""
	}
	return *req.FactoryName
}

type UpdateEntityRes bool

// 获取用户持有的 factory 总览
//
// method: get
//
// path: /api/entity/factory/all
type GetUserFactoryAllReq struct {
	Possessor *string `form:"possessor" json:"possessor" binding:"required,min=1,max=255"` // 拥有者
}

func (req *GetUserFactoryAllReq) GetPossessor() string {
	if req.Possessor == nil {
		return ""
	}
	return *req.Possessor
}

type UserFactoryInfo struct {
	ChainName        string `json:"chainName"`        // 链名
	ChainMagic       string `json:"chainMagic"`       // 链网络标识
	FactoryId        string `json:"factoryId"`        // 模板ID
	FactoryName      string `json:"factoryName"`      // 模板名称
	NumberOfEntities int    `json:"numberOfEntities"` // 持有的 entity 数量
}

type GetUserFactoryAllRes struct {
	Factories []UserFactoryInfo `json:"factories"` // factory 列表
}

// 获取用户持有的（某个 factory 下的）所有 entity
//
// method: get
//
// path: /api/entity/factory/entity/all
type GetUserFactoryEntityAllReq struct {
	Possessor  *string `form:"possessor" json:"possessor" binding:"required,min=1,max=255"`                     // 拥有者
	ChainName  *string `form:"chainName" json:"chainName" binding:"required,min=1,max=30"`                      // 链名
	ChainMagic *string `form:"chainMagic" json:"chainMagic" binding:"required,min=1,max=30"`                    // 链网络标识
	FactoryId  *string `form:"factoryId,omitempty" json:"factoryId,omitempty" binding:"omitempty,min=1,max=50"` // 模板ID
	Type       *int    `form:"type,omitempty" json:"type,omitempty" binding:"omitempty,min=1,max=2"`            // 类型(1:普通,2:限量,默认:1)
}

func (req *GetUserFactoryEntityAllReq) GetPossessor() string {
	if req.Possessor == nil {
		return ""
	}
	return *req.Possessor
}
func (req *GetUserFactoryEntityAllReq) GetChainName() string {
	if req.ChainName == nil {
		return ""
	}
	return *req.ChainName
}
func (req *GetUserFactoryEntityAllReq) GetChainMagic() string {
	if req.ChainMagic == nil {
		return ""
	}
	return *req.ChainMagic
}
func (req *GetUserFactoryEntityAllReq) GetFactoryId() string {
	if req.FactoryId == nil {
		return ""
	}
	return *req.FactoryId
}
func (req *GetUserFactoryEntityAllReq) GetType() int {
	if req.Type == nil {
		return 0
	}
	return *req.Type
}

type SubEntity struct {
	EntityId         string `json:"entityId"`         // 非同质资产ID
	TaxCollector     string `json:"taxCollector"`     // 收税人
	TaxAssetPrealnum string `json:"taxAssetPrealnum"` // 缴纳数量
}
type UserFactoryEntityInfo struct {
	ChainName        string      `json:"chainName"`        // 链名
	ChainMagic       string      `json:"chainMagic"`       // 链网络标识
	FactoryId        string      `json:"factoryId"`        // 模板ID
	FactoryName      string      `json:"factoryName"`      // 模板名称
	EntityId         string      `json:"entityId"`         // 非同质资产ID
	Possessor        string      `json:"possessor"`        // 拥有者
	TaxCollector     string      `json:"taxCollector"`     // 收税人
	TaxAssetPrealnum string      `json:"taxAssetPrealnum"` // 缴纳数量
	Type             int         `json:"type"`             // 类型(1:普通,2:限量,默认:1)
	Hash             string      `json:"hash"`             // blob哈希
	Extension        string      `json:"extension"`        // blob扩展名
	SubEntities      []SubEntity `json:"subEntities"`      // 子集
}

func NewUserFactoryEntityInfo(entity model.Entity) UserFactoryEntityInfo {
	userEntity := UserFactoryEntityInfo{}
	userEntity.Possessor = entity.Possessor
	userEntity.ChainName = entity.ChainName
	userEntity.ChainMagic = entity.ChainMagic
	userEntity.FactoryName = entity.FactoryName
	userEntity.FactoryId = entity.FactoryId
	userEntity.EntityId = entity.EntityId
	userEntity.TaxCollector = entity.TaxCollector
	userEntity.TaxAssetPrealnum = entity.TaxAssetPrealnum
	userEntity.Type = entity.Type
	userEntity.Hash = entity.Hash
	userEntity.Extension = entity.Extension
	userEntity.SubEntities = []SubEntity{}
	return userEntity
}

type GetUserFactoryEntityAllRes struct {
	Entities []UserFactoryEntityInfo `json:"entities"` // factory/entites 列表
}
