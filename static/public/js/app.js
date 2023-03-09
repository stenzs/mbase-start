const form = document.getElementById("form");
const inputFile = document.getElementById("file");
const inputValue = document.getElementById("value");


const handleSubmit = async (event) => {
	event.preventDefault();

	const formData = new FormData();
	const error = null

	formData.append("upload", inputFile.files[0]);
	formData.append("airac", inputValue.value);
	let result = await fetch("http://localhost:3000/api/v1/task", {
		method: "post",
		body: formData,
	}).catch((error) => ("Something went wrong!", error));
	let apiResponse = await result.json();

	if (error) {
		document.getElementById('apiResponse').innerHTML = JSON.stringify(error);
	} else {
		if (apiResponse["message"] !== undefined) {
			document.getElementById('apiResponse').innerHTML = JSON.stringify(apiResponse["message"]);
		} else {
			document.getElementById('apiResponse').innerHTML = JSON.stringify(apiResponse["msg"]);
		}
	}
};

form.addEventListener("submit", handleSubmit);