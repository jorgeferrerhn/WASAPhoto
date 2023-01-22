package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	//   Register routes

	rt.router.POST("/session", rt.wrap(rt.doLogin))
	rt.router.POST("/users/:id/uploadPhoto", rt.wrap(rt.uploadPhoto))
	rt.router.POST("/users/:id/uploadLogo", rt.wrap(rt.uploadLogo))
	rt.router.POST("/users/:id/commentPhoto/:photoId", rt.wrap(rt.commentPhoto))
	rt.router.PUT("/users/:id/likePhoto/:photoId", rt.wrap(rt.likePhoto))
	rt.router.PUT("/users/:id/setMyUserName", rt.wrap(rt.setMyUserName))
	rt.router.PUT("/users/:id/followUser/:id2", rt.wrap(rt.followUser))
	rt.router.PUT("/users/:id/banUser/:id2", rt.wrap(rt.banUser))
	rt.router.GET("/users/:id/getUserProfile", rt.wrap(rt.getUserProfile))
	rt.router.GET("/users/:id/getMyStream", rt.wrap(rt.getMyStream))
	rt.router.GET("/users/:id/getLogo", rt.wrap(rt.getLogo))
	rt.router.GET("/images/:id", rt.wrap(rt.getImage))
	rt.router.DELETE("/users/:id/unfollowUser/:id2", rt.wrap(rt.unfollowUser))
	rt.router.DELETE("/users/:id/unbanUser/:id2", rt.wrap(rt.unbanUser))
	rt.router.DELETE("/users/:id/unlikePhoto/:photoId", rt.wrap(rt.unlikePhoto))
	rt.router.DELETE("/users/:id/uncommentPhoto/:commentId", rt.wrap(rt.uncommentPhoto))
	rt.router.DELETE("/users/:id/deletePhoto/:photoId", rt.wrap(rt.deletePhoto))

	//   Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
