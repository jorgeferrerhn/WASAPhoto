import {createRouter, createWebHashHistory} from 'vue-router'
import HomeView from '../views/HomeView.vue'
import getUserProfile from '../views/getUserProfile.vue'
import getMyStream from '../views/getMyStream.vue'
import uploadPhoto from '../views/uploadPhoto.vue'
import getImage from '../views/getImage.vue'



const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/', component: HomeView},
		{path: '/searchUser', component: getUserProfile},
		{path: '/getUserStream', component: getMyStream},
		{path: '/uploadPhoto', component: uploadPhoto},
		{path: '/getImage', component: getImage},

	]
})

export default router
