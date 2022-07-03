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
      path: "/studentSemesters",
      name: "studentSemesters",
      component: () => import("../views/Student/SemestersView.vue"),
    },
    {
      path: "/studentCurriculum",
      name: "studentCurriculum",
      component: () => import("../views/Student/CurriculumView.vue"),
    },
    {
      path: "/studentEvaluations",
      name: "studentEvaluations",
      component: () => import("../views/Student/EvaluationsView.vue"),
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
      if (tokenObj["role"] === "Student") {
        next({ name: 'studentSemesters' });
      } else if (tokenObj["role"] === "Teacher") {
        next({ name: 'teacherEvaluations' });
      } else if (tokenObj["role"] === "Rector") {
        next({ name: 'rector' });
      }
    }

    next();
});

export default router;
