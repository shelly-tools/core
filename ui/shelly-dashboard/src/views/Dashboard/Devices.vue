<template>
  <div>
    <div v-if="loadingForDiscovery !== true">
      <button
        type="button"
        class="btn btn-primary"
        v-on:click="discoverDevices"
      >
        Primary
      </button>
      {{ devices }}
      Show Unload
    </div>
    <div v-else>
      {{ devices }}
      Show loading
    </div>
  </div>
</template>

<script>
export default {
  name: 'Devices',
  data() {
    return {
      devices: [],
      loadingForDiscovery: false,
      testString: 'Test String',
    };
  },
  methods: {
    discoverDevices() {
      this.loadingForDiscovery = true;

      fetch('http://localhost:8080/api/v1/devices/discover', {
        method: 'GET',
      })
        .then((response) => response.json())
        .then((data) => {
          this.devices = data;
        })
        .then(() => {
          this.loadingForDiscovery = false;
        });
    },
  },
};
</script>
