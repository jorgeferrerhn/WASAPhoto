<script>
export default {
  data: function() {
    return {
      errormsg: null,
      previewImage:null,
      loading: false,
      token: 0,
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





    onChangeFileUpload: async function (e) {
      let file = e.target.files[0]
      this.serverFile = file
      await this.encodeImage(file)
    },
    encodeImage: async function(input) {
      if (input) {
        let reader = new FileReader()
        reader.onload = (e) => {
          this.selectedFile = e.target.result
        }
        reader.readAsDataURL(input)
      }
    },

    uploadPhoto: async function() {

      this.loading = true;
      this.errormsg = null;

      try {


        // Let's get the cookie
        this.token = document.cookie.replace(/(?:(?:^|.*;\s*)token\s*\=\s*([^;]*).*$)|^.*$/, "$1");
        let url = "/users/"+this.token+"/photo";
        const response = await this.$axios.post(url,this.selectedFile,{
          headers:{"Authorization": this.token}
        }).then(res => res);


        let photo = response.data

        // Now, we have to add it to the photo's list

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
        <input type="file" @change="onChangeFileUpload">
        <a href="javascript:" class="btn btn-primary" @click="uploadPhoto">Upload a photo</a>


        <!-- Photo information -->
        <template v-if="photos.length > 0">
          <div class="m-3" v-for="(p,index) in photos" :key="index">
            <h5>Photo uploaded!</h5>
            <div class="card m-3" style="border-radius: 15px;">
              <div class="d-flex p-2 mb-2" style="border-radius: 15px;background-color: #efefef;">
                <div class="m-3">
                  <img v-if="p['path']" :src="p['path']" v-bind:alt="Photo" class="img-fluid m-3" style="border-radius: 10px; max-width: 100%; width: 300px; height: 200px;">


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
