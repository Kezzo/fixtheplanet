<template>
    <div class="issue-item" @click="openRepo()">
        <div class="issue-lang">
            <span class="lang-circle" v-bind:style="{ background: issue.LangColor}"></span>
            <span class="lang-text">{{issue.Language}}</span>
        </div>
        <div class="issue-title"> <p>{{issue.Title}}</p> </div>
        <div class="issue-repo"> <p>{{issue.Repo}}</p> </div>
    </div>
</template>

<script>
export default {
    name: "IssueItem",
    props: ["issue"],
    methods: {
        openRepo() {
            var win = window.open("https://github.com/" + this.issue.Repo + "/issues/" + this.issue.Number);
            win.focus();
        }
    }
}
</script>

<style scoped>
    .issue-item {
        background: #f4f4f4;
        padding: 3px;
        border-bottom: 1px #ccc solid;

        display: grid;
        grid-template-columns: 15% auto 25%;
        grid-template-rows: 100%;

        grid-column-gap: 20px;

        cursor: pointer;
    }

    .issue-item:hover {
        background: darkseagreen;
    }

    .issue-title{
        grid-column-start: 2;
        grid-column-end: 3;

        align-self: stretch;

        text-align: left;
        
    }

    .issue-repo{
        grid-column-start: 3;
        grid-column-end: 4;

        align-self: center;
        text-align: center;
    }

    .issue-lang{
        grid-column-start: 1;
        grid-column-end: 2;

        align-self: center;
        text-align: left;
        padding-left: 15px;
    }

    .lang-circle {
        display: inline-block;
        width: 15px;
        height: 15px;
        position: relative;
        top: 2px;
        border-radius: 50%;
    }

    .lang-text {
        padding-left: 5px;
        font-size: 15px;
    }
</style>