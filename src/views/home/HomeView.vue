<script setup lang="ts">
import CustomersSearchForm from '@/components/forms/CustomersSearchForm.vue'
import BaseTable from '@/components/ui/Table/BaseTable.vue'
import useApiFetch from '@/composables/useApiFetch'

const { data, error, isFetching } = useApiFetch('/customers').get().json()
</script>

<template>
  <div class="h-screen flex flex-col items-center gap-8">
    <!-- Search inputs -->
    <CustomersSearchForm />
    <!-- Table results -->
    <div class="w-full flex justify-center items-center">
      <BaseTable :data="data" :isFetching="isFetching" />
      <div v-if="error" class="">
        <p>Whoops! Something's wrong with the fetching broski...</p>
      </div>
    </div>
  </div>
</template>
