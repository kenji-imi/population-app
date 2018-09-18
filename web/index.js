var app = new Vue({
  el: "#app",
  data: {
    results: [],
    loading: true // loadingプロパティ、初期状態はtrue
  },
  mounted() {
    axios.get("http://localhost:10080/population?limit=8").then(response => {
      // console.log("status:", response.status); // 200
      // console.log("body:", response.data); // response body.
      // console.log(response.statusText);
      // console.log(response.headers);
      // console.log(response.config);
      console.log(response.data);
      this.results = response.data;
      this.loading = false;
    });
  }
});
