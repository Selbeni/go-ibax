/*---------------------------------------------------------------------------------------------
 *  Copyright (c) IBAX. All rights reserved.
 *  See LICENSE in the project root for license information.

	TxHash     string `gorm:"not null" json:"tx_hash"`
	ChainState int64  `gorm:"not null" json:"chain_state"`
	BlockId    int64  `gorm:"not null" json:"block_id"`
	ChainId    int64  `gorm:"not null" json:"chain_id"`
	ChainErr   string `gorm:"not null" json:"chain_err"`

	UpdateTime int64 `gorm:"not null" json:"update_time"`
	CreateTime int64 `gorm:"not null" json:"create_time"`
}

func (VDEDestDataLog) TableName() string {
	return "vde_dest_data_log"
}

func (m *VDEDestDataLog) Create() error {
	return DBConn.Create(&m).Error
}

func (m *VDEDestDataLog) Updates() error {
	return DBConn.Model(m).Updates(m).Error
}

func (m *VDEDestDataLog) Delete() error {
	return DBConn.Delete(m).Error
}

func (m *VDEDestDataLog) GetAll() ([]VDEDestDataLog, error) {
	var result []VDEDestDataLog
	err := DBConn.Find(&result).Error
	return result, err
}
func (m *VDEDestDataLog) GetOneByID() (*VDEDestDataLog, error) {
	err := DBConn.Where("id=?", m.ID).First(&m).Error
	return m, err
}

func (m *VDEDestDataLog) GetAllByTaskUUID(TaskUUID string) ([]VDEDestDataLog, error) {
	result := make([]VDEDestDataLog, 0)
	err := DBConn.Table("vde_dest_data_log").Where("task_uuid = ?", TaskUUID).Find(&result).Error
	return result, err
}

func (m *VDEDestDataLog) GetOneByTaskUUID(TaskUUID string) (*VDEDestDataLog, error) {
	err := DBConn.Where("task_uuid=?", TaskUUID).First(&m).Error
	return m, err
}

func (m *VDEDestDataLog) GetAllByChainState(ChainState int64) ([]VDEDestDataLog, error) {
	result := make([]VDEDestDataLog, 0)
	err := DBConn.Table("vde_dest_data_log").Where("chain_state = ?", ChainState).Find(&result).Error
	return result, err
}

func (m *VDEDestDataLog) GetOneByChainState(ChainState int64) (bool, error) {
	return isFound(DBConn.Where("chain_state = ?", ChainState).First(m))
}