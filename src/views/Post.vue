<template>
    <div class="titleBox">
        <div class="postTitle">
            {{post.title}}
        </div>
    </div>
    <div class="MainContainer">
        <div class="articleContent">
            <h2>{{post.title}}</h2>
            <p>{{post.content}}</p>
        </div>
    </div>
</template>


<script>
export default {
    name: "Post",
    data(){
        return {
            post: {
                title: "这是一个文章的标题",
                content: "这是一个文章的内容"
            }
        }
    }
}
</script>

<style>
.articleContent{
    padding: 2rem 4rem 2rem 4rem;
}

.titleBox .postTitle {
    letter-spacing: 0.2rem;
    font-size: 2.5rem;
}
</style>
