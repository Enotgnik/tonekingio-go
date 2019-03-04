<template>
  <v-app>
    <v-content>
      <MainLayout/>
    </v-content> 
  </v-app>
</template>

<script>
import MainLayout from './components/MainLayout'

//import axios from "axios"

export default {
  name: 'App',
  components: {
    MainLayout,
  }
}
</script>
