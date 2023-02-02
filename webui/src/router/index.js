import {createRouter, createWebHashHistory} from 'vue-router'
import doLogin from '../views/doLogin.vue'
import getUserProfile from '../views/getUserProfile.vue'
import getMyStream from '../views/getMyStream.vue'
import uploadPhoto from '../views/uploadPhoto.vue'
import uploadLogo from '../views/uploadLogo.vue'
import getImage from '../views/getImage.vue'
import commentPhoto from '../views/commentPhoto.vue'
import likePhoto from '../views/likePhoto.vue'
import setMyUserName from '../views/setMyUserName.vue'
import followUser from '../views/followUser.vue'
import banUser from '../views/banUser.vue'
import unfollowUser from '../views/unfollowUser.vue'
import unbanUser from '../views/unbanUser.vue'


const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/doLogin', component: doLogin},
		{path: '/searchUser', component: getUserProfile},
		{path: '/getUserStream', component: getMyStream},
		{path: '/uploadPhoto', component: uploadPhoto},
		{path: '/uploadLogo', component: uploadLogo},
		{path: '/getImage', component: getImage},
		{path: '/commentPhoto', component: commentPhoto},
		{path: '/likePhoto', component: likePhoto},
		{path: '/setMyUserName', component: setMyUserName},
		{path: '/followUser', component: followUser},
		{path: '/unfollowUser', component: unfollowUser},
		{path: '/banUser', component: banUser},
		{path: '/unbanUser', component: unbanUser},
	]
})

export default router
