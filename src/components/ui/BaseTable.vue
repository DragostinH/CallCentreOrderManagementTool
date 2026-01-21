<script setup lang="ts">
const props = defineProps({
  data: Object,
})
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
        </tr>
      </thead>
      <tbody>
        <tr>
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
        </tr>
      </tbody>
    </table>
  </div>
</template>
