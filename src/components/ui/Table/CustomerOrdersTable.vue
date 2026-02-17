<script setup lang="ts">
import type { Order } from '@/types/customerType'

defineProps<{
  orders: Order[]
  isFetching: boolean
}>()

const handleClick = (order_id: number) => {
  console.log('order ', order_id)

  return
}
</script>

<template>
  <table class="table table-zebra bg-base-100">
    <!-- head -->
    <thead class="sticky top-0 z-10 bg-secondary">
      <tr>
        <th>Date</th>
        <th>Order ID</th>
        <th>Status</th>
        <th>Amount</th>
      </tr>
    </thead>

    <tbody v-if="!isFetching">
      <tr
        @click="handleClick(order.order_id)"
        v-for="order in orders"
        :key="order.order_id"
        class="hover:bg-neutral-400 cursor-pointer font-semibold"
      >
        <td>{{ new Date(order.order_date).toLocaleDateString() }}</td>
        <td>{{ order.order_id }}</td>
        <td>{{ order.status }}</td>
        <td>{{ order.total.toFixed(2)}}</td>
      </tr>
    </tbody>
    <tbody v-else>
      <tr>
        <td colspan="7" className="bg-gray-100 p-2 text-center">Loading</td>
      </tr>
    </tbody>
  </table>
</template>
