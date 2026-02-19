<script setup lang="ts">
import CustomersSearchForm from '@/components/forms/CustomersSearchForm.vue'
import BaseTable from '@/components/ui/Table/BaseTable.vue'
import useApiFetch from '@/composables/useApiFetch'
import { buildSearchParams } from '@/helpers/buildSearchParams'
import { type Customer, type FormData as SearchFormData } from '@/types/types'
import { ref } from 'vue'
const reactiveUrl = ref('')
const { data, error, isFetching, execute } = useApiFetch(reactiveUrl, {
  immediate: false,
})
  .get()
  .json<Customer[]>()

const handleSearchCustomers = (formData: SearchFormData) => {
  const params = buildSearchParams(formData)
  reactiveUrl.value = `/customers/search?${params}`
  execute()
}
</script>

<template>
  <div class="h-screen flex flex-col items-center gap-8">
    <!-- Search inputs -->
    <CustomersSearchForm @search="handleSearchCustomers" />
    <!-- Table results -->
    <div class="w-full flex justify-center items-center">
      <BaseTable v-if="data" :data="data" :isFetching="isFetching" />
      <div v-if="error" class="">
        <p>Whoops! Something's wrong with the fetching broski...</p>
      </div>
    </div>
  </div>
</template>
