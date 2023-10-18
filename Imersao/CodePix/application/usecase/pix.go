package usecase

import (
	"github.com/alexandrecpedro/fullcycle/Imersao/CodePix/domain/model"
)

type PixUseCase struct {
	PixKeyRepository model.PixKeyRepositoryInterface
}

//* UseCase make all the flow
func (p *PixUseCase) RegisterKey(key string, kind string, accountId string) (*model.PixKey, error) {
	// COMMENT: Verify if account exists in DB
	account, err := p.PixKeyRepository.FindAccount(accountId)

	if err != nil {
		return nil, err
	}

	// COMMENT: create a new pix_key
	pixKey, err := model.NewPixKey(kind, account, key)
	
	if err != nil {
		return nil, err
	}
	
	// COMMENT: register the new pix_key on DB
	p.PixKeyRepository.RegisterKey(pixKey)
	
	if pixKey.ID == "" {
		return nil, err
	}

	return pixKey, nil
}

func (p *PixUseCase) FindKey(key string, kind string) (*model.PixKey, error) {
	pixKey, err := p.PixKeyRepository.FindKeyByKind(key, kind)
	
	if err != nil {
		return nil, err
	}
	
	return pixKey, nil
}
