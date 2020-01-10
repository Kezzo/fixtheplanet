<template>
    <div class="search-container">
        <div class="left-whitespace"></div>
        <div class="searchbar-inner-container">
            <div>
                <div class="lang-selection">
                    <div class="lang-filter" @click="select(lang)" v-bind:class="{'selected-lang-filter':lang.IsSelected }" v-bind:key="lang.Language" v-for="lang in selectedLanguages">
                        <div class="lang-filter-text">{{lang.Language}}</div>
                    </div>
                </div>
            </div>
            <div class="search-button-parent" >
                <div class="search-button" @click="searchIssues()">Search</div>
            </div>
        </div>
        <div class="right-whitespace"></div>
    </div>
</template>

<script>
export default {
    props: ["selectedLanguages"],
    methods: {
        searchIssues() {
            this.$emit('search-issues');
        },
        select(lang) {
            this.$emit('select-language', lang);
        }
    }
}
</script>

<style scoped>

.search-container{
    display: grid;
    grid-template-columns: 10% auto 10%;
    grid-template-rows: 100%;

    padding-bottom: 15px;
}

.left-whitespace{
    grid-column-start: 1;
    grid-column-end: 2;
}

.searchbar-inner-container{
    grid-column-start: 2;
    grid-column-end: 3;

    padding: 10px;

    background: #f4f4f4;

    display: grid;
    grid-template-columns: auto, 100px;
    grid-template-rows: 100%;
}

.right-whitespace{
    grid-column-start: 3;
    grid-column-end: 4;
}

.lang-selection{
    display: flex;
    flex-direction: row;
    flex-wrap: wrap;
    align-items: stretch;
    align-content: stretch;
    width: 60%;

    grid-column-start: 1;
    grid-column-end: 2;
}

.lang-filter{
    background: white;
    border: 1px solid #0a81d1;
    padding: 15px 15px;
    text-align: center;
    text-decoration: none;
    border-radius: 5px;

    cursor: pointer;

    flex-grow: 1;
}

.selected-lang-filter{
    background: #0a81d1;
    color: #f2f2f2;
}

.search-button-parent{
    grid-column-start: 2;
    grid-column-end: 3;
}

.search-button{
    background: #0a81d1;
    color: #f2f2f2;
    border: none;
    text-align: center;
    text-decoration: none;
    border-radius: 5px;
    box-sizing: border-box;

    width: 100px;
    padding-top: 10px;
    padding-bottom: 10px;

    cursor: pointer;
    float: right;
}

.search-button:hover {
    background: #3aa9f2;
}

@media screen and (max-width: 930px) {
    .search-container{
        grid-template-columns: 0% auto 0%;
    }

    .lang-selection{
        width: 100%;
    }
}

</style>