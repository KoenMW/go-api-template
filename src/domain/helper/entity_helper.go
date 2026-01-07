package helper

import "go-api/domain/model"

func EntityToDTO[E model.BaseEntity, D model.BaseDTO[E]](entity E) D {
	var dto D = model.NewEntity[D]()
	dto.RecieveEntity(entity)
	return dto
}
