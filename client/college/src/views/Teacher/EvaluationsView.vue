<script>
import { inject } from 'vue'

export default {
  name: 'teacherEvaluations',
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
        .get(`http://localhost/api/teacher/courses`)
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
    <section id="" class="login-teacher">
        <article class="teacher-course">
            <h1 class="secondary-title">Evaluations</h1>
            <div class="semestar">
                <h4 class="semestar">Choose a course</h4>
            </div>

            <table id="" class="table-informations">
                <thead>
                    <tr>
                        <th>Course</th>
                        <th>Semester</th>
                        <th>Specialty</th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="course in courses">
                        <td>
                            <RouterLink :to="{
                                path: '/teacherEvaluationsCourse',
                                query: { course: JSON.stringify(course) }
                            }">{{ course.name }}</RouterLink>
                        </td>
                        <td>{{ course.semesterNumber }}</td>
                        <td>{{ course.specialty }}</td>
                    </tr>
                </tbody>
              </table>
        </article>
    </section>
</template>
