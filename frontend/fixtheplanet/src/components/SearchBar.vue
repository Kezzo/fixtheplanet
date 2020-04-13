<template>
    <div class="search-container">
        <div class="left-whitespace"></div>
        <div class="searchbar-inner-container">
            <div class="lang-selection">
                <div class="lang-filter" @click="select(lang)" v-bind:class="{'selected-lang-filter':lang.IsSelected }" 
                    v-bind:key="lang.Language" v-for="lang in selectedLanguages" v-bind:style="{ background: lang.Color, color: lang.TextColor }">
                    <div class="lang-filter-text">{{lang.Language}}</div>
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
            this.$emit('search-issues', true);
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

    padding-bottom: 0px;
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
    grid-template-rows: 100%;
    grid-template-columns: auto, 100px;
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

    grid-column-start: 1;
    grid-column-end: 2;

    width: 70%;
}

.lang-filter{
    padding: 15px 15px;
    margin: 5px 5px;
    text-align: center;
    text-decoration: none;
    border-radius: 5px;

    cursor: pointer;

    flex-grow: 1;
}

.selected-lang-filter{
    box-shadow: 0 0 5px 3px gray;
}

.search-button-parent{
    grid-column-start: 2;
    grid-column-end: 3;

    align-self: end;
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

.lang-filter:hover {
    background: #3aa9f2;
}

@media screen and (max-width: 930px) {
    .search-container{
        grid-template-columns: auto;
        grid-template-rows: auto auto;
    }

    .lang-selection{
        width: 100%;
    }

    .search-button{
        margin-bottom: 5px;
    }
}

</style>