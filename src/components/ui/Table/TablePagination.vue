<script setup lang="ts">
import { ref } from 'vue'

defineProps<{
  isFetching: boolean
}>()

const page = ref(1)

const handlePrevPage = () => {
  --page.value
}
const handleNextPage = () => {
  ++page.value
}
</script>

<template>
  <div v-show="!isFetching" class="">
    <button @click="handlePrevPage" class="btn btn-primary uppercase">prev</button>
    <button @click="handleNextPage" class="btn btn-primary uppercase">next</button>
  </div>
</template>
