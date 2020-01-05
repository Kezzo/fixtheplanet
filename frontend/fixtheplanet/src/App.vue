<template>
  <div id="app">
    <Issues v-bind:issues="issues" />
  </div>
</template>

<script>
import Issues from './components/Issues';

export default {
  name: 'app',
  components: {
    Issues
  },
  data() {
    return {
      issues: []
    }
  },
  created() {
    fetch("http://localhost:80/issues")
    .then((resp) => {
      resp.json().then(json => {
        //console.log(json);
        this.issues = json.Issues;
      })
      .catch(err => console.log(err));
    })
    .catch(err => console.log(err));
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
