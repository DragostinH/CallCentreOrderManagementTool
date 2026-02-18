<script setup lang="ts">
import CustomerViewFields from '@/components/customer/CustomerViewFields.vue'
import CustomerOrdersTable from '@/components/ui/Table/CustomerOrdersTable.vue'
import useApiFetch from '@/composables/useApiFetch'
import { RouterView, useRoute } from 'vue-router'
const route = useRoute()
const { data, isFetching, error } = useApiFetch(`/customer/${route.params.customer_number}`)
  .get()
  .json()
</script>

<template>
  <div class="">
    <div class="grid gap-4" v-if="data">
      <CustomerViewFields :data="data" :is-fetching="isFetching" />
      <CustomerOrdersTable :orders="data?.orders" :is-fetching="isFetching" />
    </div>

    <div v-else-if="isFetching">Loading Customer Profile...</div>

    <div v-else-if="error">
      {{ error }}
    </div>
    <router-view />
  </div>
</template>
