<template>
  <div>
    <h1>{{ msg }}</h1>
    <button @click="addOne">add one</button>
    <button @click="minusOne">minus one</button>
  </div>
</template>


<script>
  export default {
    data () {
      return {
        msg: 'i am apple',
        price: 5
      }
    },
    methods: {
      addOne () {
        this.$store.dispatch('increase', this.price)
      },
      minusOne () {
        this.$store.commit('decrement', this.price)
      }
    }
  }
</script>
