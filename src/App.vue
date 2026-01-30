<script setup lang="ts">
import NavBar from '@/components/header/NavBar.vue'
</script>

<template>
  <main class="flex flex-col gap-4 h-svh">
    <!-- Header -->
    <NavBar />
    <div class="p-4">
      <RouterView />
    </div>
  </main>
</template>
