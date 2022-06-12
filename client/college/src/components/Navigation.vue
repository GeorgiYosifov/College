<template>
  <nav class="nav" id="idnav">
    <ul class="nav-menu" id="ulId">
      <li class="nav-logo"><img src="@/assets/img/logo.svg" alt="logo" /></li>
      <li class="nav-icon"><i class="fa-solid fa-bars"></i></li>
      <li class="nav-links">
        <RouterLink to="/">
          <span>Home</span>
        </RouterLink>
      </li>
      <li class="nav-links">
        <a href="#">Faculties &#9207;</a>
        <ul class="dropdown">
          <li class="nav-links"><a href="#">Faculties1</a></li>
        </ul>
      </li>
      <li class="nav-links">
        <a href="#">Departments &#9207;</a>
        <ul class="dropdown">
          <li class="nav-links">
            <a href="#">Departments1</a>
          </li>
        </ul>
      </li>
      <li class="nav-links"><a href="#">About us</a></li>
      <li class="nav-links">
        <RouterLink to="/contact">Contact us</RouterLink>
      </li>
      <li class="nav-links btn">
        <RouterLink to="/signIn">Sign in</RouterLink>
      </li>
    </ul>
  </nav>
</template>
