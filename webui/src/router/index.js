import {createRouter, createWebHashHistory} from 'vue-router'
import HomeView from '../views/HomeView.vue'

import loginPageView from '../views/loginPageView.vue'

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/', component: loginPageView},
	]
})

export default router
