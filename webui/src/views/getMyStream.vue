<script>
export default {
  data: function() {
    return {
      errormsg: null,
      loading: false,
      photos:[],
      usernames:[],
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

        for (let i = 0; i < this.photos.length; i++){
          // petition
          let url = "/username/"+this.photos[i].UserId;
          const name = await this.$axios.get(url,{
            headers:{"Authorization": this.photos[i].UserId}}).then(res => res.data);
          this.usernames.push(name)
        }

        console.log(this.photos)



      } catch (e) {
        this.errormsg = e.toString();
      }
      this.loading = false;
    },

    likePhoto: async function(p) {

      this.loading = true;
      this.errormsg = null;
      try {
        let castPhoto = JSON.parse(JSON.stringify(p))
        let url = "/users/"+this.token+"/like/"+castPhoto.ID;

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
          if (this.photos[i].ID == photo.ID) {
            // update likes list
            this.photos[i].Likes = photo.Likes
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
        if (photos[i].Id == photo["id"]){

          let likes = JSON.parse(photos[i].Likes);

          for (let j = 0; j < likes.length; j++) {
            if (likes[j] == this.token) {
              tokenLiked = true;
            }
          }
        }
      }
      return tokenLiked

    },



    commentPhoto: async function(p) {

      this.loading = true;
      this.errormsg = null;
      try {

        console.log(this.photos.length)


        // Getting the comment
        let castPhoto = JSON.parse(JSON.stringify(p))
        console.log(castPhoto);
        let url = "/users/"+this.token+"/comment/"+castPhoto.ID;
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

      await this.getUserStream();
      this.loading = false;
    }



  },
  async mounted() {
    this.token = document.cookie.replace(/(?:(?:^|.*;\s*)token\s*\=\s*([^;]*).*$)|^.*$/, "$1");
    await this.getUserStream()
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


      <template v-if="photos != null && photos.length > 0">
        <div class="m-3" v-for="(p,index) in photos" :key="index">
          <div class="card m-3" style="border-radius: 15px;">
            <div class="d-flex p-2 mb-2" style="border-radius: 15px;background-color: #efefef;">
              <div class="m-3">


                <h5>User: {{ usernames[index] }} </h5>
                <img v-if="p.Path" :src="p.Path" v-bind:alt="Photo" class="img-fluid m-3" style="border-radius: 10px; max-width: 100%; width: 300px; height: 200px;">
              </div>

              <div class="m-3">
                <h5 class="mb-1 m-3 p-3">Photo comments: </h5>

                <template v-if="JSON.parse(p.Comments).length > 0">
                  <div v-for="(c,index) in JSON.parse(p.Comments)" :key="index" >
                    <p class="">{{c['UserId']}} : {{c['Content']}}</p>
                  </div>

                </template>


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
                  <h6 class="mb-1 ">Photo ID: {{ p.ID }}</h6>
                  <p class="mb-1">Likes: {{ JSON.parse(p.Likes).length }}</p>
                  <p class="mb-1">Comments: {{ JSON.parse(p.Comments).length }}</p>
                </div>
              </div>
            </div>

          </div>
        </div>

      </template>

      <template v-else>

        <div class="card-body m-3">
          <h1 class="m-3">We're sorry, your followed users didn't upload anything yet.</h1>

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
