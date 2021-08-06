import { createRouter, createWebHistory } from 'vue-router'
import Home from '../views/Home.vue'

const routes = [
  {
    path: "",
    redirect: "/home",
    
  },
  {
    path: '/home',
    name: 'Home',
    component: Home,
  },
  {
    path: "/archive",
    name: "Archive",
    component: () => import("../views/Archive.vue"),
  },
  {
    path: "/tags",
    name: "Tags",
    component: () => import("../views/Tags.vue"),
  },
  {
    path: "/tags/:tagname",
    name: "TagsSearch",
    component: () => import("../views/TagsSearch.vue"),
  },
  {
      path: "/milestone",
      name: "Milestone",
      component: () => import("../views/Milestone.vue"),
  },
  {
    path: '/about',
    name: 'About',
    component: () => import('../views/About.vue')
  },
  {
    path: "/post/:postid",
    name: "Post",
    component: () => import("../views/Post.vue"),
  },
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
