<template>
    <div class="top-bar-container">
        <a class="github-button" v-bind:href="'https://github.com/Kezzo/fixtheplanet'">
            <svg version=1.1 width=20 height=20 viewBox="0 0 16 16" class="octocat" aria-hidden=true><path fill-rule=evenodd d="M8 0C3.58 0 0 3.58 0 8c0 3.54 2.29 6.53 5.47 7.59.4.07.55-.17.55-.38 0-.19-.01-.82-.01-1.49-2.01.37-2.53-.49-2.69-.94-.09-.23-.48-.94-.82-1.13-.28-.15-.68-.52-.01-.53.63-.01 1.08.58 1.23.82.72 1.21 1.87.87 2.33.66.07-.52.28-.87.51-1.07-1.78-.2-3.64-.89-3.64-3.95 0-.87.31-1.59.82-2.15-.08-.2-.36-1.02.08-2.12 0 0 .67-.21 2.2.82.64-.18 1.32-.27 2-.27.68 0 1.36.09 2 .27 1.53-1.04 2.2-.82 2.2-.82.44 1.1.16 1.92.08 2.12.51.56.82 1.27.82 2.15 0 3.07-1.87 3.75-3.65 3.95.29.25.54.73.54 1.48 0 1.07-.01 1.93-.01 2.2 0 .21.15.46.55.38A8.013 8.013 0 0 0 16 8c0-4.42-3.58-8-8-8z"/></svg>
            <p class="repo-text">Repository</p>
        </a>

        <div class="website-name">
            <h1 class="website-name-name">fixthepla.net</h1>
            <h4 class="website-name-desc">Helps you contribute to open source projects.</h4>
        </div>
    </div>
</template>

<style scoped>

a {
    text-decoration: none;
}

.top-bar-container {
    padding-top: 10px;
    display: grid;
    grid-template-columns: 10% auto auto 10%;
    grid-template-rows: 100%;
}

.github-button {
    display: grid;
    grid-template-columns: 16px auto;
    grid-template-rows: auto;

    align-self: right;
    justify-self: right;

    background: #333;
    color: #f2f2f2;
    border: none;
    text-align: center;
    text-decoration: none;
    border-radius: 5px;
    box-sizing: border-box;

    width: 130px;
    height: 42px;
    padding-top: 10px;
    padding-bottom: 10px;

    cursor: pointer;
    float: right;

    grid-row-start: 1;
    grid-row-end: 2;

    grid-column-start: 3;
    grid-column-end: 4;

    margin-top: 20px;
    margin-right: 10px;
}

.website-name{
    grid-column-start: 2;
    grid-column-end: 3;

    text-align: left;

    grid-row-start: 1;
    grid-row-end: 2;
    
    margin-bottom: 10px;
    margin-left: 10px;
}

.website-name-name{
    margin: 0px;
}

.website-name-desc{
    margin: 0px;
}

@media screen and (max-width: 930px) {
    
    .top-bar-container {
        grid-template-columns: auto auto;
    }

    .github-button {
        margin-right: 10px;
    }

    .website-name{
        margin-left: 10px;

        grid-column-start: 1;
        grid-column-end: 2;
    }
}

.repo-text {
    margin: 0px;
    margin-left: 10px;
}

.octocat {
    margin-left: 13px;
    margin-top: 0px;

    fill: white;
}

</style>