<template>
    <div>
        <h1>Issues:</h1>
        <div class="issues-container" v-bind:key="issue.Title" v-for="issue in issues">
            <div class="left-whitespace"> </div>
            <div class="issue-inner-container"> 
                <IssueItem v-bind:issue="issue" />
            </div>
            <div class="right-whitespace"> </div>
        </div>
    </div>
</template>

<script>
import IssueItem from './IssueItem';

export default {
    name: "Issues",
    components: {
        IssueItem
    },
    props: ["issues"]
}
</script>

<style scoped>

.issues-container{
    display: grid;
    grid-template-columns: 10% auto 10%;
    grid-template-rows: 100%;
}

.left-whitespace{
    grid-column-start: 1;
    grid-column-end: 2;
}

.issue-inner-container{
    grid-column-start: 2;
    grid-column-end: 3;
}

.right-whitespace{
    grid-column-start: 3;
    grid-column-end: 4;
}

</style>