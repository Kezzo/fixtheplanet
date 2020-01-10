<template>
  <div id="app">
    <SearchBar v-bind:selectedLanguages="selectedLanguages" v-on:search-issues="searchIssues" v-on:select-language="selectLanguage" />
    <Issues v-bind:issues="issues" />
  </div>
</template>

<script>
import SearchBar from './components/SearchBar'
import Issues from './components/Issues';

import langColors from './assets/lang-colors.json'
import mainLangs from './assets/main-langs.json'

import _ from 'underscore';

export default {
  name: 'app',
  components: {
    Issues,
    SearchBar,
  },
  methods: {
    getBaseUrl() {
      return process.env.VUE_APP_BASE_URL;
    },
    searchIssues() {
      let langFilter = "";

      this.selectedLanguages.forEach(lang => {
        if(lang.IsSelected) {
          langFilter += langFilter == "" ? "?" : "&";
          langFilter += "language=" + lang.Language;
        }
      });

      fetch(this.getBaseUrl() + "/issues" + langFilter)
      .then((resp) => {
        resp.json().then(json => {
          //console.log(json);
          this.issues = [];

          json.Issues.forEach(jsonIssue => {
            const exists = _.some(this.issues, issue => { return _.isEqual(issue, jsonIssue) })

            if(!exists) {
              this.issues.push(jsonIssue);
            }
          });

          //console.log(JSON.stringify(this.languageColors));

          this.issues.forEach(issue => {
            issue.LangColor = issue.Language in this.languageColors ? this.languageColors[issue.Language].color : "#ffffff";
          });
        })
        .catch(err => console.log(err));
      })
      .catch(err => console.log(err));
    },
    selectLanguage(selectedLang) {
      this.selectedLanguages.forEach(lang => {
        if(lang.Language == selectedLang.Language){
          lang.IsSelected = !lang.IsSelected;
        }
      });

      this.searchIssues();
    }
  },
  data() {
    return {
      issues: [],
      languageColors: langColors,
      mainLanguages: mainLangs,
      selectedLanguages: [],
    }
  },
  created() {
    this.searchIssues();

    this.mainLanguages.forEach(lang => {
      this.selectedLanguages.push({
        Language: lang,
        IsSelected: false
      })
    });
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
}

@media screen and (min-width: 931px) {
  #app {
    margin-top: 60px;
  }
}
</style>
