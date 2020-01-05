<template>
  <div id="app">
    <SearchBar v-bind:mainLanguages="mainLanguages" v-on:search-issues="searchIssues" />
    <Issues v-bind:issues="issues" />
  </div>
</template>

<script>
import SearchBar from './components/SearchBar'
import Issues from './components/Issues';

import langColors from './assets/lang-colors.json'
import mainLangs from './assets/main-langs.json'

export default {
  name: 'app',
  components: {
    Issues,
    SearchBar,
  },
  methods: {
    searchIssues() {
      fetch("http://localhost:80/issues")
      .then((resp) => {
        resp.json().then(json => {
          //console.log(json);
          this.issues = json.Issues;

          //console.log(JSON.stringify(this.languageColors));

          this.issues.forEach(issue => {
            issue.LangColor = this.languageColors[issue.Language].color;
          });
        })
        .catch(err => console.log(err));
      })
      .catch(err => console.log(err));
    }
  },
  data() {
    return {
      issues: [],
      languageColors: langColors,
      mainLanguages: mainLangs,
    }
  },
  created() {
    this.searchIssues();
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
