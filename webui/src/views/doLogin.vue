<script>
export default {
  data: function() {
    return {
      errormsg: null,
      loading: false,
      users: [],
      token: 0,
      loggedName: "",
      user: {},
      name : ""
    }
  },
  methods: {
    load() {
      return load
    },

    async refresh() {
      this.loading = true;
      this.errormsg = null;
      this.$router.push("/");
      this.loading = false;
    },

    doLogin: async function() {

      this.loading = true;
      this.errormsg = null;
      try {
        let url = "/session"
        this.name = this.search
        const response = await this.$axios.post(url, this.name).then(res => res);
        this.user = response.data;
        this.token = response.data["Id"];

        // Set the cookie for the logged user
        document.cookie = `token=`+this.token;
        document.cookie = `logged=`+this.name;

      } catch (e) {
        this.errormsg = e.toString();
      }
      this.loading = false;
    },

    getUser: async function(){
      // Getting user information
      let tokenUrl = "/users/"+this.token+"/profile";
      let userToken = await this.$axios.get(tokenUrl,{
            headers:{"Authorization": this.token}
          }
      ).then(res => res);
      console.log(userToken);
      this.user = userToken.data;

    },

    logOut: async function() {
      console.log("Logging out")
      document.cookie = `token=0`; // reset the cookie
      document.cookie = `logged=`; // reset the cookie
      this.token = document.cookie.replace(/(?:(?:^|.*;\s*)token\s*\=\s*([^;]*).*$)|^.*$/, "$1"); // update logged user
      this.token = document.cookie.replace(/(?:(?:^|.*;\s*)logged\s*\=\s*([^;]*).*$)|^.*$/, "$1"); // update logged user

    },

  },
  async mounted() {
    this.token = document.cookie.replace(/(?:(?:^|.*;\s*)token\s*\=\s*([^;]*).*$)|^.*$/, "$1"); // get token
    this.name = document.cookie.replace(/(?:(?:^|.*;\s*)logged\s*\=\s*([^;]*).*$)|^.*$/, "$1"); // get token

    }

}

</script>

<template>
  <div>
    <div
        class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
      <h1 class="h2">Login Page</h1>

      <div class="btn-toolbar mb-2 mb-md-0">



      </div>
    </div>

    <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
    <LoadingSpinner v-if="loading"></LoadingSpinner>



    <div class="card">

      <template v-if="token == 0">
        <h3 class="h3 mx-auto mt-3">Introduce your username: </h3>
        <input v-model="search" class="text-center mx-auto my-auto mt-3 mb-3" placeholder=" Username">

        <a href="javascript:" class="btn btn-primary mx-auto my-auto mb-3" @click="doLogin">Login</a>

      </template>

      <template v-else>
        <div class="card ">
          <div class="card-header">

          </div>
          <div class="card ">
            <div class="card-header">

            </div>
            <div class="card-body">
              <h3 class="text-center">Welcome to WASAPhoto, {{ this.name}} !</h3>
              <h4 class="text-center">You can now navigate through the page.</h4>
              <h5 class="text-center">Your token is: {{ token }}</h5>
            </div>
          </div>

          <a href="javascript:" class="btn btn-primary mx-auto my-auto mb-3" @click="logOut">Log Out</a>

        </div>
      </template>


    </div>


  </div>
</template>

<style scoped>
.card {
  margin-bottom: 20px;
}
</style>
