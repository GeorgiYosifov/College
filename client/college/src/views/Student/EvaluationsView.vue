<script>
import { inject } from 'vue'

export default {
  name: 'Evaluations',
  data() {
    return {
      courses: {}
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
        .get(`http://localhost/api/courses`)
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
          <tr>
            <td>001</td>
            <td>Linear algebra</td>
            <td>T1 (Test work) <br>
              T2 (Test work) <br>
              T3 (Presence and participation) <br>
              TO (Final assessment) <br>
              I1 (Exam)
            </td>
          </tr>
        </tbody>
      </table>
    </article>
  </section>
</template>