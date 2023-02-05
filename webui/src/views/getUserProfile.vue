<script>
export default {
  data: function() {
    return {
      errormsg: null,
      loading: false,
      token: 0,
      tokenUser:{},
      users:[]
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




    getUser: async function() {

      this.loading = true;
      this.errormsg = null;
      try {
        // The first thing we will do is save the token user object

        // Let's get the cookie
        this.token = document.cookie.replace(/(?:(?:^|.*;\s*)token\s*\=\s*([^;]*).*$)|^.*$/, "$1"); // update logged user
        let tokenUrl = "/users/"+this.token+"/getUserProfile";
        const userToken = await this.$axios.get(tokenUrl,{
              headers:{"Authorization": this.token}
            }
        );
        this.tokenUser = userToken.data;

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
      console.log("User to search: ",user)
      let tokenIsBanned = false;

      let banned = JSON.parse(this.tokenUser["Banned"]);
      console.log("Banned: ",banned)
      for (let j = 0; j < banned.length; j++) {
        if (banned[j] == user["Id"]) {
          tokenIsBanned = true
        }
      }


      return tokenIsBanned

    },

    getProfilePic: function(u){
      // Function that returns the profile pic path
      let user = JSON.parse(JSON.stringify(u))
      console.log("getting pic,",user)

      if (user["ProfilePic"] != 0){
        // Profile picture is not empty
        try {
          let tokenUrl = "/images/"+user["ProfilePic"];
          const image = this.$axios.get(tokenUrl, {
                headers: {"Authorization": this.token}
              }
          );
          console.log(image)

          let finalpath = string(image.data['path'])
          console.log(finalpath)
          return finalpath

        }catch (e) {
          this.errormsg = e.toString();
        }
      }
      console.log("Va a devovler la imagen de prueba")
      return "default-profile-photo.jpeg"
    }




  },

  computed: {




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

        <h3 class="h3">Introduce user ID to search: </h3>

        <input v-model="id" placeholder=" Search for a user...">



        <!-- User information -->
        <div class="col col-md-9 col-lg-7 col-xl-5" v-for="(u,index) in users" :key="index" v-if="users.length > 0">
          <div class="card" style="border-radius: 15px;">
            <div class="card-body p-4">
              <div class="d-flex text-black">
                <div class="flex-shrink-0">

                  <img :src="getProfilePic(u)"
                        class="img-fluid"
                       style="width: 180px; border-radius: 10px;">

                  <img src="/profilepics/default-profile-photo.jpeg"
                       class="img-fluid"
                       style="width: 180px; border-radius: 10px;">

                  <img :src="'/profilepics/'+getProfilePic(u)" v-bind:alt="getProfilePic(u)" class="img-fluid" sizes="(min-width: 991px) 10vw, (min-width: 768px) 20vw, 300px"
                       style="border-radius: 10px; max-width: 100%; width: 275px; height: 183px;">





                </div>
                <div class="flex-grow-1 ms-3">
                  <h5 class="mb-1">{{ u["Name"]}}</h5>
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
                      <button type="button" class="btn btn-primary flex-grow-1" @click="unfollowUser(u)">Unfollow</button>
                    </template>

                    <template v-else>
                      <button type="button" class="btn btn-primary flex-grow-1" @click="followUser(u)">Follow</button>
                    </template>

                    <!--If the user has not banned the target, a "Ban" button must be displayed. Otherwise, an "Unban" button will be displayed -->

                    <template v-if="isBanned(u)">
                      <button type="button" class="btn btn-primary flex-grow-1" @click="unbanUser(u)">Unban</button>
                    </template>

                    <template v-else>
                      <button type="button" class="btn btn-primary flex-grow-1" @click="banUser(u)">Ban</button>
                    </template>

                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>



        <a href="javascript:" class="btn btn-primary" @click="getUser">Search for user</a>
      </div>
    </div>


  </div>
</template>

<style scoped>
.card {
  margin-bottom: 20px;
}
</style>
