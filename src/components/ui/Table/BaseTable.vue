<script setup lang="ts">
import { useRouter } from 'vue-router'
import { type Customer } from '@/types/customerType'
const router = useRouter()

defineProps<{
  isFetching: boolean
  data: Customer[]
}>()

const fullName = (item: Customer) => item.first_name + ' ' + item.last_name
const handleClick = (customerNumber: string) => {
  if (router) router.push(`/customer/${customerNumber}`)
}

const lastOrder = (item: Customer) => {
  return item?.orders[0]?.order_id
}
</script>

<template>
  <div class="overflow-x-auto h-80 w-full rounded-box border border-base-content/5 bg-base-100">
    <table class="table">
      <!-- head -->
      <thead class="sticky top-0 z-10 bg-white">
        <tr>
          <th></th>
          <th>Name</th>
          <th>Address</th>
          <th>Customer Number</th>
          <th>Email</th>
          <th>Tel. Number</th>
          <th>Last Order</th>
        </tr>
      </thead>
      <tbody v-if="!isFetching">
        <tr
          @click="handleClick(item.customer_number)"
          v-for="item in data"
          :key="item.ID"
          class="hover:bg-neutral-400 cursor-pointer"
        >
          <th>{{ item.ID }}</th>
          <td>{{ fullName(item) }}</td>
          <td>{{ item.Address.street }}</td>
          <td>{{ item.customer_number }}</td>
          <td>{{ item.email }}</td>
          <td>{{ item.phone }}</td>
          <td>{{ lastOrder(item) }}</td>
        </tr>
      </tbody>
      <tbody v-else>
        <tr>
          <td colspan="7" className="bg-gray-100 p-2 text-center">Loading</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>
