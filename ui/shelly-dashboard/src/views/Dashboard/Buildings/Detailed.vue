<template>
  <div>
    <router-link
      :to="'/buildings/show/' + item.id"
      tag="div"
      class="card"
      style="width: 18rem; cursor: pointer"
      v-for="(item, index) in buildings"
      :key="index"
    >
      <img class="card-img-top" :src="item.picturePath" alt="Card image cap" />
      <div class="card-body">
        <h5 class="card-title">{{ item.buildingName }}</h5>
        <p class="card-text">
          Some quick example text to build on the card title and make up the
          bulk of the card's content.
        </p>
      </div>
    </router-link>
    <div class="form-group">
      <label for="roomName">Name of the Building</label>
      <input
        type="text"
        class="form-control"
        name="name"
        v-model="name"
        id="name"
        aria-describedby="roomNameHelp"
        placeholder="Enter name of the room"
      />
    </div>
    <div class="form-group">
      <label for="exampleFormControlFile1">Example file input</label>
      <input
        type="file"
        class="form-control-file"
        id="file"
        name="file"
        ref="file"
      />
    </div>
    <button type="submit" class="btn btn-primary" v-on:click="createBuilding">
      Submit
    </button>
  </div>
</template>

<script>
export default {
  name: 'showBuilding',
  props: ['id'],
  data() {
    return {
      building: {},
    };
  },
  created() {
    this.getBuilding();
  },
  methods: {
    async getBuilding() {
      // runs when the component is created
      // now fetch all existing buildings
      const response = await fetch('http://localhost:8080/api/v1/buildings', {
        method: 'GET',
      });

      // transform request to json and assign it to the local buildings
      const data = await response.json();
      this.building = data;
    },
  },
};
</script>
