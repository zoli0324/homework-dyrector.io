

var form = document.getElementById("cat-form")

form.onsubmit = function(event) {
	event.preventDefault()

	nameIn = document.getElementById("name").value
	fetch('http://localhost:8080/meow?name=' + nameIn).then(function (response) {
		// The API call was successful!
		return response.json();
	}).then(function (resp) {
		// This is the JSON from our response
		setContent(resp.data)
		displayCat(true)
	}).catch(function (err) {
		// There was an error
		setContent(null, err)
		displayCat(false)
		console.warn('Something went wrong.', err);
	});

}


function setContent(cat, error) {
	var contentDiv = document.getElementById("cat-content")
	var errDiv = document.getElementById("error")
	if (cat) {
		contentDiv.textContent=cat.art
	} else if (error) {
		console.error(error)
		contentDiv.textContent=""
		errDiv.textContent=error
	}
}

function displayCat(state) {
	var contentDiv = document.getElementById("cat-container")
	contentDiv.style.display= state ? "block" : "none"
}