<script>
export default {
  data: function() {
    return {
      errormsg: null,
      loading: false,
      token: 0,
      tokenUser:{},
      users:[],
      imgSrc: '/default-profile-photo.jpeg'
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
        let tokenUrl = "/users/"+this.token+"/getUserProfile";
        const userToken = await this.$axios.get(tokenUrl,{
              headers:{"Authorization": this.token}
            }
        );
        this.tokenUser = userToken.data;
        for (let i = 0; i < this.users.length; i++){
          if (this.users[i] == this.token){
            console.log(this.tokenUser); // Updating logged user
          }
        }
      } catch (e) {
        console.log("Error")
        this.errormsg = e.toString();
      }

      this.loading = false;


    },


    getUser: async function() {

      this.loading = true;
      this.errormsg = null;
      try {
        await this.updateLogged();

        // Then, we search for the requested user
        let url = "/users/"+this.id+"/getUserProfile";
        const response = await this.$axios.get(url,{
          headers:{"Authorization": this.token}
            }
        );
        let user = response.data;
        let contains = false
        this.users = JSON.parse(JSON.stringify(this.users))


        for (let i = 0; i < this.users.length; i++){
          if (this.users[i]["Id"] == user["Id"]){
            contains = true
          } // contains
        }
        if (!contains){
          this.users.push(user);
        }



      } catch (e) {
        this.errormsg = e.toString();
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
        });
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
        });
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

    isFollower(u){
      // Method to check if token user is follower of the searched user

      let user = JSON.parse(JSON.stringify(u))
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
        });
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
        });
        let user = response.data; // In this case, this returns the user 1


        this.tokenUser["Banned"] = user["Banned"];

      } catch (e) {
        this.errormsg = e.toString();
      }
      this.loading = false;
    },

    isBanned(u){
      // Method to check if token user has banned the parameter user

      let user = JSON.parse(JSON.stringify(u))
      let tokenIsBanned = false;

      let banned = JSON.parse(this.tokenUser["Banned"]);
      for (let j = 0; j < banned.length; j++) {
        if (banned[j] == user["Id"]) {
          tokenIsBanned = true
        }
      }


      return tokenIsBanned

    },


    getProfilePic: async function(u){
      // Function that returns the profile pic path
      let user = JSON.parse(JSON.stringify(u));
      let finalPath = "";
      console.log("getting pic,",user);

      if (user["ProfilePic"] != 0){
        // Profile picture is not empty
        try {
          let tokenUrl = "/images/"+user["ProfilePic"];
          const data = await this.$axios.get(tokenUrl, {
                headers: {"Authorization": this.token}
              }
          ).then((image)=>{return image.data});

          console.log(data["path"]);
          return data["path"];

        }catch (e) {
          this.errormsg = e.toString();
        }
      }
      finalPath = "default-profile-photo.jpeg";
      return finalPath;
    }




  },

  async mounted() {
    this.token = document.cookie.replace(/(?:(?:^|.*;\s*)token\s*\=\s*([^;]*).*$)|^.*$/, "$1"); // get token
    const profilePic = await getProfilePic(this.tokenUser)


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

        <h3 class="h3 mb-3 mt-3">Introduce user ID to search: </h3>

        <input class="mb-3 mt-3" v-model="id" placeholder=" Search for a user...">
        <a href="javascript:" class="btn btn-primary m-3" @click="getUser">Search for user</a>


        <!-- User information -->
        <div class="col col-md-9 col-lg-7 col-xl-5 mt-3 mb-3" v-for="(u,index) in users" :key="index" v-if="users.length > 0">
          <div class="card mt-3 d-lg-grid" style="border-radius: 15px;">


              <h5 class="m-3 p-3">{{ u["Name"]}}</h5>
            <template>
              <img :src="imgSrc" v-bind:alt="getProfilePic(u)" class="avatar img-fluid m-3"
                   style="border-radius: 10px; width: 100px; height: 100px;">

            </template>




            <div class="d-flex mb-2"
                       style="background-color: #efefef;">
                <div>
                  <p class="small text-muted m-3">Photos uploaded</p>
                  <p class="m-3">{{ JSON.parse(u["Photos"]).length }}</p>
                </div>

                <div class="px-3">
                  <p class="small text-muted m-3">Followers </p>
                  <p class="m-3">{{ JSON.parse(u["Followers"]).length }}</p>
                </div>

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

            <template v-if="isBanned(u)">
              <button type="button" class="btn btn-primary flex-grow-1 mx-auto" @click="unbanUser(u)">Unban</button>
            </template>

            <template v-else>
              <button type="button" class="btn btn-primary flex-grow-1 mx-auto" @click="banUser(u)">Ban</button>
            </template>

          </div>
        </div>




      </div>

    </div>


  </div>
</template>

<style scoped>
.card {
  margin-bottom: 20px;
}
</style>
