<script>
export default {
  data: function() {
    return {
      errormsg: null,
      loading: false,
      token: 0,
      user: {},
      path:""
    }
  },
  methods: {
    load() {
      return load
    },

    async refresh() {
      this.loading = true;
      this.errormsg = null;
      this.$router.push("/setMyUserName");
      this.loading = false;
    },

    setMyUserName: async function() {

      this.loading = true;
      this.errormsg = null;
      if (this.name != undefined){
        try {
          let url = "/users/"+this.token+"/username";
          const response = await this.$axios.put(url, this.name,{
                headers:{"Authorization": this.token}

          }
          ).then(res => res);
          this.user = response.data;
          console.log(this.user)

          this.name = this.user['Name']
          document.cookie = `logged=`+this.name;

        } catch (e) {
          this.errormsg = e.toString();
        }

      }

      this.loading = false;
    },

    getProfilePic: async function(){
      // Function that returns the profile pic path
      let user = JSON.parse(JSON.stringify(this.user));
      console.log(user)
      console.log("getting pic,",user);

      // Profile picture is not empty
      if (this.user != undefined && user["ProfilePic"] != undefined){
        try {
          // First, we have to search for the logo
          let logoUrl = "/images/"+user["ProfilePic"];
          const logo = await this.$axios.get(logoUrl, {
                headers: {"Authorization": this.token}
              }
          ).then(res => res);

          console.log(logo.data)
          this.path = logo.data["path"];

          return logo.data["path"];

        }catch (e) {
          this.errormsg = e.toString();
        }

      }



    },

    getUser: async function(){
      // Getting user information
      let tokenUrl = "/users/"+this.token+"/profile";
      let userToken = await this.$axios.get(tokenUrl,{
            headers:{"Authorization": this.token}
          }
      ).then(res => res);
      this.user = userToken.data;

    }



  },
  async mounted() {
    this.token = document.cookie.replace(/(?:(?:^|.*;\s*)token\s*\=\s*([^;]*).*$)|^.*$/, "$1"); // get token
    this.name = document.cookie.replace(/(?:(?:^|.*;\s*)logged\s*\=\s*([^;]*).*$)|^.*$/, "$1"); // get token
    await this.getUser();
  }
}
</script>

<template>
  <div>
    <div
        class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">


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
      <div class="card-body">

        <h3 class="h3">Introduce your new user name...: </h3>
        <input v-model="name" placeholder="jose">


        <!-- User information -->
        <template v-if="1">
          <div class="col m-3 col-md-9 col-lg-7 col-xl-5" >
            <div class="card" style="border-radius: 15px;">
              <div class="card-body p-4">
                <!-- User information -->
                <div class="d-flex text-black">



                  <div class="flex-shrink-0">

                    <template v-if="user['ProfilePic'] != 0">
                      <img v-if="getProfilePic()" :src="this.path" v-bind:alt="Photo" class="img-fluid m-3" style="border-radius: 10px; max-width: 100%; width: 300px; height: 200px;">
                    </template>

                  </div>
                  <div class="flex-grow-1 ms-3">
                    <h5 class="mb-1">{{ user["Name"]}}</h5>
                    <div class="d-flex justify-content-start rounded-3 p-2 mb-2"
                         style="background-color: #efefef;">
                      <!--div>
                        <p class="small text-muted mb-1">Photos</p>
                        <p class="mb-0">{{ JSON.parse(user["Photos"]).length }}</p>
                      </div>
                      <div class="px-3">
                        <p class="small text-muted mb-1">Followers </p>
                        <p class="mb-0">{{ JSON.parse(user["Followers"]).length }}</p>
                      </div-->
                    </div>

                  </div>



                </div>
              </div>
            </div>
          </div>
        </template>

        <a href="javascript:" class="btn btn-primary" @click="setMyUserName">Set new username</a>
      </div>
    </div>


  </div>
</template>

<style scoped>
.card {
  margin-bottom: 20px;
}
</style>
