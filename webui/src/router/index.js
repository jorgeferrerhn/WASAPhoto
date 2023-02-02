import {createRouter, createWebHashHistory} from 'vue-router'
import HomeView from '../views/HomeView.vue'
import getUserProfile from '../views/getUserProfile.vue'
import getMyStream from '../views/getMyStream.vue'
import uploadPhoto from '../views/uploadPhoto.vue'
import uploadLogo from '../views/uploadLogo.vue'
import getImage from '../views/getImage.vue'
import commentPhoto from '../views/commentPhoto.vue'


const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/', component: HomeView},
		{path: '/searchUser', component: getUserProfile},
		{path: '/getUserStream', component: getMyStream},
		{path: '/uploadPhoto', component: uploadPhoto},
		{path: '/uploadLogo', component: uploadLogo},
		{path: '/getImage', component: getImage},
		{path: '/commentPhoto', component: commentPhoto},

	]
})

export default router
