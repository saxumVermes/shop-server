const apiUrl = 'http://localhost:8080';

var HttpClient = function() {
  this.get = function (url, callback) {
    var request = new XMLHttpRequest();
    request.onreadystatechange = function () {
      if (request.readyState === 4 && request.status === 200)
        callback(request.responseText);
    };

    request.open("get", url, true);
    request.setRequestHeader("Access-Control-Allow-Origin", apiUrl);
    request.send(null);
  };

  this.post = function (url, data, callback) {
    var request = new XMLHttpRequest();
    request.onreadystatechange = function () {
      if (request.readyState === 4 && request.status === 200)
        callback(request.responseText);
    };

    request.open("post", url, true);
    request.setRequestHeader("Access-Control-Allow-Origin", apiUrl);
    request.send(data);
  };

  this.delete = function (url, callback) {
    var request = new XMLHttpRequest();
    request.onreadystatechange = function () {
      if (request.readyState === 4 && request.status === 200)
        callback(request.responseText);
    };

    request.open("delete", url, true);
    request.setRequestHeader("Access-Control-Allow-Origin", apiUrl);
    request.send(null);
  }
};

var client = new HttpClient();

function submitEvent() {
  var form = document.getElementById('add-shoe');
  var data = new FormData(form);
  client.post(apiUrl + '/shoes', data, function (resp) {
    console.log(resp.toString())
  })
}

function listShoes() {
  var shoes = document.getElementById("list-shoes");
  client.get(apiUrl+'/shoes', function (resp) {
    var unserJson = JSON.parse(resp);
    for (id in unserJson) {
      var ul = document.createElement("ul");
      for (key in unserJson[id]) {
        var li = document.createElement("li");
        li.appendChild(document.createTextNode(unserJson[id][key]));
        ul.appendChild(li)
      }
      ul.appendChild(createButton("Delete", id));
      shoes.appendChild(ul);
    }
  })
}

document.addEventListener('click', function(e) {
  var s = new XMLSerializer();
  var target = e.target;
  if (target.getAttribute('name') === 'delete') {
    deleteShoe(target.getAttribute('id'));
    document.getElementById(target.getAttribute('id')).parentElement.remove()
  }
});

function deleteShoe(id) {
  client.delete(apiUrl + '/shoes/' + id, function (resp) {});
}

function createButton(value, id) {
  var button = document.createElement('button');
  button.setAttribute('type', 'button');
  button.setAttribute('name', 'delete');
  button.setAttribute('id', id);
  button.appendChild(document.createTextNode(value));
  return button
}