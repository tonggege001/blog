# blog

用Vue写了个博客。  
需要安装的Vue插件：store, router, axios  

npm run build 后，将dist目录内容copy过来就可以直接用了，但是为了能够发布到GitHub Pages上，还需要做一些改变。  

## 本地生成静态博客页面
首先Github Pages是一个静态Web服务器，也就是说没有后端，所以相当于本地写一个静态html页面作为博客内容，那么简单的方法是用markdown写完以后，使用VScode的Markdown Preview Enhenced在保存时自动在同级目录下生成对应的html页面。

## 收集Meta信息  
这一步本来是可以不需要的，完全通过客户端浏览器解析下—__post目录的所有html文件就可以了，然而Github Pages好像不支持HTTP的目录请求(可能我打开方式不对？)，所以本地需要收集一些博客页面的标题、摘要描述、tags、最近的修改日期、博客文章的创建日期，然后写到一个表里记录下来。  

客户浏览器在创建首页渲染时，create的钩子首先会请求这个meta文件解析出来所有的博客标题、tags、摘要描述，然后渲染整个页面。  

## 页面跳转  
用户点击后直接路由到/__post/对应的博客.html即可  

## 注意事项  
### 请求不到meta信息或者博客的html信息
在生成的dist目录内要加上.nojekyll 空文件，显式地告诉Github Pages不使用jekyll页面生成器。Hexo这类页面生成器都是需要在本地编译出html再上传到Github的，该博客的行为方式也和Hexo类似。Jekyll因为是github自己的页面生成器，可以在线渲染markdown生成html，所以内部提供了很多方便的接口（如我之前遇到的问题，获取目录下的所有文件）。但是Jekyll缺点是太慢，而且对Latex公式以及markdown的错误容忍度较低，渲染的结果不太好，不如本地渲染后的html好，Markdown Preview Enhenced， yyds（也可能是我打开方式不太对？）。  

### 对图片的处理  
本地渲染html有一个好处是可以直接将图片转成Base64格式，这样再也不用担心图片丢失文档打开渲染不出来的情况了～，虽然markdown也有渲染base64插入图片的方式，但是会极度影响markdown文件的美观。  

### Vue页面在github pages上点击返回键后404  
这个问题源于Vue默认使用history模式，那么浏览器执行前URL:github.io\/home/，执行后URL: github.io\/post, 这些相对路径都是在本地浏览器渲染的，服务器并没有感知。然而GitHub是静态页面，返回页面后自动请求URL: github.io/home  向服务器，此时静态服务器/目录下没有/home目录，那么就会返回404了。  

解决方法：  
1. 显式使用哈希模式，让浏览器对服务器的域名固定成github.io  
2. 在本地页面放置一个404.html，GitHub Pages会自动将这个404返回给用户，此时才404里重定向到 github.io/home 的页面即可。

为了不影响我整个代码的修改，我采用了第二种方法。


### 目录下的BlogHtmlProcess  
该文件是处理meta信息的二进制的源代码，用Go写的100多行，实际上可以利用pandoc生成html，这样就可以将整个过程自动化了，但是pandoc生成html还要配置很多东西才能生成的非常好看，但是Markdown Preview Enhenced能够自己选择主题，本身就很好看了，所以采用后者。  


## Project setup
```
npm install
```

### Compiles and hot-reloads for development
```
npm run serve
```

### Compiles and minifies for production
```
npm run build
```

### Lints and fixes files
```
npm run lint
```

### Customize configuration
See [Configuration Reference](https://cli.vuejs.org/config/).
