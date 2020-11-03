<template>
  <div>
    <div class="row">
      <div class="form-group col-md-12">
        <label for="roomName">Search</label>
        <input
          type="text"
          class="form-control"
          name="search"
          v-model="searchValue"
          v-on:keyup.enter="search"
          id="search"
          aria-describedby="roomNameHelp"
          placeholder="Enter name of the room"
        />
      </div>
      <router-link
        :to="'/buildings/show/' + item.id"
        tag="div"
        class="card col-md-3 mr-3"
        style="width: 18rem; cursor: pointer"
        v-for="(item, index) in buildings"
        :key="index"
      >
        <img class="card-img-top" :src="item.picturePath" alt="Card image cap" />
        <div class="card-body">
          <h5 class="card-title">{{ item.name }}</h5>
          <p class="card-text">
            Some quick example text to build on the card title and make up the
            bulk of the card's content.
          </p>
        </div>
      </router-link>
    </div>
      <p v-on:click="filterBuildings">Hallo</p>
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
      buildingsAvailable: [],
      buildings: [],
      name: '',
      image: '',
      searchValue: '',
    };
  },
  async created() {
    await this.getBuildings();
    this.filterBuildings();
  },
  methods: {
    search() {
      if (this.searchValue !== '') {
        let filterArray = [];
        if (this.searchValue.includes(';')) {
          filterArray = this.searchValue.split(';');
        }

        if (filterArray.length === 0) {
          filterArray.push(this.searchValue);
        }

        for (let i = 0; i < filterArray.length; i += 1) {
          const split = filterArray[i].trim().split(':');
          let key = '';
          let value = '';
          if (split.length === 2) {
            key = split[0];
            value = split[1];
          } else {
            key = 'name';
            value = split[0];
          }

          this.buildings = [];
          for (let j = 0; j < this.buildingsAvailable.length; j += 1) {
            const objValue = String(this.buildingsAvailable[j][key]);
            console.log('Object', objValue);
            const matching = new RegExp(value);
            if (objValue.match(matching)) {
              this.buildings.push(this.buildingsAvailable[j]);
            }
          }
        }
      } else {
        this.buildings = this.buildingsAvailable;
      }
    },
    getBase64(file) {
      return new Promise((resolve, reject) => {
        const reader = new FileReader();
        reader.readAsDataURL(file);
        reader.onload = () => resolve(reader.result);
        reader.onerror = (error) => reject(error);
      });
    },
    async createBuilding() {
      // Creates a new Building
      if (typeof this.$refs.file.files[0] !== 'undefined') {
        this.image = this.$refs.file.files[0];
      }
      console.log(this.image);
      const jsonBody = {
        pictureData:
          this.image !== ''
            ? String(await this.getBase64(this.image)).substring(23)
            : '',
        picturePath: this.image !== '' ? this.image.name : '',
        name: this.name,
      };

      console.log('JSON', JSON.stringify(jsonBody));

      await fetch('http://localhost:8080/api/v1/buildings/create', {
        method: 'POST',
        body: JSON.stringify(jsonBody),
      });
    },
    filterBuildings() {
      // filter is set. filter available buildings for the criteria
      if (Array.isArray(this.$route.query.filter)) {
        this.buildings = [];
        for (let i = 0; i < this.$route.query.filter.length; i += 1) {
          const split = this.$route.query.filter[i].split(':');
          const key = split[0];
          const value = split[1];
          for (let j = 0; j < this.buildingsAvailable.length; j += 1) {
            const objValue = String(this.buildingsAvailable[j][key]);
            if (objValue === value) {
              this.buildings.push(this.buildingsAvailable[j]);
            }
          }
        }
      } else if (typeof this.$route.query.filter !== 'undefined') {
        this.buildings = [];
        const split = this.$route.query.filter.split(':');
        const key = split[0];
        const value = split[1];
        for (let j = 0; j < this.buildingsAvailable.length; j += 1) {
          const objValue = String(this.buildingsAvailable[j][key]);
          if (objValue === value) {
            this.buildings.push(this.buildingsAvailable[j]);
          }
        }
      } else {
        this.buildings = this.buildingsAvailable;
      }
    },
    async getBuildings() {
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
      this.buildingsAvailable = data;
    },
  },
};
</script>
