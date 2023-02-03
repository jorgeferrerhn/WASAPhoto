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
      this.path= "/home/jorge/WASAPhoto/webui/src/images/"+name; // update the path here
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
        let getToken = document.cookie.replace(/(?:(?:^|.*;\s*)token\s*\=\s*([^;]*).*$)|^.*$/, "$1");
        console.log(this.path)
        let url = "/users/"+this.id+"/uploadPhoto";
        const response = await this.$axios.post(url,this.path,{
          headers:{'Authorization': getToken}
        });
        this.photo = response.data;


        let contains = false
        for (let i = 0; i < this.photos.length; i++){
          if (this.photos[i]["id"] == this.photo["id"]){
            contains = true
          } // contains
        }
        if (!contains){
          this.photos.push(this.photo);
        }

      } catch (e) {
        this.errormsg = e.toString();
      }
      this.loading = false;
    },

  }
  /*
  mounted() {
    this.refresh()
  }*/
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

        <h3 class="h3">Introduce user id...: </h3>
        <input v-model="id" placeholder="1">

        <h3 class="h3">Introduce photo path...: </h3>

        <input type="file" @change="onFileChange"/>
        <!--v-file-input placeholder="Upload document" v-model="file" accept="image/*" label="Image" @change="onFileChange"/-->

        <p>Photo uploaded:  {{ photo }} </p>

        <!-- Photo information -->
        <div class="col col-md-9 col-lg-7 col-xl-5" v-for="p in photos" v-if="photos.length > 0">
          <div class="card" style="border-radius: 15px;">
            <div class="card-body p-4">
              <div class="d-flex text-black">
                <div class="flex-shrink-0">


                </div>

                <div>
                  <div class="imagePreviewWrapper" style="width: 200px; length: 200px; border-radius: 10px;" @click="selectImage"> </div>

                  <input ref="fileInput" type="file" @input="pickFile">
                </div>


                <div class="flex-grow-1 ms-3">
                  <h5 class="mb-1">{{ p["path"]}}</h5>
                  <div class="d-flex justify-content-start rounded-3 p-2 mb-2"
                       style="background-color: #efefef;">
                    <div>
                      <p class="small text-muted mb-1">Likes</p>
                      <p class="mb-0">{{ JSON.parse(p["likes"]).length }}</p>
                    </div>
                    <div class="px-3">
                      <p class="small text-muted mb-1">Comments </p>
                      <p class="mb-0">{{ JSON.parse(p["comments"]).length}}</p>
                    </div>

                    <div class="px-3">
                      <p class="small text-muted mb-1">Date Added </p>
                      <p class="mb-0">{{ p["date"]}}</p>
                    </div>

                  </div>
                  <div class="d-flex pt-1">

                    <button type="button" class="btn btn-primary flex-grow-1" @click="likePhoto">Like</button>
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
