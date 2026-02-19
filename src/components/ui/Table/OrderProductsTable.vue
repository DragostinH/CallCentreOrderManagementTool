<script setup lang="ts">
import router from '@/router'
import { type OrderItem } from '@/types/types'
import { ref } from 'vue'
import { useRoute } from 'vue-router'

const route = useRoute()
const customerNumber = route.params.customer_number
const orderNumber = route.params.order_number
const selectedRows = ref<number[]>([])
const selectAll = ref<boolean>(false)

const props = defineProps<{
  orderItems: OrderItem[] | null
  isFetching: boolean
}>()

const handleClick = (product_id: number) => {
  router.push(`product/${product_id}`)
}

const isSelected = (product_id: number) => {
  return selectedRows.value.includes(product_id) ? 'bg-secondary-content text-neutral-content' : ''
}

const handleSelectAll = () => {
  if (!selectAll.value) {
    if (!props.orderItems) {
      return
    }
    if (props.orderItems?.length > 0) {
      const allProductIds = props.orderItems?.map((item) => {
        return item.product_id
      })

      selectedRows.value = [...allProductIds]
    }
    console.log(selectedRows.value)
  } else {
    deselectAll()
  }
}

const deselectAll = () => {
  selectedRows.value = []
  console.log(selectedRows.value)
}
const handleSelectRow = () => {}
</script>

<template>
  <table class="table table-md table-fixed table-zebra bg-base-100">
    <!-- head -->
    <thead class="z-10 bg-secondary">
      <tr class="text-center">
        <!-- product image? -->
        <th>
          <input
            @click="handleSelectAll"
            v-model="selectAll"
            :value="selectAll"
            type="checkbox"
            className="checkbox checkbox-accent"
          />
        </th>
        <th>Name</th>
        <th>Type</th>
        <th>Unit Price</th>
        <th>Retail Price</th>
        <th>Category</th>
        <th>Quantity</th>
        <th>Total Amount</th>
      </tr>
    </thead>

    <tbody v-if="!isFetching">
      <tr
        @click="handleClick(item?.product_id)"
        v-for="item in orderItems"
        :key="item?.product_id"
        class="hover:bg-neutral-400 cursor-pointer font-semibold text-center"
        :class="isSelected(item.product_id)"
      >
        <td>
          <input
            @click.stop="handleSelectRow"
            v-model="selectedRows"
            :value="item.product_id"
            type="checkbox"
            className="checkbox checkbox-secondary"
          />
        </td>
        <td>{{ item.product.name }}</td>
        <td>{{ item.product.product_type }}</td>
        <td>{{ item.product.UnitPrice.price }}</td>
        <td>{{ item.product.RetailPrice.price }}</td>
        <td>{{ item.product.categories[0]?.name }}</td>
        <td>{{ item.quantity }}</td>
        <td class="text-center">{{ item.price }}</td>
      </tr>
    </tbody>
    <tbody v-else>
      <tr>
        <td colspan="7" className="bg-gray-100 p-2 text-center">Loading</td>
      </tr>
    </tbody>
  </table>
</template>
