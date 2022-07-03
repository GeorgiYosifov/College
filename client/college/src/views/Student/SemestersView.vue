<script>
import { inject } from 'vue'

export default {
  name: 'Semesters',
  data() {
    return {
      semesters: {}
    }
  },
  setup() {
    const axios = inject('axios')
  },
  created() {
    this.getSemesters()
  },
  methods: {
    getSemesters() {
      this.axios.defaults.headers.common["Authorization"] = "Bearer " + localStorage.getItem("token")

      this.axios
        .get(`http://localhost/api/semesters`)
        .then((response) => {
          if (response.status === 200) {
            this.semesters = response.data
          }
        });
    }
  }
}
</script>

<template>
  <ul class="secondary-menu">
    <li class="nav-links">
      <RouterLink to="/studentSemesters"><span>Semesters</span></RouterLink>
    </li>
    <li class="nav-links">
      <RouterLink to="/studentCurriculum"><span>Curriculum</span></RouterLink>
    </li>
    <li class="nav-links">
      <RouterLink to="/studentEvaluations"><span>Evaluations</span></RouterLink>
    </li>
  </ul>

  <section id="" class="login-students">
      <article class="past-semesters">
        <h1 class="secondary-title">Semesters</h1>
          <h4 class="semesters" v-for="semester in semesters">
            <RouterLink to="/studentCurriculum">semester {{ semester }}</RouterLink>
          </h4>
      </article>
  </section>
</template>