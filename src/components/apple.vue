<template>
  <div>
    <h1>{{ msg }}</h1>
    <button @click="getParam">get param</button>
  </div>
</template>


<script>
  export default {
    data () {
      return {
        msg: 'i am apple'
      }
    },
    methods: {
      getParam () {
        console.log(this.$route.params)
      }
    }
  }
</script>
