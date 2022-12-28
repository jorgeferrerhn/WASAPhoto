package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes

	rt.router.POST("/session", rt.wrap(rt.doLogin))
	rt.router.GET("/users/:id/getUserProfile", rt.wrap(rt.getUserProfile))
	rt.router.GET("/users/:id/getMyStream", rt.wrap(rt.getMyStream))
	rt.router.GET("/users/:id/getLogo", rt.wrap(rt.getLogo))

	//rt.router.PUT("/fountains/:id", rt.wrap(rt.updateFountain))
	//rt.router.DELETE("/fountains/:id", rt.wrap(rt.deleteFountain))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
