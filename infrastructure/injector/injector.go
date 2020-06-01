package injector

import (
	"go-sqlboiler/infrastructure/database/sqlboiler"
	controllers "go-sqlboiler/interfaces/controllers/person"
	useCase "go-sqlboiler/usecase/person"
)

type injector struct {
	boilInjector sqlboiler.Injector
}

func (i *injector) NewPersonController() controllers.Controller {
	return controllers.NewPersonController(useCase.NewUseCase(
		i.boilInjector.NewPersonRepository(),
		i.boilInjector.TransactionProvider()))
}