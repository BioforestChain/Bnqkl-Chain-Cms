package model

import (
	"time"
)

type TableSchema interface {
	GetTableOptions() string // 建表参数
}

type IDModel struct {
	ID uint `gorm:"type:bigint(20) unsigned NOT NULL AUTO_INCREMENT;comment:主键ID;primarykey" json:"id"`
}

func (m *IDModel) GetBaseTableOptions() string {
	return " ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT="
}

func (m *IDModel) GetIdColumnName() string {
	return "id"
}

type BasicWithTimeModel struct {
	CreatedAt time.Time `gorm:"type:datetime NOT NULL;default:CURRENT_TIMESTAMP;comment:创建时间" json:"createAt"`
	UpdatedAt time.Time `gorm:"type:datetime NOT NULL;default:CURRENT_TIMESTAMP;comment:修改时间" json:"updatedAt"`
}

func (m *BasicWithTimeModel) GetCreatedAtColumnName() string {
	return "created_at"
}
func (m *BasicWithTimeModel) GetUpdatedAtColumnName() string {
	return "updated_at"
}

type BasicWithDelFlagModel struct {
	DelFlag bool `gorm:"type:tinyint(2) unsigned NOT NULL;default:1;comment:逻辑删除标志(1:正常,2:删除,默认:1)" json:"delFlag"`
}

func (m *BasicWithTimeModel) GetDelFlagColumnName() string {
	return "del_flag"
}
