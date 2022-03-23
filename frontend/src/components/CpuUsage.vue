<template>
  <v-container>
    <v-card
      class="mx-auto"
      color="grey lighten-4"
      elevation="2"
      max-width="600"
    >
      <v-card-title>
        <v-icon
          :color="cpuList.slice(-1)[0] > 10 ? 'red' : 'green'"
          class="mr-12"
          size="64"
          align="right"
        >
          mdi-heart-pulse
        </v-icon>
        <v-row align="start">
          <div>
            <span class="text-h3">CPU</span>
            <strong> {{ cpuList.slice(-1)[0] }}% </strong>
          </div>
        </v-row>
        <v-spacer></v-spacer>

        <v-btn icon class="align-self-start" size="28">
          <v-icon>mdi-arrow-right-thick</v-icon>
        </v-btn>
      </v-card-title>

      <v-sheet color="transparent">
        <v-sparkline
          :smooth="16"
          :gradient="['#f72047', '#ffd200', '#1feaea']"
          :line-width="3"
          :value="cpuList"
          stroke-linecap="round"
        ></v-sparkline>
        <v-divider />
        <div class="subheading red--text">{{ error }}</div>
      </v-sheet>
    </v-card>
  </v-container>
</template>

<script>
const basePort = process.env.VUE_APP_BACKEND_PORT;
const baseURL = basePort
  ? `http://localhost:${basePort}/`
  : process.env.BASE_URL;

const DELAY = 2000;
const MAX_ITEMS = 100;

export default {
  name: "CPUUsage",

  data: () => ({
    gradient: ["red", "orange", "yellow", "green"],
    cpuList: [],
    error: null,
  }),

  async created() {
    setInterval(this.fetchCPUData, DELAY);
  },

  methods: {
    async fetchCPUData() {
      this.error = null;
      fetch(baseURL + "api/stats/cpu?count=" + MAX_ITEMS)
        .then(async (res) => {
          if (!res.ok) {
            this.error = `Error ${res.status}: ${res.statusText}`;
            return;
          }

          const data = await res.json();
          console.log(data);
          this.cpuList = data;
        })
        .catch((err) => {
          this.error = err;
        });
    },
  },
};
</script>
