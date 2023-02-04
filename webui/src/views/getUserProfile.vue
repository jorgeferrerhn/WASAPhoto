<script>
export default {
  data: function() {
    return {
      errormsg: null,
      loading: false,
      token: 0,
      user: {},
      users:[],
      isfollower: false
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

    isFollower: async function(){
      console.log("Calculate if it is follower: ");
      // Method to check if token user is follower of the searched user
      let nowIsFollowing = false;
      for (let i = 0; i < this.users.length; i++){
        let userFollowed = JSON.parse(JSON.stringify(this.users[i]));
        let followers = JSON.parse(userFollowed["Followers"]);

        console.log("Followers: ",followers);

        for (let j = 0; j < followers.length; j++){
          if (followers[i] == this.token){
            nowIsFollowing = true;
          }
        }

      }
      this.isfollower = nowIsFollowing;

    },

    getUser: async function() {

      this.loading = true;
      this.errormsg = null;
      try {
        let url = "/users/"+this.id+"/getUserProfile";
        // Let's get the cookie
        let getToken = document.cookie.replace(/(?:(?:^|.*;\s*)token\s*\=\s*([^;]*).*$)|^.*$/, "$1");
        this.token = getToken; // update logged user

        const response = await this.$axios.get(url,{
          headers:{'Authorization': this.token}
            }
        );
        this.user = response.data;

        let contains = false
        for (let i = 0; i < this.users.length; i++){
          if (this.users[i]["Id"] == this.user["Id"]){
            contains = true
          } // contains
        }
        if (!contains){
          this.users.push(this.user);
        }

      } catch (e) {
        this.errormsg = e.toString();
      }
      this.loading = false;
    },
    followUser: async function() {

      this.loading = true;
      this.errormsg = null;
      try {
        let url = "/users/"+this.token+"/followUser/"+this.id;
        console.log(url)
        const response = await this.$axios.put(url, "",{
          headers:{'Authorization': this.token,
          }
        });
        this.user = response.data;
        console.log(this.user)

      } catch (e) {
        this.errormsg = e.toString();
      }
      this.isFollower()
      this.loading = false;
    },
    unfollowUser: async function() {

      this.loading = true;
      this.errormsg = null;
      try {
        let url = "/users/"+this.token+"/unfollowUser/"+this.id;
        console.log(url)
        const response = await this.$axios.delete(url,{
          headers:{'Authorization': this.token,
          }
        });
        this.user = response.data;
        this.isFollower()
        console.log(this.user)

      } catch (e) {
        this.errormsg = e.toString();
      }
      this.loading = false;
    },


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
        <div class="col col-md-9 col-lg-7 col-xl-5" v-for="u in users" v-if="users.length > 0">
          <div class="card" style="border-radius: 15px;">
            <div class="card-body p-4">
              <div class="d-flex text-black">
                <div class="flex-shrink-0">
                  <img src="https://mdbcdn.b-cdn.net/img/Photos/new-templates/bootstrap-profiles/avatar-1.webp"
                       alt="Generic placeholder image" class="img-fluid"
                       style="width: 180px; border-radius: 10px;">
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

                    <!--If the user doesn't follow the target, a "Follow" button must be displayed. Otherwise, an "Unfollow" button will be displayer -->
                    <button type="button" class="btn btn-primary flex-grow-1" @click="followUser" v-if="!this.isfollower">Follow</button>
                    <button type="button" class="btn btn-primary flex-grow-1" @click="unfollowUser" v-if="this.isfollower">Unfollow</button>
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
