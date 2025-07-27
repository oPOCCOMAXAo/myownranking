package swagger

import (
	"github.com/gin-gonic/gin"
	"github.com/opoccomaxao/myownranking/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// fullInstance godoc
//
//	@title						MyOwnRanking API Full.
//	@version					1.0
//	@description				My Own Ranking API server.
//	@termsOfService				http://swagger.io/terms/
//	@basepath					/
//
//	@securityDefinitions.apikey	StdAuth
//	@name						Authorization
//	@in							header
func fullInstance() gin.HandlerFunc {
	return ginSwagger.WrapHandler(
		swaggerfiles.NewHandler(),
		ginSwagger.InstanceName(docs.SwaggerInfofull.InstanceName()),
		ginSwagger.PersistAuthorization(true),
	)
}
