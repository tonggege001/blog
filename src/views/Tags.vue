<template>
    <div class="titleBox">标签</div>
    <div class="MainContainer">
        <div class="TagButtonMenu">
            <tag-button v-for="tagname in $store.getters.tagList" :key="tagname" :tagname="tagname"></tag-button>
        </div>

    </div>
</template>


<script>
import TagButton from '../components/TagButton.vue'

export default {
  components: { TagButton },
    name: "Tags",
    
}
</script>

<style>

.MainContainer .TagButtonMenu{
    width: 100%;
    padding: 2rem 4rem 2rem 4rem;
    box-sizing: border-box;
    text-align: left;
}
</style>





