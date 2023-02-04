<script>
export default {
  data: function() {
    return {
      errormsg: null,
      loading: false,
      token: 0,
      photo: {},
    }
  },
  methods: {
    load() {
      return load
    },

    async refresh() {
      this.loading = true;
      this.errormsg = null;
      this.$router.push("/getImage");
      this.loading = false;
    },


    getImage: async function() {

      this.loading = true;
      this.errormsg = null;
      try {
        let url = "/images/"+this.id;
        const response = await this.$axios.get(url);
        console.log(response)

        this.photo = response.data;

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

        <h3 class="h3">Introduce image identifier...: </h3>

        <input v-model="id" placeholder=" Search for an image...">
        <p>The image is: {{ photo }} </p>

        <a href="javascript:" class="btn btn-primary" @click="getImage">Search for the image</a>
      </div>
    </div>


  </div>
</template>

<style scoped>
.card {
  margin-bottom: 20px;
}
</style>
