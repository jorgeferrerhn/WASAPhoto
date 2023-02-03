<script>
export default {
  data: function() {
    return {
      errormsg: null,
      loading: false,
      token: 0,
      user: {},
    }
  },
  methods: {
    load() {
      return load
    },

    getUser: async function() {

      this.loading = true;
      this.errormsg = null;
      try {
        let url = "/users/"+this.userSearch+"/getUserProfile";
        const response = await this.$axios.get(url);
        this.user = response.data;

      } catch (e) {
        this.errormsg = e.toString();
      }
      this.loading = false;
    },

  },
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

        <h3 class="h3">Introduce user ID to search: </h3>

        <input v-model="userSearch" placeholder=" Search for a user...">
        <p>The user is: {{ user }} </p>


        <!-- User information -->
        <div class="col col-md-9 col-lg-7 col-xl-5">
          <div class="card" style="border-radius: 15px;">
            <div class="card-body p-4">
              <div class="d-flex text-black">
                <div class="flex-shrink-0">
                  <img src="https://mdbcdn.b-cdn.net/img/Photos/new-templates/bootstrap-profiles/avatar-1.webp"
                       alt="Generic placeholder image" class="img-fluid"
                       style="width: 180px; border-radius: 10px;">
                </div>
                <div class="flex-grow-1 ms-3">
                  <h5 class="mb-1">{{ user["Name"]}}</h5>
                  <div class="d-flex justify-content-start rounded-3 p-2 mb-2"
                       style="background-color: #efefef;">
                    <div>
                      <p class="small text-muted mb-1">Photos uploaded</p>
                      <p class="mb-0">{{ user["Photos"].length }}</p>
                    </div>
                    <div class="px-3">
                      <p class="small text-muted mb-1">Followers </p>
                      <p class="mb-0">{{ user["Followers"].length }}</p>
                    </div>

                  </div>
                  <div class="d-flex pt-1">

                    <button type="button" class="btn btn-primary flex-grow-1">Follow</button>
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
