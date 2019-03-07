<template>
  <v-container>
    
    <v-toolbar color="blue-grey darken-3" dark fixed app>
      <!--<v-toolbar-side-icon ></v-toolbar-side-icon>-->
      <v-toolbar-title>Antone King</v-toolbar-title>
    </v-toolbar>
    
    <v-content>
 
        <!-- add body here -->
        <v-layout
          justify-center
          align-center
        >
          <v-flex text-xs-center>
            <v-card-title class="headline">
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
              disable-initial-sort
              class="elevation-1"
            >
              <template v-slot:items="props">
                <td class="text-xs-left"> <a target="_blank" v-bind:href="props.item.html_url">{{ props.item.name }}</a></td>
                <td class="text-xs-left">{{ props.item.description }}</td>
                <td class="text-xs-left" v-if="props.item.language">{{ props.item.language }}</td>
                <td class="text-xs-left" v-else>ServiceNow Scoped App</td>
                <td class="text-xs-left">{{ props.item.size }}</td>
                <td class="text-xs-left">{{ props.item.created_at }}</td>
                <td class="text-xs-left">{{ props.item.forks }}</td>
              </template>
              <v-alert slot="no-results" :value="true" color="error" icon="warning">
                Your search for "{{ search }}" found no results.
              </v-alert>
            </v-data-table>
          </v-flex>
        </v-layout>
    </v-content>
    <v-footer color="blue-grey darken-3" app inset>
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
      axios.get("http://localhost:8080/api")
        .then(function(res){
          self.repos = res.data  
        })
        .catch(function(err){
          return err.error
        })

    },
    
  }
</script>
<style>
a {
  text-decoration: none;
  
}

</style>
