<script>
export default {
  data: function() {
    return {
      errormsg: null,
      loading: false,
      token: 0,
      logo: {},
      path: {} ,
      userToken:{},
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
      this.$router.push("/uploadLogo");
      this.loading = false;
    },

    getUser: async function(){
      // Getting user information
      let tokenUrl = "/users/"+this.token+"/profile";
      let userToken = await this.$axios.get(tokenUrl,{
            headers:{"Authorization": this.token}
          }
      ).then(res => res);;
      console.log(userToken);
      this.userToken = userToken.data;


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


    uploadLogo: async function() {
      this.loading = true;
      this.errormsg = null;

      // The single user of this.userToken is the logged one. Let's search for it to get its information

        try {
          let url = "/users/"+this.token+"/newlogo";
          console.log(url)

          const response = await this.$axios.post(url, this.selectedFile,{
            headers:{"Authorization": this.token}
          }).then(res => res);

          this.path = response.data['path']
          console.log(this.path)



        } catch (e) {
          this.errormsg = e.toString();
        }



      this.loading = false;
    },






  },

  computed:{
    userNotEmpty(){
      let tok = JSON.parse(JSON.stringify(this.userToken))
     if ((Object.keys(tok).length) > 0){
       return true
     }
      return false
    }
  },
  async mounted() {
    this.token = document.cookie.replace(/(?:(?:^|.*;\s*)token\s*\=\s*([^;]*).*$)|^.*$/, "$1"); // get token
    await this.getUser();
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


        <h3 class="h3 m-3">Select your new profile picture, {{userToken["Name"]}}...: </h3>
        <input type="file" @change="onChangeFileUpload">

        <!-- User information -->
      <template v-if="userNotEmpty && this.path.length" >
        <div class="col m-3 col-md-9 col-lg-7 col-xl-5" >
          <div class="card" style="border-radius: 15px;">
            <div class="card-body p-4">
              <!-- User information -->
              <div class="d-flex text-black">



                  <div class="flex-shrink-0">

                    <template v-if="userToken['ProfilePic'] != 0">
                      <img v-if="this.path" :src="this.path" v-bind:alt="Photo" class="img-fluid m-3" style="border-radius: 10px; max-width: 100%; width: 300px; height: 200px;">
                    </template>

                  </div>
                  <div class="flex-grow-1 ms-3">
                    <h5 class="mb-1">{{ userToken["Name"]}}</h5>
                    <div class="d-flex justify-content-start rounded-3 p-2 mb-2"
                         style="background-color: #efefef;">
                      <div>
                        <p class="small text-muted mb-1">Photos</p>
                        <p class="mb-0">{{ JSON.parse(userToken["Photos"]).length }}</p>
                      </div>
                      <div class="px-3">
                        <p class="small text-muted mb-1">Followers </p>
                        <p class="mb-0">{{ JSON.parse(userToken["Followers"]).length }}</p>
                      </div>
                    </div>

                  </div>



              </div>
            </div>
          </div>
        </div>
      </template>


        <a href="javascript:" class="btn btn-primary" @click="uploadLogo">Upload logo for the user</a>
      </div>
    </div>


  </div>
</template>

<style scoped>
.card {
  margin-bottom: 20px;
}
</style>
