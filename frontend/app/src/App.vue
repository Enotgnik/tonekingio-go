<template>
  <v-app>
    <v-content>
      <section style="margin-bottom:-5%;">
        <v-parallax :src="require('./assets/hero.jpeg')" height="600">
          <v-layout
            column
            align-left
            justify-center
            class="white--text"
          >
            <div style="margin-left:25%">
              <h1 style="letter-spacing:10px !important;" class="white--text mb-1 display-3">Antone King</h1>
              <h2 class=" mb-3">Not your traditional developer. Daily self-learning through research and development.</h2>
            </div>
          </v-layout>
        </v-parallax>
      </section>
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


