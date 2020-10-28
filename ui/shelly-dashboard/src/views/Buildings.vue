<template>
  <div>
    <div
      class="card"
      style="width: 18rem"
      v-for="(item, index) in buildings"
      :key="index"
    >
      <img class="card-img-top" src="#" alt="Card image cap" />
      <div class="card-body">
        <h5 class="card-title">{{ item.buildingName }}</h5>
        <p class="card-text">
          Some quick example text to build on the card title and make up the
          bulk of the card's content.
        </p>
        <a href="#" class="btn btn-primary">Go somewhere</a>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Buildings',
  data() {
    return {
      buildings: [],
    };
  },
  created() {
    this.getAllBuildings();
  },
  methods: {
    async getAllBuildings() {
      // runs when the component is created
      // now fetch all existing buildings
      const response = await fetch(
        'http://localhost:8080/api/v1/buildings/get/all',
        {
          method: 'GET',
        },
      );

      // transform request to json and assign it to the local buildings
      const data = await response.json();
      console.log(data);
      this.buildings = data;
    },
  },
};
</script>
