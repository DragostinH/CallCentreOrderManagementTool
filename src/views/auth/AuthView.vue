<script setup lang="ts">
import { useRouter } from 'vuetify/lib/composables/router'

const router = useRouter()
const login = () => {
  router.push('/')
}
</script>

<template>
  <div class="hero min-h-screen bg-neutral">
    <div class="hero-content bg-base-200 border-base-300 rounded-box border flex-col">
      <fieldset class="fieldset">
        <legend class="fieldset-legend">Name</legend>
        <input type="text" placeholder="Name" class="input" />
      </fieldset>
      <fieldset class="fieldset">
        <legend class="fieldset-legend">Password</legend>
        <input type="text" placeholder="Password" class="input" />
      </fieldset>
      <div class="flex gap-4">
        <button @click="login" class="btn btn-primary">Login</button>
        <button class="btn btn-soft btn-secondary">Help</button>
      </div>
    </div>
  </div>
</template>
