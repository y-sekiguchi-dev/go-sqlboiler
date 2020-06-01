package person

import (
	"github.com/gin-gonic/gin"
	"go-sqlboiler/usecase/person"
)

type Controller interface {
	get(c *gin.Context)
	post(c *gin.Context)
	put(c *gin.Context)
}

type impl struct {
	useCase person.UseCase
}

func NewPersonController(useCase person.UseCase) Controller {
	return &impl{useCase: useCase}
}

func (i *impl) get(c *gin.Context) {
	// TODO
}

func (i *impl) post(c *gin.Context) {
	// TODO
}

func (i *impl) put(c *gin.Context) {
	// TODO
}