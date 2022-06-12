<template>
  <section class="signin">
    <h1>Sign up</h1>
    <form method="post">
      <div class="txt-field">
        <input type="text" required />
        <label>Username</label>
      </div>
      <div class="txt-field">
        <input type="email" required />
        <label>Email</label>
      </div>
      <div class="txt-field">
        <input type="password" required />
        <label>Password</label>
      </div>
      <div class="txt-field">
        <input type="password" required />
        <label>Confirm Password</label>
      </div>
      <div class="select">
        <select>
          <option value="0">Log in as:</option>
          <option value="1">1</option>
          <option value="2">2</option>
          <option value="3">3</option>
        </select>
      </div>
      <div class="checkbox"> <input type="checkbox"> Remember me </div>
      <input type="submit" value="Sign up" />
      <div class="signup-link">You have an account?<RouterLink to="/signIn">Sign in</RouterLink></div>
    </form>
  </section>
</template>