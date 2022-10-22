package shared

// import (
// 	"net/http"

// 	"github.com/go-openapi/runtime"
// 	"github.com/go-openapi/runtime/middleware"
// )

// type CustomResponder struct {
// 	responder middleware.Responder
// }

// func NewCustomResponder(responder middleware.Responder) *CustomResponder {
// 	return &CustomResponder{
// 		responder: responder,
// 	}
// }

// func (r *CustomResponder) WriteResponse(rw http.ResponseWriter, p runtime.Producer) {
// 	cookie := http.Cookie{Name: "robin", Value: "abcd"}
// 	http.SetCookie(rw, &cookie)
// 	r.responder.WriteResponse(rw, p)
// }
