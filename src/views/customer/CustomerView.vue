<script setup lang="ts">
import { useRoute } from 'vue-router'

console.log('customer view')
const route = useRoute()
console.log('full path: ', route.fullPath)
</script>

<template>
  <div class="h-svh">
    <p>Customer View</p>
    <p>Router full path: {{ route.params }}</p>
  </div>
</template>
