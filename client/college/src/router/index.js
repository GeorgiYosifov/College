import { createRouter, createWebHistory } from "vue-router";
import jwt_decode from "jwt-decode";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "home",
      component: () => import("../views/HomeView.vue"),
    },
    {
      path: "/homeStudent",
      name: "homeStudent",
      component: () => import("../views/ContactView.vue"),
    },
    {
      path: "/signIn",
      name: "signIn",
      component: () => import("../views/SignInView.vue"),
    },
    {
      path: "/signUp",
      name: "signUp",
      component: () => import("../views/SignUpView.vue"),
    },
    {
      path: "/contact",
      name: "contact",
      component: () => import("../views/ContactView.vue"),
    },
  ],
});

router.beforeEach((to, from, next) => {
    let token = localStorage.getItem("token")
    if (to.name === "home" && token !== null) {
      let tokenObj = jwt_decode(token)
      if (tokenObj["role"] === "Admin") {
        next({ name: 'homeStudent' });
      }
    }

    next();
});

export default router;
