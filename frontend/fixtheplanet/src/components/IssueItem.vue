<template>
    <div class="issue-item">
        <div class="issue-lang">
            <span class="lang-circle" v-bind:style="{ background: issue.LangColor}"></span>
            <span class="lang-text">{{issue.Language}}</span>
        </div>
        <a class="issue-title" v-bind:href="'https://github.com/'+issue.Repo+'/issues/'+issue.Number">
            <div class="issue-title-name"> <p>{{issue.Title}}</p> </div>

            <div class="issue-labels">
                <div class="issue-label-name" v-bind:style="{ background: '#'+label.Color, color: label.TextColor }" v-bind:key="label.Name" v-for="label in issue.Labels">
                    <p class="issue-label-name-text">{{label.Name}}</p>
                </div>
            </div>
        </a>
        <a class="issue-repo" v-bind:href="'https://github.com/'+issue.Repo">
            <span class="issue-repo-text">{{issue.Repo}}</span>
        </a>
    </div>
</template>

<script>
export default {
    name: "IssueItem",
    props: ["issue"],
}
</script>

<style scoped>

a {
    text-decoration: none;
    color: #2c3e50;
}

.issue-item {
    background: white;
    padding: 0px;
    border-bottom: 0px whitesmoke solid;
    border-radius: 5px;
    margin-bottom: 5px;

    display: grid;
    grid-template-columns: 15% auto 25%;
    grid-template-rows: 100%;

    grid-column-gap: 0px;
}

.issue-title:hover {
    background: darkseagreen;
}

.issue-repo:hover {
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

    justify-self: stretch;
    align-self: stretch;

    text-align: left;
}

.issue-labels{
    display: flex;

    flex-direction: row;
    flex-wrap: wrap;

    margin-top: -10px;
    margin-left: 5px;
}

.issue-title-name{
    margin-top: -5px;
    margin-left: 5px;
}

.issue-label-name{
    border-radius: 2px;
    padding-top: 0.15em;
    padding-right: 4px;
    padding-bottom: 0.15em;
    padding-left: 4px;

    margin-right: 5px;

    margin-top: 5px;
    margin-bottom: 5px;
}

.issue-label-name-text{
    margin: 0;
}

.issue-repo{
    display: grid;
    
    grid-column-start: 3;
    grid-column-end: 4;

    justify-self: stretch;
    align-self: stretch;
    text-align: center;
    text-anchor: middle;
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
        grid-template-columns: 20% auto;
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
        grid-template-columns: 150px auto;
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

.issue-repo-text{
    font-style: italic;
    align-self: center;
}

</style>