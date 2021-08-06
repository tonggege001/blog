<template>
  <div class="titleBox">归档</div>
    <div class="MainContainer">
        <archive-year-box v-for="year in $store.getters.yearList" v-bind:key="year" :year="year"></archive-year-box>
    </div> 
</template>

<script>
import ArchiveYearBox from '../components/ArchiveYearBox.vue'

export default {
    name: "Archive",
    components: {
        ArchiveYearBox
    },
}
</script>

<style>

</style>


