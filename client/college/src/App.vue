<script setup>
import { RouterView } from "vue-router";
import Navigation from "./components/Navigation.vue";
import Footer from "./components/Footer.vue";
</script>

<template>
  <header>
    <Navigation />
  </header>
  <RouterView />
  <Footer />
</template>
