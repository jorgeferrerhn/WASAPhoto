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

    unbanUser: async function() {

      this.loading = true;
      this.errormsg = null;
      try {
        let url = "/users/"+this.unbanner+"/unbanUser/"+this.unbanned;
        const response = await this.$axios.delete(url);
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

        <h3 class="h3">Introduce your user ID...: </h3>
        <input v-model="unbanner" placeholder="1">

        <h3 class="h3">Introduce the ID of the user you want to unban...: </h3>
        <input v-model="unbanned" placeholder="3">



        <p>User 1 updated:  {{ user }} </p>

        <a href="javascript:" class="btn btn-primary" @click="unbanUser">Unban</a>
      </div>
    </div>


  </div>
</template>

<style scoped>
.card {
  margin-bottom: 20px;
}
</style>
