<script>
import { inject } from 'vue'

export default {
  name: 'studentSemesters',
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
  <section id="" class="login-students">
      <article class="past-semesters">
        <h1 class="secondary-title">Semesters</h1>
          <h4 class="semesters" v-for="semester in semesters">
            <RouterLink :to="`/studentEvaluations/${semester}`">semester {{ semester }}</RouterLink>
          </h4>
      </article>
  </section>
</template>