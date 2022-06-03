<template>
  <section class="signin">
    <h1>Sign in</h1>
    <form method="post">
      <div class="txt-field">
        <input type="text" required />
        <label>Username</label>
      </div>
      <div class="txt-field">
        <input type="password" required />
        <label>Password</label>
      </div>
      <div class="select">
        <select>
          <option value="0">Log in as:</option>
          <option value="1">1</option>
          <option value="2">2</option>
          <option value="3">3</option>
        </select>
      </div>
      <div class="forgot-password">Forgot Password?</div>
      <input type="submit" value="Sign in" />
      <div class="signup-link">You do not have an account?<a href="signUp.html"> Sign up</a></div>
    </form>
  </section>
</template>
