<template>
  <div class="titleBox">主页</div>
  <div class="HomeMainContainer">
    <item-container v-for="item in $store.state.post_list.slice(currentPage*10,currentPage*10+10)" v-bind:key="item.title" v-bind:item="item"></item-container>
  </div>
  <div class="pageContainer">
    <div class="pageBtn" :style="currentPage>0?'':'visibility: hidden'" @click="pageLast">上一页</div> 
    <div class="pageBtn" :style="currentPage*10 + 10>=$store.state.post_list.length?'visibility: hidden':''"  @click="pageNext">下一页</div>
  </div>
</template>


<script>
import ItemContainer from "../components/ItemContainer.vue"
export default {
  name: "Home",
  components: {
    ItemContainer,
  },
  data(){
    return {
      currentPage: 0
    }
  },
  methods: {
    pageLast(){
      this.currentPage -= 1
    },
    pageNext(){
      this.currentPage += 1
    }

  }
};
</script>

<style>
.titleBox{
  width: 100%;
  height: 12rem;
  color: aliceblue;
  text-align: center;
  letter-spacing: 1.5rem;
  font-size: 4rem;
  font-weight: 900;
  background-color: #3b2120;
  line-height: 12rem;
  box-shadow: 0px 5px 5px rgba(109, 49, 14, 0.1);
}

.HomeMainContainer{
    width: 80%;
    box-sizing: border-box;
    /* 使div居中 */
    margin: 0 auto;  
}

.pageContainer{
  width: 80%;
  box-sizing: border-box;
  margin: 2rem auto;
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  color: #dddddd;
  font-size: 1.6rem;
  font-weight: 800;
}

.pageContainer .pageBtn {
  padding: 0.5rem 0.8rem 0.5rem 0.8rem;
  margin: 0
}

.pageContainer .pageBtn:hover {
  cursor: pointer;
  color: #8aaded
}


</style>
