<template>
    <section style="margin-left:10%">
        <v-layout
          column
          wrap
          class="my-5"
          align-center
          
        >
          <v-flex xs12 sm4 class="my-3" >
            <div  class="text-xs-left">
              <h1 style="letter-spacing:10px !important;" class="headline">MY WORK</h1>
              <p style="width:70%; font-weight:bold;">I spend a lot of time tinkering with electronics and building applications. This site
                  exists to share ideas and theoris that interest me. I've been working on a number of projects
                  that I just seem to loose motivation to complete; however, that is helping me pin point the type of
                  applications I want to build.</p>
            </div>
            
          </v-flex>
          
          <!--
          <v-flex xs12>
            <v-container grid-list-xl>
              <v-layout row wrap align-center>
                <v-flex xs12 md4>
                  <v-card class="elevation-0 transparent">
                   
                    
                      <div class="headline text-xs-left">Application development</div>
                  
                    <v-card-text>
                      <ul>
                        <li>
                            AI/Machine Learning    
                        </li>
                         <li>
                            Data Migraton
                        </li>
                        <li>
                            ITSM Bussiness Automation  
                        </li>
                        <li>
                            Web Applications  
                        </li>
                      </ul> 
                    </v-card-text>
                  </v-card>
                </v-flex>
                <v-flex xs12 md4>
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
                <v-flex xs12 md4>
                  <v-card class="elevation-0 transparent">
                    <v-card-title primary-title class="layout justify-left">
                      <div  class="headline">Community</div>
                    </v-card-title>
                    <v-card-text>
                      <ul >
                        <li >
                            Discord    
                        </li>
                         <li>
                            YouTube (eventually)
                        </li>
                         <li>
                            Open source projects
                        </li>
                        <li>
                            Meetups
                        </li>
                      </ul> 
                    </v-card-text>
                  </v-card>
                </v-flex>

              </v-layout>
            </v-container>
          </v-flex> -->
        </v-layout>
      </section>
</template>
<style>
    h1{
        font-size: 1.5em; /* 25px/16px */
    }
    #content       
    {
        font-size: .75em; /* 12/16 */
        width: 71.875%; /* 650/960 */
        max-width: 960px;
        text-align: left;
        
        padding: 3.84615384615385%; /* 25/650 */
    }
    /*CSS3 multiple columns.*/
    /* Get em size for colums: 184/16 */   
    .columns
    {   
        -moz-column-width: 11.5em; /* Firefox */
        -webkit-column-width: 11.5em; /* webkit, Safari, Chrome */
        column-width: 11.5em;
    }
    /*remove standard list and bullet formatting from ul*/
    .columns ul
    {
        margin: 0;
        padding: 0;
        list-style-type: none;
    }
    /* correct webkit/chrome uneven margin on the first column*/
    .columns ul li:first-child
    {
        margin-top:0px;
    }
</style>
