<template>
    <div class="issue-item" @click="openRepo()">
        <div class="issue-lang">
            <span class="lang-circle" v-bind:style="{ background: issue.LangColor}"></span>
            <span class="lang-text">{{issue.Language}}</span>
        </div>
        <div class="issue-title">
            <div class="issue-title-name"> <p>{{issue.Title}}</p> </div>

            <div class="issue-labels">
                <div class="issue-label-name" v-bind:style="{ background: '#'+label.Color, color: label.TextColor }" v-bind:key="label.Name" v-for="label in issue.Labels">
                    <p class="issue-label-name-text">{{label.Name}}</p>
                </div>
            </div>

        </div>
        <div class="issue-repo"> <span class="issue-repo-span">{{issue.Repo}}</span> </div>
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

@media (hover: none) {
    .issue-item:hover {
        background: #f4f4f4;
    }
}

    .issue-title{
        grid-column-start: 2;
        grid-column-end: 3;

        text-align: left;
    }

    .issue-labels{
        display: flex;
    
        flex-direction: row;
        flex-wrap: wrap;
    }

    .issue-title-name{
        margin-top: -5px;
    }

    .issue-label-name{
        border-radius: 2px;
        padding-top: 0.15em;
        padding-right: 4px;
        padding-bottom: 0.15em;
        padding-left: 4px;

        margin-right: 5px;

        margin-top: -5px;
        margin-bottom: 5px;
    }

    .issue-label-name-text{
        margin: 0;
    }

    .issue-repo{
        display: grid;

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
        margin-left: 15px;

        display: grid;
        grid-template-columns: 15px auto;

        width: 100px;
    }

@media screen and (max-width: 930px) {
    .issue-item{
        grid-template-columns: 15% auto;
        grid-template-rows: 30% auto;
    }   

    .issue-title{
        grid-row-start: 2;
        grid-row-end: 3;
        grid-column-start: 2;
        grid-column-end: 4;
    }

    .issue-repo{
        grid-row-start: 1;
        grid-row-end: 2;

        grid-column-start: 2;
        grid-column-end: 3;

        text-align: left;

        margin-top: 15px;
    }

    .issue-lang{
        grid-row-start: 1;
        grid-row-end: 3;
    }
}

@media screen and (max-width: 755px) {
    .issue-item {
        grid-template-columns: 100px auto;
        grid-template-rows: 30% auto;
    }

    .issue-lang{
        grid-row-start: 1;
        grid-row-end: 2;

        grid-column-start: 1;
        grid-column-end: 2;

        margin-top: 15px;
    }

    .issue-title{
        grid-column-start: 1;
        grid-column-end: 4;

        margin-left: 15px;
    }
}

    .lang-circle {
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

    .issue-repo-span{
        font-style: italic;
    }

</style>