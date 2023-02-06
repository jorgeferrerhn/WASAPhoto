<script>
export default {
  data: function() {
    return {
      file:null,
      errormsg: null,
      previewImage:null,
      loading: false,
      token: 0,
      photo: {},
      photos:[]
    }
  },
  methods: {
    load() {
      return load
    },

    async refresh() {
      this.loading = true;
      this.errormsg = null;
      this.$router.push("/uploadPhoto");
      this.loading = false;
    },



    selectImage () {
      this.$refs.fileInput.click()
    },
    pickFile () {
      let input = this.$refs.fileInput
      let file = input.files
      if (file && file[0]) {
        let reader = new FileReader
        reader.onload = e => {
          this.previewImage = e.target.result
        }
        reader.readAsDataURL(file[0])
        this.$emit('input', file[0])
      }
    },


    onFileChange: function(event){
      let name =event.target.files[0]["name"];
      this.path= name; // update the path here
      console.log("Aquí", this.path);
      return this.path

      /*
      const reader = new FileReader()
      reader.readAsDataURL(this.file)
      reader.onload = e => {
        this.image = e.target.result
        console.log(this.image)
      }*/
    },


    uploadPhoto: async function() {

      this.loading = true;
      this.errormsg = null;
      try {
        // Let's get the cookie
        this.token = document.cookie.replace(/(?:(?:^|.*;\s*)token\s*\=\s*([^;]*).*$)|^.*$/, "$1");
        console.log(this.path)
        let url = "/users/"+this.token+"/uploadPhoto";
        const response = await this.$axios.post(url,this.path,{
          headers:{"Authorization": this.token}
        });
        let photo = response.data;


        let contains = false
        for (let i = 0; i < this.photos.length; i++){
          if (this.photos[i]["id"] == photo["id"]){
            contains = true
          }
        }


        if (!contains){
          this.photos.push(photo);
        }

      } catch (e) {
        this.errormsg = e.toString();
      }
      this.loading = false;
    },

    likePhoto: async function(p) {

      this.loading = true;
      this.errormsg = null;
      try {
        let url = "/users/"+this.token+"/likePhoto/"+p["id"];
        const response = await this.$axios.put(url, "",{
          headers:{"Authorization": this.token,
          }
        });
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
        let url = "/users/"+this.token+"/unlikePhoto/"+p["id"];
        const response = await this.$axios.delete(url,{
          headers:{"Authorization": this.token,
          }
        });
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

      console.log(photos)

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

    commentPhoto: async function(p) {

      console.log("Commenting")
      this.loading = true;
      this.errormsg = null;
      try {
        let url = "/users/"+this.token+"/commentPhoto/"+p["id"];
        const response = await this.$axios.post(url, this.comment,{
          headers:{"Authorization": this.token,
          }
        });
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
    },





  },
  mounted() {
    this.token = document.cookie.replace(/(?:(?:^|.*;\s*)token\s*\=\s*([^;]*).*$)|^.*$/, "$1");
    this.refresh()
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
        <h3 class="h3">Select new picture to upload...: </h3>

        <input type="file" @change="onFileChange"/>
        <!--v-file-input placeholder="Upload document" v-model="file" accept="image/*" label="Image" @change="onFileChange"/-->


        <!-- Photo information -->
        <div class="col col-md-9 col-lg-7 col-xl-5" v-for="(p,index) in photos" :key="index" v-if="photos.length > 0">
          <div class="card" style="border-radius: 15px;">
            <div class="card-body p-4">
              <div class="d-flex text-black">
                <div class="flex-shrink-0">

                  <div class="col-lg-2">
                    <img :src="'/img/'+p['path']" v-bind:alt="p['path']" class="img-fluid" sizes="(min-width: 991px) 10vw, (min-width: 768px) 20vw, 300px"
                         style="border-radius: 10px; max-width: 100%; width: 275px; height: 183px;">
                  </div>

                  <!--img src="/home/jorge/WASAPhoto/webui/src/images/car.jpeg"
                       alt="Generic placeholder image" class="img-fluid"
                       style="width: 180px; border-radius: 10px;"-->

                  <!--img src="https://mdbcdn.b-cdn.net/img/Photos/new-templates/bootstrap-profiles/avatar-1.webp"
                       alt="Generic placeholder image" class="img-fluid"
                       style="width: 180px; border-radius: 10px;"-->
                </div>
                <div class="flex-grow-1 ms-3">
                  <h5 class="mb-1">{{ p["id"]}}</h5>
                  <div class="d-flex justify-content-start rounded-3 p-2 mb-2"
                       style="background-color: #efefef;">
                    <div>
                      <p class="small text-muted mb-1">Likes</p>
                      <p class="mb-0">{{ JSON.parse(p["likes"]).length }}</p>
                    </div>
                    <div class="px-3">
                      <p class="small text-muted mb-1">Comments </p>
                      <p class="mb-0">{{ JSON.parse(p["comments"]).length }}</p>
                    </div>

                  </div>
                  <div class="d-flex pt-1">

                    <!--If the user didn't like the photo, a "Like" button must be displayed. Otherwise, an "Unlike" button will be displayed -->

                    <template v-if="!isLiked(p)">
                      <button type="button" class="btn btn-primary flex-grow-1" @click="likePhoto(p)">Like</button>
                    </template>

                    <template v-else>
                      <button type="button" class="btn btn-primary flex-grow-1" @click="unlikePhoto(p)">Unlike</button>
                    </template>

                    <!--Comment  -->

                    <input v-model="comment" placeholder=" Add a comment...">
                    <button type="button" class="btn btn-primary flex-grow-1" @click="commentPhoto(p)">Comment</button>



                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>


        <a href="javascript:" class="btn btn-primary" @click="uploadPhoto">Upload a photo</a>
      </div>
    </div>


  </div>
</template>

<style scoped>
.card {
  margin-bottom: 20px;
}
</style>
