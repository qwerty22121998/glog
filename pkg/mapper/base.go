package mapper

import (
	"github.com/qwerty22121998/glog/pkg/dto"
	"github.com/qwerty22121998/glog/pkg/model"
)

type BaseMapper int

var Base BaseMapper

func (BaseMapper) ToDTO(base model.BaseModel) dto.BaseDTO {
	return dto.BaseDTO{
		ID:        base.ID,
		CreatedAt: base.CreatedAt,
		UpdatedAt: base.UpdatedAt,
		DeletedAt: base.DeletedAt,
		CreatedBy: base.CreatedBy,
		UpdatedBy: base.UpdatedBy,
	}
}

func (BaseMapper) ToModel(base dto.BaseDTO) model.BaseModel {
	return model.BaseModel{
		ID:        base.ID,
		CreatedAt: base.CreatedAt,
		UpdatedAt: base.UpdatedAt,
		DeletedAt: base.DeletedAt,
		CreatedBy: base.CreatedBy,
		UpdatedBy: base.UpdatedBy,
	}
}
