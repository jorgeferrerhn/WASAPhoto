import {createRouter, createWebHashHistory} from 'vue-router'
import doLogin from '../views/doLogin.vue'
import getUserProfile from '../views/getUserProfile.vue'
import getMyStream from '../views/getMyStream.vue'
import uploadPhoto from '../views/uploadPhoto.vue'
import uploadLogo from '../views/uploadLogo.vue'
import getImage from '../views/getImage.vue'
import setMyUserName from '../views/setMyUserName.vue'
import getLoggedProfile from '../views/getLoggedProfile.vue'

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/', component: doLogin},
		{path: '/getUserProfile', component: getUserProfile},
		{path: '/getMyStream', component: getMyStream},
		{path: '/uploadPhoto', component: uploadPhoto},
		{path: '/uploadLogo', component: uploadLogo},
		{path: '/getImage', component: getImage},
		{path: '/setMyUserName', component: setMyUserName},
		{path: '/getLoggedProfile', component: getLoggedProfile},



	]
})

export default router
