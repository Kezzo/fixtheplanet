<template>
  <div id="app">
    <TopBar />
    <SearchBar v-bind:selectedLanguages="selectedLanguages" v-on:search-issues="searchIssues" v-on:select-language="selectLanguage" />
    <Issues v-bind:issues="issues" />
  </div>
</template>

<script>
import TopBar from './components/TopBar';
import SearchBar from './components/SearchBar';
import Issues from './components/Issues';

import langColors from './assets/lang-colors.json';
import mainLangs from './assets/main-langs.json';

import _ from 'underscore';

export default {
  name: 'app',
  components: {
    TopBar,
    Issues,
    SearchBar,
  },
  methods: {
    getBaseUrl() {
      return process.env.VUE_APP_BASE_URL;
    },
    searchIssues(reset) {
      if(reset) {
        this.pagingSeed = 0;
        this.pagingOffset = 0;
      }

      let urlParams = "";

      this.selectedLanguages.forEach(lang => {
        if(lang.IsSelected) {
          urlParams += urlParams == "" ? "?" : "&";
          urlParams += "language=" + encodeURIComponent(lang.Language);
        }
      });

      urlParams += urlParams == "" ? "?seed=" : "&seed=";
      urlParams += this.pagingSeed;
      urlParams += "&offset=" + this.pagingOffset;

      fetch(this.getBaseUrl() + "/issues" + urlParams)
      .then((resp) => {
        resp.json().then(json => {

          if(reset) {
            this.issues = [];
          }

          this.pagingSeed = json.PagingSeed;
          this.pagingOffset = json.NextOffset;
          //console.log(json);
          //this.issues = [];

          json.Issues.forEach(jsonIssue => {
            const exists = _.some(this.issues, issue => { return _.isEqual(issue, jsonIssue) })

            if(!exists) {
              this.issues.push(jsonIssue);
            }
          });

          //console.log(JSON.stringify(this.languageColors));

          this.issues.forEach(issue => {
            issue.LangColor = issue.Language in this.languageColors ? this.languageColors[issue.Language].color : "#ffffff";
            issue.Labels.forEach(label => {
              var colorInDec = parseInt(label.Color.charAt(0), 16);
              label.TextColor = colorInDec > 8 ? "#000000" : "#ffffff";
            });
          });

          this.queryingNextPage = false;
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

      this.searchIssues(true);
    },
    onScroll() {
      var scrolledToBottom = (document.documentElement.scrollTop + window.innerHeight) >= (document.documentElement.offsetHeight - 100);
      
      if(scrolledToBottom && !this.queryingNextPage)
      {
        this.queryingNextPage = true;
        this.searchIssues(false);
      }
    },
    resetList() {
      this.issues = [];
      this.pagingSeed = 0;
      this.pagingOffset = 0;
    }
  },
  
  data() {
    return {
      issues: [],
      languageColors: langColors,
      mainLanguages: mainLangs,
      selectedLanguages: [],
      pagingSeed: 0,
      pagingOffset: 0,
      queryingNextPage: false
    }
  },
  created() {
    this.searchIssues();

    this.mainLanguages.forEach(lang => {
      let color = lang in this.languageColors ? this.languageColors[lang].color : "#ffffff";
      let colorInDec = parseInt(color.charAt(1), 16);
      let textColor = colorInDec > 8 ? "#000000" : "#ffffff";

      this.selectedLanguages.push({
        Language: lang,
        Color: color,
        TextColor: textColor,
        IsSelected: false
      })
    });

    window.onscroll = this.onScroll;
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
  background:whitesmoke;
}

body {
  margin: 0px;
}
</style>
