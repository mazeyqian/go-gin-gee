<template>
  <div>
    <h1>{{ msg }}</h1>

    <router-view></router-view>
  </div>
</template>


<script>
  export default {
    data () {
      return {
        msg: 'i am red apple'
      }
    },
    methods: {
      getParam () {
        console.log(this.$route.params)
      }
    }
  }
</script>
