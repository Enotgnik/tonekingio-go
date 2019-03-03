<template>
  <v-container>
    <v-navigation-drawer
      fixed
      v-model="drawer"
      app
    >
      <v-list dense>
        <v-list-tile>
          <v-list-tile-action>
            <v-icon>home</v-icon>
          </v-list-tile-action>
          <v-list-tile-content>
            <v-list-tile-title>Home</v-list-tile-title>
          </v-list-tile-content>
        </v-list-tile>
        <v-list-tile>
          <v-list-tile-action>
            <v-icon>contact_mail</v-icon>
          </v-list-tile-action>
          <v-list-tile-content>
            <v-list-tile-title>Contact</v-list-tile-title>
          </v-list-tile-content>
        </v-list-tile>
      </v-list>
    </v-navigation-drawer>
    <v-toolbar color="indigo" dark fixed app>
      <v-toolbar-side-icon @click.stop="drawer = !drawer"></v-toolbar-side-icon>
      <v-toolbar-title>TONEKINGIO</v-toolbar-title>
    </v-toolbar>
    <v-content>
      <v-container fluid fill-height>
        <v-layout
          justify-center
          align-center
        >
          <v-flex text-xs-center>
            <!--
            <v-tooltip left>
            
              <v-btn icon large :href="source" target="_blank" slot="activator">
                <v-icon large>code</v-icon>
              </v-btn>
              
              <span>Source</span>
            </v-tooltip> -->
            <v-data-table id="data-table"
              :headers="headers"
              :items="desserts"
              class="elevation-1"
            >
              <template v-slot:items="props">
                <td class="text-xs-left">{{ props.item.name }}</td>
                <td class="text-xs-right">{{ props.item.descr }}</td>
                <td class="text-xs-right">{{ props.item.lang }}</td>
                <td class="text-xs-right">{{ props.item.size }}</td>
                <td class="text-xs-right">{{ props.item.commit }}</td>
                <td class="text-xs-right">{{ props.item.star }}</td>
              </template>
            </v-data-table>
          </v-flex>
        </v-layout>
      </v-container>
    </v-content>
    <v-footer color="indigo" app inset>
      <span style="margin-left:5px;" class="white--text">&copy; 1776-2019 toneking.io</span>
    </v-footer>
  </v-container>
</template>

<script>
  export default {
    el: '#app',
    data: () => ({
      drawer: null,
      headers: [
        {
          text: 'Application',
          align: 'left',
          sortable: false,
          value: 'name'
        },
        
        { text: 'Description', value: 'descr' },
        { text: 'Language', value: 'lang' },
        { text: 'Size (kb)', value: 'size' },
        { text: 'Commits', value: 'commit' },
        { text: 'Stars', value: 'star' }
      ],
      desserts: [
        {
          name: 'Frozen Yogurt',
          descr: 159,
          lang: 6.0,
          size: 24,
          commit: 4.0,
          star: '1%'
        },
        {
          name: 'Ice cream sandwich',
          descr: 237,
          lang: 9.0,
          size: 37,
          commit: 4.3,
          star: '1%'
        },
        {
          name: 'Eclair',
          descr: 262,
          lang: 16.0,
          size: 23,
          commit: 6.0,
          star: '7%'
        },
        {
          name: 'Cupcake',
          descr: 305,
          lang: 3.7,
          size: 67,
          commit: 4.3,
          star: '8%'
        },
        {
          name: 'Gingerbread',
          descr: 356,
          lang: 16.0,
          size: 49,
          commit: 3.9,
          star: '16%'
        },
        {
          name: 'Jelly bean',
          descr: 375,
          lang: 0.0,
          size: 94,
          commit: 0.0,
          star: '0%'
        },
        {
          name: 'Lollipop',
          descr: 392,
          lang: 0.2,
          size: 98,
          commit: 0,
          star: '2%'
        },
        {
          name: 'Honeycomb',
          descr: 408,
          lang: 3.2,
          size: 87,
          commit: 6.5,
          star: '45%'
        },
        {
          name: 'Donut',
          descr: 452,
          lang: 25.0,
          size: 51,
          commit: 4.9,
          star: '22%'
        },
        {
          name: 'KitKat',
          descr: 518,
          lang: 26.0,
          size: 65,
          commit: 7,
          star: '6%'
        }
      ]
    
    })
  }
</script>

<style>

</style>
