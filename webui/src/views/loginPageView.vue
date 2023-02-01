<script>
export default {
	data: function() {
		return {
			errormsg: null,
			loading: false,
			users: [],
		}
	},
	methods: {
		load() {
			return load
		},

     newItem: async function() {
      console.log("Hola");
      this.loading = true;
      this.errormsg = null;
      try {
        const { data } = await this.$axios.post("/session", {
          name: "Jorge",
        });

        console.log(data);

      } catch (e) {
        this.errormsg = e.toString();
      }
      this.loading = false;
    },

    async refresh() {
      this.loading = true;
      this.errormsg = null;
      try {
        await this.$axios.post("/session", {

        });
        this.$router.push("/");
      } catch (e) {
        this.errormsg = e.toString();
      }
      this.loading = false;
    },

    /*async deleteFountain(id) {
			this.loading = true;
			this.errormsg = null;
			try {
				await this.$axios.delete("/fountains/" + id);

				await this.refresh();
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
		}

     */
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
			<h1 class="h2">LOGIN PAGE</h1>

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

        <h3 class="h3">Introduce your username: </h3>

        <input v-model="message" placeholder=" Username">
        <p>Your token is: {{ message }}</p> <!--Aquí hay que meter una variable global para el incremento del ID, o recuperarla del backend-->

				<a href="javascript:" class="btn btn-primary" @click="newItem">Create a new user</a>
			</div>
		</div>


	</div>
</template>

<style scoped>
.card {
	margin-bottom: 20px;
}
</style>
