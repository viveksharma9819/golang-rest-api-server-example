package domain

import (
	"net/http"
)

type IContext interface {
	Set(r *http.Request, key interface{}, val interface{})
	Get(r *http.Request, key interface{}) interface{}
	SetRouteCtx(r *http.Request, val *Route) *Route
	GetRouteCtx(r *http.Request) *Route
	SetDbCtx(r *http.Request, val IDatabase) IDatabase
	GetDbCtx(r *http.Request) IDatabase
	SetRendererCtx(r *http.Request, val IRenderer) IRenderer
	GetRendererCtx(r *http.Request) IRenderer
	SetTokenAuthorityCtx(r *http.Request, val ITokenAuthority) ITokenAuthority
	GetTokenAuthorityCtx(r *http.Request) ITokenAuthority
	SetAuthenticatedClaimsCtx(r *http.Request, val *TokenClaims) *TokenClaims
	GetAuthenticatedClaimsCtx(r *http.Request) *TokenClaims
	SetCurrentUserCtx(r *http.Request, val *User) *User
	GetCurrentUserCtx(r *http.Request) *User
	SetCurrentObjectCtx(r *http.Request, val interface{}) interface{}
	GetCurrentObjectCtx(r *http.Request) interface{}
	SetControllerHooksMapCtx(r *http.Request, val *ControllerHooksMap) *ControllerHooksMap
	GetControllerHooksMapCtx(r *http.Request) *ControllerHooksMap

	InjectWithNext(func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc, ctx IContext)) func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc)
	Inject(handler ContextHandlerFunc) http.HandlerFunc
}
