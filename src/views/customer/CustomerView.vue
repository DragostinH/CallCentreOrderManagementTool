<script setup lang="ts">
import { useRoute } from 'vuetify/lib/composables/router'
import useApiFetch from '@/composables/useApiFetch'
import { onBeforeMount } from 'vue'

// pulling customer data with the customerNumber from the route
const route = useRoute()
onBeforeMount(() => {
  try {
    const { data, error } = useApiFetch(
      `/customers?customerNumber=${route.value.params.customerId}`,
    )
    console.log('customer data', data.value)
  } catch (error) {
    console.error(error)
  }
})

const fetchCustomer = () => {}
</script>

<template>
  <div class="card lg:card-side bg-neutral-800 border border-secondary shadow-sm text-neutral-200">
    <div class="card-body">asdasd</div>
  </div>
</template>
