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
      photos:[],
      selectedFile:null
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

    onUpload: async function() {
      this.loading = true;
      this.errormsg = null;
      try {
        let formData = new FormData();
        formData.append('file', this.imageData);
        this.token = document.cookie.replace(/(?:(?:^|.*;\s*)token\s*\=\s*([^;]*).*$)|^.*$/, "$1");
        console.log(this.path)
        let url = "/users/" + this.token + "/photo";
        const response = await this.$axios.post(url, this.path, {
          headers: {"Authorization": this.token}
        });

      } catch (e) {
        this.errormsg = e.toString();
      }
      this.loading = false;

    },


    onFileSelected(event){
      this.selectedFile = event.target.files[0]

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

          // cast the selected file to formData
          const fd = new FormData();
          fd.append('image',this.selectedFile,this.selectedFile.name)


          // Let's get the cookie
          this.token = document.cookie.replace(/(?:(?:^|.*;\s*)token\s*\=\s*([^;]*).*$)|^.*$/, "$1");
          let url = "/users/"+this.token+"/photo";
          const response = await this.$axios.post(url,fd,{
            headers:{"Authorization": this.token}
          }).then(res => res);

          console.log(response)
          /*
          let contains = false
          for (let i = 0; i < this.photos.length; i++){
            if (this.photos[i]["id"] == photo["id"]){
              contains = true
            }
          }


          if (!contains){
            this.photos.push(photo);
          }
          */

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
        <v-file-input accept="image/png, image/jpeg, image/bmp" placeholder="Pick a photo" prepend-icon="mdi-camera" v-model="imageData">Select image...</v-file-input>

        <input type="file" @change="onFileSelected">

        <a href="javascript:" class="btn btn-primary" @click="uploadPhoto">Upload a photo</a>
        <!--v-file-input type="file" @change="onFileChange"/-->
        <!--v-file-input placeholder="Upload document" v-model="file" accept="image/*" label="Image" @change="onFileChange"/-->


        <!-- Photo information -->
        <template v-if="photos.length > 0">
          <div class="m-3" v-for="(p,index) in photos" :key="index">
            <div class="card m-3" style="border-radius: 15px;">
              <div class="d-flex p-2 mb-2" style="border-radius: 15px;background-color: #efefef;">
                <div class="m-3">
                  <img :src="'/img/'+p['path']" v-bind:alt="p['path']" class="img-fluid m-3"
                       style="border-radius: 10px; max-width: 100%; width: 300px; height: 200px;">
                </div>

                <div class="m-3">
                  <h5 class="mb-1 m-3 p-3">Photo comments: </h5>

                  <div v-for="(c,index) in JSON.parse(p['comments'])" :key="index" >
                    <p class="">{{c['UserId']}} : {{c['Content']}}</p>
                  </div>

                </div>

              </div>

              <div class="d-flex align-items-center">

                <div class="m-3 d-flex align-items-center">
                  <div class="m-3 d-flex align-items-center">
                    <h6 class="m-3 ">Photo ID: {{ p["id"]}}</h6>
                    <p class="m-3">Likes: {{ JSON.parse(p["likes"]).length }}</p>
                    <p class="m-3">Comments: {{ JSON.parse(p["comments"]).length }}</p>
                  </div>
                </div>
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
