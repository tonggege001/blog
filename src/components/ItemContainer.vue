<template>
  <div class="ItemContainer">
    <div class="ItemContainerContent">
      <p class="title" @click="ItemContainerClick()">{{item.title}}</p>
      <p class="description">{{item.desc}}</p>
      <p class="date">{{formattedDate}}</p>
    </div>
  </div>
</template>


<script>
export default {
  name: "ItemContainer",
  props: {
    item: Object,
  },
  methods: {
    ItemContainerClick(){
      window.location.href = "/__post/"+this.item.html
    },

  },
  computed: {
    formattedDate(){
      return this.item.date.year+"年"+this.item.date.month+"月"+this.item.date.day+"日"+" · "+ this.item.tags.join(" ");
    }
  }
}
</script>

<style>
.ItemContainer {
  width: 100%;
  height: auto;
  box-sizing: border-box;
  border-radius: 10px;
  background-color: #444444;
  margin: 1rem 0 10px 0;
  transition: all 0.2s;
}

.ItemContainer:hover {
  transform: scale(1.03);
  box-shadow: 1px 1px 5px gray;
}

.ItemContainerContent{
    width: 100%;
    box-sizing: border-box;
    padding: 0px 28px 0px 28px;

    display: flex;
    flex-direction: column;
    justify-content: space-around
}

.ItemContainer .ItemContainerContent .title{
    color: #dddddd;
    text-align: left;
    font-size: 1.5rem;
    font-weight: 800;

    display: inline-block;
    margin-block-start: 0.5em;
    margin-block-end: 1em;
    margin-inline-start: 0px;
    margin-inline-end: auto;


}

.ItemContainer .ItemContainerContent .title:hover{
  cursor: pointer;
}

.ItemContainer .ItemContainerContent .description{
  width: 100%;
    text-align: left;
    color: #959b9e;
    font-size: 0.8rem;
    font-weight: 500;
    box-sizing: border-box;

    margin-block-start: 0em;
    margin-block-end: 2em;
    margin-inline-start: 0px;
    margin-inline-end: 0;

}

.ItemContainer .ItemContainerContent .date{
    color: #dddddd;
    text-align: left;

    font-size: 0.9rem;
    font-weight: 600;

    display: inline-block;
    margin-block-start: 0;
    margin-block-end: 1rem;
    margin-inline-start: 0px;
    margin-inline-end: auto;


}

/* 设置p或者div的大小随着文字内容变换：display: inline-block，并且margin-inline-end设置为auto */

</style>
