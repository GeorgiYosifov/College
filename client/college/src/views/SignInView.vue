<script>
import { inject } from 'vue'

export default {
  name: 'SignIn',
  data() {
    return {
      username: "",
      password: ""
    }
  },
  setup() {
    const axios = inject('axios')
  },
  methods: {
    submit() {
      const credentials = { username: this.username, password: this.password };

      // this.axios.defaults.headers.common["Authorization"] = "Bearer " + localStorage.getItem("token")

      this.axios
        .post(`http://localhost/api/signIn`, credentials)
        .then((response) => {
          if (response.status === 200) {
            localStorage.setItem("token", response.data["token"]);
          }
        });
    }
  }
}
</script>

<template>
  <section class="signin">
    <h1>Sign in</h1>
    <form>
      <div class="txt-field">
        <input v-model="username" type="text" required />
        <label>Username</label>
      </div>
      <div class="txt-field">
        <input v-model="password" type="password" required />
        <label>Password</label>
      </div>
      <button @click="submit" type='button'>Sign in</button>
    </form>
  </section>
</template>
