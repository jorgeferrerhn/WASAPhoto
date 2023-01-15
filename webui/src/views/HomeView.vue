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
		async refresh() {
			this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.get("/fountains/");
				this.users = response.data;
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
		},
		async newItem() {
			this.$router.push("/session");
		},
		async deleteFountain(id) {
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
			<h1 class="h2">Users list</h1>
			<div class="btn-toolbar mb-2 mb-md-0">
				<div class="btn-group me-2">
					<button type="button" class="btn btn-sm btn-outline-secondary" @click="refresh">
						Refresh
					</button>
				</div>
				<div class="btn-group me-2">
					<button type="button" class="btn btn-sm btn-outline-primary" @click="newItem">
						New
					</button>
				</div>
			</div>
		</div>

		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>

		<LoadingSpinner v-if="loading"></LoadingSpinner>

		<div class="card" v-if="users.length === 0">
			<div class="card-body">
				<p>No users in the database.</p>

				<a href="javascript:" class="btn btn-primary" @click="newItem">Create a new user</a>
			</div>
		</div>

		<div class="card" v-if="!loading" v-for="u in users">
			<div class="card-header">
				Fountain
			</div>
			<div class="card-body">
				<p class="card-text">
					Name: u.name<br />
					Longitude: {{ u.longitude }}<br />
					Status: {{ u.status }}
				</p>
				<a href="javascript:" class="btn btn-danger" @click="deleteFountain(u.id)">Delete</a>
			</div>
		</div>
	</div>
</template>

<style scoped>
.card {
	margin-bottom: 20px;
}
</style>
