package bootstrap

import "github.com/go-template/internal/user"

// IMPORT OTHER REQUIRED HANDLER
// "github.com/supertruck/wallet/internal/{domain}"

type Handlers struct {
	Auth *user.Handler
	/*IMPORT OTHER HANDLER */
}
