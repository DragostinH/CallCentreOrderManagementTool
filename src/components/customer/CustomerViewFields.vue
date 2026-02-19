<script setup lang="ts">
import type { Customer } from '@/types/types'
import { computed } from 'vue'

const props = defineProps<{
  data: Customer
  isFetching: boolean
}>()

const fullName = computed(() => {
  return props.data?.first_name + ' ' + props.data?.last_name
})
</script>

<template>
  <div
    v-if="data"
    class="card-lg card-side w-3xl relative rounded-md bg-neutral-800 border border-secondary shadow-sm text-neutral-200 p-4"
  >
    <div
      class="border-primary-content border absolute right-0 top-0 bg-secondary-content rounded-md p-1"
    >
      <span class="text-primary">Customer №: </span><span>{{ data?.customer_number }}</span>
    </div>
    <div
      class="card-title -top-4 pointer-events-none bg-primary text-primary-content rounded-md p-1 absolute"
    >
      Customer Info
    </div>
    <div class="card-body">
      <p class="border-b-primary-content border-b text-xl">
        <span>Name: </span> <span class="font-bold">{{ fullName }}</span>
      </p>
      <p class="border-b-primary-content border-b text-xl">
        <span>Email: </span> <span class="font-bold">{{ data.email }}</span>
      </p>
      <p class="border-b-primary-content border-b text-xl">
        <span>Phone: </span> <span class="font-bold">{{ data.phone }}</span>
      </p>
      <p class="border-b-primary-content border-b text-xl">
        <span>Post Code: </span> <span class="font-bold">{{ data.Address.post_code }}</span>
      </p>
      <p class="border-b-primary-content border-b text-xl">
        <span>City: </span> <span class="font-bold">{{ data.Address.city }}</span>
      </p>
      <p class="border-b-primary-content border-b text-xl">
        <span>Street: </span> <span class="font-bold">{{ data.Address.street }}</span>
      </p>
    </div>
  </div>
</template>
