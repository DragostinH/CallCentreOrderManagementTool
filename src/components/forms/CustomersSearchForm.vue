<script setup lang="ts">
import { reactive } from 'vue'
import { type FormData } from '@/types/types'
const emit = defineEmits(['search'])

const formData = reactive<FormData>({
  first_name: '',
  last_name: '',
  customer_number: '',
  email: '',
  order_number: null,
  phone: null,
  post_code: '',
})

function handleSearch() {
  emit('search', { ...formData })
}
</script>

<template>
  <div
    class="bg-secondary-content border border-primary rounded-lg grid grid-cols-2 gap-4 p-4 w-1/2"
  >
    <fieldset class="fieldset">
      <legend class="fieldset-legend text-secondary text-lg">First Name</legend>
      <input
        v-model="formData.first_name"
        type="text"
        placeholder="John"
        class="input"
        name="first_name"
      />
    </fieldset>
    <fieldset class="fieldset">
      <legend class="fieldset-legend text-secondary text-lg">Last Name</legend>
      <input
        v-model="formData.last_name"
        type="text"
        placeholder="McDoe"
        class="input"
        name="last_name"
      />
    </fieldset>
    <fieldset class="fieldset">
      <legend class="fieldset-legend text-secondary text-lg">Customer Number</legend>
      <input
        v-model="formData.customer_number"
        type="text"
        placeholder="11002200"
        class="input"
        name="customer_number"
      />
    </fieldset>
    <fieldset class="fieldset">
      <legend class="fieldset-legend text-secondary text-lg">Email</legend>
      <input
        v-model="formData.email"
        type="text"
        placeholder="john.mcdoe@email.com"
        class="input"
        name="email"
      />
    </fieldset>
    <fieldset class="fieldset">
      <legend class="fieldset-legend text-secondary text-lg">Post Code</legend>
      <input
        v-model="formData.post_code"
        type="text"
        placeholder="GX 7XJ"
        class="input"
        name="post_code"
      />
    </fieldset>
    <fieldset class="fieldset">
      <legend class="fieldset-legend text-secondary text-lg">Phone</legend>
      <input
        v-model="formData.phone"
        type="text"
        placeholder="00221231094"
        class="input"
        name="phone"
      />
    </fieldset>
    <fieldset class="fieldset">
      <legend class="fieldset-legend text-secondary text-lg">Order Number</legend>
      <input
        v-model="formData.order_number"
        type="text"
        placeholder="1223334444"
        class="input"
        name="order_number"
      />
    </fieldset>
    <button @click="handleSearch" class="btn btn-primary col-span-2">Search</button>
  </div>
</template>
