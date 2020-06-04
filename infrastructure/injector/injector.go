package injector

import (
	injector2 "go-sqlboiler/infrastructure/database/sqlboiler/injector"
	controllers "go-sqlboiler/interfaces/controllers/person"
	useCase "go-sqlboiler/usecase/person"
)

type injector struct {
	boilInjector injector2.Injector
}

func (i *injector) NewPersonController() controllers.Controller {
	return controllers.NewPersonController(useCase.NewUseCase(
		i.boilInjector.NewPersonRepository(),
		i.boilInjector.TransactionProvider()))
}
