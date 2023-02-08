<script>
export default {
  data: function() {
    return {
      errormsg: null,
      loading: false,
      token: 0,
      tokenUser:{},
      users:[],
      logopaths:[]
    }
  },
  methods: {
    load() {
      return load
    },

    async refresh() {
      this.loading = true;
      this.errormsg = null;
      this.$router.push("/getUserProfile");
      this.loading = false;
    },

    updateLogged: async function(){
      // Update logged user. This function is used on mounted() to check if the logged user has changed

      this.loading = true;
      this.errormsg = null;
      try {
        this.token = document.cookie.replace(/(?:(?:^|.*;\s*)token\s*\=\s*([^;]*).*$)|^.*$/, "$1"); // update logged user
        let tokenUrl = "/users/"+this.token+"/profile";
        const userToken = await this.$axios.get(tokenUrl,{
              headers:{"Authorization": this.token}
            }
        ).then(res => res);
        this.tokenUser = userToken.data;
        for (let i = 0; i < this.users.length; i++){
          if (this.users[i] == this.token){
            this.users[i] = userToken // Updating logged user
          }
        }

        console.log(this.users)
      } catch (e) {
        console.log("Error")
        this.errormsg = e.toString();
      }

      this.loading = false;


    },


    getUser: async function() {

      this.loading = true;
      this.errormsg = null;

      if (this.search != undefined){
        try {

          // We have to search first for the user by its username
          let sessionUrl = "/users/"+this.search;
          const response = await this.$axios.get(sessionUrl,{
            headers:{"Authorization": this.token,
            }
          }).then(res => res);

          let userId = response.data["Id"];
          // After doing this, we will be able to retrieve the user by its id

          if (userId != undefined){

            let user = response.data;
            let contains = false

            let users = JSON.parse(JSON.stringify(this.users));
            console.log(users)

            for (let i = 0; i < users.length; i++){
              if (users[i]["Id"] == user["Id"]){
                contains = true
              } // contains
            }
            if (!contains){
              users.push(user);
              this.logopaths.push("");
            }

            this.users = users;
            console.log(this.users)


          }



        } catch (e) {
          this.errormsg = e.toString();
        }

      }

      this.loading = false;
    },
    followUser: async function(u) {

      this.loading = true;
      this.errormsg = null;
      try {
        let url = "/users/"+this.token+"/followUser/"+u["Id"];
        const response = await this.$axios.put(url, "",{
          headers:{"Authorization": this.token,
          }
        }).then(res => res);;
        let user = response.data;


        // update this.users list to update followers
        for (let i = 0; i < this.users.length; i++) {
          if (this.users[i]["Id"] == user["Id"]){

            // update followers list
            this.users[i]["Followers"] = user["Followers"]
        }
        }

      } catch (e) {
        this.errormsg = e.toString();
      }


      this.loading = false;
    },
    unfollowUser: async function(u) {

      this.loading = true;
      this.errormsg = null;
      try {
        let url = "/users/"+this.token+"/unfollowUser/"+u["Id"];
        const response = await this.$axios.delete(url,{
          headers:{"Authorization": this.token,
          }
        }).then(res => res);;
        let user = response.data;


        // update this.users list to update followers
        for (let i = 0; i < this.users.length; i++) {
          if (this.users[i]["Id"] == user["Id"]) {
            // update followers list
            this.users[i]["Followers"] = user["Followers"]
          }
        }

      } catch (e) {
        this.errormsg = e.toString();
      }
      this.loading = false;
    },

    isFollower(user){
      // Method to check if token user is follower of the searched user

      console.log(user)
      let users = JSON.parse(JSON.stringify(this.users))

      let tokenIsFollower = false;
      for (let i = 0; i < users.length; i++){
        if (users[i]["Id"] == user["Id"]){
          let followers = JSON.parse(users[i]["Followers"]);
          for (let j = 0; j < followers.length; j++) {
            if (followers[j] == this.token) {
              tokenIsFollower = true;
            }
          }
        }
      }
      return tokenIsFollower

    },

    banUser: async function(u) {

      console.log("Ban")

      this.loading = true;
      this.errormsg = null;
      try {
        let url = "/users/"+this.token+"/banUser/"+u["Id"];
        const response = await this.$axios.put(url, "",{
          headers:{"Authorization": this.token,
          }
        }).then(res => res);;
        let user = response.data; // In this case, this returns the user 1

        this.tokenUser["Banned"] = user["Banned"];


      } catch (e) {
        this.errormsg = e.toString();
      }


      this.loading = false;
    },
    unbanUser: async function(u) {

      this.loading = true;
      this.errormsg = null;
      try {
        let url = "/users/"+this.token+"/unbanUser/"+u["Id"];
        const response = await this.$axios.delete(url,{
          headers:{"Authorization": this.token,
          }
        }).then(res => res);;
        let user = response.data; // In this case, this returns the user 1


        this.tokenUser["Banned"] = user["Banned"];

      } catch (e) {
        this.errormsg = e.toString();
      }
      this.loading = false;
    },

    tokenIsBanned(u){
      // Method to check if user has banned token user

      let user = JSON.parse(JSON.stringify(u))
      let tokenIsBanned = false;

      console.log(user)

      let banned = JSON.parse(user["Banned"]);
      console.log(banned)
      for (let j = 0; j < banned.length; j++) {
        if (banned[j] == this.token) {
          tokenIsBanned = true
        }
      }


      return tokenIsBanned

    },

    userIsBanned(u){
      // Method to check if token user has banned a user

      let user = JSON.parse(JSON.stringify(u))
      let tokenIsBanned = false;

      console.log(user)

      let banned = JSON.parse(this.tokenUser["Banned"]);
      console.log(banned)
      for (let j = 0; j < banned.length; j++) {
        if (banned[j] == user["Id"]) {
          tokenIsBanned = true
        }
      }


      return tokenIsBanned

    },


    getProfilePic: async function(i){
      // Function that returns the profile pic path
      let user = JSON.parse(JSON.stringify(this.users[i]));
      console.log(user)
      console.log("getting pic,",user);

      // Profile picture is not empty
      try {
        // First, we have to search for the logo
        let logoUrl = "/images/"+user["ProfilePic"];
        const logo = await this.$axios.get(logoUrl, {
              headers: {"Authorization": this.token}
            }
        ).then(res => res);;

        console.log(logo.data)
        this.logopaths[i] = logo.data["path"];

        console.log(this.logopaths)
        return logo.data["path"];

      }catch (e) {
        this.errormsg = e.toString();
      }


    },





  },

  async mounted() {
    this.token = document.cookie.replace(/(?:(?:^|.*;\s*)token\s*\=\s*([^;]*).*$)|^.*$/, "$1"); // get token
    if (this.token != 0){
      await this.updateLogged()

    }



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

        <h3 class="h3 mb-3 mt-3">Introduce user name to search: </h3>
        <input class="mb-3 mt-3" v-model="search" placeholder=" Search for a user...">
        <a href="javascript:" class="btn btn-primary m-3" @click="getUser">Search for user</a>


        <!-- User information -->
        <template v-if="users.length > 0 && token != 0">
          <div v-for="(u,index) in users" :key="index"  class="flex-grow-1 ms-3">

            <template v-if="!tokenIsBanned(u)">
              <h5 class="mb-1">{{ u["Name"]}}</h5>
              <img v-if="getProfilePic(index)" :src="logopaths[index]" v-bind:alt="Logo" class="img-fluid m-3" style="border-radius: 10px; max-width: 100%; width: 300px; height: 200px;">

              <div class="flex-shrink-0">

                <template >

                </template>




              </div>


              <div class="d-flex justify-content-start rounded-3 p-2 mb-2"
                   style="background-color: #efefef;">
                <div>
                  <p class="small text-muted mb-1">Photos uploaded</p>
                  <p class="mb-0">{{ JSON.parse(u["Photos"]).length }}</p>
                </div>
                <div class="px-3">
                  <p class="small text-muted mb-1">Followers </p>
                  <p class="mb-0">{{ JSON.parse(u["Followers"]).length }}</p>
                </div>
              </div>


              <div class="d-flex pt-1">

                <!--If the user doesn't follow the target, a "Follow" button must be displayed. Otherwise, an "Unfollow" button will be displayed -->

                <template v-if="isFollower(u)">
                  <button type="button" class="btn btn-primary flex-grow-1 mx-auto" @click="unfollowUser(u)">Unfollow</button>
                </template>

                <template v-else>
                  <button type="button" class="btn btn-primary flex-grow-1 mx-auto" @click="followUser(u)">Follow</button>
                </template>

                <!--If the user has not banned the target, a "Ban" button must be displayed. Otherwise, an "Unban" button will be displayed -->

                <template v-if="userIsBanned(u)">
                  <button type="button" class="btn btn-primary flex-grow-1 mx-auto" style="background-color:red" @click="unbanUser(u)">Unban</button>
                </template>

                <template v-else>
                  <button type="button" class="btn btn-primary flex-grow-1 mx-auto" @click="banUser(u)">Ban</button>
                </template>

              </div>

            </template>


          </div>



        </template>



      </div>

    </div>


  </div>
</template>

<style scoped>
.card {
  margin-bottom: 20px;
}
</style>
