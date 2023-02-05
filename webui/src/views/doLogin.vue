<script>
export default {
  data: function() {
    return {
      errormsg: null,
      loading: false,
      users: [],
      token: 0,
      loggedName: "",
      user: {}
    }
  },
  methods: {
    load() {
      return load
    },

    async refresh() {
      this.loading = true;
      this.errormsg = null;
      this.$router.push("/doLogin");
      this.loading = false;
    },

    doLogin: async function() {

      this.loading = true;
      this.errormsg = null;
      try {
        let url = "/session"
        const response = await this.$axios.post(url, this.name);
        this.user = response.data;
        this.token = response.data["Id"];

        // Set the cookie for the logged user
        document.cookie = `token=`+this.token;

      } catch (e) {
        this.errormsg = e.toString();
      }
      this.loading = false;
    },

    logOut: async function() {
      console.log("Logging out")
      document.cookie = `token=0`; // reset the cookie
      this.token = document.cookie.replace(/(?:(?:^|.*;\s*)token\s*\=\s*([^;]*).*$)|^.*$/, "$1"); // update logged user
    },

  },
  mounted() {
    this.refresh()
  }
}

</script>

<template>
  <div>
    <div
        class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
      <h1 class="h2">Login Page</h1>

      <div class="btn-toolbar mb-2 mb-md-0">

        <div class="btn-group me-2">
          <button type="button" class="btn btn-sm btn-outline-secondary" @click="refresh">
            Refresh
          </button>
        </div>

      </div>
    </div>

    <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
    <LoadingSpinner v-if="loading"></LoadingSpinner>



    <div class="card">

      <template v-if="token == 0">
        <h3 class="h3 mx-auto mt-3">Introduce your username: </h3>
        <input v-model="name" class="text-center mx-auto my-auto mt-3 mb-3" placeholder=" Username">

        <a href="javascript:" class="btn btn-primary mx-auto my-auto mb-3" @click="doLogin">Create a new user</a>

      </template>

      <template v-else>
        <div class="card ">
          <div class="card-header">

          </div>
          <div class="card ">
            <div class="card-header">

            </div>
            <div class="card-body">
              <h3 class="text-center">Welcome, {{ this.user["Name"]}}</h3>
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
