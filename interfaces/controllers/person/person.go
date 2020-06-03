package person

import (
	"go-sqlboiler/usecase/person"

	"github.com/gin-gonic/gin"
)

type Controller interface {
	get(c *gin.Context)
	post(c *gin.Context)
	put(c *gin.Context)
}

type controller struct {
	useCase person.UseCase
}

func NewPersonController(useCase person.UseCase) Controller {
	return &controller{useCase: useCase}
}

func (i *controller) get(c *gin.Context) {
	// TODO
}

func (i *controller) post(c *gin.Context) {
	// TODO
}

func (i *controller) put(c *gin.Context) {
	// TODO
}
