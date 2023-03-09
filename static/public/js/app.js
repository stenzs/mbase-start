const form = document.getElementById("form");
const inputFile = document.getElementById("file");
const inputValue = document.getElementById("value");


const handleSubmit = (event) => {
	event.preventDefault();

	const formData = new FormData();

	formData.append("upload", inputFile.files[0]);
	formData.append("airac", inputValue.value);

	fetch("http://localhost:3000/api/v1/task", {
		method: "post",
		body: formData,
	}).catch((error) => ("Something went wrong!", error));
};

form.addEventListener("submit", handleSubmit);