package manager

import "github.com/enigma-mart/usecase"

type UseCaseManager interface {
	ProductUseCase() usecase.ProductUseCase
}

type useCaseManager struct {
	repoManeger RepositoryManager
}

func (u *useCaseManager) ProductUseCase() usecase.ProductUseCase {
	return usecase.NewProductUseCase(u.repoManeger.ProductRepo())
}

func NewUseCaseManager(repoManager RepositoryManager) UseCaseManager {
	return &useCaseManager{
		repoManeger: repoManager,
	}
}
