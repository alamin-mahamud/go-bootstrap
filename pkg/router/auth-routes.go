package router

import (
	"github.com/alamin-mahamud/xarvis/pkg/database"
	"github.com/alamin-mahamud/xarvis/pkg/handler"
	"github.com/alamin-mahamud/xarvis/pkg/repository"
	"github.com/alamin-mahamud/xarvis/pkg/usecase"
)

func getAuthUserRoutes() Routes {
	mgoUserCollection := database.NewDB()
	userRepo := repository.NewUser(mgoUserCollection)
	userUsecase := usecase.NewUser(userRepo)
	userHandler := handler.NewUser(userUsecase)

	return Routes{
		// Public Routes
		Route{"List", "GET", "", userHandler.List},
		Route{"Get", "GET", "/{id}", userHandler.Get},
		// Protected Routes
		Route{"Create", "POST", "", userHandler.Create},
		Route{"Update", "PATCH", "/{id}", userHandler.Update},
		Route{"Delete", "DELETE", "/{id}", userHandler.Delete},
	}
}

/* func getAuthJWTRoutes() Routes {
	mgoUserCollection := database.NewDB()
	jwtRepo := repository.NewAuthJWT(mgoUserCollection)
	jwtUsecase := usecase.NewAuthJWT(jwtRepo)
	jwtHandler := handler.NewAuthJWT(jwtUsecase)

	return Routes{}

}
*/
