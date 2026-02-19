<script setup lang="ts">
import OrderProductsTable from '@/components/ui/Table/OrderProductsTable.vue'
import useApiFetch from '@/composables/useApiFetch'
import { type Order } from '@/types/types'
import { useRoute } from 'vue-router'

const orderNumber = useRoute().params.order_number
const { data, error, isFetching } = useApiFetch(`order/${orderNumber}`).get().json<Order>()
</script>

<template>
  <div class="">
    <div class="">
      <OrderProductsTable v-if="data" :order-items="data?.items" :is-fetching="isFetching" />
      <div v-if="error" class="">
        <p>Something happened. Error: {{ error }}</p>
      </div>
    </div>
  </div>
</template>
