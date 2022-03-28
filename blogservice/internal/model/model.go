package model

type Model struct {
	CreatedOn  uint32 `json:"created_on"`  // 创建时间
	CreatedBy  string `json:"created_by"`  // 创建人
	ModifiedOn uint32 `json:"modified_on"` // 修改时间
	ModifiedBy string `json:"modified_by"` // 修改人
	DeletedOn  uint32 `json:"deleted_on"`  // 删除时间
	IsDel      uint8  `json:"is_del"`      // 是否删除 0 为未删除、1 为已删除
}

