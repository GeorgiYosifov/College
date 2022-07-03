<script>
import { inject } from 'vue'

export default {
  name: 'studentEvaluations',
  data() {
    return {
      courses: []
    }
  },
  setup() {
    const axios = inject('axios')
  },
  created() {
    this.getCourses()
  },
  methods: {
    getCourses() {
      this.axios.defaults.headers.common["Authorization"] = "Bearer " + localStorage.getItem("token")

      this.axios
        .get(`http://localhost/api/` + this.$route.params.semester + `/courses`)
        .then((response) => {
          if (response.status === 200) {
            this.courses = response.data
          }
        });
    }
  }
}
</script>

<template>
  <section id="" class="login-students">
    <article class="enrollment-semestar">
      <h1 class="secondary-title">Evaluations</h1>
      <table id="" class="table-informations">
        <thead>
          <tr>
            <th>Course</th>
            <th>Teacher</th>
            <th>Evaluations</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="course in courses">
            <td>{{ course.name }}</td>
            <td>{{ course.teacher }}</td>
            <td>
              <div v-for="task in course.tasks">
                {{ task }} <br>
              </div>
            </td>
          </tr>
        </tbody>
      </table>
    </article>
  </section>
</template>