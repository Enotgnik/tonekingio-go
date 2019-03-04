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
      <v-toolbar-title>TONE KING</v-toolbar-title>
    </v-toolbar>
    
    <v-content>
      <v-container fluid fill-height>
        <v-layout
          justify-center
          align-center
        >
          <v-flex text-xs-center>
            <v-card-title>
                My Public Projects
                <v-spacer></v-spacer>
                <v-text-field
                  append-icon="search"
                  label="Search"
                  single-line
                  hide-details
                  v-model="search"
                ></v-text-field>
              </v-card-title>
            <v-data-table
              :headers="headers"
              :items="repos"
              :search="search"
              :pagination.sync="pagination"
              :custom-filter="filterLanguage"
              disable-initial-sort
              class="elevation-1"
            >
              <template v-slot:items="props">
                <td class="text-xs-left"> {{ props.item.name }}</td>
                <td class="text-xs-left">{{ props.item.description }}</td>
                <td class="text-xs-right">{{ props.item.language }}</td>
                <td class="text-xs-right">{{ props.item.size }}</td>
                <td class="text-xs-right">{{ props.item.created_at }}</td>
                <td class="text-xs-right">{{ props.item.forks }}</td>
              </template>
              <v-alert slot="no-results" :value="true" color="error" icon="warning">
                Your search for "{{ search }}" found no results.
              </v-alert>
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

  import axios from "axios"
  export default {
   
  
    data () {
      return {
        pagination: {
          sortBy: 'forks',
          descending: true,
          rowsPerPage: 10
        },
        drawer: null,

        search: '',
        sortKey: '',
        headers: [
        {
          text: 'Title',
          align: 'left',
          value: 'name'
        },
        
          { text: 'Description', value: 'descr' },
          { text: 'Language', value: 'lang' },
          { text: 'Size (kb)', value: 'size' },
          { text: 'Created', value: 'created' },
          { text: 'Forks', value: 'forks' }
        ],
        repos: [],
       
      }
    },
    mounted() {
      var self = this
      axios.get("http://localhost:8000")
        .then(function(res){
          self.repos = res.data
          
        })

    },
    filterLanguage(){
      this.repos.language != ''
    }
  }
</script>
<style>
</style>
