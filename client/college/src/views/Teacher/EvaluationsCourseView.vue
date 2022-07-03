<script>
import { inject } from 'vue'

export default {
  name: 'teacherEvaluationsCourse',
  data() {
    return {
      students: [],
      course: JSON.parse(this.$route.query.course)
    }
  },
  setup() {
    const axios = inject('axios')
  },
  created() {
    this.getStudents()
  },
  methods: {
    getStudents() {
      this.axios.defaults.headers.common["Authorization"] = "Bearer " + localStorage.getItem("token")

      this.axios
        .get(`http://localhost/api/teacher/students`, { params: { course: this.course }})
        .then((response) => {
          if (response.status === 200) {
            this.students = response.data
          }
        });
    }
  }
}
</script>

<template>
    <nav class="nav" id="idnav">
        <ul class="secondary-menu">
            <li class="nav-links">
                <RouterLink :to="{
                    path: '/teacherEvaluationsCourse',
                    query: { course: JSON.stringify(this.course) }
                }"><span>Evaluations</span></RouterLink>
            </li>
            <li class="nav-links">
                <RouterLink to="/teacherCourseDescription"><span>Course description</span></RouterLink>
            </li>
        </ul>
    </nav>

    <section id="" class="login-teacher">
        <article class="teacher-course">
            <h1 class="secondary-title">Еvaluations</h1>
            <table id="" class="table-informations">
                <thead>
                    <tr>
                        <th>Name</th>
                        <th>Semester</th>
                        <th>Specialty</th>
                    </tr>
                </thead>
                <tbody>
                    <tr>
                        <td>{{ this.course.name }}</td>
                        <td>{{ this.course.semesterNumber }}</td>
                        <td>{{ this.course.specialty }}</td>
                    </tr>
                </tbody>
            </table>
        </article>

        <article class="list-students">
            <h4 class="list">List of students</h4>
            <table id="" class="table-informations">
                <thead>
                    <tr>
                        <th>Name</th>
                        <th>Tasks</th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="student in students">
                        <td>{{ student.name }}</td>
                        <td>
                            <div v-for="task in student.tasks">
                              {{ task }} <i class="fa-solid fa-pen edit"></i>
                            </div>
                            <button class="btn">Save</button>
                        </td>
                    </tr>
                </tbody>
            </table>
        </article>
    </section>
</template>
