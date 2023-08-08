<template>
  <div>
    <h1>Hello. This is stupidest thing I ever faced in my life</h1>
  </div>
</template>

<script lang="ts" setup>
const data = await $fetch("http://localhost:4000/", { method: "GET", mode: "cors" })
console.log({ data })
</script>
