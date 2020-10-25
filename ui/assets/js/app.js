function loadRooms() {
  var tpldata;
  fetch('rooms.html',{cache: "no-store"})
    .then((response) => response.text())
    .then((template) => {
      fetch('../api/v1/rooms/get/all')
        .then(res => res.json())
        .then(data => tpldata = data)
        .then(() => {
             document.getElementById('main').innerHTML = Mustache.render(template, tpldata);
        })
    });
}
function loadBuildings() {
  var tpldata;
  fetch('buildings/buildings.html',{cache: "no-store"})
    .then((response) => response.text())
    .then((template) => {
      fetch('../api/v1/buildings/get/all')
        .then(res => res.json())
        .then(data => tpldata = data)
        .then(() => {
             document.getElementById('main').innerHTML = Mustache.render(template, tpldata);
        })
    });
}

function addBuilding() {
  var tpldata;
  fetch('buildings/add_building.html',{cache: "no-store"})
    .then((response) => response.text())
    .then((template) => {
        document.getElementById('main').innerHTML = Mustache.render(template);
    });
}