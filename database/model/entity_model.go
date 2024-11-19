package model

import "time"

// 非同质资产表
type Entity struct {
	IDModel

	Possessor        string `gorm:"type:varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL;default:'';comment:拥有者;index:idx_possessor;" json:"possessor"`
	ChainName        string `gorm:"type:varchar(30) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL;default:'';comment:链名;uniqueIndex:uniq_entity;" json:"chainName"`
	ChainMagic       string `gorm:"type:varchar(30) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL;default:'';comment:链网络标识;uniqueIndex:uniq_entity;" json:"chainMagic"`
	FactoryName      string `gorm:"type:varchar(50) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL;default:'';comment:模板名称;" json:"factoryName"`
	FactoryId        string `gorm:"type:varchar(50) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL;default:'';comment:模板ID;uniqueIndex:uniq_entity;" json:"factoryId"`
	EntityId         string `gorm:"type:varchar(50) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL;default:'';comment:非同质资产ID;uniqueIndex:uniq_entity;" json:"entityId"`
	TaxCollector     string `gorm:"type:varchar(255) NOT NULL;default:'';comment:收税人;" json:"taxCollector"`
	TaxAssetPrealnum string `gorm:"type:varchar(50)  NOT NULL;default:'0';comment:缴纳数量;" json:"taxAssetPrealnum"`
	Type             int    `gorm:"type:tinyint(2) unsigned NOT NULL;default:1;comment:类型(1:普通,2:限量,默认:1);" json:"type"`
	Hash             string `gorm:"type:varchar(100)  NOT NULL;default:'';comment:哈希;" json:"hash"`
	Extension        string `gorm:"type:varchar(50) NOT NULL;default:'';comment:blob扩展名;" json:"extension"`

	BasicWithTimeModel
	BasicWithDelFlagModel
}

func (entity *Entity) TableName() string {
	return "entity"
}
func (entity *Entity) GetTableOptions() string {
	return entity.GetBaseTableOptions() + "'非同质资产表';"
}
func (entity *Entity) GetPossessorColumnName() string {
	return "possessor"
}
func (entity *Entity) GetChainNameColumnName() string {
	return "chain_name"
}
func (entity *Entity) GetChainMagicColumnName() string {
	return "chain_magic"
}
func (entity *Entity) GetFactoryNameColumnName() string {
	return "factory_name"
}
func (entity *Entity) GetFactoryIdColumnName() string {
	return "factory_id"
}
func (entity *Entity) GetEntityIdColumnName() string {
	return "entity_id"
}
func (entity *Entity) GetTaxCollectorColumnName() string {
	return "tax_collector"
}
func (entity *Entity) GetTaxAssetPrealnumColumnName() string {
	return "tax_asset_prealnum"
}
func (entity *Entity) GetTypeColumnName() string {
	return "type"
}
func (entity *Entity) GetHashColumnName() string {
	return "hash"
}
func (entity *Entity) GetExtensionColumnName() string {
	return "extension"
}

func NewEntity(possessor string, chainName string, chainMagic string, factoryName string, factoryId string, entityId string, taxCollector string, taxAssetPrealnum string, entityType int, hash string, extension string) Entity {
	entity := Entity{}
	entity.Possessor = possessor
	entity.ChainName = chainName
	entity.ChainMagic = chainMagic
	entity.FactoryName = factoryName
	entity.FactoryId = factoryId
	entity.EntityId = entityId
	entity.TaxCollector = taxCollector
	entity.TaxAssetPrealnum = taxAssetPrealnum
	entity.Type = entityType
	entity.Hash = hash
	entity.Extension = extension
	now := time.Now()
	entity.CreatedAt = now
	entity.UpdatedAt = now
	return entity
}
