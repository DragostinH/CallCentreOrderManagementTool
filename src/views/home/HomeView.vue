<script setup lang="ts">
import { onMounted, ref, Suspense } from 'vue'
import BaseTable from '../../components/ui/BaseTable.vue'
import { useFetch } from '@vueuse/core'
import useApiFetch from '@/composables/useApiFetch'
const tableData = ref([])
onMounted(() => {
  const { data, error, isFetching } = useApiFetch('/customers').get().json()
  tableData.value = data.value
  console.log('mounted data', tableData.value)
})

console.log('data', tableData.value)
</script>

<template>
  <div class="h-screen flex flex-col items-center gap-8">
    <!-- Search inputs -->
    <div class="bg-base-200 grid grid-cols-2 p-4 w-1/2">
      <fieldset class="fieldset">
        <legend class="fieldset-legend">Name</legend>
        <input type="text" placeholder="Name" class="input" />
      </fieldset>
      <fieldset class="fieldset">
        <legend class="fieldset-legend">Customer Number</legend>
        <input type="text" placeholder="Name" class="input" />
      </fieldset>
      <fieldset class="fieldset">
        <legend class="fieldset-legend">Email</legend>
        <input type="text" placeholder="Name" class="input" />
      </fieldset>
      <fieldset class="fieldset">
        <legend class="fieldset-legend">Post Code</legend>
        <input type="text" placeholder="Name" class="input" />
      </fieldset>
      <fieldset class="fieldset">
        <legend class="fieldset-legend">Telephone Number</legend>
        <input type="text" placeholder="Name" class="input" />
      </fieldset>
      <fieldset class="fieldset">
        <legend class="fieldset-legend">Order Number</legend>
        <input type="text" placeholder="Name" class="input" />
      </fieldset>
      <button class="btn btn-primary col-span-2">Search</button>
    </div>
    <!-- Table results -->
    <BaseTable :data="tableData" />
  </div>
</template>
