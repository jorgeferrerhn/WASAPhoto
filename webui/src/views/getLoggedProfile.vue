<script>
export default {
  data: function() {
    return {
      errormsg: null,
      loading: false,
      token: 0,
      tokenUser:{},
      photos:[],
      logo:""
    }
  },
  methods: {
    load() {
      return load
    },

    async refresh() {
      this.loading = true;
      this.errormsg = null;
      this.$router.push("/getLoggedProfile");
      this.loading = false;
    },

    deletePhoto: async function(i){
      this.loading = true;
      this.errormsg = null;
      try {
        let tokenUrl = "/users/"+this.token+"/photo/"+this.photos[i]["id"];
        const userToken = await this.$axios.delete(tokenUrl,{
              headers:{"Authorization": this.token}
            }
        ).then(res => res);
        this.tokenUser = userToken.data;
        console.log(this.tokenUser)

        await this.getUserStream()





      } catch (e) {
        console.log("Error")
        this.errormsg = e.toString();
      }

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
        console.log(this.tokenUser)



      } catch (e) {
        console.log("Error")
        this.errormsg = e.toString();
      }

      this.loading = false;


    },

    getUserStream: async function() {

      this.loading = true;
      this.errormsg = null;
      try {
        let url = "/users/"+this.token+"/stream";
        const photos = await this.$axios.get(url,{
          headers:{"Authorization": this.token}}).then(res => res);
        this.photos = photos.data



      } catch (e) {
        this.errormsg = e.toString();
      }
      this.loading = false;
    },





    getProfilePic: async function(i){
      // Function that returns the profile pic path
      let user = JSON.parse(JSON.stringify(this.tokenUser));

      // Profile picture is not empty
      try {

        console.log(JSON.parse(this.tokenUser["Photos"]).length)
        // First, we have to search for the logo
        if (this.tokenUser["ProfilePic"] != undefined){
          let logoUrl = "/users/"+this.tokenUser["Id"]+"/logo";
          const logo = await this.$axios.get(logoUrl, {
                headers: {"Authorization": this.token}
              }
          ).then(res => res);

          this.logo = logo.data["path"];

          return logo.data["path"];

        }
        else{
          console.log("False")
          return 0
        }



      }catch (e) {
        this.errormsg = e.toString();
      }

      return 0


    },
  },

  async mounted() {
    this.token = document.cookie.replace(/(?:(?:^|.*;\s*)token\s*\=\s*([^;]*).*$)|^.*$/, "$1"); // get token
    if (this.token != 0){
      await this.updateLogged()
      await this.getProfilePic()
      await this.getUserStream()
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




        <!-- User information -->
        <template v-if="this.token != 0">
          <div  class="flex-grow-1 ms-3">
            <h2 class="mb-1">Your profile information, {{ this.tokenUser["Name"]}} </h2>

            <img v-if="logo != null" :src="this.logo" v-bind:alt="Logo" class="img-fluid m-3" style="border-radius: 10px; max-width: 100%; width: 300px; height: 200px;">

            <div class="flex-shrink-0">
            </div>


            <template v-if="this.logo.length > 0">
              <div class="d-flex justify-content-start rounded-3 p-2 mb-2"
                   style="background-color: #efefef;">
                <div>
                  <p class="small text-muted mb-1">Photos uploaded</p>
                  <p class="mb-0">{{ JSON.parse(this.tokenUser["Photos"]).length }}</p>
                </div>
                <div class="px-3">
                  <p class="small text-muted mb-1">Followers </p>
                  <p class="mb-0">{{ JSON.parse(this.tokenUser["Followers"]).length }}</p>
                </div>
              </div>

            </template>



          </div>



        </template>

        <!-- Photo information -->
        <template v-if="photos != null">
          <div class="m-3" v-for="(p,index) in photos" :key="index">
            <h5>Photo uploaded!</h5>
            <div class="card m-3" style="border-radius: 15px;">
              <div class="d-flex p-2 mb-2" style="border-radius: 15px;background-color: #efefef;">
                <div class="m-3">
                  <img v-if="p['path']" :src="p['path']" v-bind:alt="Photo" class="img-fluid m-3" style="border-radius: 10px; max-width: 100%; width: 300px; height: 200px;">


                </div>



              </div>

              <div class="d-flex align-items-center">

                <template v-if="p['likes'].length > 0">

                <div class="m-3 d-flex align-items-center">
                  <div class="m-3 d-flex align-items-center">
                    <h6 class="m-3 ">Photo ID: {{ p["id"]}}</h6>
                    <p class="m-3">Likes: {{ JSON.parse(p["likes"]).length }}</p>
                    <p class="m-3">Comments: {{ JSON.parse(p["comments"]).length }}</p>
                    <button type="button" @click="deletePhoto(index)" class="btn btn-primary" style="background-color:red; border-color: red"><span class="bi bi-trash">Delete photo</span> </button>
                  </div>
                </div>
                </template>
              </div>

            </div>
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
