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
function deleteBuilding(id) {
  var tpldata;
  fetch('buildings/delete_building.html',{cache: "no-store"})
    .then((response) => response.text())
    .then((template) => {
      fetch('../api/v1/buildings/get/id/' + id )
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

function createBuilding(opts) {
  fetch('../api/v1/buildings/create', {
    method: 'post',
    body: JSON.stringify(opts)
  }).then(function(response) {
    return response.json();
  }).then(function(data) {
    loadBuildings();    
  });
}
function delBuilding(opts) {
  fetch('../api/v1/buildings/delete/', {
    method: 'post',
    body: JSON.stringify(opts)
  }).then(function(response) {
    return response.json();
  }).then(function(data) {
    loadBuildings();    
  });
}


function submitBuilding() {
  var content = document.querySelector('#buildingname').value;
  if (content) {
    createBuilding({"BuildingName" : content });
  }
 }

 function submitDelBuilding(content) {
  //var content = document.querySelector('#idbuilding').value;
  if (content) {
    delBuilding({"IDBuilding" : content });
  }
 }

