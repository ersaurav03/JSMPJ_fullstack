<template>
  <div>
    <nav class="navbar" role="navigation" aria-label="main navigation">
      <div class="navbar-brand">
        <a
          role="button"
          class="navbar-burger burger"
          aria-label="menu"
          aria-expanded="false"
          data-target="navbarBasicExample"
        >
          <span aria-hidden="true"></span>
          <span aria-hidden="true"></span>
          <span aria-hidden="true"></span>
        </a>
      </div>

      <div id="navbarBasicExample" class="navbar-menu">
        <div class="navbar-start">
          <a class="navbar-item">
             <button @click="home">Home</button>
          </a>
          <a class="navbar-item">
             <button @click="about">About</button>
          </a>
          <a class="navbar-item">
            <button @click="contact">Contact</button>
          </a>
          <a class="navbar-item">
            <button @click="login">Login</button>
          </a>
        </div>
      </div>
    </nav>
    <section class="hero is-success is-fullheight">
      <div class="hero-body">
        <div class="container">
          <h1 class="title">{{info.data}}</h1>
          <h2 class="subtitle"></h2>
        </div>
      </div>
    </section>
    <footer class="footer">
      <div class="content has-text-centered">
        <p>
          <strong>Change URL</strong> by
          <a href="#"> Corporation</a>. The source code is licensed
          <a href="#">MIT</a>. The website content
          is licensed
          <a href="#">CC BY NC SA 4.0</a>.
        </p>
      </div>
    </footer>
  </div>
</template>

<script>
import axios from 'axios'
export default {
data () {
    return {
      info:[]
    }
  },
  mounted () {
    
  },
 methods: {
   home: function (){
     axios
      .get('http://localhost:8080/')
      .then(response => (this.info = response))
   },
   about: function (){
     axios
      .get('http://localhost:8080/about')
      .then(response => (this.info = response))
   },
   contact: function (){
     axios
      .get('http://localhost:8080/contact')
      .then(response => (this.info = response))
   },
   login: function (){
     axios
      .get('http://localhost:8080/login')
      .then(response => (this.info = response))
   }
 }


}
</script>
