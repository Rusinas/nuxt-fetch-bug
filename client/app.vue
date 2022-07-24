<template>
  <div>
    <h1>Hello. This is stupidest thing I ever faced in my life</h1>
  </div>
</template>


<script lang="ts">
import { defineComponent } from 'vue'

export default defineComponent({
  async setup() {
    const data = await $fetch('http://localhost:4000/')
  },
})
</script>
