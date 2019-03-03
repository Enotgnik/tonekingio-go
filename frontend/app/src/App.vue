<template>
  <v-app>
    <v-content>
      <MainLayout/>
    </v-content> 
  </v-app>
</template>

<script>
import MainLayout from './components/MainLayout'

export default {
  name: 'App',
  components: {
    MainLayout
  },
  data () {
    return {
      //
    }
  }
}
</script>
