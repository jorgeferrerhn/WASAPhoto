<script>
export default {
  data: function() {
    return {
      errormsg: null,
      loading: false,
      token: 0,
      userStream: {},
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
        let url = "/users/"+this.userSearchStream+"/getMyStream";
        const response = await this.$axios.get(url,{
          headers:{"Authorization": this.token}});
        console.log(response)

        this.userStream = response.data;

      } catch (e) {
        this.errormsg = e.toString();
      }
      this.loading = false;
    },

  },
  mounted() {
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

        <h3 class="h3">Introduce user ID to search for its stream...: </h3>

        <input v-model="userSearchStream" placeholder=" Search for a user...">
        <p>The stream of the user is: {{ userStream }} </p>

        <a href="javascript:" class="btn btn-primary" @click="getUserStream">Search for user stream</a>
      </div>
    </div>


  </div>
</template>

<style scoped>
.card {
  margin-bottom: 20px;
}
</style>
