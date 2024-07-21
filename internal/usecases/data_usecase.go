package usecases

import (
	"github.com/Nchezhegova/gophkeeper/internal/entities"
	"github.com/Nchezhegova/gophkeeper/internal/interfaces/repository"
)

type DataUseCase struct {
	DataRepository repository.DataRepository
}

func (uc *DataUseCase) StoreData(userID uint32, key, dataType string, data []byte) error {
	return uc.DataRepository.StoreData(userID, key, dataType, data)
}

func (uc *DataUseCase) GetData(userID uint32, key, dataType, identifier string) ([]entities.Data, error) {
	return uc.DataRepository.GetData(userID, key, dataType, identifier)
}

func (uc *DataUseCase) UpdateData(userID uint32, key, dataType, identifier string, newData []byte) error {
	return uc.DataRepository.UpdateData(userID, key, dataType, identifier, newData)
}

func (uc *DataUseCase) DeleteData(userID uint32, key, dataType, identifier string) error {
	return uc.DataRepository.DeleteData(userID, key, dataType, identifier)
}
