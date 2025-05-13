package routes

import (
	"github.com/goravel/framework/facades"

	"goravel/app/http/controllers"
	"goravel/app/http/middleware"
	"goravel/app/repositories"
	"goravel/app/services"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/route"
)

func Api() {
	facades.Route().Prefix("/api").Group(func(group route.Router) {
		// Group 1: Auth routes không cần xác thực
		AuthRoute(group)

		// Group 2: Routes cần xác thực
		group.Middleware(middleware.Auth()).Group(func(group route.Router) {
			group.Get("/user", func(c http.Context) http.Response {
				return controllers.NewAuthController(
					services.NewAuthService(
						repositories.NewAuthRepository(),
					),
				).Me(c)
			})
		})
	})
}

func AuthRoute(group route.Router) {
	authController := controllers.NewAuthController(
		services.NewAuthService(
			repositories.NewAuthRepository(),
		),
	)

	// Use the group parameter instead of facades.Route() to prevent duplicate registrations
	group.Get("/auth", func(ctx http.Context) http.Response {
		return ctx.Response().Json(http.StatusOK, http.Json{
			"Hello": "Goravel",
		})
	})
	group.Post("/register", func(ctx http.Context) http.Response {
		return authController.Register(ctx)
	})
	group.Post("/login", func(ctx http.Context) http.Response {
		return authController.Login(ctx)
	})
	group.Post("/refresh", func(ctx http.Context) http.Response {
		return authController.RefreshToken(ctx)
	})
	group.Post("/logout", func(ctx http.Context) http.Response {
		return authController.Logout(ctx)
	})
}
