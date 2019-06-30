<template>
  <v-container>
    
    <v-toolbar color="blue-grey darken-3" dark fixed app>
      <!--<v-toolbar-side-icon ></v-toolbar-side-icon>-->
      <v-toolbar-title>Antone King</v-toolbar-title>
    </v-toolbar>
    
    <v-content>
 
        <!-- add body here -->
         <v-container grid-list-xl style="">
          <h1 style="letter-spacing:10px !important; margin-bottom:25px;" class="headline">MY WORK</h1>
        <div  class="text-xs-left" style="margin-bottom: 50px;">
              
              
              <p style="width:70%; font-weight:bold;">I spend a lot of time tinkering with electronics and building applications. This site
                  exists to share ideas and theories that interest me. I've been working on a number of projects
                  that I just seem to lose motivation to complete; however, that is helping me pin point the type of
                  applications I want to build.</p>
            </div>
              <v-layout row wrap >
                <v-flex xs12 md3>
                  <v-card class="elevation-0 transparent">
                   
                    
                      <div class="headline text-xs-left">Application Development</div>
                     
                  
                    <v-card-text>
                      <ul>
                        <li>
                            ServiceNow Custom Apps   
                        </li>
                         <li>
                            Application Interfaces
                        </li>
                        <li>
                            ITSM Business Automation  
                        </li>
                        <li>
                            Data Migration  
                        </li>
                      </ul> 
                    </v-card-text>
                  </v-card>
                </v-flex>
                <v-flex xs12 md3>
                  <v-card class="elevation-0 transparent">
                  
                      <div class="headline text-xs-left">Hardware</div>
                       
                    
                    <v-card-text>
                      <ul>
                        <li>
                            Water Flow and Pressure reading with arduino    
                        </li>
                         <li>
                            Serving arduino data with rPi
                        </li>
                        <li>
                            Leak detection with Machine Learning  
                        </li>
                        <li>
                            Home Automation
                        </li>
                      </ul> 
                    </v-card-text>
                  </v-card>
                </v-flex>
                 <v-flex xs12 md3>
                   <div  class="headline">Certifications</div>
                  <v-card class="elevation-0 transparent">
                    <v-card-text>
                      <ul >
                        <li >
                            ServiceNow CSA
                        </li>
                         <li>
                            ServiceNow CAD
                        </li>
                         <li>
                            ServiceNow CAS - Performance Analytics (In Progress)
                        </li>
                        <li>
                            AWS Cloud Practionier (In Progress)
                        </li>
                      </ul> 
                    </v-card-text>
                  </v-card>
                </v-flex>
                <v-flex xs12 md3>
                  <div  class="headline">Future Endeavors (2019)</div>
                  <v-card class="elevation-0 transparent">
                    <v-card-text>
                      <ul >
                        <li >
                            Machine Learning/AI  
                        </li>
                         <li>
                            Big Data
                        </li>
                         <li>
                            Blockchain
                        </li>
                        <li>
                            FPGA Devices
                        </li>
                        <li>
                            Distributed Systems
                        </li>
                      </ul> 
                    </v-card-text>
                  </v-card>
                </v-flex>

              </v-layout>
            </v-container>

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
      axios.get("https://api.github.com/users/enotgnik/repos")
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
