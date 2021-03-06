package route

import (
	"github.com/ethicnology/uqac-851-software-engineering-api/database/model"
	"github.com/ethicnology/uqac-851-software-engineering-api/http/controller/bank"
	"github.com/ethicnology/uqac-851-software-engineering-api/http/controller/invoice"
	"github.com/ethicnology/uqac-851-software-engineering-api/http/controller/operation"
	"github.com/ethicnology/uqac-851-software-engineering-api/http/controller/transfer"
	"github.com/ethicnology/uqac-851-software-engineering-api/http/controller/user"
	"github.com/ethicnology/uqac-851-software-engineering-api/http/middleware"

	"goyave.dev/goyave/v3"
	"goyave.dev/goyave/v3/auth"
	"goyave.dev/goyave/v3/cors"
	"goyave.dev/goyave/v3/middleware/ratelimiter"
)

// Register all the application routes. This is the main route registrer.
func Register(router *goyave.Router) {
	// Applying default CORS settings (allow all methods and all origins)
	router.CORS(cors.Default())
	router.Middleware(ratelimiter.New(model.RateLimiterFunc))

	authenticator := auth.Middleware(&model.User{}, &auth.JWTAuthenticator{})
	myRoutes(router, authenticator)
}

func myRoutes(parent *goyave.Router, authenticator goyave.Middleware) {
	authRouter := parent.Subrouter("/auth")

	jwtController := auth.NewJWTController(&model.User{})
	jwtController.UsernameField = "email"
	authRouter.Post("/register", user.Store).Validate(user.Register)
	authRouter.Post("/login", jwtController.Login).Validate(user.Login)

	userRouter := parent.Subrouter("/users/{email}")
	userRouter.Middleware(authenticator)
	userRouter.Middleware(middleware.Owner)
	userRouter.Get("/", user.Show)
	userRouter.Patch("/", user.Update)
	userRouter.Delete("/", user.Destroy)
	userRouter.Get("/verify/{verification_code}", user.Verify)

	bankRouter := userRouter.Subrouter("/banks")
	bankRouter.Get("/", bank.Index)
	bankRouter.Post("/", bank.Store)
	bankIdRouter := bankRouter.Subrouter("/{bank_id:[0-9]+}")
	bankIdRouter.Middleware(middleware.BankOwner)
	bankIdRouter.Get("/", bank.Show)
	bankIdRouter.Delete("/", bank.Destroy)

	operationRouter := bankIdRouter.Subrouter("/operations")
	operationRouter.Get("/", operation.Index)
	operationIdRouter := operationRouter.Subrouter("/{operation_id:[0-9]+}")
	operationIdRouter.Get("/", operation.Show)

	invoiceRouter := bankIdRouter.Subrouter("/invoices")
	invoiceRouter.Get("/", invoice.Index)
	invoiceRouter.Post("/", invoice.Store).Validate(invoice.Post)
	invoiceIdRouter := invoiceRouter.Subrouter("/{invoice_id:[0-9]+}")
	invoiceIdRouter.Get("/", invoice.Show)
	invoiceIdRouter.Patch("/", invoice.Update).Validate(invoice.Patch)
	invoiceIdRouter.Delete("/", invoice.Destroy)

	transferRouter := bankIdRouter.Subrouter("/transfers")
	transferRouter.Get("/", transfer.Index)
	transferRouter.Post("/", transfer.Store).Validate(transfer.Post)
	transferIdRouter := transferRouter.Subrouter("/{transfer_id:[0-9]+}")
	transferIdRouter.Get("/", transfer.Show)
	transferIdRouter.Get("/verify/{answer}", transfer.Verify)
	transferIdRouter.Delete("/", transfer.Destroy)
}
