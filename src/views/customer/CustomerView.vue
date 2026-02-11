<script setup lang="ts">
import useApiFetch from '@/composables/useApiFetch'
import { onBeforeMount } from 'vue'
import { useRoute } from 'vue-router'

// pulling customer data with the customerNumber from the route
const route = useRoute()
const { data, error } = useApiFetch(`/customers/${route.params.customerNumber}`).get().json()
</script>

<template>
  <div class="card lg:card-side bg-neutral-800 border border-secondary shadow-sm text-neutral-200">
    <div class="card-body">asdasd</div>
  </div>
</template>
