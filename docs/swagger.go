// Package classification awesome.
//
// Documentation of our awesome API.
//
//     Schemes: http
//     BasePath: /
//     Version: 1.0.0
//     Host: some-url.com
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Security:
//     - basic
//
//    SecurityDefinitions:
//    basic:
//      type: basic
//
// swagger:meta
package docs

import (
	"github.com/patriya-piyawiroj/banana-ducks/internal/app/bnnd"
)

// swagger:route GET /v1/banana-ducks banana-ducks getAllBananaDucks
// Get all banana ducks.
// responses:
//   200: bananaDuckArray
// Returns all ducks ([]BananaDuck)

// swagger:route POST /v1/banana-ducks banana-ducks createBananaDuck
// Creates one banana duck.
// responses:
//   200: bananaDuck
// Returns created banana duck (BananaDuck)

// swagger:response bananaDuckArray
type bananaDuckArrayWrapper struct {
	// in:body
	Body []bnnd.BananaDuck
}

// swagger:response bananaDuck
type bananaDuckWrapper struct {
	// in:body
	Body bnnd.BananaDuck
}

// swagger:parameters createBananaDuck
type createBananaDuckParamsWrapper struct {
	// Banana Duck to be created
	// in:body
	Body bnnd.BananaDuck
}
