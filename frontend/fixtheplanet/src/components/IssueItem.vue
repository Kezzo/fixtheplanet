<template>
    <div class="issue-item" @click="openRepo()">
        <div class="issue-title"> <h3>{{issue.Title}}</h3> </div>
        <div class="issue-repo"> <p>{{issue.Repo}}</p> </div>
        <div class="issue-lang"> <p>{{issue.Language}}</p> </div>
    </div>
</template>

<script>
export default {
    name: "IssueItem",
    props: ["issue"],
    methods: {
        openRepo() {
            var win = window.open("https://github.com/" + this.issue.Repo);
            win.focus();
        }
    }
}
</script>

<style scoped>
    .issue-item {
        background: #f4f4f4;
        padding: 3px;
        border-bottom: 1px #ccc dotted;

        display: grid;
        grid-template-columns: auto 20% 15%;
        grid-template-rows: 100%;

        grid-column-gap: 10px;

        cursor: pointer;
    }

    .issue-item:hover {
        background: darkseagreen;
    }

    .issue-title{
        grid-column-start: 1;
        grid-column-end: 2;

        align-self: stretch;

        text-align: left;
        padding-left: 15px;
    }

    .issue-repo{
        grid-column-start: 2;
        grid-column-end: 3;

        align-self: center;
    }

    .issue-lang{
        grid-column-start: 3;
        grid-column-end: 4;

        align-self: center;
    }
</style>