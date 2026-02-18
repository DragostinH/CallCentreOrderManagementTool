<script setup lang="ts">
import useApiFetch from '@/composables/useApiFetch'
import { useRoute } from 'vue-router'

const orderNumber = useRoute().params.order_number
const customerNumber = useRoute().params.customer_number
const { data, error, isFetching } = useApiFetch(`order/${orderNumber}`).get().json()
</script>

<template>
  <div class="">
    <p>Order PAge</p>
    <p>some table</p>
  </div>
</template>
