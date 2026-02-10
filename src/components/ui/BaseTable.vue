<script setup lang="ts">
import { computed, onBeforeMount, onMounted, ref } from 'vue'
import { useRouter } from 'vuetify/lib/composables/router.mjs'
import useApiFetch from '@/composables/useApiFetch'
import type { Customer } from '@/types/customerType'
const page = ref(1)
const router = useRouter()
const URL = computed(() => {
  return `customers?_page=${page.value}&_per_page=10`
  // return `https://jsonplaceholder.typicode.com/todos/${page.value}`
})

const props = defineProps({
  data: Object,
})

const handleClick = (customerNumber: string) => {
  if (router) router.push(`/customer/${customerNumber}`)
}

const handlePrevPage = () => {
  if (page.value === 1) return
  --page.value
}
const handleNextPage = () => {
  if (props.data?.length === 0) return
  ++page.value
}
</script>

<template>
  <div class="overflow-x-auto w-full rounded-box border border-base-content/5 bg-base-100">
    <table class="table">
      <!-- head -->
      <thead>
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
      <tbody>
        <tr
          @click="handleClick(item.customerNumber)"
          v-for="(item, i) in data"
          :key="item.id"
          class="hover:bg-neutral-400 cursor-pointer"
        >
          <th>{{ item.id }}</th>
          <td>{{ item.firstName }}</td>
          <td>{{ item.address }}</td>
          <td>{{ item.customerNumber }}</td>
          <td>{{ item.email }}</td>
          <td>{{ item.telNumber }}</td>
          <td>{{ item?.orders?.length > 0 ? item.orders[0] : 'N/A' }}</td>
        </tr>
        <!-- <tr>
          <th>1</th>
          <td>Cy Ganderton</td>
          <td>Quality Control Specialist</td>
          <td>5622351</td>
          <td>jon@jon.com</td>
          <td>089848409</td>
        </tr>
        <tr>
          <th>2</th>
          <td>Hart Hagerty</td>
          <td>Desktop Support Technician</td>
          <td>2351241</td>
          <td>jon@jon.com</td>
          <td>089848409</td>
        </tr>
        <tr>
          <th>3</th>
          <td>Brice Swyre</td>
          <td>Tax Accountant</td>
          <td>3451341</td>
          <td>jon@jon.com</td>
          <td>089848409</td>
        </tr> -->
      </tbody>
    </table>
    <div class="">
      <button @click="handlePrevPage" class="btn btn-primary uppercase">prev</button>
      <button @click="handleNextPage" class="btn btn-primary uppercase">next</button>
    </div>
  </div>
</template>
