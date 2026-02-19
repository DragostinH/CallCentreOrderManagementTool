<script setup lang="ts">
import { useRoute } from 'vue-router'

const route = useRoute()
</script>

<template>
  <p>product view</p>
  <p>{{ route.params.product_id }}</p>
</template>
