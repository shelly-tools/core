<template>
  <div>
    <router-view />
  </div>
</template>

<script>
export default {
  name: 'Buildings',
  data() {
    return {
      buildings: [],
      image: '',
      name: '',
    };
  },
  created() {
    this.getAllBuildings();
  },
  methods: {
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
      this.buildings = data;
    },
  },
};
</script>
