<script>
export default {
  data: function() {
    return {
      errormsg: null,
      loading: false,
      photos:[],
      token: 0
    }
  },
  methods: {
    load() {
      return load
    },

    async refresh() {
      this.loading = true;
      console.log("Refresh")
      this.errormsg = null;
      this.$router.push("/getMyStream");
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

    likePhoto: async function(p) {

      this.loading = true;
      this.errormsg = null;
      try {
        let url = "/users/"+this.token+"/like/"+p["id"];
        const response = await this.$axios.put(url, "",{
          headers:{"Authorization": this.token,
          }
        }).then(res => res);
        let photo = response.data;


        // update this.photos list to update likes
        for (let i = 0; i < this.photos.length; i++) {
          if (this.photos[i]["id"] == photo["id"]){

            // update likes list
            this.photos[i]["likes"] = photo["likes"]
          }
        }


      } catch (e) {
        this.errormsg = e.toString();
      }


      this.loading = false;
    },
    unlikePhoto: async function(p) {

      this.loading = true;
      this.errormsg = null;
      try {
        let url = "/users/"+this.token+"/dislike/"+p["id"];
        const response = await this.$axios.delete(url,{
          headers:{"Authorization": this.token,
          }
        }).then(res => res);
        let photo = response.data;


        // update this.photos list to update likes
        for (let i = 0; i < this.photos.length; i++) {
          if (this.photos[i]["id"] == photo["id"]) {
            // update likes list
            this.photos[i]["likes"] = photo["likes"]
          }
        }

      } catch (e) {
        this.errormsg = e.toString();
      }
      this.loading = false;
    },

    isLiked(p){
      // Method to check if token user liked the searched photo

      let photo = JSON.parse(JSON.stringify(p))
      let photos = JSON.parse(JSON.stringify(this.photos))
      

      let tokenLiked = false;
      for (let i = 0; i < photos.length; i++){
        if (photos[i]["id"] == photo["id"]){
          let likes = JSON.parse(photos[i]["likes"]);
          console.log(likes)
          for (let j = 0; j < likes.length; j++) {
            if (likes[j] == this.token) {
              tokenLiked = true;
            }
          }
        }
      }
      return tokenLiked

    },

    getName: async function(i){
      this.loading = true;
      this.errormsg = null;
      let user = null;
      try {
        let url = "/users/"+i+"/profile";
        user = await this.$axios.get(url,{
          headers:{"Authorization": this.token}}).then(res => res);

      } catch (e) {
        this.errormsg = e.toString();
      }
      this.loading = false;
      return user.data["Name"]
    },

    commentPhoto: async function(p) {

      this.loading = true;
      this.errormsg = null;
      try {

        console.log(this.photos.length)


        // Getting the comment

        let url = "/users/"+this.token+"/comment/"+p["id"];
        const response = await this.$axios.post(url, this.comment,{
          headers:{"Authorization": this.token,
          }
        }).then(res => res);
        let photo = response.data;
        console.log(photo)


        // update this.photos list to update likes
        for (let i = 0; i < this.photos.length; i++) {
          if (this.photos[i]["id"] == photo["id"]){

            // update likes list
            this.photos[i]["comments"] = photo["comments"]
          }
        }



      } catch (e) {
        this.errormsg = e.toString();
      }


      this.loading = false;
    }



  },
  async mounted() {
    this.token = document.cookie.replace(/(?:(?:^|.*;\s*)token\s*\=\s*([^;]*).*$)|^.*$/, "$1");
    if (this.search != undefined){
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



    <div class="card m-3">
      <div class="card-body m-3">
        <h1 class="m-3">Press button to search your stream...</h1>

        <button class="btn btn-primary m-3" @click="getUserStream">Search</button>
      </div>

      <template v-if="photos.length > 0">
        <div class="m-3" v-for="(p,index) in photos" :key="index">
          <div class="card m-3" style="border-radius: 15px;">
            <div class="d-flex p-2 mb-2" style="border-radius: 15px;background-color: #efefef;">
              <div class="m-3">
                <img v-if="p['path']" :src="p['path']" v-bind:alt="Photo" class="img-fluid m-3" style="border-radius: 10px; max-width: 100%; width: 300px; height: 200px;">
              </div>

              <div class="m-3">
                <h5 class="mb-1 m-3 p-3">Photo comments: </h5>

                <div v-for="(c,index) in JSON.parse(p['comments'])" :key="index" >
                  <p class="">{{c['UserId']}} : {{c['Content']}}</p>
                </div>

              </div>

            </div>

            <div class="d-flex align-items-center">
              <div class="m-3">
                <input v-model="comment" placeholder="Add a comment..." class="form-control">
                <div class="d-flex align-items-center">
                  <button class="btn btn-primary m-3" @click="commentPhoto(p)">Comment</button>

                  <template v-if="!isLiked(p)">
                    <button class="btn btn-primary m-3" @click="likePhoto(p)">Like</button>
                  </template>
                  <template v-else>
                    <button class="btn btn-primary m-3" @click="unlikePhoto(p)">Unlike</button>
                  </template>
                </div>


              </div>
              <div class="m-3">
                <div class="m-3">
                  <h6 class="mb-1 ">Photo ID: {{ p["id"]}}</h6>
                  <p class="mb-1">Likes: {{ JSON.parse(p["likes"]).length }}</p>
                  <p class="mb-1">Comments: {{ JSON.parse(p["comments"]).length }}</p>
                </div>
              </div>
            </div>

          </div>
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
